package jgate

type config struct {
	HTTP                    string //http端口
	Websocket               string //websocket端口
	DataBase                string //数据库名
	ClearOfflineIntervalMin int    //清理不活跃用户间隔时间
	OfflineTimeoutMin       int    //不活跃用户时间
}
