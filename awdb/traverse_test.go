package awdb

import (
	"fmt"
	"log"
	"testing"
)

func TestNetworks_Next_IPv4(t *testing.T) {
	db, err := Open("/Users/bytedance/Documents/iplib/aiwen/IP_city_single_WGS84.awdb")
	if err != nil {
		log.Fatal(err)
	}
	_ = db
	networks := db.Networks()
	if networks == nil {
		return
	}
	cnt := 0
	var record interface{}
	for networks.Next() {
		cnt++
		if cnt > 10 {
			break
		}
		if ipnet, err := networks.Network(&record); err == nil {
			if resMap, ok := record.(map[string]interface{}); !ok {
				break
			} else {
				fmt.Printf("accuracy: %s, ipnet: %v, record: %s\n", resMap["accuracy"], ipnet, record)
			}
		} else {
			break
		}
	}
}
func TestNetworks_Next_IPv6(t *testing.T) {
	db, err := Open("/Users/bytedance/Documents/iplib/aiwen/IP_city_single_BD09_WGS84_ipv6.awdb")
	if err != nil {
		log.Fatal(err)
	}
	_ = db
	networks := db.Networks()
	if networks == nil {
		return
	}
	cnt := 0
	var record interface{}
	for networks.Next() {
		cnt++
		if cnt > 40 {
			break
		}
		if ipnet, err := networks.Network(&record); err == nil {
			if resMap, ok := record.(map[string]interface{}); !ok {
				break
			} else {
				fmt.Printf("accuracy: %s, ipnet: %v, record: %s\n", resMap["accuracy"], ipnet, record)
			}
		} else {
			break
		}
	}
}
