package format

import (
    "strconv"
    "strings"
    "testing"

    //"connectorapi-go/internal/core/domain"
)


func padOrTruncatersTest(s string, length int) string {
    if len(s) > length {
        return s[:length]
    }
    return s + strings.Repeat(" ", length-len(s))
}


func buildMockResponse(idCardNo string, numRows int) string {
    header := strings.Repeat("H", 123)
    body := padOrTruncatersTest(idCardNo, 20)

    for i := 1; i <= numRows; i++ {
        agreementNoStr := padOrTruncatersTest(strconv.Itoa(i), 16)
        logRemarkStatus := padOrTruncatersTest(strconv.Itoa(i+10), 2) // แค่ตัวอย่าง
        body += agreementNoStr + logRemarkStatus
    }

    return header + body
}

func TestFormatCollectionLogResponse_ControlRows(t *testing.T) {
    
    numRows := 5
    idCardNo := "12345678901234567890"

    respData := buildMockResponse(idCardNo, numRows)
    got := FormatCollectionLogResponse(respData)

    
    t.Logf("Formatted Response: %+v", got)

    
    if got.IDCardNo != idCardNo {
        t.Errorf("IDCardNo mismatch. Got %q", got.IDCardNo)
    }

    
    if len(got.AgreementList) != numRows {
        t.Errorf("AgreementList length mismatch. Got %d, Want %d", len(got.AgreementList), numRows)
    }

    
    for i, row := range got.AgreementList {
        t.Logf("Row %d: AgreementNo=%d, LogRemarkStatus=%q", i+1, row.AgreementNo, row.LogRemarkStatus)
    }
}
