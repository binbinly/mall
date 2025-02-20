// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"
)

const TableNameAdminUser = "admin_users"

// AdminUser mapped from table <admin_users>
type AdminUser struct {
	ID            int64     `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	Username      string    `gorm:"column:username;not null" json:"username"`
	Password      string    `gorm:"column:password;not null" json:"password"`
	Name          string    `gorm:"column:name;not null" json:"name"`
	Avatar        string    `gorm:"column:avatar" json:"avatar"`
	RememberToken string    `gorm:"column:remember_token" json:"remember_token"`
	CreatedAt     time.Time `gorm:"column:created_at" json:"created_at"`
	UpdatedAt     time.Time `gorm:"column:updated_at" json:"updated_at"`
}

// TableName AdminUser's table name
func (*AdminUser) TableName() string {
	return TableNameAdminUser
}
