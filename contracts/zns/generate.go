package zns

//go:generate abigen -abi zns_registry.json -out zns_registry.go -pkg zns -type ZNSRegistry
//go:generate abigen -abi zns_base_registrar.json -out zns_base_registrar.go -pkg zns -type ZNSBaseRegistrar
//go:generate abigen -abi zns_zond_registrar_controller.json -out zns_zond_registrar_controller.go -pkg zns -type ZNSZONDRegistrarController
//go:generate abigen -abi zns_dns_registrar.json -out zns_dns_registrar.go -pkg zns -type ZNSDNSRegistrar
//go:generate abigen -abi zns_reverse_registrar.json -out zns_reverse_registrar.go -pkg zns -type ZNSReverseRegistrar
//go:generate abigen -abi zns_name_wrapper.json -out zns_name_wrapper.go -pkg zns -type ZNSNameWrapper
//go:generate abigen -abi zns_public_resolver.json -out zns_public_resolver.go -pkg zns -type ZNSPublicResolver
//go:generate abigen -abi zns_old_regstrar_controller.json -out zns_old_regstrar_controller.go -pkg zns -type ZNSOldRegistrarController
