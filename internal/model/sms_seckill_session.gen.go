// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

const TableNameSmsSeckillSession = "sms_seckill_session"

// SmsSeckillSession mapped from table <sms_seckill_session>
type SmsSeckillSession struct {
	ID        int    `gorm:"column:id;primaryKey;autoIncrement:true;comment:ID" json:"id"` // ID
	Name      string `gorm:"column:name;not null;comment:场次名" json:"name"`                 // 场次名
	StartAt   int64  `gorm:"column:start_at;not null;comment:开始时间" json:"start_at"`        // 开始时间
	EndAt     int64  `gorm:"column:end_at;not null;comment:结束时间" json:"end_at"`            // 结束时间
	IsRelease int8   `gorm:"column:is_release;not null;comment:是否发布上线" json:"is_release"`  // 是否发布上线
	CreatedAt int    `gorm:"column:created_at;not null;comment:创建时间" json:"created_at"`    // 创建时间
	UpdatedAt int    `gorm:"column:updated_at;not null;comment:更新时间" json:"updated_at"`    // 更新时间
}

// TableName SmsSeckillSession's table name
func (*SmsSeckillSession) TableName() string {
	return TableNameSmsSeckillSession
}
