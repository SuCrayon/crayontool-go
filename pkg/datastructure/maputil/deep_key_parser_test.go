package maputil

import (
	"crayontool-go/pkg/constant"
	"strings"
	"testing"
)

var (
	deepKeyTestList = []string{
		"outerKey.innerKey",
		"outerKey\\.innerKey",
		"outerKey\\\\.innerKey",
		".outerKey.innerKey",
		"outerKey..innerKey",
		"outerKey.innerKey.",
		"\\.outerKey.innerKey",
		"\\\\.outerKey.innerKey",
		"outerKey.innerKey\\.",
		"outerKey.innerKey\\\\.",
		"outer\\Key.innerKey",
		"outer\\\\Key.innerKey",
	}
)

func Test_deepKeyParser_Parse(t *testing.T) {
	parser := deepKeyParser{}
	for _, deepKey := range deepKeyTestList {
		ret, err := parser.Parse(deepKey)
		if err != nil {
			t.Error(err)
		}
		t.Logf("deepKey: %s, len: %d, ret: %v\n", deepKey, len(ret), ret)
	}
}

func Test_deepKeyParserV2_Parse(t *testing.T) {
	parser := deepKeyParserV2{}
	for _, deepKey := range deepKeyTestList {
		ret, err := parser.Parse(deepKey)
		if err != nil {
			t.Error(err)
		}
		t.Logf("deepKey: %s, len: %d, ret: %v\n", deepKey, len(ret), ret)
	}
}

func Test_deepKeyParser_ParseCompare(t *testing.T) {
	parserV1 := deepKeyParser{}
	parserV2 := deepKeyParserV2{}
	for _, deepKey := range deepKeyTestList {

		retV1, err := parserV1.Parse(deepKey)
		if err != nil {
			t.Error(err)
		}
		t.Logf("deepKey: %s, len: %d, ret: %v\n", deepKey, len(retV1), retV1)

		retV2, err := parserV2.Parse(deepKey)
		if err != nil {
			t.Error(err)
		}
		t.Logf("deepKey: %s, len: %d, ret: %v\n", deepKey, len(retV2), retV2)

		if strings.Join(retV1, constant.EmptyStr) != strings.Join(retV2, constant.EmptyStr) {
			t.Errorf("parse result is not equal, retV1: %v, retV2: %v\n", retV1, retV2)
		}
	}
}

type customDeepKeyParser struct {
	deepKeyParserV2
}

func Test_deepKeyParser_CustomParser(t *testing.T) {
	parser := customDeepKeyParser{}
	parser.SetKeySepSymbol(constant.DollarSymbol)
	for _, rawDeepKey := range deepKeyTestList {
		deepKey := strings.ReplaceAll(rawDeepKey, constant.Dot, constant.Dollar)
		ret, err := parser.Parse(deepKey)
		if err != nil {
			t.Error(err)
		}
		t.Logf("deepKey: %s, len: %d, ret: %v\n", deepKey, len(ret), ret)
	}
}
