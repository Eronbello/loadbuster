package main

import (
	"fmt"
	"os"
	"time"

	"github.com/eronbello/loadbuster/internal/application"
	"github.com/spf13/cobra"
)

var (
    url         string
    method      string
    concurrency int
    duration    time.Duration
    authHeader  string
    body        string
    contentType string
)

// rootCmd represents the base command
var rootCmd = &cobra.Command{
    Use:   "loadbuster",
    Short: "LoadBuster is a CLI-based performance and stress testing tool in Go",
    Long: `LoadBuster is an open-source tool written in Go
that allows you to performance-test your endpoints with various configurations.`,
}

// startCmd triggers the load test
var startCmd = &cobra.Command{
    Use:   "start",
    Short: "Start a load test",
    RunE: func(cmd *cobra.Command, args []string) error {
        if url == "" {
            return fmt.Errorf("url is required, use --url or -u")
        }

        // Build the scenario
        scenario := application.BuildScenario(
            url, method, concurrency, duration,
            authHeader, body, contentType,
        )

        // Run the load test
        result := application.RunLoadTest(scenario)

        // Print results
        fmt.Printf("\n==== LoadBuster Results ====\n")
        fmt.Printf("Total Requests: %d\n", result.TotalRequests)
        fmt.Printf("Successful:     %d\n", result.Successful)
        fmt.Printf("Failed:         %d\n", result.Failed)
        fmt.Printf("Min Latency:    %v\n", result.MinLatency)
        fmt.Printf("Max Latency:    %v\n", result.MaxLatency)
        fmt.Printf("Avg Latency:    %v\n", result.AvgLatency)
        fmt.Printf("===========================\n")

        return nil
    },
}

func init() {
    // Define flags for "start" command
    startCmd.Flags().StringVarP(&url, "url", "u", "", "Target URL to test (required)")
    startCmd.Flags().StringVarP(&method, "method", "X", "GET", "HTTP method (GET, POST, PUT, etc.)")
    startCmd.Flags().IntVarP(&concurrency, "concurrency", "c", 1, "Number of concurrent workers")
    startCmd.Flags().DurationVarP(&duration, "duration", "d", 0, "Duration of the test (e.g., 30s, 1m, etc.)")
    startCmd.Flags().StringVarP(&authHeader, "auth", "a", "", "Authorization header (e.g., 'Bearer <token>')")
    startCmd.Flags().StringVarP(&body, "body", "b", "", "Raw request body to send (e.g., '{\"id\":\"3\"}')")
    startCmd.Flags().StringVar(&contentType, "content-type", "application/json", "Content-Type header for the request body")

    rootCmd.AddCommand(startCmd)
}

func main() {
    if err := rootCmd.Execute(); err != nil {
        fmt.Println(err)
        os.Exit(1)
    }
}
