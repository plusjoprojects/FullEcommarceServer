// Package app ...
package app

import (
	"server/config"
	"server/controllers"
	"server/routes"
	"server/vendors"
)

// Installing ..
func Installing() {
	// Setup The DB
	config.SetupDB()

	config.ServerInformations()

	// SetupPassport
	vendors.SetupPassport()
	vendors.CreateAdmin()
	controllers.StoreStorages()

	// setup routes
	routes.Setup()
}
