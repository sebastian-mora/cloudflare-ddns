package main

import (
	"context"
	"fmt"
	"log/slog"
	"os"

	"github.com/cloudflare/cloudflare-go"
)

func main() {
	CF_TOKEN := os.Getenv("CF_TOKEN")
	CONFIG_FILE := os.Getenv("CONFIG_FILE")

	slog.SetLogLoggerLevel(slog.LevelDebug)

	if CF_TOKEN == "" {
		slog.Error("CF_TOKEN not set")
		os.Exit(1)
	}

	// Construct a new API object using a global API key
	api, err := cloudflare.NewWithAPIToken(CF_TOKEN)
	if err != nil {
		slog.Error("failed to create cloudflare api client", "error", err)
		os.Exit(1)
	}

	if CONFIG_FILE == "" {
		CONFIG_FILE = "./config.yaml"
		slog.Info("CONFIG_FILE not set, defaulting to ./config.yaml")
	}

	config, err := LoadConfig(CONFIG_FILE)
	slog.Info("loaded records", "filename", CONFIG_FILE)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(1)
	}

	externalIp, err := GetExternalIP(config.ExternalIPService)
	if err != nil {
		slog.Error("failed to fetch external ip", "error", err)
	}

	slog.Info("fetched external ip", "ip", externalIp)

	for _, zone := range config.Zones {
		ctx := context.Background()
		// Load the records for that zone
		remoteRecords, err := GetRecords(ctx, api, zone.ZoneID)
		if err != nil {
			fmt.Printf("Error fetching records for zone %s: %v\n", zone.Name, err)
			continue // Skip to the next zone if there’s an error
		}

		for _, r := range zone.Records {
			// Find the remote record by its name
			remoteRecord := FindDnsRecordInSlice(r.Name, remoteRecords)

			if remoteRecord != nil && remoteRecord.Type == "A" && remoteRecord.Content != externalIp {

				r.Content = externalIp

				err := UpdateRecord(ctx, api, zone.ZoneID, remoteRecord.ID, r)
				if err != nil {
					slog.Error("failed to update record", "error", err)
					continue
				}

				slog.Info("updated A record", "name", remoteRecord.Name, "old_ip", remoteRecord.Content, "new_ip", externalIp)
			}

			if remoteRecord == nil {
				r.Content = externalIp

				fmt.Printf("Creating DNS record with parameters: Type=%s, Name=%s, Content=%s, TTL=%d, Proxied=%v\n",
					r.Type, r.Name, r.Content, r.TTL, r.Proxied)
				err := CreateRecord(ctx, api, zone.ZoneID, r)
				if err != nil {
					slog.Error("failed to create record", "error", err)
					continue
				}

				slog.Info("created A record", "name", r.Name, "new_ip", externalIp)
			}
		}
	}
}
