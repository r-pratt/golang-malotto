/*
	@author: Robert Pratt <robert.pratt[at]homelab.farm>
	@package: client
	@description: HTTP client wrapper for MassLotto API
	@created: 2020-12-05
*/
package client

import (
	"net/http"
	"time"
)

func CreateLottoClient() (client *http.Client) {
	tr := &http.Transport{
		IdleConnTimeout: 30 * time.Second,
	}
	client = &http.Client{
		Transport: tr,
	}
	return
}
