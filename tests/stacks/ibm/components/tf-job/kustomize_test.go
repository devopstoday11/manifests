package tf_job

import (
	"github.com/kubeflow/manifests/tests"
	"testing"
)

func TestKustomize(t *testing.T) {
	testCase := &tests.KustomizeTestCase{
		Package:  "../../../../../stacks/ibm/components/tf-job",
		Expected: "test_data/expected",
	}

	tests.RunTestCase(t, testCase)
}
