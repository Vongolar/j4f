package task

import "j4f/data/command"

type Task struct {
	CMD  command.Command
	Data interface{}
}

type TaskHandler func(*Task)
