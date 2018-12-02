package main

import (
	"context"
	"flag"
	"log"
	"os"
	"os/signal"
	"runtime"
	"time"

	"k8s.io/client-go/informers"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/cache"
	"k8s.io/client-go/tools/clientcmd"
)

func newClientset() (*kubernetes.Clientset, error) {
	rules := clientcmd.NewDefaultClientConfigLoadingRules()
	config, err := clientcmd.BuildConfigFromKubeconfigGetter("", rules.Load)

	if err != nil {
		return nil, err
	}

	return kubernetes.NewForConfig(config)
}

func leak(ctx context.Context, informer cache.SharedIndexInformer) {
	ctx, cancel := context.WithTimeout(ctx, 1*time.Minute)
	defer cancel()

	for {
		select {
		case <-ctx.Done():
			break
		default:
		}

		// Each call to WaitForCacheSync will leak a goroutine until the context is closed.
		// This is not expected behavior, the stopChannel should only be needed to abort while
		// waiting for the sync.
		// 这里有两种情况 goroutine 会释放
		// 1. 当 cache sync 完成
		// 2. stopCh <-chan struct{} 被 close 掉，但是传递过去的 stopCh <-chan struct{} 可能别人也在
		// 使用或者它是一个长时间运行的上下文，你不能在 WaitForCacheSync 中 close 掉它。
		// 详情见: https://github.com/kubernetes/kubernetes/issues/71277
		// 解决方案是在内部创建一个 cancelable context，当 WaitFor() 执行成功，close 掉这个 context
		cache.WaitForCacheSync(ctx.Done(), informer.HasSynced)
		time.Sleep(200 * time.Millisecond)
	}
}

func statusPrinter(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			break
		default:
		}
		log.Printf("goroutines: %d\n", runtime.NumGoroutine())
		time.Sleep(250 * time.Millisecond)
	}
}

func main() {
	flag.Parse()

	cs, err := newClientset()
	if err != nil {
		log.Fatalf("connecting to cluster: %v", err)
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	factory := informers.NewSharedInformerFactory(cs, 3*time.Minute)
	podInformer := factory.Core().V1().Pods()
	// Must add event handler or WaitForCacheSync will never complete
	podInformer.Informer().AddEventHandler(cache.ResourceEventHandlerFuncs{})

	log.Printf("starting")
	factory.Start(ctx.Done())

	go leak(ctx, podInformer.Informer())
	go statusPrinter(ctx)

	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, os.Interrupt, os.Kill)
	go func() {
		select {
		case <-sigCh:
			cancel()
		case <-ctx.Done():
		}
		log.Printf("stopping")
	}()

	<-ctx.Done()
}
