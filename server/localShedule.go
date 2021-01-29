/*
 * @Author: Vongola
 * @FilePath: \JFFun\server\localShedule.go
 * @Date: 2021-01-24 22:43:11
 * @Description: file content
 * @描述: 文件描述
 * @LastEditTime: 2021-01-29 16:09:17
 * @LastEditors: Vongola
 */

package server

import (
	"JFFun/data"
	"JFFun/task"
)

type localSchedule struct {
	netSchedule

	mods     []*mod
	handlers map[data.Command][]*mod
}

func (s *localSchedule) Execute(task *task.Task) {

}

func (s *localSchedule) registHandlers(m *mod) {
	s.mods = append(s.mods, m)
	handlers := m.m.GetHandlers()

	if s.handlers == nil {
		s.handlers = make(map[data.Command][]*mod, len(handlers))
	}

	for cmd := range handlers {
		s.handlers[cmd] = append(s.handlers[cmd], m)
	}
}
