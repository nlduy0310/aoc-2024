package main

type Queue[T CopySafePrimitive] struct {
	queue []T
}

func NewEmptyQueue[T CopySafePrimitive]() Queue[T] {

	return Queue[T]{queue: []T{}}
}

func (q *Queue[T]) Push(element T) {

	q.queue = append(q.queue, element)
}

func (q *Queue[T]) Pop() *T {

	if len(q.queue) == 0 {
		return nil
	}

	first := q.queue[0]
	q.queue = q.queue[1:]

	return &first
}

func (q *Queue[T]) Copy() Queue[T] {

	return Queue[T]{queue: safeCopyList(q.queue)}
}
