package main

import "fmt"

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
	q := new(Queue)
	q.Enqueue(2)
	q.Enqueue(3)
	q.Enqueue(5)
	q.Enqueue(8)

	v := q.Dequeue()

	fmt.Println(v)
	fmt.Println(q.slice)
}
