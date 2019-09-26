package model

// A PriorityQueue implements heap.Interface and holds Items.
type PriorityQueue struct {
	students []*RankedStudent
}

func NewPriorityQueue(students []*RankedStudent) *PriorityQueue {
	return &PriorityQueue{
		students,
	}
}

func (pq *PriorityQueue) Len() int { return len(pq.students) }

func (pq *PriorityQueue) Less(i, j int) bool {
	students := pq.students
	// We want Pop to give us the highest, not lowest, priority so we use greater than here.
	return students[i].Rank() >= students[j].Rank()
}

func (pq *PriorityQueue) Swap(i, j int) {
	students := pq.students
	students[i], students[j] = students[j], students[i]
	students[i].SetIndex(i)
	students[j].SetIndex(j)
}

func (pq *PriorityQueue) Push(x interface{}) {
	students := pq.students
	n := len(students)
	item := x.(*RankedStudent)
	item.SetIndex(n)
	pq.students = append(students, item)
}

func (pq *PriorityQueue) Pop() interface{} {
	old := pq.students
	n := len(old)
	item := old[n-1]
	item.SetIndex(-1) // for safety
	pq.students = old[0 : n-1]
	return item
}

func (pq *PriorityQueue) Students() []*RankedStudent {
	return pq.students
}
