package nftfx

import (
	"testing"

	"github.com/ava-labs/gecko/vms/secp256k1fx"
)

func TestTransferOperationVerifyNil(t *testing.T) {
	op := (*TransferOperation)(nil)
	if err := op.Verify(); err == nil {
		t.Fatalf("nil operation should have failed verification")
	}
}

func TestTransferOperationInvalid(t *testing.T) {
	op := TransferOperation{Input: secp256k1fx.Input{
		SigIndices: []uint32{1, 0},
	}}
	if err := op.Verify(); err == nil {
		t.Fatalf("operation should have failed verification")
	}
}

func TestTransferOperationOuts(t *testing.T) {
	op := TransferOperation{
		Output: TransferOutput{},
	}
	if outs := op.Outs(); len(outs) != 1 {
		t.Fatalf("Wrong number of outputs returned")
	}
}