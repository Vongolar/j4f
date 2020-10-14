package gate

import (
	Jcommand "JFFun/data/command"
	Jerror "JFFun/data/error"
	"bufio"
	"fmt"
	"os"
	"strconv"
)

/*
控制台
长连接
*/

var consoleOut *bufio.Writer
var consoleIn *bufio.Scanner

func listenConsole(onAccept func(authorization string, conn connect) bool) {
	consoleOut = bufio.NewWriter(os.Stdout)
	consoleIn = bufio.NewScanner(os.Stdin)

Listen:
	for input.Scan() {
		if onAccept(input.Text()) {
			//鉴权成功
			break Listen
		}
	}

	//监听指令
	for input.Scan() {

	}
}

func waitConsole(onAccept func(authorization string, conn connect) bool) {

}

type consoleConn struct {
}

func (conn *consoleConn) sync(cmd Jcommand.Command, data []byte) error {

}

type consoleResp struct {
}

func (r *consoleResp) Reply(id int64, errCode Jerror.Error, data []byte) error {
	out := consoleOut
	resp := "error : " + strconv.Itoa(int(errCode)) + "\n"
	if data != nil && len(data) > 0 {
		resp += fmt.Sprintf("%s\n\n", data)
	}

	_, err := out.WriteString(resp)
	if err != nil {
		return err
	}
	return out.Flush()
}
