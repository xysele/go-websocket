package main

import (
	"log"

	core "github.com/xtls/xray-core"
)

func main() {
	// 创建 Xray 实例
	instance, err := core.New(&core.Config{
		Inbound: []*core.InboundHandlerConfig{
			{
				ReceivedPacketHostPolicy: &core.ReceivedPacketHostPolicy{
					Overrides: []*core.StringInboundDetourConfig{
						{
							Protocol: "vmess",
							PortRange: &core.PortRange{
								From: 3000,
								To:   3000,
							},
							Settings: &core.VMess{
								Clients: []*core.VMessUser{
									{
										Id: "d4bf22df-0342-462c-bb18-d615eece0c22",
										AlterId: 0,
									},
								},
							},
							StreamSettings: &core.StreamSettings{
								Network: "ws",
							},
						},
					},
				},
			},
		},
		Outbound: []*core.OutboundHandlerConfig{
			{
				Protocol: "freedom",
				Settings: &core.OutboundFreedomConfig{},
			},
		},
	})
	if err != nil {
		log.Fatalf("Failed to create Xray instance: %v", err)
	}

	// 运行 Xray 实例
	if err := instance.Start(); err != nil {
		log.Fatalf("Failed to start Xray instance: %v", err)
	}

	// 保持运行,直到被中断
	select {}
}
