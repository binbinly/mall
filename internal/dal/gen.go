// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package dal

import (
	"context"
	"database/sql"

	"gorm.io/gorm"

	"gorm.io/gen"

	"gorm.io/plugin/dbresolver"
)

var (
	Q                     = new(Query)
	AdminExceptionLog     *adminExceptionLog
	AdminExtension        *adminExtension
	AdminExtensionHistory *adminExtensionHistory
	AdminMenu             *adminMenu
	AdminOperationLog     *adminOperationLog
	AdminPermission       *adminPermission
	AdminPermissionMenu   *adminPermissionMenu
	AdminRole             *adminRole
	AdminRoleMenu         *adminRoleMenu
	AdminRolePermission   *adminRolePermission
	AdminRoleUser         *adminRoleUser
	AdminSetting          *adminSetting
	AdminUser             *adminUser
	FailedJob             *failedJob
	Migration             *migration
	OmsOrder              *omsOrder
	OmsOrderBill          *omsOrderBill
	OmsOrderItem          *omsOrderItem
	OmsOrderOperate       *omsOrderOperate
	OmsOrderRefund        *omsOrderRefund
	OmsOrderReturnApply   *omsOrderReturnApply
	OmsOrderReturnReason  *omsOrderReturnReason
	OmsPayment            *omsPayment
	PmsAttr               *pmsAttr
	PmsAttrGroup          *pmsAttrGroup
	PmsAttrRelGroup       *pmsAttrRelGroup
	PmsAttrValue          *pmsAttrValue
	PmsBrand              *pmsBrand
	PmsCategory           *pmsCategory
	PmsCategoryRelBrand   *pmsCategoryRelBrand
	PmsSku                *pmsSku
	PmsSkuAttr            *pmsSkuAttr
	PmsSpu                *pmsSpu
	PmsSpuComment         *pmsSpuComment
	PmsSpuDesc            *pmsSpuDesc
	SmsAgreement          *smsAgreement
	SmsAppNotice          *smsAppNotice
	SmsAppSetting         *smsAppSetting
	SmsAppVersion         *smsAppVersion
	SmsConfig             *smsConfig
	SmsCoupon             *smsCoupon
	SmsCouponMember       *smsCouponMember
	SmsCouponRelCat       *smsCouponRelCat
	SmsCouponRelSpu       *smsCouponRelSpu
	SmsHomeAdv            *smsHomeAdv
	SmsHomeSubject        *smsHomeSubject
	SmsHomeSubjectSpu     *smsHomeSubjectSpu
	SmsMemberPrice        *smsMemberPrice
	SmsSeckillActivity    *smsSeckillActivity
	SmsSeckillSession     *smsSeckillSession
	SmsSeckillSku         *smsSeckillSku
	SmsSeckillSkuNotice   *smsSeckillSkuNotice
	SmsSkuFullReduction   *smsSkuFullReduction
	SmsSkuLadder          *smsSkuLadder
	SmsSpuBound           *smsSpuBound
	SysCity               *sysCity
	SysCityArea           *sysCityArea
	SysExpress            *sysExpress
	UmsCollectSpu         *umsCollectSpu
	UmsCollectSubject     *umsCollectSubject
	UmsGrowthLog          *umsGrowthLog
	UmsIntegrationLog     *umsIntegrationLog
	UmsMember             *umsMember
	UmsMemberAddress      *umsMemberAddress
	UmsMemberLevel        *umsMemberLevel
	UmsMemberLog          *umsMemberLog
	UmsMemberStat         *umsMemberStat
	WmsPurchase           *wmsPurchase
	WmsPurchaseDetail     *wmsPurchaseDetail
	WmsWare               *wmsWare
	WmsWareSku            *wmsWareSku
	WmsWareTask           *wmsWareTask
	WmsWareTaskDetail     *wmsWareTaskDetail
)

