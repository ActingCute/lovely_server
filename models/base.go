package models

import (
	"github.com/astaxie/beego"
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
	helper "lovely_server/helper"
	"net"
	"os"
	"strconv"
	"time"
)

var (
	ShowSql  bool //是否显示sql日志
	db       *xorm.Engine
	ds       *xorm.Session
	DBOk     bool //数据库连接是否正常
	Domain   string
	SaltCode string //盐
)

var (
	App          *app   = new(app)
	QiNiu        *qiniu = new(qiniu)
	DefaultAdmin *admin = new(admin)
)

const DELETE_TIME_IS_NULL = " delete_time is null"

type qiniu struct {
	Enable            bool
	AccessKey         string
	SecretKey         string
	Space             string
	Domain            string
	Videoautochange   bool
	VideoTypeToChange string
	Pipeline          string
	NotifyURL         string
	Action            string
}

type app struct {
	Runmode string
	IsDebug bool
	ApiUrl  string
}

type admin struct {
	UserName string
	PassWord string
}

func init() {
	//web
	Domain = GetAppConf("web::domain")
	helper.Debug("Domain :", Domain)
	//密码
	SaltCode = GetAppConf("password::salt")
	if len(SaltCode) < 10 {
		SaltCode = "e50831141013dd23a9b07b7fdad333a4"
	}
	//admin
	DefaultAdmin.UserName = GetAppConf("admin::user_name")
	if len(DefaultAdmin.UserName) < 5 {
		DefaultAdmin.UserName = "ActingCute酱"
	}
	DefaultAdmin.PassWord = GetAppConf("admin::pass_word")
	if len(DefaultAdmin.PassWord) < 5 {
		DefaultAdmin.PassWord = "190025254"
	}
	//数据库
	dbHost := GetAppConf("db::host")
	daTable := GetAppConf("db::table")
	dbPort := GetAppConf("db::port")
	dbUser := GetAppConf("db::user")
	dbPassword := GetAppConf("db::password")
	ShowSql = GetAppConfBool("db::showsql")
	//若获取失败，则用默认值
	dbHost = If(len(dbHost) < 1, "127.0.0.1", dbHost).(string)
	daTable = If(len(daTable) < 1, "falcon_portal", daTable).(string)
	dbPort = If(len(dbPort) < 1, "3306", dbPort).(string)
	dbUser = If(len(dbUser) < 1, "root", dbUser).(string)
	dataSourceName := dbUser + ":" + dbPassword + "@tcp(" + dbHost + ":" + dbPort + ")/" + daTable + "?charset=utf8&parseTime=true&loc=Local"
	go connentDB(dataSourceName)
	//qiniu
	QiNiu.AccessKey = beego.AppConfig.String("qiniu::accessKey")
	QiNiu.SecretKey = beego.AppConfig.String("qiniu::secretKey")
	QiNiu.Space = beego.AppConfig.String("qiniu::space")
	QiNiu.Domain = beego.AppConfig.String("qiniu::domain")
	QiNiu.VideoTypeToChange = beego.AppConfig.String("qiniu::videotypetochange")
	QiNiu.Pipeline = beego.AppConfig.String("qiniu::pipeline")
	QiNiu.NotifyURL = beego.AppConfig.String("qiniu::notifyurl")
	QiNiu.Action = beego.AppConfig.String("qiniu::action")
	helper.Debug("QiNiu - ", QiNiu)

	//app
	// App Config Init
	App.Runmode = beego.AppConfig.String("runmode")
	if beego.AppConfig.String("isdebug") == "true" {
		App.IsDebug = true
	} else {
		App.IsDebug = false
	}
	App.ApiUrl = "http://" + GetIp() + ":" + strconv.Itoa(beego.BConfig.Listen.HTTPPort)

	beego.Info(App)
	beego.Info(QiNiu)
	beego.Info(DefaultAdmin)
}

