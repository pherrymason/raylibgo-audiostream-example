package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
	"math"
)

const sampleRate = 44100
const streamBufferSize = 1024

type Oscillator struct {
	phase       float64
	phaseStride float64
}

func (o *Oscillator) advance() {
	o.phase += o.phaseStride
	if o.phase >= 1 {
		o.phase -= 1
	}
}

func NewOscillator(frequency float32, sampleDuration float32) *Oscillator {
	return &Oscillator{
		phase:       0,
		phaseStride: float64(frequency * sampleDuration),
	}
}

func main() {
	rl.InitAudioDevice()

	stream := rl.LoadAudioStream(
		uint32(sampleRate),
		32,
		1,
	)
	defer rl.UnloadAudioStream(stream)
	defer rl.CloseAudioDevice()

	rl.SetAudioStreamVolume(stream, 0.10)
	rl.SetAudioStreamBufferSizeDefault(streamBufferSize)
	rl.PlayAudioStream(stream)

	samples := make([]float32, streamBufferSize)
	sampleDuration := float32(1) / sampleRate
	oscillator := NewOscillator(440, sampleDuration)

	for {
		if rl.IsAudioStreamProcessed(stream) {
			updateSamples(samples, oscillator)
			rl.UpdateAudioStream(
				stream,
				samples,
				streamBufferSize,
			)
		}
	}
}

func updateSamples(samples []float32, oscillator *Oscillator) {
	for t := 0; t < streamBufferSize; t++ {
		oscillator.advance()
		x := 2 * math.Pi * oscillator.phase
		samples[t] = float32(math.Sin(float64(x)))
	}
}
