package goTezos

/*
Author: DefinitelyNotAGoat/MagicAglet
Version: 0.0.1
Description: This file contains regex parsing for parsing RPC responses.
License: MIT
*/

import "regexp"

//Regular Expressions Used
var (
	reSnapShotNumber        = regexp.MustCompile(`([0-9]+)`)
	reGetBlockLevelHead     = regexp.MustCompile(`level":([0-9]+),"proto"`)
	reGetHash               = regexp.MustCompile(`"hash":"([0-9a-zA-Z]+)","header"`)
	reGetRandomSeed         = regexp.MustCompile(`"random_seed":"([0-9a-zA-Z]+)"`)
	reGetRollSnapShot       = regexp.MustCompile(`"roll_snapshot":([0-9]+)`)
	reGetBalance            = regexp.MustCompile(`([0-9.]+)`)
	reGetBalanceForSnapshot = regexp.MustCompile(`"([0-9]+)"`)
	reDelegatedContracts    = regexp.MustCompile(`"([A-Z0-9a-z]+)"`)
	reListKownAddresses     = regexp.MustCompile(`([a-zA-z_0-9]+): ([a-zA-z0-9]+) \(([a-zA-Z]+)`) //Group 1 = alias, group 2 = address, group 3 = encryption
	reRewards               = regexp.MustCompile(`"rewards":"([0-9]+)"`)
)
