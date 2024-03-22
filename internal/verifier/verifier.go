package verifier

import (
	"fmt"
	"log"
	"net"
	"strings"
)

// CheckDomain checks the DNS records for a domain
func CheckDomain(domain string) {
	// Variables to store the results of the DNS lookups
	var (
		// MX is the mail exchange record
		hasMX bool
		// SPF is the sender policy framework record
		hasSPF bool
		// DMARC is the domain-based message authentication, reporting, and conformance record
		hasDMARC bool
		// spfRecord is the SPF record for the domain
		spfRecord string
		// dmarcRecord is the DMARC record for the domain
		dmarcRecord string
	)

	// MX
	// - lookup the MX records for the domain
	mxRecords, err := net.LookupTXT(domain)
	if err != nil {
		log.Printf("Error looking up MX records for '%s': %v", domain, err)
		return
	}
	// - if there are MX records, the domain has mail servers
	if len(mxRecords) > 0 {
		hasMX = true
	}

	// SPF
	txtRecords, err := net.LookupTXT(domain)
	if err != nil {
		log.Printf("Error looking up TXT records for '%s': %v", domain, err)
		return
	}
	// - if there is a TXT record starting with "v=spf", the domain has an SPF record
	for _, record := range txtRecords {
		if strings.HasPrefix(record, "v=spf1") {
			hasSPF = true
			spfRecord = record
			break
		}
	}

	// DMARC
	// - lookup the DMARC record for the domain
	dmarcRecords, err := net.LookupTXT("_dmarc." + domain)
	if err != nil {
		log.Printf("Error looking up DMARC records for '%s': %v", domain, err)
		return
	}
	// - if there is a TXT record starting with "v=DMARC", the domain has a DMARC record
	for _, record := range dmarcRecords {
		if strings.HasPrefix(record, "v=DMARC1") {
			hasDMARC = true
			dmarcRecord = record
			break
		}
	}

	fmt.Printf("Domain: %v\nhasMX: %v\nhasSPF: %v\nhasDMARC: %v\nspfRecord: %v\ndmarcRecord: %v\n", domain, hasMX, hasSPF, hasDMARC, spfRecord, dmarcRecord)
	fmt.Println("--------------------------------------------------")
}
