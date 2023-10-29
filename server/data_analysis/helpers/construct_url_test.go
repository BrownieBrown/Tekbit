package helpers_test

import (
	"server/data_analysis/helpers"
	"testing"
)

func TestConstructAlphaVantageURL(t *testing.T) {
	type args struct {
		params map[string]string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Test constructing Alpha Vantage URL",
			args: args{
				params: map[string]string{
					"function": "TIME_SERIES_INTRADAY",
					"symbol":   "IBM",
					"interval": "5min",
					"apikey":   "fake_key",
				},
			},
			// Note: The order of query parameters may vary, but that's okay in URLs.
			want: "https://www.alphavantage.co/query?apikey=fake_key&function=TIME_SERIES_INTRADAY&interval=5min&symbol=IBM",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := helpers.ConstructAlphaVantageURL(tt.args.params); got != tt.want {
				t.Errorf("ConstructAlphaVantageURL() = %v, want %v", got, tt.want)
			}
		})
	}
}
