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

import "github.com/goodrain/go-demo/dbinfo"

type dbinfoUsecase struct {
	dbinfoRepo dbinfo.Repositorier
}

// NewDBInfoUsecase returns a dbinfo.Usecaser
func NewDBInfoUsecase(dbinfoRepo dbinfo.Repositorier) dbinfo.Usecaser {
	return &dbinfoUsecase{
		dbinfoRepo: dbinfoRepo,
	}
}

// Ping verifies a connection to the database is still alive,
// establishing a connection if necessary.
func (d *dbinfoUsecase) Ping() (bool, error) {
	return d.dbinfoRepo.Ping()
}

// ListTables lists tables
func (d *dbinfoUsecase) ListTables() ([]string, error) {
	return d.dbinfoRepo.ListTables()
}
