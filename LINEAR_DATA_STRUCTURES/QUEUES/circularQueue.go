package main

type CircularQueue struct {
	tamanho int
	inicio  int
	final   int
	arr     []interface{}
}

func NewCircularQueue(limite int) *CircularQueue {

	return &CircularQueue{
		tamanho: 0,
		inicio:  0,
		final:   -1,
		arr:     make([]interface{}, limite),
	}
}

func (fila *CircularQueue) Enqueue(value interface{}) bool {

	if fila.IsFull() {
		return false
	}

	fila.final = (fila.final + 1) % len(fila.arr)
	fila.arr[fila.final] = value
	fila.tamanho++

	return true
}

func (fila *CircularQueue) Dequeue() (val interface{}, ok bool) {

	if fila.IsEmpty() {
		return nil, false
	}

	val = fila.arr[fila.inicio]
	fila.arr[fila.inicio] = nil
	fila.inicio = (fila.inicio + 1) % len(fila.arr)
	fila.tamanho--

	return val, true
}

func (fila *CircularQueue) Front() (val interface{}, ok bool) {

	if fila.IsEmpty() {
		return nil, false
	}

	return fila.arr[fila.inicio], true
}

func (fila *CircularQueue) Rear() (val interface{}, ok bool) {

	if fila.IsEmpty() {
		return nil, false
	}

	return fila.arr[fila.final], true
}

func (fila *CircularQueue) IsEmpty() bool {
	return fila.tamanho == 0
}

func (fila *CircularQueue) IsFull() bool {
	return fila.tamanho == len(fila.arr)
}

func (fila *CircularQueue) Size() int {
	return len(fila.arr)
}
