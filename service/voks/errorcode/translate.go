package errorcode

import "encoding/json"

var code = map[string]string{
	"ERROR_CUSTOMER_AND_CLUSTER_NOT_BLANK_OR_NULL":                  "Customer and cluster must not be blank or null.",
	"ERROR_CLUSTER_KUBERNETES_NOT_FOUND":                            "Kubernetes cluster not found.",
	"ERROR_VPC_NOT_FOUND":                                           "VPC not found.",
	"ERROR_CONSOLE_VM":                                              "Cannot perform this action at this time.",
	"ERROR_ADDON_NOT_FOUND":                                         "Add-on not found.",
	"ERROR_VPC_CUSTOMER_NOT_FOUND":                                  "Customer's VPC not found.",
	"ERROR_CLUSTER_KUBERNETES_STATUS_IS_NOT_SUCCESS":                "Kubernetes cluster status is not successful.",
	"ERROR_ADDON_VERSION_NOT_FOUND":                                 "Add-on version not found.",
	"ERROR_CANNOT_INSTALL_ADDON":                                    "Add-on can not be installed.",
	"ERROR_CANNOT_UNINSTALL_ADDON":                                  "Add-on cannot be uninstalled.",
	"ERROR_ADDON_NAME_IS_REQUIRED":                                  "Add-on name is required.",
	"ERROR_ADDON_VERSION_IS_REQUIRED":                               "Add-on version is required.",
	"ERROR_CLUSTER_UPDATING_SO_NOT_EXECUTE":                         "Cluster update in progress. Operation cannot proceed.",
	"ERROR_KEY_OF_LABEL_AND_TAINT_MUST_LOWER_THAN_63_CHARACTERS":    "Key of label and taint must be fewer than 63 characters.",
	"ERROR_VALUE_OF_LABEL_AND_TAINT_MUST_LOWER_THAN_63_CHARACTERS":  "Value of label and taint must be fewer than 63 characters.",
	"ERROR_CUSTOMER_AND_CLUSTER_NOT_MATCH":                          "Customer and cluster do not match.",
	"ERROR_NAME_OF_NODE_GROUP_INVALID":                              "Node group name is invalid.",
	"ERROR_NAME_NODE_GROUP_EXISTED":                                 "Node group name already exists.",
	"ERROR_NAME_NODE_GROUP_INVALID":                                 "Node group name is invalid.",
	"ERROR_EFFECT_OF_TAINT_NODE_GROUP_INVALID":                      "Effect of taint on node group is invalid.",
	"ERROR_MIN_NODE_MUST_BE_GREATER_THAN_0_AND_LOWER_THAN_MAX_NODE": "Minimum number of nodes must be greater than 0 and less than maximum nodes.",
	"ERROR_MAX_NODE_MUST_BE_LOWER_THAN_OR_EQUAL_TO_10":              "Maximum number of nodes must be less than or equal to 10.",
	"ERROR_CODE_RESOURCE_TYPE_NOT_EXIST":                            "Resource type code does not exist.",
	"ERROR_VALIDATE_RESOURCE":                                       "Resource validation failed.",
	"ERROR_NODE_GROUP_HAS_BEEN_ERROR_AUTO_REPAIR":                   "Node group has encountered an error and is being auto-repaired.",
	"ERROR_MIN_NODE_NOT_CHANGEABLE_SAVING_PLAN_AUTOSCALING":         "Minimum node count cannot be changed due to saving plan or autoscaling settings.",
	"ERROR_NODE_GROUP_NOT_EXIST_IN_THIS_CLUSTER":                    "Node group not found in this cluster",
	"ERROR_NODE_GROUP_ID_NOT_EXIST":                                 "Node group ID does not exist.",
	"ERROR_field-name_IS_REQUIRED":                                  "ERROR_field-The field 'field-name' is required",
	"ERROR_QUANTITY_OF_LABEL_AND_TAINT_EXCEEDED_LIMITATION_50":      "Key of label and taint is invalid.",
	"ERROR_KEY_OF_LABEL_AND_TAINT_IS_INVALIDED":                     "Value of label and taint is invalid.",
	"ERROR_VALUE_OF_LABEL_AND_TAINT_IS_INVALIDED":                   "The value of Label or Taint is invalid.",
	"ERROR_ADD_ONS_NFS_STORAGE_NOT_IN_ENABLED_RANGE":                "NFS storage add-on is not within the enabled range.",
	"ERROR_CLUSTER_CAN_NOT_EXECUTE":                                 "Cluster update in progress. Operation cannot proceed.",
	"ERROR_NFS_STORAGE_EXCEED_LIMITATION":                           "NFS storage exceeds the allowed limit.",
	"ERROR_VOLUME_NOT_FOUND":                                        "Volume not found.",
}

type ErrorModel struct {
	Success bool       `json:"success"`
	Errors  []ErrorKey `json:"errors"`
}

type ErrorKey struct {
	Key   string   `json:"key"`
	Param []string `json:"param"`
}

func Translate(body []byte) string {
	errorRes := &ErrorModel{}
	if json.Unmarshal(body, errorRes) != nil {
		return string(body)
	}

	if len(errorRes.Errors) > 0 {
		if value, ok := code[errorRes.Errors[0].Key]; ok {
			return value
		}
	}

	return string(body)
}
