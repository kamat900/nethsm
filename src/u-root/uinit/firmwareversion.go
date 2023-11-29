// Copyright 2023 - 2023, Nitrokey GmbH
// SPDX-License-Identifier: EUPL-1.2

// tpm.go contains TPM-related functions used to provision/retrieve and delete
// the "Device Key" stored in the TPM.
package main

var firmwareVersions = map[string]string{
	"2eeb3fbb4a1e4533ab7246d05049a7676d4a378d62426976c32774263d945806": "0.9-devel",
	"db89554134cbbc1f54e625a5df8e175a1f5189a3bfdffdfe249640a99a4bbbee": "1.0-devel",
	"6164cc90d15caca5da0bbcba579c438b1e274dfb979b8aeeb5a497b2e4ab2e69": "1.0-prod",
}

const firmwarePCRIdx = 2

func getFirmwareVersion(pcr map[int]string) string {
	version, ok := firmwareVersions[pcr[firmwarePCRIdx]]
	if ok {
		return version
	}
	return "unknown"
}
