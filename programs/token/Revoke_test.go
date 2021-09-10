package token

import (
	"bytes"
	"strconv"
	"testing"

	ag_gofuzz "github.com/gagliardetto/gofuzz"
	ag_require "github.com/stretchr/testify/require"
)

func TestEncodeDecode_Revoke(t *testing.T) {
	fu := ag_gofuzz.New().NilChance(0)
	for i := 0; i < 1; i++ {
		t.Run("Revoke"+strconv.Itoa(i), func(t *testing.T) {
			{
				params := new(Revoke)
				fu.Fuzz(params)
				params.Accounts = nil
				params.Signers = nil
				buf := new(bytes.Buffer)
				err := encodeT(*params, buf)
				ag_require.NoError(t, err)
				//
				got := new(Revoke)
				err = decodeT(got, buf.Bytes())
				params.Accounts = nil
				params.Signers = nil
				ag_require.NoError(t, err)
				ag_require.Equal(t, params, got)
			}
		})
	}
}