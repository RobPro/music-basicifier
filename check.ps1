gofmt -l . | Out-String -OutVariable fmtIssues
if ($fmtIssues.Trim().Length -gt 0) { Write-Error "gofmt issues found"; exit 1 }
go vet ./...
if ($LASTEXITCODE -ne 0) { exit 1 }
golangci-lint run
if ($LASTEXITCODE -ne 0) { exit 1 }
go test ./...
if ($LASTEXITCODE -ne 0) { exit 1 }
go build -o bin\app.exe .\cmd\app