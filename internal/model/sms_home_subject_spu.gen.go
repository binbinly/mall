// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

const TableNameSmsHomeSubjectSpu = "sms_home_subject_spu"

// SmsHomeSubjectSpu mapped from table <sms_home_subject_spu>
type SmsHomeSubjectSpu struct {
	ID        int    `gorm:"column:id;primaryKey;autoIncrement:true;comment:ID" json:"id"` // ID
	Name      string `gorm:"column:name;not null;comment:名称" json:"name"`                  // 名称
	SubjectID int64  `gorm:"column:subject_id;not null;comment:专题id" json:"subject_id"`    // 专题id
	SpuID     int64  `gorm:"column:spu_id;not null;comment:spu_id" json:"spu_id"`          // spu_id
	Sort      int    `gorm:"column:sort;not null;default:50;comment:排序" json:"sort"`       // 排序
}

// TableName SmsHomeSubjectSpu's table name
func (*SmsHomeSubjectSpu) TableName() string {
	return TableNameSmsHomeSubjectSpu
}
