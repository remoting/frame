package cpuid

import (
	"errors"
	cpu "github.com/klauspost/cpuid/v2"
	"os"
	"os/exec"
	"runtime"
	"strings"
)

func getCpuInfo() cpu.CPUInfo {
	return cpu.CPU
}
func getLinuxFileContent(path string) (string, error) {
	content, err := os.ReadFile(path)
	if err != nil {
		return "", err
	}
	return strings.TrimSpace(string(content)), nil
}

func getWindowsWMIValue(class, property string) (string, error) {
	out, err := exec.Command("wmic", class, "get", property, "/value").Output()
	if err != nil {
		return "", err
	}
	result := strings.TrimSpace(strings.Split(string(out), "=")[1])
	return result, nil
}

func getUnixCmdOutput(cmd string) (string, error) {
	out, err := exec.Command("sh", "-c", cmd).Output()
	if err != nil {
		return "", err
	}
	return strings.TrimSpace(string(out)), nil
}

func GenerateMachineCode() (string, error) {
	var tmp string
	if runtime.GOOS == "darwin" {
		tmp, _ = getUnixCmdOutput("ioreg -l | grep IOPlatformSerialNumber | awk -F '=' '{print $2}' | tr -d ' ' | tr -d '\"'")
		if tmp == "" {
			tmp, _ = getUnixCmdOutput("diskutil info / | grep 'Volume UUID' | awk '{print $3}'")
		}
		if tmp == "" {
			return "", errors.New("machine code error")
		}
	} else if runtime.GOOS == "linux" {
		tmp, _ = getUnixCmdOutput("lsblk -o UUID | grep -v ^$ | awk 'NR>1 {print $1; exit}'")
		if tmp != "" {
			return getCpuInfo().BrandName+"-"+tmp,nil
		}
		tmp, _ = getLinuxFileContent("/etc/machine-id")
		if tmp == "" {
			tmp, _ = getLinuxFileContent("/var/lib/dbus/machine-id")
		}
		if tmp != "" {
			return getCpuInfo().BrandName+"-"+tmp,nil
		}
		if tmp == "" {
			return "", errors.New("machine code error")
		}
	} else if runtime.GOOS == "windows" {
		serial, _ := getWindowsWMIValue("Win32_BaseBoard", "SerialNumber")
		disk, _ := getWindowsWMIValue("Win32_DiskDrive", "SerialNumber")
		tmp = getCpuInfo().BrandName + serial + disk
		if tmp == "" {
			return "", errors.New("machine code error")
		}
	}
	return tmp, nil
}
