package activity

import (
	"errors"
	"fmt"
	"github.com/apex/log"
	"milano.gaodun.com/model/activity"
	ms "milano.gaodun.com/model/members_student"
	"milano.gaodun.com/pkg/error-code"
)

type ActivityServiceInterface interface {
	List() []activity.PriActivity
	Add(m *activity.PriActivity) (int64, error)
	AddContact(c *ContactInfo) int
	GetContact(uid int64) (*ContactInfo, int)
}

type ActivityService struct {
	activityM *activity.ActivityModel
	auM       *activity.ActivityUserModel
	msM       *ms.MembersStudentModel
	logger    *log.Entry
}
type ContactInfo struct {
	MsgCode   string
	QQCode    string
	Phone     string
	Uid       int64
	Source    string
	ProjectId int64
}

func NewActivityService(p *activity.ActivityParam, l *log.Entry) ActivityServiceInterface {
	a := ActivityService{activityM: activity.NewActivityModel(p), logger: l}
	a.msM = ms.NewMembersStudentModel()
	a.auM = activity.NewActivityUserModel()
	return &a
}

// 查询
func (a *ActivityService) List() []activity.PriActivity {
	data, err := a.activityM.List()
	if err != nil {
		a.logger.Info(err.Error())
	}

	return data
}

// 新增
func (a *ActivityService) Add(m *activity.PriActivity) (int64, error) {
	if a.IsExist() {
		return 0, errors.New("Error data exist")
	}
	row, err := a.activityM.Add(m)
	if err != nil {
		a.logger.Info(err.Error())
		return row, err
	}
	id := m.Id
	*m = activity.PriActivity{}
	a.activityM.Id(id).Get(m)

	return row, err
}
func (a *ActivityService) GetContact(uid int64) (*ContactInfo, int) {
	contactInfo := ContactInfo{}
	auList, err := a.auM.GetAll(uid)
	if err != nil {
		return &contactInfo, error_code.ACTIVITYUSERMODELERR
	}
	for _, v := range auList {
		if v.NumType == 1 {
			contactInfo.MsgCode = v.Number
		}
		if v.NumType == 2 {
			contactInfo.QQCode = v.Number
		}
		if v.NumType == 3 {
			contactInfo.Phone = v.Number
		}
		contactInfo.Source = v.Source
		contactInfo.ProjectId = v.ProjectId
		contactInfo.Uid = v.Uid
	}
	return &contactInfo, error_code.SUCCESSSTATUS
}

// 新增
func (a *ActivityService) AddContact(m *ContactInfo) int {
	actList := []activity.PriActivityUser{}
	gms, err := a.msM.Get(m.Uid)
	if err != nil {
		return error_code.MEMBERSSTUDENTMODELERR
	}
	if gms.Id == 0 {
		return error_code.MEMBERSSTUDENTINFOEMPTY
	}
	if m.MsgCode != "" {
		au := activity.PriActivityUser{
			Number:    m.MsgCode,
			NumType:   1,
			ProjectId: m.ProjectId,
			Uid:       m.Uid,
			Source:    m.Source,
		}
		actList = append(actList, au)
	}
	if m.QQCode != "" {
		au := activity.PriActivityUser{
			Number:    m.QQCode,
			NumType:   2,
			ProjectId: m.ProjectId,
			Uid:       m.Uid,
			Source:    m.Source,
		}
		actList = append(actList, au)
	}
	if gms.Phone != "" {
		au := activity.PriActivityUser{
			Number:    gms.Phone,
			NumType:   3,
			ProjectId: m.ProjectId,
			Uid:       m.Uid,
			Source:    m.Source,
		}
		actList = append(actList, au)
	}
	auList, err := a.auM.GetAll(m.Uid)
	if err != nil {
		return error_code.ACTIVITYUSERMODELERR
	}
	contactMap := map[string]activity.PriActivityUser{}
	for _, v := range auList {
		mapKey := fmt.Sprintf("%d_%d", v.Uid, v.NumType)
		contactMap[mapKey] = v
	}
	for _, v := range actList {
		mapKey := fmt.Sprintf("%d_%d", v.Uid, v.NumType)
		contact, ok := contactMap[mapKey]
		if ok && contact.Number != "" {
			a.auM.Edit(&v)
		} else {
			a.auM.Add(&v)
		}
	}
	return error_code.SUCCESSSTATUS
}
func (a *ActivityService) IsExist() bool {
	ok, err := a.activityM.IsExist()
	if err != nil {
		a.logger.Info(err.Error())
	}
	return ok
}
