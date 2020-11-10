package lostcities

import (
	"JFFun/data/DLostCities"
	"JFFun/data/Derror"
	"JFFun/data/Dgame"
	jgamecore "JFFun/games/core"
	jlog "JFFun/log"
	jtag "JFFun/log/tag"
	jserialization "JFFun/serialization"
	jtask "JFFun/task"
	"fmt"
	"math/rand"
	"time"
)

//Game 失落的城市
type Game struct {
	id        string
	state     DLostCities.State
	curPlayer int
	players   [2]player
	drops     [5][]card //弃牌
	deck      []card
	core      *jgamecore.Core
}

const (
	dealSec     = 1  //发牌时间
	actSec      = 15 //出牌时间
	actWaitSec  = 1  //出牌动画时间
	drawSec     = 10 //抽牌时间
	drawWaitSec = 1  //抽牌动画时间
	balanceSec  = 5  //结算时间
)

//Ready 准备
func (game *Game) Ready(id string, players []string) *jgamecore.Core {
	game.id = id

	game.players[0].id = players[0]
	game.players[1].id = players[1]

	game.core = new(jgamecore.Core)
	game.core.Init(game.onGetGameInfo)

	//生成牌
	game.deck = getCards()

	//洗牌
	rand.Shuffle(len(game.deck), func(i int, j int) {
		game.deck[i], game.deck[j] = game.deck[j], game.deck[i]
	})

	jlog.Info(jtag.Game(game.id), `洗牌`)

	return game.core
}

//Start 游戏开始
func (game *Game) Start() {
	// 发牌
	game.deal()

	for len(game.deck) > 0 {
		jlog.Info(jtag.Game(game.id), game.players[game.curPlayer].id, `回合开始`)
		game.turn(&game.players[game.curPlayer])
		jlog.Info(jtag.Game(game.id), game.players[game.curPlayer].id, `回合结束`)
		game.nextTurn()
	}

	// 结算
	game.balance()
}

//发牌
func (game *Game) deal() {
	game.state = DLostCities.State_deal

	for p := 0; p < len(game.players); p++ {
		cards := game.deck[len(game.deck)-8:]
		game.deck = game.deck[:len(game.deck)-8]
		for _, c := range cards {
			game.players[p].addHandCard(c)
		}
		jlog.Info(jtag.Game(game.id), `发牌`, game.players[p].id, fmt.Sprintf("%v", game.players[p].hand))
	}

	//通知玩家发牌
	for _, player := range game.players {
		cards := new(DLostCities.Cards)
		cards.Cards = make([]*DLostCities.Card, len(player.hand))
		for i, card := range player.hand {
			cards.Cards[i] = &DLostCities.Card{
				City:  int32(card.city),
				Point: int32(card.point),
			}
		}

		game.core.SyncGameState(int32(DLostCities.State_deal), cards, player.id)
	}

	//等待玩家发牌结束
	game.core.WaitStateEnd(dealSec, int32(DLostCities.State_deal), game.players[0].id, game.players[1].id)

	jlog.Info(jtag.Game(game.id), `发牌结束`)
}

