package format

import (
	"fmt"
	"strconv"
	"strings"
	"time"
	//"bytes"

	"connectorapi-go/internal/core/domain"
	"connectorapi-go/internal/adapter/utils"
)

// Converts MyCardRequest to a fixed-length string.
func FormatMyCardRequestNormal(myCardReq domain.MyCardRequest) string {
	IDCardNo        		:= utils.PadOrTruncate(myCardReq.UserRef, 20)
	CreditCardNo            := "                "
	BusinessCode            := "  "
	return IDCardNo + CreditCardNo + BusinessCode
}

func FormatMyCardRequestAll(myCardReq domain.MyCardRequest) string {
	IDCardNo        		:= utils.PadOrTruncate(myCardReq.UserRef, 20)
	CustomerNameEN          := "Y"
	CustomerNameTH          := "Y"
	return IDCardNo + CustomerNameEN + CustomerNameTH
}

func FormatMyCardResponseNormal(raw string) (domain.MyCardResponseNormal, error) {
	const headerLen = 123
	currentDate := time.Now().Format("20060102")
	currentDateInt, _ := strconv.Atoi(currentDate)

	if len(raw) <= headerLen {
		return domain.MyCardResponseNormal{}, fmt.Errorf("raw data too short for header, length=%d", len(raw))
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

	idCardNo                  := readString(0, 20)
	totalCreditCard           := readInt(20, 4)

	const agreementLen = 61
	agreementStart := 24
	agreements := make([]domain.MyCardListNormal, 0, totalCreditCard)

	for i := 0; i < totalCreditCard; i++ {
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

		rawCreditCardNo := readBlockStr(0, 16)
		cardStatus := readBlockStr(50, 2)
		expireDateStr := readBlockStr(52, 8)
		digitalCardFlag := readBlockStr(60, 1)

		maskedCreditCardNo := rawCreditCardNo
		if len(rawCreditCardNo) == 16 {
			maskedCreditCardNo = rawCreditCardNo[0:6] + "XXXXXX" + rawCreditCardNo[12:16]
		}

		status := "HLD"
		if cardStatus == "00" || cardStatus == "II" {
			status = "ACT"
		} else {
			expireDateInt, _ := strconv.Atoi(expireDateStr)
			if expireDateInt < currentDateInt {
				status = "EXP"
			}
		}

		finalDigitalCardFlag := digitalCardFlag
		if digitalCardFlag == "" {
			finalDigitalCardFlag = "N"
		}


		agreements = append(agreements, domain.MyCardListNormal{
			CreditCardNo:        maskedCreditCardNo,
			CardName:            readBlockStr(16, 30),
			ProductType:         readBlockStr(46, 2),
			BusinessCode:        readBlockStr(48, 2),
			CardStatus:          status,
			ExpireDate:          readBlockStr(52, 8),
			// DYCA:                readBlockStr(60, 1),
			DigitalCardFlag:     finalDigitalCardFlag,
		})
	}

	return domain.MyCardResponseNormal{
		IDCardNo:                 idCardNo,
		TotalCreditCard:          totalCreditCard,
		CardList:                 agreements,
	}, nil
}

func FormatMyCardResponseAll(raw string) (domain.MyCardResponseAll, error) {
	const headerLen = 123
	if len(raw) <= headerLen {
		return domain.MyCardResponseAll{}, fmt.Errorf("raw data too short for header, length=%d", len(raw))
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

	idCardNo               := readString(0, 20)
	customerNameEN         := readString(20, 30)
	customerNameTH         := readString(50, 30)
	totalCreditCard     := readInt(80, 3)

	const agreementLen = 68
	agreementStart := 83
	agreements := make([]domain.MyCardListAll, 0, totalCreditCard)

	for i := 0; i < totalCreditCard; i++ {
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

	// Raw values
	rawCreditCardNo := readBlockStr(0, 16)

	maskedCreditCardNo := rawCreditCardNo
	if len(rawCreditCardNo) == 16 {
		maskedCreditCardNo = rawCreditCardNo[0:6] + "XXXXXX" + rawCreditCardNo[12:16]
	}

	agreements = append(agreements, domain.MyCardListAll{
		CreditCardNo:        maskedCreditCardNo,
		CardCode:            readBlockStr(16, 2),
		ProductType:         readBlockStr(18, 2),
		CardType:            readBlockStr(20, 1),
		CardStatus:          readBlockStr(21, 1),
		ExpireDate:          readBlockInt(22, 8),
		HoldCode:            readBlockStr(30, 2),
		RetreatCode:         readBlockStr(32, 1),
		SendMode:            readBlockStr(33, 1),
		FirstEmbossDate:     readBlockInt(34, 8),
		FirstConfirmDate:    readBlockInt(42, 8),
		ShoppingLimit:       readBlockInt(50, 9),
		CashingLimit:        readBlockInt(59, 9),
	})
}

	return domain.MyCardResponseAll{
		IDCardNo:                 idCardNo,
		CustomerNameEN:           customerNameEN,
		CustomerNameTH:           customerNameTH,
		TotalCreditCard:          totalCreditCard,
		CardList:                 agreements,
	}, nil
}