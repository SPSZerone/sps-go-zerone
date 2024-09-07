package error

import "testing"

func TestError(t *testing.T) {
	var err Error

	err.Update(
		OptFormatPrefix("A:%s B:%s C:%s"),
		OptArgsPrefix("aaa", "bbb", "ccc"))
	e := err.Errorf(" %s", "test")
	t.Log(e)

	err.Update(
		OptArgsPrefix("xxx", "yyy", "zzz"),
		OptFormatSuffix(" X:%d Y:%d Z:%d"),
		OptArgsSuffix(1, 2, 3))
	e = err.Errorf(" %s", "test")
	t.Log(e)

	err.Reset()
	e = err.Errorf("%s", "test")
	t.Log(e)
}
