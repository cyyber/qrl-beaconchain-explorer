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
	"Q253553366Da8546fC250F225fe3d25d0C782303b": "QRLRegistrarController",
	"Q283Af0B28c62C092C9727F1Ee09c02CA627EB7F5": "OldQrnsRegistrarController",
}

var QRNSRegistryParsedABI, QRNSBaseRegistrarParsedABI, QRNSOldRegistrarControllerParsedABI, QRNSPublicResolverParsedABI, QRNSQRLRegistrarControllerParsedABI *abi.ABI

var QRNSRegistryContract, QRNSBaseRegistrarContract, QRNSOldRegistrarControllerContract, QRNSPublicResolverContract, QRNSQRLRegistrarControllerContract *bind.BoundContract

func init() {
	var err error

	QRNSRegistryParsedABI, err = QRNSRegistryMetaData.GetAbi()
	if err != nil {
		utils.LogFatal(err, "error getting qrns-registry-abi", 0)
	}
	QRNSRegistryParsedABI, err = QRNSRegistryMetaData.GetAbi()
	if err != nil {
		utils.LogFatal(err, "error getting qrns-registry-abi", 0)
	}
	QRNSBaseRegistrarParsedABI, err = QRNSBaseRegistrarMetaData.GetAbi()
	if err != nil {
		utils.LogFatal(err, "error getting qrns-base-regsitrar-abi", 0)
	}
	QRNSOldRegistrarControllerParsedABI, err = QRNSOldRegistrarControllerMetaData.GetAbi()
	if err != nil {
		utils.LogFatal(err, "error getting qrns-old-registrar-controller-abi", 0)
	}
	QRNSPublicResolverParsedABI, err = QRNSPublicResolverMetaData.GetAbi()
	if err != nil {
		utils.LogFatal(err, "error getting qrns-public-resolver-abi", 0)
	}
	QRNSQRLRegistrarControllerParsedABI, err = QRNSQRLRegistrarControllerMetaData.GetAbi()
	if err != nil {
		utils.LogFatal(err, "error getting qrns-eth-registrar-controller-abi", 0)
	}

	QRNSRegistryContract = bind.NewBoundContract(common.Address{}, *QRNSRegistryParsedABI, nil, nil, nil)
	if err != nil {
		utils.LogFatal(err, "error creating qrns-registry-contract", 0)
	}
	QRNSBaseRegistrarContract = bind.NewBoundContract(common.Address{}, *QRNSBaseRegistrarParsedABI, nil, nil, nil)
	if err != nil {
		utils.LogFatal(err, "error creating qrns-base-registrar-contract", 0)
	}
	QRNSOldRegistrarControllerContract = bind.NewBoundContract(common.Address{}, *QRNSOldRegistrarControllerParsedABI, nil, nil, nil)
	if err != nil {
		utils.LogFatal(err, "error creating qrns-old-registrar-controller-contract", 0)
	}
	QRNSPublicResolverContract = bind.NewBoundContract(common.Address{}, *QRNSPublicResolverParsedABI, nil, nil, nil)
	if err != nil {
		utils.LogFatal(err, "error creating qrns-public-resolver-contract", 0)
	}
	QRNSQRLRegistrarControllerContract = bind.NewBoundContract(common.Address{}, *QRNSQRLRegistrarControllerParsedABI, nil, nil, nil)
	if err != nil {
		utils.LogFatal(err, "error creating qrns-eth-registrar-controller-contract", 0)
	}
}
