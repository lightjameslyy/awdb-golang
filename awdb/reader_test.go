package awdb

import (
	"fmt"
	"log"
	"net"
	"testing"
)

func TestLookupIPv4(t *testing.T) {
	db, err := Open("/Users/bytedance/Documents/iplib/aiwen/IP_city_single_WGS84.awdb")
	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	ip := net.ParseIP("166.111.4.100")

	var record interface{}
	err = db.Lookup(ip, &record)
	if err != nil {
		log.Fatal(err)
	}
	var resMap = record.(map[string]interface{})
	fmt.Printf("accuracy:%s", resMap["accuracy"])
	fmt.Println()
	fmt.Printf("%s", record)
}

func TestLookupIPv6(t *testing.T) {
	db, err := Open("/Users/bytedance/Documents/iplib/aiwen/IP_city_single_BD09_WGS84_ipv6.awdb")
	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	ip := net.ParseIP("2400:da00::dbf:0:100")

	var record interface{}
	err = db.Lookup(ip, &record)
	if err != nil {
		log.Fatal(err)
	}
	var resMap = record.(map[string]interface{})
	fmt.Printf("accuracy:%s", resMap["accuracy"])
	fmt.Println()
	fmt.Printf("%s", record)
}
