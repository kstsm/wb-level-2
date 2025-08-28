package timeutil

import (
	"fmt"
	"github.com/beevik/ntp"
)

// GetCurrentTime возвращает текущее время с NTP-сервера.
func GetCurrentTime() (string, error) {
	// Это нужно было бы поместить в конфиг, но для учебного проекта я оставил так
	const (
		ntpServer  = "pool.ntp.org"
		formatTime = "02.01.2006 15:04:05"
	)

	currentTime, err := ntp.Time(ntpServer)
	if err != nil {
		return "", fmt.Errorf("GetCurrentTime: %w", err)
	}

	return currentTime.Format(formatTime), nil
}
