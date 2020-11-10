package lostcities

import (
	"fmt"
	"testing"
)

func Test_Game(t *testing.T) {
	game := new(Game)
	game.Ready("1", []string{"player1", "player2"})
	game.Start()
}

func Benchmark_Game(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			game := new(Game)
			game.Ready("1", []string{"player1", "player2"})
			game.Start()
		}
	})
}

func Test_ChannelLen(t *testing.T) {
	type s struct {
		id string
		i  int
	}

	list := []s{s{id: "111"}}
	list[0].id = "0000"
	list[0].i = 100
	fmt.Printf("%v", list[0])
}
