package domain

//Collection Dept Mediation service
type CollectionDetailRequest struct {
	IDCardNo            string `json:"IDCardNo"       validate:"required,max=20"`
	RedCaseNo 			string `json:"RedCaseNo"      validate:"max=15"`
	BlackCaseNo 		string `json:"BlackCaseNo"    validate:"max=15"`
}

type CollectionDetailResponse struct {
	IDCardNo            string	                      `json:"IDCardNo"` 
	NoOfAgreement       int 	                      `json:"NoOfAgreement"`
	AgreementList       []CollectionDetailAgreement   `json:"AgreementList"`
}

// type CollectionDetailResponse struct - nested structure for Collection details
type CollectionDetailAgreement struct {
	AgreementNo     			string 	`json:"AgreementNo"`
	SeqOfAgreement      		int 	`json:"SeqOfAgreement"`
	OutsourceID 				string 	`json:"OutsourceID"`
	OutsourceName     			string 	`json:"OutsourceName"`
	BlockCode      				string 	`json:"BlockCode"`
	CurrentSUEOSPrincipalNet 	float64 `json:"CurrentSUEOSPrincipalNet"`
	CurrentSUEOSPrincipalVAT    float64 `json:"CurrentSUEOSPrincipalVAT"`
	CurrentSUEOSInterestNet     float64 `json:"CurrentSUEOSInterestNet"`
	CurrentSUEOSInterestVAT 	float64 `json:"CurrentSUEOSInterestVAT"`
	CurrentSUEOSPenalty     	float64 `json:"CurrentSUEOSPenalty"`
	CurrentSUEOSHDCharge      	float64 `json:"CurrentSUEOSHDCharge"`
	CurrentSUEOSOtherFee 		float64 `json:"CurrentSUEOSOtherFee"`
	CurrentSUEOSTotal     		float64 `json:"CurrentSUEOSTotal"`
	TotalPaymentAmount      	float64 `json:"TotalPaymentAmount"`
	LastPaymentDate 			int 	`json:"LastPaymentDate"`
	SUESeqNo     				int 	`json:"SUESeqNo"`
	BeginSUEOSPrincipalNet      float64 `json:"BeginSUEOSPrincipalNet"`
	BeginSUEOSPrincipalVAT 		float64 `json:"BeginSUEOSPrincipalVAT"`
	BeginSUEOSInterestNet     	float64 `json:"BeginSUEOSInterestNet"`
	BeginSUEOSInterestVAT      	float64 `json:"BeginSUEOSInterestVAT"`
	BeginSUEOSPenalty 			float64 `json:"BeginSUEOSPenalty"`
	BeginSUEOSHDCharge     		float64 `json:"BeginSUEOSHDCharge"`
	BeginSUEOSOtherFee     		float64 `json:"BeginSUEOSOtherFee"`
	BeginSUEOSTotal 			float64 `json:"BeginSUEOSTotal"`
	SUEStatus     				int		`json:"SUEStatus"`
	SUEStatusDescription      	string 	`json:"SUEStatusDescription"`
	BlackCaseNo 				string 	`json:"BlackCaseNo"`
	BlackCaseDate     			int		`json:"BlackCaseDate"`
	RedCaseNo      				string 	`json:"RedCaseNo"`
	RedCaseDate 				int		`json:"RedCaseDate"`
	CourtCode     				string 	`json:"CourtCode"`
	CourtName      				string 	`json:"CourtName"`
	JudgmentDate 				int 	`json:"JudgmentDate"`
	JudgmentResultCode     		int 	`json:"JudgmentResultCode"`
	JudgmentResultDescription   string 	`json:"JudgmentResultDescription"`
	JudgmentDetail 				string 	`json:"JudgmentDetail"`
	ExpectDate     				int 	`json:"ExpectDate"`
	AssetPrice      			float64 `json:"AssetPrice"`
	JudgeAmount 				float64 `json:"JudgeAmount"`
	NoOfInstallment     		string 	`json:"NoOfInstallment"`
	InstallmentAmount      		float64 `json:"InstallmentAmount"`
	TotalCurrentPerSUESeqNo 	float64 `json:"TotalCurrentPerSUESeqNo"`
}

type CollectionLogRequest struct {
	AgreementNo         string `json:"AgreementNo"  validate:"required,max=16"`
	RemarkCode          string `json:"RemarkCode"   validate:"required,max=4"`
	LogRemark1 			string `json:"LogRemark1"   validate:"max=120"`
	LogRemark2     		string `json:"LogRemark2"   validate:"max=120"`
	LogRemark3       	string `json:"LogRemark3"   validate:"max=120"`
	LogRemark4 			string `json:"LogRemark4"   validate:"max=120"`
	LogRemark5     		string `json:"LogRemark5"   validate:"max=120"`
	InputDate           string `json:"InputDate"    validate:"required,max=12"`
	InputTime           string `json:"InputTime"    validate:"required,max=6"`
	OperatorID          string `json:"OperatorID"   validate:"required,max=15"`
}

type CollectionLogResponse struct {
	IDCardNo 			string  `json:"IDCardNo"`
	AgreementNo 		string  `json:"AgreementNo"`	
}