func SetDefault(db *gorm.DB, opts ...gen.DOOption) {
	*Q = *Use(db, opts...)
	AdminExceptionLog = &Q.AdminExceptionLog
	AdminExtension = &Q.AdminExtension
	AdminExtensionHistory = &Q.AdminExtensionHistory
	AdminMenu = &Q.AdminMenu
	AdminOperationLog = &Q.AdminOperationLog
	AdminPermission = &Q.AdminPermission
	AdminPermissionMenu = &Q.AdminPermissionMenu
	AdminRole = &Q.AdminRole
	AdminRoleMenu = &Q.AdminRoleMenu
	AdminRolePermission = &Q.AdminRolePermission
	AdminRoleUser = &Q.AdminRoleUser
	AdminSetting = &Q.AdminSetting
	AdminUser = &Q.AdminUser
	FailedJob = &Q.FailedJob
	Migration = &Q.Migration
	OmsOrder = &Q.OmsOrder
	OmsOrderBill = &Q.OmsOrderBill
	OmsOrderItem = &Q.OmsOrderItem
	OmsOrderOperate = &Q.OmsOrderOperate
	OmsOrderRefund = &Q.OmsOrderRefund
	OmsOrderReturnApply = &Q.OmsOrderReturnApply
	OmsOrderReturnReason = &Q.OmsOrderReturnReason
	OmsPayment = &Q.OmsPayment
	PmsAttr = &Q.PmsAttr
	PmsAttrGroup = &Q.PmsAttrGroup
	PmsAttrRelGroup = &Q.PmsAttrRelGroup
	PmsAttrValue = &Q.PmsAttrValue
	PmsBrand = &Q.PmsBrand
	PmsCategory = &Q.PmsCategory
	PmsCategoryRelBrand = &Q.PmsCategoryRelBrand
	PmsSku = &Q.PmsSku
	PmsSkuAttr = &Q.PmsSkuAttr
	PmsSpu = &Q.PmsSpu
	PmsSpuComment = &Q.PmsSpuComment
	PmsSpuDesc = &Q.PmsSpuDesc
	SmsAgreement = &Q.SmsAgreement
	SmsAppNotice = &Q.SmsAppNotice
	SmsAppSetting = &Q.SmsAppSetting
	SmsAppVersion = &Q.SmsAppVersion
	SmsConfig = &Q.SmsConfig
	SmsCoupon = &Q.SmsCoupon
	SmsCouponMember = &Q.SmsCouponMember
	SmsCouponRelCat = &Q.SmsCouponRelCat
	SmsCouponRelSpu = &Q.SmsCouponRelSpu
	SmsHomeAdv = &Q.SmsHomeAdv
	SmsHomeSubject = &Q.SmsHomeSubject
	SmsHomeSubjectSpu = &Q.SmsHomeSubjectSpu
	SmsMemberPrice = &Q.SmsMemberPrice
	SmsSeckillActivity = &Q.SmsSeckillActivity
	SmsSeckillSession = &Q.SmsSeckillSession
	SmsSeckillSku = &Q.SmsSeckillSku
	SmsSeckillSkuNotice = &Q.SmsSeckillSkuNotice
	SmsSkuFullReduction = &Q.SmsSkuFullReduction
	SmsSkuLadder = &Q.SmsSkuLadder
	SmsSpuBound = &Q.SmsSpuBound
	SysCity = &Q.SysCity
	SysCityArea = &Q.SysCityArea
	SysExpress = &Q.SysExpress
	UmsCollectSpu = &Q.UmsCollectSpu
	UmsCollectSubject = &Q.UmsCollectSubject
	UmsGrowthLog = &Q.UmsGrowthLog
	UmsIntegrationLog = &Q.UmsIntegrationLog
	UmsMember = &Q.UmsMember
	UmsMemberAddress = &Q.UmsMemberAddress
	UmsMemberLevel = &Q.UmsMemberLevel
	UmsMemberLog = &Q.UmsMemberLog
	UmsMemberStat = &Q.UmsMemberStat
	WmsPurchase = &Q.WmsPurchase
	WmsPurchaseDetail = &Q.WmsPurchaseDetail
	WmsWare = &Q.WmsWare
	WmsWareSku = &Q.WmsWareSku
	WmsWareTask = &Q.WmsWareTask
	WmsWareTaskDetail = &Q.WmsWareTaskDetail
}

