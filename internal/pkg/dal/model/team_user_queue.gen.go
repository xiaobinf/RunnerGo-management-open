// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"

	"gorm.io/gorm"
)

const TableNameTeamUserQueue = "team_user_queue"

// TeamUserQueue mapped from table <team_user_queue>
type TeamUserQueue struct {
	ID        int64          `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	Email     string         `gorm:"column:email;not null" json:"email"`     // 邮箱
	TeamID    string         `gorm:"column:team_id;not null" json:"team_id"` // 团队id
	CreatedAt time.Time      `gorm:"column:created_at;not null;default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt time.Time      `gorm:"column:updated_at;not null;default:CURRENT_TIMESTAMP" json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"column:deleted_at" json:"deleted_at"`
}

// TableName TeamUserQueue's table name
func (*TeamUserQueue) TableName() string {
	return TableNameTeamUserQueue
}
