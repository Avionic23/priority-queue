package priorityqueue

type PriorityQueue struct {
	heap     []interface{}
	last     int
	function func(elem1 interface{}, elem2 interface{}) bool
}

func NewPriorityQueue(length int, function func(elem1 interface{}, elem2 interface{}) bool) *PriorityQueue {
	return &PriorityQueue{
		heap:     make([]interface{}, length),
		last:     -1,
		function: function,
	}
}

func (pq *PriorityQueue) Push(elem interface{}) {
	pq.last++
	if pq.last < len(pq.heap) {
		pq.heap[pq.last] = elem
	} else {
		pq.heap = append(pq.heap, elem)
	}
	pq.moveUp(pq.last)
}

func (pq *PriorityQueue) Pop() interface{} {
	if pq.last < 0 {
		return nil
	}
	pop := pq.heap[0]
	pq.heap[0] = pq.heap[pq.last]
	pq.last--
	pq.moveDown(0)
	return pop
}

func (pq *PriorityQueue) moveDown(start int) {
	cur := start
	if 2*start+1 <= pq.last && pq.function(pq.heap[2*start+1], pq.heap[cur]) {
		cur = 2*start + 1
	}
	if 2*start+2 <= pq.last && pq.function(pq.heap[2*start+2], pq.heap[cur]) {
		cur = 2*start + 2
	}
	if cur == start {
		return
	}
	pq.heap[start], pq.heap[cur] = pq.heap[cur], pq.heap[start]
	pq.moveDown(cur)
}

func (pq *PriorityQueue) moveUp(end int) {
	start := (end - 1) / 2
	if end <= 0 {
		return
	}
	cur := end
	if pq.function(pq.heap[cur], pq.heap[start]) {
		cur = start
	}
	if cur != start {
		return
	}
	pq.heap[start], pq.heap[end] = pq.heap[end], pq.heap[start]
	pq.moveUp(start)
}
