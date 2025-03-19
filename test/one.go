package main

import (
	"fmt"
	"syscall"
	"unsafe"
)

const (
	PROCESS_ALL_ACCESS = 0x1F0FFF
)

// 加载 kernel32.dll
var (
	kernel32              = syscall.NewLazyDLL("kernel32.dll")
	procOpenProcess       = kernel32.NewProc("OpenProcess")
	procReadProcessMemory = kernel32.NewProc("ReadProcessMemory")
)

// 打开进程句柄
func openProcess(pid uint32) (syscall.Handle, error) {
	handle, _, err := procOpenProcess.Call(uintptr(PROCESS_ALL_ACCESS), uintptr(0), uintptr(pid))
	if handle == 0 {
		return 0, err
	}
	return syscall.Handle(handle), nil
}

// 读取进程内存
func readProcessMemory(hProcess syscall.Handle, lpBaseAddress uintptr, lpBuffer []byte, nSize uintptr) (int, error) {
	var bytesRead uintptr
	r, _, err := procReadProcessMemory.Call(
		uintptr(hProcess),
		lpBaseAddress,
		uintptr(unsafe.Pointer(&lpBuffer[0])),
		nSize,
		uintptr(unsafe.Pointer(&bytesRead)),
	)
	if r == 0 {
		return 0, err
	}
	return int(bytesRead), nil
}

func main() {
	// 替换为实际的进程 ID
	pid := uint32(2580)
	// 替换为实际的内存地址
	lpAddress := uintptr(0x01005361)

	// 打开进程
	hProcess, err := openProcess(pid)
	if err != nil {
		fmt.Printf("无法打开进程: %v\n", err)
		return
	}
	defer syscall.CloseHandle(hProcess)

	// 读取内存
	buffer := make([]byte, 4)
	bytesRead, err := readProcessMemory(hProcess, lpAddress, buffer, uintptr(len(buffer)))
	if err != nil {
		fmt.Printf("读取内存失败: %v\n", err)
	} else {
		fmt.Printf("读取到 %d 字节: %v\n", bytesRead, buffer)
	}
}
