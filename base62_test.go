package base62

import (
	"github.com/bmizerany/assert"
	"testing"
)

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
}
