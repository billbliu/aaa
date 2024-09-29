package business

import (
	"context"
	"errors"
	"strconv"

	"github.com/billbliu/tarot-admin/global"
	"github.com/billbliu/tarot-admin/model/business/request"
	"github.com/billbliu/tarot-admin/model/business/response"
	"github.com/billbliu/tarot-admin/model/tables"
	"github.com/billbliu/tarot-admin/utils"
	"github.com/gin-gonic/gin"
	"github.com/golang-module/carbon/v2"
	"github.com/shopspring/decimal"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type OperationService struct{}

var OperationServiceApp = new(OperationService)

// @author: billbliu
// @function: GetStatisticsCount
// @description: 会员、充值、金币等统计信息
// @return: res response.GetStatisticsCountResponse, err error
func (e *OperationService) GetStatisticsCount(c *gin.Context) (res response.GetStatisticsCountResponse, err error) {
	now := utils.CarbonNow()
	systemStart := carbon.CreateFromDate(2023, 1, 1)
	start := now.StartOfDay()

	memberTotalCount, err := tables.TUserDao.GetRangeUserCount(global.GVA_DB, systemStart, now.Carbon)
	if err != nil {
		global.GVA_LOG.Error("memberTotalCount err", zap.Error(err))
	}
	res.MemberTotalCount = memberTotalCount

	memberDayRegisterCount, err := tables.TUserDao.GetRangeUserCount(global.GVA_DB, start, now.Carbon)
	if err != nil {
		global.GVA_LOG.Error("memberDayRegisterCount err", zap.Error(err))
	}
	res.MemberDayRegisterCount = memberDayRegisterCount

	// 查询集合中键的个数
	memberDayActiveCount, err := global.GVA_REDIS.Keys(c, "actUser*").Result()
	if err != nil {
		global.GVA_LOG.Error(err.Error())
	}

	res.MemberDayActiveCount = len(memberDayActiveCount)

	start = now.SubYears(30)
	totalUsdtAmount, err := tables.TUserAssetOrderDao.GetAllUserDepositUsdtAmountByTime(global.GVA_DB, start, now.Carbon)
	if err != nil {
		return res, err
	}
	res.UsdtTotalAmount = totalUsdtAmount

	start = now.StartOfDay()
	thisDayUsdtAmount, err := tables.TUserAssetOrderDao.GetAllUserDepositUsdtAmountByTime(global.GVA_DB, start, now.Carbon)
	if err != nil {
		return res, err
	}
	res.UsdtDayAmount = thisDayUsdtAmount

	totalCoin, err := tables.TUserAssetDao.GetTotalCoin(global.GVA_DB)
	if err != nil {
		return res, err
	}
	res.CoinTotalAmount = totalCoin

	totalUseCoin, err := tables.TUserAssetOrderDao.GetTotalUseCoin(global.GVA_DB)
	if err != nil {
		return res, err
	}
	res.CoinUseTotalAmount = totalUseCoin

	start = now.StartOfDay()
	behaviors := []tables.BehaviorType{
		tables.BEHAVIOR_DIVINATION_CHOOSE_ONE_OF_THE_TWO,
		tables.BEHAVIOR_DIVINATION_NO_CARD_ARRAY,
		tables.BEHAVIOR_DIVINATION_HOLY_TRIANGLE,
		tables.BEHAVIOR_DIVINATION_CELTIC_CROSS,
	}
	dayUseCoin, err := tables.TUserAssetOrderDao.GetRangeTotalUseCoin(global.GVA_DB, start, now.Carbon, behaviors)
	if err != nil {
		return res, err
	}
	res.CoinDayUseAmount = dayUseCoin

	return res, err
}

// @author: billbliu
// @function: GetMembersCount
// @description: 获取会员数量信息
// @return: res response.GetMembersCountResponse, err error
func (e *OperationService) GetDepositAmount() (res response.GetDepositAmountResponse, err error) {
	now := utils.CarbonNow()
	totalStart := now.SubYears(30)
	totalUsdtAmount, err := tables.TUserAssetOrderDao.GetAllUserDepositUsdtAmountByTime(global.GVA_DB, totalStart, now.Carbon)
	if err != nil {
		return res, err
	}
	res.TotalUsdtAmount = totalUsdtAmount

	thisDayStart := now.StartOfDay()
	thisDayUsdtAmount, err := tables.TUserAssetOrderDao.GetAllUserDepositUsdtAmountByTime(global.GVA_DB, thisDayStart, now.Carbon)
	if err != nil {
		return res, err
	}
	res.ThisDayUsdtAmount = thisDayUsdtAmount
	return res, err
}

// @author: billbliu
// @function: UpdateAllUserAssetTotalInfo
// @description: 更新所有会员资产总信息
// @return: error
func (e *OperationService) UpdateAllUserAssetTotalInfo() error {
	users := []tables.TUser{}
	if err := global.GVA_DB.Model(&tables.TUser{}).Find(&users).Error; err != nil {
		return err
	}

	for _, user := range users {
		totalCoinBehaviors := []tables.BehaviorType{
			tables.BEHAVIOR_DEPOSIT,
			tables.BEHAVIOR_ROYALTY_CONVERT_COIN,
			tables.BEHAVIOR_ROYALTY_SEND_COIN,
			tables.BEHAVIOR_REGISTER_SEND_COIN,
			tables.BEHAVIOR_MONEY_CONVERT_COIN,
		}
		totalCoin, err := tables.TUserAssetOrderDao.GetUserTotalAmountByBehaviors(global.GVA_DB, user.ID, totalCoinBehaviors)
		if err != nil {
			global.GVA_LOG.Error("GetUserTotalAmountByBehaviors err", zap.Error(err))
		}

		totalReturnBehaviors := []tables.BehaviorType{
			tables.BEHAVIOR_ROYALTY,
		}
		totalReturn, err := tables.TUserAssetOrderDao.GetUserTotalAmountByBehaviors(global.GVA_DB, user.ID, totalReturnBehaviors)
		if err != nil {
			global.GVA_LOG.Error("GetUserTotalAmountByBehaviors err", zap.Error(err))
		}

		totalMoneyBehaviors := []tables.BehaviorType{
			tables.BEHAVIOR_ROYALTY_CONVERT_MONEY,
		}
		totalMoney, err := tables.TUserAssetOrderDao.GetUserTotalAmountByBehaviors(global.GVA_DB, user.ID, totalMoneyBehaviors)
		if err != nil {
			global.GVA_LOG.Error("GetUserTotalAmountByBehaviors err", zap.Error(err))
		}

		totalUsdt, err := tables.TUserAssetOrderDao.GetUserTotalUsdtAmountByUserID(global.GVA_DB, user.ID)
		if err != nil {
			global.GVA_LOG.Error("GetUserTotalAmountByBehaviors err", zap.Error(err))
		}

		tUserAsset := tables.TUserAsset{}
		if err := global.GVA_DB.Model(&tables.TUserAsset{}).Where("user_id = ?", user.ID).First(&tUserAsset).Error; err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				// 不存在则创建
				tUserAsset.UserId = user.ID
				if err = global.GVA_DB.Create(&tUserAsset).Error; err != nil {
					global.GVA_LOG.Error("Create err", zap.Error(err))
				}
			} else {
				global.GVA_LOG.Error("Query err", zap.Error(err))
			}
		} else {
			// 存在则更新
			tUserAsset.TotalCoin = totalCoin
			tUserAsset.TotalReturnCash = totalReturn
			tUserAsset.TotalUsdt = totalUsdt
			tUserAsset.TotalMoney = totalMoney
			if err = global.GVA_DB.Save(&tUserAsset).Error; err != nil {
				global.GVA_LOG.Error("Save err", zap.Error(err))
			}
		}

		// if err := global.GVA_DB.Model(&tables.TUserAsset{}).Where("user_id = ?", user.ID).Updates(tables.TUserAsset{
		// 	TotalCoin:       totalCoin,
		// 	TotalReturnCash: totalReturn,
		// 	TotalUsdt:       totalUsdt,
		// 	TotalMoney:      totalMoney,
		// }).Error; err != nil {
		// 	global.GVA_LOG.Error("GetUserTotalAmountByBehaviors err", zap.Error(err))
		// }
	}

	return nil
}

