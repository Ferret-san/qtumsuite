// Copyright (c) 2013-2016 The btcsuite developers
// Use of this source code is governed by an ISC
// license that can be found in the LICENSE file.

package bloom_test

import (
	"bytes"
	"encoding/hex"
	"testing"

	"github.com/qtumproject/qtumsuite"
	"github.com/qtumproject/qtumsuite/bloom"
	"github.com/qtumproject/qtumsuite/chaincfg/chainhash"
	"github.com/qtumproject/qtumsuite/wire"
)

func TestMerkleBlock3(t *testing.T) {
	blockStr := "01000000" +
		"79cda856b143d9db2c1caff01d1aecc8630d30625d10e8b4b8b0000000000000" +
		"b50cc069d6a3e33e3ff84a5c41d9d3febe7c770fdcc96b2c3ff60abe184f1963" +
		"67291b4d" +
		"4c86041b" +
		"8fa45d63" +
		"e965ffd002cd6ad0e2dc402b8044de833e06b23127ea8c3d80aec91410771495" +
		"56e81f171bcc55a6ff8345e692c0f86e5b48e01b996cadc001622fb5e363b421" +
		"0000000000000000000000000000000000000000000000000000000000000000" +
		"ffffffff" +
		"00" +
		"010100000001000000000000000000000000000000000000000000000000000000000000" +
		"0000ffffffff08044c86041b020a02ffffffff0100f2052a01000000434" +
		"104ecd3229b0571c3be876feaac0442a9f13c5a572742927af1dc623353" +
		"ecf8c202225f64868137a18cdd85cbbb4c74fbccfd4f49639cf1bdc94a5" +
		"672bb15ad5d4cac00000000"
	blockBytes, err := hex.DecodeString(blockStr)
	if err != nil {
		t.Errorf("TestMerkleBlock3 DecodeString failed: %v", err)
		return
	}
	blk, err := qtumsuite.NewBlockFromBytes(blockBytes)
	if err != nil {
		t.Errorf("TestMerkleBlock3 NewBlockFromBytes failed: %v", err)
		return
	}

	f := bloom.NewFilter(10, 0, 0.000001, wire.BloomUpdateAll)

	inputStr := "63194f18be0af63f2c6bc9dc0f777cbefed3d9415c4af83f3ee3a3d669c00cb5"
	hash, err := chainhash.NewHashFromStr(inputStr)
	if err != nil {
		t.Errorf("TestMerkleBlock3 NewHashFromStr failed: %v", err)
		return
	}

	f.AddHash(hash)

	mBlock, _ := bloom.NewMerkleBlock(blk, f)

	wantStr := "0100000079cda856b143d9db2c1caff01d1aecc8630d30625d10e8b4b8b0" +
		"000000000000b50cc069d6a3e33e3ff84a5c41d9d3febe7c770fdcc96b2c3ff60ab" +
		"e184f196367291b4d4c86041b8fa45d63e965ffd002cd6ad0e2dc402b8044de833e" +
		"06b23127ea8c3d80aec9141077149556e81f171bcc55a6ff8345e692c0f86e5b48e" +
		"01b996cadc001622fb5e363b4210000000000000000000000000000000000000000" +
		"000000000000000000000000ffffffff000100000001b50cc069d6a3e33e3ff84a5" +
		"c41d9d3febe7c770fdcc96b2c3ff60abe184f19630101"
	want, err := hex.DecodeString(wantStr)
	if err != nil {
		t.Errorf("TestMerkleBlock3 DecodeString failed: %v", err)
		return
	}

	got := bytes.NewBuffer(nil)
	err = mBlock.BtcEncode(got, wire.ProtocolVersion, wire.LatestEncoding)
	if err != nil {
		t.Errorf("TestMerkleBlock3 BtcEncode failed: %v", err)
		return
	}

	if !bytes.Equal(want, got.Bytes()) {
		t.Errorf("TestMerkleBlock3 failed merkle block comparison: "+
			"got %v want %v", got.Bytes(), want)
		return
	}
}
