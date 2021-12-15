package validators

import (
	"encoding/json"
	"fmt"

	"github.com/mt-sre/addon-metadata-operator/pkg/utils"
	operatorsv1alpha1 "github.com/operator-framework/api/pkg/operators/v1alpha1"
	"github.com/operator-framework/operator-registry/pkg/registry"
)

func Validate007CSVInstallModes(mb utils.MetaBundle) (bool, string, error) {
	installMode := mb.AddonMeta.InstallMode

	for _, bundle := range mb.Bundles {
		bundleName, err := utils.GetBundleNameVersion(bundle)
		if err != nil {
			return false, "", err
		}
		spec, err := extractCSVSpec(bundle)
		if err != nil {
			return false, "", fmt.Errorf("Can't extract CSV for %v, got %v.", bundleName, err)
		}
		if success, failureMsg := isInstallModeSupported(spec.InstallModes, installMode); !success {
			return false, fmt.Sprintf("Bundle %v failed CSV validation: %v.", bundleName, failureMsg), nil
		}
	}
	return true, "", nil
}

func isInstallModeSupported(installModes []operatorsv1alpha1.InstallMode, target string) (bool, string) {
	var allSupported []string
	for _, im := range installModes {
		imType := string(im.Type)
		if im.Supported {
			allSupported = append(allSupported, imType)
		}

		if imType == target {
			return im.Supported, ""
		}
	}
	return false, fmt.Sprintf("Target installMode %v is not supported. CSV only supports these installModes %v.", target, allSupported)
}

type CSVSpec struct {
	InstallModes []operatorsv1alpha1.InstallMode `json:"installModes"`
}

func extractCSVSpec(b registry.Bundle) (*CSVSpec, error) {
	csv, err := b.ClusterServiceVersion()
	if err != nil {
		return nil, err
	}
	var res *CSVSpec
	if err := json.Unmarshal(csv.Spec, res); err != nil {
		return nil, err
	}

	return res, nil
}
