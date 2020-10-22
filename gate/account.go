package gate

type account struct {
	auth authorization
	id   string
	conn connect
}
