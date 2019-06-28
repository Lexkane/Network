package ping

import (
	"fmt"
	"github.com/google/gopacket"
	"github.com/google/gopacket/layers"
	"github.com/google/gopacket/pcap"
	"log"
	"time"
)

var (
	device  string = "en0"
	snaplen int32  = 65535
	promisc bool   = false
	err     error
	timeout time.Duration = -1 * time.Second
	handle  *pcap.Handle
)

func main() {
	handle, err = pcap.OpenLive(device, snaplen, promisc, timeout)
	if err != nil {
		log.Fatal(err)
	}
	defer handle.Close()

	var filter string = "src host 172.20.1.8 and icmp"
	err = handle.SetBPFFilter(filter)
	if err != nil {
		log.Fatal(err)
	}

	packetSource := gopacket.NewPacketSource(handle, handle.LinkType())
	for packet := range packetSource.Packets() {
		fmt.Println("Someone ping me")
		fmt.Println("----------------")
		fmt.Println(packet)

		ethernet_layer := packet.Layer(layers.LayerTypeEthernet)
		ethernet_frame := ethernet_layer.(*layers.Ethernet)
		ip_layer := packet.Layer(layers.LayerTypeIPv4)
		ip_packet, _ := ip_layer.(*layers.IPv4)
		icmp_layer := packet.Layer(layers.LayerTypeICMPv4)
		icmp_packet := icmp_layer.(*layers.ICMPv4)

		if icmp_packet.TypeCode.String() == "EchoRequest" {
			if len(icmp_packet.Payload) > 0 {
				log.Println("Info:EchoRequest Received")
			} else {
				log.Println("Wan: Empty EchoRequest Received")
				ethernet_frame_copy := *ethernet_frame
				ip_packet_copy := *ip_packet
				icmp_packet_copy := *icmp_packet

				ethernet_frame_copy.SrcMAC = ethernet_frame.DstMAC
				ethernet_frame_copy.DstMAC = ethernet_frame.SrcMAC

				ip_packet_copy.SrcIP = ip_packet.DstIP
				ip_packet_copy.DstIP = ip_packet.SrcIP

				icmp_packet_copy.TypeCode = layers.ICMPv4TypeEchoReply
				var buffer gopacket.SerializeBuffer
				var options gopacket.SerializeOptions
				options.ComputeChecksums = true

				buffer = gopacket.NewSerializeBuffer()
				err := gopacket.SerializeLayers(buffer,
					options,
					&ethernet_frame_copy,
					&ip_packet_copy,
					gopacket.Payload(icmp_packet_copy.Payload), )
				if err != nil {
					log.Fatal(err)
				}

				new_message := buffer.Bytes()
				err = handle.WritePacketData(new_message)
				if err != nil {
					log.Fatal(err)
				}
			}
		}
		fmt.Println("------------------")
		fmt.Println("Source Adress" + ip_packet.SrcIP.String())
		fmt.Println("Destination Adress" + ip_packet.DstIP.String())
		fmt.Println("Protocol: ", ip_packet.Protocol)
		fmt.Println("------------------")
		fmt.Println("ICMP Code", icmp_packet.TypeCode)
		fmt.Println("ICMP Sequence Number", icmp_packet.Seq)
		fmt.Println("Payload data length", len(icmp_packet.Payload))
		fmt.Println("Payload data", icmp_packet.Payload)
		fmt.Println("Payload data to string ", string(icmp_packet.Payload))
		fmt.Println("------------------\n")

	}

}
