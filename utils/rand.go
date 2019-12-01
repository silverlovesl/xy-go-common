package utils

import (
	"math/rand"
	"time"
)

var randSrc = rand.NewSource(time.Now().UnixNano())

const (
	//rsLetters       = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	// 0, O など紛らわしい文字は避ける.
	rsLetters       = "123456789ABCDEFGHIJKLMNPQRSTUVWXYZ"
	rsLetterIdxBits = 6
	rsLetterIdxMask = 1<<rsLetterIdxBits - 1
	rsLetterIdxMax  = 63 / rsLetterIdxBits

	shortenLetters       = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
	shortenLetterIdxBits = 7
	shortenLetterIdxMask = 1<<rsLetterIdxBits - 1
	shortenLetterIdxMax  = 127 / rsLetterIdxBits
)

// RandStringFriendly は引き継ぎコードや問い合わせコードなど視認しやすいランダムコードを得る.
func RandStringFriendly(n int) string {
	b := make([]byte, n)
	cache, remain := randSrc.Int63(), rsLetterIdxMax
	for i := n - 1; i >= 0; {
		if remain == 0 {
			cache, remain = randSrc.Int63(), rsLetterIdxMax
		}
		idx := int(cache & rsLetterIdxMask)
		if idx < len(rsLetters) {
			b[i] = rsLetters[idx]
			i--
		}
		cache >>= rsLetterIdxBits
		remain--
	}
	return string(b)
}

// RandStringForShortenURL は短縮URL用のランダム文字を返す.
func RandStringForShortenURL(n int) string {
	b := make([]byte, n)
	cache, remain := randSrc.Int63(), shortenLetterIdxMax
	for i := n - 1; i >= 0; {
		if remain == 0 {
			cache, remain = randSrc.Int63(), shortenLetterIdxMax
		}
		idx := int(cache & shortenLetterIdxMask)
		if idx < len(shortenLetters) {
			b[i] = shortenLetters[idx]
			i--
		}
		cache >>= shortenLetterIdxBits
		remain--
	}
	return string(b)
}