// @author: billbliu
// @function: UpdateOperationReport
// @description: 更新所有会员资产总信息
// @return: error
func (e *OperationService) UpdateOperationReport(req request.UpdateOperationReportRequest) error {
	now := req.Start
	for now.Lte(req.End.Carbon) {
		report := tables.TOperationReport{}
		if err := global.GVA_DB.Model(&tables.TOperationReport{}).Where("which_day = ?", now).First(&report).Error; err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				res := e.getWhiceDayOperationReport(now)
				// 不存在则创建
				report.WhichDay = now
				report.TotalMembers = res.TotalMembers
				report.DayRegisterMembers = res.DayRegisterMembers
				report.DayActiveMembers = res.DayActiveMembers

				report.TotalDepositUsdtAmount = res.TotalDepositUsdtAmount
				report.DayDepositUsdtAmount = res.DayDepositUsdtAmount

				report.TotalDepositCoin = res.TotalDepositCoin
				report.DayBinaryDivinationCoin = res.DayBinaryDivinationCoin
				report.DayRandomDivinationCoin = res.DayRandomDivinationCoin
				report.DayTrinityDivinationCoin = res.DayTrinityDivinationCoin
				report.DayCelticCrossDivinationCoin = res.DayCelticCrossDivinationCoin
				if err = global.GVA_DB.Create(&report).Error; err != nil {
					global.GVA_LOG.Error("Create err", zap.Error(err))
				}
			} else {
				global.GVA_LOG.Error("Query err", zap.Error(err))
			}
		} else {
			// 存在则更新
			res := e.getWhiceDayOperationReport(now)
			report.WhichDay = now
			report.TotalMembers = res.TotalMembers
			report.DayRegisterMembers = res.DayRegisterMembers
			report.DayActiveMembers = res.DayActiveMembers

			report.TotalDepositUsdtAmount = res.TotalDepositUsdtAmount
			report.DayDepositUsdtAmount = res.DayDepositUsdtAmount

			report.TotalDepositCoin = res.TotalDepositCoin
			report.DayBinaryDivinationCoin = res.DayBinaryDivinationCoin
			report.DayRandomDivinationCoin = res.DayRandomDivinationCoin
			report.DayTrinityDivinationCoin = res.DayTrinityDivinationCoin
			report.DayCelticCrossDivinationCoin = res.DayCelticCrossDivinationCoin
			if err = global.GVA_DB.Save(&report).Error; err != nil {
				global.GVA_LOG.Error("Save err", zap.Error(err))
			}
		}
		now = now.AddDay().ToDateStruct()
	}
	return nil
}

