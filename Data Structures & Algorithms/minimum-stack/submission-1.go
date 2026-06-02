type MinStack struct {
    n []int
    top int
    mins []int    
    mintop int
}

func Constructor() MinStack {
    return MinStack{
        n : []int{},
        mins: []int{},
        top: 0,
        mintop: -1,
    }
}

func (this *MinStack) Push(val int) {
    this.n = append(this.n, val)
    this.top++
    if len(this.mins) == 0 || val <= this.mins[this.mintop] {
        this.mins = append(this.mins, val)
        this.mintop++
    }
}

func (this *MinStack) Pop() {    
    popped, remaining := this.n[this.top-1],this.n[:this.top-1]
    this.n = remaining
    this.top--   
    if this.mintop >= 0 && popped == this.mins[this.mintop]{
        this.mins = this.mins[:this.mintop]
        this.mintop--
    }
}

func (this *MinStack) Top() int {
    return this.n[this.top-1]
}

func (this *MinStack) GetMin() int {
    return this.mins[this.mintop]
}
