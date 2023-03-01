#!/bin/sh

abigen --abi ../../contracts/build/POGToken.abi --pkg contract --type POGToken --out POGToken.go --bin ../../contracts/build/POGToken.bin
