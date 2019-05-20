package params
type StudentId int64
type Page struct {
	PageNum int    `form:"page_num" json:"page_num"`
	Offset  int    `form:"offset" json:"offset"`
	Source  string `form:"source" json:"source"`
}

func (a *Page) HandlePage() {
	if a.PageNum < 1 {
		a.PageNum = 1
	}
	if a.Offset < 1 || a.Offset > 10000 {
		a.Offset = 20
	}
}
