/*
	@author: Robert Pratt <robert.pratt[at]homelab.farm>
	@package: drawings
	@description: Struct definitions for /api/v1/draw-results and children.
	@created: 2020-12-05
*/
package drawings

import (
	"../client"
	"net/http"
	"net/url"
)

func CreateDrawingRequest(date string, game Game) (req *http.Request, err error) {
	params := url.Values{}
	params.Set("draw_date", date)

	reqUrl, err := url.Parse(client.ApiHost + drawResultsUrlPath + "/" + string(game))
	if err != nil {
		return
	}

	reqUrl.RawQuery = params.Encode()

	req = &http.Request{
		URL:    reqUrl,
		Method: "GET",
		Form:   params,
	}
	return
}
