package main

type CQueue struct {
	in, out []int
}

func Constructor() CQueue {
	return CQueue{}
}

func (this *CQueue) AppendTail(value int) {
	this.in = append(this.in, value)
}

func (this *CQueue) DeleteHead() int {
	if len(this.out) == 0 {
		for len(this.in) > 0 {
			this.out = append(this.out, this.in[len(this.in)-1])
			this.in = this.in[:len(this.in)-1]
		}
	}
	if len(this.out) == 0 {
		return -1
	}
	ans := this.out[len(this.out)-1]
	this.out = this.out[:len(this.out)-1]
	return ans
}
