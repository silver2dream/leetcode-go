package s232

type MyQueue struct {
	arr []int
}

func Constructor() MyQueue {
	return MyQueue{}
}

func (this *MyQueue) Push(x int) {
	this.arr = append(this.arr, x)
}

func (this *MyQueue) Pop() int {
	pop := this.arr[0]
	this.arr = this.arr[1:len(this.arr)]
	return pop
}

func (this *MyQueue) Peek() int {
	peak := this.arr[0]
	return peak
}

func (this *MyQueue) Empty() bool {
	return len(this.arr) == 0
}

/**
 * Your MyQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Peek();
 * param_4 := obj.Empty();
 */
