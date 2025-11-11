package mysql

import (
	"gorm.io/gorm"
	"time"
)

type User struct {
	UserId         string         `gorm:"primaryKey;size:50;column:user_id"`
	Username       string         `gorm:"size:50;not null;column:username"`
	Password       string         `gorm:"size:100;not null;column:password"`
	Phone          string         `gorm:"size:20;not null;uniqueIndex;column:phone"`
	BudgetMin      float64        `gorm:"type:decimal(10,2);default:0.00;column:budget_min"`
	BudgetMax      float64        `gorm:"type:decimal(10,2);default:0.00;column:budget_max"`
	PreferredType  string         `gorm:"size:20;default:'';column:preferred_type"`
	PreferredBrand string         `gorm:"size:50;default:'';column:preferred_brand"`
	Status         int8           `gorm:"default:1;column:status"`
	Address        string         `gorm:"size:255;column:address"`
	CreatedAt      time.Time      `gorm:"column:created_at"`
	UpdatedAt      time.Time      `gorm:"column:updated_at"`
	DeletedAt      gorm.DeletedAt `gorm:"index;column:deleted_at"`
}

// Consultation 咨询记录
type Consultation struct {
	ConsultID       int        `gorm:"primaryKey" json:"consult_id"`
	UserID          string     `gorm:"not null" json:"user_id"`
	BudgetRange     string     `gorm:"size:50;not null" json:"budget_range"`
	PreferredType   string     `gorm:"size:20;not null" json:"preferred_type"`
	UseCase         string     `gorm:"size:30;not null" json:"use_case"`
	FuelType        string     `gorm:"size:20;not null" json:"fuel_type"`
	BrandPreference string     `gorm:"size:50" json:"brand_preference"`
	ConsultContent  string     `gorm:"type:text" json:"consult_content"`
	CreatedAt       time.Time  `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt       time.Time  `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
	DeletedAt       *time.Time `gorm:"index" json:"deleted_at,omitempty"`
}

// ConsultResult 咨询结果
type ConsultResult struct {
	ResultID      int       `gorm:"primaryKey" json:"result_id"`
	ConsultID     int       `gorm:"not null" json:"consult_id"`
	Analysis      string    `gorm:"type:text;not null" json:"analysis"`
	Proposal      string    `gorm:"type:text;not null" json:"proposal"`
	RecommendCars string    `gorm:"type:json;not null" json:"recommend_cars"` // 存储Car数组的JSON
	CreatedAt     time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt     time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
}
type Points struct {
	PointID    int        `gorm:"primaryKey;column:point_id;autoIncrement" json:"point_id"`
	UserID     string     `gorm:"column:user_id;not null;index:idx_user_id" json:"user_id"`
	Points     int        `gorm:"column:points;not null" json:"points"`
	Reason     string     `gorm:"column:reason;size:50;not null" json:"reason"`
	CreateTime time.Time  `gorm:"column:create_time;default:CURRENT_TIMESTAMP" json:"create_time"`
	UpdateTime time.Time  `gorm:"column:update_time;default:CURRENT_TIMESTAMP;autoUpdateTime" json:"update_time"`
	DeletedAt  *time.Time `gorm:"column:deleted_at;index" json:"deleted_at,omitempty"`
}

type Gift struct {
	GiftID         int64     `gorm:"primaryKey;column:gift_id;autoIncrement" json:"gift_id"`
	GiftName       string    `gorm:"column:gift_name;size:100;not null" json:"gift_name"`
	RequiredPoints int       `gorm:"column:required_points;not null;index:idx_points" json:"required_points"`
	StockQuantity  int       `gorm:"column:stock_quantity;not null;default:0" json:"stock_quantity"`
	CoverImageURL  string    `gorm:"column:cover_image_url;size:500" json:"cover_image_url"`
	IsOnline       bool      `gorm:"column:is_online;not null;default:true" json:"is_online"`
	CreateTime     time.Time `gorm:"column:create_time;default:CURRENT_TIMESTAMP;not null" json:"create_time"`
	UpdateTime     time.Time `gorm:"column:update_time;default:CURRENT_TIMESTAMP;autoUpdateTime;not null" json:"update_time"`
}

type Exchange struct {
	ExchangeID   int       `gorm:"primaryKey;column:exchange_id;autoIncrement" json:"exchange_id"`
	UserID       string    `gorm:"column:user_id;not null;index:idx_user_id" json:"user_id"`
	GiftName     string    `gorm:"column:gift_name;size:50;not null" json:"gift_name"`
	NeedPoints   int       `gorm:"column:need_points;not null" json:"need_points"`
	ExchangeTime time.Time `gorm:"column:exchange_time;default:CURRENT_TIMESTAMP" json:"exchange_time"`
	Status       int8      `gorm:"column:status;default:0" json:"status"` // 0-待发货/1-已完成
}
