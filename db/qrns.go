package db

// TODO(now.youtrack.cloud/issue/TZB-1)
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

	"github.com/sirupsen/logrus"
	"github.com/theQRL/go-zond/qrlclient"
	"github.com/theQRL/qrl-beaconchain-explorer/metrics"
	"github.com/theQRL/qrl-beaconchain-explorer/types"
	"github.com/theQRL/qrl-beaconchain-explorer/utils"

	gcp_bigtable "cloud.google.com/go/bigtable"
	"golang.org/x/sync/errgroup"

	"github.com/coocood/freecache"
	"github.com/theQRL/go-zond/accounts/abi/bind"
	"github.com/theQRL/go-zond/common"

	zond_types "github.com/theQRL/go-zond/core/types"
	go_qrns "github.com/wealdtech/go-ens/v3"
)

// TransformQrnsNameRegistered accepts an execution block and creates bigtable mutations for QRNS Name events.
// It transforms the logs contained within a block and indexes qrns relevant transactions and tags changes (to be verified from the node in a separate process)
// ==================================================
//
// It indexes transactions
//
// - by hashed qrns name
// Row:    <chainID>:QRNS:I:H:<nameHash>:<txHash>
// Family: f
// Column: nil
// Cell:   nil
// Example scan: "5:QRNS:I:H:4ae569dd0aa2f6e9207e41423c956d0d27cbc376a499ee8d90fe1d84489ae9d1:e627ae94bd16eb1ed8774cd4003fc25625159f13f8a2612cc1c7f8d2ab11b1d7"
//
// - by address
// Row:    <chainID>:QRNS:I:A:<address>:<txHash>
// Family: f
// Column: nil
// Cell:   nil
// Example scan: "5:QRNS:I:A:05579fadcf7cc6544f7aa018a2726c85251600c5:e627ae94bd16eb1ed8774cd4003fc25625159f13f8a2612cc1c7f8d2ab11b1d7"
//
// ==================================================
//
// Track for later verification via the node ("set dirty")
//
// - by name
// Row:    <chainID>:QRNS:V:N:<name>
// Family: f
// Column: nil
// Cell:   nil
// Example scan: "5:QRNS:V:N:somename"
//
// - by name hash
// Row:    <chainID>:QRNS:V:H:<nameHash>
// Family: f
// Column: nil
// Cell:   nil
// Example scan: "5:QRNS:V:H:6f5d9cc23e60abe836401b4fd386ec9280a1f671d47d9bf3ec75dab76380d845"
//
// - by address
// Row:    <chainID>:QRNS:V:A:<address>
// Family: f
// Column: nil
// Cell:   nil
// Example scan: "5:QRNS:V:A:27234cb8734d5b1fac0521c6f5dc5aebc6e839b6"
//
// ==================================================

