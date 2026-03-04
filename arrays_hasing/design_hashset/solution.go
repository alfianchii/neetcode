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
	this.bucket[key] = key
}

func (this *MyHashSet) Remove(key int) {
	this.bucket[key] = -1
}

func (this *MyHashSet) Contains(key int) bool {
	if this.bucket[key] == -1 {
		return false
	}

	return true
}
