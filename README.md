<div align="center">

# 🔫 killport

### The fastest way to kill processes hogging your ports.

**English** · [Türkçe](#-türkçe)

<br/>

[![Go](https://img.shields.io/badge/Go-1.21+-00ADD8?style=for-the-badge&logo=go&logoColor=white&labelColor=1a1b26)](https://go.dev)
[![License](https://img.shields.io/badge/License-MIT-50FA7B?style=for-the-badge&labelColor=1a1b26)](LICENSE)
[![Platforms](https://img.shields.io/badge/macOS%20·%20Linux%20·%20Windows-FFA500?style=for-the-badge&labelColor=1a1b26)](#-installation)

<br/>
</div>

<br/>

## ✨ Why killport?

Ever got a "port already in use" error? Finding the PID and killing it manually is annoying. **killport** does it for you in a single command, supporting multiple languages out of the box.

<table>
<tr>
<td width="50%" valign="top">

### ⚡ Blazing fast
Written in Go, executes instantly without any heavy runtime or dependencies.

### 🌍 Speaks your language
Automatically detects your system language (`LANG`) and replies in English or Turkish.

</td>
<td width="50%" valign="top">

### 🎯 Dead simple
No complex flags. Just type `killport 8080` and the port is free.

### 📦 Cross-Platform
Native support for macOS, Linux, and Windows. Single compiled binary for your OS.

</td>
</tr>
</table>

## 🚀 Quick start

```bash
killport 8080       # Kills the process running on port 8080
killport -h         # Shows the interactive help menu
```

## 📦 Installation

<details open>
<summary><b>🪟 Windows (PowerShell)</b></summary>

```powershell
Invoke-WebRequest -Uri "https://github.com/Emiran404/killport/releases/latest/download/killport-windows-amd64.exe" -OutFile "killport.exe"
mkdir "$env:USERPROFILE\bin" -Force
move killport.exe "$env:USERPROFILE\bin\"
setx PATH "$env:PATH;$env:USERPROFILE\bin"
# open a new terminal, then:
killport -h
```
</details>

<details>
<summary><b>🍎 macOS</b></summary>

```bash
# Apple Silicon (M1/M2/M3): arm64  ·  Intel: amd64
curl -L -o killport https://github.com/Emiran404/killport/releases/latest/download/killport-darwin-arm64
chmod +x killport
sudo mv killport /usr/local/bin/
```
</details>

<details>
<summary><b>🐧 Linux</b></summary>

```bash
# amd64 for most machines  ·  ARM: arm64
curl -L -o killport https://github.com/Emiran404/killport/releases/latest/download/killport-linux-amd64
chmod +x killport
sudo mv killport /usr/local/bin/
```
</details>

<details>
<summary><b>🌍 Any platform — with Go</b></summary>

```bash
go install github.com/Emiran404/killport@latest
```
> If `killport: command not found` appears, make sure your Go bin directory (e.g. `~/go/bin`) is in your `PATH`.
</details>

<br/>

---

<br/>

<div align="center">

# 🇹🇷 Türkçe

### Portlarını işgal eden işlemleri kapatmanın en hızlı yolu.

</div>

## ✨ Neden killport?

"Port already in use" (Port kullanımda) hatası hepimizin canını sıkar. PID bulup manuel kapatmak yerine **killport** ile bunu tek satırda ve kendi dilinizde yapın.

- ⚡ **İnanılmaz hızlı** — Go ile yazıldığı için milisaniyeler içinde tepki verir.
- 🌍 **Sizin dilinizi konuşur** — Sistem dilini otomatik algılar (Türkçe/İngilizce destekli).
- 🎯 **Çok basit** — Karmaşık parametreler yok. Sadece `killport 3000` yazın.
- 📦 **Çapraz Platform** — macOS, Linux ve Windows'ta yerleşik çalışır. Bağımlılık gerektirmez.

## 🚀 Hızlı başlangıç

```bash
killport 3000       # 3000 portunu kullanan işlemi kapatır
killport -h         # Yardım menüsünü gösterir
```

## 📦 Kurulum

<details open>
<summary><b>🪟 Windows (PowerShell)</b></summary>

```powershell
Invoke-WebRequest -Uri "https://github.com/Emiran404/killport/releases/latest/download/killport-windows-amd64.exe" -OutFile "killport.exe"
mkdir "$env:USERPROFILE\bin" -Force
move killport.exe "$env:USERPROFILE\bin\"
setx PATH "$env:PATH;$env:USERPROFILE\bin"
# yeni bir terminal aç, sonra:
killport -h
```
</details>

<details>
<summary><b>🍎 macOS</b></summary>

```bash
# Apple Silicon (M1/M2/M3): arm64  ·  Intel: amd64
curl -L -o killport https://github.com/Emiran404/killport/releases/latest/download/killport-darwin-arm64
chmod +x killport
sudo mv killport /usr/local/bin/
```
</details>

<details>
<summary><b>🐧 Linux</b></summary>

```bash
# Çoğu makine için: amd64  ·  ARM: arm64
curl -L -o killport https://github.com/Emiran404/killport/releases/latest/download/killport-linux-amd64
chmod +x killport
sudo mv killport /usr/local/bin/
```
</details>

<details>
<summary><b>🌍 Tüm Platformlar — Go ile</b></summary>

```bash
go install github.com/Emiran404/killport@latest
```
</details>

<br/>

<div align="center">

**⭐ Faydalı bulduysan bir yıldız bırak — If you find killport useful, give it a star!**

</div>
