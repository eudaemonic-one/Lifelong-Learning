type MyQueue struct {
    stack1 []int
    stack2 []int
}


/** Initialize your data structure here. */
func Constructor() MyQueue {
    return MyQueue{make([]int, 0), make([]int, 0)}
}


/** Push element x to the back of queue. */
func (this *MyQueue) Push(x int)  {
    this.stack1 = append(this.stack1, x)
}


/** Removes the element from in front of queue and returns that element. */
func (this *MyQueue) Pop() int {
    length := len(this.stack1)
    for i := 0; i < length; i++ {
        this.stack2 = append(this.stack2, this.stack1[len(this.stack1)-1])
        this.stack1 = this.stack1[:len(this.stack1)-1]
    }
    top := this.stack2[len(this.stack2)-1]
    this.stack2 = this.stack2[:len(this.stack2)-1]
    for i := 0; i < length-1; i++ {
        this.stack1 = append(this.stack1, this.stack2[len(this.stack2)-1])
        this.stack2 = this.stack2[:len(this.stack2)-1]
    }
    return top
}


/** Get the front element. */
func (this *MyQueue) Peek() int {
    for i := 0; i < len(this.stack1); i++ {
        this.stack2 = append(this.stack2, this.stack1[len(this.stack1)-i-1])
    }
    top := this.stack2[len(this.stack2)-1]
    this.stack2 = make([]int, 0)
    return top
}


/** Returns whether the queue is empty. */
func (this *MyQueue) Empty() bool {
    return len(this.stack1) == 0
}


/**
 * Your MyQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Peek();
 * param_4 := obj.Empty();
 */
