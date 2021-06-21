package main

import (
	"runtime"
	"fmt"
	"github.com/google/gopacket"
)

func main() {
	cap1 := 100000000
	pkts := make(chan gopacket.Packet, cap1)

	var m runtime.MemStats
        runtime.ReadMemStats(&m)
	fmt.Printf("capacity %v\n", cap1)
        fmt.Printf("Alloc = %v B\n", m.Alloc)
        fmt.Printf("TotalAlloc = %v B\n", m.TotalAlloc)
	
	//cap2 := 100000000
	//pkts = make(chan gopacket.Packet, cap2)
	//fmt.Printf("capacity %v\n", cap2)

        runtime.ReadMemStats(&m)
        fmt.Printf("Alloc = %v B\n", m.Alloc)
        fmt.Printf("TotalAlloc = %v B\n", m.TotalAlloc)

	<- pkts
}

