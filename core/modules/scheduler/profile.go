package mscheduler

func createFixIntergerQueue(length int) *fixIntergerQueue {
	return &fixIntergerQueue{
		q:    make([]int, length),
		head: -1,
		tail: -1,
	}
}

func createFixBooleanQueue(length int) *fixBoolenQueue {
	return &fixBoolenQueue{
		q:    make([]bool, length),
		head: -1,
		tail: -1,
	}
}

type fixIntergerQueue struct {
	q    []int
	head int
	tail int
}

func (q *fixIntergerQueue) getLength() int {
	if q.head < 0 || q.tail < 0 {
		return 0
	}

	if q.head <= q.tail {
		return q.tail - q.head + 1
	}

	return len(q.q) - q.head + q.tail + 1
}

func (q *fixIntergerQueue) isFull() bool {
	return q.getLength() == len(q.q)
}

func (q *fixIntergerQueue) add(v int) {
	if q.isFull() {
		q.dequeue()
	}

	if q.getLength() == 0 {
		q.head = 0
		q.tail = 0
	} else {
		q.tail++
	}

	if q.head < 0 {
		q.head = 0
	}

	if q.tail == len(q.q) {
		q.tail = 0
	}

	q.q[q.tail] = v
}

func (q *fixIntergerQueue) dequeue() {
	l := q.getLength()
	if l == 0 {
		return
	}

	if l == 1 {
		q.head = -1
		q.tail = -1
		return
	}

	q.head++
}

func (q *fixIntergerQueue) get(index int) int {
	i, l := index+q.head, len(q.q)
	if i >= l {
		i = l - index - 1
	}
	return q.q[i]
}

func (q *fixIntergerQueue) average() int {
	l := q.getLength()
	if l == 0 {
		return 0
	}

	sum := 0
	for i := q.head; ; i++ {
		if i == l {
			i = 0
		}

		sum += q.q[i]

		if i == q.tail {
			break
		}
	}

	return sum / l
}

type fixBoolenQueue struct {
	q    []bool
	head int
	tail int
}

func (q *fixBoolenQueue) getLength() int {
	if q.head < 0 || q.tail < 0 {
		return 0
	}

	if q.head <= q.tail {
		return q.tail - q.head + 1
	}

	return len(q.q) - q.head + q.tail + 1
}

func (q *fixBoolenQueue) isFull() bool {
	return q.getLength() == len(q.q)
}

func (q *fixBoolenQueue) add(v bool) {
	if q.isFull() {
		q.dequeue()
	}

	if q.getLength() == 0 {
		q.head = 0
		q.tail = 0
	} else {
		q.tail++
	}

	if q.head < 0 {
		q.head = 0
	}

	if q.tail == len(q.q) {
		q.tail = 0
	}

	q.q[q.tail] = v
}

func (q *fixBoolenQueue) dequeue() {
	l := q.getLength()
	if l == 0 {
		return
	}

	if l == 1 {
		q.head = -1
		q.tail = -1
		return
	}

	q.head++
}

func (q *fixBoolenQueue) getValueLength(v bool) int {
	l := q.getLength()
	if l == 0 {
		return 0
	}

	sum := 0
	for i := q.head; ; i++ {
		if i == l {
			i = 0
		}

		if q.q[i] == v {
			sum++
		}

		if i == q.tail {
			break
		}
	}

	return sum
}