func (bigtable *Bigtable) TransformQrnsNameRegistered(blk *types.ExecutionBlock, cache *freecache.Cache) (bulkData *types.BulkMutations, bulkMetadataUpdates *types.BulkMutations, err error) {
	startTime := time.Now()
	defer func() {
		metrics.TaskDuration.WithLabelValues("bt_transform_qrns").Observe(time.Since(startTime).Seconds())
	}()

	var qrnsCrontractAddresses map[string]string
	switch bigtable.chainId {
	case "1":
		qrnsCrontractAddresses = qrnsContracts.QRNSCrontractAddressesEthereum
	default:
		return nil, nil, nil
	}

	bulkData = &types.BulkMutations{}
	bulkMetadataUpdates = &types.BulkMutations{}
	keys := make(map[string]bool)
	ethLog := qrl_types.Log{}

	for i, tx := range blk.GetTransactions() {
		if i >= TX_PER_BLOCK_LIMIT {
			return nil, nil, fmt.Errorf("unexpected number of transactions in block expected at most %d but got: %v, tx: %x", TX_PER_BLOCK_LIMIT-1, i, tx.GetHash())
		}
		for j, log := range tx.GetLogs() {
			if j >= ITX_PER_TX_LIMIT {
				return nil, nil, fmt.Errorf("unexpected number of logs in block expected at most %d but got: %v tx: %x", ITX_PER_TX_LIMIT-1, j, tx.GetHash())
			}
			qrnsContract := qrnsCrontractAddresses[common.BytesToAddress(log.Address).String()]

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
					"block":        blk.GetNumber(),
					"tx":           tx.GetHash(),
					"logIndex":     j,
					"qrnsContract": qrnsContract,
				}

				if qrnsContract == "Registry" {
					if bytes.Equal(lTopic, qrnsContracts.QRNSRegistryParsedABI.Events["NewResolver"].ID.Bytes()) {
						logFields["event"] = "NewResolver"
						r := &qrnsContracts.QRNSRegistryNewResolver{}
						err = qrnsContracts.QRNSRegistryContract.UnpackLog(r, "NewResolver", ethLog)
						if err != nil {
							utils.LogWarn(err, "error unpacking qrns-log", 0, logFields)
							continue
						}
						keys[fmt.Sprintf("%s:QRNS:V:H:%x", bigtable.chainId, r.Node)] = true
					} else if bytes.Equal(lTopic, qrnsContracts.QRNSRegistryParsedABI.Events["NewOwner"].ID.Bytes()) {
						logFields["event"] = "NewOwner"
						r := &qrnsContracts.QRNSRegistryNewOwner{}
						err = qrnsContracts.QRNSRegistryContract.UnpackLog(r, "NewOwner", ethLog)
						if err != nil {
							utils.LogWarn(err, "error unpacking qrns-log", 0, logFields)
							continue
						}
						keys[fmt.Sprintf("%s:QRNS:V:A:%x", bigtable.chainId, r.Owner)] = true
					} else if bytes.Equal(lTopic, qrnsContracts.QRNSRegistryParsedABI.Events["NewTTL"].ID.Bytes()) {
						logFields["event"] = "NewTTL"
						r := &qrnsContracts.QRNSRegistryNewTTL{}
						err = qrnsContracts.QRNSRegistryContract.UnpackLog(r, "NewTTL", ethLog)
						if err != nil {
							utils.LogWarn(err, "error unpacking qrns-log", 0, logFields)
							continue
						}
						keys[fmt.Sprintf("%s:QRNS:V:H:%x", bigtable.chainId, r.Node)] = true
					}
				} else if qrnsContract == "QRLRegistrarController" {
					if bytes.Equal(lTopic, qrnsContracts.QRNSQRLRegistrarControllerParsedABI.Events["NameRegistered"].ID.Bytes()) {
						logFields["event"] = "NameRegistered"
						r := &qrnsContracts.QRNSQRLRegistrarControllerNameRegistered{}
						err = qrnsContracts.QRNSQRLRegistrarControllerContract.UnpackLog(r, "NameRegistered", ethLog)
						if err != nil {
							utils.LogWarn(err, "error unpacking qrns-log", 0, logFields)
							continue
						}
						keys[fmt.Sprintf("%s:QRNS:V:N:%s", bigtable.chainId, r.Name)] = true
						keys[fmt.Sprintf("%s:QRNS:V:A:%x", bigtable.chainId, r.Owner)] = true
					} else if bytes.Equal(lTopic, qrnsContracts.QRNSQRLRegistrarControllerParsedABI.Events["NameRenewed"].ID.Bytes()) {
						logFields["event"] = "NameRenewed"
						r := &qrnsContracts.QRNSQRLRegistrarControllerNameRenewed{}
						err = qrnsContracts.QRNSQRLRegistrarControllerContract.UnpackLog(r, "NameRenewed", ethLog)
						if err != nil {
							utils.LogWarn(err, "error unpacking qrns-log", 0, logFields)
							continue
						}
						keys[fmt.Sprintf("%s:QRNS:V:N:%s", bigtable.chainId, r.Name)] = true
					}
				} else if qrnsContract == "OldQrnsRegistrarController" {
					if bytes.Equal(lTopic, qrnsContracts.QRNSOldRegistrarControllerParsedABI.Events["NameRegistered"].ID.Bytes()) {
						logFields["event"] = "NameRegistered"
						r := &qrnsContracts.QRNSOldRegistrarControllerNameRegistered{}
						err = qrnsContracts.QRNSOldRegistrarControllerContract.UnpackLog(r, "NameRegistered", ethLog)
						if err != nil {
							utils.LogWarn(err, "error unpacking qrns-log", 0, logFields)
							continue
						}
						if err = verifyName(r.Name); err != nil {
							utils.LogWarn(err, "error verifying qrns-name", 0, logFields)
							continue
						}
						keys[fmt.Sprintf("%s:QRNS:V:N:%s", bigtable.chainId, r.Name)] = true
						keys[fmt.Sprintf("%s:QRNS:V:A:%x", bigtable.chainId, r.Owner)] = true
					} else if bytes.Equal(lTopic, qrnsContracts.QRNSOldRegistrarControllerParsedABI.Events["NameRenewed"].ID.Bytes()) {
						logFields["event"] = "NameRenewed"
						r := &qrnsContracts.QRNSOldRegistrarControllerNameRenewed{}
						err = qrnsContracts.QRNSOldRegistrarControllerContract.UnpackLog(r, "NameRenewed", ethLog)
						if err != nil {
							utils.LogWarn(err, "error unpacking qrns-log", 0, logFields)
							continue
						}
						if err = verifyName(r.Name); err != nil {
							utils.LogWarn(err, "error verifying qrns-name", 0, logFields)
							continue
						}
						keys[fmt.Sprintf("%s:QRNS:V:N:%s", bigtable.chainId, r.Name)] = true
					}
				} else {
					if bytes.Equal(lTopic, qrnsContracts.QRNSPublicResolverParsedABI.Events["NameChanged"].ID.Bytes()) {
						logFields["event"] = "NameChanged"
						r := &qrnsContracts.QRNSPublicResolverNameChanged{}
						err = qrnsContracts.QRNSPublicResolverContract.UnpackLog(r, "NameChanged", ethLog)
						if err != nil {
							utils.LogWarn(err, "error unpacking qrns-log", 0, logFields)
							continue
						}
						if err = verifyName(r.Name); err != nil {
							utils.LogWarn(err, "error verifying qrns-name", 0, logFields)
							continue
						}
						keys[fmt.Sprintf("%s:QRNS:V:N:%s", bigtable.chainId, r.Name)] = true
					} else if bytes.Equal(lTopic, qrnsContracts.QRNSPublicResolverParsedABI.Events["AddressChanged"].ID.Bytes()) {
						logFields["event"] = "AddressChanged"
						r := &qrnsContracts.QRNSPublicResolverAddressChanged{}
						err = qrnsContracts.QRNSPublicResolverContract.UnpackLog(r, "AddressChanged", ethLog)
						if err != nil {
							utils.LogWarn(err, "error unpacking qrns-log", 0, logFields)
							continue
						}
						keys[fmt.Sprintf("%s:QRNS:V:H:%x", bigtable.chainId, r.Node)] = true
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

type QrnsCheckedDictionary struct {
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

func (bigtable *Bigtable) ImportQrnsUpdates(client *qrlclient.Client, readBatchSize int64) error {
	startTime := time.Now()
	defer func() {
		metrics.TaskDuration.WithLabelValues("bt_import_qrns_updates").Observe(time.Since(startTime).Seconds())
	}()

	key := fmt.Sprintf("%s:QRNS:V", bigtable.chainId)

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
		logger.Info("No QRNS entries to validate")
		return nil
	}

	logger.Infof("Validating %v QRNS entries", len(keys))
	alreadyChecked := QrnsCheckedDictionary{
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
		logger.Infof("Batching QRNS entries %v:%v of %v", i, to, total)

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
						qrns_name
					FROM qrns
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
					err := validateQrnsName(client, name, &alreadyChecked)
					if err != nil {
						return fmt.Errorf("error validating new name [%v]: %w", name, err)
					}
				} else if address != nil {
					err := validateQrnsAddress(client, *address, &alreadyChecked)
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

	logger.WithField("updates", total).Info("Import of QRNS updates completed")
	return nil
}

func validateQrnsAddress(client *qrlclient.Client, address common.Address, alreadyChecked *QrnsCheckedDictionary) error {
	alreadyChecked.mux.Lock()
	if alreadyChecked.address[address] {
		alreadyChecked.mux.Unlock()
		return nil
	}
	alreadyChecked.address[address] = true
	alreadyChecked.mux.Unlock()

	names := []string{}
	err := ReaderDb.Select(&names, `SELECT qrns_name FROM qrns WHERE address = $1 AND is_primary_name AND valid_to >= now()`, address.Bytes())
	if err != nil && err != sql.ErrNoRows {
		return err
	}

	for _, name := range names {
		if name != "" {
			err = validateQrnsName(client, name, alreadyChecked)
			if err != nil {
				return err
			}
		}
		reverseName, err := go_qrns.ReverseResolve(client, address)
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
			err = validateQrnsName(client, reverseName, alreadyChecked)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func validateQrnsName(client *qrlclient.Client, name string, alreadyChecked *QrnsCheckedDictionary) error {
	if name == "" || name == ".qrl" {
		return nil
	}
	// For now only .qrl is supported other qrns domains use different techniques and require and individual implementation
	if !strings.HasSuffix(name, ".qrl") {
		name = fmt.Sprintf("%s.qrl", name)
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
		metrics.TaskDuration.WithLabelValues("qrns_validate_qrns_name").Observe(time.Since(startTime).Seconds())
	}()

	nameHash, err := go_qrns.NameHash(name)
	if err != nil {
		logger.Warnf("error could not hash name [%v]: %v -> removing qrns entry", name, err)
		err = removeQrnsName(client, name)
		if err != nil {
			return fmt.Errorf("error removing qrns name [%v]: %w", name, err)
		}
		return nil
	}

	addr, err := go_qrns.Resolve(client, name)
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
			err = removeQrnsName(client, name)
			if err != nil {
				return fmt.Errorf("error removing qrns name after resolve failed [%v]: %w", name, err)
			}
			return nil
		}
		return fmt.Errorf("error could not resolve name [%v]: %w", name, err)
	}

	// we need to get the main domain to get the expiration date
	parts := strings.Split(name, ".")
	mainName := strings.Join(parts[len(parts)-2:], ".")

	expires, err := GetQrnsExpiration(client, mainName)
	if err != nil {
		return fmt.Errorf("error could not get qrns expire date for [%v]: %w", name, err)
	}

	// qrnsName, err := go_qrns.NewName(client, mainName)
	// if err != nil {
	// 	if strings.HasPrefix(err.Error(), "name is not valid") {
	// 		logger.WithField("error", err).WithField("name", name).Warnf("could not create name")
	// 		return nil
	// 	}
	// 	return fmt.Errorf("error could not create name via go_qrns.NewName for [%v]: %w", name, err)
	// }
	// expires, err := qrnsName.Expires()
	// if err != nil {
	// 	return fmt.Errorf("error could not get qrns expire date for [%v]: %w", name, err)
	// }

	isPrimary := false
	reverseName, err := go_qrns.ReverseResolve(client, addr)
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
	INSERT INTO qrns (
		name_hash,
		qrns_name,
		address,
		is_primary_name,
		valid_to)
	VALUES ($1, $2, $3, $4, $5)
	ON CONFLICT
		(name_hash)
	DO UPDATE SET
		qrns_name = excluded.qrns_name,
		address = excluded.address,
		is_primary_name = excluded.is_primary_name,
		valid_to = excluded.valid_to
	`, nameHash[:], name, addr.Bytes(), isPrimary, expires)
	if err != nil {
		if strings.Contains(fmt.Sprintf("%v", err), "invalid byte sequence") {
			logger.Warnf("could not insert qrns name [%v]: %v", name, err)
			return nil
		}
		return fmt.Errorf("error writing qrns data for name [%v]: %w", name, err)
	}

	logrus.WithFields(logrus.Fields{
		"name":        name,
		"address":     addr,
		"expires":     expires,
		"reverseName": reverseName,
	}).Infof("validated qrns name")
	return nil
}

func GetQrnsExpiration(client *qrlclient.Client, name string) (time.Time, error) {
	startTime := time.Now()
	defer func() {
		metrics.TaskDuration.WithLabelValues("qrns_get_expiration").Observe(time.Since(startTime).Seconds())
	}()

	normName, err := go_qrns.NormaliseDomain(name)
	if err != nil {
		return time.Time{}, fmt.Errorf("error calling go_qrns.NormaliseDomain: %w", err)
	}
	domain := go_qrns.Domain(normName)
	label, err := go_qrns.DomainPart(normName, 1)
	if err != nil {
		return time.Time{}, fmt.Errorf("error calling go_qrns.DomainPart: %w", err)
	}
	registrar, err := go_qrns.NewBaseRegistrar(client, domain)
	if err != nil {
		return time.Time{}, fmt.Errorf("error calling go_qrns.NewBaseRegistrar: %w", err)
	}
	uqName, err := go_qrns.UnqualifiedName(label, domain)
	if err != nil {
		return time.Time{}, fmt.Errorf("error calling go_qrns.UnqualifiedName: %w", err)
	}
	labelHash, err := go_qrns.LabelHash(uqName)
	if err != nil {
		return time.Time{}, fmt.Errorf("error calling go_qrns.LabelHash: %w", err)
	}
	id := new(big.Int).SetBytes(labelHash[:])
	ts, err := registrar.Contract.NameExpires(&bind.CallOpts{}, id)
	if err != nil {
		return time.Time{}, err
	}
	return time.Unix(ts.Int64(), 0), nil
}

func GetAddressForQrnsName(name string) (address *common.Address, err error) {
	addressBytes := []byte{}
	err = ReaderDb.Get(&addressBytes, `
	SELECT address
	FROM qrns
	WHERE
		qrns_name = $1 AND
		valid_to >= now()
	`, name)
	if err == nil && addressBytes != nil {
		add := common.BytesToAddress(addressBytes)
		address = &add
	}
	return address, err
}

func GetQrnsNameForAddress(address common.Address) (name string, err error) {
	err = ReaderDb.Get(&name, `
	SELECT qrns_name
	FROM qrns
	WHERE
		address = $1 AND
		is_primary_name AND
		valid_to >= now()
	;`, address.Bytes())
	return name, err
}

func GetQrnsNamesForAddress(addressMap map[string]string) error {
	if len(addressMap) == 0 {
		return nil
	}
	type pair struct {
		Address  []byte `db:"address"`
		QrnsName string `db:"qrns_name"`
	}
	dbAddresses := []pair{}
	addresses := make([][]byte, 0, len(addressMap))
	for add := range addressMap {
		addresses = append(addresses, []byte(add))
	}

	err := ReaderDb.Select(&dbAddresses, `
	SELECT address, qrns_name
	FROM qrns
	WHERE
		address = ANY($1) AND
		is_primary_name AND
		valid_to >= now()
	;`, addresses)
	if err != nil {
		return err
	}
	for _, foundling := range dbAddresses {
		addressMap[string(foundling.Address)] = foundling.QrnsName
	}
	return nil
}

func removeQrnsName(client *qrlclient.Client, name string) error {
	_, err := WriterDb.Exec(`
	DELETE FROM qrns
	WHERE
		qrns_name = $1
	;`, name)
	if err != nil && strings.Contains(fmt.Sprintf("%v", err), "invalid byte sequence") {
		logger.Warnf("could not delete qrns name [%v]: %v", name, err)
		return nil
	} else if err != nil {
		return fmt.Errorf("error deleting qrns name [%v]: %v", name, err)
	}
	logger.Infof("Qrns name removed from db: %v", name)
	return nil
}
*/
