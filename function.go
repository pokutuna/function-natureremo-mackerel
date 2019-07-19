package function

import (
	"context"

	"github.com/kelseyhightower/envconfig"
	"github.com/mackerelio/mackerel-client-go"
	"github.com/tenntenn/natureremo"
)

type config struct {
	NatureRemoAccessToken string `required:"true" envconfig:"NATUREREMO_ACCESS_TOKEN"`
	MackerelAPIKey        string `required:"true" envconfig:"MACKEREL_API_KEY"`
	ServiceName           string `required:"true" envconfig:"SERVICE_NAME"`
	MetricPrefix          string `default:"natureremo" envconfig:"METRIC_PREFIX"`
	UseSensorEventTime    bool   `default:"false" envconfig:"USE_SENSOR_EVENT_TIME"`
}

var c config

// RemoToMackerel is a entry point for Cloud Functions
func RemoToMackerel(ctx context.Context, m interface{}) error {
	if err := envconfig.Process("", &c); err != nil {
		return err
	}

	remo := natureremo.NewClient(c.NatureRemoAccessToken)
	ds, err := remo.DeviceService.GetAll(ctx)
	if err != nil {
		return err
	}

	mkr := mackerel.NewClient(c.MackerelAPIKey)
	err = mkr.PostServiceMetricValues(c.ServiceName, devices(ds).ToMetricValues())
	if err != nil {
		return err
	}

	return nil
}
