package gosort

import "fmt"

type Node struct {
	value int
	next  *Node
}

func BucketSort(nums []int, bucketSize int) {
	if len(nums) == 0 {
		return
	}
	minValue, maxValue := MinAndMaxValue(nums)
	buckets := make([]*Node, (maxValue-minValue)/bucketSize+1)

	for _, num := range nums {
		bucket_num := (num - minValue) / bucketSize
		node := buckets[bucket_num]
		if buckets[bucket_num] == nil {
			buckets[bucket_num] = &Node{
				value: num,
			}
			continue
		}
		if node.value > num {
			buckets[bucket_num] = &Node{
				value: num,
				next:  node,
			}
			continue
		}
		for {
			if node.next == nil {
				node.next = &Node{
					value: num,
				}
				break
			}
			if node.next.value > num {
				node.next = &Node{
					value: num,
					next:  node.next,
				}
				break
			}
			node = node.next
		}

	}

	index := 0
	for _, bucket := range buckets {
		node := bucket
		for node != nil {
			nums[index] = node.value
			index++
			node = node.next
		}
	}
	fmt.Print()
}
