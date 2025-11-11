package pack

import (
	"CarBuyerAssitance/biz/dal/mysql"
	resp "CarBuyerAssitance/biz/model/model"
	"CarBuyerAssitance/biz/service/model"
	"strconv"
)

func User(user *mysql.User) *resp.UserInfo {
	return &resp.UserInfo{
		Username:       user.Username,
		UserId:         user.UserId,
		Phone:          user.Phone,
		BudgeMin:       user.BudgetMin,
		BudgeMax:       user.BudgetMax,
		PreferredType:  user.PreferredType,
		PreferredBrand: user.PreferredBrand,
		CreatedAt:      strconv.FormatInt(user.CreatedAt.Unix(), 10),
		UpdatedAt:      strconv.FormatInt(user.UpdatedAt.Unix(), 10),
		DeletedAt:      strconv.FormatInt(0, 10),
	}
}

func ConsultResult(data *model.ConsultResult) *resp.ConsultResult {
	result := &resp.ConsultResult{
		Analysis: data.Analysis,
		Proposal: data.Proposal,
	}

	// 转换Car列表
	if data.Result != nil {
		result.Result = make([]*resp.Car, len(data.Result))
		for i, car := range data.Result {
			result.Result[i] = &resp.Car{
				ImageUrl:          car.ImageUrl,
				CarName:           car.CarName,
				FuelConsumption:   car.FuelConsumption,
				Power:             car.Power,
				Seat:              car.Seat,
				Drive:             car.Drive,
				RecommendedReason: car.RecommendedReason,
			}
		}
	}

	return result
}
func Consultation(data *model.AllConsulation) *resp.Consultation {
	if data == nil {
		return nil
	}

	result := &resp.Consultation{
		Consult: &resp.Consult{
			UserId:          data.Consultation.UserId,
			ConsultId:       int64(data.Consultation.ConsultId),
			BudgetRange:     data.Consultation.BudgetRange,
			PreferredType:   data.Consultation.PreferredType,
			UseCase:         data.Consultation.UseCase,
			FuelType:        data.Consultation.FuelType,
			BrandPreference: data.Consultation.BrandPreference,
		},
	}

	// 转换ConsultResult
	if data.ConsultResult.Result != nil {
		result.ConsultResult = &resp.ConsultResult{
			Analysis: data.ConsultResult.Analysis,
			Proposal: data.ConsultResult.Proposal,
			Result:   make([]*resp.Car, len(data.ConsultResult.Result)),
		}

		// 转换Car列表
		for i, car := range data.ConsultResult.Result {
			result.ConsultResult.Result[i] = &resp.Car{
				ImageUrl:          car.ImageUrl,
				CarName:           car.CarName,
				FuelConsumption:   car.FuelConsumption,
				Power:             car.Power,
				Seat:              car.Seat,
				Drive:             car.Drive,
				RecommendedReason: car.RecommendedReason,
			}
		}
	} else {
		result.ConsultResult = &resp.ConsultResult{
			Analysis: data.ConsultResult.Analysis,
			Proposal: data.ConsultResult.Proposal,
			Result:   []*resp.Car{},
		}
	}

	return result
}

func PointList(data []*mysql.Points) *resp.PointList {
	if data == nil {
		return &resp.PointList{
			Item: []*resp.Point{},
			Num:  0,
			Sum:  0,
		}
	}

	// 计算总积分
	var totalSum int64 = 0
	for _, point := range data {
		totalSum += int64(point.Points)
	}

	// 转换每个Point
	points := make([]*resp.Point, len(data))
	for i, point := range data {
		points[i] = &resp.Point{
			PointID:   int64(point.PointID),
			UserID:    point.UserID,
			Points:    int64(point.Points),
			Reason:    point.Reason,
			CreatedAt: point.CreateTime.Format("2006-01-02 15:04:05"),
			UpdatedAt: point.UpdateTime.Format("2006-01-02 15:04:05"),
		}
	}

	return &resp.PointList{
		Item: points,
		Num:  int64(len(data)),
		Sum:  totalSum,
	}
}

func Gift(data []*mysql.Gift) *resp.GiftList {
	if data == nil {
		return &resp.GiftList{
			Item:  []*resp.Gift{},
			Total: 0,
		}
	}

	// 转换每个Gift
	gifts := make([]*resp.Gift, len(data))
	for i, gift := range data {
		// 将bool转换为int64
		var isOnline int64 = 0
		if gift.IsOnline {
			isOnline = 1
		}

		gifts[i] = &resp.Gift{
			GiftID:         gift.GiftID,
			GiftName:       gift.GiftName,
			RequiredPoints: int64(gift.RequiredPoints),
			StockQuantity:  int64(gift.StockQuantity),
			CoverImageURL:  gift.CoverImageURL,
			IsOnline:       isOnline,
			CreatedAt:      gift.CreateTime.Format("2006-01-02 15:04:05"),
			UpdatedAt:      gift.UpdateTime.Format("2006-01-02 15:04:05"),
		}
	}

	return &resp.GiftList{
		Item:  gifts,
		Total: int64(len(data)),
	}
}

func Order(data *mysql.Exchange) *resp.Order {
	return &resp.Order{
		UserID:     data.UserID,
		GiftName:   data.GiftName,
		NeedPoints: int64(data.NeedPoints),
		Status:     int64(data.Status),
		OrderTime:  data.ExchangeTime.Format("2006-01-02 15:04:05"),
		Id:         int64(data.ExchangeID),
	}
}
