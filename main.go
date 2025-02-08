package main

import (
	"context"
	"fmt"
	"strconv"
	"time"

	// "github.com/aws/aws-sdk-go-v2/config"
	"github.com/omniful/go_commons/config"
	"github.com/omniful/go_commons/db/sql/migration"
	"github.com/omniful/go_commons/http"
	"github.com/omniful/go_commons/log"
	"github.com/varun-singhal-omniful/wms-service/appinit"
	"github.com/varun-singhal-omniful/wms-service/router"
)

const (
	modeMigration  = "migration"
	upMigration    = "up"
	downMigration  = "down"
	forceMigration = "force"
)

func main() {
	// fmt.Println("edfws")
	err := config.Init(10 * time.Second)
	server := http.InitializeServer(":8081", 0, 0, 70)
	context := context.Background()
	runMigration(context, upMigration, "0")
	appinit.Initialize(context)
	err = router.Initialize(context, server)
	if err != nil {
		log.Errorf(err.Error())
		return
	}
	err = server.StartServer("WMS")
	if err != nil {
		fmt.Println("Error in starting the server")
		return
	}

	fmt.Println("Server started")
}

func runMigration(ctx context.Context, migrationType string, number string) {
	database := config.GetString(ctx, "postgresql.database")
	mysqlWriteHost := config.GetString(ctx, "postgresql.master.host")
	mysqlWritePort := config.GetString(ctx, "postgresql.master.port")
	mysqlWritePassword := config.GetString(ctx, "postgresql.master.password")
	mysqlWriterUsername := config.GetString(ctx, "postgresql.master.username")

	m, err := migration.InitializeMigrate("file://deployment/migration", "postgres://"+mysqlWriteHost+":"+mysqlWritePort+"/"+database+"?user="+mysqlWriterUsername+"&password="+mysqlWritePassword+"&sslmode=disable")
	if err != nil {
		panic(err)
	}

	switch migrationType {
	case upMigration:
		err = m.Up()
		if err != nil {
			panic(err)
		}
		break
	case downMigration:
		err = m.Down()
		if err != nil {
			panic(err)
		}
		break
	case forceMigration:
		version, parseErr := strconv.Atoi(number)
		if parseErr != nil {
			panic(parseErr)
		}

		err = m.ForceVersion(version)
		if err != nil {
			return
		}
		break
	default:
		err = m.Up()
		if err != nil {
			panic(err)
		}
		break
	}
}
