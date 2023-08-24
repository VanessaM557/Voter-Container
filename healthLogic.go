package main
import (
	"time"
)

type HealthData struct {
	BootTime              time.Time `json:"bootTime"`
	TotalAPICalls         int       `json:"totalApiCalls"`
	TotalAPICallsWithError int      `json:"totalApiErrors"`
	Uptime                string    `json:"uptime"`
}

func InitializeHealth() HealthData {
	return HealthData{
		BootTime:              time.Now(),
		TotalAPICalls:         0,
		TotalAPICallsWithError: 0,
	}
}

func IncrementTotalAPICalls(healthData *HealthData) {
	healthData.TotalAPICalls++
}

func IncrementTotalAPICallsWithError(healthData *HealthData) {
	healthData.TotalAPICallsWithError++
}
