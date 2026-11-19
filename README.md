# Telescope Proxy

Reverse proxy HTTP ringan berbasis **Golang (Clean Architecture)** dengan dukungan *Dynamic Sticky Failover* SOCKS5 dan `curl-impersonate`.

## Tech Stack

- **Language**: [Go](https://go.dev/)
- **HTTP Framework**: [GoFiber v3](https://github.com/gofiber/fiber)
- **Configuration**: [Viper](https://github.com/spf13/viper)
- **Logger**: [Logrus](https://github.com/sirupsen/logrus) (Structured JSON Logging)
- **Testing**: [Testify](https://github.com/stretchr/testify)
- **HTTP Egress Engine**: `curl` (local development) / `curl-impersonate-chrome` (production)

## Requirements

Pastikan perangkat lokal sudah terinstal:
- **Go** (minimal versi `1.25+`)
- **cURL** (`curl` standar untuk development lokal, atau `curl_chrome116` untuk production)
- **Git**

## Installation & Setup

1. **Clone repository**:
   ```bash
   git clone <repository-url>
   cd telescope-proxy
   ```

2. **Salin file konfigurasi `.env`**:
   ```bash
   cp .env.example .env
   ```
   Sesuaikan nilai variabel di dalam `.env` (`UPSTREAM_HOST`, `SOCKS_TARGETS`, `PORT`, dll.) dengan kebutuhan environment lokal.

3. **Unduh dan rapikan dependencies**:
   ```bash
   go mod tidy
   ```

## Development Commands

- **Menjalankan server di lokal**:
  ```bash
  go run cmd/web/main.go
  ```

- **Build binary aplikasi**:
  ```bash
  go build -o bin/telescope-proxy cmd/web/main.go
  ```

- **Pemeriksaan statis (Static Analysis / Vet)**:
  ```bash
  go vet ./...
  ```

## Testing Commands

- **Menjalankan seluruh unit test**:
  ```bash
  go test -v ./...
  ```

- **Menjalankan unit test beserta persentase coverage**:
  ```bash
  go test -v -cover ./...
  ```

- **Membuat laporan coverage (`coverage.out`) untuk SonarQube / HTML**:
  ```bash
  go test -coverprofile=coverage.out ./...
  go tool cover -html=coverage.out
  ```

- **Menjalankan test dengan Race Condition Detector**:
  ```bash
  go test -race ./...
  ```

- **Menjalankan Benchmark Performa & Memori**:
  ```bash
  go test -bench=. -benchmem ./...
  ```