package main

import (
	"fmt"
	"net/http"

	"github.com/FATHOM5/godbus"
	"github.com/FATHOM5/godbus/spec"
	"github.com/FATHOM5/rest"
)

func (h *handler) routes() {
	h.router.Handle("/pdu", rest.M.StrictPOST(h.pduHandler()))
}

func (handler) pduHandler() http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		// parse request pdu
		var pdu godbus.PDU
		err := rest.ParseClientRequest(r, &pdu)
		if err != nil {
			_ = rest.RespondUnknownClientParams(w, rest.ErrorFields{
				"reason": err,
			})
			return
		}

		fmt.Println("PDU Request", pdu)

		//
		// logic to read/write from the pdu goes here
		//

		// create response pdu
		respPDU := godbus.PDU{
			FunctionCode: pdu.FunctionCode,
			Data:         spec.ReadHoldingRegistersResp(4, []byte{0, 0, 0, 1}),
		}

		// response with the pdu
		rest.RespondOK(w, respPDU)
	})
}
