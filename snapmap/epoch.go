package snapmap

import "github.com/dghubble/sling"

type EpochService struct {
	sling *sling.Sling
}

func newEpochService(sling *sling.Sling) *EpochService {
	return &EpochService{
		sling: sling,
	}
}

type TileSetType string

const (
	HeatTileSetType TileSetType = "HEAT"
	PointOfInterestTileSetType TileSetType = "POI"
)

func (e *EpochService) GetEpoch(setType TileSetType) interface{} {
	return nil
}
