// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"

	"gorm.io/gorm"
)

const TableNameTeam = "team"

// Team mapped from table <team>
type Team struct {
	ID                 int64          `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	Name               string         `gorm:"column:name;not null" json:"name"`                       // 团队名称
	Type               int32          `gorm:"column:type;not null" json:"type"`                       // 团队类型 1: 私有团队；2: 普通团队
	CreatedUserID      int64          `gorm:"column:created_user_id;not null" json:"created_user_id"` // 创建者id
	CreateUserIdentify string         `gorm:"column:create_user_identify;not null" json:"create_user_identify"`
	CreatedAt          time.Time      `gorm:"column:created_at;not null;default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt          time.Time      `gorm:"column:updated_at;not null;default:CURRENT_TIMESTAMP" json:"updated_at"`
	DeletedAt          gorm.DeletedAt `gorm:"column:deleted_at" json:"deleted_at"`
}

// TableName Team's table name
func (*Team) TableName() string {
	return TableNameTeam
}
