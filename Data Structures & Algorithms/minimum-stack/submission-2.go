type ValMin struct{
    val, min int
}

type MinStack struct {
    n []ValMin
    top int    
}

func Constructor() MinStack {
    return MinStack{
        n : []ValMin{},        
        top: -1,
    }
}

func (this *MinStack) Push(val int) {
    this.top++
    if this.top == 0 || val < this.n[this.top-1].min {
        this.n = append(this.n, ValMin{val,val})        
    } else {
        this.n = append(this.n, ValMin{val,this.n[this.top-1].min})
    }       
}

func (this *MinStack) Pop() { 
    if this.top < 0 {
        return
    }
    this.n = this.n[:this.top]
    this.top--     
}

func (this *MinStack) Top() int {
    return this.n[this.top].val
}

func (this *MinStack) GetMin() int {
    return this.n[this.top].min
}
