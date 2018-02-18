package block

import "testing"

func TestValidate(t *testing.T) {
	AdjustDifficulty(8)
	bike := Bike{"SN123545"}
	block, _ := NewBlock(&bike, []byte("SN123544"))
	pow := NewProofOfWork(block)

	if !pow.Validate() {
		t.Error("Validation of block failed")
	}
}