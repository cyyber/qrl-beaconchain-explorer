package db

// TODO(rgeraldes24): unused for now
/*
import (
	"bytes"
	"context"
	"database/sql"
	"encoding/hex"
	"fmt"
	"math/big"
	"strings"
	"sync"
	"time"

	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/sirupsen/logrus"
	znsContracts "github.com/theQRL/zond-beaconchain-explorer/contracts/zns"
	"github.com/theQRL/zond-beaconchain-explorer/metrics"
	"github.com/theQRL/zond-beaconchain-explorer/types"
	"github.com/theQRL/zond-beaconchain-explorer/utils"

	gcp_bigtable "cloud.google.com/go/bigtable"
	"golang.org/x/sync/errgroup"

	"github.com/coocood/freecache"
	"github.com/theQRL/go-zond/accounts/abi/bind"
	"github.com/theQRL/go-zond/common"

	zond_types "github.com/theQRL/go-zond/core/types"
	go_zns "github.com/wealdtech/go-zns/v3"
)

// TransformZnsNameRegistered accepts an eth1 block and creates bigtable mutations for ZNS Name events.
// It transforms the logs contained within a block and indexes zns relevant transactions and tags changes (to be verified from the node in a separate process)
// ==================================================
//
// It indexes transactions
//
// - by hashed zns name
// Row:    <chainID>:ZNS:I:H:<nameHash>:<txHash>
// Family: f
// Column: nil
// Cell:   nil
// Example scan: "5:ZNS:I:H:4ae569dd0aa2f6e9207e41423c956d0d27cbc376a499ee8d90fe1d84489ae9d1:e627ae94bd16eb1ed8774cd4003fc25625159f13f8a2612cc1c7f8d2ab11b1d7"
//
// - by address
// Row:    <chainID>:ZNS:I:A:<address>:<txHash>
// Family: f
// Column: nil
// Cell:   nil
// Example scan: "5:ZNS:I:A:05579fadcf7cc6544f7aa018a2726c85251600c5:e627ae94bd16eb1ed8774cd4003fc25625159f13f8a2612cc1c7f8d2ab11b1d7"
//
// ==================================================
//
// Track for later verification via the node ("set dirty")
//
// - by name
// Row:    <chainID>:ZNS:V:N:<name>
// Family: f
// Column: nil
// Cell:   nil
// Example scan: "5:ZNS:V:N:somename"
//
// - by name hash
// Row:    <chainID>:ZNS:V:H:<nameHash>
// Family: f
// Column: nil
// Cell:   nil
// Example scan: "5:ZNS:V:H:6f5d9cc23e60abe836401b4fd386ec9280a1f671d47d9bf3ec75dab76380d845"
//
// - by address
// Row:    <chainID>:ZNS:V:A:<address>
// Family: f
// Column: nil
// Cell:   nil
// Example scan: "5:ZNS:V:A:27234cb8734d5b1fac0521c6f5dc5aebc6e839b6"
//
// ==================================================

func (bigtable *Bigtable) TransformZnsNameRegistered(blk *types.Eth1Block, cache *freecache.Cache) (bulkData *types.BulkMutations, bulkMetadataUpdates *types.BulkMutations, err error) {
	startTime := time.Now()
	defer func() {
		metrics.TaskDuration.WithLabelValues("bt_transform_zns").Observe(time.Since(startTime).Seconds())
	}()

	var znsCrontractAddresses map[string]string
	switch bigtable.chainId {
	case "1":
		znsCrontractAddresses = znsContracts.ZNSCrontractAddressesEthereum
	default:
		return nil, nil, nil
	}

	bulkData = &types.BulkMutations{}
	bulkMetadataUpdates = &types.BulkMutations{}
	keys := make(map[string]bool)
	ethLog := zond_types.Log{}

	for i, tx := range blk.GetTransactions() {
		if i >= TX_PER_BLOCK_LIMIT {
			return nil, nil, fmt.Errorf("unexpected number of transactions in block expected at most %d but got: %v, tx: %x", TX_PER_BLOCK_LIMIT-1, i, tx.GetHash())
		}
		for j, log := range tx.GetLogs() {
			if j >= ITX_PER_TX_LIMIT {
				return nil, nil, fmt.Errorf("unexpected number of logs in block expected at most %d but got: %v tx: %x", ITX_PER_TX_LIMIT-1, j, tx.GetHash())
			}
			znsContract := znsCrontractAddresses[common.BytesToAddress(log.Address).String()]

			topics := log.GetTopics()
			ethTopics := make([]common.Hash, 0, len(topics))
			for _, t := range topics {
				ethTopics = append(ethTopics, common.BytesToHash(t))
			}

			ethLog.Address = common.BytesToAddress(log.GetAddress())
			ethLog.Data = log.Data
			ethLog.Topics = ethTopics
			ethLog.BlockNumber = blk.GetNumber()
			ethLog.TxHash = common.BytesToHash(tx.GetHash())
			ethLog.TxIndex = uint(i)
			ethLog.BlockHash = common.BytesToHash(blk.GetHash())
			ethLog.Index = uint(j)
			ethLog.Removed = log.GetRemoved()

			for _, lTopic := range topics {
				logFields := map[string]interface{}{
					"block":       blk.GetNumber(),
					"tx":          tx.GetHash(),
					"logIndex":    j,
					"znsContract": znsContract,
				}

				if znsContract == "Registry" {
					if bytes.Equal(lTopic, znsContracts.ZNSRegistryParsedABI.Events["NewResolver"].ID.Bytes()) {
						logFields["event"] = "NewResolver"
						r := &znsContracts.ZNSRegistryNewResolver{}
						err = znsContracts.ZNSRegistryContract.UnpackLog(r, "NewResolver", ethLog)
						if err != nil {
							utils.LogWarn(err, "error unpacking zns-log", 0, logFields)
							continue
						}
						keys[fmt.Sprintf("%s:ZNS:V:H:%x", bigtable.chainId, r.Node)] = true
					} else if bytes.Equal(lTopic, znsContracts.ZNSRegistryParsedABI.Events["NewOwner"].ID.Bytes()) {
						logFields["event"] = "NewOwner"
						r := &znsContracts.ZNSRegistryNewOwner{}
						err = znsContracts.ZNSRegistryContract.UnpackLog(r, "NewOwner", ethLog)
						if err != nil {
							utils.LogWarn(err, "error unpacking zns-log", 0, logFields)
							continue
						}
						keys[fmt.Sprintf("%s:ZNS:V:A:%x", bigtable.chainId, r.Owner)] = true
					} else if bytes.Equal(lTopic, znsContracts.ZNSRegistryParsedABI.Events["NewTTL"].ID.Bytes()) {
						logFields["event"] = "NewTTL"
						r := &znsContracts.ZNSRegistryNewTTL{}
						err = znsContracts.ZNSRegistryContract.UnpackLog(r, "NewTTL", ethLog)
						if err != nil {
							utils.LogWarn(err, "error unpacking zns-log", 0, logFields)
							continue
						}
						keys[fmt.Sprintf("%s:ZNS:V:H:%x", bigtable.chainId, r.Node)] = true
					}
				} else if znsContract == "ZONDRegistrarController" {
					if bytes.Equal(lTopic, znsContracts.ZNSZONDRegistrarControllerParsedABI.Events["NameRegistered"].ID.Bytes()) {
						logFields["event"] = "NameRegistered"
						r := &znsContracts.ZNSZONDRegistrarControllerNameRegistered{}
						err = znsContracts.ZNSZONDRegistrarControllerContract.UnpackLog(r, "NameRegistered", ethLog)
						if err != nil {
							utils.LogWarn(err, "error unpacking zns-log", 0, logFields)
							continue
						}
						keys[fmt.Sprintf("%s:ZNS:V:N:%s", bigtable.chainId, r.Name)] = true
						keys[fmt.Sprintf("%s:ZNS:V:A:%x", bigtable.chainId, r.Owner)] = true
					} else if bytes.Equal(lTopic, znsContracts.ZNSZONDRegistrarControllerParsedABI.Events["NameRenewed"].ID.Bytes()) {
						logFields["event"] = "NameRenewed"
						r := &znsContracts.ZNSZONDRegistrarControllerNameRenewed{}
						err = znsContracts.ZNSZONDRegistrarControllerContract.UnpackLog(r, "NameRenewed", ethLog)
						if err != nil {
							utils.LogWarn(err, "error unpacking zns-log", 0, logFields)
							continue
						}
						keys[fmt.Sprintf("%s:ZNS:V:N:%s", bigtable.chainId, r.Name)] = true
					}
				} else if znsContract == "OldZnsRegistrarController" {
					if bytes.Equal(lTopic, znsContracts.ZNSOldRegistrarControllerParsedABI.Events["NameRegistered"].ID.Bytes()) {
						logFields["event"] = "NameRegistered"
						r := &znsContracts.ZNSOldRegistrarControllerNameRegistered{}
						err = znsContracts.ZNSOldRegistrarControllerContract.UnpackLog(r, "NameRegistered", ethLog)
						if err != nil {
							utils.LogWarn(err, "error unpacking zns-log", 0, logFields)
							continue
						}
						if err = verifyName(r.Name); err != nil {
							utils.LogWarn(err, "error verifying zns-name", 0, logFields)
							continue
						}
						keys[fmt.Sprintf("%s:ZNS:V:N:%s", bigtable.chainId, r.Name)] = true
						keys[fmt.Sprintf("%s:ZNS:V:A:%x", bigtable.chainId, r.Owner)] = true
					} else if bytes.Equal(lTopic, znsContracts.ZNSOldRegistrarControllerParsedABI.Events["NameRenewed"].ID.Bytes()) {
						logFields["event"] = "NameRenewed"
						r := &znsContracts.ZNSOldRegistrarControllerNameRenewed{}
						err = znsContracts.ZNSOldRegistrarControllerContract.UnpackLog(r, "NameRenewed", ethLog)
						if err != nil {
							utils.LogWarn(err, "error unpacking zns-log", 0, logFields)
							continue
						}
						if err = verifyName(r.Name); err != nil {
							utils.LogWarn(err, "error verifying zns-name", 0, logFields)
							continue
						}
						keys[fmt.Sprintf("%s:ZNS:V:N:%s", bigtable.chainId, r.Name)] = true
					}
				} else {
					if bytes.Equal(lTopic, znsContracts.ZNSPublicResolverParsedABI.Events["NameChanged"].ID.Bytes()) {
						logFields["event"] = "NameChanged"
						r := &znsContracts.ZNSPublicResolverNameChanged{}
						err = znsContracts.ZNSPublicResolverContract.UnpackLog(r, "NameChanged", ethLog)
						if err != nil {
							utils.LogWarn(err, "error unpacking zns-log", 0, logFields)
							continue
						}
						if err = verifyName(r.Name); err != nil {
							utils.LogWarn(err, "error verifying zns-name", 0, logFields)
							continue
						}
						keys[fmt.Sprintf("%s:ZNS:V:N:%s", bigtable.chainId, r.Name)] = true
					} else if bytes.Equal(lTopic, znsContracts.ZNSPublicResolverParsedABI.Events["AddressChanged"].ID.Bytes()) {
						logFields["event"] = "AddressChanged"
						r := &znsContracts.ZNSPublicResolverAddressChanged{}
						err = znsContracts.ZNSPublicResolverContract.UnpackLog(r, "AddressChanged", ethLog)
						if err != nil {
							utils.LogWarn(err, "error unpacking zns-log", 0, logFields)
							continue
						}
						keys[fmt.Sprintf("%s:ZNS:V:H:%x", bigtable.chainId, r.Node)] = true
					}
				}
			}
		}
	}
	for key := range keys {
		mut := gcp_bigtable.NewMutation()
		mut.Set(DEFAULT_FAMILY, key, gcp_bigtable.Timestamp(0), nil)

		bulkData.Keys = append(bulkData.Keys, key)
		bulkData.Muts = append(bulkData.Muts, mut)
	}

	return bulkData, bulkMetadataUpdates, nil
}

func verifyName(name string) error {
	// limited by max capacity of db (caused by btrees of indexes); tests showed maximum of 2684 (added buffer)
	if len(name) > 2048 {
		return fmt.Errorf("name too long: %v", name)
	}
	return nil
}

type ZnsCheckedDictionary struct {
	mux     sync.Mutex
	address map[common.Address]bool
	name    map[string]bool
}

func (bigtable *Bigtable) GetRowsByPrefix(prefix string) ([]string, error) {
	ctx, done := context.WithTimeout(context.Background(), time.Second*30)
	defer done()

	rowRange := gcp_bigtable.PrefixRange(prefix)
	keys := []string{}

	err := bigtable.tableData.ReadRows(ctx, rowRange, func(row gcp_bigtable.Row) bool {
		row_ := row[DEFAULT_FAMILY][0]
		keys = append(keys, row_.Row)
		return true
	}, gcp_bigtable.LimitRows(1000))
	if err != nil {
		return nil, err
	}

	return keys, nil
}

func (bigtable *Bigtable) ImportZnsUpdates(client *ethclient.Client, readBatchSize int64) error {
	startTime := time.Now()
	defer func() {
		metrics.TaskDuration.WithLabelValues("bt_import_zns_updates").Observe(time.Since(startTime).Seconds())
	}()

	key := fmt.Sprintf("%s:ZNS:V", bigtable.chainId)

	ctx, done := context.WithTimeout(context.Background(), time.Second*30)
	defer done()

	rowRange := gcp_bigtable.PrefixRange(key)
	keys := []string{}

	err := bigtable.tableData.ReadRows(ctx, rowRange, func(row gcp_bigtable.Row) bool {
		row_ := row[DEFAULT_FAMILY][0]
		keys = append(keys, row_.Row)
		return true
	}, gcp_bigtable.LimitRows(readBatchSize)) // limit to max 1000 entries to avoid blocking the import of new blocks
	if err != nil {
		return err
	}

	if len(keys) == 0 {
		logger.Info("No ZNS entries to validate")
		return nil
	}

	logger.Infof("Validating %v ZNS entries", len(keys))
	alreadyChecked := ZnsCheckedDictionary{
		address: make(map[common.Address]bool),
		name:    make(map[string]bool),
	}

	mutDelete := gcp_bigtable.NewMutation()
	mutDelete.DeleteRow()

	batchSize := 100
	total := len(keys)
	for i := 0; i < total; i += batchSize {
		to := i + batchSize
		if to > total {
			to = total
		}
		batch := keys[i:to]
		logger.Infof("Batching ZNS entries %v:%v of %v", i, to, total)

		g := new(errgroup.Group)
		g.SetLimit(10) // limit load on the node
		mutsDelete := &types.BulkMutations{
			Keys: make([]string, 0, 1),
			Muts: make([]*gcp_bigtable.Mutation, 0, 1),
		}

		for _, k := range batch {
			key := k
			var name string
			var address *common.Address
			split := strings.Split(key, ":")
			value := split[4]

			switch split[3] {
			case "H":
				// if we have a hash we look if we find a name in the db. If not we can ignore it.
				nameHash, err := hex.DecodeString(value)
				if err != nil {
					utils.LogError(err, fmt.Errorf("name hash could not be decoded: %v", value), 0)
				} else {
					err := ReaderDb.Get(&name, `
					SELECT
						zns_name
					FROM zns
					WHERE name_hash = $1
					`, nameHash[:])
					if err != nil && err != sql.ErrNoRows {
						return err
					}
				}
			case "A":
				addressHash, err := hex.DecodeString(value)
				if err != nil {
					utils.LogError(err, fmt.Errorf("address hash could not be decoded: %v", value), 0)
				} else {
					add := common.BytesToAddress(addressHash)
					address = &add
				}
			case "N":
				name = value
			}

			g.Go(func() error {
				if name != "" {
					err := validateZnsName(client, name, &alreadyChecked)
					if err != nil {
						return fmt.Errorf("error validating new name [%v]: %w", name, err)
					}
				} else if address != nil {
					err := validateZnsAddress(client, *address, &alreadyChecked)
					if err != nil {
						return fmt.Errorf("error validating new address [%v]: %w", address, err)
					}
				}
				return nil
			})

			mutsDelete.Keys = append(mutsDelete.Keys, key)
			mutsDelete.Muts = append(mutsDelete.Muts, mutDelete)
		}

		if err := g.Wait(); err != nil {
			return err
		}

		// After processing a batch of keys we remove them from bigtable
		err = bigtable.WriteBulk(mutsDelete, bigtable.tableData, DEFAULT_BATCH_INSERTS)
		if err != nil {
			return err
		}

		// give node some time for other stuff between batches
		time.Sleep(time.Millisecond * 100)
	}

	logger.WithField("updates", total).Info("Import of ZNS updates completed")
	return nil
}

func validateZnsAddress(client *ethclient.Client, address common.Address, alreadyChecked *ZnsCheckedDictionary) error {
	alreadyChecked.mux.Lock()
	if alreadyChecked.address[address] {
		alreadyChecked.mux.Unlock()
		return nil
	}
	alreadyChecked.address[address] = true
	alreadyChecked.mux.Unlock()

	names := []string{}
	err := ReaderDb.Select(&names, `SELECT zns_name FROM zns WHERE address = $1 AND is_primary_name AND valid_to >= now()`, address.Bytes())
	if err != nil && err != sql.ErrNoRows {
		return err
	}

	for _, name := range names {
		if name != "" {
			err = validateZnsName(client, name, alreadyChecked)
			if err != nil {
				return err
			}
		}
		reverseName, err := go_zns.ReverseResolve(client, address)
		if err != nil {
			if err.Error() == "not a resolver" ||
				err.Error() == "no resolution" ||
				err.Error() == "execution reverted" ||
				strings.HasPrefix(err.Error(), "name is not valid") {
				logger.Warnf("reverse resolving address [%v] resulted in a skippable error [%s], skipping it", address, err.Error())
			} else {
				return fmt.Errorf("error could not reverse resolve address [%v]: %w", address, err)
			}
		}

		if reverseName != name {
			err = validateZnsName(client, reverseName, alreadyChecked)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func validateZnsName(client *ethclient.Client, name string, alreadyChecked *ZnsCheckedDictionary) error {
	if name == "" || name == ".eth" {
		return nil
	}
	// For now only .eth is supported other zns domains use different techniques and require and individual implementation
	if !strings.HasSuffix(name, ".eth") {
		name = fmt.Sprintf("%s.eth", name)
	}
	alreadyChecked.mux.Lock()
	if alreadyChecked.name[name] {
		alreadyChecked.mux.Unlock()
		return nil
	}
	alreadyChecked.name[name] = true
	alreadyChecked.mux.Unlock()

	startTime := time.Now()
	defer func() {
		metrics.TaskDuration.WithLabelValues("zns_validate_zns_name").Observe(time.Since(startTime).Seconds())
	}()

	nameHash, err := go_zns.NameHash(name)
	if err != nil {
		logger.Warnf("error could not hash name [%v]: %v -> removing zns entry", name, err)
		err = removeZnsName(client, name)
		if err != nil {
			return fmt.Errorf("error removing zns name [%v]: %w", name, err)
		}
		return nil
	}

	addr, err := go_zns.Resolve(client, name)
	if err != nil {
		if err.Error() == "unregistered name" ||
			err.Error() == "no address" ||
			err.Error() == "no resolver" ||
			err.Error() == "abi: attempting to unmarshall an empty string while arguments are expected" ||
			strings.Contains(err.Error(), "execution reverted") ||
			err.Error() == "invalid jump destination" ||
			err.Error() == "invalid opcode: INVALID" {
			// the given name is not available anymore or resolving it did not work properly => we can remove it from the db (if it is there)
			logger.WithField("error", err).WithField("name", name).Warnf("could not resolve name")
			err = removeZnsName(client, name)
			if err != nil {
				return fmt.Errorf("error removing zns name after resolve failed [%v]: %w", name, err)
			}
			return nil
		}
		return fmt.Errorf("error could not resolve name [%v]: %w", name, err)
	}

	// we need to get the main domain to get the expiration date
	parts := strings.Split(name, ".")
	mainName := strings.Join(parts[len(parts)-2:], ".")

	expires, err := GetZnsExpiration(client, mainName)
	if err != nil {
		return fmt.Errorf("error could not get zns expire date for [%v]: %w", name, err)
	}

	// znsName, err := go_zns.NewName(client, mainName)
	// if err != nil {
	// 	if strings.HasPrefix(err.Error(), "name is not valid") {
	// 		logger.WithField("error", err).WithField("name", name).Warnf("could not create name")
	// 		return nil
	// 	}
	// 	return fmt.Errorf("error could not create name via go_zns.NewName for [%v]: %w", name, err)
	// }
	// expires, err := znsName.Expires()
	// if err != nil {
	// 	return fmt.Errorf("error could not get zns expire date for [%v]: %w", name, err)
	// }

	isPrimary := false
	reverseName, err := go_zns.ReverseResolve(client, addr)
	if err != nil {
		if err.Error() == "not a resolver" || err.Error() == "no resolution" || err.Error() == "execution reverted" {
			logger.Warnf("reverse resolving address [%v] for name [%v] resulted in an error [%s], marking entry as not primary", addr, name, err.Error())
		} else {
			return fmt.Errorf("error could not reverse resolve address [%v]: %w", addr, err)
		}
	}
	if reverseName == name {
		isPrimary = true
	}

	_, err = WriterDb.Exec(`
	INSERT INTO zns (
		name_hash,
		zns_name,
		address,
		is_primary_name,
		valid_to)
	VALUES ($1, $2, $3, $4, $5)
	ON CONFLICT
		(name_hash)
	DO UPDATE SET
		zns_name = excluded.zns_name,
		address = excluded.address,
		is_primary_name = excluded.is_primary_name,
		valid_to = excluded.valid_to
	`, nameHash[:], name, addr.Bytes(), isPrimary, expires)
	if err != nil {
		if strings.Contains(fmt.Sprintf("%v", err), "invalid byte sequence") {
			logger.Warnf("could not insert zns name [%v]: %v", name, err)
			return nil
		}
		return fmt.Errorf("error writing zns data for name [%v]: %w", name, err)
	}

	logrus.WithFields(logrus.Fields{
		"name":        name,
		"address":     addr,
		"expires":     expires,
		"reverseName": reverseName,
	}).Infof("validated zns name")
	return nil
}

func GetZnsExpiration(client *ethclient.Client, name string) (time.Time, error) {
	startTime := time.Now()
	defer func() {
		metrics.TaskDuration.WithLabelValues("zns_get_expiration").Observe(time.Since(startTime).Seconds())
	}()

	normName, err := go_zns.NormaliseDomain(name)
	if err != nil {
		return time.Time{}, fmt.Errorf("error calling go_zns.NormaliseDomain: %w", err)
	}
	domain := go_zns.Domain(normName)
	label, err := go_zns.DomainPart(normName, 1)
	if err != nil {
		return time.Time{}, fmt.Errorf("error calling go_zns.DomainPart: %w", err)
	}
	registrar, err := go_zns.NewBaseRegistrar(client, domain)
	if err != nil {
		return time.Time{}, fmt.Errorf("error calling go_zns.NewBaseRegistrar: %w", err)
	}
	uqName, err := go_zns.UnqualifiedName(label, domain)
	if err != nil {
		return time.Time{}, fmt.Errorf("error calling go_zns.UnqualifiedName: %w", err)
	}
	labelHash, err := go_zns.LabelHash(uqName)
	if err != nil {
		return time.Time{}, fmt.Errorf("error calling go_zns.LabelHash: %w", err)
	}
	id := new(big.Int).SetBytes(labelHash[:])
	ts, err := registrar.Contract.NameExpires(&bind.CallOpts{}, id)
	if err != nil {
		return time.Time{}, err
	}
	return time.Unix(ts.Int64(), 0), nil
}

func GetAddressForZnsName(name string) (address *common.Address, err error) {
	addressBytes := []byte{}
	err = ReaderDb.Get(&addressBytes, `
	SELECT address
	FROM zns
	WHERE
		zns_name = $1 AND
		valid_to >= now()
	`, name)
	if err == nil && addressBytes != nil {
		add := common.BytesToAddress(addressBytes)
		address = &add
	}
	return address, err
}

func GetZnsNameForAddress(address common.Address) (name string, err error) {
	err = ReaderDb.Get(&name, `
	SELECT zns_name
	FROM zns
	WHERE
		address = $1 AND
		is_primary_name AND
		valid_to >= now()
	;`, address.Bytes())
	return name, err
}

func GetZnsNamesForAddress(addressMap map[string]string) error {
	if len(addressMap) == 0 {
		return nil
	}
	type pair struct {
		Address []byte `db:"address"`
		ZnsName string `db:"zns_name"`
	}
	dbAddresses := []pair{}
	addresses := make([][]byte, 0, len(addressMap))
	for add := range addressMap {
		addresses = append(addresses, []byte(add))
	}

	err := ReaderDb.Select(&dbAddresses, `
	SELECT address, zns_name
	FROM zns
	WHERE
		address = ANY($1) AND
		is_primary_name AND
		valid_to >= now()
	;`, addresses)
	if err != nil {
		return err
	}
	for _, foundling := range dbAddresses {
		addressMap[string(foundling.Address)] = foundling.ZnsName
	}
	return nil
}

func removeZnsName(client *ethclient.Client, name string) error {
	_, err := WriterDb.Exec(`
	DELETE FROM zns
	WHERE
		zns_name = $1
	;`, name)
	if err != nil && strings.Contains(fmt.Sprintf("%v", err), "invalid byte sequence") {
		logger.Warnf("could not delete zns name [%v]: %v", name, err)
		return nil
	} else if err != nil {
		return fmt.Errorf("error deleting zns name [%v]: %v", name, err)
	}
	logger.Infof("Zns name removed from db: %v", name)
	return nil
}
*/
