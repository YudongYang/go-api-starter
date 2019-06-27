package init

import (
	_ "api-starter/models"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"

	//_ "github.com/mattn/go-sqlite3"
	// _ github.com/lib/pq
	utils "api-starter/utils"

	_ "github.com/go-sql-driver/mysql"
)

//初始化数据连接
func init() {
	//读取配置文件，设置数据库参数
	//数据库类别
	dbType := beego.AppConfig.String("db.type")
	//连接名称
	dbAlias := beego.AppConfig.String(dbType + "::db.alias")
	//数据库名称
	dbName := beego.AppConfig.String(dbType + "::db.name")
	//数据库连接用户名
	dbUser := beego.AppConfig.String(dbType + "::db.user")
	//数据库连接用户名
	dbPwd := beego.AppConfig.String(dbType + "::db.pwd")
	//数据库IP（域名）
	dbHost := beego.AppConfig.String(dbType + "::db.host")
	//数据库端口
	dbPort := beego.AppConfig.String(dbType + "::db.port")
	switch dbType {
	case "sqlite3":
		orm.RegisterDataBase(dbAlias, dbType, dbName)
	case "mysql":
		dbCharset := beego.AppConfig.String(dbType + "::db.charset")
		orm.RegisterDataBase(dbAlias, dbType, dbUser+":"+dbPwd+"@tcp("+dbHost+":"+
			dbPort+")/"+dbName+"?charset="+dbCharset, 30)
	}
	//如果是开发模式，则显示命令信息
	isDev := (beego.AppConfig.String("runmode") == "dev")
	//自动建表
	orm.RunSyncdb("default", false, isDev)
	utils.Logger.Info("Init: Database(%s): %s", isDev, dbHost)
	if isDev {
		orm.Debug = isDev
	}
}
