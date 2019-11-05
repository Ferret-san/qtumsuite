// Copyright (c) 2014-2016 The btcsuite developers
// Use of this source code is governed by an ISC
// license that can be found in the LICENSE file.

package chaincfg

import (
	"time"

	"github.com/qtumproject/qtumsuite/chaincfg/chainhash"
	"github.com/qtumproject/qtumsuite/wire"
)

// genesisCoinbaseTx is the coinbase transaction for the genesis blocks for
// the main network, regression test network, and test network (version 3).
var genesisCoinbaseTx = wire.MsgTx{
	Version: 1,
	TxIn: []*wire.TxIn{
		{
			PreviousOutPoint: wire.OutPoint{
				Hash:  chainhash.Hash{},
				Index: 0xffffffff,
			},
			SignatureScript: []byte{
				0x00, 0x04, 0xbf, 0x91, 0x22, 0x1d, 0x01, 0x04,
				0x39, 0x53, 0x65, 0x70, 0x20, 0x30, 0x32, 0x2c,
				0x20, 0x32, 0x30, 0x31, 0x37, 0x20, 0x42, 0x69,
				0x74, 0x63, 0x6f, 0x69, 0x6e, 0x20, 0x62, 0x72,
				0x65, 0x61, 0x6b, 0x73, 0x20, 0x24, 0x35, 0x2c,
				0x30, 0x30, 0x30, 0x20, 0x69, 0x6e, 0x20, 0x6c,
				0x61, 0x74, 0x65, 0x73, 0x74, 0x20, 0x70, 0x72,
				0x69, 0x63, 0x65, 0x20, 0x66, 0x72, 0x65, 0x6e,
				0x7a, 0x79,
			},
			Sequence: 0xffffffff,
		},
	},
	TxOut: []*wire.TxOut{
		{
			Value: 0x12a05f200,
			PkScript: []byte{
				0x41, 0x04, 0x0d, 0x61, 0xd8, 0x65, 0x34, 0x48,
				0xc9, 0x87, 0x31, 0xee, 0x5f, 0xff, 0xd3, 0x03,
				0xc1, 0x5e, 0x71, 0xec, 0x20, 0x57, 0xb7, 0x7f,
				0x11, 0xab, 0x36, 0x01, 0x97, 0x97, 0x28, 0xcd,
				0xaf, 0xf2, 0xd6, 0x8a, 0xfb, 0xba, 0x14, 0xe4,
				0xfa, 0x0b, 0xc4, 0x4f, 0x20, 0x72, 0xb0, 0xb2,
				0x3e, 0xf6, 0x37, 0x17, 0xf8, 0xcd, 0xfb, 0xe5,
				0x8d, 0xcd, 0x33, 0xf3, 0x2b, 0x6a, 0xfe, 0x98,
				0x74, 0x1a, 0xac,
			},
		},
	},
	LockTime: 0,
}

// genesisHash is the hash of the first block in the block chain for the main
// network (genesis block).
var genesisHash = chainhash.Hash([chainhash.HashSize]byte{ // Make go vet happy.
	0x6c, 0x98, 0xed, 0x82, 0x55, 0x5e, 0xf9, 0xe3,
	0xd6, 0x34, 0x83, 0xd9, 0x21, 0xfa, 0x6c, 0x09,
	0xc3, 0xf8, 0xe6, 0x8c, 0xae, 0xf8, 0x80, 0x35,
	0x85, 0xf2, 0x3c, 0xf8, 0xae, 0x75, 0x00, 0x00,
})

// genesisMerkleRoot is the hash of the first transaction in the genesis block
// for the main network.
var genesisMerkleRoot = chainhash.Hash([chainhash.HashSize]byte{ // Make go vet happy.
	0x6d, 0xb9, 0x05, 0x14, 0x23, 0x82, 0x32, 0x4d,
	0xb4, 0x17, 0x76, 0x18, 0x91, 0xf2, 0xd2, 0xf3,
	0x55, 0xea, 0x92, 0xf2, 0x7a, 0xb0, 0xfc, 0x35,
	0xe5, 0x9e, 0x90, 0xb5, 0x0e, 0x05, 0x34, 0xed,
})

var genesisHashStateRoot = chainhash.Hash([chainhash.HashSize]byte{
	0xe9, 0x65, 0xff, 0xd0, 0x02, 0xcd, 0x6a, 0xd0,
	0xe2, 0xdc, 0x40, 0x2b, 0x80, 0x44, 0xde, 0x83,
	0x3e, 0x06, 0xb2, 0x31, 0x27, 0xea, 0x8c, 0x3d,
	0x80, 0xae, 0xc9, 0x14, 0x10, 0x77, 0x14, 0x95,
})

var genesisHashUTXORoot = chainhash.Hash([chainhash.HashSize]byte{
	0x56, 0xe8, 0x1f, 0x17, 0x1b, 0xcc, 0x55, 0xa6,
	0xff, 0x83, 0x45, 0xe6, 0x92, 0xc0, 0xf8, 0x6e,
	0x5b, 0x48, 0xe0, 0x1b, 0x99, 0x6c, 0xad, 0xc0,
	0x01, 0x62, 0x2f, 0xb5, 0xe3, 0x63, 0xb4, 0x21,
})