// @author: billbliu
// @function: getWhiceDayOperationReport
// @description: 获取某天运维报表数据
// @return: tables.TOperationReport
func (e *OperationService) getWhiceDayOperationReport(day carbon.Date) tables.TOperationReport {
	systemStart := carbon.CreateFromDate(2023, 1, 1)
	dayStart, dayEnd := day.StartOfDay(), day.EndOfDay()
	res := tables.TOperationReport{}

	memberTotalCount, err := tables.TUserDao.GetRangeUserCount(global.GVA_DB, systemStart, dayEnd)
	if err != nil {
		global.GVA_LOG.Error("memberTotalCount err", zap.Error(err))
	}
	res.TotalMembers = memberTotalCount

	memberDayRegisterCount, err := tables.TUserDao.GetRangeUserCount(global.GVA_DB, dayStart, dayEnd)
	if err != nil {
		global.GVA_LOG.Error("memberDayRegisterCount err", zap.Error(err))
	}
	res.DayRegisterMembers = memberDayRegisterCount

	ctx := context.Background()
	memberDayActiveCount, err := global.GVA_REDIS.Keys(ctx, "actUser*").Result()
	if err != nil {
		global.GVA_LOG.Error(err.Error())
	}

	res.DayActiveMembers = int64(len(memberDayActiveCount))

	totalUsdtAmount, err := tables.TUserAssetOrderDao.GetAllUserDepositUsdtAmountByTime(global.GVA_DB, systemStart, dayEnd)
	if err != nil {
		global.GVA_LOG.Error("totalUsdtAmount err", zap.Error(err))
	}
	res.TotalDepositUsdtAmount = totalUsdtAmount

	thisDayUsdtAmount, err := tables.TUserAssetOrderDao.GetAllUserDepositUsdtAmountByTime(global.GVA_DB, dayStart, dayEnd)
	if err != nil {
		global.GVA_LOG.Error("thisDayUsdtAmount err", zap.Error(err))
	}
	res.DayDepositUsdtAmount = thisDayUsdtAmount

	binaryBehaviors := []tables.BehaviorType{
		tables.BEHAVIOR_DIVINATION_CHOOSE_ONE_OF_THE_TWO,
	}
	dayBinaryDivinationCoin, err := tables.TUserAssetOrderDao.GetRangeTotalUseCoin(global.GVA_DB, dayStart, dayEnd, binaryBehaviors)
	if err != nil {
		global.GVA_LOG.Error("dayBinaryDivinationCoin err", zap.Error(err))
	}
	res.DayBinaryDivinationCoin = dayBinaryDivinationCoin

	randomBehaviors := []tables.BehaviorType{
		tables.BEHAVIOR_DIVINATION_NO_CARD_ARRAY,
	}
	dayRandomDivinationCoin, err := tables.TUserAssetOrderDao.GetRangeTotalUseCoin(global.GVA_DB, dayStart, dayEnd, randomBehaviors)
	if err != nil {
		global.GVA_LOG.Error("dayRandomDivinationCoin err", zap.Error(err))
	}
	res.DayRandomDivinationCoin = dayRandomDivinationCoin

	trinityBehaviors := []tables.BehaviorType{
		tables.BEHAVIOR_DIVINATION_HOLY_TRIANGLE,
	}
	dayTrinityDivinationCoin, err := tables.TUserAssetOrderDao.GetRangeTotalUseCoin(global.GVA_DB, dayStart, dayEnd, trinityBehaviors)
	if err != nil {
		global.GVA_LOG.Error("dayTrinityDivinationCoin err", zap.Error(err))
	}
	res.DayTrinityDivinationCoin = dayTrinityDivinationCoin

	celticBehaviors := []tables.BehaviorType{
		tables.BEHAVIOR_DIVINATION_CELTIC_CROSS,
	}
	dayCelticCrossDivinationCoin, err := tables.TUserAssetOrderDao.GetRangeTotalUseCoin(global.GVA_DB, dayStart, dayEnd, celticBehaviors)
	if err != nil {
		global.GVA_LOG.Error("dayCelticCrossDivinationCoin err", zap.Error(err))
	}
	res.DayCelticCrossDivinationCoin = dayCelticCrossDivinationCoin

	return res
}

