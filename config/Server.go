// Package config ..
package config

// Server ..
type Server struct {
	PublicPath string
	Dev        bool
	ServerName string
	ServerURI  string
}

// ServerInfo ..
var ServerInfo Server

// ServerInformations ..
func ServerInformations() Server {
	dev := true
	serverName := "ecommarce"

	var publicPath string
	var serverURI string
	if dev {
		publicPath = "./"
		serverURI = "http://192.168.1.152:8082/"

	} else {
		publicPath = "/var/www/" + serverName + "/"
		serverURI = "http://35.224.42.25:8081/"
	}

	ServerInfo = Server{
		Dev:        dev,
		PublicPath: publicPath,
		ServerName: serverName,
		ServerURI:  serverURI,
	}
	return Server{
		Dev:        dev,
		PublicPath: publicPath,
		ServerName: serverName,
		ServerURI:  serverURI,
	}
}
