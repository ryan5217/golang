package study_rights

import (
	"github.com/go-xorm/xorm"
	"milano.gaodun.com/pkg/utils"
)

type TkStudyRightsManage struct {
	Id          int64
	ProjectId   int64
	SubjectId   int64
	GradationId int64 `xorm:"stage"`
	Description string
	ModuleId    int64  `xorm:"type"`
	GoodsIds    string `xorm:"course_ids"`
	CourseIds   string `xorm:"required_courses"`
	PaperType   int64
	ResourceId  int64 `xorm:"syllabus_id"`
	Status      int64
	CreatedTime string
	UpdatedTime string
}
type SearchParam struct {
	ProjectId int64
	SubjectId int64
	Uid       int64
}
type StudyRightsManageModel struct {
	*xorm.Engine
	s *xorm.Session
}

func NewStudyRightsManageModel() *StudyRightsManageModel {
	return &StudyRightsManageModel{Engine: utils.NewtikuDb}
}
func (g *StudyRightsManageModel) Get(id int64) (*TkStudyRightsManage, error) {
	srm := TkStudyRightsManage{}
	_, err := g.Id(id).Get(&srm)
	return &srm, err
}
func (g *StudyRightsManageModel) List(param SearchParam) (*[]TkStudyRightsManage, error) {
	srm := []TkStudyRightsManage{}
	err := g.Where("project_id=?", param.ProjectId).
		Where("subject_id=?", param.SubjectId).Find(&srm)
	return &srm, err
}
func (g *StudyRightsManageModel) SubjectList(projectIdList []string) (map[int64][]int64, error) {
	srm := []TkStudyRightsManage{}
	err := g.In("project_id", projectIdList).Where("status<>?", 2).Where("status<>?", 4).Find(&srm)
	res := map[int64][]int64{}
	for _, v := range projectIdList {
		res[utils.String2Int64(v)] = []int64{}
	}
	subjectExist := map[int64]bool{}
	for _, v := range srm {
		if !subjectExist[v.SubjectId] {
			subjectExist[v.SubjectId] = true
			res[v.ProjectId] = append(res[v.ProjectId], v.SubjectId)
		}
	}
	return res, err
}
