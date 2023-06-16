// Code generated : DO NOT EDIT.
// Copyright (c) 2022 Jean-Francois SMIGIELSKI
// Distributed under the MIT License

package media2

import (
	"context"
	"github.com/juju/errors"
	"github.com/use-go/onvif"
	"github.com/use-go/onvif/sdk"
	"github.com/use-go/onvif/media2"
)

// Call_GetOSDs forwards the call to dev.CallMethod() then parses the payload of the reply as a GetOSDsResponse.
func Call_GetOSDs(ctx context.Context, dev *onvif.Device, request media2.GetOSDs) (media2.GetOSDsResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			GetOSDsResponse media2.GetOSDsResponse
		}
	}
	var reply Envelope
	if httpReply, err := dev.CallMethod(request); err != nil {
		return reply.Body.GetOSDsResponse, errors.Annotate(err, "call")
	} else {
		err = sdk.ReadAndParse(ctx, httpReply, &reply, "GetOSDs")
		return reply.Body.GetOSDsResponse, errors.Annotate(err, "reply")
	}
}
