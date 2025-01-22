package application

import (
	"time"

	"github.com/eronbello/loadbuster/internal/domain"
	"github.com/eronbello/loadbuster/internal/infrastructure"
)

// Result captures the outcome of a load test run.
type Result struct {
    TotalRequests int
    Successful    int
    Failed        int
    MinLatency    time.Duration
    MaxLatency    time.Duration
    AvgLatency    time.Duration
}

// BuildScenario creates a Scenario from user inputs.
func BuildScenario(url, method string, concurrency int, duration time.Duration, authHeader, body, contentType string) domain.Scenario {
    return domain.Scenario{
        URL:         url,
        Method:      method,
        Concurrency: concurrency,
        Duration:    duration,
        AuthHeader:  authHeader,
        Body:        body,
        ContentType: contentType,
    }
}

// RunLoadTest orchestrates the load test by launching concurrent workers.
func RunLoadTest(scenario domain.Scenario) Result {
    return infrastructure.RunHTTPTest(scenario)
}
