// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package dal

import (
	"context"
	"fmt"
	"testing"

	"project-layout/internal/model"

	"gorm.io/gen"
	"gorm.io/gen/field"
	"gorm.io/gorm/clause"
)

func init() {
	InitializeDB()
	err := _gen_test_db.AutoMigrate(&model.SmsHomeAdv{})
	if err != nil {
		fmt.Printf("Error: AutoMigrate(&model.SmsHomeAdv{}) fail: %s", err)
	}
}

func Test_smsHomeAdvQuery(t *testing.T) {
	smsHomeAdv := newSmsHomeAdv(_gen_test_db)
	smsHomeAdv = *smsHomeAdv.As(smsHomeAdv.TableName())
	_do := smsHomeAdv.WithContext(context.Background()).Debug()

	primaryKey := field.NewString(smsHomeAdv.TableName(), clause.PrimaryKey)
	_, err := _do.Unscoped().Where(primaryKey.IsNotNull()).Delete()
	if err != nil {
		t.Error("clean table <sms_home_adv> fail:", err)
		return
	}

	_, ok := smsHomeAdv.GetFieldByName("")
	if ok {
		t.Error("GetFieldByName(\"\") from smsHomeAdv success")
	}

	err = _do.Create(&model.SmsHomeAdv{})
	if err != nil {
		t.Error("create item in table <sms_home_adv> fail:", err)
	}

	err = _do.Save(&model.SmsHomeAdv{})
	if err != nil {
		t.Error("create item in table <sms_home_adv> fail:", err)
	}

	err = _do.CreateInBatches([]*model.SmsHomeAdv{{}, {}}, 10)
	if err != nil {
		t.Error("create item in table <sms_home_adv> fail:", err)
	}

	_, err = _do.Select(smsHomeAdv.ALL).Take()
	if err != nil {
		t.Error("Take() on table <sms_home_adv> fail:", err)
	}

	_, err = _do.First()
	if err != nil {
		t.Error("First() on table <sms_home_adv> fail:", err)
	}

	_, err = _do.Last()
	if err != nil {
		t.Error("First() on table <sms_home_adv> fail:", err)
	}

	_, err = _do.Where(primaryKey.IsNotNull()).FindInBatch(10, func(tx gen.Dao, batch int) error { return nil })
	if err != nil {
		t.Error("FindInBatch() on table <sms_home_adv> fail:", err)
	}

	err = _do.Where(primaryKey.IsNotNull()).FindInBatches(&[]*model.SmsHomeAdv{}, 10, func(tx gen.Dao, batch int) error { return nil })
	if err != nil {
		t.Error("FindInBatches() on table <sms_home_adv> fail:", err)
	}

	_, err = _do.Select(smsHomeAdv.ALL).Where(primaryKey.IsNotNull()).Order(primaryKey.Desc()).Find()
	if err != nil {
		t.Error("Find() on table <sms_home_adv> fail:", err)
	}

	_, err = _do.Distinct(primaryKey).Take()
	if err != nil {
		t.Error("select Distinct() on table <sms_home_adv> fail:", err)
	}

	_, err = _do.Select(smsHomeAdv.ALL).Omit(primaryKey).Take()
	if err != nil {
		t.Error("Omit() on table <sms_home_adv> fail:", err)
	}

	_, err = _do.Group(primaryKey).Find()
	if err != nil {
		t.Error("Group() on table <sms_home_adv> fail:", err)
	}

	_, err = _do.Scopes(func(dao gen.Dao) gen.Dao { return dao.Where(primaryKey.IsNotNull()) }).Find()
	if err != nil {
		t.Error("Scopes() on table <sms_home_adv> fail:", err)
	}

	_, _, err = _do.FindByPage(0, 1)
	if err != nil {
		t.Error("FindByPage() on table <sms_home_adv> fail:", err)
	}

	_, err = _do.ScanByPage(&model.SmsHomeAdv{}, 0, 1)
	if err != nil {
		t.Error("ScanByPage() on table <sms_home_adv> fail:", err)
	}

	_, err = _do.Attrs(primaryKey).Assign(primaryKey).FirstOrInit()
	if err != nil {
		t.Error("FirstOrInit() on table <sms_home_adv> fail:", err)
	}

	_, err = _do.Attrs(primaryKey).Assign(primaryKey).FirstOrCreate()
	if err != nil {
		t.Error("FirstOrCreate() on table <sms_home_adv> fail:", err)
	}

	var _a _another
	var _aPK = field.NewString(_a.TableName(), "id")

	err = _do.Join(&_a, primaryKey.EqCol(_aPK)).Scan(map[string]interface{}{})
	if err != nil {
		t.Error("Join() on table <sms_home_adv> fail:", err)
	}

	err = _do.LeftJoin(&_a, primaryKey.EqCol(_aPK)).Scan(map[string]interface{}{})
	if err != nil {
		t.Error("LeftJoin() on table <sms_home_adv> fail:", err)
	}

	_, err = _do.Not().Or().Clauses().Take()
	if err != nil {
		t.Error("Not/Or/Clauses on table <sms_home_adv> fail:", err)
	}
}
