package snapmap

type GeoService struct {

}

var API_BASE_URL = "https://ms.sc-jpl.com/web/"

func newGeoService() *GeoService {
	return &GeoService{}
}

//func (g *GeoService)