package primary

import (
	"github.com/go-xorm/xorm"
	"milano.gaodun.com/pkg/utils"
)

type PriUserStage struct {
	Id        int64
	StudentId int64
	Stage     int
	Locked    int
}

type UserStageModel struct {
	*xorm.Engine
	s   *xorm.Session
	pus PriUserStage
}

func NewUserStageModel() *UserStageModel {
	return &UserStageModel{Engine: utils.GaodunPrimaryDb}
}

func (b *UserStageModel) Edit(pus *PriUserStage) (int64, error) {
	return b.Where("student_id=?", pus.StudentId).Update(pus)
}
func (b *UserStageModel) Add(pus *PriUserStage) (int64, error) {
	row, err := b.InsertOne(pus)
	return row, err
}
func (b *UserStageModel) GetOne(studentId int64) (*PriUserStage, error) {
	user_stage := PriUserStage{}
	_, err := b.Where("student_id=?", studentId).Get(&user_stage)
	return &user_stage, err
}
