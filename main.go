package main

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"strings"
)

var messages = map[string]map[string]string{
	"tr": {
		"usage":    "❌ Hata: Lütfen bir port numarası girin. Kullanım: killport <port>",
		"notFound": "⚠️ %s portunu kullanan herhangi bir işlem bulunamadı.\n",
		"errKill":  "❌ %s portu (PID: %s) kapatılırken hata oluştu: %v\n",
		"success":  "✅ %s portu (PID: %s) başarıyla kapatıldı!\n",
		"failAll":  "⚠️ %s portunda çalışan işlemler kapatılamadı.\n",
		"help":     "🔫 KillPort - Pratik Port Kapatıcı\n\nKullanım:\n  killport <port_numarası>    Belirtilen portu kullanan işlemleri kapatır.\n  killport -h, --help         Bu yardım mesajını gösterir.\n\nÖrnek:\n  killport 8080\n",
	},
	"en": {
		"usage":    "❌ Error: Please provide a port number. Usage: killport <port>",
		"notFound": "⚠️ No process found running on port %s.\n",
		"errKill":  "❌ Error killing process on port %s (PID: %s): %v\n",
		"success":  "✅ Successfully killed process on port %s (PID: %s)!\n",
		"failAll":  "⚠️ Could not kill processes running on port %s.\n",
		"help":     "🔫 KillPort - Practical Port Killer\n\nUsage:\n  killport <port_number>      Kills the processes running on the specified port.\n  killport -h, --help         Shows this help message.\n\nExample:\n  killport 8080\n",
	},
}

func getLang() string {
	lang := os.Getenv("LANG")
	if strings.HasPrefix(lang, "tr") {
		return "tr"
	}
	// Geri kalan her şey için varsayılan dil İngilizce
	return "en"
}

func main() {
	lang := getLang()
	msg := messages[lang]

	// 1. Kullanıcıdan port numarasını al
	if len(os.Args) < 2 {
		fmt.Println(msg["usage"])
		os.Exit(1)
	}
	port := os.Args[1]
	
	// Yardım parametresi kontrolü
	if port == "-h" || port == "--help" {
		fmt.Print(msg["help"])
		os.Exit(0)
	}

	var pids []string

	// 2. İşletim sistemine o portu kimin kullandığını sor
	if runtime.GOOS == "windows" {
		out, _ := exec.Command("cmd", "/c", fmt.Sprintf("netstat -ano | findstr :%s", port)).Output()
		lines := strings.Split(string(out), "\n")
		for _, line := range lines {
			fields := strings.Fields(line)
			if len(fields) >= 4 {
				pid := fields[len(fields)-1]
				if pid != "0" {
					// Duplicate kontrolü
					exists := false
					for _, p := range pids {
						if p == pid {
							exists = true
							break
						}
					}
					if !exists {
						pids = append(pids, pid)
					}
				}
			}
		}
	} else { // macOS / Linux
		out, err := exec.Command("lsof", "-t", "-i", ":"+port).Output()
		if err == nil && len(out) > 0 {
			pids = strings.Split(strings.TrimSpace(string(out)), "\n")
		}
	}

	if len(pids) == 0 {
		fmt.Printf(msg["notFound"], port)
		os.Exit(0)
	}

	// 3. Bulunan işlemleri öldür
	successCount := 0
	for _, pid := range pids {
		if pid == "" {
			continue
		}

		var killCmd *exec.Cmd
		if runtime.GOOS == "windows" {
			killCmd = exec.Command("taskkill", "/F", "/PID", pid)
		} else {
			killCmd = exec.Command("kill", "-9", pid)
		}

		err := killCmd.Run()
		if err != nil {
			fmt.Printf(msg["errKill"], port, pid, err)
		} else {
			fmt.Printf(msg["success"], port, pid)
			successCount++
		}
	}
	
	if successCount == 0 {
		fmt.Printf(msg["failAll"], port)
	}
}
