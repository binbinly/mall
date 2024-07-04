package repository

import "fmt"

func BuildCartKey(uid int) string {
	return fmt.Sprintf("mall:cart:%v", uid)
}

// buildMemberCacheKey 构建会员缓存key
func buildMemberCacheKey(id int) string {
	return fmt.Sprintf("mall:member:%d", id)
}

// buildAddressCacheKey 构建收货地址缓存键
func buildAddressCacheKey(id int) string {
	return fmt.Sprintf("mall:member_address:%d", id)
}

// buildAddressAllCacheKey 当前用户所有收货地址键
func buildAddressAllCacheKey(uid int) string {
	return fmt.Sprintf("mall:member_address_all:%d", uid)
}

// buildBrandCacheKey 构建品牌模型缓存键
func buildBrandCacheKey(id int) string {
	return fmt.Sprintf("mall:brand:%d", id)
}

// buildCategoryCacheKey 构建分类模型缓存键
func buildCategoryCacheKey(id int) string {
	return fmt.Sprintf("mall:category:%d", id)
}

func buildSkuCacheKey(id int) string {
	return fmt.Sprintf("mall:sku:%d", id)
}

func stockCacheKey(skuID int) string {
	return fmt.Sprintf("mall:sku_stock:%d", skuID)
}
