type MinStack struct {
    top int
    items []int
    minTop int
    mins []int
}

func Constructor() MinStack {
    return MinStack{}
}

func (this *MinStack) Push(val int) {
    this.items = append(this.items, val)
    this.top++

    if len(this.mins) == 0 || val <= this.mins[this.minTop - 1]{
        this.mins = append(this.mins, val)
        this.minTop++        
    }
}

func (this *MinStack) Pop() {
    if len(this.items) == 0{
        return
    }
    out := this.items[len(this.items)-1]
    this.items = this.items[:len(this.items)-1]
    this.top--

    if out == this.mins[this.minTop - 1]{
        this.mins = this.mins[:len(this.mins) - 1]
        this.minTop--
    }
}

func (this *MinStack) Top() int {
    return this.items[this.top-1]
}

func (this *MinStack) GetMin() int {
    return this.mins[this.minTop-1]
}
