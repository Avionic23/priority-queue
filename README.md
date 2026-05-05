# Priority Queue

A generic priority queue implementation in Go backed by a binary heap.

## Implementation Details

**Generic via `interface{}`**
The queue stores elements as `interface{}`. The types used in `Push`, `Pop`, and the comparison function must be consistent — callers are responsible for casting to the correct type when reading values from `Pop`.

**Heap storage with a last-index pointer**
The underlying slice is never cropped on `Pop`. Instead, a `last` index tracks the logical end of the heap. Popping swaps the root with the last element and decrements `last`, leaving the tail of the slice untouched.

**Dynamic reallocation on overflow**
If `Push` is called when the number of elements exceeds the current slice length, the slice is grown automatically via `append`. No manual intervention is needed.

**Pre-allocating capacity**
If the maximum number of additional elements is known upfront, pass it as the `length` argument to `NewPriorityQueue`. This pre-allocates extra slots in the backing slice, avoiding reallocation during pushes as long as the number of new elements stays within that bound.

**Initialising from existing data with `Heapify`**
Pass a non-nil slice as the first argument to `NewPriorityQueue` to seed the heap with existing elements. The constructor runs `Heapify` (Floyd's O(n) algorithm) to arrange them into a valid heap without pushing each element individually.

## Usage

```go
// Empty heap, pre-allocate for up to 100 pushes
pq := priorityqueue.NewPriorityQueue(nil, 100, func(a, b interface{}) bool {
    return a.(int) < b.(int) // min-heap
})

pq.Push(5)
pq.Push(2)
pq.Push(1)

pq.Pop().(int) // returns 1
pq.Pop().(int) // returns 2
pq.Pop().(int) // returns 5

// Initialise from existing elements (heapified in O(n))
pq2 := priorityqueue.NewPriorityQueue([]interface{}{4, 3, 2, 1}, 0, func(a, b interface{}) bool {
    return a.(int) < b.(int)
})

pq2.Pop().(int) // returns 1
```

Pass any comparison function to control ordering (min-heap, max-heap, custom structs, etc.). When using custom types, ensure the comparison function and all `Push`/`Pop` calls use the same underlying type.

## Run

```bash
go run main.go
```
