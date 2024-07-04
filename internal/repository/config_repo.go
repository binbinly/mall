package repository

import (
	"context"
	"encoding/json"
	"fmt"
	"gorm.io/gorm"
	"project-layout/internal/model"

	"github.com/pkg/errors"
)

type IConfig interface {
	GetNoticeList(ctx context.Context, offset, limit int) (list []*model.SmsAppNotice, err error)
	AppPageData(ctx context.Context, page int) (list []*model.SmsAppSetting, err error)
	AppHomePageData(ctx context.Context, catID int) (list []*model.SmsAppSetting, err error)
	GetConfigByName(ctx context.Context, name string, v any) (err error)
}

// GetNoticeList 公告列表
func (r *Repo) GetNoticeList(ctx context.Context, offset, limit int) (list []*model.SmsAppNotice, err error) {
	if err = r.db.WithContext(ctx).Scopes(model.OffsetPage(offset, limit)).Order(model.DefaultOrder).Find(&list).Error; err != nil {
		return nil, errors.Wrap(err, "[repo.notice] db find")
	}
	return
}

// AppPageData app其他页面配置数据
func (r *Repo) AppPageData(ctx context.Context, page int) (list []*model.SmsAppSetting, err error) {
	doKey := fmt.Sprintf("app_page:%d", page)
	if err = r.QueryCache(ctx, doKey, &list, 0, func() (any, error) {
		// 从数据库中获取
		if err = r.db.WithContext(ctx).Model(&model.SmsAppSetting{}).
			Where("page=?", page).Order(model.DefaultOrderSort).Scan(&list).Error; err != nil {
			return nil, err
		}
		if len(list) == 0 {
			return nil, gorm.ErrEmptySlice
		}
		return list, nil
	}); err != nil {
		return nil, errors.Wrapf(err, "[repo.market] query cache")
	}

	return
}

// AppHomePageData app首页配置数据
func (r *Repo) AppHomePageData(ctx context.Context, catID int) (list []*model.SmsAppSetting, err error) {
	doKey := fmt.Sprintf("app_page:%d_%d", model.AppPageHome, catID)
	if err = r.QueryCache(ctx, doKey, &list, 0, func() (any, error) {
		// 从数据库中获取
		if err = r.db.WithContext(ctx).Model(&model.SmsAppSetting{}).
			Where("page=? and cat_id=?", model.AppPageHome, catID).
			Order(model.DefaultOrderSort).Scan(&list).Error; err != nil {
			return nil, err
		}
		if len(list) == 0 {
			return nil, gorm.ErrEmptySlice
		}
		return list, nil
	}); err != nil {
		return nil, errors.Wrapf(err, "[repo.market] query cache")
	}

	return
}

// GetConfigByName 获取配置
func (r *Repo) GetConfigByName(ctx context.Context, name string, v any) (err error) {
	doKey := "config:" + name
	var config *model.SmsConfig
	if err = r.QueryCache(ctx, doKey, &config, 0, func() (any, error) {
		// 从数据库中获取
		if err = r.db.WithContext(ctx).Model(&model.SmsConfig{}).
			Where("name=?", name).First(&config).Error; err != nil {
			return nil, err
		}
		return config, nil
	}); err != nil {
		return errors.Wrapf(err, "[repo.market] query cache")
	}

	if config.Value == "" {
		return nil
	}
	if v != nil {
		if err = json.Unmarshal([]byte(config.Value), &v); err != nil {
			return errors.Wrapf(err, "[repo.market] json unmarshal by value: %v", config.Value)
		}
	} else {
		v = config.Value
	}
	return nil
}
