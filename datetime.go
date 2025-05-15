package utils

import "time"

// 获取一个时间的零点时间
func DayZero(moment time.Time) time.Time {
	zero := time.Date(moment.Year(), moment.Month(), moment.Day(), 0, 0, 0, 0, time.Local)
	return zero
}

// 计算两次零点时间差
func SubZero(lastTime time.Time) float64 {
	lastZero := DayZero(lastTime)
	todayZero := DayZero(time.Now())
	subHours := todayZero.Sub(lastZero).Hours()
	return subHours
}

// 获取当前时间到下一个零点的duration
func ToNextZero() time.Duration {
	now := time.Now()
	endOfDay := time.Date(now.Year(), now.Month(), now.Day(), 23, 59, 59, 0, time.Local)
	// 计算距离下个零点的时间
	duration := endOfDay.Sub(now)
	return duration
}
