package app

// 用来定义一些公共的常量
const (
	// DefaultLimit 默认分页数
	DefaultLimit = 20
)

// GetPageOffset 获取分页起始偏移量
func GetPageOffset(page int) int {
	var offset int
	if page > 0 {
		offset = (page - 1) * DefaultLimit
	}
	return offset
}
