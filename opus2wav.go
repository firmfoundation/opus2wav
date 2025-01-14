package opus2wav

type Converter struct {
	// Internal fields, e.g., RTP stream settings or codec state
}

// NewConverter initializes a new Converter instance.
func NewConverter() *Converter {
	return &Converter{}
}

// ProcessRTP processes an RTP packet and decodes Opus audio.
func (c *Converter) ProcessRTP(packet []byte) error {
	// Implementation
	return nil
}

// WriteWAV writes decoded audio to a WAV file.
func (c *Converter) WriteWAV(filename string) error {
	// Implementation

	return nil
}
