package main

type MyHashSet struct {
	bucket []int
}

func Constructor() MyHashSet {
	bucket := make([]int, 1000001)

	for i := range bucket {
		bucket[i] = -1
	}

	return MyHashSet{bucket: bucket}
}

func (this *MyHashSet) Add(key int) {
	if !this.Contains(key) {
		this.bucket = append(this.bucket, key)
	}
}

func (this *MyHashSet) Remove(key int) {
	for i, v := range this.bucket {
		if v == key {
			this.bucket = append(this.bucket[:i], this.bucket[i+1:]...)
			return
		}
	}
}

func (this *MyHashSet) Contains(key int) bool {
	for _, v := range this.bucket {
		if v == key {
			return true
		}
	}
	return false
}
