//go:build darwin || linux || bsd

package service

import "golang.org/x/sys/unix"

func GetDiskSpace(path string) (uint64, uint64, error) {
	var stats unix.Statfs_t
	err := unix.Statfs(path, &stats)
	if err != nil {
		return 0, 0, err
	}

	totalDisk := stats.Blocks * uint64(stats.Bsize)
	freeDisk := stats.Bavail * uint64(stats.Bsize)

	return totalDisk, freeDisk, nil
}
