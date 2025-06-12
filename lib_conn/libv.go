package common

import (
	"log"
	"sync"

	"libvirt.org/go/libvirt"
)

var (
	conn *libvirt.Connect
	once sync.Once
)

const libConnectionString = "qemu:///system"

// InitConnection initializes the libvirt connection
func initConnection() *libvirt.Connect {
	var err error
	conn, err = libvirt.NewConnect(libConnectionString)
	if err != nil {
		log.Fatalf("Failed to connect to libvirt: %v", err)
	}
	return conn
}

// GetConnection returns the existing connection
func getConnection() *libvirt.Connect {
	if conn == nil {
		log.Fatal("Libvirt connection not initialized. Call InitConnection first.")
	}
	return conn
}

// CloseConnection closes the libvirt connection
func CloseConnection() {
	if conn != nil {
		conn.Close()
	}
}

// GetLibvClient initializes the connection once and returns it
func GetLibvClient() *libvirt.Connect {
	once.Do(func() {
		initConnection()
	})
	return getConnection()
}
