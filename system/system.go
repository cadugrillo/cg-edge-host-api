package system

import (
	"fmt"
	"syscall"
	"time"

	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/disk"
	"github.com/shirou/gopsutil/mem"
)

func RestartHost() string {
	syscall.Sync()
	err := syscall.Reboot(syscall.LINUX_REBOOT_CMD_RESTART)
	return err.Error()
}

func ShutdownHost() string {
	syscall.Sync()
	err := syscall.Reboot(syscall.LINUX_REBOOT_CMD_POWER_OFF)
	return err.Error()
}

func GetHostStats() string {
	cpuPercent, err := cpu.Percent(time.Second, true)
	if err != nil {
		panic(err)
	}
	usedPercent0 := fmt.Sprintf("%.2f", cpuPercent[0])
	usedPercent1 := fmt.Sprintf("%.2f", cpuPercent[1])
	//usedPercent2 := fmt.Sprintf("%.2f", cpuPercent[2])
	//usedPercent3 := fmt.Sprintf("%.2f", cpuPercent[3])
	fmt.Println("CPU Usage: ", usedPercent0+"%")
	fmt.Println("CPU Usage: ", usedPercent1+"%")
	//fmt.Println("CPU Usage: ", usedPercent2+"%")
	//fmt.Println("CPU Usage: ", usedPercent3+"%")

	m, err := mem.VirtualMemory()
	if err != nil {
		panic(err)
	}

	usedMemory := fmt.Sprintf(
		"%s (%.2f%%)",
		getReadableSize(m.Used),
		m.UsedPercent,
	)
	fmt.Println("Ram Total: ", getReadableSize(m.Total))
	fmt.Println("Ram Used: ", usedMemory)
	fmt.Println("Ram Available: ", getReadableSize(m.Available))
	fmt.Println("Ram Free: ", getReadableSize(m.Free))

	diskUsage, err := disk.Usage("/")
	if err != nil {
		panic(err)
	}
	usedPercent := fmt.Sprintf("%.2f", diskUsage.UsedPercent)
	fmt.Println("Disk Usage: ", usedPercent+"%")
	fmt.Println("Disk Space Available: ", getReadableSize(diskUsage.Free))

	return "done"
}

func getReadableSize(sizeInBytes uint64) (readableSizeString string) {
	var (
		units = []string{"B", "KB", "MB", "GB", "TB", "PB"}
		size  = float64(sizeInBytes)
		i     = 0
	)
	for ; i < len(units) && size >= 1024; i++ {
		size = size / 1024
	}
	readableSizeString = fmt.Sprintf("%.2f %s", size, units[i])
	return
}
