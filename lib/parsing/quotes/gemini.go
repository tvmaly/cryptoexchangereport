package quotes

import "strconv"

// GeminiQuote is not an ideal structure as they support only 3 pairs of coin/currency
// and they return the symbol as a key in the JSON response
// this requires us to represent the volume as a map rather than a struct
// to abstract the representation of the data to a single form
type GeminiQuote struct {
	Ask    float64           `json:"ask,string"`
	Bid    float64           `json:"bid,string"`
	Last   float64           `json:"last,string"`
	Volume map[string]string `json:"volume"`
}

func (t *GeminiQuote) Timestamp() float64 {

	for k, v := range t.Volume {

		if k == "timestamp" {

			ts, err := strconv.ParseFloat(v, 64)

			if err != nil {
				return 0
			}

			return ts

		}

	}

	return 0

}
