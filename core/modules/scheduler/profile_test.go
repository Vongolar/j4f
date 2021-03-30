package mscheduler

import "fmt"

func Example_Queue() {
	q := createFixIntergerQueue(5)
	q.add(0)
	q.add(1)
	q.add(2)
	q.add(3)
	q.add(4)
	q.add(5)
	q.add(6)
	q.add(7)
	q.add(8)
	q.add(9)
	q.add(10)
	fmt.Println(q.get(3))
	//Output: 9
}

func Example_Ave() {
	q := createFixIntergerQueue(5)
	q.add(1)
	q.add(2)
	q.add(3)
	q.add(4)
	q.add(5)
	q.add(6)
	q.add(7)
	q.add(8)
	q.add(9)
	fmt.Println(q.average())
	//Output: 7
}
