package app

import (
	"github.com/jarcoal/httpmock"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestApplication_Search(t *testing.T) {
	type queryValues struct {
		query string
		searchType string
		limit  string
		offset  string
	}
	tests := []struct {
		name             string
		spotifyResponder httpmock.Responder
		queryValues      queryValues
		expectedCode     int
		expectedBody     string
	}{
		{
			name:             "success",
			spotifyResponder: httpmock.NewStringResponder(200, `{ "artists" : { "href" : "https://api.spotify.com/v1/search?query=tania+bowra&type=artist&offset=0&limit=20", "items" : [ { "external_urls" : { "spotify" : "https://open.spotify.com/artist/08td7MxkoHQkXnWAYD8d6Q" }, "followers" : { "href" : null, "total" : 197 }, "genres" : [ ], "href" : "https://api.spotify.com/v1/artists/08td7MxkoHQkXnWAYD8d6Q", "id" : "08td7MxkoHQkXnWAYD8d6Q", "images" : [ { "height" : 640, "url" : "https://i.scdn.co/image/ab67616d0000b2731ae2bdc1378da1b440e1f610", "width" : 640 }, { "height" : 300, "url" : "https://i.scdn.co/image/ab67616d00001e021ae2bdc1378da1b440e1f610", "width" : 300 }, { "height" : 64, "url" : "https://i.scdn.co/image/ab67616d000048511ae2bdc1378da1b440e1f610", "width" : 64 } ], "name" : "Tania Bowra", "popularity" : 1, "type" : "artist", "uri" : "spotify:artist:08td7MxkoHQkXnWAYD8d6Q" } ], "limit" : 20, "next" : null, "offset" : 0, "previous" : null, "total" : 2 } }`),
			queryValues:      queryValues{query: "tania bowra", searchType: "artist"},
			expectedCode:     200,
			expectedBody:     `{ "artists" : { "href" : "https://api.spotify.com/v1/search?query=tania+bowra&type=artist&offset=0&limit=20", "items" : [ { "external_urls" : { "spotify" : "https://open.spotify.com/artist/08td7MxkoHQkXnWAYD8d6Q" }, "followers" : { "href" : null, "total" : 197 }, "genres" : [ ], "href" : "https://api.spotify.com/v1/artists/08td7MxkoHQkXnWAYD8d6Q", "id" : "08td7MxkoHQkXnWAYD8d6Q", "images" : [ { "height" : 640, "url" : "https://i.scdn.co/image/ab67616d0000b2731ae2bdc1378da1b440e1f610", "width" : 640 }, { "height" : 300, "url" : "https://i.scdn.co/image/ab67616d00001e021ae2bdc1378da1b440e1f610", "width" : 300 }, { "height" : 64, "url" : "https://i.scdn.co/image/ab67616d000048511ae2bdc1378da1b440e1f610", "width" : 64 } ], "name" : "Tania Bowra", "popularity" : 1, "type" : "artist", "uri" : "spotify:artist:08td7MxkoHQkXnWAYD8d6Q" } ], "limit" : 20, "next" : null, "offset" : 0, "previous" : null, "total" : 2 } }`,
		}, {
			name:             "missing query param",
			spotifyResponder: nil,
			queryValues:      queryValues{searchType: "artist"},
			expectedCode:     400,
			expectedBody:     `{"description":"Missing 'query' param"}`,
		}, {
			name:             "missing type param",
			spotifyResponder: nil,
			queryValues:      queryValues{query: "tania bowra"},
			expectedCode:     400,
			expectedBody:     `{"description":"Missing 'type' param"}`,
		}, {
			name:             "invalid limit param",
			spotifyResponder: nil,
			queryValues:      queryValues{query: "tania bowra", searchType: "artist", limit: "A"},
			expectedCode:     400,
			expectedBody:     `{"description":"Invalid 'limit' param"}`,
		}, {
			name:             "invalid offset param",
			spotifyResponder: nil,
			queryValues:      queryValues{query: "tania bowra", searchType: "artist", offset: "B"},
			expectedCode:     400,
			expectedBody:     `{"description":"Invalid 'offset' param"}`,
		},
	}

	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	tokenResponder := httpmock.NewStringResponder(200, `{"access_token":"d3a2ded039c2c2851c97d940910dbc3a","token_type": "bearer", "expires_in":3600, "scope": ""}`)
	httpmock.RegisterResponder("POST", "https://accounts.spotify.com/api/token", tokenResponder)

	a := CreateApplication()

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req, err := http.NewRequest("GET", "/search", nil)
			if err != nil {
				t.Errorf("Search(). Error in creating request: %v", err)
			}

			q := req.URL.Query()
			q.Add("query", tt.queryValues.query)
			q.Add("type", tt.queryValues.searchType)
			q.Add("limit", tt.queryValues.limit)
			q.Add("offset", tt.queryValues.offset)
			req.URL.RawQuery = q.Encode()

			w := httptest.NewRecorder()

			httpmock.RegisterResponder("GET", "https://api.spotify.com/v1/search", tt.spotifyResponder)

			handler := http.HandlerFunc(a.Search)
			handler.ServeHTTP(w, req)

			if w.Code != tt.expectedCode {
				t.Errorf("Search() code = %v, expectedCode %v", w.Code, tt.expectedCode)
			}
		})
	}
}