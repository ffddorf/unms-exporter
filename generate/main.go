package main

import (
	"github.com/go-swagger/go-swagger/generator"
)

const specLocation = "openapi.json"

var includedOperations = []string{
	"getDevices",
}

// manually compiled list of models which getDevices needs
var includedModels = []string{
	"addresses",
	"antenna",
	"antenna 1",
	"apDevice",
	"capabilities",
	"capacities",
	"capacities 1",
	"current",
	"DeviceAttributes",
	"DeviceFirmware",
	"DeviceIdentification",
	"DeviceIdentification 1",
	"DeviceInterfaceListSchema",
	"DeviceInterfaceSchema",
	"DeviceMeta",
	"DeviceOverview",
	"DeviceStatusOverview",
	"DeviceUpgrade",
	"dfsLockouts",
	"discovery",
	"Error",
	"eswitch",
	"Firmware",
	"FirmwareIdentification",
	"FirmwareSemVer",
	"FirmwareSupport",
	"interface",
	"InterfaceIdentification",
	"InterfaceOspf",
	"InterfacePoe",
	"InterfaceSpeeds",
	"InterfaceStatistics",
	"InterfaceStatus",
	"lag",
	"latest",
	"latestBackup",
	"linkScore",
	"ListOfDevices",
	"loadBalanceValues",
	"Model 2",
	"Model 3",
	"models",
	"ospfConfig",
	"parent 1",
	"pingWatchdog",
	"port",
	"ports",
	"ports 1",
	"prerelease",
	"rxChain",
	"semver",
	"site",
	"speedLimit",
	"Station",
	"stations",
	"statistics",
	"stp",
	"stp 1",
	"switch",
	"txChain",
	"unmss",
	"validation",
	"vid",
	"visibleBy",
	"vlans",
	"wireless",
	"wirelessActiveInterfaceIds",
}

func main() {
	opts := &generator.GenOpts{
		IncludeModel:      true,
		IncludeValidator:  true,
		IncludeHandler:    true,
		IncludeParameters: true,
		IncludeResponses:  true,
		IncludeURLBuilder: false,
		IncludeMain:       true,
		IncludeSupport:    true,
		ValidateSpec:      false,
		IsClient:          true,
		Spec:              specLocation,
		Target:            ".",
		ModelPackage:      "models",
		ClientPackage:     "client",
	}
	if err := opts.EnsureDefaults(); err != nil {
		panic(err)
	}

	err := generator.GenerateClient("", includedModels, includedOperations, opts)
	if err != nil {
		panic(err)
	}
}
