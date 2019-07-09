// Code generated by "enumer -type ResourceType -addprefix google_ -transform snake -linecomment"; DO NOT EDIT.

package google

import (
	"fmt"
)

const _ResourceTypeName = "google_not_implemented"

var _ResourceTypeIndex = [...]uint8{0, 22}

const _ResourceTypeLowerName = "google_not_implemented"

func (i ResourceType) String() string {
	if i < 0 || i >= ResourceType(len(_ResourceTypeIndex)-1) {
		return fmt.Sprintf("ResourceType(%d)", i)
	}
	return _ResourceTypeName[_ResourceTypeIndex[i]:_ResourceTypeIndex[i+1]]
}

var _ResourceTypeValues = []ResourceType{0}

var _ResourceTypeNameToValueMap = map[string]ResourceType{
	_ResourceTypeName[0:22]:      0,
	_ResourceTypeLowerName[0:22]: 0,
}

var _ResourceTypeNames = []string{
	_ResourceTypeName[0:22],
}

// ResourceTypeString retrieves an enum value from the enum constants string name.
// Throws an error if the param is not part of the enum.
func ResourceTypeString(s string) (ResourceType, error) {
	if val, ok := _ResourceTypeNameToValueMap[s]; ok {
		return val, nil
	}
	return 0, fmt.Errorf("%s does not belong to ResourceType values", s)
}

// ResourceTypeValues returns all values of the enum
func ResourceTypeValues() []ResourceType {
	return _ResourceTypeValues
}

// ResourceTypeStrings returns a slice of all String values of the enum
func ResourceTypeStrings() []string {
	strs := make([]string, len(_ResourceTypeNames))
	copy(strs, _ResourceTypeNames)
	return strs
}

// IsAResourceType returns "true" if the value is listed in the enum definition. "false" otherwise
func (i ResourceType) IsAResourceType() bool {
	for _, v := range _ResourceTypeValues {
		if i == v {
			return true
		}
	}
	return false
}