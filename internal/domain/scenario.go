package domain

import "time"

// Scenario holds the configuration for a load test.
type Scenario struct {
    URL         string
    Method      string
    Concurrency int
    Duration    time.Duration
    AuthHeader  string
    Body        string
    ContentType string
}
