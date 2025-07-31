package qrns

import (
	"github.com/theQRL/qrl-beaconchain-explorer/utils"

	"github.com/theQRL/go-zond/accounts/abi"
	"github.com/theQRL/go-zond/accounts/abi/bind"
	"github.com/theQRL/go-zond/common"
)

// TODO(now.youtrack.cloud/issue/TZB-1)
var ZNSCrontractAddressesZond = map[string]string{
	"Q00000000000C2E074eC69A0dFb2997BA6C7d2e1e": "Registry",
	"Q253553366Da8546fC250F225fe3d25d0C782303b": "ZONDRegistrarController",
	"Q283Af0B28c62C092C9727F1Ee09c02CA627EB7F5": "OldZnsRegistrarController",
}

var ZNSRegistryParsedABI, ZNSBaseRegistrarParsedABI, ZNSOldRegistrarControllerParsedABI, ZNSPublicResolverParsedABI, ZNSZONDRegistrarControllerParsedABI *abi.ABI

var ZNSRegistryContract, ZNSBaseRegistrarContract, ZNSOldRegistrarControllerContract, ZNSPublicResolverContract, ZNSZONDRegistrarControllerContract *bind.BoundContract

func init() {
	var err error

	ZNSRegistryParsedABI, err = ZNSRegistryMetaData.GetAbi()
	if err != nil {
		utils.LogFatal(err, "error getting zns-registry-abi", 0)
	}
	ZNSRegistryParsedABI, err = ZNSRegistryMetaData.GetAbi()
	if err != nil {
		utils.LogFatal(err, "error getting zns-registry-abi", 0)
	}
	ZNSBaseRegistrarParsedABI, err = ZNSBaseRegistrarMetaData.GetAbi()
	if err != nil {
		utils.LogFatal(err, "error getting zns-base-regsitrar-abi", 0)
	}
	ZNSOldRegistrarControllerParsedABI, err = ZNSOldRegistrarControllerMetaData.GetAbi()
	if err != nil {
		utils.LogFatal(err, "error getting zns-old-registrar-controller-abi", 0)
	}
	ZNSPublicResolverParsedABI, err = ZNSPublicResolverMetaData.GetAbi()
	if err != nil {
		utils.LogFatal(err, "error getting zns-public-resolver-abi", 0)
	}
	ZNSZONDRegistrarControllerParsedABI, err = ZNSZONDRegistrarControllerMetaData.GetAbi()
	if err != nil {
		utils.LogFatal(err, "error getting zns-eth-registrar-controller-abi", 0)
	}

	ZNSRegistryContract = bind.NewBoundContract(common.Address{}, *ZNSRegistryParsedABI, nil, nil, nil)
	if err != nil {
		utils.LogFatal(err, "error creating zns-registry-contract", 0)
	}
	ZNSBaseRegistrarContract = bind.NewBoundContract(common.Address{}, *ZNSBaseRegistrarParsedABI, nil, nil, nil)
	if err != nil {
		utils.LogFatal(err, "error creating zns-base-registrar-contract", 0)
	}
	ZNSOldRegistrarControllerContract = bind.NewBoundContract(common.Address{}, *ZNSOldRegistrarControllerParsedABI, nil, nil, nil)
	if err != nil {
		utils.LogFatal(err, "error creating zns-old-registrar-controller-contract", 0)
	}
	ZNSPublicResolverContract = bind.NewBoundContract(common.Address{}, *ZNSPublicResolverParsedABI, nil, nil, nil)
	if err != nil {
		utils.LogFatal(err, "error creating zns-public-resolver-contract", 0)
	}
	ZNSZONDRegistrarControllerContract = bind.NewBoundContract(common.Address{}, *ZNSZONDRegistrarControllerParsedABI, nil, nil, nil)
	if err != nil {
		utils.LogFatal(err, "error creating zns-eth-registrar-controller-contract", 0)
	}
}
