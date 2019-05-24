package queuearray

const size = 100

type Queue struct{
	top int
	rear int
	items [size]interface{}
}

func New(size int)*Queue{
	return &Queue{
	}
}

func (s *Queue) Push(item interface{})bool{
	if s.top==len(s.items){
		return false
	}
	s.items[s.top]=item
	s.top +=1
	return true
}

func (s *Queue) Pop()(interface{},bool){
	if s.top==s.rear{
		return nil,false
	}
	item := s.items[s.rear]
	s.rear +=1
	return item,true
}