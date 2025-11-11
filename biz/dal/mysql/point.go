package mysql

import "context"

func CreatePointsRecord(ctx context.Context, userID string, points int, reason string) error {
	record := &Points{
		UserID: userID,
		Points: points,
		Reason: reason,
	}
	return db.WithContext(ctx).Create(record).Error
}

// 查询用户的积分记录
func GetUserPoints(ctx context.Context, userID string) ([]*Points, error) {
	var records []*Points
	err := db.WithContext(ctx).
		Where("user_id = ? AND deleted_at IS NULL", userID).
		Order("create_time DESC").
		Find(&records).Error
	return records, err
}

// 计算用户总积分
func GetUserTotalPoints(ctx context.Context, userID string) (int, error) {
	var total int
	err := db.WithContext(ctx).Model(&Points{}).
		Where("user_id = ? AND deleted_at IS NULL", userID).
		Select("COALESCE(SUM(points), 0)").
		Scan(&total).Error
	return total, err
}