func Use(db *gorm.DB, opts ...gen.DOOption) *Query {
	return &Query{
		db:                    db,
		AdminExceptionLog:     newAdminExceptionLog(db, opts...),
		AdminExtension:        newAdminExtension(db, opts...),
		AdminExtensionHistory: newAdminExtensionHistory(db, opts...),
		AdminMenu:             newAdminMenu(db, opts...),
		AdminOperationLog:     newAdminOperationLog(db, opts...),
		AdminPermission:       newAdminPermission(db, opts...),
		AdminPermissionMenu:   newAdminPermissionMenu(db, opts...),
		AdminRole:             newAdminRole(db, opts...),
		AdminRoleMenu:         newAdminRoleMenu(db, opts...),
		AdminRolePermission:   newAdminRolePermission(db, opts...),
		AdminRoleUser:         newAdminRoleUser(db, opts...),
		AdminSetting:          newAdminSetting(db, opts...),
		AdminUser:             newAdminUser(db, opts...),
		FailedJob:             newFailedJob(db, opts...),
		Migration:             newMigration(db, opts...),
		OmsOrder:              newOmsOrder(db, opts...),
		OmsOrderBill:          newOmsOrderBill(db, opts...),
		OmsOrderItem:          newOmsOrderItem(db, opts...),
		OmsOrderOperate:       newOmsOrderOperate(db, opts...),
		OmsOrderRefund:        newOmsOrderRefund(db, opts...),
		OmsOrderReturnApply:   newOmsOrderReturnApply(db, opts...),
		OmsOrderReturnReason:  newOmsOrderReturnReason(db, opts...),
		OmsPayment:            newOmsPayment(db, opts...),
		PmsAttr:               newPmsAttr(db, opts...),
		PmsAttrGroup:          newPmsAttrGroup(db, opts...),
		PmsAttrRelGroup:       newPmsAttrRelGroup(db, opts...),
		PmsAttrValue:          newPmsAttrValue(db, opts...),
		PmsBrand:              newPmsBrand(db, opts...),
		PmsCategory:           newPmsCategory(db, opts...),
		PmsCategoryRelBrand:   newPmsCategoryRelBrand(db, opts...),
		PmsSku:                newPmsSku(db, opts...),
		PmsSkuAttr:            newPmsSkuAttr(db, opts...),
		PmsSpu:                newPmsSpu(db, opts...),
		PmsSpuComment:         newPmsSpuComment(db, opts...),
		PmsSpuDesc:            newPmsSpuDesc(db, opts...),
		SmsAgreement:          newSmsAgreement(db, opts...),
		SmsAppNotice:          newSmsAppNotice(db, opts...),
		SmsAppSetting:         newSmsAppSetting(db, opts...),
		SmsAppVersion:         newSmsAppVersion(db, opts...),
		SmsConfig:             newSmsConfig(db, opts...),
		SmsCoupon:             newSmsCoupon(db, opts...),
		SmsCouponMember:       newSmsCouponMember(db, opts...),
		SmsCouponRelCat:       newSmsCouponRelCat(db, opts...),
		SmsCouponRelSpu:       newSmsCouponRelSpu(db, opts...),
		SmsHomeAdv:            newSmsHomeAdv(db, opts...),
		SmsHomeSubject:        newSmsHomeSubject(db, opts...),
		SmsHomeSubjectSpu:     newSmsHomeSubjectSpu(db, opts...),
		SmsMemberPrice:        newSmsMemberPrice(db, opts...),
		SmsSeckillActivity:    newSmsSeckillActivity(db, opts...),
		SmsSeckillSession:     newSmsSeckillSession(db, opts...),
		SmsSeckillSku:         newSmsSeckillSku(db, opts...),
		SmsSeckillSkuNotice:   newSmsSeckillSkuNotice(db, opts...),
		SmsSkuFullReduction:   newSmsSkuFullReduction(db, opts...),
		SmsSkuLadder:          newSmsSkuLadder(db, opts...),
		SmsSpuBound:           newSmsSpuBound(db, opts...),
		SysCity:               newSysCity(db, opts...),
		SysCityArea:           newSysCityArea(db, opts...),
		SysExpress:            newSysExpress(db, opts...),
		UmsCollectSpu:         newUmsCollectSpu(db, opts...),
		UmsCollectSubject:     newUmsCollectSubject(db, opts...),
		UmsGrowthLog:          newUmsGrowthLog(db, opts...),
		UmsIntegrationLog:     newUmsIntegrationLog(db, opts...),
		UmsMember:             newUmsMember(db, opts...),
		UmsMemberAddress:      newUmsMemberAddress(db, opts...),
		UmsMemberLevel:        newUmsMemberLevel(db, opts...),
		UmsMemberLog:          newUmsMemberLog(db, opts...),
		UmsMemberStat:         newUmsMemberStat(db, opts...),
		WmsPurchase:           newWmsPurchase(db, opts...),
		WmsPurchaseDetail:     newWmsPurchaseDetail(db, opts...),
		WmsWare:               newWmsWare(db, opts...),
		WmsWareSku:            newWmsWareSku(db, opts...),
		WmsWareTask:           newWmsWareTask(db, opts...),
		WmsWareTaskDetail:     newWmsWareTaskDetail(db, opts...),
	}
}

