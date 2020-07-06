package timer

import (
	"context"
	"testing"
	"time"

	"github.com/zssky/log"
	"goprogramming/src/mysql/database"
)

func TestNewTimer(t *testing.T) {
	// 生产者
	//prod1 := func(mc chan interface{}) {
	//	for i := 0; i < 10; i++ {
	//		mc <- &database.UserInfo{
	//			Name: fmt.Sprintf("user-%d", i),
	//			Age:  10 * i,
	//		}
	//	}
	//}
	prod2 := Produce // 不需要带括号(),因为这是声明的函数类型而不是直接使用的函数。

	// 消费者
	c := func(m interface{}) {
		log.Infof("channel message is %v", m.(*database.UserInfo))
	}

	//ctx, _ := context.WithCancel(context.Background())

	tm := NewTimer(10, 10, context.Background(), prod2, c)
	tm.Run()

	time.Sleep(20 * time.Second)
}
