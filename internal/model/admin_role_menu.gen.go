// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"
)

const TableNameAdminRoleMenu = "admin_role_menu"

// AdminRoleMenu mapped from table <admin_role_menu>
type AdminRoleMenu struct {
	RoleID    int64     `gorm:"column:role_id;primaryKey" json:"role_id"`
	MenuID    int64     `gorm:"column:menu_id;primaryKey" json:"menu_id"`
	CreatedAt time.Time `gorm:"column:created_at" json:"created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at" json:"updated_at"`
}

// TableName AdminRoleMenu's table name
func (*AdminRoleMenu) TableName() string {
	return TableNameAdminRoleMenu
}
