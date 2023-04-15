package maputil

import "sync"

type IDeepKeyOperator interface {
	Set(target map[string]interface{}, deepKey string, value interface{}) error
	SetIfAbsent(target map[string]interface{}, deepKey string, value interface{}) error
	Gen(deepKey string, value interface{}) (map[string]interface{}, error)
	GenSilently(deepKey string, value interface{}) map[string]interface{}
}

type deepKeyOperator struct {
	mutex sync.RWMutex
}

var (
	DeepKeyOperator = deepKeyOperator{}
)

func (d *deepKeyOperator) Set(target map[string]interface{}, deepKey string, value interface{}) error {
	split, err := DeepKeyParser.Parse(deepKey)
	if err != nil {
		return err
	}
	if len(split) == 1 {
		target[deepKey] = value
		return nil
	}
	cur := target
	for i := 0; i < len(split)-1; i++ {
		curKey := split[i]
		var m map[string]interface{}
		t, ok := cur[curKey]

		if ok {
			tm, assertOK := t.(map[string]interface{})
			if assertOK {
				m = tm
			} else {
				m = make(map[string]interface{})
			}
		} else {
			m = make(map[string]interface{})
		}
		cur[curKey] = m
		cur = m
	}

	cur[split[len(split)-1]] = value
	return nil
}

func (d *deepKeyOperator) SetIfAbsent(target map[string]interface{}, deepKey string, value interface{}) error {
	split, err := DeepKeyParser.Parse(deepKey)
	if err != nil {
		return err
	}
	if len(split) == 1 {
		_, ok := target[deepKey]
		if ok {
			return nil
		}
		target[deepKey] = value
		return nil
	}
	cur := target
	for i := 0; i < len(split)-1; i++ {
		curKey := split[i]
		var m map[string]interface{}
		t, ok := cur[curKey]
		if ok {
			tm, assertOK := t.(map[string]interface{})
			if assertOK {
				m = tm
			} else {
				m = make(map[string]interface{})
			}
		} else {
			m = make(map[string]interface{})
		}
		cur[curKey] = m
		cur = m
	}

	_, ok := cur[split[len(split)-1]]
	if ok {
		return nil
	}
	cur[split[len(split)-1]] = value
	return nil
}

func (d *deepKeyOperator) Gen(deepKey string, value interface{}) (map[string]interface{}, error) {
	var ret = make(map[string]interface{})
	return ret, d.Set(ret, deepKey, value)
}

func (d *deepKeyOperator) GenSilently(deepKey string, value interface{}) map[string]interface{} {
	gen, _ := d.Gen(deepKey, value)
	return gen
}
