package schedule

import (
	Jcommand "JFFun/data/command"
	Jtask "JFFun/task"
)

type task struct {
	cmd  Jcommand.Command
	task *Jtask.Task
}
