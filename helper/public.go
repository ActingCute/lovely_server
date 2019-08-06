package helper

import (
	"fmt"
	"github.com/astaxie/beego"
	"math/rand"
	"os"
	"runtime"
	"strconv"
	"strings"
	"time"
	"encoding/hex"
	"crypto/md5"
)

var (
	IsDebug = GetBool(GetAppConf("IsDebug"))
	TimeZone = "Asia/Chongqing"
	Layout = "2006-01-02 15:04:05"
)

//error
func Error(err ...interface{}) bool {
	if err[0] != nil {
		_, files, line, ok := runtime.Caller(1)
		if !ok {
			fmt.Println(fmt.Errorf("Error : Cant not print!"))
			return true
		}
		fs := strings.Split(files, "/")
		file := ""
		file = fs[0]
		if len(fs) > 2 {
			file = fs[len(fs)-2] + "/" + fs[len(fs)-1]
		}
		fileline := "[" + file + ":" + strconv.Itoa(line) + "]"
		go beego.Error(fileline, err, "\r\n")
		return true
	}
	return false
}

//debug
func Debug(debug ...interface{}) {
	if IsDebug {
		_, files, line, ok := runtime.Caller(1)
		if !ok {
			fmt.Println(fmt.Errorf("Error: Cant not print!"))
			return
		}
		fs := strings.Split(files, "/")
		file := ""
		file = fs[0]
		if len(fs) > 2 {
			file = fs[len(fs)-2] + "/" + fs[len(fs)-1]
		}
		fileline := "[" + file + ":" + strconv.Itoa(line) + "]"
		go beego.Debug(fileline, debug, "\r\n")
	}
}

//system

func IsExist(filename string) bool {
	_, err := os.Stat(filename)
	return err == nil || os.IsExist(err)

}

// config
func GetAppConf(name string) string {
	return beego.AppConfig.String(name)
}

func GetBool(str string) bool {

	boolean := false

	if str == "1" || str == "true" {
		boolean = true
	}

	return boolean

}

const (
	KC_RAND_KIND_NUM   = 0 // 纯数字
	KC_RAND_KIND_LOWER = 1 // 小写字母
	KC_RAND_KIND_UPPER = 2 // 大写字母
	KC_RAND_KIND_ALL   = 3 // 数字、大小写字母
)

func GetToken(size int, kind int) string {
	ikind, kinds, result := kind, [][]int{[]int{10, 48}, []int{26, 97}, []int{26, 65}}, make([]byte, size)
	is_all := kind > 2 || kind < 0
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < size; i++ {
		if is_all { // random ikind
			ikind = rand.Intn(3)
		}
		scope, base := kinds[ikind][0], kinds[ikind][1]
		result[i] = uint8(base + rand.Intn(scope))
	}
	return string(result)
}

func Md5(str string) string {

	if len(str) == 0 {
		return str
	}

	md := md5.Sum([]byte(str))
	mdpw := hex.EncodeToString(md[:])
	return mdpw

}

//随机字符
func GetRandomString(strlen int) string {
	str := "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	bytes := []byte(str)
	result := []byte{}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < strlen; i++ {
		result = append(result, bytes[r.Intn(len(bytes))])
	}
	return string(result)
}

//string时间戳转化int64
func StringDateFormatInt(t string, layout ...string) (unix_time int64) {
	beego.Info("StringDateFormatInt = ", t)
	layoutstr := Layout
	if len(layout) > 0 {
		layoutstr = layout[0]
	}
	datetime, err := time.Parse(layoutstr, t)
	if Error(err) {
		return 0
	}
	return datetime.Unix()
}