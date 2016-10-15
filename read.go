package memory

// #define _GNU_SOURCE
// #include <sys/uio.h>
// #include <errno.h>
// #include <stdio.h>
// #include <stdlib.h>
// #include <string.h>
//
// ssize_t
// goMemoryRead(pid_t pid,int64_t addr, void *buf, size_t bufSize)
// {
// struct iovec local[1];
// struct iovec remote[1];
// ssize_t nread;
//
// local[0].iov_base = buf;
// local[0].iov_len = bufSize;
// remote[0].iov_base = (void *) addr;
// remote[0].iov_len = bufSize;
//
// nread = process_vm_readv(pid, local, 1, remote, 1, 0);
// //printf("XXXX  READ %d from pid %d at %x \n", nread, pid, (void *)addr);
// if (nread <= 0) {
//  printf("XXX ERROR IS %s\n", strerror(errno));
// }
// return nread;
// }
//
import "C"
import (
	"bytes"
	"encoding/binary"
	"fmt"
)

func (p *Process) Read(addr int64, bufSize uint64, result interface{}) error {
	inBuf := make([]byte, bufSize)
	ptrInBuf := C.CBytes(inBuf)
	defer C.free(ptrInBuf)

	lenOutBuf := C.goMemoryRead(p.Pid, C.int64_t(addr), ptrInBuf, C.size_t(bufSize))

	if uint64(lenOutBuf) == bufSize {
		outBuf := C.GoBytes(ptrInBuf, C.int(lenOutBuf))

		b := bytes.NewReader(outBuf)

		fmt.Println("XXX %v", outBuf)
		if err := binary.Read(b, binary.LittleEndian, result); err != nil {
			return fmt.Errorf("Could not read from bytes: %s", err)
		}

		return nil
	}

	return fmt.Errorf("Buffer size mismatch size=%d, requested=%d", int(lenOutBuf), bufSize)
}

func (p *Process) ReadInt64(addr int64) (int64, error) {
	var result int64
	inBufSize := uint64(8)

	err := p.Read(addr, inBufSize, &result)

	if err != nil {
		return 0, fmt.Errorf("Could not read int64: %s", err)
	}

	fmt.Println("XXX GOT VAL", result)
	return result, nil
}

func (p *Process) ReadInt32(addr int64) (int32, error) {
	var result int32
	inBufSize := uint64(4)

	err := p.Read(addr, inBufSize, &result)

	if err != nil {
		return 0, fmt.Errorf("Could not read int32: %s", err)
	}

	return result, nil
}

func (p *Process) ReadBool(addr int64) (bool, error) {
	var result int8
	inBufSize := uint64(1)

	err := p.Read(addr, inBufSize, &result)

	if err != nil {
		return false, fmt.Errorf("Could not read bool: %s", err)
	}

	return result != 0, nil
}
