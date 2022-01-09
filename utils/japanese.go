package utils

import "unicode"

const (
	// http://www.unicodemap.org/range/62/Hiragana/
	// ぁ
	hiraganaLo = 0x3041

	// http://www.unicodemap.org/range/63/Katakana/
	// ァ
	katakanaLo = 0x30a1

	codeDiff = katakanaLo - hiraganaLo
)

// HiraganaToKatakana converts Hiragana str to Katakana.
// Refer to https://qiita.com/yoheimuta/items/35c8dfb36cddfe19a64e
func HiraganaToKatakana(str string) string {
	src := []rune(str)
	dst := make([]rune, len(src))
	for i, r := range src {
		switch {
		case unicode.In(r, unicode.Hiragana):
			dst[i] = r + codeDiff
		default:
			dst[i] = r
		}
	}
	return string(dst)
}

// HiraganaToKatakana converts Hiragana str to Katakana.
// Refer to https://qiita.com/yoheimuta/items/35c8dfb36cddfe19a64e
func KatakanaToHiragana(str string) string {
	src := []rune(str)
	dst := make([]rune, len(src))
	for i, r := range src {
		switch {
		case unicode.In(r, unicode.Katakana):
			dst[i] = r - codeDiff
		default:
			dst[i] = r
		}
	}
	return string(dst)
}
