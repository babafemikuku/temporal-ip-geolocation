package iplocate

import (
	"fmt"
	"time"

	"go.temporal.io/sdk/temporal"
	"go.temporal.io/sdk/workflow"
)

// GetAddressFromIP is the Temporal Workflow that retrieves the IP address and location info.
func GetAddressFromIp(ctx workflow.Context, name string) (string, error) {
	ao := workflow.ActivityOptions{
		StartToCloseTimeout: time.Minute,
		RetryPolicy: &temporal.RetryPolicy{
			InitialInterval:    time.Second,
			MaximumInterval:    time.Minute,
			BackoffCoefficient: 2,
		},
	}

	ctx = workflow.WithActivityOptions(ctx, ao)

	var ipActivities *IPActivities

	var ip string
	if err := workflow.ExecuteActivity(ctx, ipActivities.GetIP).Get(ctx, &ip); err != nil {
		return "", fmt.Errorf("Failed to get IP address: %s", err)
	}

	var location string
	if err := workflow.ExecuteActivity(ctx, ipActivities.GetLocationInfo, ip).Get(ctx, &location); err != nil {
		return "", fmt.Errorf("Failed to get location: %s", err)
	}

	return fmt.Sprintf("Hello %s! Your IP is %s and your location is %s.", name, ip, location), nil
}
