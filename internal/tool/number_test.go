package tool

import "testing"

func TestNumber(t *testing.T) {
	t.Log(ConvertGeo("121.4356079"))
	t.Log(ConvertGeo("031.2945919"))
	t.Log(ConvertGeo("-031.2945919"))
	t.Log(ConvertGeo(""))
	t.Log(ConvertGeo(nil))
}
