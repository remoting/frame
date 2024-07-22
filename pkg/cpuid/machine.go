package cpuid

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	cpuinfo "github.com/klauspost/cpuid/v2"
	"os"
	"os/exec"
	"runtime"
	"strings"
)

func getProcessorID() (string, error) {
	//cat /proc/cpuinfo | grep name | awk -F ':' '{print $2}' | sed 's/^ //;s/ $//'
	return cpuinfo.CPU.BrandName, nil
}

func getBaseBoardSerialNumber() (string, error) {
	switch runtime.GOOS {
	case "windows":
		return getWindowsWMIValue("Win32_BaseBoard", "SerialNumber")
		//return getBaseBoardSerialNumber()
	case "linux":
		// Check for multiple possible sources
		if serial, err := getLinuxFileContent("/sys/class/dmi/id/board_serial"); err == nil && serial != "" {
			return serial, nil
		}
		if machineID, err := getLinuxFileContent("/etc/machine-id"); err == nil && machineID != "" {
			return machineID, nil
		}
		if machineID, err := getLinuxFileContent("/var/lib/dbus/machine-id"); err == nil && machineID != "" {
			return machineID, nil
		}
		return "", fmt.Errorf("baseboard serial number not available")
	case "darwin":
		return getUnixCmdOutput("ioreg -l | grep IOPlatformSerialNumber | awk -F '=' '{print $2}' | tr -d ' ' | tr -d '\"'")
	default:
		return "", fmt.Errorf("unsupported platform")
	}
}
func getDiskDriveSerialNumber() (string, error) {
	switch runtime.GOOS {
	case "windows":
		return getWindowsWMIValue("Win32_DiskDrive", "SerialNumber")
	case "linux":
		return getUnixCmdOutput("lsblk -o UUID | grep -v ^$ | awk 'NR>1 {print $1; exit}'")
	case "darwin":
		return getUnixCmdOutput("diskutil info / | grep 'Volume UUID' | awk '{print $3}'")
	default:
		return "", fmt.Errorf("unsupported platform")
	}
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
	processorID, err := getProcessorID()
	if err != nil {
		return "", err
	}

	baseBoardSerial, err := getBaseBoardSerialNumber()
	if err != nil {
		return "", err
	}
	diskSerial, err := getDiskDriveSerialNumber()
	if err != nil {
		return "", err
	}

	// Concatenate all parts
	combined := processorID + baseBoardSerial + diskSerial

	// Hash the combined string to generate a machine code
	hash := md5.New()
	hash.Write([]byte(combined))
	machineCode := hex.EncodeToString(hash.Sum(nil))

	return machineCode, nil
}
