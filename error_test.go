package getstream_test

import (
	"encoding/json"
	"testing"
	"time"

	getstream "github.com/GetStream/stream-go"
	"github.com/stretchr/testify/assert"
)

func TestError(t *testing.T) {

	errorResponse := "{\"code\": 5, \"detail\": \"some detail\", \"duration\": \"36ms\", \"exception\": \"an exception\", \"status_code\": 400}"

	var getStreamError getstream.Error
	err := json.Unmarshal([]byte(errorResponse), &getStreamError)
	if err != nil {
		t.Fatal(err)
	}

	testError := getstream.Error{
		Code:        5,
		Detail:      "some detail",
		RawDuration: "36ms",
		Exception:   "an exception",
		StatusCode:  400,
	}

	assert.Equal(t, getStreamError, testError)

	assert.Equal(t, getStreamError.Duration(), time.Millisecond*36)

	assert.Equal(t, getStreamError.Error(), "an exception (36ms): some detail")
}

func TestErrorBadDuration(t *testing.T) {

	testError := getstream.Error{
		Code:        5,
		Detail:      "some detail",
		RawDuration: "36blah",
		Exception:   "an exception",
		StatusCode:  400,
	}

	if testError.Duration() != time.Duration(0) {
		t.Fatal()
	}

}
