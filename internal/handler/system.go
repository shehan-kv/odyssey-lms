package handler

import (
	"encoding/json"
	"net/http"
	"os"
	"path"
	"runtime"

	"github.com/pbnjay/memory"
	"odyssey.lms/internal/service"
)

type systemInfo struct {
	Cpu          int    `json:"cpu"`
	Memory       uint64 `json:"memory"`
	UploadsFree  uint64 `json:"uploadsFree"`
	UploadsTotal uint64 `json:"uploadsTotal"`
}

func GetSystemInfo(w http.ResponseWriter, h *http.Request) {
	var info systemInfo

	info.Cpu = runtime.NumCPU()

	var m runtime.MemStats
	runtime.ReadMemStats(&m)

	info.Memory = memory.TotalMemory()

	cwd, _ := os.Getwd()
	uploadsPath := path.Join(cwd, "uploads")

	totalDisk, freeDisk, err := service.GetDiskSpace(uploadsPath)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	info.UploadsTotal = totalDisk
	info.UploadsFree = freeDisk

	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(&info)
}
