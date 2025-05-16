package date

import "time"

func FormatTime(timeStruct time.Time) string {
	return timeStruct.Format("2006-01-02 15:04:05")
}
