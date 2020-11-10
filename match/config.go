package jmatch

type config struct {
	Buff                  int
	SureDurSec            int64
	ClearMatchSec         int64
	ClearMatchIntervalMin int64
	Game                  map[string]gameConfig
}

type gameConfig struct {
	Player int //游戏人数
}
