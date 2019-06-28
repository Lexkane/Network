package filters

import (
	"fmt"
	"github.com/google/gopacket"
	"github.com/google/gopacket/layers"
	"github.com/google/gopacket/pcap"
	"github.com/google/gopacket/pcapgo"
	"log"
	"os"
	"time"
)

var (
	deviceName   string = "en0"
	dumpFilename string = "test.pcap"
	snapshotLen  int32  = 65535
	promiscuous  bool   = false
	err          error
	timeout      time.Duration = -1 * time.Second
	handle       *pcap.Handle
	packetCount  int = 0
)

func main() {
	//Open output pcap ile and write header
	dumpFile, _ := os.Create(dumpFilename)
	packetWriter := pcapgo.NewWriter(dumpFile)
	packetWriter.WriteFileHeader(65535, layers.LinkTypeEthernet)
	defer dumpFile.Close()

	//Open device for capturing
	handle, err = pcap.OpenLive(deviceName, snapshotLen, promiscuous, timeout)
	if err != nil {
		log.Fatal("Error opening device %s : %v", deviceName, err)
	}
	defer handle.Close()
	packetSource := gopacket.NewPacketSource(handle, handle.LinkType())

	//Start processing packets
	for packet := range packetSource.Packets() {
		fmt.Println(packet)
		packetWriter.WritePacket(packet.Metadata().CaptureInfo, packet.Data())
		packetCount++

		//Only capture 100 packets
		if packetCount > 100 {
			break
		}
	}

}
