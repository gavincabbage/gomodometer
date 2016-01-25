package gomodometer

import (
	"errors"
	"os"
)

const conversionCoefficient float64 = 0.000254

type MouseOdometer struct {
	deviceFile *os.File
	readings   chan []byte
	quit       chan struct{}
	x, y       int
}

func NewMouseOdometer(deviceFileName string) *MouseOdometer {
	f, err := os.Open(deviceFileName)
	if err != nil {
		return nil
	}
	return &MouseOdometer{deviceFile: f}
}

func (o *MouseOdometer) Start() {
	o.setup()
	go o.readMouseDevice()
	go o.sumReadingsUntilQuit()
}

func (o *MouseOdometer) Stop() (float64, float64) {
	if o.quit != nil {
		o.quit <- struct{}{}
	}
	return convertToCentimeters(o.x), convertToCentimeters(o.y)
}

func (o *MouseOdometer) setup() {
	o.x, o.y = 0.0, 0.0
	o.readings = make(chan []byte)
	o.quit = make(chan struct{})
}

func (o *MouseOdometer) readMouseDevice() {
	buffer := make([]byte, 3)
	for {
		num, err := o.deviceFile.Read(buffer)
		if err != nil || num != 3 {
			panic(errors.New("problem reading file"))
		}
		o.readings <- buffer
	}
}

func (o *MouseOdometer) sumReadingsUntilQuit() {
	for {
		select {
		case <-o.quit:
			return
		case reading := <-o.readings:
			o.x += normalizeReading(reading[1])
			o.y += normalizeReading(reading[2])
			//fmt.Println(convertToCentimeters(o.x), " ", convertToCentimeters(o.y)) // gnuplot> plot 'mouse.dat' u 1:2 w l
		}
	}
}

func normalizeReading(raw byte) int {
	return int(raw) - 128
}

func convertToCentimeters(raw int) float64 {
	return float64(raw) * conversionCoefficient
}
