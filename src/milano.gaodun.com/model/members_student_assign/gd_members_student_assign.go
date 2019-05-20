package members_student_assign

import (
	"github.com/go-xorm/xorm"
	"milano.gaodun.com/pkg/utils"
	"time"
)

type GdMembersStudentAssign struct {
	Id          int64
	StudentId   int64
	CourseId    int64
	InvalidTime int64
}

type MembersStudentAssignModel struct {
	*xorm.Engine
	s *xorm.Session
}

func NewMembersStudentAssignModel() *MembersStudentAssignModel {
	return &MembersStudentAssignModel{Engine: utils.GaodunDb2}
}

func (g *MembersStudentAssignModel) GetAssigns(uid int64, coursesIds []int64) ([]GdMembersStudentAssign, error) {
	courseAssigns := []GdMembersStudentAssign{}
	err := g.Where("student_id=?", uid).Where("(invalid_time = 0 or invalid_time >?)", time.Now().Unix()).In("course_id", coursesIds).Find(&courseAssigns)
	return courseAssigns, err
}
