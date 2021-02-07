/*
 * @Author: Vongola
 * @LastEditTime: 2021-02-05 12:07:21
 * @LastEditors: Vongola
 * @Description: file content
 * @FilePath: \JFFun\modules\http\http.go
 * @Date: 2021-02-04 14:53:32
 * @描述: 文件描述
 */

package http

import (
	"context"
	"fmt"
	"io/ioutil"
	jconfig "j4f/core/config"
	"j4f/core/scheduler"
	"j4f/core/task"
	"net/http"
	"strconv"
)

type M_Http struct {
	name      string
	ctx       context.Context
	cfg       config
	scheduler scheduler.Scheduler
}

func (m *M_Http) Init(ctx context.Context, name string, cfgPath string) error {
	m.name = name
	m.ctx = ctx

	if err := jconfig.DecodeConfigFromFile(cfgPath, &m.cfg); err != nil {
		return err
	}

	return nil
}

func (m *M_Http) Run(c chan *task.TaskHandleTuple, s scheduler.Scheduler) {
	m.scheduler = s
	close(c)

	go m.Listen()
}

func (m *M_Http) Listen() {
	mux := http.NewServeMux()
	mux.HandleFunc(`/`, func(w http.ResponseWriter, r *http.Request) {
		b, _ := ioutil.ReadAll(r.Body)
		fmt.Println(string(b))
		w.WriteHeader(http.StatusOK)
	})

	if err := http.ListenAndServe(":"+strconv.Itoa(m.cfg.Port), mux); err != nil {
		m.scheduler.ErrorTag(m.name, err)
	}
}
