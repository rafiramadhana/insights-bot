package utils

import (
	"unicode"
)

func ContainsCJKChar(s string) bool {
	for _, r := range s {
		if unicode.Is(unicode.Han, r) {
			return true
		}
		if unicode.Is(unicode.Hangul, r) {
			return true
		}
		if unicode.Is(unicode.Hiragana, r) {
			return true
		}
		if unicode.Is(unicode.Katakana, r) {
			return true
		}

		/*
			U+3001  、
			U+3002  。
			U+3003  〃
			U+3008  〈
			U+3009  〉
			U+300A  《
			U+300B  》
			U+300C  「
			U+300D  」
			U+300E  『
			U+300F  』
			U+3010  【
			U+3011  】
			U+3014  〔
			U+3015  〕
			U+3016  〖
			U+3017  〗
			U+3018  〘
			U+3019  〙
			U+301A  〚
			U+301B  〛
			U+301C  〜
			U+301D  〝
			U+301E  〞
			U+301F  〟
			U+3030  〰
			U+303D  〽
		*/
		if r >= 0x3001 && r <= 0x303D {
			return true
		}
	}

	return false
}
