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
	"fmt"
	"github.com/goodrain/go-demo/foobar"
	"github.com/goodrain/go-demo/model"
	"github.com/labstack/echo"
	"github.com/sirupsen/logrus"
	"net/http"
)

type foobarHandler struct {
	foobarUcase foobar.Usecaser
}

// NewFoobarHandler creates a new foobarHandler struct
func NewFoobarHandler(e *echo.Echo, foobarUcase foobar.Usecaser) {
	handler := foobarHandler{
		foobarUcase: foobarUcase,
	}

	g := e.Group("/foobar")
	g.GET("/list-env", handler.ListEnv)
	g.POST("/greeting", handler.Greeting)
}

// ListEnv lists all environments
func (h *foobarHandler) ListEnv(c echo.Context) error {
	return c.JSON(http.StatusOK, model.NewResponseVO(0, "3000", "", h.foobarUcase.ListEnv()))
}

// Greeting -
func (h *foobarHandler) Greeting(c echo.Context) error {
	g := model.Greeting{}
	if err := c.Bind(g); err != nil {
		logrus.Errorf("error binding the request body into provided type `model.Greeting`: %v", err)
		return c.JSON(http.StatusInternalServerError, model.NewResponseVO(1, "4001",
			fmt.Sprintf("error binding the request body into provided type `model.Greeting`: %v", err), nil))
	}
	return c.JSON(http.StatusOK, model.NewResponseVO(0, "4000", "", g.Name))
}
