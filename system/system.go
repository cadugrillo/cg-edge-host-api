package system

import (
	"fmt"
	"syscall"
	"time"

	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/disk"
	"github.com/shirou/gopsutil/mem"
)

type HostStats struct {
	CpuUsage      []float64
	RamTotal      float64
	RamUsed       float64
	RamUsedPct    float64
	RamAvailable  float64
	RamFree       float64
	DiskUsage     float64
	DiskAvailable float64
	DiskTotal     float64
}

var (
	hostStats HostStats
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

func GetHostStats() HostStats {

	cpuPercent, err := cpu.Percent(time.Second, true)
	if err != nil {
		fmt.Println(err.Error())
		return hostStats
	}
	hostStats.CpuUsage = append(hostStats.CpuUsage, cpuPercent...)

	m, err := mem.VirtualMemory()
	if err != nil {
		fmt.Println(err.Error())
		return hostStats
	}
	hostStats.RamTotal = getReadableSize(m.Total)
	hostStats.RamUsed = getReadableSize(m.Used)
	hostStats.RamUsedPct = m.UsedPercent
	hostStats.RamAvailable = getReadableSize(m.Available)
	hostStats.RamFree = getReadableSize(m.Free)

	fmt.Println("Ram Total: ", getReadableSize(m.Total))
	fmt.Println("Ram Used: ", getReadableSize(m.Used))
	fmt.Println("Ram Used %: ", m.UsedPercent)
	fmt.Println("Ram Available: ", getReadableSize(m.Available))
	fmt.Println("Ram Free: ", getReadableSize(m.Free))

	diskUsage, err := disk.Usage("/")
	if err != nil {
		fmt.Println(err.Error())
		return hostStats
	}

	hostStats.DiskUsage = diskUsage.UsedPercent
	hostStats.DiskAvailable = getReadableSize(diskUsage.Free)
	hostStats.DiskTotal = getReadableSize(diskUsage.Total)

	return hostStats
}

func getReadableSize(sizeInBytes uint64) (readableSize float64) {
	var (
		units = []string{"B", "KB", "MB", "GB", "TB", "PB"}
		size  = float64(sizeInBytes)
		i     = 0
	)
	for ; i < len(units) && size >= 1000; i++ {
		size = size / 1000
	}

	return size
}
