package appprotect

import (
	"fmt"
	"net"
	"regexp"
	"strconv"
	"strings"

	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
)



var dataProfilesIntMin = 0
var dataProfilesIntMax = 2147483647
var dataProfilesAllowedStr = "any"

var jsonprofilesSlicePath = []string{"spec","policy","json-profiles"}
var jsonprofilesPaths = [][]string{
	{"defenseAttributes","maximumArrayLength"},
	{"defenseAttributes","maximumStructureDepth"},
	{"defenseAttributes","maximumTotalLengthOfJSONData"},
	{"defenseAttributes","maximumValueLength"},
}

var xMLProfilesSlicePath = []string{"spec","policy","xml-profiles"}
var xMLProfilesPaths = [][]string{
	{"defenseAttributes","maximumAttributeValueLength"},
	{"defenseAttributes","maximumAttributesPerElement"},
	{"defenseAttributes","maximumChildrenPerElement"},
	{"defenseAttributes","maximumDocumentDepth"},
	{"defenseAttributes","maximumDocumentSize"},
	{"defenseAttributes","maximumElements"},
	{"defenseAttributes","maximumNSDeclarations"},
	{"defenseAttributes","maximumNameLength"},
	{"defenseAttributes","maximumNamespaceLength"},
} 

var maximumHeaderLengthIntMin = 1
var maximumHeaderLengthIntMax = 65536
var maximumHeaderLengthAllowedStr = "any"

var maximumHTTPHeaderLengthPath = []string{"spec","policy","header-settings","maximumHttpHeaderLength"}

var maximumCookieHeaderLengthPath = []string{"spec","policy","cookie-settings","maximumCookieHeaderLength"}

var scoreThresholdIntMin = 0
var scoreThresholdIntMax = 150
var scoreThresholdAllowedStr = "default"
var scoreThresholdAllowedSlicePath = []string{"spec","policy","bot-defense","mitigations","anomalies"}
var scoreThresholdAllowedPaths = [][]string{
	{"scoreThreshold"},
}

var appProtectPolicyRequiredFields = [][]string{
	{"spec", "policy"},
}

var appProtectLogConfRequiredFields = [][]string{
	{"spec", "content"},
	{"spec", "filter"},
}

var appProtectUserSigRequiredSlices = [][]string{
	{"spec", "signatures"},
}

func validateRequiredFields(policy *unstructured.Unstructured, fieldsList [][]string) error {
	for _, fields := range fieldsList {
		field, found, err := unstructured.NestedMap(policy.Object, fields...)
		if err != nil {
			return fmt.Errorf("Error checking for required field %v: %v", field, err)
		}
		if !found {
			return fmt.Errorf("Required field %v not found", field)
		}
	}
	return nil
}

func validateRequiredSlices(policy *unstructured.Unstructured, fieldsList [][]string) error {
	for _, fields := range fieldsList {
		field, found, err := unstructured.NestedSlice(policy.Object, fields...)
		if err != nil {
			return fmt.Errorf("Error checking for required field %v: %v", field, err)
		}
		if !found {
			return fmt.Errorf("Required field %v not found", field)
		}
	}
	return nil
}

func validateRequiredStrings(policy *unstructured.Unstructured, fieldsList [][]string) error {
	for _, fields := range fieldsList {
		field, found, err := unstructured.NestedString(policy.Object, fields...)
		if err != nil {
			return fmt.Errorf("Error checking for required field %v: %v", field, err)
		}
		if !found {
			return fmt.Errorf("Required field %v not found", field)
		}
	}
	return nil
}

// ValidateAppProtectPolicy validates Policy resource
func ValidateAppProtectPolicy(policy *unstructured.Unstructured) error {
	polName := policy.GetName()

	err := validateRequiredFields(policy, appProtectPolicyRequiredFields)
	if err != nil {
		return fmt.Errorf("Error validating App Protect Policy %s: %v", polName, err)
	}

	return nil
}

func validateIntOrStringFields(policy *unstructured.Unstructured) error {	
	err := validateIntorStringFieldsInSlice(policy.Object, jsonprofilesSlicePath, jsonprofilesPaths, dataProfilesAllowedStr, dataProfilesIntMin, dataProfilesIntMax)
	if err != nil {
		return err
	}
	err = validateIntorStringFieldsInSlice(policy.Object, xMLProfilesSlicePath, xMLProfilesPaths, dataProfilesAllowedStr, dataProfilesIntMin, dataProfilesIntMax)
	if err != nil {
		return err
	}
	err = validateIntorStringFieldsInSlice(policy.Object, scoreThresholdAllowedSlicePath, scoreThresholdAllowedPaths, scoreThresholdAllowedStr, scoreThresholdIntMin, scoreThresholdIntMax)
		if err != nil {
		return err
		}	
	err = validateIntOrStringFieldInMap(policy.Object, maximumHTTPHeaderLengthPath, maximumHeaderLengthAllowedStr, maximumHeaderLengthIntMin, maximumHeaderLengthIntMax)
	if err != nil {
		return err
	}
	err = validateIntOrStringFieldInMap(policy.Object, maximumCookieHeaderLengthPath, maximumHeaderLengthAllowedStr, maximumHeaderLengthIntMin, maximumHeaderLengthIntMax)
	if err != nil {
		return err
	}

	return nil
}

