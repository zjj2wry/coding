package stackarray

const size = 100

type Stack struct{
	top int
	items [size]interface{}
}

func New(size int)*Stack{
	return &Stack{
	}
}

func (s *Stack) Push(item interface{})bool{
	if s.top==len(s.items){
		return false
	}
	s.items[s.top]=item
	s.top +=1
	return true
}

func (s *Stack) Pop()(interface{},bool){
	if s.top==0{
		return nil,false
	}
	item := s.items[s.top-1]
	s.top -=1
	return item,true
}

