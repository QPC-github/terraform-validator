// ----------------------------------------------------------------------------
//
//     ***     AUTO GENERATED CODE    ***    Type: MMv1     ***
//
// ----------------------------------------------------------------------------
//
//     This file is automatically generated by Magic Modules and manual
//     changes will be clobbered when the file is regenerated.
//
//     Please read more about how to change this file in
//     .github/CONTRIBUTING.md.
//
// ----------------------------------------------------------------------------

package google

import "reflect"

const ApigeeOrganizationAssetType string = "apigee.googleapis.com/Organization"

func resourceConverterApigeeOrganization() ResourceConverter {
	return ResourceConverter{
		AssetType: ApigeeOrganizationAssetType,
		Convert:   GetApigeeOrganizationCaiObject,
	}
}

func GetApigeeOrganizationCaiObject(d TerraformResourceData, config *Config) ([]Asset, error) {
	name, err := assetName(d, config, "//apigee.googleapis.com/organizations/{{name}}")
	if err != nil {
		return []Asset{}, err
	}
	if obj, err := GetApigeeOrganizationApiObject(d, config); err == nil {
		return []Asset{{
			Name: name,
			Type: ApigeeOrganizationAssetType,
			Resource: &AssetResource{
				Version:              "v1",
				DiscoveryDocumentURI: "https://www.googleapis.com/discovery/v1/apis/apigee/v1/rest",
				DiscoveryName:        "Organization",
				Data:                 obj,
			},
		}}, nil
	} else {
		return []Asset{}, err
	}
}

func GetApigeeOrganizationApiObject(d TerraformResourceData, config *Config) (map[string]interface{}, error) {
	obj := make(map[string]interface{})
	displayNameProp, err := expandApigeeOrganizationDisplayName(d.Get("display_name"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("display_name"); !isEmptyValue(reflect.ValueOf(displayNameProp)) && (ok || !reflect.DeepEqual(v, displayNameProp)) {
		obj["displayName"] = displayNameProp
	}
	descriptionProp, err := expandApigeeOrganizationDescription(d.Get("description"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("description"); !isEmptyValue(reflect.ValueOf(descriptionProp)) && (ok || !reflect.DeepEqual(v, descriptionProp)) {
		obj["description"] = descriptionProp
	}
	analyticsRegionProp, err := expandApigeeOrganizationAnalyticsRegion(d.Get("analytics_region"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("analytics_region"); !isEmptyValue(reflect.ValueOf(analyticsRegionProp)) && (ok || !reflect.DeepEqual(v, analyticsRegionProp)) {
		obj["analyticsRegion"] = analyticsRegionProp
	}
	authorizedNetworkProp, err := expandApigeeOrganizationAuthorizedNetwork(d.Get("authorized_network"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("authorized_network"); !isEmptyValue(reflect.ValueOf(authorizedNetworkProp)) && (ok || !reflect.DeepEqual(v, authorizedNetworkProp)) {
		obj["authorizedNetwork"] = authorizedNetworkProp
	}
	runtimeTypeProp, err := expandApigeeOrganizationRuntimeType(d.Get("runtime_type"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("runtime_type"); !isEmptyValue(reflect.ValueOf(runtimeTypeProp)) && (ok || !reflect.DeepEqual(v, runtimeTypeProp)) {
		obj["runtimeType"] = runtimeTypeProp
	}
	billingTypeProp, err := expandApigeeOrganizationBillingType(d.Get("billing_type"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("billing_type"); !isEmptyValue(reflect.ValueOf(billingTypeProp)) && (ok || !reflect.DeepEqual(v, billingTypeProp)) {
		obj["billingType"] = billingTypeProp
	}
	runtimeDatabaseEncryptionKeyNameProp, err := expandApigeeOrganizationRuntimeDatabaseEncryptionKeyName(d.Get("runtime_database_encryption_key_name"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("runtime_database_encryption_key_name"); !isEmptyValue(reflect.ValueOf(runtimeDatabaseEncryptionKeyNameProp)) && (ok || !reflect.DeepEqual(v, runtimeDatabaseEncryptionKeyNameProp)) {
		obj["runtimeDatabaseEncryptionKeyName"] = runtimeDatabaseEncryptionKeyNameProp
	}

	return resourceApigeeOrganizationEncoder(d, config, obj)
}

func resourceApigeeOrganizationEncoder(d TerraformResourceData, meta interface{}, obj map[string]interface{}) (map[string]interface{}, error) {
	obj["name"] = d.Get("project_id").(string)
	return obj, nil
}

func expandApigeeOrganizationDisplayName(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandApigeeOrganizationDescription(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandApigeeOrganizationAnalyticsRegion(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandApigeeOrganizationAuthorizedNetwork(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandApigeeOrganizationRuntimeType(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandApigeeOrganizationBillingType(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandApigeeOrganizationRuntimeDatabaseEncryptionKeyName(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}
