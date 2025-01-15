package opus2wav

import (
	"errors"

	"github.com/hraban/opus"
	"github.com/pion/rtp"
)

type Converter struct {
	decoder *opus.Decoder
}

// NewConverter initializes a new Converter instance.
func NewConverter() *Converter {
	return &Converter{}
}

func (c *Converter) ProcessRTP(rtpPacket *rtp.Packet) ([]int16, error) {
	if rtpPacket == nil {
		return nil, errors.New("received nil RTP packet")
	}

	if len(rtpPacket.Payload) == 0 {
		return nil, errors.New("received RTP packet has no payload")
	}

	pcm := make([]int16, 1920*2) //buffer
	n, err := c.decoder.Decode(rtpPacket.Payload, pcm)
	if err != nil {
		return nil, err
	}

	return pcm[:n], nil
}

// WriteWAV writes decoded audio to a WAV file.
func (c *Converter) WriteWAV(filename string) error {
	// Implementation

	return nil
}
