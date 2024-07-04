// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"

	"gorm.io/gorm"
)

const TableNameUmsMember = "ums_member"

// UmsMember mapped from table <ums_member>
type UmsMember struct {
	ID          int            `gorm:"column:id;primaryKey;autoIncrement:true;comment:ID" json:"id"`                        // ID
	LevelID     int            `gorm:"column:level_id;not null;comment:会员等级id" json:"level_id"`                             // 会员等级id
	Username    string         `gorm:"column:username;not null;comment:用户名" json:"username"`                                // 用户名
	Nickname    string         `gorm:"column:nickname;not null;comment:昵称" json:"nickname"`                                 // 昵称
	Password    string         `gorm:"column:password;not null;comment:密码" json:"password"`                                 // 密码
	Phone       int64          `gorm:"column:phone;not null;comment:手机号" json:"phone"`                                      // 手机号
	Email       string         `gorm:"column:email;not null;comment:邮箱" json:"email"`                                       // 邮箱
	Avatar      string         `gorm:"column:avatar;not null;comment:头像" json:"avatar"`                                     // 头像
	Gender      int8           `gorm:"column:gender;not null;default:1;comment:性别" json:"gender"`                           // 性别
	Birth       time.Time      `gorm:"column:birth;comment:生日" json:"birth"`                                                // 生日
	Area        string         `gorm:"column:area;not null;comment:城市" json:"area"`                                         // 城市
	Job         string         `gorm:"column:job;not null;comment:职业" json:"job"`                                           // 职业
	SourceType  int8           `gorm:"column:source_type;not null;comment:用户来源" json:"source_type"`                         // 用户来源
	Integration int64          `gorm:"column:integration;not null;comment:积分" json:"integration"`                           // 积分
	Growth      int64          `gorm:"column:growth;not null;comment:成长值" json:"growth"`                                    // 成长值
	Status      int8           `gorm:"column:status;not null;default:1;comment:状态" json:"status"`                           // 状态
	Sign        string         `gorm:"column:sign;not null;comment:签名" json:"sign"`                                         // 签名
	CreatedAt   time.Time      `gorm:"column:created_at;not null;default:CURRENT_TIMESTAMP;comment:创建时间" json:"created_at"` // 创建时间
	UpdatedAt   time.Time      `gorm:"column:updated_at;not null;default:CURRENT_TIMESTAMP;comment:更新时间" json:"updated_at"` // 更新时间
	DeletedAt   gorm.DeletedAt `gorm:"column:deleted_at;comment:删除时间" json:"deleted_at"`                                    // 删除时间
}

// TableName UmsMember's table name
func (*UmsMember) TableName() string {
	return TableNameUmsMember
}
