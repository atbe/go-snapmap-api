package snapmap

import "github.com/dghubble/sling"

type PlaylistService struct {
	sling *sling.Sling
}

func newPlaylistService(sling *sling.Sling) *PlaylistService {
	return &PlaylistService{
		sling: sling,
	}
}

func (p *PlaylistService) GetPlaylist(latitude float32, longitude float32, radiusInMeters int, zoomLevel int) interface{} {
	// get epoch

	// do the post request

	return nil
}