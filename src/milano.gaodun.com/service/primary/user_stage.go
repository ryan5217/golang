package primary

import (
	"github.com/apex/log"
	"milano.gaodun.com/model/api/primary"
)

type UserStageService struct {
	usm    *primary.UserStageModel
	logger *log.Entry
}

func NewUserStageService(logger *log.Entry) *UserStageService {
	var g UserStageService
	g.logger = logger
	g.usm = primary.NewUserStageModel()
	return &g
}

func (g *UserStageService) Save(us *primary.PriUserStage) (*primary.PriUserStage, error) {
	pus := primary.PriUserStage{}
	g.usm.Where("student_id=?", us.StudentId).Get(&pus)
	if us.Id > 0 || pus.Id > 0 {
		//if(pus.Id > 0){
		//	us.Id = pus.Id
		//}
		_, err := g.usm.Edit(us)
		g.usm.Id(pus.Id).Get(us)
		return us, err
	} else {
		_, err := g.usm.Add(us)
		g.usm.Id(us.Id).Get(us)
		return us, err
	}
}
func (g *UserStageService) GetStage(studentId int64) (*primary.PriUserStage, error) {
	user_stage, err := g.usm.GetOne(studentId)
	return user_stage, err
}
