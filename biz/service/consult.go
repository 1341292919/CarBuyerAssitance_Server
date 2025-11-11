package service

import (
	"CarBuyerAssitance/biz/dal/mysql"
	"CarBuyerAssitance/biz/service/model"
	"CarBuyerAssitance/pkg/utils"
	"context"
	"fmt"
	"github.com/cloudwego/hertz/pkg/app"
)

type ConsultService struct {
	ctx context.Context
	c   *app.RequestContext
}

func NewConsultService(ctx context.Context, c *app.RequestContext) *ConsultService {
	return &ConsultService{
		ctx: ctx,
		c:   c,
	}
}

func (svc *ConsultService) Consult(consult *model.Consult) (*model.ConsultResult, error) {
	id := GetUserIDFromContext(svc.c)
	con, err := mysql.CreateConsultation(svc.ctx, id, consult)
	if err != nil {
		return nil, err
	}
	consultResult, err := utils.CallOpenAIWithConsult(svc.ctx, consult)
	if err != nil {
		return nil, fmt.Errorf("call openai error:" + err.Error())
	}
	err = mysql.SaveConsultResult(svc.ctx, con.ConsultId, consultResult)
	if err != nil {
		return nil, err
	}
	return consultResult, nil
}

func (svc *ConsultService) QueryConsult(consult_id int) (*model.AllConsulation, error) {
	return mysql.QueryConsultMessage(svc.ctx, consult_id)
}
