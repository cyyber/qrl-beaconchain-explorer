#! /bin/bash
set -e

clean_up () {
    ARG=$?
    rm -rf /tmp/full_withdrawal
    exit $ARG
} 
trap clean_up EXIT

while getopts b:i:m: flag
do
    case "${flag}" in
        b) bn_endpoint=${OPTARG};;
        w) wallet_dir=${OPTARG};;
    esac
done
echo "Wallet Dir: $wallet_dir";
echo "BN Endpoint: $bn_endpoint";

echo "submitting exit message"
prysmctl validator exit \
    --wallet-dir="$wallet_dir" \
    --beacon-rpc-provider="$bn_endpoint"