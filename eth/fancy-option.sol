// SPDX-License-Identifier: BSD-3-Clause-Modification
pragma solidity ^0.8.3;

contract HashDictionary {
	event AddedEntry(bytes entry);
	enum VoteChoice { None, Right, Wrong }

	struct HashPair {
		bytes phash;
		bytes chash;
		uint8 chash_interp; // 1 - chash is cryptographic hash, 2 - chash is IPFS CID
	}

	struct Entry {
		HashPair hpair;
		int votes_right;
		int votes_wrong;
	}

	Entry[] public entries_list;
	mapping(bytes => uint) public hashpair2idx;
	mapping(bytes => uint) public phash2idx;
	mapping(bytes => uint) public chash2idx;
	mapping(address => VoteChoice) votes;

	constructor() {
		entries_list.push(sha3(Entry()));
	}

	function getEntryIdx(HashPair calldata hpair) public returns (uint) {
		return hashpair2idx[sha3(hpair)];
	}

	function getEntryByPHash(bytes calldata hash) public {
		uint entry_idx = phash2idx[hash];
		require(entry_idx != 0, "entry does not exists");

		// ...
	}

	function getEntryByCHash(bytes calldata hash) public {
		uint entry_idx = chash2idx[hash];
		require(entry_idx != 0, "entry does not exists");

		// ...
	}

	function addEntry(HashPair calldata hpair) public {
		uint entry_idx = getEntryIdx(hpair);
		require(entry_idx == 0, "entry already exists");

		// ...
	}

	function voteEntry(HashPair calldata hpair, VoteChoice vote_val) public {
		uint entry_idx = getEntryIdx(hpair);
		require(entry_idx != 0, "entry does not exists");
		require(vote_val != 0, "vote cannot be null");

		// ...
	}
}