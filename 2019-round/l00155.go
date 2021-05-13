package main

type MinStack struct {
	Num *Node
	Min *Node
}

type Node struct {
	Val int
	Next *Node
}


/** initialize your data structure here. */
func Constructor() MinStack {
	return MinStack{Num:nil,Min:nil}
}


func (this *MinStack) Push(x int)  {
	newNode := &Node{Val:x,Next:(*this).Num}
	this.Num = newNode
	if this.Min ==nil || x <= this.GetMin() {
		newNode := &Node{Val:x,Next:(*this).Min}
		this.Min = newNode
	}
}


func (this *MinStack) Pop()  {
	if this.Num == nil {
		panic("the stack is empty")
	}
	x:= this.Num.Val
	if this.GetMin() == x {
		this.Min = this.Min.Next
	}
	this.Num = this.Num.Next
}


func (this *MinStack) Top() int {
	if this.Num == nil {
		panic("the stack is empty")
	}
	return this.Num.Val
}


func (this *MinStack) GetMin() int {
	if this.Num == nil {
		panic("the stack is empty")
	}
	return this.Min.Val
}


func (this *MinStack) PrintStack() {
	print("Num is:  ")
	for a := this.Num;a!=nil;a=a.Next{
		print(a.Val," ")
	}
	println()
	print("Min is:  ")
	for a := this.Min;a!=nil;a=a.Next{
		print(a.Val," ")
	}
	println()	
}

func main() {
	aa := Constructor()
	aa.Push(2147483646)
	aa.Push(2147483646)
	aa.Push(2147483647)
	aa.PrintStack()
	println(aa.Top())
	aa.Pop()
	println(aa.GetMin())
	aa.Pop()
	println(aa.GetMin())
	aa.PrintStack()
	aa.Pop()
	aa.Push(2147483647)
	println(aa.Top())
	println(aa.GetMin())
	aa.Push(-2147483648)
	println(aa.Top())
	println(aa.GetMin())
	aa.Pop()
	println(aa.GetMin())
}
