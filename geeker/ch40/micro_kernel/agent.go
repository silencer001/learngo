package agent

import (
	"context"
	"errors"
)

type Event struct {
}

type EventReceiver interface {
	OnEvent(evt Event)
}

type Collector interface {
	Init(eventReceiver EventReceiver) error
	Start(agtCtx context.Context) error
	Stop() error
	Destroy() error
}

type Agent struct {
	collectors map[string]Collector
	evtBuf     chan Event
	cancel     context.CancelFunc
	ctx        context.Context
	state      int
}

var WrongStateError = errors.New("Agent state in error state.")

func (agt *Agent) RegisterCollector(name string, collector Collector) error {
	if agt.state != Waiting {
		return WrongStateError
	}
	agt.collectors[name] = collector
	return collector.Init(agt)
}
