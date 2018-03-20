package exchangedata

import (
	"fmt"
	"gopkg.in/resty.v1"
	"time"
)

type RequestSender interface {
	SendRequest(string) (Response, error)
}

type Response struct {
	ReceivedTimestamp float64 `json:receivedtimestamp`
	SentTimestamp     float64 `json:senttimestamp`
	Response          []byte  `json:response`
}

type RestyAdapter string

func NewRequestSender() RequestSender {
	var ra RestyAdapter = "resty"
	return ra
}

func (ra RestyAdapter) SendRequest(url string) (Response, error) {
	restyResp, err := resty.R().Get(url)

	if err != nil {
		fmt.Printf("Error while sending GET request to %s", url)
		return Response{}, err
	}

	return ra.WrapRestyResponse(restyResp), nil
}

// Wraps gopkg.in/resty.v1.Responce type in our Response type
// with received and sent timestamp as UTC Unix timestamp with milliseconds,
// and response body
func (ra RestyAdapter) WrapRestyResponse(resp *resty.Response) Response {
	var requested float64 = UnixNanoToUnixMilli(resp.Request.Time.UTC().UnixNano())
	var received float64 = UnixNanoToUnixMilli(resp.ReceivedAt().UTC().UnixNano())

	var wrappedResp Response = Response{
		ReceivedTimestamp: received,
		SentTimestamp:     requested,
		Response:          resp.Body(),
	}

	return wrappedResp
}

func UnixNanoToUnixMilli(unixnano int64) float64 {
	return float64(unixnano/int64(time.Millisecond)) / 1000.0
}
