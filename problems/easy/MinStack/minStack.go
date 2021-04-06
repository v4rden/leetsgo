package MinStack

type MinStack struct {
	Elements, mins []int
}

func Constructor() MinStack {
	return MinStack{Elements: make([]int, 0), mins: make([]int, 0)}
}

func (this *MinStack) Push(val int) {
	this.Elements = append(this.Elements, val)

	if len(this.mins) > 0 && this.GetMin() < val {
		this.mins = append(this.mins, this.GetMin())
	} else {
		this.mins = append(this.mins, val)
	}
}

func (this *MinStack) Pop() {
	this.Elements = this.Elements[:len(this.Elements)-1]
	this.mins = this.mins[:len(this.mins)-1]
}

func (this *MinStack) Top() int {
	return this.Elements[len(this.Elements)-1]
}

func (this *MinStack) GetMin() int {
	return this.mins[len(this.mins)-1]
}
