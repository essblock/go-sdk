package compliance

import (
	"github.com/asaskevich/govalidator"
	"github.com/essblock/go-sdk/address"
)

func init() {
	govalidator.SetFieldsRequiredByDefault(true)
	govalidator.CustomTypeTagMap.Set("ess_address", govalidator.CustomTypeValidator(isEssAddress))
}

func isEssAddress(i interface{}, context interface{}) bool {
	addr, ok := i.(string)

	if !ok {
		return false
	}

	_, _, err := address.Split(addr)

	if err == nil {
		return true
	}

	return false
}
