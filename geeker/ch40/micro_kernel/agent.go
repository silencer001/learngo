package agent

import "context"

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
	collector map[string]Collector
	evtBuf    chan Event
	cancel    context.CancelFunc
	ctx       context.Context
	state     int
}
