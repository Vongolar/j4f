package server

import (
	"context"
	"j4f/schedule"
	"log"
	"os"
	"os/signal"
	"sync"
)

type Server struct {
}

func (s *Server) Run() error {
	log.Println("sever start")
	ctx, cancel := context.WithCancel(context.Background())
	wg := new(sync.WaitGroup)

	err := s.startSchedule(ctx, wg)
	if err != nil {
		log.Println("sever start error", err)
		cancel()
		return err
	}

	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, os.Interrupt)
	<-signalChan

	log.Println("sever closing")

	cancel()
	wg.Wait()

	log.Println("sever closed")
	return nil
}

func (s *Server) startSchedule(ctx context.Context, wg *sync.WaitGroup) error {
	module := new(schedule.MSchedule)
	return module.Run(ctx, wg)
}
