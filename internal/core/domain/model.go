package domain

	//Collection Dept Mediation service

type CollectionDetailRequest struct {
	IDCardNo 			string `json:"IDCardNo" 	validate:"required"`
	RedCaseNo 			string `json:"RedCaseNo"`
	BlackCaseNo 		string `json:"BlackCaseNo"`
}

type CollectionDetailResponse struct {
	IDCardNo			string	`len:"20"  json:"IDCardNo"` 
	NoOfAgreement 		int 	`len:"2"   json:"NoOfAgreement"`
	AgreementList       []CollectionDetailAgreement   `json:"AgreementList"`
}

// type CollectionDetailResponse struct - nested structure for Collection details
type CollectionDetailAgreement struct {
	AgreementNo     			int 	`len:"16"   json:"AgreementNo"`
	SeqOfAgreement      		int 	`len:"2"    json:"SeqOfAgreement"`
	OutsourceID 				string 	`len:"4"  	json:"OutsourceID"`
	OutsourceName     			string 	`len:"30"   json:"OutsourceName"`
	BlockCode      				string 	`len:"2"   	json:"BlockCode"`
	CurrentSUEOSPrincipalNet 	float64 `len:"10"   json:"CurrentSUEOSPrincipalNet"`
	CurrentSUEOSPrincipalVAT    float64 `len:"10"   json:"CurrentSUEOSPrincipalVAT"`
	CurrentSUEOSInterestNet     float64 `len:"10"   json:"CurrentSUEOSInterestNet"`
	CurrentSUEOSInterestVAT 	float64 `len:"10"   json:"CurrentSUEOSInterestVAT"`
	CurrentSUEOSPenalty     	float64 `len:"9"    json:"CurrentSUEOSPenalty"`
	CurrentSUEOSHDCharge      	float64 `len:"9"    json:"CurrentSUEOSHDCharge"`
	CurrentSUEOSOtherFee 		float64 `len:"9"    json:"CurrentSUEOSOtherFee"`
	CurrentSUEOSTotal     		float64 `len:"10"   json:"CurrentSUEOSTotal"`
	TotalPaymentAmount      	float64 `len:"10"   json:"TotalPaymentAmount"`
	LastPaymentDate 			int 	`len:"8"   	json:"LastPaymentDate"`
	SUESeqNo     				int 	`len:"2"   	json:"SUESeqNo"`
	BeginSUEOSPrincipalNet      float64 `len:"10"   json:"BeginSUEOSPrincipalNet"`
	BeginSUEOSPrincipalVAT 		float64 `len:"10"   json:"BeginSUEOSPrincipalVAT"`
	BeginSUEOSInterestNet     	float64 `len:"10"   json:"BeginSUEOSInterestNet"`
	BeginSUEOSInterestVAT      	float64 `len:"10"   json:"BeginSUEOSInterestVAT"`
	BeginSUEOSPenalty 			float64 `len:"10"   json:"BeginSUEOSPenalty"`
	BeginSUEOSHDCharge     		float64 `len:"9"    json:"BeginSUEOSHDCharge"`
	BeginSUEOSOtherFee     		float64 `len:"9"    json:"BeginSUEOSOtherFee"`
	BeginSUEOSTotal 			float64 `len:"10"   json:"BeginSUEOSTotal"`
	SUEStatus     				int		`len:"2"   	json:"SUEStatus"`
	SUEStatusDescription      	string 	`len:"30"   json:"SUEStatusDescription"`
	BlackCaseNo 				string 	`len:"15"   json:"BlackCaseNo"`
	BlackCaseDate     			int		`len:"8"   	json:"BlackCaseDate"`
	RedCaseNo      				string 	`len:"15"   json:"RedCaseNo"`
	RedCaseDate 				int		`len:"8"   	json:"RedCaseDate"`
	CourtCode     				string 	`len:"4"   	json:"CourtCode"`
	CourtName      				string 	`len:"30"   json:"CourtName"`
	JudgmentDate 				int 	`len:"8"  	json:"JudgmentDate"`
	JudgmentResultCode     		int 	`len:"1"   	json:"JudgmentResultCode"`
	JudgmentResultDescription   string 	`len:"40"   json:"JudgmentResultDescription"`
	JudgmentDetail 				string 	`len:"500"  json:"JudgmentDetail"`
	ExpectDate     				int 	`len:"8"   	json:"ExpectDate"`
	AssetPrice      			float64 `len:"10"   json:"AssetPrice"`
	JudgeAmount 				float64 `len:"10"  	json:"JudgeAmount"`
	NoOfInstallment     		string 	`len:"3"   	json:"NoOfInstallment"`
	InstallmentAmount      		float64 `len:"10"   json:"InstallmentAmount"`
	TotalCurrentPerSUESeqNo 	float64 `len:"11"  	json:"TotalCurrentPerSUESeqNo"`
}

type CollectionLogRequest struct {
	CollectionLogAgreementRq       []CollectionLogAgreementRq   `json:"AgreementList"`
}

// CollectionLogRequest - nested structure for CollectionLog details
type CollectionLogAgreementRq struct {
	AgreementNo     	int    `json:"AgreementNo" 	validate:"required"`
	RemarkCode       	string `json:"RemarkCode" 	validate:"required"`
	LogRemark1 			string `json:"LogRemark1"`
	LogRemark2     		string `json:"LogRemark2"`
	LogRemark3       	string `json:"LogRemark3"`
	LogRemark4 			string `json:"LogRemark4"`
	LogRemark5     		string `json:"LogRemark5"`
	InputDate       	string `json:"InputDate" 	validate:"required"`
	InputTime 			string `json:"InputTime" 	validate:"required"`
	OperatorID     		string `json:"OperatorID" 	validate:"required"`
}

type CollectionLogResponse struct {
	IDCardNo 			string  `len:"20"   json:"IDCardNo"`
	AgreementList    	[]CollectionLogAgreementRs	`json:"AgreementList"`	
}

// type CollectionLogResponse struct - nested structure for CollectionLog details
type CollectionLogAgreementRs struct {
	AgreementNo 		int  	`len:"16"   json:"AgreementNo"`
	LogRemarkStatus    	string	`len:"2"    json:"LogRemarkStatus"`	
}