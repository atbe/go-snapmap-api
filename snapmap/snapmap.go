package snapmap

import (
	"github.com/dghubble/sling"
	"net/http"
)

const SnapmapApiBaseUrl = "https://ms.sc-jpl.com/web"

type Client struct {
	sling *sling.Sling
	EpochService *EpochService
	PlaylistService *PlaylistService
	TileSetService *TileSetService
}

func NewClient(client *http.Client) *Client {
	base :=  sling.New().Client(client).Base(SnapmapApiBaseUrl)
	base = base.Set("Content-Type", "application/json")
	return &Client{
		sling: base,
		EpochService: newEpochService(client),
		PlaylistService: newPlaylistService(client),
		TileSetService: newTileSetService(client),
	}
}

