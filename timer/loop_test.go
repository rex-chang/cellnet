package timer

import (
	"fmt"
	"github.com/davyxu/cellnet"
	"testing"
	"time"
)

func TestLoopPanic(t *testing.T) {

	q := cellnet.NewEventQueue()
	q.EnableCapturePanic(true)

	var times = 3

	NewLoop(q, time.Millisecond*100, func(loop *Loop) {

		times--
		if times == 0 {
			loop.Stop()
			q.StopLoop(0)
		}

		fmt.Println("before")
		panic("panic")
		fmt.Println("after")

	}, nil).Start()

	q.StartLoop()

	q.Wait()
}
