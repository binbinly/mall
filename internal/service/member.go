package service

import (
	"context"
	"github.com/binbinly/gin-pkg/app"
	"github.com/binbinly/pkg/auth"
	"project-layout/internal/model"
	"time"

	"github.com/pkg/errors"
)

// MemberRegister 注册用户
func (s *Service) MemberRegister(ctx context.Context, phone int64, username, password string) error {
	u := &model.UmsMember{
		Username: username,
		Password: password,
		Phone:    phone,
		Status:   1,
	}
	exist, err := s.repo.MemberExist(ctx, username, phone)
	if err != nil {
		return errors.Wrapf(err, "[service.member] member exist")
	}
	if exist {
		return ErrMemberExisted
	}
	if err = s.repo.MemberCreate(ctx, u); err != nil {
		return errors.Wrapf(err, "[service.member] create member")
	}

	return nil
}

// UsernameLogin 用户名密码登录
func (s *Service) UsernameLogin(ctx context.Context, username, password string) (*model.UmsMember, string, error) {
	// 如果是已经注册用户，则通过用户名获取用户信息
	member, err := s.repo.GetMemberByUsername(ctx, username)
	if err != nil {
		return nil, "", errors.Wrapf(err, "[service.member] from db by username: %l", username)
	}

	// 否则新建用户信息, 并取得用户信息
	if member.ID == 0 {
		return nil, "", ErrMemberNotFound
	}

	if member.Status != 1 {
		return nil, "", ErrMemberFrozen
	}

	// Compare the login password with the user password.
	if err = member.Compare(password); err != nil {
		return nil, "", ErrMemberNotMatch
	}

	token, err := s.generateToken(ctx, member.ID)
	if err != nil {
		return nil, "", err
	}

	return member, token, nil
}

// PhoneLogin 邮箱登录
func (s *Service) PhoneLogin(ctx context.Context, phone int64) (*model.UmsMember, string, error) {
	// 如果是已经注册用户，则通过手机号获取用户信息
	member, err := s.repo.GetMemberByPhone(ctx, phone)
	if err != nil {
		return nil, "", errors.Wrapf(err, "[service.member] from db by phone: %d", phone)
	}

	// 否则新建用户信息, 并取得用户信息
	if member.ID == 0 {
		return nil, "", ErrMemberNotFound
	}

	if member.Status != 1 {
		return nil, "", ErrMemberFrozen
	}

	token, err := s.generateToken(ctx, member.ID)
	if err != nil {
		return nil, "", err
	}

	return member, token, nil
}

// MemberEdit 修改会员信息
func (s *Service) MemberEdit(ctx context.Context, id int, um map[string]any) error {
	return s.repo.MemberUpdate(ctx, id, um)
}

// MemberEditPwd 修改用户密码
func (s *Service) MemberEditPwd(ctx context.Context, id int, oldPassword, password string) error {
	member, err := s.memberInfo(ctx, id)
	if err != nil {
		return err
	}

	if err = member.Compare(oldPassword); err != nil {
		return ErrMemberPwd
	}

	member.Password = password
	if err = s.repo.MemberSave(ctx, member); err != nil {
		return errors.Wrapf(err, "[service.member] save pwd by id:%v", id)
	}
	return nil
}

// MemberInfo 获取用户信息
func (s *Service) MemberInfo(ctx context.Context, id int) (*model.UmsMember, error) {
	return s.memberInfo(ctx, id)
}

// MemberLogout 用户登出
func (l *Service) MemberLogout(ctx context.Context, id int) error {
	if err := l.rdb.Del(ctx, BuildUserTokenKey(id)).Err(); err != nil {
		return err
	}

	return nil
}

// memberInfo 获取会员模型
func (s *Service) memberInfo(ctx context.Context, id int) (*model.UmsMember, error) {
	member, err := s.repo.GetMemberByID(ctx, id)
	if err != nil {
		return nil, errors.Wrapf(err, "[service.member] from repo by id: %d", id)
	}
	if member.ID == 0 {
		return nil, ErrMemberNotFound
	}
	if member.Status != 1 {
		return nil, ErrMemberFrozen
	}
	return member, nil
}

// generateToken 生成token
func (s *Service) generateToken(ctx context.Context, uid int) (string, error) {
	// 签名
	payload := map[string]any{"user_id": uid}
	token, err := auth.Sign(ctx, payload, app.Conf.JwtSecret, app.Conf.JwtTimeout)
	if err != nil {
		return "", errors.Wrapf(err, "[service.member] gen token sign")
	}

	// 设置新令牌，用户单点登录
	if err = s.rdb.Set(ctx, BuildUserTokenKey(uid), token, time.Duration(app.Conf.JwtTimeout)*time.Second).Err(); err != nil {
		return "", errors.Wrapf(err, "[service.member] set token to redis")
	}
	return token, nil
}
