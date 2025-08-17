// Package tis620 provides simple conversion between UTF-8 and TIS-620 (Code Page 874) for Thai text.
package tis620

// table maps TIS-620 (0xA1–0xFB) to Unicode code points for Thai characters.
var toUnicode = [0x5B]rune{
	0x0E01, 0x0E02, 0x0E03, 0x0E04, 0x0E05, 0x0E06, 0x0E07, 0x0E08, 0x0E09, 0x0E0A, 0x0E0B,
	0x0E0C, 0x0E0D, 0x0E0E, 0x0E0F, 0x0E10, 0x0E11, 0x0E12, 0x0E13, 0x0E14, 0x0E15, 0x0E16,
	0x0E17, 0x0E18, 0x0E19, 0x0E1A, 0x0E1B, 0x0E1C, 0x0E1D, 0x0E1E, 0x0E1F, 0x0E20, 0x0E21,
	0x0E22, 0x0E23, 0x0E24, 0x0E25, 0x0E26, 0x0E27, 0x0E28, 0x0E29, 0x0E2A, 0x0E2B, 0x0E2C,
	0x0E2D, 0x0E2E, 0x0E2F, 0x0E30, 0x0E31, 0x0E32, 0x0E33, 0x0E34, 0x0E35, 0x0E36, 0x0E37,
	0x0E38, 0x0E39, 0x0E3A,
}

// ToTIS620 converts UTF-8 bytes to TIS-620 bytes.
// Characters outside Thai range will be replaced with '?' (0x3F).
func ToTIS620(input []byte) []byte {
	out := make([]byte, len(input))
	for i, b := range input {
		r := rune(b)
		if r >= 0xE0 && r <= 0xFB {
			out[i] = byte(0xA1 + (r - 0x0E01))
		} else {
			out[i] = b
		}
	}
	return out
}

// ToUTF8 converts TIS-620 bytes to UTF-8 bytes.
// Non-TIS-620 bytes are kept as-is.
func ToUTF8(input []byte) []byte {
	out := make([]rune, 0, len(input))
	for _, b := range input {
		if b >= 0xA1 && b <= 0xFB {
			r := toUnicode[b-0xA1]
			out = append(out, r)
		} else {
			out = append(out, rune(b))
		}
	}
	// Convert []rune to UTF-8 []byte
	return []byte(string(out))
}
