package binary

import "testing"

func TestEndian(t *testing.T) {
	t.Logf("%v", IsLittleEndian())
}
