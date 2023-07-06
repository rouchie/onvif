// Code generated : DO NOT EDIT.
// Copyright (c) 2022 Jean-Francois SMIGIELSKI
// Distributed under the MIT License

package imaging

import (
	"context"
	"github.com/juju/errors"
	"github.com/use-go/onvif"
	"github.com/use-go/onvif/sdk"
	"github.com/use-go/onvif/imaging"
)

// Call_Move forwards the call to dev.CallMethod() then parses the payload of the reply as a MoveResponse.
func Call_Move(ctx context.Context, dev *onvif.Device, request imaging.Move) (imaging.MoveResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			MoveResponse imaging.MoveResponse
		}
	}
	var reply Envelope
	if httpReply, err := dev.CallMethod(request); err != nil {
		return reply.Body.MoveResponse, errors.Annotate(err, "call")
	} else {
		err = sdk.ReadAndParse(ctx, httpReply, &reply, "Move")
		return reply.Body.MoveResponse, errors.Annotate(err, "reply")
	}
}