type Query struct {
	db *gorm.DB

	AdminExceptionLog     adminExceptionLog
	AdminExtension        adminExtension
	AdminExtensionHistory adminExtensionHistory
	AdminMenu             adminMenu
	AdminOperationLog     adminOperationLog
	AdminPermission       adminPermission
	AdminPermissionMenu   adminPermissionMenu
	AdminRole             adminRole
	AdminRoleMenu         adminRoleMenu
	AdminRolePermission   adminRolePermission
	AdminRoleUser         adminRoleUser
	AdminSetting          adminSetting
	AdminUser             adminUser
	FailedJob             failedJob
	Migration             migration
	OmsOrder              omsOrder
	OmsOrderBill          omsOrderBill
	OmsOrderItem          omsOrderItem
	OmsOrderOperate       omsOrderOperate
	OmsOrderRefund        omsOrderRefund
	OmsOrderReturnApply   omsOrderReturnApply
	OmsOrderReturnReason  omsOrderReturnReason
	OmsPayment            omsPayment
	PmsAttr               pmsAttr
	PmsAttrGroup          pmsAttrGroup
	PmsAttrRelGroup       pmsAttrRelGroup
	PmsAttrValue          pmsAttrValue
	PmsBrand              pmsBrand
	PmsCategory           pmsCategory
	PmsCategoryRelBrand   pmsCategoryRelBrand
	PmsSku                pmsSku
	PmsSkuAttr            pmsSkuAttr
	PmsSpu                pmsSpu
	PmsSpuComment         pmsSpuComment
	PmsSpuDesc            pmsSpuDesc
	SmsAgreement          smsAgreement
	SmsAppNotice          smsAppNotice
	SmsAppSetting         smsAppSetting
	SmsAppVersion         smsAppVersion
	SmsConfig             smsConfig
	SmsCoupon             smsCoupon
	SmsCouponMember       smsCouponMember
	SmsCouponRelCat       smsCouponRelCat
	SmsCouponRelSpu       smsCouponRelSpu
	SmsHomeAdv            smsHomeAdv
	SmsHomeSubject        smsHomeSubject
	SmsHomeSubjectSpu     smsHomeSubjectSpu
	SmsMemberPrice        smsMemberPrice
	SmsSeckillActivity    smsSeckillActivity
	SmsSeckillSession     smsSeckillSession
	SmsSeckillSku         smsSeckillSku
	SmsSeckillSkuNotice   smsSeckillSkuNotice
	SmsSkuFullReduction   smsSkuFullReduction
	SmsSkuLadder          smsSkuLadder
	SmsSpuBound           smsSpuBound
	SysCity               sysCity
	SysCityArea           sysCityArea
	SysExpress            sysExpress
	UmsCollectSpu         umsCollectSpu
	UmsCollectSubject     umsCollectSubject
	UmsGrowthLog          umsGrowthLog
	UmsIntegrationLog     umsIntegrationLog
	UmsMember             umsMember
	UmsMemberAddress      umsMemberAddress
	UmsMemberLevel        umsMemberLevel
	UmsMemberLog          umsMemberLog
	UmsMemberStat         umsMemberStat
	WmsPurchase           wmsPurchase
	WmsPurchaseDetail     wmsPurchaseDetail
	WmsWare               wmsWare
	WmsWareSku            wmsWareSku
	WmsWareTask           wmsWareTask
	WmsWareTaskDetail     wmsWareTaskDetail
}

func (q *Query) Available() bool { return q.db != nil }

func (q *Query) clone(db *gorm.DB) *Query {
	return &Query{
		db:                    db,
		AdminExceptionLog:     q.AdminExceptionLog.clone(db),
		AdminExtension:        q.AdminExtension.clone(db),
		AdminExtensionHistory: q.AdminExtensionHistory.clone(db),
		AdminMenu:             q.AdminMenu.clone(db),
		AdminOperationLog:     q.AdminOperationLog.clone(db),
		AdminPermission:       q.AdminPermission.clone(db),
		AdminPermissionMenu:   q.AdminPermissionMenu.clone(db),
		AdminRole:             q.AdminRole.clone(db),
		AdminRoleMenu:         q.AdminRoleMenu.clone(db),
		AdminRolePermission:   q.AdminRolePermission.clone(db),
		AdminRoleUser:         q.AdminRoleUser.clone(db),
		AdminSetting:          q.AdminSetting.clone(db),
		AdminUser:             q.AdminUser.clone(db),
		FailedJob:             q.FailedJob.clone(db),
		Migration:             q.Migration.clone(db),
		OmsOrder:              q.OmsOrder.clone(db),
		OmsOrderBill:          q.OmsOrderBill.clone(db),
		OmsOrderItem:          q.OmsOrderItem.clone(db),
		OmsOrderOperate:       q.OmsOrderOperate.clone(db),
		OmsOrderRefund:        q.OmsOrderRefund.clone(db),
		OmsOrderReturnApply:   q.OmsOrderReturnApply.clone(db),
		OmsOrderReturnReason:  q.OmsOrderReturnReason.clone(db),
		OmsPayment:            q.OmsPayment.clone(db),
		PmsAttr:               q.PmsAttr.clone(db),
		PmsAttrGroup:          q.PmsAttrGroup.clone(db),
		PmsAttrRelGroup:       q.PmsAttrRelGroup.clone(db),
		PmsAttrValue:          q.PmsAttrValue.clone(db),
		PmsBrand:              q.PmsBrand.clone(db),
		PmsCategory:           q.PmsCategory.clone(db),
		PmsCategoryRelBrand:   q.PmsCategoryRelBrand.clone(db),
		PmsSku:                q.PmsSku.clone(db),
		PmsSkuAttr:            q.PmsSkuAttr.clone(db),
		PmsSpu:                q.PmsSpu.clone(db),
		PmsSpuComment:         q.PmsSpuComment.clone(db),
		PmsSpuDesc:            q.PmsSpuDesc.clone(db),
		SmsAgreement:          q.SmsAgreement.clone(db),
		SmsAppNotice:          q.SmsAppNotice.clone(db),
		SmsAppSetting:         q.SmsAppSetting.clone(db),
		SmsAppVersion:         q.SmsAppVersion.clone(db),
		SmsConfig:             q.SmsConfig.clone(db),
		SmsCoupon:             q.SmsCoupon.clone(db),
		SmsCouponMember:       q.SmsCouponMember.clone(db),
		SmsCouponRelCat:       q.SmsCouponRelCat.clone(db),
		SmsCouponRelSpu:       q.SmsCouponRelSpu.clone(db),
		SmsHomeAdv:            q.SmsHomeAdv.clone(db),
		SmsHomeSubject:        q.SmsHomeSubject.clone(db),
		SmsHomeSubjectSpu:     q.SmsHomeSubjectSpu.clone(db),
		SmsMemberPrice:        q.SmsMemberPrice.clone(db),
		SmsSeckillActivity:    q.SmsSeckillActivity.clone(db),
		SmsSeckillSession:     q.SmsSeckillSession.clone(db),
		SmsSeckillSku:         q.SmsSeckillSku.clone(db),
		SmsSeckillSkuNotice:   q.SmsSeckillSkuNotice.clone(db),
		SmsSkuFullReduction:   q.SmsSkuFullReduction.clone(db),
		SmsSkuLadder:          q.SmsSkuLadder.clone(db),
		SmsSpuBound:           q.SmsSpuBound.clone(db),
		SysCity:               q.SysCity.clone(db),
		SysCityArea:           q.SysCityArea.clone(db),
		SysExpress:            q.SysExpress.clone(db),
		UmsCollectSpu:         q.UmsCollectSpu.clone(db),
		UmsCollectSubject:     q.UmsCollectSubject.clone(db),
		UmsGrowthLog:          q.UmsGrowthLog.clone(db),
		UmsIntegrationLog:     q.UmsIntegrationLog.clone(db),
		UmsMember:             q.UmsMember.clone(db),
		UmsMemberAddress:      q.UmsMemberAddress.clone(db),
		UmsMemberLevel:        q.UmsMemberLevel.clone(db),
		UmsMemberLog:          q.UmsMemberLog.clone(db),
		UmsMemberStat:         q.UmsMemberStat.clone(db),
		WmsPurchase:           q.WmsPurchase.clone(db),
		WmsPurchaseDetail:     q.WmsPurchaseDetail.clone(db),
		WmsWare:               q.WmsWare.clone(db),
		WmsWareSku:            q.WmsWareSku.clone(db),
		WmsWareTask:           q.WmsWareTask.clone(db),
		WmsWareTaskDetail:     q.WmsWareTaskDetail.clone(db),
	}
}

