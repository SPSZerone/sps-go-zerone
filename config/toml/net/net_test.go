package net

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestAddr(t *testing.T) {
	content := `
type = "tcp4"
host = "localhost"
port = 8080
`
	t.Log("========== content ==========")
	t.Logf("%v", content)

	addr, err := NewAddrFromContent(content)
	require.Nilf(t, err, "content:%v err:%v", content, err)
	t.Log("========== toml.Decode ==========")
	t.Logf("%+v", addr)
	t.Logf("%#v", addr)
}
