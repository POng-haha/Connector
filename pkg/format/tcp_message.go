package format

import (
	"fmt"
	"time"

	"connectorapi-go/internal/core/domain"
)

// PadOrTruncate pads a string with spaces on the right or truncates it to a given length.
func PadOrTruncate(s string, length int) string {
	if len(s) > length {
		return s[:length]
	}
	return fmt.Sprintf("%-*s", length, s)
}

// FormatGetCustomerInfoRequestFixedLength converts GetCustomerInfoRequest to a fixed-length string.
func FormatGetCustomerInfoRequestFixedLength(reqData domain.GetCustomerInfoRequest) string {
	IDCardNo := PadOrTruncate(reqData.IDCardNo, 20)
	AeonID := PadOrTruncate(reqData.AeonID, 20)
	AgreementNo := PadOrTruncate(reqData.AgreementNo, 16)
	return IDCardNo + AeonID + AgreementNo
}

// FormatUpdateAddressRequestFixedLength converts UpdateAddressRequest to a fixed-length string.
func FormatUpdateAddressRequestFixedLength(reqData domain.UpdateAddressRequest) string {
	customerID := PadOrTruncate(reqData.CustomerID, 36)
	city := PadOrTruncate(reqData.Address.City, 100)
	street := PadOrTruncate(reqData.Address.Street, 100)
	postalCode := PadOrTruncate(reqData.Address.PostalCode, 5)
	return customerID + city + street + postalCode
}

// FormatGetCustomerInfo004RequestFixedLength converts GetCustomerInfo004Request to a fixed-length string.
func FormatGetCustomerInfo004RequestFixedLength(reqData domain.GetCustomerInfo004Request) string {
	IDCardNo := PadOrTruncate(reqData.IDCardNo, 20)
	AeonID := PadOrTruncate(reqData.AeonID, 20)
	AgreementNo := PadOrTruncate(reqData.AgreementNo, 16)
	return IDCardNo + AeonID + AgreementNo
}

// FormatCollectionLogRequestFixedLength converts CollectionLogRequest to a fixed-length string.
func FormatCollectionLogRequestFixedLength(reqData domain.CollectionLogRequest) string {
	CustomerID := PadOrTruncate(reqData.CustomerID, 20)
	Address := PadOrTruncate(reqData.Address, 20)
	return CustomerID + Address
}

// FormatCollectionDetailRequestFixedLength converts CollectionDetailRequest to a fixed-length string.
func FormatCollectionDetailRequestFixedLength(reqData domain.CollectionDetailRequest) string {
	CustomerID := PadOrTruncate(reqData.CustomerID, 20)
	Address := PadOrTruncate(reqData.Address, 20)
	return CustomerID + Address
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
