/*
 * @Author: Vongola
 * @LastEditTime: 2021-01-29 18:58:29
 * @LastEditors: Vongola
 * @Description: file content
 * @FilePath: \JFFun\server\server_test.go
 * @Date: 2021-01-29 16:21:22
 * @描述: 文件描述
 */

package server

import (
	"fmt"
	"math/rand"
	"strconv"
	"testing"
	"time"
)

type node struct {
	id string

	links map[string]link
}

type link struct {
	delay int
	paths []*node
}

func (l *link) point() *node {
	return l.paths[len(l.paths)-1]
}

func Test_Graph(t *testing.T) {
	var nodes []*node
	for i := 0; i < 10; i++ {
		nodes = append(nodes, &node{id: strconv.Itoa(i), links: make(map[string]link)})
	}

	rand.Seed(time.Now().UnixNano())
	for _, n := range nodes {
		linkCnt := rand.Intn(len(nodes)/2) + 1
		for i := 0; i < linkCnt; i++ {
			choose := rand.Intn(len(nodes))
			for nodes[choose].id == n.id {
				choose = rand.Intn(len(nodes))
			}

			n.links[nodes[choose].id] = link{delay: rand.Intn(1000) + 1, paths: []*node{nodes[choose]}}
		}
	}

	fmt.Println("----------------------------------------------------------------------")
	fmt.Println("node 直连")
	for _, n := range nodes {
		fmt.Print(n.id, " -> ")
		for id, l := range n.links {
			fmt.Print(id, ":", l.delay, " , ")
		}
		fmt.Print("\n")
	}
	fmt.Println("node 直连 END")
	fmt.Println("----------------------------------------------------------------------")

	add := true
	for add {
		add = false
		for _, n := range nodes {
			n.ping()
		}
	}

}

func (n *node) ping() {
	for id, link := range n.links {
		for id2, link2 := range n.links {
			if id != id2 {
				point := link.point()
			}
		}
	}
}
