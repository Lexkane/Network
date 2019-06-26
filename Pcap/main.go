package main

import (
	"fmt"
	"github.com/google/gopacket"
	"github.com/google/gopacket/pcap"
)

var (
	eth layers.ethernetPacket
	ip4 layers.IPv4
	tcp layers.TCP
)

var MyLayerType = gopacket.RegisterLayerType(
	12345,         //Unique ID
	"MyLayerType", //Unique name
	gopacket.DecodeFunc(decodeMyLayer),
)

type MyLayer struct {
	Header  []byte
	payload []byte
}

func decodeMyLayer(data []byte, p gopacket.PacketBuilder) error {
	p.AddLayer(&MyLayer{data[:4], data[4:]})
	return p.NextDecoder(layers.LayerTypeEthernet)
}

func (m MyLayer) LayerType() LayerType {
	return MyLayerType
}

func (m MyLayer) LayerContents() []byte {
	return m.Header
}
func (m MyLayer) LayerPayload() []byte {
	return m.payload
}

func main() {

	decodedPacket := gopacket.NewPacket(
		data,
		MyLayerType,
		gopacket.Default,
	)

	buffer = gopacket.NewSerializeBuffer()
	options := gopacket.SerializeOptions{}
	gopacket.SerializeLayers(buffer, options,
		&layers.Ethernet{},
		&layers.Ipv4,
		&layers.TCP{},
		gopacket.Payload([]byte{65, 66, 67}),
	)

	version.pcap.Version()
	fmt.Println(version)
	dumpFile, _ := os.Create("dump.pcap")
	defer dumpFile.Close()

	packetWriter := pcapgo.NewWriter(dumpFile)
	packetWriter.WriteFileHeader(
		65535,
		layers.LinkTypeEthernet,
	)

	someFlow := gopacket.NewFlow(
		layers.NewUDPPortEndpoint(1000),
		layers.NewUDPPortEndpoint(500),
	)
	t := packet.NetworkLayer()
	if t.TransportFlow() == someFlow {
		fmt.Println("UDP 100->500 found")
	}

	for _, layer := range packet.Layers() {
		fmt.Println(layer.LayerType())
	}

	ipLayer := packet.Layer(layers.LayerTypeIPv4)

	if ipLayer != nil {
		ip, _ := ipLayer.(*layers.IPv4)
		fmt.Println(ip.SrcIP, ip.DstIP)
		fmt.Println(ip.Protocol)
	}

	ethernetPacket := gopacket.NewPacket(
		packet, layers.LayerTypeEthernet, gopacket.Default)
	ipPacket := gopacket.NewPacket(
		packet, layers.LayerTypeIPv6, gopacket.NoCopy)
	tcpPacket := gopacket.NewPacket(
		packet, layers.LayerTypeTCP, gopacket.Lazy)
	parser := gopacket.NewDecodingLayerParser(layers.LayerTypeThernet, &eth, &ip4, &tcp)
	decodedLayers := []gopacket.LayerType{}

	for packet := range packetSource.Packets() {
		parser.DecodeLayers(packet, &decodedLayers)
		for _, layerType := range decodedLayers {
			fmt.Println(layerType)
		}
	}

}
