// Copyright (c) 2013-2016 The btcsuite developers
// Use of this source code is governed by an ISC
// license that can be found in the LICENSE file.

package wire

import (
	"bytes"
	"reflect"
	"testing"
	"time"

	"github.com/davecgh/go-spew/spew"
)

// TestBlockHeader tests the BlockHeader API.
func TestBlockHeader(t *testing.T) {
	nonce64, err := RandomUint64()
	if err != nil {
		t.Errorf("RandomUint64: Error generating nonce: %v", err)
	}
	nonce := uint32(nonce64)

	hash := mainNetGenesisHash
	merkleHash := mainNetGenesisMerkleRoot
	bits := uint32(0x1d00ffff)
	bh := NewBlockHeader(1, &hash, &merkleHash, bits, nonce)

	// Ensure we get the same data back out.
	if !bh.PrevBlock.IsEqual(&hash) {
		t.Errorf("NewBlockHeader: wrong prev hash - got %v, want %v",
			spew.Sprint(bh.PrevBlock), spew.Sprint(hash))
	}
	if !bh.MerkleRoot.IsEqual(&merkleHash) {
		t.Errorf("NewBlockHeader: wrong merkle root - got %v, want %v",
			spew.Sprint(bh.MerkleRoot), spew.Sprint(merkleHash))
	}
	if bh.Bits != bits {
		t.Errorf("NewBlockHeader: wrong bits - got %v, want %v",
			bh.Bits, bits)
	}
	if bh.Nonce != nonce {
		t.Errorf("NewBlockHeader: wrong nonce - got %v, want %v",
			bh.Nonce, nonce)
	}
}

