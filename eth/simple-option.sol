// SPDX-License-Identifier: BSD-3-Clause-Modification
pragma solidity ^0.8.3;

contract SimpleIpnsDirectory {
	event AddedEntry(bytes entry);

	bytes[] public entries_list;
	mapping(bytes => bool) public entries_set;
	string public description;

	constructor(string memory desc) {
		description = desc;
	}

	function addEntry(bytes calldata ipns_key) public {
		bool entry_exists = entries_set[ipns_key];
		require(!entry_exists, "IPNS entry already exists");

		entries_set[ipns_key] = true;
		entries_list.push(ipns_key);
		emit AddedEntry(ipns_key);
	}
}