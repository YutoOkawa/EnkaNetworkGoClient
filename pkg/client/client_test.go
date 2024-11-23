package client

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestClient_FetchData(t *testing.T) {
	tests := []struct {
		name        string
		handlerFunc func(*testing.T) http.HandlerFunc
		want        []byte
		wantErr     error
	}{
		{
			name: "ShouldFetchDataSuccessfully",
			handlerFunc: func(t *testing.T) http.HandlerFunc {
				return func(w http.ResponseWriter, r *http.Request) {
					w.Write([]byte("Success!"))
				}
			},
			want:    []byte("Success!"),
			wantErr: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			server := httptest.NewServer(tt.handlerFunc(t))

			c := NewClient()
			got, err := c.fetchData(server.URL)
			if err != tt.wantErr {
				t.Errorf("Client.FetchData() error = %v, wantErr %v", err, tt.wantErr)
			}
			if string(got) != string(tt.want) {
				t.Errorf("Client.FetchData() = %v, want %v", got, tt.want)
			}
		})
	}
}
