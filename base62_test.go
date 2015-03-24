package base62

import (
	u "github.com/araddon/gou"
	"github.com/bmizerany/assert"
	"testing"
)

func init() {
	u.SetupLogging("debug")
	u.SetColorOutput()
}

func TestBase62(t *testing.T) {
	urlVal := "http://www.lytics.io/?utm_content=content&utm_campaign=campaign"
	encodedValue := StdEncoding.EncodeToString([]byte(urlVal))
	assert.Tf(t, encodedValue != "", "%v", encodedValue)
	rtVal, err := StdEncoding.DecodeString(encodedValue)
	assert.T(t, err == nil)
	assert.Tf(t, string(rtVal) == urlVal, "%v", encodedValue)

	intVal := "999"
	encodedValue = StdEncoding.EncodeToString([]byte(intVal))
	assert.Tf(t, encodedValue == "OTk4", "%v", encodedValue)
	rtVal, err = StdEncoding.DecodeString(encodedValue)
	assert.T(t, err == nil)
	assert.Tf(t, string(rtVal) == intVal, "%v", encodedValue)

	urlVal = `http://our-uploads.s3.amazonaws.com/file-export/stuff-1427217700-12.csv?AWSAccessKeyId=AAAAIIG7MAQRTHTD7CLP&Expires=1427304113&Signature=VQsRAhgamiw1RVtbrCXOsMu%2BgFo`
	encodedValue = StdEncoding.EncodeToString([]byte(urlVal))
	// We are just ensuring it doesn't panic
	assert.T(t, true == true)
}
