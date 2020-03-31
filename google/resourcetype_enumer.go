// Code generated by "enumer -type ResourceType -addprefix google_ -transform snake -linecomment"; DO NOT EDIT.

package google

import (
	"fmt"
)

const _ResourceTypeName = "google_compute_instancegoogle_compute_firewallgoogle_compute_networkgoogle_compute_health_checkgoogle_compute_instance_groupgoogle_compute_backend_bucketgoogle_compute_backend_servicegoogle_compute_ssl_certificategoogle_compute_target_http_proxygoogle_compute_target_https_proxygoogle_compute_url_mapgoogle_compute_global_forwarding_rulegoogle_compute_forwarding_rulegoogle_compute_diskgoogle_dns_managed_zonegoogle_dns_record_setgoogle_storage_bucketgoogle_sql_database_instance"

var _ResourceTypeIndex = [...]uint16{0, 23, 46, 68, 95, 124, 153, 183, 213, 245, 278, 300, 337, 367, 386, 409, 430, 451, 479}

const _ResourceTypeLowerName = "google_compute_instancegoogle_compute_firewallgoogle_compute_networkgoogle_compute_health_checkgoogle_compute_instance_groupgoogle_compute_backend_bucketgoogle_compute_backend_servicegoogle_compute_ssl_certificategoogle_compute_target_http_proxygoogle_compute_target_https_proxygoogle_compute_url_mapgoogle_compute_global_forwarding_rulegoogle_compute_forwarding_rulegoogle_compute_diskgoogle_dns_managed_zonegoogle_dns_record_setgoogle_storage_bucketgoogle_sql_database_instance"

func (i ResourceType) String() string {
	if i < 0 || i >= ResourceType(len(_ResourceTypeIndex)-1) {
		return fmt.Sprintf("ResourceType(%d)", i)
	}
	return _ResourceTypeName[_ResourceTypeIndex[i]:_ResourceTypeIndex[i+1]]
}

var _ResourceTypeValues = []ResourceType{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17}

var _ResourceTypeNameToValueMap = map[string]ResourceType{
	_ResourceTypeName[0:23]:         0,
	_ResourceTypeLowerName[0:23]:    0,
	_ResourceTypeName[23:46]:        1,
	_ResourceTypeLowerName[23:46]:   1,
	_ResourceTypeName[46:68]:        2,
	_ResourceTypeLowerName[46:68]:   2,
	_ResourceTypeName[68:95]:        3,
	_ResourceTypeLowerName[68:95]:   3,
	_ResourceTypeName[95:124]:       4,
	_ResourceTypeLowerName[95:124]:  4,
	_ResourceTypeName[124:153]:      5,
	_ResourceTypeLowerName[124:153]: 5,
	_ResourceTypeName[153:183]:      6,
	_ResourceTypeLowerName[153:183]: 6,
	_ResourceTypeName[183:213]:      7,
	_ResourceTypeLowerName[183:213]: 7,
	_ResourceTypeName[213:245]:      8,
	_ResourceTypeLowerName[213:245]: 8,
	_ResourceTypeName[245:278]:      9,
	_ResourceTypeLowerName[245:278]: 9,
	_ResourceTypeName[278:300]:      10,
	_ResourceTypeLowerName[278:300]: 10,
	_ResourceTypeName[300:337]:      11,
	_ResourceTypeLowerName[300:337]: 11,
	_ResourceTypeName[337:367]:      12,
	_ResourceTypeLowerName[337:367]: 12,
	_ResourceTypeName[367:386]:      13,
	_ResourceTypeLowerName[367:386]: 13,
	_ResourceTypeName[386:409]:      14,
	_ResourceTypeLowerName[386:409]: 14,
	_ResourceTypeName[409:430]:      15,
	_ResourceTypeLowerName[409:430]: 15,
	_ResourceTypeName[430:451]:      16,
	_ResourceTypeLowerName[430:451]: 16,
	_ResourceTypeName[451:479]:      17,
	_ResourceTypeLowerName[451:479]: 17,
}

var _ResourceTypeNames = []string{
	_ResourceTypeName[0:23],
	_ResourceTypeName[23:46],
	_ResourceTypeName[46:68],
	_ResourceTypeName[68:95],
	_ResourceTypeName[95:124],
	_ResourceTypeName[124:153],
	_ResourceTypeName[153:183],
	_ResourceTypeName[183:213],
	_ResourceTypeName[213:245],
	_ResourceTypeName[245:278],
	_ResourceTypeName[278:300],
	_ResourceTypeName[300:337],
	_ResourceTypeName[337:367],
	_ResourceTypeName[367:386],
	_ResourceTypeName[386:409],
	_ResourceTypeName[409:430],
	_ResourceTypeName[430:451],
	_ResourceTypeName[451:479],
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
