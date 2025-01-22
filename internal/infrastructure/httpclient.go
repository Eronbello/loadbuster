package infrastructure

import (
	"bytes"
	"net/http"
	"sync"
	"time"

	"github.com/eronbello/loadbuster/internal/domain"
)

// RunHTTPTest executes the load test using the given Scenario.
func RunHTTPTest(scenario domain.Scenario) (result struct {
    TotalRequests int
    Successful    int
    Failed        int
    MinLatency    time.Duration
    MaxLatency    time.Duration
    AvgLatency    time.Duration
}) {
    if scenario.Concurrency < 1 {
        scenario.Concurrency = 1
    }

    // If no duration is specified, default to zero (single iteration).
    testDuration := scenario.Duration
    if testDuration <= 0 {
        testDuration = 0
    }

    var (
        wg          sync.WaitGroup
        totalMu     sync.Mutex
        successMu   sync.Mutex
        failMu      sync.Mutex
        latencies   []time.Duration
        latenciesMu sync.Mutex
        stopChan    = make(chan struct{})
    )

    worker := func() {
        defer wg.Done()

        for {
            select {
            case <-stopChan:
                return
            default:
                startTime := time.Now()

                req, err := http.NewRequest(scenario.Method, scenario.URL, bytes.NewBufferString(scenario.Body))
                if err != nil {
                    failMu.Lock()
                    result.Failed++
                    failMu.Unlock()
                    continue
                }

                // Headers
                if scenario.AuthHeader != "" {
                    req.Header.Set("Authorization", scenario.AuthHeader)
                }
                if scenario.Body != "" {
                    req.Header.Set("Content-Type", scenario.ContentType)
                }

                client := &http.Client{}
                resp, err := client.Do(req)
                if err != nil {
                    failMu.Lock()
                    result.Failed++
                    failMu.Unlock()
                } else {
                    if resp.StatusCode >= 200 && resp.StatusCode < 300 {
                        successMu.Lock()
                        result.Successful++
                        successMu.Unlock()
                    } else {
                        failMu.Lock()
                        result.Failed++
                        failMu.Unlock()
                    }
                    _ = resp.Body.Close()
                }

                latency := time.Since(startTime)

                // Update counters
                totalMu.Lock()
                result.TotalRequests++
                totalMu.Unlock()

                // Track latency
                latenciesMu.Lock()
                latencies = append(latencies, latency)
                latenciesMu.Unlock()
            }
        }
    }

    // Start the workers
    for i := 0; i < scenario.Concurrency; i++ {
        wg.Add(1)
        go worker()
    }

    // If a duration is specified, run for that duration; otherwise do a single pass
    if testDuration > 0 {
        time.Sleep(testDuration)
        close(stopChan)
    } else {
        // No duration: close immediately after starting => each worker does one iteration
        close(stopChan)
    }

    wg.Wait()

    // Calculate min, max, and avg latencies
    if len(latencies) > 0 {
        var total time.Duration
        min := latencies[0]
        max := latencies[0]

        for _, l := range latencies {
            if l < min {
                min = l
            }
            if l > max {
                max = l
            }
            total += l
        }
        result.MinLatency = min
        result.MaxLatency = max
        result.AvgLatency = total / time.Duration(len(latencies))
    }

    return
}
