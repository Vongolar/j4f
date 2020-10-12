package gate

import (
	Jerror "JFFun/data/error"
	"bufio"
	"fmt"
	"os"
	"strconv"
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

type consoleResp struct {
}

func (req *consoleResp) Reply(id int64, errCode Jerror.Error, data []byte) error {
	out := consoleOut
	if errCode != Jerror.Error_ok {
		out = consoleErr
	}
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
