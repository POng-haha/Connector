package format

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"connectorapi-go/internal/core/domain"
)

func PadOrTruncate(s string, length int) string {
	runes := []rune(s)
	if len(runes) > length {
		runes = runes[:length]
	}
	padded := string(runes)
	if len(runes) < length {
		padded = padded + strings.Repeat(" ", length-len(runes))
	}
	return padded
}

// helper for trimming and converting string to int
func parseIntField(raw string) int {
	raw = strings.TrimSpace(raw)
	i, _ := strconv.Atoi(raw)
	return i
}

// helper for trimming and converting string to float64
func parseFloatField(raw string) float64 {
	raw = strings.TrimSpace(raw)
	f, _ := strconv.ParseFloat(raw, 64)
	return f
}

// BuildFixedLengthHeader constructs the fixed-length header.
func BuildFixedLengthHeader(
	routeSystem, routeService, routeFormat, requestID string, fixedLengthData string,
) string {
	now := time.Now()
	requestDate := now.Format("20060102")
	requestTime := now.Format("150405")

	responseCode := PadOrTruncate("", 6)
	responseMessage := PadOrTruncate("", 50)

	const fixedHeaderLength = 10 + 15 + 3 + 20 + 8 + 6 + 5 + 6 + 50 // 123 characters
	totalMessageLength := fixedHeaderLength + len(fixedLengthData)
	requestLengthStr := fmt.Sprintf("%05d", totalMessageLength)

	header := PadOrTruncate(routeSystem, 10) +
		PadOrTruncate(routeService, 15) +
		PadOrTruncate(routeFormat, 3) +
		PadOrTruncate(requestID, 20) +
		PadOrTruncate(requestDate, 8) +
		PadOrTruncate(requestTime, 6) +
		requestLengthStr +
		responseCode +
		responseMessage

	return header
}

// Converts CollectionDetailRequest to a fixed-length string.
func FormatCollectionDetailRequest(reqData domain.CollectionDetailRequest) string {
	IDCardNo := PadOrTruncate(reqData.IDCardNo, 20)
	RedCaseNo := PadOrTruncate(reqData.RedCaseNo, 15)
	BlackCaseNo := PadOrTruncate(reqData.BlackCaseNo, 15)
	return IDCardNo + RedCaseNo + BlackCaseNo
}

