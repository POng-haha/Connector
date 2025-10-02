package format

import (
	"fmt"
	// "strconv"
	"strings"
	//"bytes"

	"connectorapi-go/internal/core/domain"
	"connectorapi-go/internal/adapter/utils"
)

// Converts UpdateConsentRequest to a fixed-length string.
func FormatUpdateConsentRequest(req domain.UpdateConsentRequest) string {
	var builder strings.Builder

	builder.WriteString(utils.PadOrTruncate(req.IDCardNo, 20))
	builder.WriteString(utils.PadOrTruncate(req.ActionChannel, 3))
	builder.WriteString(utils.PadOrTruncate(req.ActionDateTime, 14))
	builder.WriteString(utils.PadOrTruncate(req.ApplicationNo, 20))
	builder.WriteString(utils.PadOrTruncate(req.ApplicationVersion, 13))
	builder.WriteString(utils.PadOrTruncate(req.IPAddress, 50))
	builder.WriteString(utils.PadOrTruncate(req.ATMNo, 5))
	builder.WriteString(utils.PadOrTruncate(req.BranchCode, 4))
	builder.WriteString(utils.PadOrTruncate(req.VoicePath, 150))
	builder.WriteString(utils.PadIntWithZero(req.TotalOfConsentCode, 2))

	for _, item := range req.ConsentLists {
		builder.WriteString(utils.PadOrTruncate(item.ConsentForm, 3))
		builder.WriteString(utils.PadOrTruncate(item.ConsentCode, 3))
		builder.WriteString(utils.PadOrTruncate(item.ConsentFormVersion, 13))
		builder.WriteString(utils.PadOrTruncate(item.ConsentLanguage, 1))
		builder.WriteString(utils.PadOrTruncate(item.ConsentStatus, 2))
	}

	return builder.String()
}

func FormatUpdateConsentResponse(raw string) (domain.UpdateConsentResponse, error) {
	const headerLen = 123
	const dataLen = 122

	if len(raw) <= headerLen {
		return domain.UpdateConsentResponse{}, fmt.Errorf("raw data too short for header, length=%d", len(raw))
	}

	data := raw[headerLen:]
	if len(data) < dataLen {
		return domain.UpdateConsentResponse{}, fmt.Errorf("raw data too short for body, length=%d, need %d", len(data), dataLen)
	}

	parser := utils.NewFixedParser(data)

	// iDCardNo                 := parser.ReadString(0,20)
	// applicationNo            := parser.ReadString(20,20)
	status                   := parser.ReadString(40,2)
	// filler                   := parser.ReadString(42,80)

	var Finalstatus string
	switch status {
		case "00":
			Finalstatus = "C"
		default:
			Finalstatus = "N"
	}

	return domain.UpdateConsentResponse{
		Status:                    Finalstatus,
	}, nil
}