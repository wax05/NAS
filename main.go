package main

import (
	"fmt"

	DB "github.com/wax05/NAS/Database"

	"database/sql"

	_ "github.com/go-sql-driver/mysql"
	//"github.com/gin-gonic/gin"
	//"net/http"
)

func main() {
	Db, DbErr := sql.Open("mysql", "root:0000@tcp(localhost:3306)/nas")
	if DbErr != nil {
		panic(DbErr.Error())
	}
	defer Db.Close()
	DbVal, err := DB.AllUserData(*Db)
	if err != nil {
		panic(err.Error())
	}

	fmt.Println(DbVal...)
	// r := gin.Default()
	// r.GET("/someJSON", func(c *gin.Context) {
	// 	data := map[string]interface{}{
	// 		"Lang":      "GoLang",
	// 		"FrameWork": "Gin",
	// 	}
	// 	c.AsciiJSON(http.StatusOK, data)
	// })
	// r.Run(":5000")
}
