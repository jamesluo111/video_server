package taskrunner

import (
	"errors"
	"fmt"
	"log"
	"os"
	"sync"
	"video_server/scheduler/dbops"
)

func deleteVideo(vid string) error {
	fmt.Println(VIDEO_DIR + vid)
	err := os.Remove(VIDEO_DIR + vid)
	if err != nil && !os.IsNotExist(err) {
		return err
	}
	return nil
}

func VideoClearDispatcher(ch dataChan) error {
	//从数据库中获取数据
	data, err := dbops.GetDeleteVideoRecord(3)
	if err != nil {
		log.Printf("error to get delete video:%v", err)
	}
	//数据库没有数据直接return nil
	if len(data) == 0 {
		return errors.New("All tasks finished")
		//return nil
	}
	for _, id := range data {
		ch <- id
	}
	return nil
}

func VideoClearExecutor(ch dataChan) error {
	errMap := sync.Map{}
	var err error
forloop:
	for {
		select {
		case id := <-ch:
			go func(vid interface{}) {
				if err := deleteVideo(id.(string)); err != nil {
					errMap.Store(id, err)
					return
				}
				if err := dbops.DeleteVideoRecord(id.(string)); err != nil {
					errMap.Store(id, err)
					return
				}
			}(id)
		default:
			break forloop
		}
	}
	errMap.Range(func(key, value interface{}) bool {
		err = value.(error)
		if err != nil {
			return false
		}
		return true
	})
	return err

}
