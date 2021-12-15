package validators

import (
	"log"

	"github.com/mt-sre/addon-metadata-operator/internal/testutils"
	"github.com/mt-sre/addon-metadata-operator/pkg/utils"
)

// check interface implemented
var _ = utils.ValidatorTest(Validator003CSVNames{})

type Validator003CSVNames struct{}

func (v Validator003CSVNames) Name() string {
	return "003_csv_names"
}

func (v Validator003CSVNames) Run(mb utils.MetaBundle) (bool, string, error) {
	return Validate003CSVNames(mb)
}

func (v Validator003CSVNames) SucceedingCandidates() []utils.MetaBundle {
	res, err := testutils.DefaultSucceedingCandidates()
	if err != nil {
		log.Fatalf("Could not load default succeeding candidates, got %v. Exiting.", err)
	}
	return res
}

// not implemented
func (v Validator003CSVNames) FailingCandidates() []utils.MetaBundle {
	return []utils.MetaBundle{}
}
