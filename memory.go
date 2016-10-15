package memory

// #include <sys/uio.h>
import "C"

type Process struct {
	Pid C.pid_t
}
