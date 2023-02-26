package converter

import (
	"bytes"
	"encoding/json"
	"net/url"
	"strconv"
)

func StructToJsonTagMap(data interface{}) (map[string]interface{}, error) {
	jsonStr, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}
	out := new(bytes.Buffer)
	// JSONの整形
	err = json.Indent(out, jsonStr, "", "    ")
	if err != nil {
		return nil, err
	}
	var mapData map[string]interface{}
	if err := json.Unmarshal([]byte(out.String()), &mapData); err != nil {
		return nil, err
	}
	return mapData, err
}

func StructToJsonTagQueryStr(data interface{}) (string, error) {
	jsonStr, err := json.Marshal(data)
	if err != nil {
		return "", err
	}
	out := new(bytes.Buffer)
	// JSONの整形
	err = json.Indent(out, jsonStr, "", "    ")
	if err != nil {
		return "", err
	}
	var mapData map[string]interface{}
	if err := json.Unmarshal([]byte(out.String()), &mapData); err != nil {
		return "", err
	}

	values := url.Values{}
	for k, v := range mapData {
		switch val := v.(type) {
		case string:
			values.Set(k, val)
		case int:
			values.Set(k, strconv.Itoa(val))
		}
	}
	return values.Encode(), err
}