// TestBlockHeaderWire tests the BlockHeader wire encode and decode for various
// protocol versions.
func TestBlockHeaderWire(t *testing.T) {
	nonce := uint32(123123) // 0x1e0f3
	pver := uint32(70001)

	// baseBlockHdr is used in the various tests as a baseline BlockHeader.
	bits := uint32(0x1d00ffff)
	baseBlockHdr := &BlockHeader{
		Version:       1,
		PrevBlock:     mainNetGenesisHash,
		MerkleRoot:    mainNetGenesisMerkleRoot,
		Timestamp:     time.Unix(0x495fab29, 0), // 2009-01-03 12:15:05 -0600 CST
		Bits:          bits,
		Nonce:         nonce,
		HashStateRoot: mainNetGenesisHashStateRoot,
		HashUTXORoot:  mainNetGenesisHashUTXORoot,
		BlockSig:      make([]byte, 0),
	}

	// baseBlockHdrEncoded is the wire encoded bytes of baseBlockHdr.
	baseBlockHdrEncoded := []byte{
		0x01, 0x00, 0x00, 0x00, // Version 1
		0x6c, 0x98, 0xed, 0x82, 0x55, 0x5e, 0xf9, 0xe3,
		0xd6, 0x34, 0x83, 0xd9, 0x21, 0xfa, 0x6c, 0x09,
		0xc3, 0xf8, 0xe6, 0x8c, 0xae, 0xf8, 0x80, 0x35,
		0x85, 0xf2, 0x3c, 0xf8, 0xae, 0x75, 0x00, 0x00, // PrevBlock
		0x6d, 0xb9, 0x05, 0x14, 0x23, 0x82, 0x32, 0x4d,
		0xb4, 0x17, 0x76, 0x18, 0x91, 0xf2, 0xd2, 0xf3,
		0x55, 0xea, 0x92, 0xf2, 0x7a, 0xb0, 0xfc, 0x35,
		0xe5, 0x9e, 0x90, 0xb5, 0x0e, 0x05, 0x34, 0xed, // MerkleRoot
		0x29, 0xab, 0x5f, 0x49, // Timestamp
		0xff, 0xff, 0x00, 0x1d, // Bits
		0xf3, 0xe0, 0x01, 0x00, // Nonce
		0xe9, 0x65, 0xff, 0xd0, 0x02, 0xcd, 0x6a, 0xd0,
		0xe2, 0xdc, 0x40, 0x2b, 0x80, 0x44, 0xde, 0x83,
		0x3e, 0x06, 0xb2, 0x31, 0x27, 0xea, 0x8c, 0x3d,
		0x80, 0xae, 0xc9, 0x14, 0x10, 0x77, 0x14, 0x95, // staste root
		0x56, 0xe8, 0x1f, 0x17, 0x1b, 0xcc, 0x55, 0xa6,
		0xff, 0x83, 0x45, 0xe6, 0x92, 0xc0, 0xf8, 0x6e,
		0x5b, 0x48, 0xe0, 0x1b, 0x99, 0x6c, 0xad, 0xc0,
		0x01, 0x62, 0x2f, 0xb5, 0xe3, 0x63, 0xb4, 0x21, // utxo root
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, // prevout stake
		0x00, 0x00, 0x00, 0x00, // prevout N
		0x00, // BlockSig
	}

	tests := []struct {
		in   *BlockHeader    // Data to encode
		out  *BlockHeader    // Expected decoded data
		buf  []byte          // Wire encoding
		pver uint32          // Protocol version for wire encoding
		enc  MessageEncoding // Message encoding variant to use
	}{
		// Latest protocol version.
		{
			baseBlockHdr,
			baseBlockHdr,
			baseBlockHdrEncoded,
			ProtocolVersion,
			BaseEncoding,
		},

		// Protocol version BIP0035Version.
		{
			baseBlockHdr,
			baseBlockHdr,
			baseBlockHdrEncoded,
			BIP0035Version,
			BaseEncoding,
		},

		// Protocol version BIP0031Version.
		{
			baseBlockHdr,
			baseBlockHdr,
			baseBlockHdrEncoded,
			BIP0031Version,
			BaseEncoding,
		},

		// Protocol version NetAddressTimeVersion.
		{
			baseBlockHdr,
			baseBlockHdr,
			baseBlockHdrEncoded,
			NetAddressTimeVersion,
			BaseEncoding,
		},

		// Protocol version MultipleAddressVersion.
		{
			baseBlockHdr,
			baseBlockHdr,
			baseBlockHdrEncoded,
			MultipleAddressVersion,
			BaseEncoding,
		},
	}

	t.Logf("Running %d tests", len(tests))
	for i, test := range tests {
		// Encode to wire format.
		var buf bytes.Buffer
		err := writeBlockHeader(&buf, test.pver, test.in)
		if err != nil {
			t.Errorf("writeBlockHeader #%d error %v", i, err)
			continue
		}
		if !bytes.Equal(buf.Bytes(), test.buf) {
			t.Errorf("writeBlockHeader #%d\n got: %s want: %s", i,
				spew.Sdump(buf.Bytes()), spew.Sdump(test.buf))
			continue
		}

		buf.Reset()
		err = test.in.BtcEncode(&buf, pver, 0)
		if err != nil {
			t.Errorf("BtcEncode #%d error %v", i, err)
			continue
		}
		if !bytes.Equal(buf.Bytes(), test.buf) {
			t.Errorf("BtcEncode #%d\n got: %s want: %s", i,
				spew.Sdump(buf.Bytes()), spew.Sdump(test.buf))
			continue
		}

		// Decode the block header from wire format.
		var bh BlockHeader
		rbuf := bytes.NewReader(test.buf)
		err = readBlockHeader(rbuf, test.pver, &bh)
		if err != nil {
			t.Errorf("readBlockHeader #%d error %v", i, err)
			continue
		}
		if !reflect.DeepEqual(&bh, test.out) {
			t.Errorf("readBlockHeader #%d\n got: %s want: %s", i,
				spew.Sdump(&bh), spew.Sdump(test.out))
			continue
		}

		rbuf = bytes.NewReader(test.buf)
		err = bh.BtcDecode(rbuf, pver, test.enc)
		if err != nil {
			t.Errorf("BtcDecode #%d error %v", i, err)
			continue
		}
		if !reflect.DeepEqual(&bh, test.out) {
			t.Errorf("BtcDecode #%d\n got: %s want: %s", i,
				spew.Sdump(&bh), spew.Sdump(test.out))
			continue
		}
	}
}

