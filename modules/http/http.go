/*
 * @Author: Vongola
 * @LastEditTime: 2021-02-08 18:53:49
 * @LastEditors: Vongola
 * @Description: file content
 * @FilePath: \JFFun\modules\http\http.go
 * @Date: 2021-02-04 14:53:32
 * @描述: 文件描述
 */

package http

import (
	"context"
	"io"
	jconfig "j4f/core/config"
	"j4f/core/message"
	"j4f/core/scheduler"
	"j4f/core/task"
	"j4f/data"
	"net/http"
	"strconv"
)

type M_Http struct {
	name      string
	ctx       context.Context
	scheduler scheduler.Scheduler
	cfg       config
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

	go m.Listen()
}

func (m *M_Http) Listen() {
	mux := http.NewServeMux()
	mux.HandleFunc(`/`, func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add(`Access-Control-Allow-Methods`, `OPTIONS, POST`)
		w.Header().Add(`Access-Control-Allow-Origin`, `*`)
		w.Header().Add(`Access-Control-Allow-Headers`, `*`)

		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusOK)
			return
		}

		if r.Method != http.MethodPost {
			w.WriteHeader(http.StatusMethodNotAllowed)
			return
		}

		cmdStr := r.Header.Get(`Command`)
		cmdInt, err := strconv.Atoi(cmdStr)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		seria := message.DefaultSeria
		seriaStr := r.Header.Get("Seria")
		if tmpSeria, err := strconv.Atoi(seriaStr); err != nil {
			seria = tmpSeria
		}

		var content []byte
		if r.ContentLength > 0 {
			content = make([]byte, r.ContentLength)
			_, err = r.Body.Read(content)
			if err != nil && err != io.EOF {
				w.WriteHeader(http.StatusBadRequest)
				return
			}
		}

		cr := task.NewChannelRequest(w)
		t := &task.Task{
			CMD:     data.Command(cmdInt),
			Request: cr,
			Seria:   seria,
		}

		if r.ContentLength > 0 {
			if t.Data, err = message.Decode(seria, r.Body, t.CMD); err != nil {
				w.WriteHeader(http.StatusBadRequest)
				return
			}
		}

		m.scheduler.Exec(t)

		errCode := cr.Wait()
		w.Header().Add("ErrorCode", strconv.Itoa(int(errCode)))
		w.WriteHeader(http.StatusOK)
	})

	m.scheduler.InfoTag(m.name, `HTTP Listen Port:`, m.cfg.Port)
	if err := http.ListenAndServe(":"+strconv.Itoa(m.cfg.Port), mux); err != nil {
		m.scheduler.ErrorTag(m.name, err)
	}
}
