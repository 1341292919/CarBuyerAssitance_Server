package taskqueue

import (
	"CarBuyerAssitance/pkg/constants"
	"CarBuyerAssitance/pkg/taskqueue"
)

var taskQueue *taskqueue.BaseTaskQueue

func Init() {
	taskQueue = taskqueue.NewBaseTaskQueue()
	Work(constants.TaskQueue)
}
