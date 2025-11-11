package mysql

import (
	"CarBuyerAssitance/biz/service/model"
	"CarBuyerAssitance/pkg/constants"
	"CarBuyerAssitance/pkg/errno"
	"context"
	"encoding/json"
	"errors"
	"gorm.io/gorm"
)

func CreateConsultation(ctx context.Context, userID string, consult *model.Consult) (*model.Consultation, error) {
	consultation := &Consultation{
		UserID:          userID,
		BudgetRange:     consult.BudgetRange,
		PreferredType:   consult.PreferredType,
		UseCase:         consult.UseCase,
		FuelType:        consult.FuelType,
		BrandPreference: consult.BrandPreference,
	}

	// 保存到数据库
	if err := db.WithContext(ctx).Table(constants.TableConsult).Create(consultation).Error; err != nil {
		return nil, err
	}

	return &model.Consultation{
		UserId:          userID,
		BudgetRange:     consultation.BudgetRange,
		PreferredType:   consultation.PreferredType,
		UseCase:         consultation.UseCase,
		FuelType:        consultation.FuelType,
		BrandPreference: consultation.BrandPreference,
		ConsultId:       consultation.ConsultID,
	}, nil
}

func SaveConsultResult(ctx context.Context, consultID int, result *model.ConsultResult) error {
	// 将Car数组转换为JSON
	carsJSON, err := json.Marshal(result.Result)
	if err != nil {
		return err
	}

	consultResult := &ConsultResult{
		ConsultID:     consultID,
		Analysis:      result.Analysis,
		Proposal:      result.Proposal,
		RecommendCars: string(carsJSON),
	}

	return db.WithContext(ctx).Table(constants.TableConsultResult).Create(consultResult).Error
}

func QueryConsultMessage(ctx context.Context, consultID int) (*model.AllConsulation, error) {
	var (
		dbConsult Consultation
		dbResult  ConsultResult
	)

	// 查询咨询记录
	if err := db.WithContext(ctx).Table(constants.TableConsult).
		Where("consult_id = ?", consultID).
		First(&dbConsult).Error; err != nil {
		return nil, err
	}

	// 查询咨询结果
	if err := db.WithContext(ctx).Table(constants.TableConsultResult).
		Where("consult_id = ?", consultID).
		First(&dbResult).Error; err != nil {
		return nil, err
	}

	// 解析JSON格式的推荐车辆
	var cars []model.Car
	if err := json.Unmarshal([]byte(dbResult.RecommendCars), &cars); err != nil {
		return nil, err
	}

	// 构建咨询结果
	consultResult := model.ConsultResult{
		Analysis: dbResult.Analysis,
		Proposal: dbResult.Proposal,
		Result:   cars,
	}

	// 构建咨询信息
	consultation := model.Consultation{
		UserId:          dbConsult.UserID,
		ConsultId:       dbConsult.ConsultID,
		BudgetRange:     dbConsult.BudgetRange,
		PreferredType:   dbConsult.PreferredType,
		UseCase:         dbConsult.UseCase,
		FuelType:        dbConsult.FuelType,
		BrandPreference: dbConsult.BrandPreference,
	}

	// 构建完整结果
	allConsultation := &model.AllConsulation{
		Consultation:  consultation,
		ConsultResult: consultResult,
	}

	return allConsultation, nil
}

func GetOnlineGifts(ctx context.Context) ([]*Gift, error) {
	var gifts []*Gift
	err := db.WithContext(ctx).
		Table(constants.TableGift).
		Where("is_online = ?", true).
		Order("required_points ASC").
		Find(&gifts).
		Error
	return gifts, err
}

func IsGiftExist(ctx context.Context, gift_id int64) (bool, error) {
	var userInfo *Gift
	err := db.WithContext(ctx).
		Table(constants.TableGift).
		Where("gift_id = ?", gift_id).
		First(&userInfo).
		Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) { //没找到了说明用户不存在
			return false, nil
		}
		return false, errno.Errorf(errno.InternalDatabaseErrorCode, "mysql: failed to query gift: %v", err)
	}
	return true, nil
}
func QueryGiftById(ctx context.Context, gift_id int64) (*Gift, error) {
	var userInfo *Gift
	err := db.WithContext(ctx).
		Table(constants.TableGift).
		Where("gift_id = ?", gift_id).
		First(&userInfo).
		Error
	if err != nil {
		return nil, errno.NewErrNo(errno.InternalDatabaseErrorCode, "query GIFT Info error:"+err.Error())
	}
	return userInfo, nil
}

func BuyGift(ctx context.Context, gift_id int64) error {
	err := db.WithContext(ctx).
		Table(constants.TableGift).
		Where("gift_id = ?", gift_id).
		Update("stock_quantity", gorm.Expr("stock_quantity - ?", 1)).
		Error
	if err != nil {
		return errno.NewErrNo(errno.InternalDatabaseErrorCode, "update stockquantity error: "+err.Error())
	}
	return nil
}

func CreateExchange(ctx context.Context, exchange *Exchange) (*Exchange, error) {
	err := db.WithContext(ctx).
		Table(constants.TableExchange).
		Create(&exchange).Error
	if err != nil {
		return nil, errno.NewErrNo(errno.InternalDatabaseErrorCode, " create exchange"+err.Error())
	}
	return exchange, err
}
