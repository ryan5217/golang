package questionnaire

import (
	"github.com/apex/log"
	"milano.gaodun.com/model/questionnaire"
	//"encoding/json"
)

type QuestionnaireServiceInterface interface {
	AddQuestionnaire(qu *questionnaire.PriQuestionnaire) *questionnaire.PriQuestionnaire
	Edit(qu *questionnaire.PriQuestionnaire) (int64, error)
}

type QuestionnaireService struct {
	QuestionnaireM *questionnaire.QuestionnaireModel
	logger         *log.Entry
}

func NewQuestionnaireService(logger *log.Entry) QuestionnaireServiceInterface {
	return &QuestionnaireService{QuestionnaireM: questionnaire.NewQuestionnaireModel(), logger: logger}
}
func (i *QuestionnaireService) AddQuestionnaire(qu *questionnaire.PriQuestionnaire) *questionnaire.PriQuestionnaire {
	_, err := i.QuestionnaireM.Add(qu)
	if err != nil {
		i.logger.Error(err.Error())
		return qu
	}
	return qu
}

//保存更新内容
func (i *QuestionnaireService) Edit(qu *questionnaire.PriQuestionnaire) (int64, error) {
	row, err := i.QuestionnaireM.Edit(qu)
	if err != nil {
		i.logger.Error(err.Error())
	}
	i.QuestionnaireM.Id(qu.Id).Get(qu)
	return row, err
}
