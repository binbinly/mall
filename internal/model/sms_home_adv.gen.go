// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

const TableNameSmsHomeAdv = "sms_home_adv"

// SmsHomeAdv mapped from table <sms_home_adv>
type SmsHomeAdv struct {
	ID         int    `gorm:"column:id;primaryKey;autoIncrement:true;comment:ID" json:"id"` // ID
	Title      string `gorm:"column:title;not null;comment:标题" json:"title"`                // 标题
	Img        string `gorm:"column:img;not null;comment:图片" json:"img"`                    // 图片
	StartAt    int64  `gorm:"column:start_at;not null;comment:开始时间" json:"start_at"`        // 开始时间
	EndAt      int64  `gorm:"column:end_at;not null;comment:结束时间" json:"end_at"`            // 结束时间
	ClickCount int64  `gorm:"column:click_count;not null;comment:点击数" json:"click_count"`   // 点击数
	URL        string `gorm:"column:url;not null;comment:链接地址" json:"url"`                  // 链接地址
	Note       string `gorm:"column:note;not null;comment:备注" json:"note"`                  // 备注
	IsRelease  int8   `gorm:"column:is_release;not null;comment:是否发布上线" json:"is_release"`  // 是否发布上线
	Sort       int    `gorm:"column:sort;not null;default:50;comment:排序" json:"sort"`       // 排序
	CreateBy   int    `gorm:"column:create_by;not null;comment:创建者" json:"create_by"`       // 创建者
	UpdateBy   int    `gorm:"column:update_by;not null;comment:更新者" json:"update_by"`       // 更新者
	CreatedAt  int    `gorm:"column:created_at;not null;comment:创建时间" json:"created_at"`    // 创建时间
	UpdatedAt  int    `gorm:"column:updated_at;not null;comment:更新时间" json:"updated_at"`    // 更新时间
}

// TableName SmsHomeAdv's table name
func (*SmsHomeAdv) TableName() string {
	return TableNameSmsHomeAdv
}
