package of0009

type CQueue struct {
	input  []int
	output []int
}

func Constructor() CQueue {
	return CQueue{}
}

func (this *CQueue) AppendTail(value int) {
	this.input = append(this.input, value)
}

func (this *CQueue) DeleteHead() int {
	if len(this.input) == 0 && len(this.output) == 0 {
		return -1
	}
	if len(this.output) == 0 {
		for len(this.input) > 0 {
			this.output = append(this.output, this.input[len(this.input)-1])
			this.input = this.input[:len(this.input)-1]
		}
	}
	value := this.output[len(this.output)-1]
	this.output = this.output[:len(this.output)-1]
	return value
}

/**
 * Your CQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AppendTail(value);
 * param_2 := obj.DeleteHead();
 */
