package networkops

import (
	"fmt"
	"log"
	"strings"
	"time"

	lib_conn "abc/lib_conn"

	"github.com/go-playground/validator"
	"github.com/google/uuid"
	libvirtxml "github.com/libvirt/libvirt-go-xml"
)

/*
Name: mandatory

Bridge conf:

	BridgeName: mandatory
	stp: default "on"
	delay: default "0"

ForwardMode:

	ForwardMode: type of network (mandatory)
	ForwardDev(More related to super user or admin user)

DHCP conf:

	IPAddress: Address
	Netmask: Netmask
	DHCP:  Start, END

Static IP conf:

	StaticMAC: MAC
	StaticIP: IP
	StaticHostName: MAC
*/
const (
	ForwardModeNAT     = "nat"
	ForwardModeRoute   = "route"
	ForwardModeIsolate = "isolate"
	ForwardModeBridge  = "bridge"
)

type NetworkValidationErrors []string // Custom type to be used for error during Network Configuration validations

func (ne NetworkValidationErrors) Error() string {
	return strings.Join(ne, "; ")
}

var (
	STP        string = "on"
	Delay      string = "0"
	ForwardDev string = "enp1s0"
)

type StaticConfig struct {
	MAC      string `validate:"required,mac"`
	IP       string `validate:"required,ip"`
	HostName string `validate:"required,hostname"`
}

type IPConfig struct {
	IP             string `validate:"required,ip"`
	Netmask        string `validate:"required,ip"`
	DHCPRangeStart string `validate:"required,ip"`
	DHCPRangeEnd   string `validate:"required,ip"`
}

type NetworkConf struct {
	Name        string `validate:"required"`
	BridgeName  string
	ForwardMode string        `validate:"required,oneof=nat route isolate bridge"`
	IPConf      *IPConfig     `validate:"omitempty,dive"`
	Static      *StaticConfig `validate:"omitempty,dive"` // Only validate if not nil
}

type NetworkDbData struct {
	Id            uuid.UUID `validate:"required"`
	Name          string    `validate:"required"`
	Autostart     bool      `validate:"required"`
	IsPersistent  bool      `validate:"required"`
	State         bool      `validate:"required"`
	GetDHCPLeases []string
	Ports         []string
	CreatedAt     time.Time `validate:"required"`
}

func ValidateNetworkConfig(cfg *NetworkConf) error {
	validate := validator.New()
	var errs NetworkValidationErrors

	// First, validate basic rules
	if err := validate.Struct(cfg); err != nil {
		return err
	}

	// Conditional validations
	switch cfg.ForwardMode {

	case ForwardModeBridge:
		if cfg.IPConf != nil {
			errs = append(errs, fmt.Sprintf("No IP configuration is required if network mode is: %s", cfg.ForwardMode))
		}

	default:
		return nil
	}

	return NetworkValidationErrors(errs)
}

func NetworkXMLBuilder(cfg *NetworkConf) *libvirtxml.Network {
	var forwardConfig *libvirtxml.NetworkForward = nil // stores forward mode configuration
	var bridgeConfig *libvirtxml.NetworkBridge = nil   // stores bridge configuration
	var ipConfig []libvirtxml.NetworkIP = nil

	forwardConfig = getForwardConfig(cfg)
	bridgeConfig = getbBridgeConfig(cfg)
	ipConfig = getIPConfig(cfg)

	nwXML := &libvirtxml.Network{
		Name:    cfg.Name,
		Forward: forwardConfig,
		Bridge:  bridgeConfig,
		IPs:     ipConfig,
	}

	return nwXML
}

func getIPConfig(cfg *NetworkConf) []libvirtxml.NetworkIP {
	// helper function for 'NetworkXMLBuilder' to configure IP
	if cfg.ForwardMode == ForwardModeBridge {
		return nil
	}

	dhcpConfig := &libvirtxml.NetworkDHCP{
		Ranges: []libvirtxml.NetworkDHCPRange{
			{
				Start: cfg.IPConf.DHCPRangeStart,
				End:   cfg.IPConf.DHCPRangeEnd,
			},
		},
	}
	// If Static config is provided, add it to DHCP hosts
	if cfg.Static != nil {
		dhcpConfig.Hosts = []libvirtxml.NetworkDHCPHost{
			{
				MAC:  cfg.Static.MAC,
				Name: cfg.Static.HostName,
				IP:   cfg.Static.IP,
			},
		}
	}
	ipConf := []libvirtxml.NetworkIP{
		{
			Address: cfg.IPConf.IP,
			Netmask: cfg.IPConf.Netmask,
			DHCP:    dhcpConfig,
		},
	}
	return ipConf

}

func getForwardConfig(cfg *NetworkConf) *libvirtxml.NetworkForward {
	// helper function for 'NetworkXMLBuilder' to configure Forward mode

	return &libvirtxml.NetworkForward{
		Mode: cfg.ForwardMode,
		Dev:  ForwardDev,
	} // common for "routed" and "NAT" network and will be omitted for "isolated" network
}

func getbBridgeConfig(cfg *NetworkConf) *libvirtxml.NetworkBridge {
	// helper function for 'NetworkXMLBuilder' to configure Bridge

	if cfg.ForwardMode == ForwardModeBridge {
		return nil
	}

	return &libvirtxml.NetworkBridge{
		Name:  cfg.BridgeName,
		STP:   STP,
		Delay: Delay,
	}
}

func NetworkCreate(cfg *NetworkConf) error {
	conn := lib_conn.GetLibvClient()
	// defer client.CloseConnection()

	// if err := ValidateNetworkConfig(cfg); err != nil {
	// 	return err
	// }
	// creating XML for network
	err := ValidateNetworkConfig(cfg)
	if err != nil {
		return err
	}

	netXML := NetworkXMLBuilder(cfg)
	// Marshal the network definition to XML
	networkXML, err := netXML.Marshal()

	if err != nil {
		return err
	}

	networkLib, err := conn.NetworkDefineXML(networkXML)
	if err != nil {
		return err
	}

	// Start the network
	err = networkLib.Create()
	if err != nil {
		nw, _ := conn.LookupNetworkByName(cfg.Name)
		checkActive, _ := nw.IsActive()
		if !checkActive {
			_ = nw.Undefine()
			return err
		}
	}
	// Set the network to autostart
	if err := networkLib.SetAutostart(true); err != nil {
		return err
	}
	log.Printf("Network %s created and started successfully.", netXML.Name)

	// Fetching network details
	_, err = conn.LookupNetworkByName(cfg.Name)
	if err != nil {
		return err
	}

	return nil
}
