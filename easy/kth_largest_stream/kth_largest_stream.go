package main

import (
	"fmt"
	"math"
	"sort"
)

type KthLargest struct {
	k                 int
	nums              []int
	currentKthLargest int
}

func (kl *KthLargest) filterKElements() {
	sort.Ints(kl.nums)
	if len(kl.nums) >= kl.k {
		kl.nums = kl.nums[len(kl.nums)-kl.k:]
		kl.currentKthLargest = kl.nums[len(kl.nums)-kl.k]
	}
}

func Constructor(k int, nums []int) KthLargest {
	kthLargest := KthLargest{
		k:                 k,
		nums:              nums,
		currentKthLargest: math.MinInt64,
	}
	kthLargest.filterKElements()
	return kthLargest
}

func (kl *KthLargest) Add(val int) int {
	if val >= kl.currentKthLargest {
		kl.nums = append(kl.nums, val)
		kl.filterKElements()
	}

	return kl.currentKthLargest
}

func main() {
	k := 3
	arr := []int{4, 5, 8, 2}
	kthLargest := Constructor(k, arr)

	fmt.Println("+++ SOLUTION 1 +++")
	fmt.Println(kthLargest.Add(3))
	fmt.Println(kthLargest.Add(5))
	fmt.Println(kthLargest.Add(10))
	fmt.Println(kthLargest.Add(9))
	fmt.Println(kthLargest.Add(4))

	k = 1
	arr = []int{}
	kthLargest = Constructor(k, arr)

	fmt.Println("+++ SOLUTION 2 +++")
	fmt.Println(kthLargest.Add(-3))
	fmt.Println(kthLargest.Add(-2))
	fmt.Println(kthLargest.Add(-4))
	fmt.Println(kthLargest.Add(0))
	fmt.Println(kthLargest.Add(4))
}
