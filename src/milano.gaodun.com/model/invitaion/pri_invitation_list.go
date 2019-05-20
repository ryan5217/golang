package invitaion

import (
	"github.com/go-xorm/xorm"
	"milano.gaodun.com/pkg/utils"
)

type PriInvitationList struct {
	Id             int64
	StudentId      int64
	InvitedStudent int64
	Status         int32  `xorm:"-"`
	Message        string `xorm:"-"`
}

type InvitationlistModel struct {
	*xorm.Engine
	s   *xorm.Session
	pic PriInvitationList
}

func NewInvitationlistModel() *InvitationlistModel {
	return &InvitationlistModel{Engine: utils.GaodunPrimaryDb}
}
func (b *InvitationlistModel) GetInvit(invitedStudent int64) (PriInvitationList, error) {
	invitation_list := PriInvitationList{}

	_, err := b.Where("invited_student=?", invitedStudent).Get(&invitation_list)
	return invitation_list, err
}
func (b *InvitationlistModel) GetInvitationList(studentId int64) ([]PriInvitationList, error) {
	invitation_list := []PriInvitationList{}

	err := b.Where("student_id=?", studentId).Asc("created_time").Find(&invitation_list)
	return invitation_list, err
}

//保存邀请码
func (b *InvitationlistModel) Add(invitL *PriInvitationList) (int64, error) {
	row, err := b.InsertOne(invitL)
	return row, err
}