func FormatCollectionDetailResponse(respData string) domain.CollectionDetailResponse {
	const headerLen = 123
	body := respData[headerLen:]

	idCardNo := strings.TrimSpace(body[0:20])
	noOfAgreement := parseIntField(body[20:22])

	const agreementLen = 942
	agreementStart := 22
	agreements := make([]domain.CollectionDetailAgreement, 0, noOfAgreement)

	for i := 0; i < noOfAgreement; i++ {
		start := agreementStart + i*agreementLen
		end := start + agreementLen
		if end > len(body) {
			break
		}
		block := body[start:end]

		agreements = append(agreements, domain.CollectionDetailAgreement{
    	AgreementNo:               parseIntField(block[0:16]),        
    	SeqOfAgreement:            parseIntField(block[16:18]),       
    	OutsourceID:               strings.TrimSpace(block[18:22]),  
    	OutsourceName:             strings.TrimSpace(block[22:52]),   
    	BlockCode:                 strings.TrimSpace(block[52:54]),  
    	CurrentSUEOSPrincipalNet:  parseFloatField(block[54:64]),   
    	CurrentSUEOSPrincipalVAT:  parseFloatField(block[64:74]),   
    	CurrentSUEOSInterestNet:   parseFloatField(block[74:84]),   
    	CurrentSUEOSInterestVAT:   parseFloatField(block[84:94]),   
    	CurrentSUEOSPenalty:       parseFloatField(block[94:103]),  
    	CurrentSUEOSHDCharge:      parseFloatField(block[103:112]), 
    	CurrentSUEOSOtherFee:      parseFloatField(block[112:121]), 
    	CurrentSUEOSTotal:         parseFloatField(block[121:131]), 
    	TotalPaymentAmount:        parseFloatField(block[131:141]), 
    	LastPaymentDate:           parseIntField(block[141:149]),   
    	SUESeqNo:                  parseIntField(block[149:151]),   
    	BeginSUEOSPrincipalNet:    parseFloatField(block[151:161]), 
    	BeginSUEOSPrincipalVAT:    parseFloatField(block[161:171]), 
    	BeginSUEOSInterestNet:     parseFloatField(block[171:181]), 
    	BeginSUEOSInterestVAT:     parseFloatField(block[181:191]), 
    	BeginSUEOSPenalty:         parseFloatField(block[191:201]), 
    	BeginSUEOSHDCharge:        parseFloatField(block[201:210]), 
    	BeginSUEOSOtherFee:        parseFloatField(block[210:219]), 
    	BeginSUEOSTotal:           parseFloatField(block[219:229]), 
    	SUEStatus:                 parseIntField(block[229:231]),   
    	SUEStatusDescription:      strings.TrimSpace(block[231:261]), 
    	BlackCaseNo:               strings.TrimSpace(block[261:276]), 
    	BlackCaseDate:             parseIntField(block[276:284]),  
    	RedCaseNo:                 strings.TrimSpace(block[284:299]), 
    	RedCaseDate:               parseIntField(block[299:307]),   
    	CourtCode:                 strings.TrimSpace(block[307:311]), 
    	CourtName:                 strings.TrimSpace(block[311:341]), 
    	JudgmentDate:              parseIntField(block[341:349]),   
    	JudgmentResultCode:        parseIntField(block[349:350]),   
    	JudgmentResultDescription: strings.TrimSpace(block[350:390]), 
    	JudgmentDetail:            strings.TrimSpace(block[390:890]), 
    	ExpectDate:                parseIntField(block[890:898]),   
    	AssetPrice:                parseFloatField(block[898:908]), 
    	JudgeAmount:               parseFloatField(block[908:918]), 
    	NoOfInstallment:           strings.TrimSpace(block[918:921]), 
    	InstallmentAmount:         parseFloatField(block[921:931]), 
    	TotalCurrentPerSUESeqNo:   parseFloatField(block[931:942]), 
		})
	}

	return domain.CollectionDetailResponse{
		IDCardNo:      idCardNo,
		NoOfAgreement: noOfAgreement,
		AgreementList: agreements,
	}
}

// Converts CollectionLogRequest to a fixed-length string.
func FormatCollectionLogRequest(reqData domain.CollectionLogRequest) string {
	var result string

	for _, item := range reqData.CollectionLogAgreementRq {
		AgreementNo := PadOrTruncate(strconv.Itoa(item.AgreementNo), 16)
		RemarkCode := PadOrTruncate(item.RemarkCode, 4)
		LogRemark1 := PadOrTruncate(item.LogRemark1, 120)
		LogRemark2 := PadOrTruncate(item.LogRemark2, 120)
		LogRemark3 := PadOrTruncate(item.LogRemark3, 120)
		LogRemark4 := PadOrTruncate(item.LogRemark4, 120)
		LogRemark5 := PadOrTruncate(item.LogRemark5, 120)
		InputDate := PadOrTruncate(item.InputDate, 12)
		InputTime := PadOrTruncate(item.InputTime, 6)
		OperatorID := PadOrTruncate(item.OperatorID, 15)

		result += AgreementNo + RemarkCode + LogRemark1 + LogRemark2 + LogRemark3 +
			LogRemark4 + LogRemark5 + InputDate + InputTime + OperatorID
	}

	return result
}

func FormatCollectionLogResponse(respData string) domain.CollectionLogResponse {
	const headerLen = 123
	body := respData[headerLen:]

	idCardNo := strings.TrimSpace(body[0:20])

	const rowLen = 18
	content := body[20:]
	numRows := len(content) / rowLen
	agreements := make([]domain.CollectionLogAgreementRs, 0, numRows)

	for i := 0; i < numRows; i++ {
		start := i * rowLen
		end := start + rowLen
		if end > len(content) {
			break
		}
		row := content[start:end]

		agreements = append(agreements, domain.CollectionLogAgreementRs{
			AgreementNo:     parseIntField(row[0:16]),
			LogRemarkStatus: strings.TrimSpace(row[16:18]),
		})
	}

	return domain.CollectionLogResponse{
		IDCardNo:      idCardNo,
		AgreementList: agreements,
	}
}
