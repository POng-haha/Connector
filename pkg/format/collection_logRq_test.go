package format

import (
    "strconv"
    "strings"
    "testing"

    "connectorapi-go/internal/core/domain"
)

func padOrTruncateTest(s string, length int) string {
    if len(s) > length {
        return s[:length]
    }
    return s + strings.Repeat(" ", length-len(s))
}

func TestFormatCollectionLogRequest(t *testing.T) {
    req := domain.CollectionLogRequest{
        CollectionLogAgreementRq: []domain.CollectionLogAgreementRq{
            {
                AgreementNo: 123,
                RemarkCode:  "AB",
                LogRemark1:  "Remark1",
                LogRemark2:  "",
                LogRemark3:  "Remark3",
                LogRemark4:  "",
                LogRemark5:  "",
                InputDate:   "20250820",
                InputTime:   "123456",
                OperatorID:  "OP01",
            },
        },
    }

    expected := padOrTruncateTest(strconv.Itoa(123), 16) +
        padOrTruncateTest("AB", 4) +
        padOrTruncateTest("Remark1", 120) +
        padOrTruncateTest("", 120) +
        padOrTruncateTest("Remark3", 120) +
        padOrTruncateTest("", 120) +
        padOrTruncateTest("", 120) +
        padOrTruncateTest("20250820", 12) +
        padOrTruncateTest("123456", 6) +
        padOrTruncateTest("OP01", 15)

    got := FormatCollectionLogRequest(req)

	t.Log("Formatted string:", got)

    if got != expected {
        t.Errorf("Formatted string does not match expected.\nGot:  %q\nWant: %q", got, expected)
    }
}
