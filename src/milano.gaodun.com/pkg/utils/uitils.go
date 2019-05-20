package utils

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"regexp"
	"strconv"
	"strings"
	"time"

	"math/rand"

	"bytes"

	"crypto/sha1"

	"github.com/gin-gonic/gin"
	"gitlab.gaodun.com/golib/filetool"
	"milano.gaodun.com/conf"
	"sort"
)

var (
	reLineBreak = regexp.MustCompile("\n")
	strPol      = "ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789abcdefghijklmnopqrstuvwxyz"
)

func GetStatus(c *gin.Context) {
	dir, _ := os.Getwd()
	fileOperate := filetool.FileOperate{Filename: dir + string(os.PathSeparator) + "DEPLOY"}
	if fileOperate.CheckFileExist(fileOperate.Filename) {
		text, _ := fileOperate.ReadFile(fileOperate.Filename)
		lineList := strings.Split(string(text), "\n")
		begin := "{"
		tmpStr := ""
		for _, line := range lineList {
			if len(line) == 0 {
				break
			}
			n := strings.Split(line, "|")
			jsonKey := "\"" + n[0] + "\""
			jsonValue := ":\"" + n[1] + "\""
			tmpStr += jsonKey + jsonValue + ","
		}
		mStr := strings.TrimRight(tmpStr, ",")
		end := "}"
		info := begin + mStr + end
		c.String(http.StatusOK, "{\"status\": \"1\",\"data\":"+info+"}")
		return
	}

	c.String(http.StatusOK, "{\"status\": \"1\"}")
	return
}

