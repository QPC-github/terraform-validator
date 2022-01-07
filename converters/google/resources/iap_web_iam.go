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

import "fmt"

// Provide a separate asset type constant so we don't have to worry about name conflicts between IAM and non-IAM converter files
const IapWebIAMAssetType string = "iap.googleapis.com/Web"

func resourceConverterIapWebIamPolicy() ResourceConverter {
	return ResourceConverter{
		AssetType:         IapWebIAMAssetType,
		Convert:           GetIapWebIamPolicyCaiObject,
		MergeCreateUpdate: MergeIapWebIamPolicy,
	}
}

func resourceConverterIapWebIamBinding() ResourceConverter {
	return ResourceConverter{
		AssetType:         IapWebIAMAssetType,
		Convert:           GetIapWebIamBindingCaiObject,
		FetchFullResource: FetchIapWebIamPolicy,
		MergeCreateUpdate: MergeIapWebIamBinding,
		MergeDelete:       MergeIapWebIamBindingDelete,
	}
}

func resourceConverterIapWebIamMember() ResourceConverter {
	return ResourceConverter{
		AssetType:         IapWebIAMAssetType,
		Convert:           GetIapWebIamMemberCaiObject,
		FetchFullResource: FetchIapWebIamPolicy,
		MergeCreateUpdate: MergeIapWebIamMember,
		MergeDelete:       MergeIapWebIamMemberDelete,
	}
}

func GetIapWebIamPolicyCaiObject(d TerraformResourceData, config *Config) ([]Asset, error) {
	return newIapWebIamAsset(d, config, expandIamPolicyBindings)
}

func GetIapWebIamBindingCaiObject(d TerraformResourceData, config *Config) ([]Asset, error) {
	return newIapWebIamAsset(d, config, expandIamRoleBindings)
}

func GetIapWebIamMemberCaiObject(d TerraformResourceData, config *Config) ([]Asset, error) {
	return newIapWebIamAsset(d, config, expandIamMemberBindings)
}

func MergeIapWebIamPolicy(existing, incoming Asset) Asset {
	existing.IAMPolicy = incoming.IAMPolicy
	return existing
}

func MergeIapWebIamBinding(existing, incoming Asset) Asset {
	return mergeIamAssets(existing, incoming, mergeAuthoritativeBindings)
}

func MergeIapWebIamBindingDelete(existing, incoming Asset) Asset {
	return mergeDeleteIamAssets(existing, incoming, mergeDeleteAuthoritativeBindings)
}

func MergeIapWebIamMember(existing, incoming Asset) Asset {
	return mergeIamAssets(existing, incoming, mergeAdditiveBindings)
}

func MergeIapWebIamMemberDelete(existing, incoming Asset) Asset {
	return mergeDeleteIamAssets(existing, incoming, mergeDeleteAdditiveBindings)
}

func newIapWebIamAsset(
	d TerraformResourceData,
	config *Config,
	expandBindings func(d TerraformResourceData) ([]IAMBinding, error),
) ([]Asset, error) {
	bindings, err := expandBindings(d)
	if err != nil {
		return []Asset{}, fmt.Errorf("expanding bindings: %v", err)
	}

	name, err := assetName(d, config, "//iap.googleapis.com/projects/{{project}}/iap_web")
	if err != nil {
		return []Asset{}, err
	}

	return []Asset{{
		Name: name,
		Type: IapWebIAMAssetType,
		IAMPolicy: &IAMPolicy{
			Bindings: bindings,
		},
	}}, nil
}

func FetchIapWebIamPolicy(d TerraformResourceData, config *Config) (Asset, error) {
	// Check if the identity field returns a value

	return fetchIamPolicy(
		IapWebIamUpdaterProducer,
		d,
		config,
		"//iap.googleapis.com/projects/{{project}}/iap_web",
		IapWebIAMAssetType,
	)
}
