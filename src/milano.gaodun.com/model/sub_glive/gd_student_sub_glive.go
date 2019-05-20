package banner

import (
	"github.com/go-xorm/xorm"
	"milano.gaodun.com/pkg/utils"
)

type GdStudentSubGlive struct {
	Id        int64
	StudentId int64
	GliveId   int64
}
type GdStudentSubGliveModel struct {
	*xorm.Engine
	s *xorm.Session
}

func NewGdStudentSubGliveModel() *GdStudentSubGliveModel {
	return &GdStudentSubGliveModel{Engine: utils.GaodunDb}
}

// 获取列表 按条件获取
func (g *GdStudentSubGliveModel) GetCount(GliveId int64, ProjectId int64) (int64, error) {
	var r GdStudentSubGlive
	cou, err := g.Where("glive_id=?", GliveId).Where("project_id=?", ProjectId).Count(r)
	return cou, err
}
