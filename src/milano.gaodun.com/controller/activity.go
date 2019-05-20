package controller

import (
	"github.com/gin-gonic/gin"
	actM "milano.gaodun.com/model/activity"
	"milano.gaodun.com/pkg/error-code"
	"milano.gaodun.com/pkg/setting"
	actSer "milano.gaodun.com/service/activity"
	"regexp"
)

var ActivityApi = NewActivityApi()

func NewActivityApi() *Activity {
	return &Activity{}
}

type Activity struct {
	Base
}

// 查询
func (a Activity) GetList(c *gin.Context) {
	param := actM.ActivityParam{}

	param.Id = a.GetInt64(c, "id")
	param.SubjectId = a.GetInt64(c, "subject_id")
	param.ProjectId = a.GetInt64(c, "project_id")
	param.StudentId = a.GetInt64(c, "student_id")
	param.ActType = a.GetInt32(c, "act_type")
	param.ActName = a.GetString(c, "act_name")

	act := actSer.NewActivityService(&param, setting.GinLogger(c))
	data := act.List()

	a.ServerJSONSuccess(c, &data)
}
func (a Activity) AddUser(c *gin.Context) {
	qq, _ := regexp.Compile("[1-9][0-9]{4,14}")
	phone, _ := regexp.Compile("^(13[0-9]|14[579]|15[0-3,5-9]|16[6]|17[0135678]|18[0-9]|19[89])\\d{8}$")
	act_user := actM.PriActivityUser{}
	act_user.Number = a.GetString(c, "number")
	act_user.NumType = a.GetInt(c, "num_type")
	act_user.ProjectId = a.GetInt64(c, "project_id")
	if act_user.ProjectId == 0 {
		act_user.ProjectId = 14
	}
	if (act_user.NumType == 2 && !qq.MatchString(act_user.Number)) ||
		(act_user.NumType == 3 && !phone.MatchString(act_user.Number)) ||
		act_user.NumType > 5 || act_user.NumType < 1 {
		a.ServerJSONError(c, nil, error_code.PARAMETERERR)
		return
	}

	act_userM := actM.NewActivityUserModel()
	_, err := act_userM.Add(&act_user)
	if err != nil {
		setting.GinLogger(c).Info(err.Error())
		a.ServerJSONError(c, nil, error_code.FAIL)
	} else {
		a.ServerJSONSuccess(c, act_user)
	}
}
func (a Activity) GetContact(c *gin.Context) {
	uid := a.GetInt64(c, "uid", true)
	auS := actSer.NewActivityService(&actM.ActivityParam{}, setting.GinLogger(c))
	res,code := auS.GetContact(uid)
	if code != error_code.SUCCESSSTATUS {
		a.ServerJSONError(c, nil, code)
	} else {
		a.ServerJSONSuccess(c, res)
	}
}
func (a Activity) AddContact(c *gin.Context) {
	qq, _ := regexp.Compile("[1-9][0-9]{4,14}")
	act_user := actSer.ContactInfo{}
	act_user.Uid = a.GetInt64(c, "uid", true)
	act_user.MsgCode = a.GetString(c, "msg_code")
	act_user.QQCode = a.GetString(c, "qq_code")
	act_user.Source = a.GetString(c, "source", true)
	act_user.ProjectId = a.GetInt64(c, "project_id")
	if act_user.ProjectId == 0 {
		act_user.ProjectId = 14
	}
	if (act_user.MsgCode == "" && !qq.MatchString(act_user.QQCode)) ||
		(act_user.MsgCode == "" && act_user.QQCode == "") {
		a.ServerJSONError(c, nil, error_code.PARAMETERERR)
		return
	}

	act_userS := actSer.NewActivityService(&actM.ActivityParam{}, setting.GinLogger(c))
	code := act_userS.AddContact(&act_user)
	if code != error_code.SUCCESSSTATUS {
		a.ServerJSONError(c, nil, code)
	} else {
		a.ServerJSONSuccess(c, "添加成功")
	}
}
func (a Activity) Add(c *gin.Context) {
	pri := actM.PriActivity{}

	pri.ProjectId = a.GetInt64(c, "project_id", true)
	pri.ActName = a.GetString(c, "act_name", true)
	pri.SubjectId = a.GetInt64(c, "subject_id", true)
	pri.StudentId = a.GetInt64(c, "student_id", true)
	pri.ActType = a.GetInt32(c, "act_type")
	pri.ActState = a.GetInt32(c, "act_state")
	pri.Remark = a.GetString(c, "remark")

	par := pri.ActivityParam
	par.ForceUpdateCol = a.PostMustCols(c, true)
	// must be judged field if a required item
	if c.GetBool(Verify) {
		return
	}

	priS := actSer.NewActivityService(&par, setting.GinLogger(c))
	_, err := priS.Add(&pri)
	if err != nil {
		setting.GinLogger(c).Info(err.Error())
		a.ServerJSONError(c, nil, error_code.FAIL)
	} else {
		a.ServerJSONSuccess(c, pri)
	}
}
