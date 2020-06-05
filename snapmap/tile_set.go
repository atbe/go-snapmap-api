package snapmap

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

type TileSetService struct {
	http *http.Client
}

func newTileSetService(http *http.Client) *TileSetService {
	return &TileSetService{
		http: http,
	}
}

type TileSetId struct {
	Type string `json:"type"`
	Flavor string `json:"flavor"`
	Epoch string `json:"epoch"`
}

type TileSetInfo struct {
	Id TileSetId `json:"id"`
	State string `json:"state"`
	StartTime int `json:"startTime"`
	LastUpdateTime int `json:"lastUpdateTime"`
	PoiTileSetInfo interface{} `json:"poiTileSetInfo"`
}

type Response struct {
	RenderConfig struct {
		HeatmapGradient  interface{} `json:"heatmapGradient"`
		HeatNormalizationPeak int `json:"heatNormalizationPeak"`
		FuzzNormalizationPeak float32 `json:"fuzzNormalizationPeak"`
		HeatPointBaseRadius int `json:"heatPointBaseRadius"`
	} `json:"renderConfig"`

	TileSetInfos []TileSetInfo `json:"tileSetInfos"`
}

func (t *TileSetService) GetLatest(target *Response) (error) {
	response, err := t.http.Post(fmt.Sprintf("%v/getLatestTileSet", SnapmapApiBaseUrl), "application/json", strings.NewReader("{}"))
	if err != nil {
		return err
	}
	defer response.Body.Close()

	return json.NewDecoder(response.Body).Decode(target)
}