// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

const TableNameMigration = "migrations"

// Migration mapped from table <migrations>
type Migration struct {
	ID        int    `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	Migration string `gorm:"column:migration;not null" json:"migration"`
	Batch     int    `gorm:"column:batch;not null" json:"batch"`
}

// TableName Migration's table name
func (*Migration) TableName() string {
	return TableNameMigration
}
