package validators

import (
	"fmt"
	"strings"

	"github.com/mt-sre/addon-metadata-operator/pkg/utils"
	"github.com/operator-framework/operator-registry/pkg/registry"
	"golang.org/x/mod/semver"
)

func Validate003CSVNames(mb utils.MetaBundle) (bool, string, error) {
	operatorName := mb.AddonMeta.OperatorName

	for _, bundle := range mb.Bundles {
		bundleName, err := utils.GetBundleNameVersion(bundle)
		if err != nil {
			return false, "", err
		}
		csv, err := bundle.ClusterServiceVersion()
		if err != nil {
			return false, "", fmt.Errorf("Can't get CSV for %v, got %v.", bundleName, err)
		}
		if success, failureMsg := checkCSVNames(csv, operatorName); !success {
			return false, fmt.Sprintf("Bundle %v failed CSV validation: %v.", bundleName, failureMsg), nil
		}
	}
	return true, "", nil
}

func checkCSVNames(csv *registry.ClusterServiceVersion, operatorName string) (bool, string) {
	// also validate replaces if present, ignore errors
	csvNames := []string{csv.Name}
	if replaces, err := csv.GetReplaces(); err != nil && replaces != "" {
		csvNames = append(csvNames, replaces)
	}

	for _, csvName := range csvNames {
		parts := strings.SplitN(csvName, ".", 2)
		if len(parts) != 2 {
			return false, fmt.Sprintf("Could not split the csvName '%v' in two parts.", csvName)
		}
		if parts[0] != operatorName {
			return false, fmt.Sprintf("csvName '%v' first part does not match the operatorName '%v'.", csvName, operatorName)
		}

		// semver requires v prefix, make sure we only add it once
		versionNoV := strings.TrimPrefix(parts[1], "v")
		if !semver.IsValid(fmt.Sprintf("v%v", versionNoV)) {
			return false, fmt.Sprintf("csvName '%v' second part is an invalid semver 'v%v'.", csvName, versionNoV)
		}
	}
	return true, ""
}
