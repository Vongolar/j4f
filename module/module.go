package module

//Module 模块接口
type Module interface {
	Init(cfg string) error
}
