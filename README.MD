## PingPong Web Service

This repository contains the implementation of a simple web service called "PingPong". It responds with "pong" when it receives a GET request on the "/ping" path.

### Project Structure

The project is organized as follows:

```
cmd/
│   └── pingpong/
│       └── main.go                # Entry point of the application
internal/
│   ├── handlers/
│   │   ├── ping.go                # Handler for the /ping endpoint
│   │   ├── ping_test.go           # Tests for the ping handler
│   │   ├── health.go              # Handler for the /health endpoint
│   │   └── health_test.go         # Tests for the health handler
│   ├── middleware/
│   │   ├── logging.go             # Logging middleware
│   │   └── logging_test.go        # Tests for the logging middleware
│   └── response/
│       ├── http_response.go       # Helper functions for HTTP responses
│       └── http_response_test.go  # Tests for the response helpers
```

### How to Run

1. **Build the application**:

   ```bash
   make build
   ```

2. **Run the application**:

   ```bash
   make run 
   ```

3. **Test the application**:

   In the root directory of the project, run:

   ```bash
   make test
   ```

### How to Check Test Coverage

   **Generate a coverage profile**:

   ```bash
   make coverage
   ```
