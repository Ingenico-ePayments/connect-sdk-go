package logging

import (
	"encoding/json"
	"errors"
	"strings"

	"github.com/Jeffail/gabs"
)

var errInvalidDataType = errors.New("invalid data type - should never happen")

type propertyObfuscator struct {
	obfuscators valueObfuscatorMap
}

func newPropertyObfuscator(obfuscators valueObfuscatorMap) propertyObfuscator {
	return propertyObfuscator{obfuscators}
}

func (po *propertyObfuscator) navigateJSON(cont *gabs.Container) error {
	contVec := append([]*gabs.Container(nil), cont)

	for len(contVec) > 0 {
		var currentCont *gabs.Container
		currentCont, contVec = contVec[len(contVec)-1], contVec[:len(contVec)-1]

		children, sErr := currentCont.ChildrenMap()

		if sErr != nil {
			if sErr == gabs.ErrNotObj {
				continue
			}

			return sErr
		}

		for key, value := range children {
			obfuscator, ok := po.obfuscators[key]
			if ok == true {
				strVal, cErr := value.Data().(string)

				if cErr != true {
					return errInvalidDataType
				}

				newVal, oErr := obfuscator.obfuscateValue(strVal)

				if oErr != nil {
					return oErr
				}

				currentCont.Set(newVal, key)
			}

			contVec = append(contVec, value)
		}
	}

	return nil
}

func (po *propertyObfuscator) obfuscate(body string) (string, error) {
	if strings.TrimSpace(body) == "" {
		return body, nil
	}

	parsedJSON, pErr := gabs.ParseJSON([]byte(body))

	if _, ok := pErr.(*json.SyntaxError); ok {
		return body, nil
	}

	if pErr != nil {
		return body, pErr
	}

	nError := po.navigateJSON(parsedJSON)

	if nError != nil {
		return body, nError
	}

	return parsedJSON.StringIndent("", "    "), nil
}
