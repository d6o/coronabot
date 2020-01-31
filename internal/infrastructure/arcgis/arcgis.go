package arcgis

import (
	"encoding/json"
	"net/http"

	"github.com/disiqueira/coronabot/internal/domain/model"
)

type (
	CoranaREST struct {
		httpClient *http.Client
	}

	arcgisReponse struct {
		Features []struct {
			Attributes struct {
				Confirmed     int    `json:"Confirmed"`
				CountryRegion string `json:"Country_Region"`
				Deaths        int    `json:"Deaths"`
				LastUpdate    int64  `json:"Last_Update"`
				Recovered     int    `json:"Recovered"`
			} `json:"attributes"`
		} `json:"features"`
	}
)

const (
	arcgisURL = "https://services1.arcgis.com/0MSEUqKaxRlEPj5g/arcgis/rest/services/ncov_cases/FeatureServer/2/query?f=json&where=1%3D1&returnGeometry=false&spatialRel=esriSpatialRelIntersects&outFields=*&orderByFields=Confirmed%20desc&resultOffset=0&resultRecordCount=250&cacheHint=true"
)

func New(httpClient *http.Client) *CoranaREST {
	return &CoranaREST{
		httpClient: httpClient,
	}
}

func (c *CoranaREST) StatusPerCountry() ([]model.Status, error) {
	req, err := http.NewRequest("GET", arcgisURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}

	var response arcgisReponse
	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		return nil, err
	}

	statusList := make([]model.Status, len(response.Features))
	for i, feature := range response.Features {
		statusList[i] = model.NewStatus(
			feature.Attributes.CountryRegion,
			feature.Attributes.Deaths,
			feature.Attributes.Confirmed,
			feature.Attributes.Recovered,
			feature.Attributes.LastUpdate,
		)
	}

	return statusList, nil
}
