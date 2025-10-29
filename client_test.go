package mdns

import (
	"net"
	"testing"
	"time"

	"github.com/johannes-kuhfuss/services_utils/logger"
)

func TestClient(t *testing.T) {

	iface, _ := net.InterfaceByName("Ethernet")

	entriesCh := make(chan *ServiceEntry, 32)
	go func() {
		for entry := range entriesCh {
			logger.Infof("Found device %v\r\n", entry.Name)
		}
	}()

	queryParams := &QueryParam{
		Service:             "_services._dns-sd._udp",
		Domain:              "local",
		Timeout:             5 * time.Second,
		Interface:           iface,
		Entries:             entriesCh,
		WantUnicastResponse: false,
		DisableIPv4:         false,
		DisableIPv6:         false,
	}
	logger.Infof("Starting to query services (%v) on interface %v", queryParams.Service, queryParams.Interface.Name)
	err := Query(queryParams)
	if err != nil {
		logger.Errorf("Error while querying audio devices: %v", err)
	}
	close(entriesCh)
}
