// SPDX-License-Identifier: BSD-3-Clause-Modification
pragma solidity ^0.8.4;

/// @title A simple dictionary of hash equivalences based on users contributions and votes.
/// @author G. Queiroz <gabrieljvnq@gmail.com>
/// @custom:experimental This is an experimental contract.
contract HashDictionary {
	event AddedEntry(Entry);
	enum VoteChoice { None, Right, Wrong }

	struct HashPair {
		bytes phash;
		bytes chash;
	}

	struct Entry {
		HashPair hpair;
		int votes_right;
		int votes_wrong;
	}

	Entry[] public entries_list;
	string public phash_alg;
	string public chash_alg;
	mapping(bytes => uint) public hashpair2idx;
	mapping(bytes => uint[]) public phash2idx;
	mapping(bytes => uint[]) public chash2idx;
	mapping(bytes => VoteChoice) votes;

	/// @notice Initializes the contract.
    /// @param alg_phash The name of the (usually perceptual) hash algorithm. Example: "pHash"
    /// @param alg_chash The name of the (usually cryptographic) hash algorithm. Example: "multihash"
	constructor(string memory alg_phash, string memory alg_chash) {
		require(bytes(alg_phash).length != 0, "phash algorithm must be specified");
		require(bytes(alg_chash).length != 0, "chash algorithm must be specified");
		phash_alg = alg_phash;
		chash_alg = alg_chash;

		// Add a dummy entry so that the index 0 means NULL
		entries_list.push(Entry({
			hpair: HashPair({
				phash: "",
				chash: ""
			}),
			votes_right: 0,
			votes_wrong: 0
		}));
	}

	/// @notice Converts a HashPair to a bytes array so it can be used as a mapping key.
	function hpair2bytes(HashPair calldata hpair) public pure returns (bytes memory) {
		return bytes.concat(hpair.phash, hex"00", hpair.chash);
	}

	/// @notice Converts a pair of HashPair and address to a bytes array so it can be used as a mapping key.
	function getVoteKey(HashPair calldata hpair, address voter) public pure returns (bytes memory) {
		bytes memory voter_addr = abi.encodePacked(voter);
		return bytes.concat(hpair2bytes(hpair), hex"00", voter_addr);
	}

	/// @notice Gets the index that hold the HashPair entry.
	function getEntryIdx(HashPair calldata hpair) public view returns (uint) {
		return hashpair2idx[hpair2bytes(hpair)];
	}

	/// @notice Gets the HashPair entry associated with a particular phash.
	function getEntriesByPHash(bytes calldata hash) public view returns (Entry[] memory) {
		return getEntriesByIdxList(phash2idx[hash]);
	}

	/// @notice Gets the HashPair entry associated with a particular chash.
	function getEntriesByCHash(bytes calldata hash) public view returns (Entry[] memory) {
		return getEntriesByIdxList(chash2idx[hash]);
	}

	function getEntriesByIdxList(uint[] storage idx_list) internal view returns (Entry[] memory) {
		Entry[] memory ans = new Entry[](idx_list.length);

		for (uint i = 0; i < idx_list.length; i++) {
			ans[i] = entries_list[idx_list[i]];
		}

		return ans;
	}

	/// @notice Adds an equivalence entry and votes for it.
    /// @param hpair The entry being added.
	function addEntry(HashPair calldata hpair) public {
		// No duplicates
		uint entry_idx = getEntryIdx(hpair);
		require(entry_idx == 0, "entry already exists");

		hashpair2idx[hpair2bytes(hpair)] = entries_list.length;
		// We know we won't make duplicates because we already checked for it
		phash2idx[hpair.phash].push(entries_list.length);
		chash2idx[hpair.chash].push(entries_list.length);
		Entry memory new_entry = Entry({
			hpair: hpair,
			votes_right: 1,
			votes_wrong: 0
		});
		entries_list.push(new_entry);

		bytes memory vote_key = getVoteKey(hpair, msg.sender);
		votes[vote_key] = VoteChoice.Right;

		emit AddedEntry(new_entry);
	}

	/// @notice Votes on the accuracy of an entry.
    /// @param hpair The entry being voted on.
    /// @param new_vote The user's vote choice.
	function voteEntry(HashPair calldata hpair, VoteChoice new_vote) public {
		uint entry_idx = getEntryIdx(hpair);
		if (entry_idx == 0) {
			addEntry(hpair);
			return;
		}

		bytes memory vote_key = getVoteKey(hpair, msg.sender);
		VoteChoice old_vote = votes[vote_key];
		votes[vote_key] = new_vote;
		if (old_vote == VoteChoice.Right) {
			entries_list[entry_idx].votes_right--;
		}
		if (old_vote == VoteChoice.Wrong) {
			entries_list[entry_idx].votes_wrong--;
		}
		if (new_vote == VoteChoice.Right) {
			entries_list[entry_idx].votes_right++;
		}
		if (new_vote == VoteChoice.Wrong) {
			entries_list[entry_idx].votes_wrong++;
		}
	}
}