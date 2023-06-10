package time

import (
	"time"

	ptime "github.com/yaa110/go-persian-calendar"
)

func ToPersian(t time.Time) string {
	return ptime.New(t).Format("yyyy/MM/dd")
}
