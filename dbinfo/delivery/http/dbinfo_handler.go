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
	"github.com/goodrain/go-demo/dbinfo"
	"github.com/goodrain/go-demo/model"
	"github.com/labstack/echo"
	"net/http"
)

// DBInfoHandler  represents the http handler for dbinfo
type DBInfoHandler struct {
	DBInfoUcaser dbinfo.Usecaser
}

// NewDBInfoHTTPHandler creates a DBInfoHandler and sets up Echo’s router
func NewDBInfoHTTPHandler(e *echo.Echo, dbinfoUcaser dbinfo.Usecaser) {
	handler := &DBInfoHandler{
		DBInfoUcaser: dbinfoUcaser,
	}
	g := e.Group("/dbinfo")
	g.GET("/ping", handler.Ping)
	g.GET("/list-tables", handler.ListTables)
}

// Ping verifies a connection to the database is still alive,
// establishing a connection if necessary.
func (h *DBInfoHandler) Ping(c echo.Context) error {
	b, err := h.DBInfoUcaser.Ping()
	if err != nil {
		return c.JSON(http.StatusOK, model.NewResponseVO(1, "1001", err.Error(), b))
	}
	return c.JSON(http.StatusOK, model.NewResponseVO(0, "1000", "", b))
}

// ListTables lists tables
func (h *DBInfoHandler) ListTables(c echo.Context) error {
	tables, err := h.DBInfoUcaser.ListTables()
	if err != nil {
		return c.JSON(http.StatusOK, model.NewResponseVO(1, "2001", err.Error(), nil))
	}
	return c.JSON(http.StatusOK, model.NewResponseVO(0, "2000", "", tables))
}