var genesisPrevoutStake = wire.OutPoint{*newHashFromStr("0000000000000000000000000000000000000000000000000000000000000000"), 0xffffffff}

// genesisBlock defines the genesis block of the block chain which serves as the
// public transaction ledger for the main network.
var genesisBlock = wire.MsgBlock{
	Header: wire.BlockHeader{
		Version:       1,
		PrevBlock:     chainhash.Hash{},         // 0000000000000000000000000000000000000000000000000000000000000000
		MerkleRoot:    genesisMerkleRoot,        // ed34050eb5909ee535fcb07af292ea55f3d2f291187617b44d3282231405b96d
		Timestamp:     time.Unix(0x59afd2f5, 0), // 2009-01-03 18:15:05 +0000 UTC
		Bits:          0x1f00ffff,               // 486604799 [00000000ffff0000000000000000000000000000000000000000000000000000]
		Nonce:         0x007a78f9,               // 8026361
		HashStateRoot: genesisHashStateRoot,
		HashUTXORoot:  genesisHashUTXORoot,
		PrevoutStake:  genesisPrevoutStake,
		BlockSig:      make([]byte, 0),
	},
	Transactions: []*wire.MsgTx{&genesisCoinbaseTx},
}

// regTestGenesisHash is the hash of the first block in the block chain for the
// regression test network (genesis block).
var regTestGenesisHash = chainhash.Hash([chainhash.HashSize]byte{ // Make go vet happy.
	0x43, 0xe9, 0xda, 0xfa, 0x2e, 0x97, 0xfb, 0x58,
	0x9a, 0xee, 0xb7, 0x78, 0x72, 0x8a, 0x3e, 0x36,
	0x94, 0x29, 0x33, 0x26, 0x89, 0x7d, 0xc3, 0xef,
	0x44, 0x0b, 0xac, 0x02, 0xb4, 0xd5, 0x5e, 0x66,
})

// regTestGenesisMerkleRoot is the hash of the first transaction in the genesis
// block for the regression test network.  It is the same as the merkle root for
// the main network.
var regTestGenesisMerkleRoot = genesisMerkleRoot

// regTestGenesisBlock defines the genesis block of the block chain which serves
// as the public transaction ledger for the regression test network.
var regTestGenesisBlock = wire.MsgBlock{
	Header: wire.BlockHeader{
		Version:       1,
		PrevBlock:     chainhash.Hash{},         // 0000000000000000000000000000000000000000000000000000000000000000
		MerkleRoot:    regTestGenesisMerkleRoot, // 4a5e1e4baab89f3a32518a88c31bc87f618f76673e2cc77ab2127b7afdeda33b
		Timestamp:     time.Unix(1504695029, 0), // 2011-02-02 23:16:42 +0000 UTC
		Bits:          0x207fffff,               // 545259519 [7fffff0000000000000000000000000000000000000000000000000000000000]
		Nonce:         17,
		HashStateRoot: genesisHashStateRoot,
		HashUTXORoot:  genesisHashUTXORoot,
		PrevoutStake:  genesisPrevoutStake,
		BlockSig:      make([]byte, 0),
	},
	Transactions: []*wire.MsgTx{&genesisCoinbaseTx},
}

// testNet3GenesisHash is the hash of the first block in the block chain for the
// test network (version 3).
var testNet3GenesisHash = chainhash.Hash([chainhash.HashSize]byte{ // Make go vet happy.
	0x22, 0x42, 0xe1, 0xf5, 0x51, 0xba, 0xb2, 0xfe,
	0x66, 0xad, 0x2a, 0x97, 0x17, 0x86, 0x82, 0x3f,
	0x4d, 0x59, 0x20, 0x92, 0x2f, 0x0d, 0xca, 0x84,
	0x06, 0x5c, 0x21, 0xee, 0x03, 0xe8, 0x00, 0x00,
})

// testNet3GenesisMerkleRoot is the hash of the first transaction in the genesis
// block for the test network (version 3).  It is the same as the merkle root
// for the main network.
var testNet3GenesisMerkleRoot = genesisMerkleRoot

// testNet3GenesisBlock defines the genesis block of the block chain which
// serves as the public transaction ledger for the test network (version 3).
var testNet3GenesisBlock = wire.MsgBlock{
	Header: wire.BlockHeader{
		Version:       1,
		PrevBlock:     chainhash.Hash{},          // 0000000000000000000000000000000000000000000000000000000000000000
		MerkleRoot:    testNet3GenesisMerkleRoot, // 4a5e1e4baab89f3a32518a88c31bc87f618f76673e2cc77ab2127b7afdeda33b
		Timestamp:     time.Unix(1504695029, 0),  // 2011-02-02 23:16:42 +0000 UTC
		Bits:          0x1f00ffff,                // 486604799 [00000000ffff0000000000000000000000000000000000000000000000000000]
		Nonce:         7349697,                   // 414098458
		HashStateRoot: genesisHashStateRoot,
		HashUTXORoot:  genesisHashUTXORoot,
		PrevoutStake:  genesisPrevoutStake,
		BlockSig:      make([]byte, 0),
	},
	Transactions: []*wire.MsgTx{&genesisCoinbaseTx},
}
