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
	err := _gen_test_db.AutoMigrate(&model.WmsPurchaseDetail{})
	if err != nil {
		fmt.Printf("Error: AutoMigrate(&model.WmsPurchaseDetail{}) fail: %s", err)
	}
}

func Test_wmsPurchaseDetailQuery(t *testing.T) {
	wmsPurchaseDetail := newWmsPurchaseDetail(_gen_test_db)
	wmsPurchaseDetail = *wmsPurchaseDetail.As(wmsPurchaseDetail.TableName())
	_do := wmsPurchaseDetail.WithContext(context.Background()).Debug()

	primaryKey := field.NewString(wmsPurchaseDetail.TableName(), clause.PrimaryKey)
	_, err := _do.Unscoped().Where(primaryKey.IsNotNull()).Delete()
	if err != nil {
		t.Error("clean table <wms_purchase_detail> fail:", err)
		return
	}

	_, ok := wmsPurchaseDetail.GetFieldByName("")
	if ok {
		t.Error("GetFieldByName(\"\") from wmsPurchaseDetail success")
	}

	err = _do.Create(&model.WmsPurchaseDetail{})
	if err != nil {
		t.Error("create item in table <wms_purchase_detail> fail:", err)
	}

	err = _do.Save(&model.WmsPurchaseDetail{})
	if err != nil {
		t.Error("create item in table <wms_purchase_detail> fail:", err)
	}

	err = _do.CreateInBatches([]*model.WmsPurchaseDetail{{}, {}}, 10)
	if err != nil {
		t.Error("create item in table <wms_purchase_detail> fail:", err)
	}

	_, err = _do.Select(wmsPurchaseDetail.ALL).Take()
	if err != nil {
		t.Error("Take() on table <wms_purchase_detail> fail:", err)
	}

	_, err = _do.First()
	if err != nil {
		t.Error("First() on table <wms_purchase_detail> fail:", err)
	}

	_, err = _do.Last()
	if err != nil {
		t.Error("First() on table <wms_purchase_detail> fail:", err)
	}

	_, err = _do.Where(primaryKey.IsNotNull()).FindInBatch(10, func(tx gen.Dao, batch int) error { return nil })
	if err != nil {
		t.Error("FindInBatch() on table <wms_purchase_detail> fail:", err)
	}

	err = _do.Where(primaryKey.IsNotNull()).FindInBatches(&[]*model.WmsPurchaseDetail{}, 10, func(tx gen.Dao, batch int) error { return nil })
	if err != nil {
		t.Error("FindInBatches() on table <wms_purchase_detail> fail:", err)
	}

	_, err = _do.Select(wmsPurchaseDetail.ALL).Where(primaryKey.IsNotNull()).Order(primaryKey.Desc()).Find()
	if err != nil {
		t.Error("Find() on table <wms_purchase_detail> fail:", err)
	}

	_, err = _do.Distinct(primaryKey).Take()
	if err != nil {
		t.Error("select Distinct() on table <wms_purchase_detail> fail:", err)
	}

	_, err = _do.Select(wmsPurchaseDetail.ALL).Omit(primaryKey).Take()
	if err != nil {
		t.Error("Omit() on table <wms_purchase_detail> fail:", err)
	}

	_, err = _do.Group(primaryKey).Find()
	if err != nil {
		t.Error("Group() on table <wms_purchase_detail> fail:", err)
	}

	_, err = _do.Scopes(func(dao gen.Dao) gen.Dao { return dao.Where(primaryKey.IsNotNull()) }).Find()
	if err != nil {
		t.Error("Scopes() on table <wms_purchase_detail> fail:", err)
	}

	_, _, err = _do.FindByPage(0, 1)
	if err != nil {
		t.Error("FindByPage() on table <wms_purchase_detail> fail:", err)
	}

	_, err = _do.ScanByPage(&model.WmsPurchaseDetail{}, 0, 1)
	if err != nil {
		t.Error("ScanByPage() on table <wms_purchase_detail> fail:", err)
	}

	_, err = _do.Attrs(primaryKey).Assign(primaryKey).FirstOrInit()
	if err != nil {
		t.Error("FirstOrInit() on table <wms_purchase_detail> fail:", err)
	}

	_, err = _do.Attrs(primaryKey).Assign(primaryKey).FirstOrCreate()
	if err != nil {
		t.Error("FirstOrCreate() on table <wms_purchase_detail> fail:", err)
	}

	var _a _another
	var _aPK = field.NewString(_a.TableName(), "id")

	err = _do.Join(&_a, primaryKey.EqCol(_aPK)).Scan(map[string]interface{}{})
	if err != nil {
		t.Error("Join() on table <wms_purchase_detail> fail:", err)
	}

	err = _do.LeftJoin(&_a, primaryKey.EqCol(_aPK)).Scan(map[string]interface{}{})
	if err != nil {
		t.Error("LeftJoin() on table <wms_purchase_detail> fail:", err)
	}

	_, err = _do.Not().Or().Clauses().Take()
	if err != nil {
		t.Error("Not/Or/Clauses on table <wms_purchase_detail> fail:", err)
	}
}
