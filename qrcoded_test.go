package main

import (
	"bytes"
	"image/png"
	"testing"
)

func TestGenerateQRCodeReturnsValue(t *testing.T) {
	buffer := new(bytes.Buffer)
	GenerateQRCode(buffer, "555-2368")

	if buffer.Len() == 0 {
		t.Errorf("Generated QR code has no data")
	}
}

func TestGenerateQrCodeGeneratesPNG(t *testing.T) {
	buffer := new(bytes.Buffer)
	GenerateQRCode(buffer, "555-2368")
	_, err := png.Decode(buffer)

	if err != nil {
		t.Errorf("Generated QR code is not a PNG: %s", err)
	}
}
