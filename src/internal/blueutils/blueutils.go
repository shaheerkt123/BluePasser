package blueutils

import (
	"encoding/json"
	"log"

	"tinygo.org/x/bluetooth"
)

var (
	Adapter     = bluetooth.DefaultAdapter
	ServiceUUID = bluetooth.NewUUID([16]byte{0xf0, 0xb2, 0xa5, 0xa2, 0x3c, 0x9a, 0x4f, 0x4a, 0x8f, 0x8b, 0x2e, 0x4b, 0x7e, 0x3b, 0x1b, 0x1a})
	CharUUID    = bluetooth.NewUUID([16]byte{0xf0, 0xb2, 0xa5, 0xa2, 0x3c, 0x9a, 0x4f, 0x4a, 0x8f, 0x8b, 0x2e, 0x4b, 0x7e, 0x3b, 0x1b, 0x1b})
)

type WifiCredentials struct {
	SSID     string `json:"ssid"`
	Password string `json:"password"`
}

func StartBroadcasting(ssid, password string) error {
	if err := Adapter.Enable(); err != nil {
		return err
	}

	creds := WifiCredentials{SSID: ssid, Password: password}
	data, err := json.Marshal(creds)
	if err != nil {
		return err
	}

	adv := Adapter.DefaultAdvertisement()
	if err := adv.Configure(bluetooth.AdvertisementOptions{
		LocalName: "BluePasser",
	}); err != nil {
		return err
	}

	if err := adv.Start(); err != nil {
		return err
	}

	err = Adapter.AddService(&bluetooth.Service{
		UUID: ServiceUUID,
		Characteristics: []bluetooth.CharacteristicConfig{
			{
				UUID:  CharUUID,
				Value: data,
				Flags: bluetooth.CharacteristicReadPermission,
			},
		},
	})

	if err != nil {
		return err
	}

	log.Println("Broadcasting... Press Ctrl+C to stop")
	select {}
}

func ScanForCredentials(handler func(creds WifiCredentials)) error {
	if err := Adapter.Enable(); err != nil {
		return err
	}

	log.Println("Scanning...")
	return Adapter.Scan(func(adapter *bluetooth.Adapter, device bluetooth.ScanResult) {
		if device.LocalName() == "BluePasser" {
			log.Println("Found BluePasser device")
			adapter.StopScan()
			dev, err := adapter.Connect(device.Address, bluetooth.ConnectionParams{})
			if err != nil {
				log.Println(err)
				return
			}

			log.Println("Connected")

			services, err := dev.DiscoverServices([]bluetooth.UUID{ServiceUUID})
			if err != nil {
				log.Println(err)
				return
			}

			if len(services) == 0 {
				log.Println("Could not find BluePasser service")
				return
			}

			service := services[0]
			chars, err := service.DiscoverCharacteristics([]bluetooth.UUID{CharUUID})
			if err != nil {
				log.Println(err)
				return
			}

			if len(chars) == 0 {
				log.Println("Could not find characteristic")
				return
			}

			char := chars[0]
			buf := make([]byte, 256)
			n, err := char.Read(buf)
			if err != nil {
				log.Println(err)
				return
			}

			var creds WifiCredentials
			if err := json.Unmarshal(buf[:n], &creds); err == nil {
				handler(creds)
			}

			dev.Disconnect()
		}
	})
}
