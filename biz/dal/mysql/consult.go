package mysql

import (
	"CarBuyerAssitance/biz/service/model"
	"CarBuyerAssitance/pkg/constants"
	"context"
	"encoding/json"
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
