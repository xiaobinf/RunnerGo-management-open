// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"

	"gorm.io/gorm"
)

const TableNameTarget = "target"

// Target mapped from table <target>
type Target struct {
	ID            int64          `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`                      // id
	TargetID      string         `gorm:"column:target_id;not null" json:"target_id"`                             // 全局唯一ID
	TeamID        string         `gorm:"column:team_id;not null" json:"team_id"`                                 // 团队id
	TargetType    string         `gorm:"column:target_type;not null" json:"target_type"`                         // 类型：文件夹，接口，分组，场景,测试用例
	Name          string         `gorm:"column:name;not null" json:"name"`                                       // 名称
	ParentID      string         `gorm:"column:parent_id;not null;default:0" json:"parent_id"`                   // 父级ID
	Method        string         `gorm:"column:method;not null" json:"method"`                                   // 方法
	Sort          int32          `gorm:"column:sort;not null" json:"sort"`                                       // 排序
	TypeSort      int32          `gorm:"column:type_sort;not null" json:"type_sort"`                             // 类型排序
	Status        int32          `gorm:"column:status;not null;default:1" json:"status"`                         // 回收站状态：1-正常，2-回收站
	Version       int32          `gorm:"column:version;not null" json:"version"`                                 // 产品版本号
	CreatedUserID string         `gorm:"column:created_user_id;not null" json:"created_user_id"`                 // 创建人ID
	RecentUserID  string         `gorm:"column:recent_user_id;not null" json:"recent_user_id"`                   // 最近修改人ID
	Description   string         `gorm:"column:description" json:"description"`                                  // 备注
	Source        int32          `gorm:"column:source;default:1" json:"source"`                                  // 数据来源：1-正常来源，2-性能，3-自动化测试
	PlanID        string         `gorm:"column:plan_id;not null;default:0" json:"plan_id"`                       // 计划id
	SourceID      string         `gorm:"column:source_id;not null" json:"source_id"`                             // 引用来源ID
	IsChecked     int32          `gorm:"column:is_checked;not null;default:1" json:"is_checked"`                 // 是否开启：1-开启，2-关闭
	CreatedAt     time.Time      `gorm:"column:created_at;not null;default:CURRENT_TIMESTAMP" json:"created_at"` // 创建时间
	UpdatedAt     time.Time      `gorm:"column:updated_at;not null;default:CURRENT_TIMESTAMP" json:"updated_at"` // 更新时间
	DeletedAt     gorm.DeletedAt `gorm:"column:deleted_at" json:"deleted_at"`                                    // 删除时间
}

// TableName Target's table name
func (*Target) TableName() string {
	return TableNameTarget
}
