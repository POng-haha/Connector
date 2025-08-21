package domain

// ---------- API CheckDopa ---------
type CheckDOPARequest struct {
	IDCardNo          string `json:"IDCardNo"    validate:"required,max=20"`
	FirstName         string `json:"FirstName"   validate:"required,max=50"`
	LastName          string `json:"LastName"    validate:"required,max=50"`
	BirthDate         string `json:"BirthDate"   validate:"required,max=8"`
	LaserID           string `json:"LaserID"     validate:"required,max=20"`
	Reference         string `json:"Reference"   validate:"max=50"`
}

type CheckDOPAResponse struct {
	Result            bool   `json:"Result"`
	Reference         string `json:"Reference"`
}

