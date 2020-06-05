package snapmap

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

type PlaylistService struct {
	http *http.Client
}

func newPlaylistService(http *http.Client) *PlaylistService {
	return &PlaylistService{
		http: http,
	}
}

type GeoPoint struct {
	Latitude float32 `json:"lat"`
	Longitude float32 `json:"lon"`
}

type PlaylistRequestBody struct {
	RequestGeoPoint GeoPoint `json:"requestGeoPoint"`
	ZoomLevel int `json:"zoomLevel"`
	TileSetId TileSetId `json:"tileSetId"`
	RadiusMeters int `json:"radiusMeters"`
}

type PlaylistResponse struct {
	Manifest struct {
		Elements []SnapElement `json:"elements"`
	} `json:"manifest"`
}

type SnapElement struct {
	Id string `json:"id"`
	Duration float32 `json:"duration"`
	Timestamp string `json:"timestamp"`
	SnapInfo SnapInfo `json:"snapInfo"`
}

type SnapInfo struct {
	SnapMediaTitle string `json:"snapMediaTitle"`
	Title struct {
		Fallback string `json:"fallback"`
	} `json:"title"`
	StreamingMediaInfo struct {
		PrefixUrl string `json:"prefixUrl"`
		MediaUrl string `json:"mediaUrl"`
	} `json:"streamingMediaInfo"`
}

func (p *PlaylistService) GetPlaylist(latitude float32, longitude float32, radiusInMeters int, zoomLevel int, id TileSetId) (*PlaylistResponse, error) {
	requestBody := PlaylistRequestBody{
		RequestGeoPoint: GeoPoint{
			Latitude:  latitude,
			Longitude: longitude,
		},
		ZoomLevel: zoomLevel,
		TileSetId: id,
		RadiusMeters: radiusInMeters,
	}

	buf := new(bytes.Buffer)
	json.NewEncoder(buf).Encode(requestBody)
	response, err := p.http.Post(fmt.Sprintf("%v/getPlaylist", SnapmapApiBaseUrl), "application/json", buf)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	//json.NewEncoder(buf).Encode(requestBody)
	//body, err := ioutil.ReadAll(buf)

	var responseBody PlaylistResponse
	err = json.NewDecoder(response.Body).Decode(&responseBody)
	if err != nil {
		return nil, err
	}

	return &responseBody, nil
}