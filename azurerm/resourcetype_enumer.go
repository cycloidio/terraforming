// Code generated by "enumer -type ResourceType -addprefix azurerm_ -transform snake -linecomment"; DO NOT EDIT.

package azurerm

import (
	"fmt"
	"strings"
)

const _ResourceTypeName = "azurerm_resource_groupazurerm_subnetazurerm_virtual_desktop_host_poolazurerm_virtual_desktop_application_groupazurerm_logic_app_trigger_customazurerm_logic_app_action_customazurerm_logic_app_workflowazurerm_network_interfaceazurerm_network_security_groupazurerm_virtual_machineazurerm_virtual_machine_extensionazurerm_virtual_machine_scale_setazurerm_virtual_network"

var _ResourceTypeIndex = [...]uint16{0, 22, 36, 69, 110, 142, 173, 199, 224, 254, 277, 310, 343, 366}

const _ResourceTypeLowerName = "azurerm_resource_groupazurerm_subnetazurerm_virtual_desktop_host_poolazurerm_virtual_desktop_application_groupazurerm_logic_app_trigger_customazurerm_logic_app_action_customazurerm_logic_app_workflowazurerm_network_interfaceazurerm_network_security_groupazurerm_virtual_machineazurerm_virtual_machine_extensionazurerm_virtual_machine_scale_setazurerm_virtual_network"

func (i ResourceType) String() string {
	if i < 0 || i >= ResourceType(len(_ResourceTypeIndex)-1) {
		return fmt.Sprintf("ResourceType(%d)", i)
	}
	return _ResourceTypeName[_ResourceTypeIndex[i]:_ResourceTypeIndex[i+1]]
}

// An "invalid array index" compiler error signifies that the constant values have changed.
// Re-run the stringer command to generate them again.
func _ResourceTypeNoOp() {
	var x [1]struct{}
	_ = x[ResourceGroup-(0)]
	_ = x[Subnet-(1)]
	_ = x[VirtualDesktopHostPool-(2)]
	_ = x[VirtualDesktopApplicationGroup-(3)]
	_ = x[LogicAppTriggerCustom-(4)]
	_ = x[LogicAppActionCustom-(5)]
	_ = x[LogicAppWorkflow-(6)]
	_ = x[NetworkInterface-(7)]
	_ = x[NetworkSecurityGroup-(8)]
	_ = x[VirtualMachine-(9)]
	_ = x[VirtualMachineExtension-(10)]
	_ = x[VirtualMachineScaleSet-(11)]
	_ = x[VirtualNetwork-(12)]
}

var _ResourceTypeValues = []ResourceType{ResourceGroup, Subnet, VirtualDesktopHostPool, VirtualDesktopApplicationGroup, LogicAppTriggerCustom, LogicAppActionCustom, LogicAppWorkflow, NetworkInterface, NetworkSecurityGroup, VirtualMachine, VirtualMachineExtension, VirtualMachineScaleSet, VirtualNetwork}

var _ResourceTypeNameToValueMap = map[string]ResourceType{
	_ResourceTypeName[0:22]:         ResourceGroup,
	_ResourceTypeLowerName[0:22]:    ResourceGroup,
	_ResourceTypeName[22:36]:        Subnet,
	_ResourceTypeLowerName[22:36]:   Subnet,
	_ResourceTypeName[36:69]:        VirtualDesktopHostPool,
	_ResourceTypeLowerName[36:69]:   VirtualDesktopHostPool,
	_ResourceTypeName[69:110]:       VirtualDesktopApplicationGroup,
	_ResourceTypeLowerName[69:110]:  VirtualDesktopApplicationGroup,
	_ResourceTypeName[110:142]:      LogicAppTriggerCustom,
	_ResourceTypeLowerName[110:142]: LogicAppTriggerCustom,
	_ResourceTypeName[142:173]:      LogicAppActionCustom,
	_ResourceTypeLowerName[142:173]: LogicAppActionCustom,
	_ResourceTypeName[173:199]:      LogicAppWorkflow,
	_ResourceTypeLowerName[173:199]: LogicAppWorkflow,
	_ResourceTypeName[199:224]:      NetworkInterface,
	_ResourceTypeLowerName[199:224]: NetworkInterface,
	_ResourceTypeName[224:254]:      NetworkSecurityGroup,
	_ResourceTypeLowerName[224:254]: NetworkSecurityGroup,
	_ResourceTypeName[254:277]:      VirtualMachine,
	_ResourceTypeLowerName[254:277]: VirtualMachine,
	_ResourceTypeName[277:310]:      VirtualMachineExtension,
	_ResourceTypeLowerName[277:310]: VirtualMachineExtension,
	_ResourceTypeName[310:343]:      VirtualMachineScaleSet,
	_ResourceTypeLowerName[310:343]: VirtualMachineScaleSet,
	_ResourceTypeName[343:366]:      VirtualNetwork,
	_ResourceTypeLowerName[343:366]: VirtualNetwork,
}

var _ResourceTypeNames = []string{
	_ResourceTypeName[0:22],
	_ResourceTypeName[22:36],
	_ResourceTypeName[36:69],
	_ResourceTypeName[69:110],
	_ResourceTypeName[110:142],
	_ResourceTypeName[142:173],
	_ResourceTypeName[173:199],
	_ResourceTypeName[199:224],
	_ResourceTypeName[224:254],
	_ResourceTypeName[254:277],
	_ResourceTypeName[277:310],
	_ResourceTypeName[310:343],
	_ResourceTypeName[343:366],
}

// ResourceTypeString retrieves an enum value from the enum constants string name.
// Throws an error if the param is not part of the enum.
func ResourceTypeString(s string) (ResourceType, error) {
	if val, ok := _ResourceTypeNameToValueMap[s]; ok {
		return val, nil
	}
	s = strings.ToLower(s)
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
