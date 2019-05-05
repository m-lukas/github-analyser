package util

type Queue struct {
	Elements []interface{}
}

func (queue *Queue) Push(elem interface{}) {
	queue.Elements = append(queue.Elements, elem)
}

func (queue *Queue) Pop() interface{} {
	var elem interface{}
	elem, queue.Elements = queue.Elements[0], queue.Elements[1:]
	return elem
}
