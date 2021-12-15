package validators

import (
	"log"

	"github.com/mt-sre/addon-metadata-operator/internal/testutils"
	"github.com/mt-sre/addon-metadata-operator/pkg/utils"
)

// check interface implemented
var _ = utils.ValidatorTest(Validator007CSVInstallModes{})

type Validator007CSVInstallModes struct{}

func (v Validator007CSVInstallModes) Name() string {
	return "007_csv_install_modes"
}

func (v Validator007CSVInstallModes) Run(mb utils.MetaBundle) (bool, string, error) {
	return Validate007CSVInstallModes(mb)
}

func (v Validator007CSVInstallModes) SucceedingCandidates() []utils.MetaBundle {
	res, err := testutils.DefaultSucceedingCandidates()
	if err != nil {
		log.Fatalf("Could not load default succeeding candidates, got %v. Exiting.", err)
	}
	return res
}

// not implemented
func (v Validator007CSVInstallModes) FailingCandidates() []utils.MetaBundle {
	return []utils.MetaBundle{}
}
