package util

type Queue struct {
	elements []interface{}
}

func (queue *Queue) Push(elem interface{}) {
	queue.elements = append(queue.elements, elem)
}

func (queue *Queue) Pop() interface{} {
	var elem interface{}
	elem, queue.elements = queue.elements[0], queue.elements[1:]
	return elem
}