func (q *Query) ReadDB() *Query {
	return q.ReplaceDB(q.db.Clauses(dbresolver.Read))
}

func (q *Query) WriteDB() *Query {
	return q.ReplaceDB(q.db.Clauses(dbresolver.Write))
}

func (q *Query) ReplaceDB(db *gorm.DB) *Query {
	return &Query{
		db:                    db,
		AdminExceptionLog:     q.AdminExceptionLog.replaceDB(db),
		AdminExtension:        q.AdminExtension.replaceDB(db),
		AdminExtensionHistory: q.AdminExtensionHistory.replaceDB(db),
		AdminMenu:             q.AdminMenu.replaceDB(db),
		AdminOperationLog:     q.AdminOperationLog.replaceDB(db),
		AdminPermission:       q.AdminPermission.replaceDB(db),
		AdminPermissionMenu:   q.AdminPermissionMenu.replaceDB(db),
		AdminRole:             q.AdminRole.replaceDB(db),
		AdminRoleMenu:         q.AdminRoleMenu.replaceDB(db),
		AdminRolePermission:   q.AdminRolePermission.replaceDB(db),
		AdminRoleUser:         q.AdminRoleUser.replaceDB(db),
		AdminSetting:          q.AdminSetting.replaceDB(db),
		AdminUser:             q.AdminUser.replaceDB(db),
		FailedJob:             q.FailedJob.replaceDB(db),
		Migration:             q.Migration.replaceDB(db),
		OmsOrder:              q.OmsOrder.replaceDB(db),
		OmsOrderBill:          q.OmsOrderBill.replaceDB(db),
		OmsOrderItem:          q.OmsOrderItem.replaceDB(db),
		OmsOrderOperate:       q.OmsOrderOperate.replaceDB(db),
		OmsOrderRefund:        q.OmsOrderRefund.replaceDB(db),
		OmsOrderReturnApply:   q.OmsOrderReturnApply.replaceDB(db),
		OmsOrderReturnReason:  q.OmsOrderReturnReason.replaceDB(db),
		OmsPayment:            q.OmsPayment.replaceDB(db),
		PmsAttr:               q.PmsAttr.replaceDB(db),
		PmsAttrGroup:          q.PmsAttrGroup.replaceDB(db),
		PmsAttrRelGroup:       q.PmsAttrRelGroup.replaceDB(db),
		PmsAttrValue:          q.PmsAttrValue.replaceDB(db),
		PmsBrand:              q.PmsBrand.replaceDB(db),
		PmsCategory:           q.PmsCategory.replaceDB(db),
		PmsCategoryRelBrand:   q.PmsCategoryRelBrand.replaceDB(db),
		PmsSku:                q.PmsSku.replaceDB(db),
		PmsSkuAttr:            q.PmsSkuAttr.replaceDB(db),
		PmsSpu:                q.PmsSpu.replaceDB(db),
		PmsSpuComment:         q.PmsSpuComment.replaceDB(db),
		PmsSpuDesc:            q.PmsSpuDesc.replaceDB(db),
		SmsAgreement:          q.SmsAgreement.replaceDB(db),
		SmsAppNotice:          q.SmsAppNotice.replaceDB(db),
		SmsAppSetting:         q.SmsAppSetting.replaceDB(db),
		SmsAppVersion:         q.SmsAppVersion.replaceDB(db),
		SmsConfig:             q.SmsConfig.replaceDB(db),
		SmsCoupon:             q.SmsCoupon.replaceDB(db),
		SmsCouponMember:       q.SmsCouponMember.replaceDB(db),
		SmsCouponRelCat:       q.SmsCouponRelCat.replaceDB(db),
		SmsCouponRelSpu:       q.SmsCouponRelSpu.replaceDB(db),
		SmsHomeAdv:            q.SmsHomeAdv.replaceDB(db),
		SmsHomeSubject:        q.SmsHomeSubject.replaceDB(db),
		SmsHomeSubjectSpu:     q.SmsHomeSubjectSpu.replaceDB(db),
		SmsMemberPrice:        q.SmsMemberPrice.replaceDB(db),
		SmsSeckillActivity:    q.SmsSeckillActivity.replaceDB(db),
		SmsSeckillSession:     q.SmsSeckillSession.replaceDB(db),
		SmsSeckillSku:         q.SmsSeckillSku.replaceDB(db),
		SmsSeckillSkuNotice:   q.SmsSeckillSkuNotice.replaceDB(db),
		SmsSkuFullReduction:   q.SmsSkuFullReduction.replaceDB(db),
		SmsSkuLadder:          q.SmsSkuLadder.replaceDB(db),
		SmsSpuBound:           q.SmsSpuBound.replaceDB(db),
		SysCity:               q.SysCity.replaceDB(db),
		SysCityArea:           q.SysCityArea.replaceDB(db),
		SysExpress:            q.SysExpress.replaceDB(db),
		UmsCollectSpu:         q.UmsCollectSpu.replaceDB(db),
		UmsCollectSubject:     q.UmsCollectSubject.replaceDB(db),
		UmsGrowthLog:          q.UmsGrowthLog.replaceDB(db),
		UmsIntegrationLog:     q.UmsIntegrationLog.replaceDB(db),
		UmsMember:             q.UmsMember.replaceDB(db),
		UmsMemberAddress:      q.UmsMemberAddress.replaceDB(db),
		UmsMemberLevel:        q.UmsMemberLevel.replaceDB(db),
		UmsMemberLog:          q.UmsMemberLog.replaceDB(db),
		UmsMemberStat:         q.UmsMemberStat.replaceDB(db),
		WmsPurchase:           q.WmsPurchase.replaceDB(db),
		WmsPurchaseDetail:     q.WmsPurchaseDetail.replaceDB(db),
		WmsWare:               q.WmsWare.replaceDB(db),
		WmsWareSku:            q.WmsWareSku.replaceDB(db),
		WmsWareTask:           q.WmsWareTask.replaceDB(db),
		WmsWareTaskDetail:     q.WmsWareTaskDetail.replaceDB(db),
	}
}

