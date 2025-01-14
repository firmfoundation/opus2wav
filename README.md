# what is it?

opus2wav

opus2wav is a Go package designed for WebRTC projects, enabling seamless conversion of RTP audio streams into WAV format. This package decodes Opus-encoded audio payloads from RTP streams and outputs high-quality WAV files, making it ideal for integration with external AI models and speech recognition systems.
Key Features

    RTP Stream Handling: Receives and processes RTP audio streams from WebRTC clients.
    Opus Decoding: Extracts and decodes Opus-encoded audio payloads.
    WAV Conversion: Outputs audio in widely supported WAV format, suitable for AI models like Google Speech-to-Text, AWS Transcribe, and more.
    Ease of Integration: Designed to integrate seamlessly into real-time and batch audio processing pipelines.

Use Cases

    Preparing audio for speech-to-text conversion.
    Processing live audio streams in WebRTC applications.
    Converting Opus audio to WAV for archiving or further analysis.