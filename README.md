# music-basicifier

Music Basicifier is a Windows desktop app (Go + Fyne) that aims to:
- take music from a YouTube URL or local audio file,
- extract a melody,
- and generate outputs for QBASIC and MakeCode Adafruit.

Current implemented scope (roadmap progress):
- YouTube URL input field in the UI
- Local audio file input field in the UI (`.wav`, `.m4u`, `.mp3`)

## Platform Target

This project currently targets:
- Windows 10/11
- amd64

The GUI stack uses Fyne and requires cgo-enabled builds on Windows.

## Environment Requirements

## 1) Go

Required Go version (from `go.mod`):
- Go 1.26.5

Verify:

```powershell
go version
```

## 2) C Toolchain for cgo (required)

Because Fyne's desktop rendering path needs cgo on Windows, you must have a C compiler on `PATH`.

A working option is LLVM MinGW (UCRT):

```powershell
winget install --id MartinStorsjo.LLVM-MinGW.UCRT --accept-package-agreements --accept-source-agreements
```

After install, restart your shell/terminal, then verify:

```powershell
gcc --version
```

If `gcc` is not found, close and reopen VS Code, then try again.

## 3) golangci-lint (required for checks)

The project check script runs `golangci-lint`.
Install it if needed:

```powershell
go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest
```

Ensure your Go bin directory is on `PATH` (typically `%USERPROFILE%\go\bin`).
Verify:

```powershell
golangci-lint version
```

## 4) make (optional)

`make` is optional on Windows.
If you do not have `make`, use `check.ps1` instead.

## Clone and Install Dependencies

```powershell
git clone https://github.com/RobPro/music-basicifier.git
cd music-basicifier
go mod download
```

## Build

### Windows-native build (recommended)

```powershell
$env:CGO_ENABLED = "1"
go build -o bin\app.exe .\cmd\app
```

### Via Makefile (if `make` is available)

```powershell
make build
```

The Makefile build target already sets:
- `CGO_ENABLED=1`
- `GOOS=windows`
- `GOARCH=amd64`

## Run

From source:

```powershell
go run .\cmd\app
```

Or run built executable:

```powershell
.\bin\app.exe
```

## Quality Checks

### Preferred on Windows

```powershell
.\check.ps1
```

This runs:
- `gofmt` check
- `go vet`
- `golangci-lint run`
- `go test ./...`
- `go build -o bin\app.exe .\cmd\app` (with `CGO_ENABLED=1`)

### Alternative

```powershell
make check
```

## Troubleshooting

## Error: "ui requires cgo-enabled build on Windows"

Cause:
- The app was built with `CGO_ENABLED=0` or without a working C compiler.

Fix:
1. Install a GCC-compatible toolchain (example above).
2. Confirm `gcc --version` works in your terminal.
3. Build with cgo enabled:

```powershell
$env:CGO_ENABLED = "1"
go build -o bin\app.exe .\cmd\app
```

## Error: `cgo: C compiler "gcc" not found`

Cause:
- No C compiler on `PATH`.

Fix:
1. Install LLVM MinGW via `winget`.
2. Restart terminal/VS Code so `PATH` updates are loaded.
3. Re-run `gcc --version` and then re-run `./check.ps1`.

## Project Layout (high level)

- `cmd/app`: application entrypoint
- `internal/ui`: Fyne window and UI wiring
- `internal/domain`: domain-layer package placeholder
- `internal/platform`: platform-layer package placeholder
- `example-input`: sample input audio
- `example-output`: sample expected output format

## Notes

- No admin rights are required to run the app itself.
- Roadmap and scope details are documented in `PROJECT.md`.