// TestBlockHeaderSerialize tests BlockHeader serialize and deserialize.
func TestBlockHeaderSerialize(t *testing.T) {
	nonce := uint32(123123) // 0x1e0f3

	// baseBlockHdr is used in the various tests as a baseline BlockHeader.
	bits := uint32(0x1d00ffff)
	baseBlockHdr := &BlockHeader{
		Version:       1,
		PrevBlock:     mainNetGenesisHash,
		MerkleRoot:    mainNetGenesisMerkleRoot,
		Timestamp:     time.Unix(0x495fab29, 0), // 2009-01-03 12:15:05 -0600 CST
		Bits:          bits,
		Nonce:         nonce,
		HashStateRoot: mainNetGenesisHashStateRoot,
		HashUTXORoot:  mainNetGenesisHashUTXORoot,
		BlockSig:      make([]byte, 0),
	}

	// baseBlockHdrEncoded is the wire encoded bytes of baseBlockHdr.
	// baseBlockHdrEncoded is the wire encoded bytes of baseBlockHdr.
	baseBlockHdrEncoded := []byte{
		0x01, 0x00, 0x00, 0x00, // Version 1
		0x6c, 0x98, 0xed, 0x82, 0x55, 0x5e, 0xf9, 0xe3,
		0xd6, 0x34, 0x83, 0xd9, 0x21, 0xfa, 0x6c, 0x09,
		0xc3, 0xf8, 0xe6, 0x8c, 0xae, 0xf8, 0x80, 0x35,
		0x85, 0xf2, 0x3c, 0xf8, 0xae, 0x75, 0x00, 0x00, // PrevBlock
		0x6d, 0xb9, 0x05, 0x14, 0x23, 0x82, 0x32, 0x4d,
		0xb4, 0x17, 0x76, 0x18, 0x91, 0xf2, 0xd2, 0xf3,
		0x55, 0xea, 0x92, 0xf2, 0x7a, 0xb0, 0xfc, 0x35,
		0xe5, 0x9e, 0x90, 0xb5, 0x0e, 0x05, 0x34, 0xed, // MerkleRoot
		0x29, 0xab, 0x5f, 0x49, // Timestamp
		0xff, 0xff, 0x00, 0x1d, // Bits
		0xf3, 0xe0, 0x01, 0x00, // Nonce
		0xe9, 0x65, 0xff, 0xd0, 0x02, 0xcd, 0x6a, 0xd0,
		0xe2, 0xdc, 0x40, 0x2b, 0x80, 0x44, 0xde, 0x83,
		0x3e, 0x06, 0xb2, 0x31, 0x27, 0xea, 0x8c, 0x3d,
		0x80, 0xae, 0xc9, 0x14, 0x10, 0x77, 0x14, 0x95, // staste root
		0x56, 0xe8, 0x1f, 0x17, 0x1b, 0xcc, 0x55, 0xa6,
		0xff, 0x83, 0x45, 0xe6, 0x92, 0xc0, 0xf8, 0x6e,
		0x5b, 0x48, 0xe0, 0x1b, 0x99, 0x6c, 0xad, 0xc0,
		0x01, 0x62, 0x2f, 0xb5, 0xe3, 0x63, 0xb4, 0x21, // utxo root
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, // prevout stake
		0x00, 0x00, 0x00, 0x00, // prevout N
		0x00, // BlockSig
	}

	tests := []struct {
		in  *BlockHeader // Data to encode
		out *BlockHeader // Expected decoded data
		buf []byte       // Serialized data
	}{
		{
			baseBlockHdr,
			baseBlockHdr,
			baseBlockHdrEncoded,
		},
	}

	t.Logf("Running %d tests", len(tests))
	for i, test := range tests {
		// Serialize the block header.
		var buf bytes.Buffer
		err := test.in.Serialize(&buf)
		if err != nil {
			t.Errorf("Serialize #%d error %v", i, err)
			continue
		}
		if !bytes.Equal(buf.Bytes(), test.buf) {
			t.Errorf("Serialize #%d\n got: %s want: %s", i,
				spew.Sdump(buf.Bytes()), spew.Sdump(test.buf))
			continue
		}

		// Deserialize the block header.
		var bh BlockHeader
		rbuf := bytes.NewReader(test.buf)
		err = bh.Deserialize(rbuf)
		if err != nil {
			t.Errorf("Deserialize #%d error %v", i, err)
			continue
		}
		if !reflect.DeepEqual(&bh, test.out) {
			t.Errorf("Deserialize #%d\n got: %s want: %s", i,
				spew.Sdump(&bh), spew.Sdump(test.out))
			continue
		}
	}
}
