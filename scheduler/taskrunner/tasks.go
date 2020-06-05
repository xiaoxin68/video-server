package taskrunner

import (
	"errors"
	"log"
	"os"
	"sync"
	"video-server/scheduler/handlers"
)

//删除磁盘中的video
func deleteVedio(vid string) error {
	err := os.Remove(VIDEO_PATH + vid)
	if err != nil && !os.IsNotExist(err) {
		log.Printf("Deleting video error: %v", err)
		return err
	}
	return nil
}

//conf数据库中读取指定条数的记录，并将id存入dataChan
func VideoClearDispatcher(dc dataChan) error {
	res, err := handlers.ReadVideoDeletionRecord(3)
	if err != nil {
		log.Printf("Video clear dispatcher error: %v", err)
		return err
	}

	if len(res) == 0 {
		return errors.New("All tasks finished")
	}

	for _, id := range res {
		dc <- id
	}

	return nil
}

//读取dataChan中的数据（存的是vid），并删除（数据库+磁盘）
func VideoClearExecutor(dc dataChan) error {
	errMap := &sync.Map{} //存储错误信息
	var err error

forloop:
	for {
		select {
		case vid := <-dc:
			go func(id interface{}) {
				//从磁盘删除
				if err := deleteVedio(id.(string)); err != nil {
					errMap.Store(id, err)
					return
				}
				//从数据库删除
				if err := handlers.DelVideoDeletionRecord(id.(string)); err != nil {
					errMap.Store(id, err)
					return
				}
			}(vid)
		default:
			break forloop
		}
	}

	//遍历错误的map，只掉有一个错误存在，就停止range过程，直接返回false
	errMap.Range(func(key, value interface{}) bool {
		err = value.(error)
		if err != nil {
			return false
		}
		return true
	})

	return err
}
