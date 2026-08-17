package iplocate_test

import (
	"testing"

	"github.com/babafemikuku/temporal-ip-geolocation/iplocate"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"go.temporal.io/sdk/testsuite"
)

func Test_Workflow(t *testing.T) {
	testsuite := &testsuite.WorkflowTestSuite{}
	env := testsuite.NewTestWorkflowEnvironment()
	activities := &iplocate.IPActivities{}

	env.OnActivity(activities.GetIP, mock.Anything).Return("1.1.1.1", nil)
	env.OnActivity(activities.GetLocationInfo, mock.Anything, "1.1.1.1").Return("Nigeria", nil)

	env.ExecuteWorkflow(iplocate.GetAddressFromIp, "Babafemi Kuku")

	var result string
	assert.NoError(t, env.GetWorkflowResult(&result))
	assert.Equal(t, "Hello Babafemi Kuku! Your IP is 1.1.1.1 and your location is Nigeria.", result)
}
