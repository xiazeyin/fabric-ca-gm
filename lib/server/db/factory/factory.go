/*
Copyright IBM Corp. All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/

package factory

import (
	"context"

	"github.com/pkg/errors"
	"github.com/xiazeyin/fabric-ca-gm/lib/server/db"
	"github.com/xiazeyin/fabric-ca-gm/lib/server/db/mysql"
	"github.com/xiazeyin/fabric-ca-gm/lib/server/db/postgres"
	"github.com/xiazeyin/fabric-ca-gm/lib/server/db/sqlite"
	"github.com/xiazeyin/fabric-ca-gm/lib/tls"
	"github.com/xiazeyin/fabric-gm/bccsp"
)

// DB is interface that defines the functions on a database
type DB interface {
	Connect() error
	PingContext(ctx context.Context) error
	Create() (*db.DB, error)
}

// New returns a DB interface for the request database type
func New(
	dbType,
	datasource,
	caName string,
	tlsConfig *tls.ClientTLSConfig,
	csp bccsp.BCCSP,
	metrics *db.Metrics,
) (DB, error) {
	switch dbType {
	case "sqlite3":
		return sqlite.NewDB(datasource, caName, metrics), nil
	case "postgres":
		return postgres.NewDB(datasource, caName, tlsConfig, metrics), nil
	case "mysql":
		return mysql.NewDB(datasource, caName, tlsConfig, csp, metrics), nil
	default:
		return nil, errors.Errorf("Invalid db.type in config file: '%s'; must be 'sqlite3', 'postgres', or 'mysql'", dbType)
	}
}
