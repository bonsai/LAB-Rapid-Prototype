[CmdletBinding()]
param(
  [string]$ExtensionDir = (Join-Path $PSScriptRoot '..\ext'),
  [string]$Name = 'Rapid Prototype Actions'
)

$ErrorActionPreference = 'Stop'
$ExtensionDir = (Resolve-Path $ExtensionDir).Path
$manifest = Join-Path $ExtensionDir 'manifest.json'

if (-not (Test-Path $manifest)) {
  throw "manifest.json not found: $manifest"
}

Write-Host "Installing Chrome extension: $Name"
Write-Host "Extension directory: $ExtensionDir"

$chrome = @(
  "$env:ProgramFiles\Google\Chrome\Application\chrome.exe",
  "${env:ProgramFiles(x86)}\Google\Chrome\Application\chrome.exe",
  "$env:LOCALAPPDATA\Google\Chrome\Application\chrome.exe",
  "$env:ProgramFiles\Microsoft\Edge\Application\msedge.exe",
  "${env:ProgramFiles(x86)}\Microsoft\Edge\Application\msedge.exe",
  "$env:LOCALAPPDATA\Microsoft\Edge\Application\msedge.exe"
) | Where-Object { $_ -and (Test-Path $_) } | Select-Object -First 1

if ($chrome) {
  Write-Host "Browser: $chrome"
}

Write-Host ''
Write-Host 'Load unpacked extension:'
Write-Host '  1. Open chrome://extensions (or edge://extensions)'
Write-Host '  2. Enable Developer mode'
Write-Host "  3. Select 'Load unpacked' and choose: $ExtensionDir"
Write-Host ''
Write-Host 'CRX installation is handled by the Build CRX GitHub Action.'
Write-Host 'This script installs the development/unpacked extension.'

if ($chrome) {
  Start-Process $chrome 'chrome://extensions/'
}
