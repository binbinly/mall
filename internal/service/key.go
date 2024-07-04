package service

import "fmt"

const (
	// _userPrefix 用户令牌前缀
	_userPrefix = "user:token:%d"
)

// BuildUserTokenKey 用户令牌键
func BuildUserTokenKey(uid int) string {
	return fmt.Sprintf(_userPrefix, uid)
}
