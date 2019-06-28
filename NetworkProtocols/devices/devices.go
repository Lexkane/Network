package devices

import (
	"fmt"
	"github.com/google/gopacket/pcap"
	"log"
)

func main() {
	var devices []pcap.Interface
	devices, err := pcap.FindAllDevs()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Devices found:")
	for _, device := range devices {
		fmt.Println("\nName", device.Name)
		fmt.Println("Description:", device.Description)
		fmt.Println("Devices adresses:")
		for _, adress := range device.Addresses {
			fmt.Println("-IP adress", adress.IP)
			fmt.Println("-Subnet mask", adress.Netmask)
		}
	}
}
