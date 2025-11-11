package taskqueue

import (
	"CarBuyerAssitance/biz/dal/mysql"
	"CarBuyerAssitance/pkg/taskqueue"
	"context"
)

func Work(key string) {
	taskQueue.Start()
}
func AddUpdateScoreTask(ctx context.Context, key string, p *mysql.Points) {
	taskQueue.Add(key, taskqueue.QueueTask{Execute: func() error {
		return updateScoreTask(ctx, p)
	}})
}
func updateScoreTask(ctx context.Context, p *mysql.Points) error {
	err := mysql.CreatePointsRecord(ctx, p.UserID, p.Points, p.Reason)
	if err != nil {
		return err
	}
	return nil
}
