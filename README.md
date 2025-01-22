
# LoadBuster

**LoadBuster** is a CLI-based performance and stress testing tool written in Go. It allows you to quickly configure concurrency, duration, HTTP method, headers, and raw request bodies for load tests against any HTTP endpoint.

## Features

- **Simulates concurrent users** with configurable concurrency.
- **Duration-based tests**, letting you run short or extended load scenarios.
- **Supports all HTTP methods** (GET, POST, PUT, DELETE, etc.).
- **Customizable headers**, including authorization tokens.
- **Raw request bodies**, such as JSON or any text, directly through a command-line flag.
- **Basic metrics collection**, including total requests, successful and failed requests, and min/avg/max latencies.
- **Docker and Docker Compose compatibility** for consistent, containerized testing environments.

## Quick Start

1. **Clone the repository:**
   ```bash
   git clone https://github.com/yourusername/loadbuster.git
   cd loadbuster
   ```

2. **Build from source:**
   ```bash
   go build -o loadbuster ./cmd/loadbuster
   ```

3. **Run a sample load test:**
   ```bash
   ./loadbuster start --url https://api.example.com --method POST --concurrency 10 --duration 15s --auth "Bearer xyz123" --body '{"id":"3"}'
   ```
   This example sends a JSON body to `https://api.example.com` using a POST request, with 10 concurrent workers for 15 seconds, including a bearer token in the authorization header.

## Docker Usage

1. **Build the Docker image:**
   ```bash
   docker build -t loadbuster .
   ```

2. **Run LoadBuster in a container:**
   ```bash
   docker run --rm loadbuster start --url https://api.example.com --method POST --concurrency 10 --duration 10s --auth "Bearer xyz123" --body '{"id":"3"}'
   ```

## Docker Compose

1. **Customize the `docker-compose.yml` file** to specify your preferred URL, method, concurrency, duration, and request body.

2. **Run LoadBuster:**
   ```bash
   docker-compose up --build
   ```

## Contributing

Contributions, bug reports, and feature requests are welcome! Please open an issue or submit a pull request following the project's guidelines for style, testing, and commit practices.

## License

**LoadBuster** is licensed under the MIT License. Use, modify, and distribute freely while respecting the terms detailed in the LICENSE file.
