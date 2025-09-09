package qrns

//go:generate abigen -abi qrns_registry.json -out qrns_registry.go -pkg qrns -type QRNSRegistry
//go:generate abigen -abi qrns_base_registrar.json -out qrns_base_registrar.go -pkg qrns -type QRNSBaseRegistrar
//go:generate abigen -abi qrns_qrl_registrar_controller.json -out qrns_qrl_registrar_controller.go -pkg qrns -type QRNSQRLRegistrarController
//go:generate abigen -abi qrns_dns_registrar.json -out qrns_dns_registrar.go -pkg qrns -type QRNSDNSRegistrar
//go:generate abigen -abi qrns_reverse_registrar.json -out qrns_reverse_registrar.go -pkg qrns -type QRNSReverseRegistrar
//go:generate abigen -abi qrns_name_wrapper.json -out qrns_name_wrapper.go -pkg qrns -type QRNSNameWrapper
//go:generate abigen -abi qrns_public_resolver.json -out qrns_public_resolver.go -pkg qrns -type QRNSPublicResolver
//go:generate abigen -abi qrns_old_regstrar_controller.json -out qrns_old_regstrar_controller.go -pkg qrns -type QRNSOldRegistrarController
