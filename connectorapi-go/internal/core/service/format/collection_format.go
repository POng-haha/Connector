package format

import (
	"fmt"
	"strconv"
	"strings"
	// "time"
	//"bytes"

	"connectorapi-go/internal/core/domain"
	"connectorapi-go/internal/adapter/utils"
)

// Converts CollectionDetailRequest to a fixed-length string.
func FormatCollectionDetailRequest(reqData domain.CollectionDetailRequest) string {
	IDCardNo    := utils.PadOrTruncate(reqData.IDCardNo, 20)
	RedCaseNo   := utils.PadOrTruncate(reqData.RedCaseNo, 15)
	BlackCaseNo := utils.PadOrTruncate(reqData.BlackCaseNo, 15)
	return IDCardNo + RedCaseNo + BlackCaseNo
}

func FormatCollectionDetailResponse(raw string) (domain.CollectionDetailResponse, error) {
	const headerLen = 123
	if len(raw) <= headerLen {
		return domain.CollectionDetailResponse{}, fmt.Errorf("raw data too short for header, length=%d", len(raw))
	}
	body := raw[headerLen:]
	runes := []rune(body)

	readRunes := func(start, length int) []rune {
		if start >= len(runes) {
			return []rune{}
		}
		end := start + length
		if end > len(runes) {
			end = len(runes)
		}
		return runes[start:end]
	}

	readString := func(start, length int) string {
		rs := readRunes(start, length)
		return strings.TrimSpace(string(rs))
	}

	readInt := func(start, length int) int {
		s := readString(start, length)
		if s == "" {
			return 0
		}
		i, _ := strconv.Atoi(s)
		return i
	}

	// readFloat := func(start, length int) float64 {
	// 	s := readString(start, length)
	// 	if s == "" {
	// 		return 0
	// 	}
	// 	f, _ := strconv.ParseFloat(s, 64)
	// 	return f
	// }

	idCardNo := readString(0, 20)
	noOfAgreement := readInt(20, 2)

	const agreementLen = 942
	agreementStart := 22
	agreements := make([]domain.CollectionDetailAgreement, 0, noOfAgreement)

	for i := 0; i < noOfAgreement; i++ {
		start := agreementStart + i*agreementLen
		if start >= len(runes) {
			break
		}
		blockRunes := readRunes(start, agreementLen)

		readBlockStr := func(startField, length int) string {
			if startField >= len(blockRunes) {
				return ""
			}
			end := startField + length
			if end > len(blockRunes) {
				end = len(blockRunes)
			}
			return strings.TrimSpace(string(blockRunes[startField:end]))
		}

		readBlockInt := func(startField, length int) int {
			s := readBlockStr(startField, length)
			if s == "" {
				return 0
			}
			i, _ := strconv.Atoi(s)
			return i
		}

        // readBlockFloat := func(startField, length int) domain.DecimalString {
        //     s := readBlockStr(startField, length)
        //     if s == "" {
        //     return 0.00
        // }
        // i, _ := strconv.ParseInt(s, 10, 64)      // แปลง string เป็น int
        // f := float64(i) / 100.0                  // หาร 100 เพื่อทศนิยม 2 ตำแหน่ง
        // return domain.DecimalString(f)
        // }

		readBlockFloat := func(startField, length int) utils.DecimalString {
            s := readBlockStr(startField, length)
            i, _ := strconv.ParseInt(s, 10, 64)
            f := float64(i) / 100.0
            return utils.DecimalString(f)
       }


		agreements = append(agreements, domain.CollectionDetailAgreement{
			AgreementNo:              readBlockStr(0, 16),
			SeqOfAgreement:           readBlockInt(16, 2),
			OutsourceID:              readBlockStr(18, 4),
			OutsourceName:            readBlockStr(22, 30),
			BlockCode:                readBlockStr(52, 2),
			CurrentSUEOSPrincipalNet: readBlockFloat(54, 10),
			CurrentSUEOSPrincipalVAT: readBlockFloat(64, 10),
			CurrentSUEOSInterestNet:  readBlockFloat(74, 10),
			CurrentSUEOSInterestVAT:  readBlockFloat(84, 10),
			CurrentSUEOSPenalty:      readBlockFloat(94, 9),
			CurrentSUEOSHDCharge:     readBlockFloat(103, 9),
			CurrentSUEOSOtherFee:     readBlockFloat(112, 9),
			CurrentSUEOSTotal:        readBlockFloat(121, 10),
			TotalPaymentAmount:       readBlockFloat(131, 10),
			LastPaymentDate:          readBlockInt(141, 8),
			SUESeqNo:                 readBlockInt(149, 2),
			BeginSUEOSPrincipalNet:   readBlockFloat(151, 10),
			BeginSUEOSPrincipalVAT:   readBlockFloat(161, 10),
			BeginSUEOSInterestNet:    readBlockFloat(171, 10),
			BeginSUEOSInterestVAT:    readBlockFloat(181, 10),
			BeginSUEOSPenalty:        readBlockFloat(191, 10),
			BeginSUEOSHDCharge:       readBlockFloat(201, 9),
			BeginSUEOSOtherFee:       readBlockFloat(210, 9),
			BeginSUEOSTotal:          readBlockFloat(219, 10),
			SUEStatus:                readBlockInt(229, 2),
			SUEStatusDescription:     readBlockStr(231, 30),
			BlackCaseNo:              readBlockStr(261, 15),
			BlackCaseDate:            readBlockInt(276, 8),
			RedCaseNo:                readBlockStr(284, 15),
			RedCaseDate:              readBlockInt(299, 8),
			CourtCode:                readBlockStr(307, 4),
			CourtName:                readBlockStr(311, 30),
			JudgmentDate:             readBlockInt(341, 8),
			JudgmentResultCode:       readBlockInt(349, 1),
			JudgmentResultDescription: readBlockStr(350, 40),
			JudgmentDetail:           readBlockStr(390, 500),
			ExpectDate:               readBlockInt(890, 8),
			AssetPrice:               readBlockFloat(898, 10),
			JudgeAmount:              readBlockFloat(908, 10),
			NoOfInstallment:          readBlockStr(918, 3),
			InstallmentAmount:        readBlockFloat(921, 10),
			TotalCurrentPerSUESeqNo:  readBlockFloat(931, 11),
		})
	}

	return domain.CollectionDetailResponse{
		IDCardNo:      idCardNo,
		NoOfAgreement: noOfAgreement,
		AgreementList: agreements,
	}, nil
}


