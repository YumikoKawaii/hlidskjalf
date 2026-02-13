//go:build linux

package proxy

import (
	"encoding/binary"
	"fmt"
	"net"
	"unsafe"

	"golang.org/x/sys/unix"
)

const soOriginalDst = 80

func GetOriginalDst(conn net.Conn) (string, error) {
	tc, ok := conn.(*net.TCPConn)
	if !ok {
		return "", fmt.Errorf("not a TCP connection")
	}

	raw, err := tc.SyscallConn()
	if err != nil {
		return "", fmt.Errorf("syscall conn: %w", err)
	}

	var dst string
	var controlErr error

	err = raw.Control(func(fd uintptr) {
		var addr unix.RawSockaddrInet4
		addrLen := uint32(unsafe.Sizeof(addr))

		_, _, errno := unix.Syscall6(
			unix.SYS_GETSOCKOPT,
			fd,
			unix.SOL_IP,
			soOriginalDst,
			uintptr(unsafe.Pointer(&addr)),
			uintptr(unsafe.Pointer(&addrLen)),
			0,
		)
		if errno != 0 {
			controlErr = fmt.Errorf("getsockopt SO_ORIGINAL_DST: %w", errno)
			return
		}

		port := binary.BigEndian.Uint16((*[2]byte)(unsafe.Pointer(&addr.Port))[:])
		dst = fmt.Sprintf("%s:%d", net.IP(addr.Addr[:]).String(), port)
	})
	if err != nil {
		return "", fmt.Errorf("raw control: %w", err)
	}
	if controlErr != nil {
		return "", controlErr
	}

	return dst, nil
}
