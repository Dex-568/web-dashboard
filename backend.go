package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	_ "github.com/microsoft/go-mssqldb"
)

type Device struct {
	Name         string `json:"name"`
	Status       string `json:"status"`
	TotalMemory  string `json:"totalMemory"`
	MemoryUsage  string `json:"memoryUsage"`
	MemWarnThres string `json:"memWarnThres"`
	CpuCount     string `json:"cpuCount"`
	CpuUsage     string `json:"cpuUsage"`
	CpuWarnThres string `json:"cpuWarnThres"`
}

var db *sql.DB

func grabDBDevices() ([]Device, error) {

	// grab all the data, maybe get the id as well?
	rows, err := db.Query(`
	SELECT name, status, totalMemory, memoryUsage, memWarnThres,
	cpuCount, cpuUsage, cpuWarnThres FROM Device`)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	// init device struct, put the rows into something i can work with
	var devices []Device
	for rows.Next() {
		var d Device
		err := rows.Scan(
			&d.Name,
			&d.Status,
			&d.TotalMemory,
			&d.MemoryUsage,
			&d.MemWarnThres,
			&d.CpuCount,
			&d.CpuUsage,
			&d.CpuWarnThres,
		)
		if err != nil {
			return nil, err
		}
		devices = append(devices, d)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return devices, nil
}

func jsonDevices(c *gin.Context) {
	devices, err := grabDBDevices()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to load devices: " + err.Error()})
		return
	}
	// serve 200, and the devices for processing
	c.IndentedJSON(http.StatusOK, devices)
}

func main() {
	password := os.Getenv("DB_PASSWORD")
	connectString := fmt.Sprintf("sqlserver://sa:%s@localhost:1433?database=TestDB&connection+timeout=30&encrypt=true&trustservercertificate=true",
		password)

	var err error
	db, err = sql.Open("sqlserver", connectString)
	if err != nil {
		log.Fatal("Failed to open database: ", err)
	}
	defer db.Close()

	if err := db.Ping(); err != nil {
		log.Fatal("Database connection failed: ", err)
	}

	router := gin.Default()
	// handle cross-origin res sharing, not really necessary for a localhost app but
	router.Use(cors.Default())
	router.GET("/devices", jsonDevices)

	router.Run("localhost:8080")
}
