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
	err := _gen_test_db.AutoMigrate(&model.UmsMemberStat{})
	if err != nil {
		fmt.Printf("Error: AutoMigrate(&model.UmsMemberStat{}) fail: %s", err)
	}
}

func Test_umsMemberStatQuery(t *testing.T) {
	umsMemberStat := newUmsMemberStat(_gen_test_db)
	umsMemberStat = *umsMemberStat.As(umsMemberStat.TableName())
	_do := umsMemberStat.WithContext(context.Background()).Debug()

	primaryKey := field.NewString(umsMemberStat.TableName(), clause.PrimaryKey)
	_, err := _do.Unscoped().Where(primaryKey.IsNotNull()).Delete()
	if err != nil {
		t.Error("clean table <ums_member_stat> fail:", err)
		return
	}

	_, ok := umsMemberStat.GetFieldByName("")
	if ok {
		t.Error("GetFieldByName(\"\") from umsMemberStat success")
	}

	err = _do.Create(&model.UmsMemberStat{})
	if err != nil {
		t.Error("create item in table <ums_member_stat> fail:", err)
	}

	err = _do.Save(&model.UmsMemberStat{})
	if err != nil {
		t.Error("create item in table <ums_member_stat> fail:", err)
	}

	err = _do.CreateInBatches([]*model.UmsMemberStat{{}, {}}, 10)
	if err != nil {
		t.Error("create item in table <ums_member_stat> fail:", err)
	}

	_, err = _do.Select(umsMemberStat.ALL).Take()
	if err != nil {
		t.Error("Take() on table <ums_member_stat> fail:", err)
	}

	_, err = _do.First()
	if err != nil {
		t.Error("First() on table <ums_member_stat> fail:", err)
	}

	_, err = _do.Last()
	if err != nil {
		t.Error("First() on table <ums_member_stat> fail:", err)
	}

	_, err = _do.Where(primaryKey.IsNotNull()).FindInBatch(10, func(tx gen.Dao, batch int) error { return nil })
	if err != nil {
		t.Error("FindInBatch() on table <ums_member_stat> fail:", err)
	}

	err = _do.Where(primaryKey.IsNotNull()).FindInBatches(&[]*model.UmsMemberStat{}, 10, func(tx gen.Dao, batch int) error { return nil })
	if err != nil {
		t.Error("FindInBatches() on table <ums_member_stat> fail:", err)
	}

	_, err = _do.Select(umsMemberStat.ALL).Where(primaryKey.IsNotNull()).Order(primaryKey.Desc()).Find()
	if err != nil {
		t.Error("Find() on table <ums_member_stat> fail:", err)
	}

	_, err = _do.Distinct(primaryKey).Take()
	if err != nil {
		t.Error("select Distinct() on table <ums_member_stat> fail:", err)
	}

	_, err = _do.Select(umsMemberStat.ALL).Omit(primaryKey).Take()
	if err != nil {
		t.Error("Omit() on table <ums_member_stat> fail:", err)
	}

	_, err = _do.Group(primaryKey).Find()
	if err != nil {
		t.Error("Group() on table <ums_member_stat> fail:", err)
	}

	_, err = _do.Scopes(func(dao gen.Dao) gen.Dao { return dao.Where(primaryKey.IsNotNull()) }).Find()
	if err != nil {
		t.Error("Scopes() on table <ums_member_stat> fail:", err)
	}

	_, _, err = _do.FindByPage(0, 1)
	if err != nil {
		t.Error("FindByPage() on table <ums_member_stat> fail:", err)
	}

	_, err = _do.ScanByPage(&model.UmsMemberStat{}, 0, 1)
	if err != nil {
		t.Error("ScanByPage() on table <ums_member_stat> fail:", err)
	}

	_, err = _do.Attrs(primaryKey).Assign(primaryKey).FirstOrInit()
	if err != nil {
		t.Error("FirstOrInit() on table <ums_member_stat> fail:", err)
	}

	_, err = _do.Attrs(primaryKey).Assign(primaryKey).FirstOrCreate()
	if err != nil {
		t.Error("FirstOrCreate() on table <ums_member_stat> fail:", err)
	}

	var _a _another
	var _aPK = field.NewString(_a.TableName(), "id")

	err = _do.Join(&_a, primaryKey.EqCol(_aPK)).Scan(map[string]interface{}{})
	if err != nil {
		t.Error("Join() on table <ums_member_stat> fail:", err)
	}

	err = _do.LeftJoin(&_a, primaryKey.EqCol(_aPK)).Scan(map[string]interface{}{})
	if err != nil {
		t.Error("LeftJoin() on table <ums_member_stat> fail:", err)
	}

	_, err = _do.Not().Or().Clauses().Take()
	if err != nil {
		t.Error("Not/Or/Clauses on table <ums_member_stat> fail:", err)
	}
}