// @author: billbliu
// @function: GetMemberCountChart
// @description: 会员数量图表数据
// @return: res response.GetMemberChartResponse, err error
func (e *OperationService) GetMemberCountChart(req request.GetMemberCountChartRequest) (response.GetMemberCountChartResponse, error) {
	if req.Days == 0 {
		req.Days = 7
	}

	res := response.GetMemberCountChartResponse{}

	end := utils.CarbonNow().SubDay()
	operationReports := []tables.TOperationReport{}
	if err := global.GVA_DB.Debug().Model(&tables.TOperationReport{}).Where("which_day BETWEEN ? AND ?", end.SubDays(req.Days), end).Find(&operationReports).Error; err != nil {
		return res, err
	}

	resDays := make([]string, len(operationReports))
	resTotalCounts := make([]int64, len(operationReports))
	resRegisterCounts := make([]int64, len(operationReports))
	resActiveCounts := make([]int64, len(operationReports))
	for i := 0; i < len(operationReports); i++ {
		resDays[i] = operationReports[i].WhichDay.String()
		resTotalCounts[i] = operationReports[i].TotalMembers
		resRegisterCounts[i] = operationReports[i].DayRegisterMembers
		resActiveCounts[i] = operationReports[i].DayActiveMembers
	}

	res.DayList = resDays
	res.TotalCountList = resTotalCounts
	res.DayRegisterCountList = resRegisterCounts
	res.DayActiveCountList = resActiveCounts
	return res, nil
}

// @author: billbliu
// @function: GetDepositUsdtChart
// @description: 充值图表数据
// @return: res response.GetDepositUsdtChartResponse, err error
func (e *OperationService) GetDepositUsdtChart(req request.GetDepositUsdtChartRequest) (response.GetDepositUsdtChartResponse, error) {
	if req.Days == 0 {
		req.Days = 7
	}

	res := response.GetDepositUsdtChartResponse{}

	end := utils.CarbonNow().SubDay()
	operationReports := []tables.TOperationReport{}
	if err := global.GVA_DB.Debug().Model(&tables.TOperationReport{}).Where("which_day BETWEEN ? AND ?", end.SubDays(req.Days), end).Find(&operationReports).Error; err != nil {
		return res, err
	}

	resDays := make([]string, len(operationReports))
	resTotalCounts := make([]decimal.Decimal, len(operationReports))
	resDayCounts := make([]decimal.Decimal, len(operationReports))
	for i := 0; i < len(operationReports); i++ {
		resDays[i] = operationReports[i].WhichDay.String()
		resTotalCounts[i] = operationReports[i].TotalDepositUsdtAmount
		resDayCounts[i] = operationReports[i].DayDepositUsdtAmount
	}

	res.DayList = resDays
	res.TotalCountList = resTotalCounts
	res.DayCountList = resDayCounts
	return res, nil
}

