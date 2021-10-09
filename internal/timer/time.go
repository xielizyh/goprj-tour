package timer

import "time"

// GetNowTime 获取当前本地时间
func GetNowTime() time.Time {
	location, _ := time.LoadLocation("Asia/Shanghai")
	return time.Now().In(location)
}

// GetCalculateTime 获取持续d后的时间点
func GetCalculateTime(currentTimer time.Time, d string) (time.Time, error) {
	duration, err := time.ParseDuration(d)
	if err != nil {
		return time.Time{}, err
	}
	// 当前时间+duration为最终时间
	return currentTimer.Add(duration), nil
}
