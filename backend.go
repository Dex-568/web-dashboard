package main

import (
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
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

// mock data, could do a db later
var devices = []Device{
	{
		Name:         "Device 1",
		Status:       "Online",
		TotalMemory:  "8192MB",
		MemoryUsage:  "42%",
		MemWarnThres: "No",
		CpuCount:     "8",
		CpuUsage:     "62%",
		CpuWarnThres: "No",
	},

	{
		Name:         "Device 2",
		Status:       "Online",
		TotalMemory:  "16382MB",
		MemoryUsage:  "99%",
		MemWarnThres: "Yes",
		CpuCount:     "16",
		CpuUsage:     "22%",
		CpuWarnThres: "No",
	},

	{
		Name:         "Device 3",
		Status:       "Online",
		TotalMemory:  "4092MB",
		MemoryUsage:  "42%",
		MemWarnThres: "No",
		CpuCount:     "4",
		CpuUsage:     "82%",
		CpuWarnThres: "Yes",
	},
}

func getDevices(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, devices)
}

func main() {
	router := gin.Default()
	// handle cross-origin res sharing, not really necessary for a localhost app but
	router.Use(cors.Default())
	router.GET("/devices", getDevices)

	router.Run("localhost:8080")
}
