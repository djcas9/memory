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
}
