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
const ComputeImageIAMAssetType string = "compute.googleapis.com/Image"

func resourceConverterComputeImageIamPolicy() ResourceConverter {
	return ResourceConverter{
		AssetType:         ComputeImageIAMAssetType,
		Convert:           GetComputeImageIamPolicyCaiObject,
		MergeCreateUpdate: MergeComputeImageIamPolicy,
	}
}

func resourceConverterComputeImageIamBinding() ResourceConverter {
	return ResourceConverter{
		AssetType:         ComputeImageIAMAssetType,
		Convert:           GetComputeImageIamBindingCaiObject,
		FetchFullResource: FetchComputeImageIamPolicy,
		MergeCreateUpdate: MergeComputeImageIamBinding,
		MergeDelete:       MergeComputeImageIamBindingDelete,
	}
}

func resourceConverterComputeImageIamMember() ResourceConverter {
	return ResourceConverter{
		AssetType:         ComputeImageIAMAssetType,
		Convert:           GetComputeImageIamMemberCaiObject,
		FetchFullResource: FetchComputeImageIamPolicy,
		MergeCreateUpdate: MergeComputeImageIamMember,
		MergeDelete:       MergeComputeImageIamMemberDelete,
	}
}

func GetComputeImageIamPolicyCaiObject(d TerraformResourceData, config *Config) ([]Asset, error) {
	return newComputeImageIamAsset(d, config, expandIamPolicyBindings)
}

func GetComputeImageIamBindingCaiObject(d TerraformResourceData, config *Config) ([]Asset, error) {
	return newComputeImageIamAsset(d, config, expandIamRoleBindings)
}

func GetComputeImageIamMemberCaiObject(d TerraformResourceData, config *Config) ([]Asset, error) {
	return newComputeImageIamAsset(d, config, expandIamMemberBindings)
}

func MergeComputeImageIamPolicy(existing, incoming Asset) Asset {
	existing.IAMPolicy = incoming.IAMPolicy
	return existing
}

func MergeComputeImageIamBinding(existing, incoming Asset) Asset {
	return mergeIamAssets(existing, incoming, mergeAuthoritativeBindings)
}

func MergeComputeImageIamBindingDelete(existing, incoming Asset) Asset {
	return mergeDeleteIamAssets(existing, incoming, mergeDeleteAuthoritativeBindings)
}

func MergeComputeImageIamMember(existing, incoming Asset) Asset {
	return mergeIamAssets(existing, incoming, mergeAdditiveBindings)
}

func MergeComputeImageIamMemberDelete(existing, incoming Asset) Asset {
	return mergeDeleteIamAssets(existing, incoming, mergeDeleteAdditiveBindings)
}

func newComputeImageIamAsset(
	d TerraformResourceData,
	config *Config,
	expandBindings func(d TerraformResourceData) ([]IAMBinding, error),
) ([]Asset, error) {
	bindings, err := expandBindings(d)
	if err != nil {
		return []Asset{}, fmt.Errorf("expanding bindings: %v", err)
	}

	name, err := assetName(d, config, "//compute.googleapis.com/projects/{{project}}/global/images/{{image}}")
	if err != nil {
		return []Asset{}, err
	}

	return []Asset{{
		Name: name,
		Type: ComputeImageIAMAssetType,
		IAMPolicy: &IAMPolicy{
			Bindings: bindings,
		},
	}}, nil
}

func FetchComputeImageIamPolicy(d TerraformResourceData, config *Config) (Asset, error) {
	// Check if the identity field returns a value
	if _, ok := d.GetOk("{{image}}"); !ok {
		return Asset{}, ErrEmptyIdentityField
	}

	return fetchIamPolicy(
		ComputeImageIamUpdaterProducer,
		d,
		config,
		"//compute.googleapis.com/projects/{{project}}/global/images/{{image}}",
		ComputeImageIAMAssetType,
	)
}
