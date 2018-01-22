package base62

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestBase62(t *testing.T) {
	urlVal := "http://www.lytics.io/?utm_content=content&utm_campaign=campaign"
	encodedValue := StdEncoding.EncodeToString([]byte(urlVal))
	assert.NotEqual(t, "", encodedValue)
	rtVal, err := StdEncoding.DecodeString(encodedValue)
	assert.Equal(t, nil, err)
	assert.Equal(t, string(rtVal), urlVal)

	intVal := "999"
	encodedValue = StdEncoding.EncodeToString([]byte(intVal))
	assert.Equal(t, "OTk4", encodedValue, "%v", encodedValue)
	rtVal, err = StdEncoding.DecodeString(encodedValue)
	assert.Equal(t, nil, err)
	assert.Equal(t, string(rtVal), intVal)

	urlVal = `http://our-uploads.s3.amazonaws.com/file-export/stuff-1427217700-12.csv?AWSAccessKeyId=AAAAIIG7MAQRTHTD7CLP&Expires=1427304113&Signature=VQsRAhgamiw1RVtbrCXOsMu%2BgFo`
	encodedValue = StdEncoding.EncodeToString([]byte(urlVal))
	// We are just ensuring it doesn't panic
	assert.True(t, len(encodedValue) > 0)
}
