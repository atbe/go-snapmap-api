package snapmap

import (
	"fmt"
	"log"
	"net/http"
	"testing"
)

func TestIntMinBasic(t *testing.T) {
	httpClient := http.Client{}
	apiClient := NewClient(&httpClient)

	//41.856922, -87.644096

	responseBody := new(Response)
	err := apiClient.TileSetService.GetLatest(responseBody)

	if err != nil {
		panic(err)
	}


	// get a tile set info object
	log.Print(len(responseBody.TileSetInfos))
	log.Print(responseBody.RenderConfig)
	log.Print(err)
	tileSetInfo := responseBody.TileSetInfos[1]
	log.Print(tileSetInfo)

	// get the epoch object

	log.Print("Getting response")
	playlistResponse, err := apiClient.PlaylistService.GetPlaylist(41.856922, -87.644096, 10000, 12, tileSetInfo.Id)
	if err != nil {
		panic(err)
	}

	for _, snapInfo := range playlistResponse.Manifest.Elements {
		log.Print(fmt.Sprintf("%v%v", snapInfo.SnapInfo.StreamingMediaInfo.PrefixUrl, snapInfo.SnapInfo.StreamingMediaInfo.MediaUrl))
	}

}