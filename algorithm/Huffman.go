package main

import (
	"fmt"
)

// HuffmanNode 定义霍夫曼树中的节点
type HuffmanNode struct {
	Weight int
	Value  int
	Left   *HuffmanNode
	Right  *HuffmanNode
}

// implement the heap.Interface for a HuffmanHeap
type HuffmanHeap []*HuffmanNode

func (h HuffmanHeap) Len() int            { return len(h) }
func (h HuffmanHeap) Less(i, j int) bool  { return h[i].Weight < h[j].Weight }
func (h HuffmanHeap) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *HuffmanHeap) Push(x interface{}) { *h = append(*h, x.(*HuffmanNode)) }
func (h *HuffmanHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

// BuildHuffmanTree 构建霍夫曼树
func BuildHuffmanTree(weights []int) *HuffmanNode {
	heap := &HuffmanHeap{}
	heap.Init()

	// 将权重转换为HuffmanNode并添加到堆中
	for _, weight := range weights {
		heap.Push(&HuffmanNode{Weight: weight, Value: weight})
	}

	// 构建霍夫曼树
	for heap.Len() > 1 {
		left := heap.Pop().(*HuffmanNode)
		right := heap.Pop().(*HuffmanNode)

		newNode := &HuffmanNode{
			Weight: left.Weight + right.Weight,
			Left:   left,
			Right:  right,
		}

		heap.Push(newNode)
	}

	return heap.Pop().(*HuffmanNode)
}

// calculateWPL 计算霍夫曼树的带权外部路径长度
func calculateWPL(node *HuffmanNode, depth int) int {
	if node.Left == nil && node.Right == nil {
		return node.Weight * depth
	}
	return calculateWPL(node.Left, depth+1) + calculateWPL(node.Right, depth+1)
}

func main() {
	weights := []int{2, 4, 5, 9}
	root := BuildHuffmanTree(weights)
	wpl := calculateWPL(root, 1) // 从1开始计数
	fmt.Printf("The weighted path length of the Huffman tree is: %d\n", wpl)
}
