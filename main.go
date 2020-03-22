package main

// Queue is your basic queue data structure
type Queue struct {
	slice []int
}

// Enqueue adds the integer to the back of the queue
func (q *Queue) Enqueue(i int) {
	q.slice = append(q.slice, i)
}

// Dequeue returns the first item in the queue
func (q *Queue) Dequeue() int {
	v := q.slice[0]
	q.slice = q.slice[1:len(q.slice)]
	return v
}

func main() {

}
