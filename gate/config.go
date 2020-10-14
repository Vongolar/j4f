package gate

type config struct {
	Console   bool
	HTTP      string
	Websocket string

	FitPlayerCount int
	CommandBuffer  int
}
