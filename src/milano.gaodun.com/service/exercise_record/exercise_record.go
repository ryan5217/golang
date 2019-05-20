package exercise_record

import (
	"fmt"
	"github.com/apex/log"
	tk "milano.gaodun.com/model/api/tiku"
	em "milano.gaodun.com/model/exercise_record"
	sr "milano.gaodun.com/model/study_rights"
	sb "milano.gaodun.com/model/subject"
	"milano.gaodun.com/pkg/error-code"
	"milano.gaodun.com/pkg/utils"
)

type ExerciseRecordService struct {
	erm    *em.ExerciseRecordModel
	srm    *sr.StudyRightsModulesModel
	tka    *tk.PaperRecordApi
	sb     *sb.SubjectModel
	logger *log.Entry
}

func NewExerciseRecordService(logger *log.Entry) *ExerciseRecordService {
	var g ExerciseRecordService
	h := utils.HttpHandle
	g.logger = logger
	g.erm = em.NewExerciseRecordModel()
	g.srm = sr.NewStudyRightsModulesModel()
	g.tka = tk.NewPaperRecordApi(h)
	g.sb = sb.NewSubjectModel()
	return &g
}

type PaperRecord struct {
	Exercise em.TkExerciseRecord `json:"exercise"`
	Resource interface{}         `json:"resource"`
}

func (g *ExerciseRecordService) AddRecord(record *em.TkExerciseRecord) int {
	rec, err := g.erm.GetOne(record.Uid)
	if err != nil {
		return error_code.DBERR
	}
	if record.ModuleId > 0 {
		mm, err := g.srm.Get(record.ModuleId)
		if err != nil {
			return error_code.DBERR
		}
		record.ModuleName = mm.Name
	}
	if rec.Id > 0 {
		record.Id = rec.Id
		_, err = g.erm.Update(record)
	} else {
		_, err = g.erm.Add(record)
	}
	if err != nil {
		return error_code.DBERR
	}
	return error_code.SUCCESSSTATUS
}
func (g *ExerciseRecordService) GetRecord(uid int64) (interface{}, int) {
	res := PaperRecord{Resource:map[int64]int64{}}
	rec, err := g.erm.GetOne(uid)
	if rec.ResourceId == 0 {
		return &res, error_code.SUCCESSSTATUS
	}
	if err != nil {
		return &res, error_code.DBERR
	}
	mm, err := g.srm.Get(rec.ModuleId)
	if err != nil {
		return &res, error_code.DBERR
	}
	sb, err := g.sb.Get(mm.SubjectId)
	res.Exercise = *rec
	res.Exercise.SubjectName = sb.Name
	if err != nil {
		return map[int64]int64{}, error_code.DBERR
	}
	//试卷类型学习记录，需要获取做题状态
	if rec.Type == 2 {
		ps, code := g.tka.GetPaperStatus(uid, mm.ProjectId, mm.SubjectId, fmt.Sprintf("%d", rec.ResourceId))
		index := len(ps.MustArray())
		res.Resource = ps.GetIndex(index-1)
		return &res, code
	}
	return &res, error_code.SUCCESSSTATUS
}
