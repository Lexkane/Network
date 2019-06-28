package offline

import (
	"fmt"
	"github.com/google/gopacket"
	"github.com/google/gopacket/pcap"
	"log"
)

var (
	handle *pcap.Handle
	err    error
)

func main() {
	//Open pcap file
	handle, err = pcap.OpenOffline("test.pcap")
	if err != nil {
		log.Fatal(err)
	}

	defer handle.Close()

	//Process packets

	packetSource := gopacket.NewPacketSource(handle, handle.LinkType())
	for packet := range packetSource.Packets() {
		fmt.Println(packet)
	}
}
