package states

import (
	"errors"
	"fmt"

	"github.com/davecgh/go-spew/spew"
	"github.com/sirupsen/logrus"
)

type BrainwaveSignal struct {
	AttentionIntensity int
	RawData            []byte
}

type BrainwaveSignalCollector struct {
	DeviceID string
	Signal   BrainwaveSignal
}

func (c *BrainwaveSignalCollector) CollectSignal() error {
	logrus.Info("Collecting brainwave signal...")

	// Simulate signal collection process
	c.Signal = BrainwaveSignal{
		AttentionIntensity: 3,
		RawData:            []byte{1, 2, 3, 4, 5},
	}

	logrus.Info("Collected brainwave signal:")
	spew.Dump(c.Signal)

	return nil
}

func (c *BrainwaveSignalCollector) TransmitSignal() error {
	if c.Signal.RawData == nil {
		return errors.New("Cannot transmit nil signal")
	}

	logrus.Info("Transmitting brainwave signal...")

	// Simulate signal transmission process
	fmt.Println("Transmitting signal: ", c.Signal.RawData)

	return nil
}
