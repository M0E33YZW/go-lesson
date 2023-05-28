package main

import (
	"container/heap"
	"fmt"
)

// 優先度付きキュー内の要素を表す構造体
type Queue struct {
	value string // 要素
	prio  int    // 優先度
}

// 優先度付きキューを表す型
type PrioQueue []*Queue

// PrioQueueのコンストラクタ
func NewPrioQueue() PrioQueue {
	return make(PrioQueue, 0)
}

// enqueue
func (pq *PrioQueue) enqueue(s string, prio int) {
	item := &Queue{
		value: s,
		prio:  prio,
	}
	*pq = append(*pq, item)
	heap.Fix(pq, len(*pq)-1)
}

// dequeue
func (pq *PrioQueue) dequeue() string {
	n := len(*pq)
	item := (*pq)[0]
	*pq = (*pq)[1:n]
	return item.value
}

// Lenは要素数を返すメソッド
func (pq PrioQueue) Len() int {
	return len(pq)
}

// Lessは比較関数
func (pq PrioQueue) Less(i, j int) bool {
	// 優先度の高い順にソートするため、prioの値が小さいものを優先する
	return pq[i].prio < pq[j].prio
}

// Swapは要素を交換するメソッド
func (pq PrioQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

// Popは要素を取り出すメソッド
func (pq *PrioQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	*pq = old[0 : n-1]
	return item
}

// Pushは要素を追加するメソッド
func (pq *PrioQueue) Push(x interface{}) {
	item := x.(*Queue)
	*pq = append(*pq, item)
}

func prioSched() []string {
	prioQueue := NewPrioQueue()
	prioQueue.enqueue("A", 1)
	prioQueue.enqueue("B", 2)
	prioQueue.enqueue("C", 2)
	prioQueue.enqueue("D", 3)
	prioQueue.dequeue()
	prioQueue.dequeue()
	prioQueue.enqueue("D", 3)
	prioQueue.enqueue("B", 2)
	prioQueue.dequeue()
	prioQueue.dequeue()
	prioQueue.enqueue("C", 2)
	prioQueue.enqueue("A", 1)
	var result []string
	for len(prioQueue) > 0 {
		item := prioQueue.dequeue()
		result = append(result, item)
	}
	return result
}

func main() {
	fmt.Println(prioSched())
}
