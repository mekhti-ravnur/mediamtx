package handshake

import (
	"bufio"
	"bytes"
	"testing"

	"github.com/stretchr/testify/require"
)

var c0s0enc = []byte{0x03}

var c0s0dec = C0S0{}

func TestC0S0Read(t *testing.T) {
	var c0s0 C0S0
	err := c0s0.Read(bufio.NewReader(bytes.NewReader(c0s0enc)))
	require.NoError(t, err)
	require.Equal(t, c0s0dec, c0s0)
}

func TestC0S0Write(t *testing.T) {
	var buf bytes.Buffer
	err := c0s0dec.Write(&buf)
	require.NoError(t, err)
	require.Equal(t, c0s0enc, buf.Bytes())
}
