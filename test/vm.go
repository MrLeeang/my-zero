package main

import (
	"fmt"

	libvirt "github.com/libvirt/libvirt-go"
)

func main() {

	conn, err := libvirt.NewConnect("qemu:///system")

	if err != nil {
		panic(err)
	}

	defer conn.Close()

	fmt.Println(conn)
}
