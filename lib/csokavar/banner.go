package csokavar

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/encse/altnet/lib/io"
	"github.com/hako/durafmt"
	"github.com/shirou/gopsutil/v3/cpu"
	"github.com/shirou/gopsutil/v3/host"
	"github.com/shirou/gopsutil/v3/load"
)

func Banner(screenWidth int) string {
	var out strings.Builder

	arch := "-"
	platform := "-"
	cpus := "-"
	load1 := "-"
	load5 := "-"
	load15 := "-"
	uptime := "-"

	infoStat, err := host.Info()
	if err == nil {
		arch = infoStat.KernelArch
		platform = infoStat.OS
		uptime = fmt.Sprintf("%v", durafmt.Parse(time.Duration(infoStat.Uptime)*time.Second))
	}

	cpuCount, err := cpu.Counts(false)
	if err == nil {
		cpus = strconv.Itoa(cpuCount)
	}

	loadAvg, err := load.Avg()
	if err == nil {
		load1 = fmt.Sprintf("%.2f", loadAvg.Load1)
		load5 = fmt.Sprintf("%.2f", loadAvg.Load5)
		load15 = fmt.Sprintf("%.2f", loadAvg.Load15)
	}

	out.WriteString(io.Center("Connected to CSOKAVAR, Encse's home on the web. Happy surfing.", screenWidth))
	out.WriteString("\n")
	out.WriteString("\n")
	out.WriteString(
		io.Center(
			fmt.Sprintf("Server: %v %v with %v cpu(s), load average: %v, %v, %v", arch, platform, cpus, load1, load5, load15),
			screenWidth,
		))
	out.WriteString("\n")
	out.WriteString(io.Center(fmt.Sprintf("uptime: %v", uptime), screenWidth))
	out.WriteString("\n")
	out.WriteString("\n")
	out.WriteString(io.Center("SysOp: encse", screenWidth))
	out.WriteString("\n")
	out.WriteString("\n")

	return out.String()
}
