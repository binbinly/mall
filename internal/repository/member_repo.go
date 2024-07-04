package repository

import (
	"context"
	"gorm.io/gorm"
	"project-layout/internal/model"

	"github.com/pkg/errors"
)

// IMember 会员接口定义
type IMember interface {
	MemberCreate(ctx context.Context, member *model.UmsMember) error
	MemberUpdate(ctx context.Context, id int, um map[string]any) error
	MemberSave(ctx context.Context, member *model.UmsMember) error
	GetMemberByID(ctx context.Context, id int) (member *model.UmsMember, err error)
	GetMemberByUsername(ctx context.Context, username string) (*model.UmsMember, error)
	GetMemberByPhone(ctx context.Context, phone int64) (*model.UmsMember, error)
	MemberExist(ctx context.Context, username string, phone int64) (bool, error)
	MemberDelete(ctx context.Context, member *model.UmsMember) error
}

// MemberCreate 创建用户
func (r *Repo) MemberCreate(ctx context.Context, member *model.UmsMember) (err error) {
	if err = r.UmsMember.WithContext(ctx).Create(member); err != nil {
		return errors.Wrap(err, "[repo.member] Create err")
	}
	r.DelCache(ctx, buildMemberCacheKey(member.ID))

	return nil
}

// MemberUpdate 更新用户信息
func (r *Repo) MemberUpdate(ctx context.Context, id int, um map[string]any) error {
	if _, err := r.UmsMember.WithContext(ctx).Where(r.UmsMember.ID.Eq(id)).Updates(um); err != nil {
		return errors.Wrapf(err, "[repo.member] updates")
	}
	r.DelCache(ctx, buildMemberCacheKey(id))

	return nil
}

// MemberSave 保存
func (r *Repo) MemberSave(ctx context.Context, member *model.UmsMember) error {
	if err := r.UmsMember.WithContext(ctx).Save(member); err != nil {
		return errors.Wrapf(err, "[repo.member] save")
	}
	r.DelCache(ctx, buildMemberCacheKey(member.ID))

	return nil
}

// GetMemberByID 获取用户
func (r *Repo) GetMemberByID(ctx context.Context, id int) (member *model.UmsMember, err error) {
	if err = r.QueryCache(ctx, buildMemberCacheKey(id), &member, 0, func() (any, error) {
		// 从数据库中获取
		data, err := r.UmsMember.WithContext(ctx).Where(r.UmsMember.ID.Eq(id)).First()
		if err != nil {
			return nil, err
		}
		return data, nil
	}); err != nil {
		return nil, errors.Wrapf(err, "[repo.member] query cache")
	}

	return member, nil
}

// GetMemberByUsername 根据账号获取用户
func (r *Repo) GetMemberByUsername(ctx context.Context, username string) (*model.UmsMember, error) {
	member, err := r.UmsMember.WithContext(ctx).Where(r.UmsMember.Username.Eq(username)).First()
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, errors.Wrapf(err, "[repo.member] get member by username %v: ", username)
	}

	return member, nil
}

// GetMemberByPhone 根据手机号获取用户
func (r *Repo) GetMemberByPhone(ctx context.Context, phone int64) (*model.UmsMember, error) {
	member, err := r.UmsMember.WithContext(ctx).Where(r.UmsMember.Phone.Eq(phone)).First()
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, errors.Wrapf(err, "[repo.member] get member by phone %v: ", phone)
	}

	return member, nil
}

// MemberExist 用户是否已存在
func (r *Repo) MemberExist(ctx context.Context, username string, phone int64) (bool, error) {
	c, err := r.UmsMember.WithContext(ctx).Where(r.UmsMember.Phone.Eq(phone)).Or(r.UmsMember.Username.Eq(username)).Count()
	if err != nil {
		return false, errors.Wrapf(err, "[repo.member] username %v or phone %v does not exist", username, phone)
	}
	return c > 0, nil
}

// MemberDelete 用户删除
func (r *Repo) MemberDelete(ctx context.Context, member *model.UmsMember) error {
	if _, err := r.UmsMember.WithContext(ctx).Delete(member); err != nil {
		return errors.Wrapf(err, "[repo.member] delete")
	}
	r.DelCache(ctx, buildMemberCacheKey(member.ID))

	return nil
}
