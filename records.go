package main

import (
	"context"
	"fmt"

	"github.com/cloudflare/cloudflare-go"
)

// UpdateRecord updates a DNS record in Cloudflare.
func UpdateRecord(ctx context.Context, api *cloudflare.API, zoneID string, recordId string, r DNSRecord) error {

	// create the zone
	zone := cloudflare.ZoneIdentifier(zoneID)
	// Prepare the record
	record := cloudflare.UpdateDNSRecordParams{
		ID:      recordId,
		Type:    r.Type,
		Name:    r.Name,
		Content: r.Content,
		TTL:     r.TTL,
		Proxied: &r.Proxied,
	}

	_, err := api.UpdateDNSRecord(ctx, zone, record)

	return err
}

// // CreateRecord creates a new DNS record in Cloudflare.
func CreateRecord(ctx context.Context, api *cloudflare.API, zoneID string, r DNSRecord) error {

	record := cloudflare.CreateDNSRecordParams{
		Type:    r.Type,
		Content: "73.70.144.73",
		Name:    r.Name,
		Proxied: &r.Proxied,
	}

	// create zoneId
	zone := cloudflare.ZoneIdentifier(zoneID)

	_, err := api.CreateDNSRecord(ctx, zone, record)
	if err != nil {
		return fmt.Errorf("failed to create DNS record: %v", err)
	}
	fmt.Printf("Created DNS record: %s\n", r.Name)
	return nil
}

// GetRecord retrieves a DNS record from Cloudflare.
func GetRecords(ctx context.Context, api *cloudflare.API, zoneID string) ([]cloudflare.DNSRecord, error) {
	zone := cloudflare.ZoneIdentifier(zoneID)

	records, _, err := api.ListDNSRecords(ctx, zone, cloudflare.ListDNSRecordsParams{})

	if err != nil {
		return nil, fmt.Errorf("failed to fetch DNS records: %v", err)
	}

	return records, err
}
