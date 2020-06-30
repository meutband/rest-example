package main

import (
	"fmt"
	"net/http"

	"github.com/FATHOM5/godbus"
	"github.com/FATHOM5/godbus/spec"
	"github.com/FATHOM5/rest"
)

const serverAddress = "localhost:8000"

// Client is a struct that contains the rest client and server address
type Client struct {
	client        *rest.RestfulClient
	serverAddress string
}

// NewClient creates a new Client type
func NewClient(addr string) *AddClient {
	c := new(Client)
	c.client = rest.NewRestClient(http.DefaultClient)
	c.serverAddress = addr

	return c
}

// Request is a method of Client that POSTs to the client serverAddress.
func (cli Client) Request(pdu godbus.PDU) (reply godbus.PDU, err error) {
	rc := cli.client
	_, err = rc.Post(rc.Join(cli.serverAddress, "pdu"), pdu, &reply)
	return
}

func main() {

	// Fake read of PDU from modbus
	testPDU := godbus.PDU{
		FunctionCode: spec.ReadHoldingRegisters,
		Data:         spec.ReadHoldingRegistersReq(3001, 2),
	}

	// create new client type
	cli := NewAddClient(serverAddress)

	// post the pdu to the server
	reply, err := cli.Request(testPDU)
	if err != nil {
		fmt.Println("Request failed")
		fmt.Println(err)
		return
	}

	fmt.Printf("%v = %v\n", testPDU, reply)

}
