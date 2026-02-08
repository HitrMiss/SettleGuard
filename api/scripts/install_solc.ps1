# scripts/install_solc.ps1
$version = "v0.8.33"

$scriptDir = Split-Path -Parent $MyInvocation.MyCommand.Definition

$apiDir = (Get-Item (Join-Path $scriptDir "..")).FullName

$binDir = Join-Path $apiDir "bin"

if (!(Test-Path $binDir)) {
    New-Item -ItemType Directory -Path $binDir | Out-Null
    Write-Host "Created directory: $binDir" -ForegroundColor Gray
}

Write-Host "Downloading Solidity $version standalone binary..." -ForegroundColor Cyan

$url = "https://github.com/ethereum/solidity/releases/download/$version/solc-windows.exe"
$destFile = Join-Path $binDir "solc.exe"

try {
    Invoke-WebRequest -Uri $url -OutFile $destFile -ErrorAction Stop
    Write-Host "Success! solc installed locally at: $destFile" -ForegroundColor Green
    Write-Host "You can now run: (my) (generation) gen.go see README.md" -ForegroundColor Yellow
}
catch {
    Write-Host "Download failed: $($_.Exception.Message)" -ForegroundColor Red
}