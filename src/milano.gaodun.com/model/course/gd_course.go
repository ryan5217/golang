package course

import (
	"github.com/go-xorm/xorm"
	"milano.gaodun.com/pkg/setting"
	"milano.gaodun.com/pkg/utils"
)

type GdCourse struct {
	Id             int64
	Name           string
	Isbig          int64
	Courses        string
	RelationCourse int64
	Studentsnum    int64
	Istasks        int64
}

type CourseModel struct {
	*xorm.Engine
	s *xorm.Session
}

func NewCourseModel() *CourseModel {
	return &CourseModel{Engine: utils.GaodunDb2}
}

func (g *CourseModel) Get(courseId int64) (*GdCourse, error) {
	course := GdCourse{}
	_, err := g.Id(courseId).Get(&course)
	if err != nil {
		setting.Logger.Infof("CourseModel_Get_%s", err.Error())
	}
	return &course, err
}
func (g *CourseModel) GetCourseBySaasCourseId(saasId int64) (*GdCourse, error) {
	course := GdCourse{}
	_, err := g.Where("relation_course=?", saasId).Get(&course)
	if err != nil {
		setting.Logger.Infof("CourseModel_GetCourseBySaasCourseId_%s", err.Error())
	}
	return &course, err
}
func (g *CourseModel) GetCourseByIds(courseIds string) (*[]GdCourse, error) {
	courses := []GdCourse{}
	err := g.Where("id in(" + courseIds + ")").Find(&courses)
	if err != nil {
		setting.Logger.Infof("CourseModel_GetCourseByIds_%s", err.Error())
	}
	return &courses, err
}
