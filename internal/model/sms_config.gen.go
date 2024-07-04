// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

const TableNameSmsConfig = "sms_config"

// SmsConfig mapped from table <sms_config>
type SmsConfig struct {
	ID        int    `gorm:"column:id;primaryKey;autoIncrement:true;comment:ID" json:"id"` // ID
	Name      string `gorm:"column:name;not null;comment:配置键" json:"name"`                 // 配置键
	Value     string `gorm:"column:value;not null;comment:配置值" json:"value"`               // 配置值
	Desc      string `gorm:"column:desc;not null;comment:描述" json:"desc"`                  // 描述
	CreateBy  int    `gorm:"column:create_by;not null;comment:创建者" json:"create_by"`       // 创建者
	UpdateBy  int    `gorm:"column:update_by;not null;comment:更新者" json:"update_by"`       // 更新者
	CreatedAt int    `gorm:"column:created_at;not null;comment:创建时间" json:"created_at"`    // 创建时间
	UpdatedAt int    `gorm:"column:updated_at;not null;comment:更新时间" json:"updated_at"`    // 更新时间
}

// TableName SmsConfig's table name
func (*SmsConfig) TableName() string {
	return TableNameSmsConfig
}
