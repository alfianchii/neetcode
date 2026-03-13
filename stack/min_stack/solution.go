package main

type MinStack struct {
	bucket []int
}

func Constructor() MinStack {
	bucket := []int{}

	return MinStack{bucket: bucket}
}

func (this *MinStack) Push(val int) {
	this.bucket = append(this.bucket, val)
}

func (this *MinStack) Pop() {
	this.bucket = this.bucket[:len(this.bucket)-1]
}

func (this *MinStack) Top() int {
	return this.bucket[len(this.bucket)-1]
}

func (this *MinStack) GetMin() int {
	if len(this.bucket) == 0 {
		return 0
	}

	min := this.bucket[0]
	for i := 1; i < len(this.bucket); i++ {
		nextNum := this.bucket[i]

		if nextNum < min {
			min = nextNum
		}
	}

	return min
}