//turn 单个玩家回合
func (game *Game) turn(player *player) {
	game.state = DLostCities.State_play
	//通知玩家出牌
	game.core.SyncGameState(int32(DLostCities.State_play), &Dgame.WaitDeadLine{
		Player: player.id,
		Ts:     actSec + time.Now().Unix(),
	}, game.players[0].id, game.players[1].id)

	//等待玩家出牌
	var act DLostCities.Act
	var curCard card
	game.core.WaitAct(actSec, func() {
		// 超时随机丢弃一张
		act = DLostCities.Act_drop
		curCard = player.hand[rand.Intn(len(player.hand))]
	}, func(task *jtask.Task) bool {
		if data, ok := task.Data.(*Dgame.Action); ok {
			if !(data.Act == int32(DLostCities.Act_invest) || data.Act == int32(DLostCities.Act_drop)) {
				task.Error(Derror.Error_badRequest)
				return false
			}
			c := &DLostCities.Card{}
			err := jserialization.UnMarshal(jserialization.DefaultMode, data.Data, c)
			if err != nil {
				task.Error(Derror.Error_badRequest)
				return false
			}
			curCard = card{
				city:  int(c.City),
				point: int(c.Point),
			}
			if !player.inHand(curCard) {
				task.Error(Derror.Error_noInHand)
				return false
			}

			if data.Act == int32(DLostCities.Act_invest) {
				if !player.canInvest(curCard) {
					task.Error(Derror.Error_invaildInvest)
					return false
				}
			}

			act = DLostCities.Act(data.Act)
			task.OK()
			return true
		}
		task.Error(Derror.Error_badRequest)
		return false
	}, player.id)

	switch act {
	case DLostCities.Act_invest:
		game.invest(player, curCard)
	case DLostCities.Act_drop:
		game.drop(player, curCard)
	}

	game.core.WaitStateEnd(actWaitSec, int32(DLostCities.State_play), game.players[0].id, game.players[1].id)

	game.state = DLostCities.State_getCard
	//通知玩家拿牌
	game.core.SyncGameState(int32(DLostCities.State_getCard), &Dgame.WaitDeadLine{
		Player: player.id,
		Ts:     drawSec + time.Now().Unix(),
	}, game.players[0].id, game.players[1].id)
	//等待玩家拿牌
	game.core.WaitAct(drawSec, func() {
		act = DLostCities.Act_draw
	}, func(task *jtask.Task) bool {
		if data, ok := task.Data.(*Dgame.Action); ok {
			if data.Act == int32(DLostCities.Act_draw) {
				act = DLostCities.Act_draw
				task.OK()
				return true
			}
			if data.Act == int32(DLostCities.Act_choose) {
				c := &DLostCities.Card{}
				err := jserialization.UnMarshal(jserialization.DefaultMode, data.Data, c)
				if err != nil {
					task.Error(Derror.Error_badRequest)
					return false
				}
				curCard = card{
					city:  int(c.City),
					point: int(c.Point),
				}
				if !game.inDropTop(curCard) {
					task.Error(Derror.Error_noInDropTop)
					return false
				}
			}
		}
		task.Error(Derror.Error_badRequest)
		return false
	}, player.id)

	switch act {
	case DLostCities.Act_draw:
		game.draw(player)
	case DLostCities.Act_choose:
		game.choose(player, curCard)
	}

	game.core.WaitStateEnd(drawWaitSec, int32(DLostCities.State_getCard), game.players[0].id, game.players[1].id)
}

//投资
func (game *Game) invest(player *player, card card) {
	jlog.Info(jtag.Game(game.id), player.id, `投资`, fmt.Sprintf("%v", card))
	player.invest(card)

	game.core.SyncPlayerAct(int32(DLostCities.Act_invest), player.id, &DLostCities.Card{
		City:  int32(card.city),
		Point: int32(card.point),
	}, game.players[0].id, game.players[1].id)
}

//弃牌
func (game *Game) drop(player *player, card card) {
	jlog.Info(jtag.Game(game.id), player.id, `弃牌`, fmt.Sprintf("%v", card))
	player.subHandCard(card)
	game.drops[card.city] = append(game.drops[card.city], card)

	game.core.SyncPlayerAct(int32(DLostCities.Act_drop), player.id, &DLostCities.Card{
		City:  int32(card.city),
		Point: int32(card.point),
	}, game.players[0].id, game.players[1].id)
}

//选牌
func (game *Game) choose(player *player, card card) {
	jlog.Info(jtag.Game(game.id), player.id, `选牌`, fmt.Sprintf("%v", card))
	l := len(game.drops[card.city])
	game.drops[card.city] = game.drops[card.city][:l-1]
	player.addHandCard(card)

	game.core.SyncPlayerAct(int32(DLostCities.Act_choose), player.id, &DLostCities.Card{
		City:  int32(card.city),
		Point: int32(card.point),
	}, game.players[0].id, game.players[1].id)
}

