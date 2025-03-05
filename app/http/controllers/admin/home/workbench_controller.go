package home

import (
	"fmt"
	"strconv"
	"time"

	"goravel/app/http/controllers/admin"

	"github.com/gookit/goutil/fmtutil"
	"github.com/goravel/framework/contracts/http"
	"github.com/goravel/framework/facades"
	"github.com/goravel/framework/support"
	"github.com/goravel/framework/support/carbon"
	"github.com/shirou/gopsutil/v3/cpu"
	"github.com/shirou/gopsutil/v3/host"
	"github.com/shirou/gopsutil/v3/mem"
)

type WorkbenchController struct {
	response admin.Response
}

func NewWorkbenchController() *WorkbenchController {
	return &WorkbenchController{}
}

type Monitor struct {
	Version Version `json:"version"`
	Today   Today   `json:"today"`
	Memory  Memory  `json:"memory"`
	CPU     CPU     `json:"cpu"`
}

type Version struct {
	Framework string `json:"framework"`
	Admin     string `json:"admin"`
	OsArch    string `json:"osArch"`
}

type Today struct {
	Time string `json:"time"`
}

type Memory struct {
	Total       string  `json:"total"`
	Available   string  `json:"available"`
	Free        string  `json:"free"`
	Used        string  `json:"used"`
	UsedPercent float64 `json:"usedPercent"`
}

type CPU struct {
	Cores       int     `json:"cores"`
	Model       string  `json:"model"`
	Mhz         float64 `json:"mhz"`
	UsedPercent float64 `json:"usedPercent"`
}

func (c *WorkbenchController) Home(ctx http.Context) http.Response {
	var monitor Monitor
	monitor.Version = c.GetVersion()
	monitor.Today = c.GetToday()
	monitor.Memory = c.GetMemory()
	monitor.CPU = c.GetCPU()
	return c.response.Success(ctx, monitor)
}

func (c *WorkbenchController) GetVersion() (v Version) {
	hostInfo, err := host.Info()
	if err != nil {
		return
	}
	v.Framework = support.Version
	v.Admin = facades.Config().GetString("app.version")
	v.OsArch = hostInfo.Platform + hostInfo.PlatformVersion
	return
}

func (c *WorkbenchController) GetToday() (t Today) {
	t.Time = carbon.Now().ToDateTimeString(carbon.PRC)
	return
}

func (c *WorkbenchController) GetMemory() (m Memory) {
	v, _ := mem.VirtualMemory()
	m = Memory{
		Total:     fmtutil.DataSize(v.Total),
		Available: fmtutil.DataSize(v.Available),
		Free:      fmtutil.DataSize(v.Free),
		Used:      fmtutil.DataSize(v.Used),
	}
	m.UsedPercent, _ = strconv.ParseFloat(fmt.Sprintf("%.2f", v.UsedPercent), 64)
	return
}

func (c *WorkbenchController) GetCPU() (u CPU) {
	u.Cores, _ = cpu.Counts(true)
	cpuInfo, err := cpu.Percent(time.Second, false)
	if err == nil {
		u.UsedPercent, _ = strconv.ParseFloat(fmt.Sprintf("%.2f", cpuInfo[0]), 64)
	}

	cpuInfos, _ := cpu.Info()
	u.Model = cpuInfos[0].ModelName
	u.Mhz = cpuInfos[0].Mhz
	return
}
