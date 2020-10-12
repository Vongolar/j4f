package gate

import (
	Jerror "JFFun/data/error"
	"bufio"
	"os"
	"unsafe"
)

var consoleOut *bufio.Writer
var consoleErr *bufio.Writer

func listenConsole(on func(string)) {
	consoleOut = bufio.NewWriter(os.Stdout)
	consoleErr = bufio.NewWriter(os.Stderr)
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		on(input.Text())
	}
}

type consoleReq struct {
}

func (req *consoleReq) Reply(errCode Jerror.Error, msg []byte) error {
	out := consoleOut
	if errCode != Jerror.Error_ok {
		out = consoleErr
	}
	_, err := out.WriteString(*(*string)(unsafe.Pointer(&msg)))
	if err != nil {
		return err
	}
	return out.Flush()
}
