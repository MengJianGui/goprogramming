package timer

import (
	"context"
	"fmt"
	"time"

	"github.com/zssky/log"
	"goprogramming/src/mysql/database"
)

/***
timer for time scheduler
*/

// Timer for schedule
type Timer struct {
	period int64                  // cycle period
	concur int                    // concurrent numbers for consumer
	ctx    context.Context        // using when task is finished
	mc     chan interface{}       // message channel
	p      func(chan interface{}) // producer get tasks from DBA
	c      func(interface{})      // consumer do tasks from producer
}

// NewTimer: period：cycle period seconds; con: concurrent thread; ctx:context; p: produce func; c: consumer func
func NewTimer(period int64, concu int, ctx context.Context, p func(chan interface{}), c func(interface{})) *Timer {
	return &Timer{
		period: period,
		concur: concu,
		ctx:    ctx,
		mc:     make(chan interface{}, 64),
		p:      p,
		c:      c,
	}
}

// Run for time schedule using eg: t.Run()
func (t *Timer) Run() {
	go t.produce()

	log.Debugf("run timer %d", t.period)
	// holding more than 20 concurrences
	for i := 0; i < t.concur; i++ {
		go t.consumer()
	}
}

// produce message into channel
func (t *Timer) produce() {
	log.Debugf("timer running %d seconds", t.period)
	for {
		select {
		case <-t.ctx.Done():
			log.Warnf("timer %d is closed", t.period)
			return
		default:
			t.p(t.mc)
		}

		time.Sleep(time.Duration(t.period) * time.Second)
	}
}

// consumer from channel messages
func (t *Timer) consumer() {
	for {
		select {
		case <-t.ctx.Done():
			log.Warnf("timer %d is closed", t.period)
			return
		case m, has := <-t.mc:
			if !has {
				// has more data on this channel
				log.Warnf("channel is closed")
				return
			}

			t.c(m)
		}
	}
}

func Produce(m chan interface{}) {
	for i := 0; i < 10; i++ {
		m <- &database.UserInfo{
			Name: fmt.Sprintf("user-%d", i),
			Age:  10 * i,
		}
	}
}