type queryCtx struct {
	AdminExceptionLog     IAdminExceptionLogDo
	AdminExtension        IAdminExtensionDo
	AdminExtensionHistory IAdminExtensionHistoryDo
	AdminMenu             IAdminMenuDo
	AdminOperationLog     IAdminOperationLogDo
	AdminPermission       IAdminPermissionDo
	AdminPermissionMenu   IAdminPermissionMenuDo
	AdminRole             IAdminRoleDo
	AdminRoleMenu         IAdminRoleMenuDo
	AdminRolePermission   IAdminRolePermissionDo
	AdminRoleUser         IAdminRoleUserDo
	AdminSetting          IAdminSettingDo
	AdminUser             IAdminUserDo
	FailedJob             IFailedJobDo
	Migration             IMigrationDo
	OmsOrder              IOmsOrderDo
	OmsOrderBill          IOmsOrderBillDo
	OmsOrderItem          IOmsOrderItemDo
	OmsOrderOperate       IOmsOrderOperateDo
	OmsOrderRefund        IOmsOrderRefundDo
	OmsOrderReturnApply   IOmsOrderReturnApplyDo
	OmsOrderReturnReason  IOmsOrderReturnReasonDo
	OmsPayment            IOmsPaymentDo
	PmsAttr               IPmsAttrDo
	PmsAttrGroup          IPmsAttrGroupDo
	PmsAttrRelGroup       IPmsAttrRelGroupDo
	PmsAttrValue          IPmsAttrValueDo
	PmsBrand              IPmsBrandDo
	PmsCategory           IPmsCategoryDo
	PmsCategoryRelBrand   IPmsCategoryRelBrandDo
	PmsSku                IPmsSkuDo
	PmsSkuAttr            IPmsSkuAttrDo
	PmsSpu                IPmsSpuDo
	PmsSpuComment         IPmsSpuCommentDo
	PmsSpuDesc            IPmsSpuDescDo
	SmsAgreement          ISmsAgreementDo
	SmsAppNotice          ISmsAppNoticeDo
	SmsAppSetting         ISmsAppSettingDo
	SmsAppVersion         ISmsAppVersionDo
	SmsConfig             ISmsConfigDo
	SmsCoupon             ISmsCouponDo
	SmsCouponMember       ISmsCouponMemberDo
	SmsCouponRelCat       ISmsCouponRelCatDo
	SmsCouponRelSpu       ISmsCouponRelSpuDo
	SmsHomeAdv            ISmsHomeAdvDo
	SmsHomeSubject        ISmsHomeSubjectDo
	SmsHomeSubjectSpu     ISmsHomeSubjectSpuDo
	SmsMemberPrice        ISmsMemberPriceDo
	SmsSeckillActivity    ISmsSeckillActivityDo
	SmsSeckillSession     ISmsSeckillSessionDo
	SmsSeckillSku         ISmsSeckillSkuDo
	SmsSeckillSkuNotice   ISmsSeckillSkuNoticeDo
	SmsSkuFullReduction   ISmsSkuFullReductionDo
	SmsSkuLadder          ISmsSkuLadderDo
	SmsSpuBound           ISmsSpuBoundDo
	SysCity               ISysCityDo
	SysCityArea           ISysCityAreaDo
	SysExpress            ISysExpressDo
	UmsCollectSpu         IUmsCollectSpuDo
	UmsCollectSubject     IUmsCollectSubjectDo
	UmsGrowthLog          IUmsGrowthLogDo
	UmsIntegrationLog     IUmsIntegrationLogDo
	UmsMember             IUmsMemberDo
	UmsMemberAddress      IUmsMemberAddressDo
	UmsMemberLevel        IUmsMemberLevelDo
	UmsMemberLog          IUmsMemberLogDo
	UmsMemberStat         IUmsMemberStatDo
	WmsPurchase           IWmsPurchaseDo
	WmsPurchaseDetail     IWmsPurchaseDetailDo
	WmsWare               IWmsWareDo
	WmsWareSku            IWmsWareSkuDo
	WmsWareTask           IWmsWareTaskDo
	WmsWareTaskDetail     IWmsWareTaskDetailDo
}

