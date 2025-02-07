package appinit

import (
	"context"
	"fmt"
	"time"

	"github.com/omniful/go_commons/config"
	opostgres "github.com/omniful/go_commons/db/sql/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Initialize(ctx context.Context) {
	initializeDB(ctx)
}
func initializeDB(ctx context.Context) {
	maxOpenConnections := config.GetInt(ctx, "postgresql.maxOpenConns")
	maxIdleConnections := config.GetInt(ctx, "postgresql.maxIdleConns")
	database := config.GetString(ctx, "postgresql.database")
	connIdleTimeout := 10 * time.Minute
	mysqlWriteServer := config.GetString(ctx, "postgresql.master.host")
	mysqlWritePort := config.GetString(ctx, "postgresql.master.port")
	mysqlWritePassword := config.GetString(ctx, "postgresql.master.password")
	mysqlWriterUsername := config.GetString(ctx, "postgresql.master.username")

	debugMode := config.GetBool(ctx, "postgresql.debugMode")

	// Master config i.e. - Write endpoint
	masterConfig := opostgres.DBConfig{
		Host:               mysqlWriteServer,
		Port:               mysqlWritePort,
		Username:           mysqlWriterUsername,
		Password:           mysqlWritePassword,
		Dbname:             database,
		MaxOpenConnections: maxOpenConnections,
		MaxIdleConnections: maxIdleConnections,
		ConnMaxLifetime:    connIdleTimeout,
		DebugMode:          debugMode,
	}

	slavesConfig := make([]opostgres.DBConfig, 0)
	db := opostgres.InitializeDBInstance(masterConfig, &slavesConfig)
	DB = db.GetMasterDB(ctx)
	fmt.Println("db connected")
	// log.InfofWithContext(ctx, "Initialized Postgres DB client")
	// postgres.SetCluster(db)
}
