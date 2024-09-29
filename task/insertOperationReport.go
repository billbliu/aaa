package task

import (
	"github.com/billbliu/tarot-admin/global"
	"github.com/billbliu/tarot-admin/model/business/request"
	"github.com/billbliu/tarot-admin/service/business"
	"github.com/golang-module/carbon/v2"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

//@author: [billbliu]
//@function: InsertDayOperationReport
//@description: 定时更新运营报表
//@param: db(数据库对象) *gorm.DB
//@return: error

func InsertDayOperationReport(db *gorm.DB) error {
	now := carbon.Date{Carbon: carbon.Now(carbon.PRC).StartOfDay()}
	req := request.UpdateOperationReportRequest{
		Start: now,
		End:   now,
	}
	if err := business.OperationServiceApp.UpdateOperationReport(req); err != nil {
		global.GVA_LOG.Error("UpdateOperationReport", zap.Error(err))
	}
	return nil
}
