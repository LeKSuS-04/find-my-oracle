package fetch

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strings"
)

const (
	PUBLIC_IPS_URL = "https://docs.oracle.com/en-us/iaas/tools/public_ip_ranges.json"
)

type RegionError struct {
	AvailableRegions []string
}

func (e *RegionError) Error() string {
	return "Invalid region. Available Regions:\n" +
		"    " + strings.Join(e.AvailableRegions, ",\n    ") + "."
}

type cidr struct {
	IPMask string   `json:"cidr"`
	Tags   []string `json:"tags"`
}
type region struct {
	Name  string `json:"region"`
	CIDRs []cidr `json:"cidrs"`
}
type oracleResponse struct {
	LastUpdatedTimestamp string   `json:"last_updated_timestamp"`
	Regions              []region `json:"regions"`
}

func FetchIPMasks(searchRegion string) ([]string, error) {
	r, err := http.Get(PUBLIC_IPS_URL)
	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return nil, err
	}

	var data oracleResponse
	json.Unmarshal(body, &data)

	var availableRegions []string
	for _, receivedRegion := range data.Regions {
		if receivedRegion.Name == searchRegion {
			ips := []string{}
			for _, regionCIDR := range receivedRegion.CIDRs {
				ips = append(ips, regionCIDR.IPMask)
			}
			return ips, nil
		}
		availableRegions = append(availableRegions, receivedRegion.Name)
	}

	return []string{}, &RegionError{AvailableRegions: availableRegions}
}
