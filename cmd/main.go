package main

import (
	"log"

	"github.com/go-sql-driver/mysql"
	"github.com/jumystap/jumystap-core/cmd/jumystap"
	"github.com/jumystap/jumystap-core/config"
	"github.com/jumystap/jumystap-core/database"
)

func main() {
    db, err := database.NewMySQLStorage(mysql.Config{
        User:                   config.Envs.DBUser,
        Passwd:                 config.Envs.DBPassword,
        Addr:                   config.Envs.DBAddress,
        DBName:                 config.Envs.DBName,

        Net:                    "tcp",
        AllowNativePasswords:   true,
        ParseTime:              true,
    })
    if err != nil {
        log.Fatal(err)
    }

    server := jumystap.NewAPIServer(":8090", db)

	if err := server.Run(); err != nil {
		log.Fatal(err)
	}    
}
