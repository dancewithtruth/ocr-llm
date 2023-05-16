package ocr

import (
	"io"
	"io/ioutil"
	"os"

	"github.com/otiai10/gosseract/v2"
)

type OCR interface {
	ImageToText(file io.ReadCloser) (string, error)
	Close() error
}

type ocr struct {
	client *gosseract.Client
}

func New() OCR {
	client := gosseract.NewClient()
	return &ocr{client: client}
}

// ImageToText uses an image path to load an image and extract the text into a string
func (o *ocr) ImageToText(file io.ReadCloser) (string, error) {
	defer file.Close()
	// Create physical file
	tempfile, err := ioutil.TempFile("", "extract"+"-")
	if err != nil {
		return "", err
	}
	defer func() {
		tempfile.Close()
		os.Remove(tempfile.Name())
	}()

	// Make uploaded physical
	if _, err = io.Copy(tempfile, file); err != nil {
		return "", err
	}

	if err := o.client.SetImage(tempfile.Name()); err != nil {
		return "", err
	}
	text, err := o.client.Text()
	if err != nil {
		return "", err
	}
	return text, nil
}

// Close closes the ocr connection
func (o *ocr) Close() error {
	err := o.client.Close()
	return err
}

type MockOCR struct {
}

// ImageToText uses an image path to load an image and extract the text into a string
func (o *MockOCR) ImageToText(file io.ReadCloser) (string, error) {
	return "", nil
}

// Close closes the ocr connection
func (o *MockOCR) Close() error {
	return nil
}
