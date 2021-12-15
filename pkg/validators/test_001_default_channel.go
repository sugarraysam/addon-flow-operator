package validators

import (
	"log"

	"github.com/mt-sre/addon-metadata-operator/internal/testutils"
	"github.com/mt-sre/addon-metadata-operator/pkg/utils"
)

// check interface implemented
var _ = utils.ValidatorTest(ValidatorTest001DefaultChannel{})

type ValidatorTest001DefaultChannel struct{}

func (v ValidatorTest001DefaultChannel) Name() string {
	return "Addon Default Channel Validator"
}

func (v ValidatorTest001DefaultChannel) Run(mb utils.MetaBundle) (bool, string, error) {
	return Validate001DefaultChannel(mb)
}

func (v ValidatorTest001DefaultChannel) SucceedingCandidates() []utils.MetaBundle {
	res, err := testutils.DefaultSucceedingCandidates()
	if err != nil {
		log.Fatalf("Could not load default succeeding candidates, got %v. Exiting.", err)
	}
	return res
}

// not implemented
func (v ValidatorTest001DefaultChannel) FailingCandidates() []utils.MetaBundle {
	return []utils.MetaBundle{}
}