// PathExists ...
func PathExists(path string) (bool, error) {
	dir, _ := os.Getwd() //当前的目录
	path = dir + path
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

// MakeDir ...
func MakeDir(path string) error {
	dir, _ := os.Getwd() //当前的目录
	err := os.Mkdir(dir+path, os.ModePerm)
	return err
}

// ChangeStruct2OtherStruct ...
func ChangeStruct2OtherStruct(item interface{}, toItem interface{}) {

	j, _ := json.Marshal(item)
	json.Unmarshal(j, &toItem)
}

// ChangeRedis2OtherStruct ...
func ChangeRedis2OtherStruct(item interface{}) []uint8 {
	// var toItem []uint8
	// j, _ := json.Marshal(item)
	// json.Unmarshal(j, &toItem)
	return item.([]uint8)
}

// ChangeUint82OtherStruct ...
func ChangeUint82OtherStruct(item interface{}, toItem interface{}) {
	j, _ := json.Marshal(item)
	json.Unmarshal(j, &toItem)
}

// ChangeByteStruct2OtherStruct ...
func ChangeByteStruct2OtherStruct(str []byte, toItem interface{}) error {

	if erri := json.Unmarshal([]byte(str), &toItem); erri != nil && toItem == nil {
		return erri
	}
	return nil
}

// ChangeArrayString2Int ...
func ChangeArrayString2Int(before []string) []int {
	after := make([]int, len(before))
	for k, item := range before {
		after[k], _ = strconv.Atoi(item)
	}
	return after
}

// ChangeArrayString2Int64 ...
func ChangeArrayString2Int64(before []string) []int64 {
	after := make([]int64, len(before))
	for k, item := range before {
		after[k], _ = strconv.ParseInt(item, 10, 64)
	}
	return after
}

// JoinInt64Array2String 拆分int64的数组为string
func JoinInt64Array2String(before []int64, sep string) string {
	after := ""
	for _, item := range before {
		after = after + strconv.Itoa(int(item)) + sep
	}
	return strings.Trim(after, sep)
}

// String2Int ...
func String2Int(before string) int {
	after, _ := strconv.Atoi(before)
	return after
}

// String2Int64 ...
func String2Int64(before string) int64 {
	after, _ := strconv.ParseInt(before, 10, 64)
	return after
}

// String2Float64 ...
func String2Float64(before string) float64 {
	after, _ := strconv.ParseFloat(before, 64)
	return after
}

// String2Int32 ...
func String2Int32(before string) int32 {
	after, _ := strconv.Atoi(before)
	return int32(after)
}

// String2Int8 ...
func String2Int8(before string) int8 {
	after, _ := strconv.Atoi(before)
	return int8(after)
}

// Slise2Map ....
func Slise2Map(originData []string) map[string]int {
	retData := make(map[string]int, 0)
	if originData == nil {
		return retData
	}
	// count := len(originData)
	for i, v := range originData {
		retData[v] = i
	}
	// for i := 0; i < count; i++ {
	// 	// if mapContains(retData, originData[i]) { //前后取前

	// 	// }
	// }
	return retData
}

//  InArrayInt64 是否存在数组中 int64
func InArrayInt64(arr []int64, val int64) bool {
	flag := false
	for _, item := range arr {
		if item == val {
			flag = true
			break
		}
	}
	return flag
}

//  InArrayInt  是否存在数组中 int
func InArrayInt(arr []int, val int) bool {
	flag := false
	for _, item := range arr {
		if item == val {
			flag = true
			break
		}
	}
	return flag
}

// InArrayString 是否存在数组中 字符串
func InArrayString(arr []string, val string) bool {
	flag := false
	for _, item := range arr {
		if item == val {
			flag = true
			break
		}
	}
	return flag
}

//判断key是否存在
func MapContains(needMap map[string]int, key string) bool {
	if _, ok := needMap[key]; ok {
		return true
	}
	return false
}

// FormartDate2Time
func FormartDate2Time(dataTimeStr, ms string) int64 {
	dataTime, _ := time.Parse(ms, dataTimeStr)
	return dataTime.Unix()
}

func StrToUnix(ms string) int64 {
	timeLayout := "2006-01-02 15:04:05"  //转化所需模板
	loc, _ := time.LoadLocation("Local") //获取时区
	tmp, _ := time.ParseInLocation(timeLayout, ms, loc)
	return tmp.Unix()
}

// ABCToRune ...
func ABCToRune(abc string) rune {
	abcrune := []rune(abc)
	return abcrune[0]
}

func ContactInterfaceMap(params ...map[string]interface{}) map[string]interface{} {
	reParam := make(map[string]interface{})
	for _, param := range params {
		for k, v := range param {
			reParam[k] = v
		}
	}
	return reParam
}

func ContactStrMap(params ...map[string]string) map[string]string {
	reParam := make(map[string]string)
	for _, param := range params {
		for k, v := range param {
			reParam[k] = v
		}
	}
	return reParam
}

func IntSliceToStrSlice(array []int) []string {
	var re []string
	for _, v := range array {
		id := strconv.Itoa(v)
		re = append(re, id)
	}
	return re
}

func GenRandStr(strLen int) string {
	var buffer bytes.Buffer
	max := len(strPol) - 1

	for i := 0; i < strLen; i++ {
		index := rand.Intn(max)
		buffer.WriteByte(strPol[index])
	}

	return buffer.String()
}

func Signature(timestamp, nonce, token string) string {
	a := []string{
		timestamp,
		nonce,
		token,
	}
	sort.Strings(a)
	str := strings.Join(a, "")

	var buffer bytes.Buffer
	buffer.WriteString(str)

	sha1Hash := sha1.New()
	sha1Hash.Write(buffer.Bytes())
	data := sha1Hash.Sum(nil)
	return fmt.Sprintf("%x", data)
}
func GetRandomString(l int) string {
	str := "0123456789abcdefghijklmnopqrstuvwxyz"
	bytes := []byte(str)
	result := []byte{}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < l; i++ {
		result = append(result, bytes[r.Intn(len(bytes))])
	}
	return strings.ToUpper(string(result))
}

func GetAvatarUrl(avatar string) string {
	var resAvatar string
	if avatar == "" {
		resAvatar = "http://simg01.gaodunwangxiao.com/v/Uploads/avatar/default.jpg"
	} else {
		resAvatar = conf.SIMG_DOMAIN + "/v" + avatar
	}
	return resAvatar
}

func VersionCompare(oldVersion string, newVersion string) bool {
	if oldVersion == "" {
		return true
	}
	if newVersion == "" {
		return false
	}
	oldStr := strings.Split(oldVersion,".")
	newStr := strings.Split(newVersion,".")
	length := len(newStr)
	if length > len(oldStr) {
		length = len(oldStr)
	}
	for i := 0; i < length; i++ {
		newInt, _ := strconv.ParseInt(newStr[i], 10, 64)
		oldInt, _ := strconv.ParseInt(oldStr[i], 10, 64)
		if newInt > oldInt {
			return true
		} else if i < length-1 {
			continue
		} else {
			return false
		}
	}
	return false
}
