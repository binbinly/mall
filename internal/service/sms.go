package service

import (
	"context"
	"fmt"
	"github.com/binbinly/gin-pkg/app"
	"github.com/binbinly/pkg/logger"
	"strconv"
	"time"

	"github.com/pkg/errors"
	"github.com/redis/go-redis/v9"

	"math/rand"
)

const (
	_verifyCodeRedisKey = "app:vcode:%v"   // 验证码key
	_maxDurationTime    = 10 * time.Minute // 验证码有效期
)

// 限制规则
var rules = []*rule{{
	Count: 1,
	TTL:   60,
	Key:   "sms_limit:rule_minute:",
	Err:   ErrVerifyCodeRuleMinute,
}, {
	Count: 5,
	TTL:   3600,
	Key:   "sms_limit:rule_hour:",
	Err:   ErrVerifyCodeRuleHour,
}, {
	Count: 10,
	TTL:   86400,
	Key:   "sms_limit:rule_day:",
	Err:   ErrVerifyCodeRuleDay,
}}

// 限制规则
type rule struct {
	Count int           //限制次数
	TTL   time.Duration //限制时间
	Key   string        //key
	Err   error         //错误
}

// SendSMS 发送短信
func (s *Service) SendSMS(ctx context.Context, phone string) (string, error) {
	code, err := s.genVCode(ctx, phone)
	if err != nil {
		return "", err
	}
	if app.Conf.Sms.IsReal { // 真实发送短信场景下的限流控制
		if err = s.checkRules(ctx, phone); err != nil {
			return "", err
		}

		if err = s.realSend(phone); err != nil {
			return "", err
		}
		s.execRules(ctx, phone)

		return "", nil
	}

	return code, nil
}

// CheckVCode 验证校验码是否正确
func (s *Service) CheckVCode(ctx context.Context, phone int64, vCode string) error {
	oldVCode, err := s.getVCode(ctx, phone)
	if err != nil {
		return errors.Wrapf(err, "[third.sms] get verify code")
	}

	if vCode != oldVCode {
		return ErrVerifyCodeNotMatch
	}

	return nil
}

// checkRules 验证规则
func (s *Service) checkRules(ctx context.Context, phone string) (err error) {
	if !app.Conf.Sms.IsReal { // 非真实发送短信场景无需验证
		return nil
	}

	var r string
	for _, v := range rules {
		r, err = s.rdb.Get(ctx, v.Key+phone).Result()
		if errors.Is(err, redis.Nil) {
			return nil
		} else if err != nil {
			return errors.Wrap(err, "[third.sms] redis get rule err")
		}
		num, err := strconv.Atoi(r)
		if err != nil {
			return errors.Wrapf(err, "[third.sms] atoi r:%v", r)
		}
		if num >= v.Count {
			return v.Err
		}
	}
	return nil
}

// execRules 发送成功,执行限流规则
func (s *Service) execRules(ctx context.Context, phone string) {
	if !app.Conf.Sms.IsReal {
		return
	}
	var err error
	for _, v := range rules {
		pipe := s.rdb.Pipeline()
		pipe.Incr(ctx, v.Key+phone)
		pipe.Expire(ctx, v.Key+phone, v.TTL*time.Second)
		if _, err = pipe.Exec(ctx); err != nil {
			logger.Warnf("[third.sms] redis pipe exec err:%v", err)
		}
	}
}

// realSend 真实发送短信
func (s *Service) realSend(phone string) error {
	//TODO 调用第三方短信接口
	return nil
}

// genVCode 生成验证码
func (s *Service) genVCode(ctx context.Context, phone string) (string, error) {
	// step1: 生成随机数
	vCodeStr := fmt.Sprintf("%06v", rand.New(rand.NewSource(time.Now().UnixNano())).Int31n(1000000))

	// step2: 写入到redis里
	// 使用set, key使用前缀+手机号 缓存10分钟）
	key := fmt.Sprintf(_verifyCodeRedisKey, phone)
	if err := s.rdb.Set(ctx, key, vCodeStr, _maxDurationTime).Err(); err != nil {
		return "", errors.Wrap(err, "[third.sms] redis set verify code err")
	}

	return vCodeStr, nil
}

// getVCode 获取验证码
func (s *Service) getVCode(ctx context.Context, phone int64) (string, error) {
	// 直接从redis里获取
	key := fmt.Sprintf(_verifyCodeRedisKey, phone)
	verifyCode, err := s.rdb.Get(ctx, key).Result()
	if errors.Is(err, redis.Nil) {
		return "", nil
	} else if err != nil {
		return "", errors.Wrap(err, "[third.sms] redis get verify code err")
	}

	return verifyCode, nil
}