// @author: billbliu
// @function: GetCoinUseChart
// @description: 金币消费图表数据
// @return: res response.GetDepositUsdtChartResponse, err error
func (e *OperationService) GetCoinUseChart(req request.GetCoinUseChartRequest) (response.GetCoinUseChartResponse, error) {
	if req.Days == 0 {
		req.Days = 7
	}

	res := response.GetCoinUseChartResponse{}

	end := utils.CarbonNow().SubDay()
	operationReports := []tables.TOperationReport{}
	if err := global.GVA_DB.Debug().Model(&tables.TOperationReport{}).Where("which_day BETWEEN ? AND ?", end.SubDays(req.Days), end).Find(&operationReports).Error; err != nil {
		return res, err
	}

	resDays := make([]string, len(operationReports))
	resTotalUseCounts := make([]decimal.Decimal, len(operationReports))
	resBinaryUseCounts := make([]decimal.Decimal, len(operationReports))
	resRandomUseCounts := make([]decimal.Decimal, len(operationReports))
	resTrinityUseCounts := make([]decimal.Decimal, len(operationReports))
	resCelticCrossUseCounts := make([]decimal.Decimal, len(operationReports))
	for i := 0; i < len(operationReports); i++ {
		resDays[i] = operationReports[i].WhichDay.String()
		resTotalUseCounts[i] = operationReports[i].DayBinaryDivinationCoin.Add(operationReports[i].DayRandomDivinationCoin).Add(operationReports[i].DayTrinityDivinationCoin).Add(operationReports[i].DayCelticCrossDivinationCoin)
		resBinaryUseCounts[i] = operationReports[i].DayBinaryDivinationCoin
		resRandomUseCounts[i] = operationReports[i].DayRandomDivinationCoin
		resTrinityUseCounts[i] = operationReports[i].DayTrinityDivinationCoin
		resCelticCrossUseCounts[i] = operationReports[i].DayCelticCrossDivinationCoin
	}

	res.DayList = resDays
	res.TotalUseCountList = resTotalUseCounts
	res.DayBinaryUseCountList = resBinaryUseCounts
	res.DayRandomUseCountList = resRandomUseCounts
	res.DayTrinityUseCountList = resTrinityUseCounts
	res.DayCelticCrossUseCountList = resCelticCrossUseCounts
	return res, nil
}

// @author: billbliu
// @function: GetCountryMemberChart
// @description: 国家会员数图表数据
// @return: res response.GetCountryMemberChartResponse, err error
func (e *OperationService) GetCountryMemberChart(req request.GetCountryMemberChartRequest) (response.GetCountryMemberChartResponse, error) {
	if req.TopSeveral == 0 {
		req.TopSeveral = 10
	}

	res := response.GetCountryMemberChartResponse{}

	type Result struct {
		CountryID   int
		ChineseName string
		Count       int
	}

	var results []Result

	if err := global.GVA_DB.Debug().Table("t_user").
		Select("t_user.country_id, t_country.chinese_name, COUNT(*) as count").
		Joins("left join t_country on t_country.id = t_user.country_id").
		Group("t_user.country_id, t_country.chinese_name").Order("count desc").
		Scan(&results).Error; err != nil {
		return res, err
	}

	count := req.TopSeveral
	if len(results) < count {
		count = len(results)
	}

	countryList := make([]string, count)
	memberCountList := make([]int64, count)

	for i := 0; i < count; i++ {
		countryList[i] = results[i].ChineseName
		memberCountList[i] = int64(results[i].Count)
	}

	res.CountryList = countryList
	res.MemberCountList = memberCountList

	return res, nil
}

