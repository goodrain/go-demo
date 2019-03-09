// RAINBOND, Application Management Platform
// Copyright (C) 2014-2017 Goodrain Co., Ltd.

// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version. For any non-GPL usage of Rainbond,
// one or multiple Commercial Licenses authorized by Goodrain Co., Ltd.
// must be obtained first.

// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU General Public License for more details.

// You should have received a copy of the GNU General Public License
// along with this program. If not, see <http://www.gnu.org/licenses/>.

package http

import (
	"bytes"
	"fmt"
	"github.com/goodrain/go-demo/proxy"
	"github.com/labstack/echo"
	"github.com/sirupsen/logrus"
	"io/ioutil"
	"net/http"
)

type ProxyHandler struct {
	ProxyUcaser proxy.Usecaser
}

// NewProxyHandler creates a ProxyHandler and sets up Echo’s router
func NewProxyHandler(e *echo.Echo, proxyUcaser proxy.Usecaser) {
	handler := &ProxyHandler{
		ProxyUcaser: proxyUcaser,
	}
	e.POST("/proxy", handler.Proxy)
}

type ProxyInfo struct {
	ProxyMethod string      `json:"proxy_method"`
	ProxyURL    string      `json:"proxy_url"`
	Data        interface{} `json:"data"`
}

func (p *ProxyHandler) Proxy(c echo.Context) error {
	// we need to buffer the body if we want to read it here and send it
	// in the request.
	body, err := ioutil.ReadAll(c.Request().Body)
	if err != nil {
		logrus.Errorf("error reading data from request body: %v", err)
		return c.JSON(http.StatusInternalServerError,
			fmt.Errorf("error reading data from request body: %v", err))
	}

	// you can reassign the body if you need to parse it as multipart
	c.Request().Body = ioutil.NopCloser(bytes.NewReader(body))

	pi := new(ProxyInfo)
	if err = c.Bind(pi); err != nil {
		logrus.Errorf("error binding the request body into ProxyInfo: %v", err)
		return c.JSON(http.StatusInternalServerError,
			fmt.Errorf("error binding the request body into ProxyInfo: %v", err))
	}

	proxyReq, err := http.NewRequest(pi.ProxyMethod, pi.ProxyURL, bytes.NewReader(body))

	// We may want to filter some headers, otherwise we could just use a shallow copy
	// proxyReq.Header = req.Header
	proxyReq.Header = make(http.Header)
	for h, val := range c.Request().Header {
		proxyReq.Header[h] = val
	}

	httpCli := http.Client{}
	resp, err := httpCli.Do(proxyReq)
	if resp != nil {
		defer resp.Body.Close()
	}
	if err != nil {
		logrus.Errorf("error doing http proxy: %v", err)
		return c.JSON(http.StatusBadGateway, fmt.Sprintf("error doing http proxy: %v", err))
	}
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		logrus.Errorf("error reading data from response body")
		return c.JSON(http.StatusInternalServerError, fmt.Errorf("error reading data from response body"))
	}

	return c.String(http.StatusOK, string(b))
}
