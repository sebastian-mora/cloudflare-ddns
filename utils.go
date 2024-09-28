package main

import "github.com/cloudflare/cloudflare-go"

// FindDnsRecordInSlice searches for a DNS record by name in a slice.
func FindDnsRecordInSlice(name string, records []cloudflare.DNSRecord) *cloudflare.DNSRecord {
	for _, record := range records {
		if record.Name == name {
			return &record
		}
	}
	return nil
}
