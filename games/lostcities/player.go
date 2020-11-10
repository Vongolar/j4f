package lostcities

type player struct {
	id    string
	hand  []card //手牌
	table []card //出牌
}

// 添加到手牌
func (player *player) addHandCard(card card) {
	player.hand = append(player.hand, card)
}

// 减少手牌
func (player *player) subHandCard(card card) {
	for i, c := range player.hand {
		if c == card {
			player.hand = append(player.hand[:i], player.hand[i+1:]...)
			return
		}
	}
}

func (player *player) invest(card card) {
	player.subHandCard(card)
	player.table = append(player.table, card)
}

func (player *player) inHand(card card) bool {
	for _, c := range player.hand {
		if c == card {
			return true
		}
	}
	return false
}

func (player *player) canInvest(card card) bool {
	for _, c := range player.hand {
		if c.city == card.city && c.point > card.point {
			return false
		}
	}
	return true
}

// 积分
func (player *player) point() ([5]int, int) {
	score := 0
	var detail [5]int
	for i := 0; i < 5; i++ {
		invest, cnt, point := 0, 0, 0
		for j := 0; j < len(player.table) && player.table[j].city == i; j++ {
			cnt++
			if player.table[j].isInvest() {
				invest++
			} else {
				point += player.table[j].point
			}
		}
		if cnt > 0 {
			s := -20
			s += point
			s *= (invest + 1)
			if cnt >= 8 {
				s += 20
			}
			detail[i] = score
			score += s
		}
	}
	return detail, score
}