func validateIntorStringFieldsInSlice(fieldMap map[string]interface{}, slicePath []string, fieldPaths [][]string, allowedStrVal string, intMin, intMax int) error {
	policySlice, present, err := unstructured.NestedSlice(fieldMap, slicePath...)
	if err != nil {
		return fmt.Errorf("Error retrieving slice %s: %v", slicePath[len(slicePath)-1], err)
	}
	if present {
		for _, policySliceMap := range policySlice {
			for _, path := range fieldPaths {
				err = validateIntOrStringFieldInMap(policySliceMap.(map[string]interface{}), path, allowedStrVal, intMin, intMax)
				if err != nil {
					return err
				}
			}
		}
	}
	return nil
}

func validateIntOrStringFieldInMap(fieldMap map[string]interface{} , fieldPath []string, allowedStrVal string, intMin, intMax int) error {
	policyField, present, err := unstructured.NestedFieldNoCopy(fieldMap, fieldPath...)
	if err != nil {
		return fmt.Errorf("Error retrieving field %s : %v", fieldPath[len(fieldPath)-1], err)
	}
	if present {
		valid, err := validateIntOrStringField(policyField, allowedStrVal, intMin, intMax)
		if err != nil {
			return fmt.Errorf("Error validating field %s : %v", fieldPath[len(fieldPath)-1], err)
		}
		if ! valid {
			return fmt.Errorf("Error validating field %s : field must be string == \"%s\" or %d < int < %d", fieldPath[len(fieldPath)-1], allowedStrVal, intMin, intMax)
		}
	}
	return nil
}

func validateIntOrStringField(field interface{}, allowedStrVal string, intMin, intMax int) (bool, error) {
	switch f := field.(type) {
	case int:
		return (f > intMin && f < intMax), nil
	case string:
		return f == allowedStrVal, nil
	default:
		return false, fmt.Errorf("Unsupported type")
	}
}

// ValidateAppProtectLogConf validates LogConfiguration resource
func ValidateAppProtectLogConf(logConf *unstructured.Unstructured) error {
	lcName := logConf.GetName()
	err := validateRequiredFields(logConf, appProtectLogConfRequiredFields)
	if err != nil {
		return fmt.Errorf("Error validating App Protect Log Configuration %v: %v", lcName, err)
	}

	return nil
}

var logDstEx = regexp.MustCompile(`(?:syslog:server=((?:\d{1,3}\.){3}\d{1,3}|localhost):\d{1,5})|stderr|(?:\/[\S]+)+`)
var logDstFileEx = regexp.MustCompile(`(?:\/[\S]+)+`)

// ValidateAppProtectLogDestination validates destination for log configuration
func ValidateAppProtectLogDestination(dstAntn string) error {
	errormsg := "Error parsing App Protect Log config: Destination must follow format: syslog:server=<ip-address | localhost>:<port> or stderr or absolute path to file"
	if !logDstEx.MatchString(dstAntn) {
		return fmt.Errorf("%s Log Destination did not follow format", errormsg)
	}
	if dstAntn == "stderr" {
		return nil
	}

	if logDstFileEx.MatchString(dstAntn) {
		return nil
	}

	dstchunks := strings.Split(dstAntn, ":")

	// This error can be ignored since the regex check ensures this string will be parsable
	port, _ := strconv.Atoi(dstchunks[2])

	if port > 65535 || port < 1 {
		return fmt.Errorf("Error parsing port: %v not a valid port number", port)
	}

	ipstr := strings.Split(dstchunks[1], "=")[1]
	if ipstr == "localhost" {
		return nil
	}

	if net.ParseIP(ipstr) == nil {
		return fmt.Errorf("Error parsing host: %v is not a valid ip address", ipstr)
	}

	return nil
}

// ParseResourceReferenceAnnotation returns a namespace/name string
func ParseResourceReferenceAnnotation(ns, antn string) string {
	if !strings.Contains(antn, "/") {
		return ns + "/" + antn
	}
	return antn
}

func validateAppProtectUserSig(userSig *unstructured.Unstructured) error {
	sigName := userSig.GetName()
	err := validateRequiredSlices(userSig, appProtectUserSigRequiredSlices)
	if err != nil {
		return fmt.Errorf("Error validating App Protect User Signature %v: %v", sigName, err)
	}

	return nil
}

// GetNsName gets the key of a resource in the format: "resNamespace/resName"
func GetNsName(obj *unstructured.Unstructured) string {
	return obj.GetNamespace() + "/" + obj.GetName()
}
