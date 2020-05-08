type MyStack struct {
    Queue []int
}


/** Initialize your data structure here. */
func Constructor() MyStack {
    return MyStack{[]int{}}
}


/** Push element x onto stack. */
func (this *MyStack) Push(x int)  {
    this.Queue = append(this.Queue, x)
}


/** Removes the element on top of the stack and returns that element. */
func (this *MyStack) Pop() int {
    for i := 0; i < len(this.Queue)-1; i++ {
        elem := this.Queue[0]
        this.Queue = this.Queue[1:]
        this.Queue = append(this.Queue, elem)
    }
    top := this.Queue[0]
    this.Queue = this.Queue[1:]
    return top
}


/** Get the top element. */
func (this *MyStack) Top() int {
    for i := 0; i < len(this.Queue)-1; i++ {
        elem := this.Queue[0]
        this.Queue = this.Queue[1:]
        this.Queue = append(this.Queue, elem)
    }
    top := this.Queue[0]
    this.Queue = this.Queue[1:]
    this.Queue = append(this.Queue, top)
    return top
}


/** Returns whether the stack is empty. */
func (this *MyStack) Empty() bool {
    return len(this.Queue) == 0
}


/**
 * Your MyStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.Empty();
 */
