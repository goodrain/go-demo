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

package usecase

import (
	"fmt"
	"github.com/goodrain/go-demo/proxy"
	"github.com/sirupsen/logrus"
	"io/ioutil"
	"net/http"
)

type proxyUsecase struct {
}

// NewProxyUsecase returns a proxy.Usecaser
func NewProxyUsecase() proxy.Usecaser {
	return &proxyUsecase{}
}

// Get proxies get action
func (p *proxyUsecase) Get(url string) ([]byte, error) {
	res, err := http.Get(url)
	if err != nil {
		logrus.Errorf("error proxying get action: %v", err)
		return nil, fmt.Errorf("error proxying get action: %v", err)
	}
	result, err := ioutil.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		logrus.Errorf("error read data from response body: %v", err)
		return nil, fmt.Errorf("error read data from response body: %v", err)
	}
	return result, nil
}