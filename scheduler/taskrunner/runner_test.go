package taskrunner

import (
	"errors"
	"log"
	"testing"
	"time"
)

func TestRunner(t *testing.T) {
	//定义生产者方法
	d := func(dc dataChan) error {
		for i := 0; i < 30; i++ {
			//将数据交给dataChan
			dc <- i
			log.Printf("dispatch sent:%v", i)
		}
		return nil
	}

	e := func(dc dataChan) error {
	forloop:
		for {
			select {
			case d := <-dc:
				log.Printf("exceutor received:%v", d)
			default:
				break forloop
			}
		}
		return errors.New("Exector")
	}
	runner := NewRunner(30, false, d, e)
	go runner.startAllDispatch()
	time.Sleep(3 * time.Second)
}
