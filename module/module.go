package module

type Module interface {
	Init(configPath string) error
}
