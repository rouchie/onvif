package media2

import (
	"github.com/use-go/onvif/xsd/onvif"
)

type GetStreamUri struct {
	XMLName      string               `xml:"tr2:GetStreamUri"`
	Protocol     string               `xml:"tr2:Protocol"`
	ProfileToken onvif.ReferenceToken `xml:"tr2:ProfileToken"`
}

type GetStreamUriResponse struct {
	Uri string
}

type GetOSDs struct {
	XMLName            string               `xml:"tr2:GetOSDs"`
	OSDToken           onvif.ReferenceToken `xml:"tr2:OSDToken"`
	ConfigurationToken onvif.ReferenceToken `xml:"tr2:ConfigurationToken"`
}

type GetOSDsResponse struct {
	OSDs []onvif.OSDConfiguration
}

type CreateOSD struct {
	XMLName string                 `xml:"tr2:CreateOSD"`
	OSD     onvif.OSDConfiguration `xml:"tr2:OSD"`
}

type CreateOSDResponse struct {
	OSDToken onvif.ReferenceToken
}
