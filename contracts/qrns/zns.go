package qrns

import (
	"github.com/theQRL/qrl-beaconchain-explorer/utils"

	"github.com/theQRL/go-zond/accounts/abi"
	"github.com/theQRL/go-zond/accounts/abi/bind"
	"github.com/theQRL/go-zond/common"
)

// TODO(rgeraldes24)
var QRNSCrontractAddressesZond = map[string]string{
	"Q00000000000C2E074eC69A0dFb2997BA6C7d2e1e": "Registry",
	"Q253553366Da8546fC250F225fe3d25d0C782303b": "ZONDRegistrarController",
	"Q283Af0B28c62C092C9727F1Ee09c02CA627EB7F5": "OldZnsRegistrarController",
}

var QRNSRegistryParsedABI, QRNSBaseRegistrarParsedABI, QRNSOldRegistrarControllerParsedABI, QRNSPublicResolverParsedABI, QRNSZONDRegistrarControllerParsedABI *abi.ABI

var QRNSRegistryContract, QRNSBaseRegistrarContract, QRNSOldRegistrarControllerContract, QRNSPublicResolverContract, QRNSZONDRegistrarControllerContract *bind.BoundContract

func init() {
	var err error

	QRNSRegistryParsedABI, err = QRNSRegistryMetaData.GetAbi()
	if err != nil {
		utils.LogFatal(err, "error getting zns-registry-abi", 0)
	}
	QRNSRegistryParsedABI, err = QRNSRegistryMetaData.GetAbi()
	if err != nil {
		utils.LogFatal(err, "error getting zns-registry-abi", 0)
	}
	QRNSBaseRegistrarParsedABI, err = QRNSBaseRegistrarMetaData.GetAbi()
	if err != nil {
		utils.LogFatal(err, "error getting zns-base-regsitrar-abi", 0)
	}
	QRNSOldRegistrarControllerParsedABI, err = QRNSOldRegistrarControllerMetaData.GetAbi()
	if err != nil {
		utils.LogFatal(err, "error getting zns-old-registrar-controller-abi", 0)
	}
	QRNSPublicResolverParsedABI, err = QRNSPublicResolverMetaData.GetAbi()
	if err != nil {
		utils.LogFatal(err, "error getting zns-public-resolver-abi", 0)
	}
	QRNSZONDRegistrarControllerParsedABI, err = QRNSZONDRegistrarControllerMetaData.GetAbi()
	if err != nil {
		utils.LogFatal(err, "error getting zns-eth-registrar-controller-abi", 0)
	}

	QRNSRegistryContract = bind.NewBoundContract(common.Address{}, *QRNSRegistryParsedABI, nil, nil, nil)
	if err != nil {
		utils.LogFatal(err, "error creating zns-registry-contract", 0)
	}
	QRNSBaseRegistrarContract = bind.NewBoundContract(common.Address{}, *QRNSBaseRegistrarParsedABI, nil, nil, nil)
	if err != nil {
		utils.LogFatal(err, "error creating zns-base-registrar-contract", 0)
	}
	QRNSOldRegistrarControllerContract = bind.NewBoundContract(common.Address{}, *QRNSOldRegistrarControllerParsedABI, nil, nil, nil)
	if err != nil {
		utils.LogFatal(err, "error creating zns-old-registrar-controller-contract", 0)
	}
	QRNSPublicResolverContract = bind.NewBoundContract(common.Address{}, *QRNSPublicResolverParsedABI, nil, nil, nil)
	if err != nil {
		utils.LogFatal(err, "error creating zns-public-resolver-contract", 0)
	}
	QRNSZONDRegistrarControllerContract = bind.NewBoundContract(common.Address{}, *QRNSZONDRegistrarControllerParsedABI, nil, nil, nil)
	if err != nil {
		utils.LogFatal(err, "error creating zns-eth-registrar-controller-contract", 0)
	}
}
