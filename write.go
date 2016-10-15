package memory

// #define _GNU_SOURCE
// #include <sys/uio.h>
// #include <errno.h>
// #include <stdio.h>
// #include <stdlib.h>
// #include <string.h>
//
// ssize_t
// goMemoryWrite(pid_t pid,int64_t addr, void *buf, size_t bufSize)
// {
// struct iovec local[1];
// struct iovec remote[1];
// ssize_t nread;
//
//
//
// local[0].iov_base = buf;
// local[0].iov_len = bufSize;
// remote[0].iov_base = (void *) addr;
// remote[0].iov_len = bufSize;
//
// //printf("XXXX  WRITE %d %d to pid %d at %x \n", bufSize, *(int *)buf, pid, (void *)addr);
// nread = process_vm_writev(pid, local, 1, remote, 1, 0);
// if (nread <= 0) {
//  printf("XXX WRITE ERROR IS %s\n", strerror(errno));
// }
// return nread;
// }
import "C"
import (
	"bytes"
	"encoding/binary"
	"fmt"
)

func (p *Process) Write(addr int64, bufSize uint64, input interface{}) error {

	var b bytes.Buffer

	if err := binary.Write(&b, binary.LittleEndian, input); err != nil {
		return fmt.Errorf("Could not write from bytes: %s", err)
	}

	inBuf := b.Bytes()
	ptrInBuf := C.CBytes(inBuf)
	defer C.free(ptrInBuf)

	lenOutBuf := C.goMemoryWrite(p.Pid, C.int64_t(addr), ptrInBuf, C.size_t(bufSize))

	if uint64(lenOutBuf) == bufSize {
		return nil
	}

	return fmt.Errorf("Buffer size mismatch size=%d, requested=%d", int(lenOutBuf), bufSize)
}

func (p *Process) WriteInt64(addr int64, input int64) error {
	inBufSize := uint64(8)
	return p.Write(addr, inBufSize, &input)
}

func (p *Process) WriteInt32(addr int64, input int32) error {
	inBufSize := uint64(4)
	return p.Write(addr, inBufSize, &input)
}
