package jlogin

import (
	jconfig "JFFun/serialization/config"
	"context"
)

//MLogin 登陆模块
type MLogin struct {
	cfg config
}

//Init 初始化
func (m *MLogin) Init(cfg string) error {
	if err := jconfig.LoadConfig(cfg, &m.cfg); err != nil {
		return err
	}

	if err := m.createDBTables(); err != nil {
		return err
	}
	return nil
}

//Run 运行
func (m *MLogin) Run(ctx context.Context, name string) {
}
