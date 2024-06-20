package service

import "golang.org/x/sys/windows"

func GetDiskSpace(path string) (uint64, uint64, error) {
	var totalDisk uint64
	var freeDisk uint64

	err := windows.GetDiskFreeSpaceEx(
		windows.StringToUTF16Ptr(path),
		nil,
		&totalDisk,
		&freeDisk,
	)
	if err != nil {
		return 0, 0, nil
	}

	return totalDisk, freeDisk, nil

}
