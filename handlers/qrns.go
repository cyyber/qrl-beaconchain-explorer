package handlers

import (
	"database/sql"
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/theQRL/qrl-beaconchain-explorer/cache"
	"github.com/theQRL/qrl-beaconchain-explorer/db"
	"github.com/theQRL/qrl-beaconchain-explorer/types"
	"github.com/theQRL/qrl-beaconchain-explorer/utils"

	"github.com/theQRL/go-zond/common"
)

func GetQrnsDomain(search string) (*types.QrnsDomainResponse, error) {
	data := &types.QrnsDomainResponse{}
	var returnError error

	if utils.IsValidQrnsDomain(search) {
		cacheKey := fmt.Sprintf("%d:ens:address:%v", utils.Config.Chain.ClConfig.DepositChainID, search)

		if address, err := cache.TieredCache.GetStringWithLocalTimeout(cacheKey, time.Minute); err == nil && len(address) > 0 {
			data.Address = address
			return data, nil
		}

		address, err := db.GetAddressForQrnsName(search)
		if err != nil {
			data.Domain = search
			return data, err // We want to return the data if it was a valid domain even if there was an error getting the address from bigtable. A valid domain might be enough for the caller.
		}
		data.Address = address.Hex()

		name, err := db.GetQrnsNameForAddress(*address)
		if err != nil && err != sql.ErrNoRows {
			return data, err // We want to return the data if it was a valid address even if there was an error getting the domain from bigtable. A valid address might be enough for the caller.
		}
		data.Domain = name

		err = cache.TieredCache.SetString(cacheKey, data.Address, time.Minute)
		if err != nil {
			logger.Errorf("error caching ens address: %v", err)
		}

	} else if utils.IsValidExecutionAddress(search) {
		data.Address = search

		cacheKey := fmt.Sprintf("%d:ens:domain:%v", utils.Config.Chain.ClConfig.DepositChainID, search)

		if domain, err := cache.TieredCache.GetStringWithLocalTimeout(cacheKey, time.Minute); err == nil && len(domain) > 0 {
			data.Domain = domain
			return data, nil
		}
		name, err := db.GetQrnsNameForAddress(common.HexToAddress(search))
		if err != nil && err != sql.ErrNoRows {
			return data, err // We want to return the data if it was a valid address even if there was an error getting the domain from bigtable. A valid address might be enough for the caller.
		}
		data.Domain = name
		err = cache.TieredCache.SetString(cacheKey, data.Domain, time.Minute)
		if err != nil {
			logger.Errorf("error caching ens address: %v", err)
		}
	} else {
		returnError = errors.New("not an ens domain or address")
	}
	return data, returnError //We always want to return the data if it was a valid address/domain even if there was an error getting data. A valid address might be enough for the caller.
}

func ReplaceQrnsNameWithAddress(search string) string {
	if utils.IsValidQrnsDomain(search) {
		ensData, _ := GetQrnsDomain(search)
		if len(ensData.Address) > 0 {
			search = strings.Replace(ensData.Address, "0x", "", -1)
		}
	}
	return search
}
