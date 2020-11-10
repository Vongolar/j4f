package lostcities

var cards []card

type card struct {
	city  int
	point int
}

func getCards() []card {
	if len(cards) == 0 {
		cards = make([]card, 60)
		index := 0
		for i := 0; i < 5; i++ {
			for j := 0; j < 3; j++ {
				cards[index] = card{
					city: i,
				}
				index++
			}
			for j := 2; j <= 10; j++ {
				cards[index] = card{
					city:  i,
					point: j,
				}
				index++
			}
		}
	}
	return cards
}

func (card card) isInvest() bool {
	return card.point == 0
}