//抽牌
func (game *Game) draw(player *player) {
	card := game.deck[len(game.deck)-1]
	player.addHandCard(card)
	game.deck = game.deck[:len(game.deck)-1]

	jlog.Info(jtag.Game(game.id), player.id, `抽牌`, fmt.Sprintf("%v", card))

	game.core.SyncPlayerAct(int32(DLostCities.Act_draw), player.id, &DLostCities.Card{
		City:  int32(card.city),
		Point: int32(card.point),
	}, player.id)

	otherPlayer := game.players[0].id
	if otherPlayer == player.id {
		otherPlayer = game.players[1].id
	}
	game.core.SyncPlayerAct(int32(DLostCities.Act_draw), player.id, &DLostCities.Card{
		City:  -1,
		Point: -1,
	}, otherPlayer)
}

func (game *Game) balance() {
	game.state = DLostCities.State_balance

	balance := new(DLostCities.BalanceResult)
	for _, player := range game.players {
		ds, score := player.point()
		detail := &DLostCities.Score{
			Player:  player.id,
			Score:   int32(score),
			Details: make([]int32, len(ds)),
		}

		for i := 0; i < len(ds); i++ {
			detail.Details[i] = int32(ds[i])
		}
		balance.Details = append(balance.Details, detail)
	}

	jlog.Info(jtag.Game(game.id), `结算`, fmt.Sprintf("%v", balance))

	game.core.SyncGameState(int32(DLostCities.State_balance), balance, game.players[0].id, game.players[1].id)

	game.core.WaitStateEnd(balanceSec, int32(DLostCities.State_balance), game.players[0].id, game.players[1].id)
}

func (game *Game) nextTurn() {
	game.curPlayer++
	if game.curPlayer+1 > len(game.players) {
		game.curPlayer = 0
	}
}

func (game *Game) onGetGameInfo(task *jtask.Task) {
	info := game.getGameInfo(task.PlayerID)
	task.Reply(Derror.Error_ok, info)
}

func (game *Game) getGameInfo(playerID string) *DLostCities.GameInfo {
	info := new(DLostCities.GameInfo)
	info.State = game.state

	info.Players = make([]*DLostCities.PlayerInfo, len(game.players))
	for pi, player := range game.players {
		info.Players[pi].PlayerId = player.id

		info.Players[pi].Table = make([]*DLostCities.Card, len(player.table))
		for ci, card := range player.table {
			info.Players[pi].Table[ci].City = int32(card.city)
			info.Players[pi].Table[ci].Point = int32(card.point)
		}

		info.Players[pi].Hands = make([]*DLostCities.Card, len(player.hand))
		for ci, card := range player.hand {
			info.Players[pi].Hands[ci].City = int32(card.city)
			info.Players[pi].Hands[ci].Point = int32(card.point)
		}
	}

	info.DropArea = make([]*DLostCities.DropArea, len(game.drops))
	for di, area := range game.drops {
		info.DropArea[di].City = int32(di)
		info.DropArea[di].Cards = make([]*DLostCities.Card, len(area))
		for ci, card := range area {
			info.DropArea[di].Cards[ci] = &DLostCities.Card{
				City:  int32(card.city),
				Point: int32(card.point),
			}
		}
	}

	info.Deck = int32(len(game.deck))
	info.CurPlayer = game.players[game.curPlayer].id

	//上帝视角
	if len(playerID) == 0 {
		return info
	}

	//玩家视角,隐藏非自己的手牌，旁观玩家完全看不到手牌
	for _, player := range info.Players {
		if player.PlayerId != playerID {
			for _, card := range player.Hands {
				// -1 表示隐藏
				card.Point = -1
				card.City = -1
			}
		}
	}

	return info
}

func (game *Game) inDropTop(card card) bool {
	l := len(game.drops[card.city])
	if l == 0 {
		return false
	}
	return game.drops[card.city][l] == card
}
