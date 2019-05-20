package crm_order_push

const (
	CrmName        = "有赞项目负责-杨璇"
	CrmId          = 2821
	CrmPayType     = 8 // crm 给的值 不要随意改
	CrmPayTypeName = "学习卡(微店)"
)

type PostData struct {
	ProductList  []string `json:"productList"`
	DiscountList []string `json:"discountList"`
	StudentInfo  struct {
		StudentGuid     string
		UcenterUid      int
		TrueName        string
		Telphone        string
		Email           string
		CertificateType int
		CertificateNo   string
		IDcardFace      string
		IDcardBack      string
		UrgentName      string
		UrgentPhone     string
		SignUpSchoolID  int
		AttendSchoolID  int
		StartTime       string
		Create_Time     string
		Create_By       int
	} `json:"studentInfo"`
	OrderInfo struct {
		CourseType    int
		OrderType     int
		Remark        string
		Source        int
		SellPrise     string
		DiscountPrise int
	} `json:"orderInfo"`
	ClueInfo struct {
		City           string `json:"city"`
		Province       string `json:"province"`
		AgentName      string
		CustomerSource string
		CrmId          int
		ErpContactId   string `json:"erpContactId"`
		FollowRecords  string
	} `json:"clueInfo"`
	SerialList []Serial `json:"serialList"`
	VOrderNo   string
}

type Serial struct {
	OrderNo     string
	PaySerialNo string
	PayCount    string
	PayType     int
	PayTypeName string
	PayAccount  string
	PayTime     string
	Remark      string
}

func newPostData() PostData {
	p := PostData{}
	p.DiscountList = []string{}
	p.StudentInfo.CertificateType = 1000201
	p.StudentInfo.SignUpSchoolID = 83
	p.StudentInfo.AttendSchoolID = 83
	p.StudentInfo.Create_By = 1
	p.OrderInfo.OrderType = 1000400
	p.OrderInfo.Source = 1000443
	p.ClueInfo.AgentName = CrmName
	p.ClueInfo.CustomerSource = "网校用户/网校社群"
	p.ClueInfo.CrmId = CrmId
	p.ClueInfo.ErpContactId = "31288"
	p.ClueInfo.FollowRecords = "-有赞营销"
	p.SerialList = []Serial{}

	return p
}