func connentDB(dataSourceName string) {
	t1 := time.NewTimer(time.Second * 1)
	timer := 1
	for {
		select {
		case <-t1.C:
			timer++
			if timer > 15 {
				beego.Info("多次尝试连接数据库失败，请检查数据库配置信息！")
				break
			}
			beego.Debug("timer --- ", timer)
			var err error
			db, err = xorm.NewEngine("mysql", dataSourceName)
			if helper.Error(err) {
				beego.Error("1.数据库初始化失败 ----", err)
				t1.Reset(time.Second * 10)
				continue
			}
			err = db.Ping()
			if helper.Error(err) {
				beego.Error("2.数据库初始化失败 ----", err)
				t1.Reset(time.Second * 10)
				continue
			}
			maxOpenConns := GetAppConfInt("db::maxOpenConns")
			maxIdleConns := GetAppConfInt("db::maxIdleConns")
			db.SetMaxOpenConns(maxOpenConns)
			db.SetMaxIdleConns(maxIdleConns)

			db.ShowSQL(ShowSql)
			db.DatabaseTZ = time.Local
			db.TZLocation = time.Local
			// LRU 缓存
			cacher := xorm.NewLRUCacher(xorm.NewMemoryStore(), 10000)
			db.SetDefaultCacher(cacher)
			DBOk = true
			InitDB()
			if ShowSql {
				//initDBLogger()
			} else {
				beego.Info("sql 日志已经关闭")
			}
			break
		}
	}
}

func InitDB() {
	//所有的业务都应该等待数据表同步完成再执行
	//只同步表结构
	//db.Table("activate").Sync2(new(Activate))
	err := db.Sync2(
		new(Comment),
		new(Count),
		new(User),
		new(Twitter),
		new(Admin),
	)
	if !helper.Error(err) {
		initDBLogger()
		initDefaultValue()
	}
}

func initDefaultValue() {
	defAdmin := Admin{
		UserName: DefaultAdmin.UserName,
		PassWord: DefaultAdmin.PassWord}
	AddAdmin(&defAdmin)
}

func initDBLogger() {
	t1 := time.NewTimer(time.Second * 1)
	logpath := "./Log/sql.log"
	go func() {
		for {
			select {
			case <-t1.C:
				beego.Info("初始化SQL日志 --- ok")
				if helper.IsExist(logpath) {
					f, err := os.OpenFile(logpath, os.O_RDWR, 0775)
					if helper.Error(err) {
						beego.Error("open sql log error:", err)
						return
					}
					os.Stdout = f
					logger := xorm.NewSimpleLogger(os.Stdout)
					db.SetLogger(logger)
					break
				} else {
					beego.Info("初始化SQL日志 --- ing")
					t1.Reset(time.Second * 5)
				}
			}
		}
	}()

}

func If(condition bool, trueVal, falseVal interface{}) interface{} {
	if condition {
		return trueVal
	}
	return falseVal
}

func GetAppConf(name string) string {
	return beego.AppConfig.String(name)
}

func GetAppConfWithDefault(name, def string) (ret string) {
	ret = beego.AppConfig.String(name)
	if len(ret) == 0 {
		ret = def
	}
	return ret
}

func GetAppConfInt(name string) int {
	num, _ := beego.AppConfig.Int(name)
	return num
}

func GetAppConfInt64(name string) int64 {
	num, _ := beego.AppConfig.Int64(name)
	return num
}

func GetAppConfBool(name string) bool {
	isok, err := beego.AppConfig.Bool(name)
	if helper.Error(err) {
		isok = false
	}
	return isok
}

func Insert(data interface{}) error {
	_, err := db.Omit("delete_time").Insert(data)
	helper.Error(err)
	return err
}

func Update(data interface{}, Where string, whereData []interface{}, cols ...string) error {
	helper.Debug("Update whereData -- ", whereData)
	_, err := db.Where(Where, whereData...).Cols(cols...).Update(data)
	helper.Error(err)
	return err
}

func Delete(data interface{}, Where string, whereData []interface{}) error {
	helper.Debug("Delete whereData -- ", whereData)
	_, err := db.Where(Where, whereData...).Delete(data)
	helper.Error(err)
	return err
}

//获取本地IP
func GetIp() string {
	if App.IsDebug {
		return "127.0.0.1"
	}
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return "127.0.0.1"
	}
	for _, address := range addrs {
		// 检查ip地址判断是否回环地址
		if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				return ipnet.IP.String()
			}
		}
	}
	return "127.0.0.1"
}

func GetPassword(password string, uid int64) string {
	pass := helper.Md5(strconv.FormatInt(uid, 10) + password + SaltCode)
	return pass
}
