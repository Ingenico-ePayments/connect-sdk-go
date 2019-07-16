package logging

import (
	"encoding/json"
	"errors"
	"strconv"
	"strings"
)

var errInvalidDataType = errors.New("invalid data type - should never happen")

type propertyObfuscator struct {
	obfuscators valueObfuscatorMap
}

func newPropertyObfuscator(obfuscators valueObfuscatorMap) propertyObfuscator {
	return propertyObfuscator{obfuscators}
}

func (po *propertyObfuscator) obfuscateValue(value interface{}, obfuscator *valueObfuscator) (string, error) {
	if strVal, ok := value.(string); ok {
		return obfuscator.obfuscateValue(strVal)
	}
	if intVal, ok := value.(int32); ok {
		return obfuscator.obfuscateValue(strconv.FormatInt(int64(intVal), 10))
	}
	if intVal, ok := value.(int64); ok {
		return obfuscator.obfuscateValue(strconv.FormatInt(intVal, 10))
	}
	if floatVal, ok := value.(float32); ok {
		return obfuscator.obfuscateValue(strconv.FormatFloat(float64(floatVal), 'f', -1, 32))
	}
	if floatVal, ok := value.(float64); ok {
		return obfuscator.obfuscateValue(strconv.FormatFloat(floatVal, 'f', -1, 64))
	}
	if boolVal, ok := value.(bool); ok {
		return obfuscator.obfuscateValue(strconv.FormatBool(boolVal))
	}
	return "", errInvalidDataType
}

func (po *propertyObfuscator) navigateJSON(content interface{}) error {
	if content == nil {
		return nil
	}

	if contentMap, ok := content.(map[string]interface{}); ok {
		for name, obj := range contentMap {
			if obfuscator, ok := po.obfuscators[name]; ok {
				obfuscatedValue, oErr := po.obfuscateValue(obj, &obfuscator)
				if oErr != nil {
					return oErr
				}

				contentMap[name] = obfuscatedValue

			} else {
				nErr := po.navigateJSON(obj)
				if nErr != nil {
					return nErr
				}
			}
		}
	}

	if contentSlice, ok := content.([]interface{}); ok {
		for _, obj := range contentSlice {
			nErr := po.navigateJSON(obj)
			if nErr != nil {
				return nErr
			}
		}
	}

	return nil
}

func (po *propertyObfuscator) obfuscate(body string) (string, error) {
	if strings.TrimSpace(body) == "" {
		return body, nil
	}

	var parsedJSON interface{}
	pErr := json.Unmarshal([]byte(body), &parsedJSON)
	if _, ok := pErr.(*json.SyntaxError); ok {
		return body, nil
	}
	if pErr != nil {
		return body, pErr
	}

	nErr := po.navigateJSON(parsedJSON)
	if nErr != nil {
		return body, nErr
	}

	obfuscatedBody, mErr := json.MarshalIndent(parsedJSON, "", "    ")
	if mErr != nil {
		return body, mErr
	}

	return string(obfuscatedBody), nil
}