// @author: billbliu
// @function: GetAllCountryMemberChart
// @description: 获取所有国家会员数图表数据
// @return: []response.CountryMemberInfo, error
func (e *OperationService) GetAllCountryMemberChart() ([]response.CountryMemberInfo, error) {
	type Result struct {
		CountryID   int
		ChineseName string
		Count       int64
		Latitude    float32
		Longitude   float32
		Code        uint
	}

	var results []Result

	if err := global.GVA_DB.Debug().Table("t_country").
		Select("t_user.country_id, t_country.chinese_name, COUNT(t_user.country_id) as count, t_country.latitude, t_country.longitude, t_country.code").
		Joins("left join t_user  on t_country.id = t_user.country_id").
		Where("t_country.latitude != 0").
		Group("t_user.country_id, t_country.chinese_name, t_country.latitude, t_country.longitude, t_country.code").Order("count desc").
		Scan(&results).Error; err != nil {
		return nil, err
	}

	res := make([]response.CountryMemberInfo, len(results))

	for i := 0; i < len(results); i++ {
		res[i].CountryCode = results[i].Code
		res[i].CountryName = results[i].ChineseName
		res[i].MemberCount = results[i].Count
		res[i].Latitude = results[i].Latitude
		res[i].Longitude = results[i].Longitude
	}

	return res, nil
}

// @author: billbliu
// @function: GetMemberDepositUsdtChart
// @description: 会员充值Usdt图表数据
// @return: res response.GetMemberDepositUsdtChartResponse, err error
func (e *OperationService) GetMemberDepositUsdtChart(req request.GetMemberDepositUsdtChartRequest) (response.GetMemberDepositUsdtChartResponse, error) {
	if req.TopSeveral == 0 {
		req.TopSeveral = 10
	}

	res := response.GetMemberDepositUsdtChartResponse{}

	userAssets := []tables.TUserAsset{}
	if err := global.GVA_DB.Debug().Model(&tables.TUserAsset{}).Where("total_usdt != 0").Order("total_usdt desc").Find(&userAssets).Error; err != nil {
		return res, err
	}

	count := req.TopSeveral
	if len(userAssets) < count {
		count = len(userAssets)
	}

	userNameList := make([]string, count)
	depositUsdtCountList := make([]decimal.Decimal, count)

	for i := 0; i < count; i++ {
		userNameList[i] = strconv.FormatUint(uint64(userAssets[i].UserId), 10)
		depositUsdtCountList[i] = userAssets[i].TotalUsdt
	}

	res.UserNameList = userNameList
	res.DepositUsdtCountList = depositUsdtCountList

	return res, nil
}

// @author: billbliu
// @function: GetMemberInviteChart
// @description: 获取用户邀请图表数据
// @return: res response.GetMemberInviteChartResponse, err error
func (e *OperationService) GetMemberInviteChart(req request.GetMemberInviteChartRequest) (response.GetMemberInviteChartResponse, error) {
	if req.TopSeveral == 0 {
		req.TopSeveral = 10
	}

	res := response.GetMemberInviteChartResponse{}

	type Result struct {
		InvitationUserId int64
		Count            int64
	}

	var results []Result

	if err := global.GVA_DB.Debug().Table("t_user").
		Select("invitation_user_id, COUNT(*) as count").Where("invitation_user_id != 0").
		Group("invitation_user_id").Order("count desc").
		Scan(&results).Error; err != nil {
		return res, err
	}

	count := req.TopSeveral
	if len(results) < count {
		count = len(results)
	}

	userNameList := make([]string, count)
	inviteCountList := make([]int64, count)

	for i := 0; i < count; i++ {
		userNameList[i] = strconv.FormatInt(results[i].InvitationUserId, 10)
		inviteCountList[i] = results[i].Count
	}

	res.UserNameList = userNameList
	res.InviteCountList = inviteCountList

	return res, nil
}

// @author: billbliu
// @function: GetMemberRebateChart
// @description: 获取用户返利图表数据
// @return: res response.GetMemberRebateChartResponse, err error
func (e *OperationService) GetMemberRebateChart(req request.GetMemberRebateChartRequest) (response.GetMemberRebateChartResponse, error) {
	if req.TopSeveral == 0 {
		req.TopSeveral = 10
	}

	res := response.GetMemberRebateChartResponse{}

	userAssets := []tables.TUserAsset{}
	if err := global.GVA_DB.Debug().Model(&tables.TUserAsset{}).Where("total_return_cash != 0").
		Order("total_return_cash desc").Find(&userAssets).Error; err != nil {
		return res, err
	}

	count := req.TopSeveral
	if len(userAssets) < count {
		count = len(userAssets)
	}

	userNameList := make([]string, count)
	rebateCountList := make([]decimal.Decimal, count)

	for i := 0; i < count; i++ {
		userNameList[i] = strconv.FormatUint(uint64(userAssets[i].UserId), 10)
		rebateCountList[i] = userAssets[i].TotalReturnCash
	}

	res.UserNameList = userNameList
	res.RebateCountList = rebateCountList

	return res, nil
}

