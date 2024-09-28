package main

import "github.com/cloudflare/cloudflare-go"

func FindDnsRecordInSlice(name string, records []cloudflare.DNSRecord) *cloudflare.DNSRecord {
	for _, v := range records {
		if v.Name == name {
			return &v
		}
	}
	return nil
}
