package gomodometer

import (
	"os"
	"testing"
)

const testFileName = "/tmp/gomodometer_test_file"

func TestNewMouseOdometer(t *testing.T) {
	createTestFile()
	o := NewMouseOdometer(testFileName)
	if o == nil || o.deviceFile == nil {
		t.Fail()
	}
	os.Remove(testFileName)
}

func TestNormalizeReading(t *testing.T) {
	var raw byte = 200
	result := normalizeReading(raw)
	if result != 72 {
		t.Fail()
	}
}

func TestConvertToCentimeters(t *testing.T) {
	var raw int = 100
	result := convertToCentimeters(raw)
	if result > 0.03 || result < 0.02 { // should be 0.0254, epsilon = .01
		t.Fail()
	}
}

func createTestFile() {
	os.Create(testFileName)
	f, _ := os.Open(testFileName)
	defer f.Close()
	f.WriteString("111222333")
}
