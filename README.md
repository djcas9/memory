# memory
Linux memory read/write lib using process_vm_readv and process_vm_writev.

## Features

  * r/w float32/float64
  * r/w int32/int64
  * r/w Bool

## Example

```go
package main

import (
	"fmt"

	"github.com/mephux/memory"
)

func main() {

	m := memory.Process{
		// change this
		// pmap pid
		// addr must have r/w perms
		Pid: 19249,
	}

	// change this
	var addr int64 = 0x00007f733416f000

	i, err := m.ReadInt32(addr)

	fmt.Println(i, err)

	if err := m.WriteInt32(addr, 55); err != nil {
		panic(err)
	}

	i, err = m.ReadInt32(addr)

	fmt.Println(i, err)

	// m := C.read(pid, 0x10000)

	// fmt.Println(m)
}
```

## TODO

  * Windows
  * Darwin

## Self-Promotion

Like memory? Follow the repository on
[GitHub](https://github.com/mephux/memory) and if
you would like to stalk me, follow [mephux](http://dweb.io/) on
[Twitter](http://twitter.com/mephux) and
[GitHub](https://github.com/mephux).
