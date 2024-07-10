package main

import (
	"fmt"
)
type Container interface{
	Push(x interface{})
	Pop() interface{}
	Size() int
	ShowCont() [] interface{}
}
type Stack struct{
	DB [] interface{}
}
func (s *Stack) Push(x interface{}){
	s.DB=append(s.DB,x)
}
func(s *Stack) Pop() interface{}{
	if s.Size() == 0{
		return nil
	}
	x := s.DB[len(s.DB)-1]
	s.DB = s.DB[:len(s.DB)-1]
	return x
}
func (s *Stack) Size()int{
	return len(s.DB)
}
func (s *Stack) ShowCont(){
	return s.DB
}
func NewStack(DB [] interface) *Stack{
	return&Stack{
		DB [] interface: DB [] interface
	}
}
func main(){
	var mas [] int
	var i int
	var action string
	fmt.Println("Enter is Stack :")
	for{
		fmt.Scan(&i)
		if i == nil{
			return	
		}
		
		mas = append(mas, i)
	}
	StackMas:=NewStack(mas)
	fmt.Println("Select the action you want to do with the Stack ")
	for{
		switch action {
		case "push":
			
			
		}
	}
}
