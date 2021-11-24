package timer

import "time"

// 获取时间
func GetNowTime() time.Time {
	// 设置当前时区
	location, _ := time.LoadLocation("Asia/Shanghai")
	return time.Now().In(location)
}

// 推算时间
func GetCalcalateTime(currentTimer time.Time, d string) (time.Time, error) {
	duration, err := time.ParseDuration(d)
	if err != nil {
		return time.Time{}, err
	}

	return currentTimer.Add(duration), nil
}

/*
// 如果预先知道准确的duration，且不需要适配，那么即可直接使用Add方法进行代理
const (
	Nanosecond  time.Duration = 1
	Microsecond               = 1000 * Nanosecond
	Millisecond               = 1000 * Microsecond
	Second                    = 1000 * Millisecond
	Minute                    = 60 * Second
	Hour                      = 60 * Minute
)

func GetCalcalateTime()  {
	GetNowTime().Add(time.Second * 60)
}
*/