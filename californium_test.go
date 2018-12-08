package coap

import (
	"fmt"
	"testing"
)

var californiumServer = "californium.eclipse.org:5683"

func TestCalifornium(t *testing.T) {
	szx := BlockWiseSzx16
	client := &Client{Net: "udp", BlockWiseTransferSzx: &szx}
	co, err := client.Dial(californiumServer)
	if err != nil {
		t.Fatalf("Error dialing: %v", err)
	}

	fmt.Printf("conn: %v\n", co.LocalAddr())

	resp, err := co.Get("")
	if err != nil {
		t.Fatalf("Cannot post exchange")
	}
	decodeMsg(resp)
}

func TestCaliforniumTCP(t *testing.T) {
	enable := true
	client := &Client{Net: "tcp", BlockWiseTransfer: &enable}
	co, err := client.Dial(californiumServer)
	if err != nil {
		t.Fatalf("Error dialing: %v", err)
	}

	fmt.Printf("conn: %v\n", co.LocalAddr())

	resp, err := co.Get("")
	if err != nil {
		t.Fatalf("Cannot post exchange")
	}
	decodeMsg(resp)
}
