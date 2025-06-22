# APIStorm

Lightweight, powerful API performance testing tool that measures what matters.

## Overview

APIStorm is a command-line tool for API performance testing. It allows you to send multiple concurrent requests to your API endpoints and measure response times.

## Installation

```bash
go install github.com/yourusername/apistorm/cmd/cli@latest
```

Or build from source:

```bash
git clone https://github.com/yourusername/apistorm.git
cd apistorm
go build -o apistorm ./cmd/cli
```

## Usage

### Basic Examples

Test a GET endpoint with 100 requests:

```bash
apistorm --method GET --url https://api.example.com/users --totalcalls 100
```

POST request with a JSON body:

```bash
apistorm --method POST --url https://api.example.com/users --body '{"name": "John Doe"}' --totalcalls 50 --delay 10
```

### Available Options

| Flag         | Description                            | Default | Required |
|--------------|----------------------------------------|---------|----------|
| --method     | HTTP method (GET, POST, PATCH, DELETE) | -       | Yes      |
| --url        | API endpoint URL                       | -       | Yes      |
| --body       | Request body (for POST/PATCH)          | "{}"    | No       |
| --totalcalls | Number of requests to make             | -       | Yes      |
| --delay      | Delay between requests (milliseconds)  | 0       | No       |

## Features

- Concurrent API requests
- Support for GET, POST, PATCH, and DELETE methods
- Custom request bodies
- Configurable delay between requests
- Performance metrics reporting

## License

[MIT](LICENSE)