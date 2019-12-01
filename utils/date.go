package utils

import (
	"strconv"
	"time"

	"github.com/guregu/null"
)

var jst = time.FixedZone("Asia/Tokyo", 9*60*60)

const (
	dateFormat               = "2006-01-02"
	dateTimeFormat           = "2006-01-02 15:04:05"
	dateFormatOnlyNumber     = "20060102"
	dateTimeFormatOnlyNumber = "20060102150405"
)

// FormatDateJST は JST で日付をフォーマットする.
func FormatDateJST(date time.Time) string {
	return date.In(jst).Format(dateFormat)
}

// FormatDateTimeJST は JST で日付をフォーマットする.
func FormatDateTimeJST(date time.Time) string {
	return date.In(jst).Format(dateTimeFormat)
}

// FormatDateTimeOnlyNumber は 日付をフォーマットする.
func FormatDateTimeOnlyNumber(date time.Time) string {
	return date.Format(dateTimeFormatOnlyNumber)
}

// FormatDateJSTByNullTime は null.Time を JST で日付フォーマットした日付を返す. null の場合は空文字を返す.
func FormatDateJSTByNullTime(date null.Time) string {
	if date.Valid {
		return FormatDateJST(date.Time)
	}
	return ""
}

// FormatDateTimeJSTByNullTime は null.Time を JST で日付フォーマットした日付(時間を含む)を返す. null の場合は空文字を返す.
func FormatDateTimeJSTByNullTime(date null.Time) string {
	if date.Valid {
		return FormatDateTimeJST(date.Time)
	}
	return ""
}

// CalcDateSub 日時の差分をそれぞれ返却する.
func CalcDateSub(from, to time.Time) (days, hours, mins, secs int) {
	sub := to.Sub(from)
	days = int(sub.Hours()) / 24
	hours = int(sub.Hours()) % 24
	mins = int(sub.Minutes()) % 60
	secs = int(sub.Seconds()) % 60
	return
}

// NumberOfTheWeekInMonth は月中の第何週かを返す.
func NumberOfTheWeekInMonth(now time.Time) int {
	beginningOfTheMonth := time.Date(now.Year(), now.Month(), 1, 1, 1, 1, 1, now.Location())
	_, thisWeek := now.ISOWeek()
	_, beginningWeek := beginningOfTheMonth.ISOWeek()
	return 1 + thisWeek - beginningWeek
}

// BeginningOfMonth 今月の月初めを返す.
func BeginningOfMonth(t time.Time) time.Time {
	beginningOfTheMonth := time.Date(t.Year(), t.Month(), 1, 0, 0, 0, 0, t.Location())
	return beginningOfTheMonth
}

// BeginningOfLastMonth 先月の月初めを返す..
func BeginningOfLastMonth(t time.Time) time.Time {
	t = t.AddDate(0, -1, 0)
	return BeginningOfMonth(t)
}

// BeginningOfNextMonth 来月の月初めを返す..
func BeginningOfNextMonth(t time.Time) time.Time {
	t = t.AddDate(0, 1, 0)
	return BeginningOfMonth(t)
}

// BeginningOfWeek 指定日時の週初めを返す.
func BeginningOfWeek(t time.Time) time.Time {
	t = BeginningOfDay(t)
	return t.AddDate(0, 0, -int(t.Weekday()))
}

// BeginningOfLastWeek 指定日時の前週の初めを返す.
func BeginningOfLastWeek(t time.Time) time.Time {
	t = BeginningOfDay(t).AddDate(0, 0, 7)
	return t.AddDate(0, 0, -int(t.Weekday()))
}

// BeginningOfDay 指定日の開始を返す.
func BeginningOfDay(t time.Time) time.Time {
	return time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, t.Location())
}

// BeginningOfLastDay 前日の開始を返す.
func BeginningOfLastDay(t time.Time) time.Time {
	return time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, t.Location())
}

// JSTNow 現在時間を返す.
func JSTNow() time.Time {
	return time.Now().In(JSTLocation())
}

// JSTLocation 現在時間を返す.
func JSTLocation() *time.Location {
	return time.FixedZone("Asia/Tokyo", 9*60*60)
}

// GetAge 誕生日から現在の年齢を返す.
func GetAge(t null.Time) null.Int {
	// 誕生日未入力の場合はnil
	if !t.Valid {
		return null.IntFromPtr(nil)
	}

	// ( 今日の日付(YYYYMMDD) - 誕生日(YYYYMMD) ) / 10000 で年齢を求める
	now := JSTNow().Format(dateFormatOnlyNumber)
	birthday := t.Time.In(jst).Format(dateFormatOnlyNumber)

	nowInt, err := strconv.Atoi(now)
	if err != nil {
		// 数値化に失敗したらnilを返しておく
		return null.IntFromPtr(nil)
	}
	birthdayInt, err := strconv.Atoi(birthday)
	if err != nil {
		return null.IntFromPtr(nil)
	}

	age := (nowInt - birthdayInt) / 10000
	return null.IntFrom(int64(age))
}

// DateUTC 年月日からtime.Time型を生成する
func DateUTC(year, month, date int) time.Time {
	return time.Date(year, time.Month(month), date, 0, 0, 0, 0, time.UTC)
}
