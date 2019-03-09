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

package repository

import (
	"database/sql"
	"fmt"
	"github.com/goodrain/go-demo/dbinfo"
	"github.com/sirupsen/logrus"
)

type mysqlDBInfoRepo struct {
	DB *sql.DB
}

// NewMysqlDBInfoRepository will create an implementation of author.Repositorier
func NewMysqlDBInfoRepository(db *sql.DB) dbinfo.Repositorier {
	return &mysqlDBInfoRepo{
		DB: db,
	}
}

// Ping verifies a connection to the database is still alive,
// establishing a connection if necessary.
func (m *mysqlDBInfoRepo) Ping() (bool, error) {
	if m.DB == nil {
		return false, fmt.Errorf("*sql.DB is nil")
	}

	err := m.DB.Ping()
	if err != nil {
		logrus.Debugf("error pinging sql.DB: %v", err)
		return false, fmt.Errorf("error pinging sql.DB: %v", err)
	}
	return true, nil
}

// ListTables lists tables
func (m *mysqlDBInfoRepo) ListTables() ([]string, error) {
	query := `show tables;`
	rows, err := m.DB.Query(query)
	if err != nil {
		logrus.Error(err)
		return nil, err
	}
	defer rows.Close()

	// Get column names
	columns, err := rows.Columns()
	if err != nil {
		logrus.Errorf("error getting columns: %v", err)
		return nil, err
	}

	// Make a slice for the values
	values := make([]sql.RawBytes, len(columns))

	// rows.Scan wants '[]interface{}' as an argument, so we must copy the
	// references into such a slice
	// See http://code.google.com/p/go-wiki/wiki/InterfaceSlice for details
	scanArgs := make([]interface{}, len(values))
	for i := range values {
		scanArgs[i] = &values[i]
	}

	var result []string
	// Fetch rows
	for rows.Next() {
		// get RawBytes from data
		err = rows.Scan(scanArgs...)
		if err != nil {
			logrus.Errorf("error getting RawBytes from data: %v", err)
		}

		// Now do something with the data.
		// Here we just print each column as a string.
		var value string
		for i, col := range values {
			// Here we can check if the value is nil (NULL value)
			if col == nil {
				value = "NULL"
			} else {
				value = string(col)
			}
			logrus.Debugf("%s:%s\n", columns[i], value)
			result = append(result, value)
		}
	}
	if err = rows.Err(); err != nil {
		logrus.Debugf("error listing tables: %v", err)
		return nil, err
	}
	return result, nil
}