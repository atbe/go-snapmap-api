package snapmap

import (
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

	log.Print(len(responseBody.TileSetInfos))
	log.Print(responseBody.RenderConfig)
	log.Print(err)
	tileSetInfo := responseBody.TileSetInfos[0]
	log.Print(tileSetInfo)
}