// Converts CollectionLogRequest to a fixed-length string.
func FormatCollectionLogRequest(reqData domain.CollectionLogRequest) string {
	AgreementNo := utils.PadOrTruncate(reqData.AgreementNo, 16)
	RemarkCode  := utils.PadOrTruncate(reqData.RemarkCode, 4)
	LogRemark1  := utils.PadOrTruncate(reqData.LogRemark1, 120)
	LogRemark2  := utils.PadOrTruncate(reqData.LogRemark2, 120)
	LogRemark3  := utils.PadOrTruncate(reqData.LogRemark3, 120)
	LogRemark4  := utils.PadOrTruncate(reqData.LogRemark4, 120)
	LogRemark5  := utils.PadOrTruncate(reqData.LogRemark5, 120)
	InputDate   := utils.PadOrTruncate(reqData.InputDate, 8)
	InputTime   := utils.PadOrTruncate(reqData.InputTime, 6)
	OperatorID  := utils.PadOrTruncate(reqData.OperatorID, 15)
	return AgreementNo + RemarkCode + LogRemark1 + LogRemark2 + LogRemark3 + LogRemark4 + LogRemark5 + InputDate + InputTime + OperatorID
}

func FormatCollectionLogResponse(raw string) (domain.CollectionLogResponse, error) {
	const headerLen = 123
	const dataLen = 36 

	if len(raw) <= headerLen {
		return domain.CollectionLogResponse{}, fmt.Errorf("raw data too short for header, length=%d", len(raw))
	}

	data := raw[headerLen:]
	if len(data) < dataLen {
		return domain.CollectionLogResponse{}, fmt.Errorf("raw data too short for body, length=%d, need %d", len(data), dataLen)
	}

	idCardNo := strings.TrimSpace(data[:20])
	agreementNo := strings.TrimSpace(data[20:36])

	return domain.CollectionLogResponse{
		IDCardNo:    idCardNo,
		AgreementNo: agreementNo,
	}, nil
}