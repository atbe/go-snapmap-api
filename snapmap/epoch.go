package snapmap

import (
	"net/http"
)

type EpochService struct {
	http *http.Client
}

func newEpochService(http *http.Client) *EpochService {
	return &EpochService{
		http: http,
	}
}

type TileSetType string

const (
	HeatTileSetType TileSetType = "HEAT"
	PointOfInterestTileSetType TileSetType = "POI"
)

type EpochResponse struct {

}

func (e *EpochService) GetEpoch(target interface{}, setType TileSetType) error {
	return nil
}