// @author: billbliu
// @function: GetCountryDepositUsdtChart
// @description: 获取国家充值图表数据
// @return: res response.GetCountryDepositUsdtChartResponse, err error
func (e *OperationService) GetCountryDepositUsdtChart(req request.GetCountryDepositUsdtChartRequest) (response.GetCountryDepositUsdtChartResponse, error) {
	if req.TopSeveral == 0 {
		req.TopSeveral = 10
	}

	res := response.GetCountryDepositUsdtChartResponse{}

	type Result struct {
		CountryID   int
		ChineseName string
		Count       decimal.Decimal
	}

	var results []Result

	if err := global.GVA_DB.Debug().Table("t_user").
		Select("t_user.country_id, t_country.chinese_name, SUM(t_user_asset.total_usdt) as count").
		Joins("left join t_country on t_country.id = t_user.country_id").
		Joins("left join t_user_asset on t_user_asset.user_id = t_user.id").
		Group("t_user.country_id, t_country.chinese_name").Order("count desc").
		Scan(&results).Error; err != nil {
		return res, err
	}

	count := req.TopSeveral
	if len(results) < count {
		count = len(results)
	}

	countryNameList := make([]string, 0)
	depositUsdtCountList := make([]decimal.Decimal, 0)

	for i := 0; i < count; i++ {
		if results[i].Count.Equal(decimal.Zero) {
			break
		}
		countryNameList = append(countryNameList, results[i].ChineseName)
		depositUsdtCountList = append(depositUsdtCountList, results[i].Count)
	}

	res.CountryNameList = countryNameList
	res.DepositUsdtCountList = depositUsdtCountList

	return res, nil
}

// @author: billbliu
// @function: GetDivinationPercentChart
// @description: 获取国家充值图表数据
// @return: []response.GetCountryDepositUsdtCharGetDivinationPercentChartResponsetResponse, error
func (e *OperationService) GetDivinationPercentChart() ([]response.GetDivinationPercentChartResponse, error) {

	type Result struct {
		TotalBinary      decimal.Decimal
		TotalRandom      decimal.Decimal
		TotalTrinity     decimal.Decimal
		TotalCelticCross decimal.Decimal
	}
	var result Result

	end := utils.CarbonNow().SubDay()
	// operationReports := []tables.TOperationReport{}
	if err := global.GVA_DB.Debug().Model(&tables.TOperationReport{}).
		Select("SUM(day_binary_divination_coin) as total_binary, SUM(day_random_divination_coin) as total_random,SUM(day_trinity_divination_coin) as total_trinity, SUM(day_celtic_cross_divination_coin) as total_celtic_cross").
		Where("which_day BETWEEN ? AND ?", end.SubYears(30), end).Scan(&result).Error; err != nil {
		return nil, err
	}

	total := result.TotalBinary.Add(result.TotalRandom).Add(result.TotalTrinity).Add(result.TotalCelticCross)

	res := make([]response.GetDivinationPercentChartResponse, 0)
	res = append(res, response.GetDivinationPercentChartResponse{
		CardFormationName: "二选一",
		Percent:           result.TotalBinary.Div(total).Mul(decimal.NewFromInt(100)).Round(2),
	})
	res = append(res, response.GetDivinationPercentChartResponse{
		CardFormationName: "无牌阵",
		Percent:           result.TotalRandom.Div(total).Mul(decimal.NewFromInt(100)).Round(2),
	})
	res = append(res, response.GetDivinationPercentChartResponse{
		CardFormationName: "圣三角",
		Percent:           result.TotalTrinity.Div(total).Mul(decimal.NewFromInt(100)).Round(2),
	})
	res = append(res, response.GetDivinationPercentChartResponse{
		CardFormationName: "凯尔特十字",
		Percent:           result.TotalCelticCross.Div(total).Mul(decimal.NewFromInt(100)).Round(2),
	})

	return res, nil
}
