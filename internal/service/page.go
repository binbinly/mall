package service

import (
	"context"
	"project-layout/internal/app"
	"project-layout/internal/model"

	"github.com/pkg/errors"
	"github.com/redis/go-redis/v9"
)

// GetHomeCat 首页分类
func (s *Service) GetHomeCat(ctx context.Context) ([]*model.ConfigHomeCat, error) {
	var cats []*model.ConfigHomeCat
	if err := s.repo.GetConfigByName(ctx, model.ConfigKeyHomeCat, &cats); err != nil {
		return nil, errors.Wrapf(err, "[service.page] get home cat")
	}

	return cats, nil
}

// GetHomeCatData 首页分类下的配置数据
func (s *Service) GetHomeCatData(ctx context.Context, cid int) ([]*model.SmsAppSetting, error) {
	data, err := s.repo.AppHomePageData(ctx, cid)
	if err != nil {
		return nil, errors.Wrapf(err, "[service.page] home data by cid: %v", cid)
	}
	return data, nil
}

// GetNoticeList 公告列表
func (s *Service) GetNoticeList(ctx context.Context, page int) ([]*model.SmsAppNotice, error) {
	list, err := s.repo.GetNoticeList(ctx, app.GetPageOffset(page), app.DefaultLimit)
	if err != nil {
		return nil, errors.Wrapf(err, "[service.page] notice list page: %v", page)
	}
	return list, nil
}

// GetSearchData 搜索页配置数据
func (s *Service) GetSearchData(ctx context.Context) ([]*model.SmsAppSetting, []string, error) {
	data, err := s.repo.AppPageData(ctx, model.AppPageSearch)
	if err != nil {
		return nil, nil, errors.Wrap(err, "[service.page] search data")
	}
	// 搜索热词
	hot, err := s.rdb.ZRevRangeByScore(ctx, "search_hot", &redis.ZRangeBy{
		Min:    "-inf",
		Max:    "+inf",
		Offset: 0,
		Count:  10,
	}).Result()
	if err != nil {
		return nil, nil, errors.Wrap(err, "[service.page] hot keyword by redis")
	}
	return data, hot, nil
}

// GetPayConfig 支付配置
func (s *Service) GetPayConfig(ctx context.Context) ([]*model.ConfigPayList, error) {
	var pays []*model.ConfigPayList
	if err := s.repo.GetConfigByName(ctx, model.ConfigKeyPayList, &pays); err != nil {
		return nil, errors.Wrapf(err, "[service.page] get pay config")
	}
	return pays, nil
}
