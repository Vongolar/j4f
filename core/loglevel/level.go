package loglevel

type Level = int

const (
	INFO Level = iota
	WARNING
	ERROR
)

var names = map[Level]string{
	INFO:    `INFO`,
	WARNING: `WARN`,
	ERROR:   `ERRO`,
}

func GetLevelTag(lvl Level) string {
	return names[lvl]
}
