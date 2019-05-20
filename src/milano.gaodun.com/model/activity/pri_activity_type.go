package activity

type PriActivity struct {
	ActivityParam `xorm:"extends"`
	ActState      int32  `xorm:"default 1"`
	Remark        string `json:"Remark,omitempty"`
	CreatedTime   string `xorm:"created"`
	UpdatedTime   string `xorm:"updated"`
}

type ActivityParam struct {
	Id             int64
	StudentId      int64
	ActName        string
	ActType        int32 `xorm:"default 1"`
	ProjectId      int64
	SubjectId      int64
	ForceUpdateCol []string `xorm:"-" json:"-"`
}
