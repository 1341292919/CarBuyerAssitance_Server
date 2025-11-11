package service

import (
	"CarBuyerAssitance/pkg/constants"
	"fmt"
	"github.com/cloudwego/hertz/pkg/app"
)

// 这里其实不允许返回为空值，返回空值必然导致后续业务错误
func GetUserIDFromContext(c *app.RequestContext) string {
	if c == nil || c.Keys == nil {
		panic(fmt.Errorf("stream c or c.key is nil"))
		return ""
	}
	data, exists := c.Keys[constants.ContextUserId]
	if !exists {
		panic(fmt.Errorf("userId is nil"))
		return ""
	}

	// 类型断言确保返回的是 string
	if userID, ok := data.(string); ok {
		return userID
	}

	panic(fmt.Errorf("userId is not string"))
	return ""
}
