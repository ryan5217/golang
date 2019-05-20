package exercise_record

import (
	"github.com/go-xorm/xorm"
	"milano.gaodun.com/pkg/setting"
	"milano.gaodun.com/pkg/utils"
)

type ExerciseRecordModel struct {
	*xorm.Engine
	s *xorm.Session
}
type TkExerciseRecord struct {
	Id          int64  `json:"id"`
	Uid         int64  `json:"uid"`
	ResourceId  int64  `json:"resource_id"`
	Type        int64  `json:"type"`
	ModuleId    int64  `json:"module_id"`
	ModuleName  string `json:"module_name"`
	ChapterName string `json:"chapter_name"`
	SubjectName string `json:"subject_name" xorm:"-"`
}

func NewExerciseRecordModel() *ExerciseRecordModel {
	return &ExerciseRecordModel{Engine: utils.NewtikuDb}
}

func (g *ExerciseRecordModel) Add(er *TkExerciseRecord) (int64, error) {
	num, err := g.InsertOne(er)
	if err != nil {
		setting.Logger.Infof("ExerciseRecordModel_Add_%s", err.Error())
	}
	return num, err
}
func (g *ExerciseRecordModel) Update(er *TkExerciseRecord) (int64, error) {
	num, err := g.Id(er.Id).Cols("module_id", "module_name", "resource_id", "type","chapter_name").Update(er)
	if err != nil {
		setting.Logger.Infof("ExerciseRecordModel_Update_%s", err.Error())
	}
	return num, err
}
func (g *ExerciseRecordModel) GetOne(uid int64) (*TkExerciseRecord, error) {
	er := TkExerciseRecord{}
	_, err := g.Where("uid=?", uid).Get(&er)
	if err != nil {
		setting.Logger.Infof("ExerciseRecordModel_GetOne_%s", err.Error())
	}
	return &er, err
}