func (q *Query) WithContext(ctx context.Context) *queryCtx {
	return &queryCtx{
		AdminExceptionLog:     q.AdminExceptionLog.WithContext(ctx),
		AdminExtension:        q.AdminExtension.WithContext(ctx),
		AdminExtensionHistory: q.AdminExtensionHistory.WithContext(ctx),
		AdminMenu:             q.AdminMenu.WithContext(ctx),
		AdminOperationLog:     q.AdminOperationLog.WithContext(ctx),
		AdminPermission:       q.AdminPermission.WithContext(ctx),
		AdminPermissionMenu:   q.AdminPermissionMenu.WithContext(ctx),
		AdminRole:             q.AdminRole.WithContext(ctx),
		AdminRoleMenu:         q.AdminRoleMenu.WithContext(ctx),
		AdminRolePermission:   q.AdminRolePermission.WithContext(ctx),
		AdminRoleUser:         q.AdminRoleUser.WithContext(ctx),
		AdminSetting:          q.AdminSetting.WithContext(ctx),
		AdminUser:             q.AdminUser.WithContext(ctx),
		FailedJob:             q.FailedJob.WithContext(ctx),
		Migration:             q.Migration.WithContext(ctx),
		OmsOrder:              q.OmsOrder.WithContext(ctx),
		OmsOrderBill:          q.OmsOrderBill.WithContext(ctx),
		OmsOrderItem:          q.OmsOrderItem.WithContext(ctx),
		OmsOrderOperate:       q.OmsOrderOperate.WithContext(ctx),
		OmsOrderRefund:        q.OmsOrderRefund.WithContext(ctx),
		OmsOrderReturnApply:   q.OmsOrderReturnApply.WithContext(ctx),
		OmsOrderReturnReason:  q.OmsOrderReturnReason.WithContext(ctx),
		OmsPayment:            q.OmsPayment.WithContext(ctx),
		PmsAttr:               q.PmsAttr.WithContext(ctx),
		PmsAttrGroup:          q.PmsAttrGroup.WithContext(ctx),
		PmsAttrRelGroup:       q.PmsAttrRelGroup.WithContext(ctx),
		PmsAttrValue:          q.PmsAttrValue.WithContext(ctx),
		PmsBrand:              q.PmsBrand.WithContext(ctx),
		PmsCategory:           q.PmsCategory.WithContext(ctx),
		PmsCategoryRelBrand:   q.PmsCategoryRelBrand.WithContext(ctx),
		PmsSku:                q.PmsSku.WithContext(ctx),
		PmsSkuAttr:            q.PmsSkuAttr.WithContext(ctx),
		PmsSpu:                q.PmsSpu.WithContext(ctx),
		PmsSpuComment:         q.PmsSpuComment.WithContext(ctx),
		PmsSpuDesc:            q.PmsSpuDesc.WithContext(ctx),
		SmsAgreement:          q.SmsAgreement.WithContext(ctx),
		SmsAppNotice:          q.SmsAppNotice.WithContext(ctx),
		SmsAppSetting:         q.SmsAppSetting.WithContext(ctx),
		SmsAppVersion:         q.SmsAppVersion.WithContext(ctx),
		SmsConfig:             q.SmsConfig.WithContext(ctx),
		SmsCoupon:             q.SmsCoupon.WithContext(ctx),
		SmsCouponMember:       q.SmsCouponMember.WithContext(ctx),
		SmsCouponRelCat:       q.SmsCouponRelCat.WithContext(ctx),
		SmsCouponRelSpu:       q.SmsCouponRelSpu.WithContext(ctx),
		SmsHomeAdv:            q.SmsHomeAdv.WithContext(ctx),
		SmsHomeSubject:        q.SmsHomeSubject.WithContext(ctx),
		SmsHomeSubjectSpu:     q.SmsHomeSubjectSpu.WithContext(ctx),
		SmsMemberPrice:        q.SmsMemberPrice.WithContext(ctx),
		SmsSeckillActivity:    q.SmsSeckillActivity.WithContext(ctx),
		SmsSeckillSession:     q.SmsSeckillSession.WithContext(ctx),
		SmsSeckillSku:         q.SmsSeckillSku.WithContext(ctx),
		SmsSeckillSkuNotice:   q.SmsSeckillSkuNotice.WithContext(ctx),
		SmsSkuFullReduction:   q.SmsSkuFullReduction.WithContext(ctx),
		SmsSkuLadder:          q.SmsSkuLadder.WithContext(ctx),
		SmsSpuBound:           q.SmsSpuBound.WithContext(ctx),
		SysCity:               q.SysCity.WithContext(ctx),
		SysCityArea:           q.SysCityArea.WithContext(ctx),
		SysExpress:            q.SysExpress.WithContext(ctx),
		UmsCollectSpu:         q.UmsCollectSpu.WithContext(ctx),
		UmsCollectSubject:     q.UmsCollectSubject.WithContext(ctx),
		UmsGrowthLog:          q.UmsGrowthLog.WithContext(ctx),
		UmsIntegrationLog:     q.UmsIntegrationLog.WithContext(ctx),
		UmsMember:             q.UmsMember.WithContext(ctx),
		UmsMemberAddress:      q.UmsMemberAddress.WithContext(ctx),
		UmsMemberLevel:        q.UmsMemberLevel.WithContext(ctx),
		UmsMemberLog:          q.UmsMemberLog.WithContext(ctx),
		UmsMemberStat:         q.UmsMemberStat.WithContext(ctx),
		WmsPurchase:           q.WmsPurchase.WithContext(ctx),
		WmsPurchaseDetail:     q.WmsPurchaseDetail.WithContext(ctx),
		WmsWare:               q.WmsWare.WithContext(ctx),
		WmsWareSku:            q.WmsWareSku.WithContext(ctx),
		WmsWareTask:           q.WmsWareTask.WithContext(ctx),
		WmsWareTaskDetail:     q.WmsWareTaskDetail.WithContext(ctx),
	}
}

func (q *Query) Transaction(fc func(tx *Query) error, opts ...*sql.TxOptions) error {
	return q.db.Transaction(func(tx *gorm.DB) error { return fc(q.clone(tx)) }, opts...)
}

func (q *Query) Begin(opts ...*sql.TxOptions) *QueryTx {
	tx := q.db.Begin(opts...)
	return &QueryTx{Query: q.clone(tx), Error: tx.Error}
}

type QueryTx struct {
	*Query
	Error error
}

func (q *QueryTx) Commit() error {
	return q.db.Commit().Error
}

func (q *QueryTx) Rollback() error {
	return q.db.Rollback().Error
}

func (q *QueryTx) SavePoint(name string) error {
	return q.db.SavePoint(name).Error
}

func (q *QueryTx) RollbackTo(name string) error {
	return q.db.RollbackTo(name).Error
}
