package models

import (
	"fmt"
	"github.com/EDDYCJY/go-gin-demo/pkg/setting"
	"github.com/jinzhu/gorm"
	"log"
	"time"
)

var db *gorm.DB

func Setup() {
	var err error
	// 创建gorm实例 -> 数据库表映射
	db, err = gorm.Open(
		setting.DatabaseSetting.Type,
		// 该语句最后返回string-> 顺便打印sql语句输出 -> 创建链路
		fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local",
			setting.DatabaseSetting.User,
			setting.DatabaseSetting.Password,
			setting.DatabaseSetting.Host,
			setting.DatabaseSetting.Name),
	)

	if err != nil {
		log.Fatalf("model.Setup err: %v", err)
	}
	// 默认表名处理 -> 此处做表名拼接 -> 最终会根据使用哪个表返回model
	gorm.DefaultTableNameHandler = func(db *gorm.DB, defaultTableName string) string {
		return setting.DatabaseSetting.TablePrefix + defaultTableName
	}
	// 是否默认使用单表
	db.SingularTable(true)
	// 设置注册时间(避免分布式系统的时间带来很多问题，采取时间统一方式)
	db.Callback().Create().Replace("gorm:update_time_stamp", updateTimeStampForCreateCallback)
	db.Callback().Update().Replace("gorm:update_time_stamp", updateTimeStampForUpdateCallback)
	// 删除原有的，注册新的回调方式
	// 设置删除回滚能力
	db.Callback().Delete().Replace("gorm:delete", deleteCallback)

	// 给数据库限制连接数
	db.DB().SetMaxIdleConns(10)
	db.DB().SetMaxOpenConns(100)
}

// 给表的创建、更新时间赋值
func updateTimeStampForCreateCallback(scope *gorm.Scope) {
	if !scope.HasError() {
		nowTime := time.Now().Unix()
		// 查询FieldByName 中的表字段，是否空白。
		if createTimeField, ok := scope.FieldByName("CreatedOn"); ok {
			if createTimeField.IsBlank {
				createTimeField.Set(nowTime)
			}
		}

		if modifyTimeField, ok := scope.FieldByName("ModifiedOn"); ok {
			// 查询FieldByName 中的表字段，是否空白。
			if modifyTimeField.IsBlank {
				modifyTimeField.Set(nowTime)
			}
		}
	}
}

func deleteCallback(scope *gorm.Scope) {
	if !scope.HasError() {
		var extraOption string
		if str, ok := scope.Get("gorm: delete_option"); ok {
			extraOption = fmt.Sprint(str)
		}
		// 查找已删除的数据
		deleteOnField, hasDeleteOnField := scope.FieldByName("DeletedOn")

		// 做一些处理 如果搜索到该表且表已删除 -> 给当前表设置删除时间
		if !scope.Search.Unscoped && hasDeleteOnField {
			scope.Raw(fmt.Sprintf(
				"UPDATE %v SET %v=%v%v%v",
				scope.QuotedTableName(),
				scope.Quote(deleteOnField.DBName),
				scope.AddToVars(time.Now().Unix()),
				addExtraSpaceIfExist(scope.CombinedConditionSql()),
				addExtraSpaceIfExist(extraOption),
			))
		} else {
			// 删除该表
			scope.Raw(fmt.Sprintf("DELETED from :%v%v%v",
				scope.QuotedTableName(),
				addExtraSpaceIfExist(scope.CombinedConditionSql()),
				addExtraSpaceIfExist(extraOption),
			)).Exec()
		}
	}
}

// updateTimeStampForUpdateCallback will set `ModifiedOn` when updating
func updateTimeStampForUpdateCallback(scope *gorm.Scope) {
	if _, ok := scope.Get("gorm:update_column"); !ok {
		scope.SetColumn("ModifiedOn", time.Now().Unix())
	}
}

// 做一些SQL的空格处理
func addExtraSpaceIfExist(str string) string {
	if str != "" {
		return " " + str
	}
	return ""
}
