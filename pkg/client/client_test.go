package client

import (
	"context"
	"fmt"
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
		{
			name: "ShouldReturnErrorWhenServerReturns400",
			handlerFunc: func(t *testing.T) http.HandlerFunc {
				return func(w http.ResponseWriter, r *http.Request) {
					w.WriteHeader(http.StatusBadRequest)
				}
			},
			want:    nil,
			wantErr: fmt.Errorf("Client error: %d", http.StatusBadRequest),
		},
		{
			name: "ShouldReturnErrorWhenServerReturns404",
			handlerFunc: func(t *testing.T) http.HandlerFunc {
				return func(w http.ResponseWriter, r *http.Request) {
					w.WriteHeader(http.StatusNotFound)
				}
			},
			want:    nil,
			wantErr: fmt.Errorf("Client error: %d", http.StatusNotFound),
		},
		{
			name: "ShouldReturnErrorWhenServerReturns424",
			handlerFunc: func(t *testing.T) http.HandlerFunc {
				return func(w http.ResponseWriter, r *http.Request) {
					w.WriteHeader(424)
				}
			},
			want:    nil,
			wantErr: fmt.Errorf("Client error: %d", http.StatusFailedDependency),
		},
		{
			name: "ShouldReturnErrorWhenServerReturns429",
			handlerFunc: func(t *testing.T) http.HandlerFunc {
				return func(w http.ResponseWriter, r *http.Request) {
					w.WriteHeader(http.StatusTooManyRequests)
				}
			},
			want:    nil,
			wantErr: fmt.Errorf("Client error: %d", http.StatusTooManyRequests),
		},
		{
			name: "ShouldReturnErrorWhenServerReturns500",
			handlerFunc: func(t *testing.T) http.HandlerFunc {
				return func(w http.ResponseWriter, r *http.Request) {
					w.WriteHeader(http.StatusInternalServerError)
				}
			},
			want:    nil,
			wantErr: fmt.Errorf("Server error: %d", http.StatusInternalServerError),
		},
		{
			name: "ShouldReturnErrorWhenServerReturns503",
			handlerFunc: func(t *testing.T) http.HandlerFunc {
				return func(w http.ResponseWriter, r *http.Request) {
					w.WriteHeader(http.StatusServiceUnavailable)
				}
			},
			want:    nil,
			wantErr: fmt.Errorf("Server error: %d", http.StatusServiceUnavailable),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			server := httptest.NewServer(tt.handlerFunc(t))
			defer server.Close()

			c := NewClient()
			got, err := c.fetchData(context.Background(), server.URL)
			if err == nil && tt.wantErr != nil {
				t.Errorf("Client.FetchData() error = %v, wantErr %v", err, tt.wantErr)
			}
			if tt.wantErr != nil {
				if err == nil {
					t.Fatal("error is expected but got nil")
				}
				if err.Error() != tt.wantErr.Error() {
					t.Errorf("Client.FetchData() error = %v, wantErr %v", err, tt.wantErr)
				}
			} else {
				if err != nil {
					t.Errorf("error is not expected but got %v", err)
				}
			}
			if string(got) != string(tt.want) {
				t.Errorf("Client.FetchData() = %v, want %v", got, tt.want)
			}
		})
	}
}
