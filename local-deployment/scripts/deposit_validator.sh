#! /bin/bash
set -e

clean_up () {
    ARG=$?
    rm -rf /tmp/deposit
    exit $ARG
} 
trap clean_up EXIT

while getopts a:b:e:i:m: flag
do
    case "${flag}" in
        b) bn_endpoint=${OPTARG};;
        e) el_endpoint=${OPTARG};;
        i) index=${OPTARG};;
        m) mnemonic=${OPTARG};;
    esac
done

echo "Validator Index: $index";
echo "Mnemonic: $mnemonic";
echo "BN Endpoint: $bn_endpoint";
echo "EL Endpoint: $el_endpoint";

mkdir -p /tmp/deposit

deposit_contract_address=$(curl -s $bn_endpoint/zond/v1/config/spec | jq -r '.data.DEPOSIT_CONTRACT_ADDRESS')

deposit new-seed \
    --validator-start-index="$index" \
    --num-validators=1 \
    --folder="/tmp/deposit" \
    --chain-name="mainnet" \
    --execution-address="$deposit_contract_address" \
    --mnemonic="$mnemonic"

echo "" > /tmp/deposit/staking_wallet.seed

deposit submit \
    --validator-keys-dir=/tmp/deposit/validator_keys \
    --zond-seed-file=/tmp/deposit/staking_wallet.seed \
    --http-web3provider="$el_endpoint" \
    --deposit-contract="$deposit_contract_address"

exit;
rm -rf /tmp/deposit