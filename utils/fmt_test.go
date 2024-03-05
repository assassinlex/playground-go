package utils

import (
	"fmt"
	"testing"
	"time"
)

func TestTime(t *testing.T) {
	now := time.Now()
	fmt.Println(now.Day())
	year, month, day := now.Date()
	fmt.Println(year, int(month), day)
	fmt.Println(fmt.Sprintf("%s 00:00:00", now.Format("2006-01-02")))
	fmt.Println(fmt.Sprintf("%s 00:00:00", now.Add(24*time.Hour).Format("2006-01-02")))
}
