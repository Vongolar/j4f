package request

import "j4f/data/errCode"

type Request interface {
	Reply(code errCode.Code, ext interface{})
}
