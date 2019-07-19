package function

import (
	"fmt"
	"regexp"
	"time"

	"github.com/mackerelio/mackerel-client-go"
	"github.com/tenntenn/natureremo"
)

var metricNamePattern = regexp.MustCompile(`\A[a-zA-Z0-9._-]+\z`)
var window = 5 * time.Minute

type devices []*natureremo.Device

func (ds devices) ToMetricValues() []*mackerel.MetricValue {
	var values []*mackerel.MetricValue
	now := time.Now()

	for _, d := range ds {
		var deviceName string
		if metricNamePattern.Match([]byte(d.Name)) {
			deviceName = d.Name
		} else {
			deviceName = d.ID
		}

		for s, v := range d.NewestEvents {
			sensorLabel, boolValue := sensorLabelAndType(s)
			name := fmt.Sprintf("%s.%s.%s", c.MetricPrefix, sensorLabel, deviceName)

			var time int64
			if c.UseSensorEventTime {
				time = v.CreatedAt.Unix()
			} else {
				time = now.Unix()
			}

			var value interface{}
			if boolValue {
				if now.Sub(v.CreatedAt) <= window {
					value = 1
				} else {
					value = 0
				}
			} else {
				value = v.Value
			}

			values = append(values, &mackerel.MetricValue{
				Name:  name,
				Time:  time,
				Value: value,
			})
		}
	}

	return values
}

func sensorLabelAndType(s natureremo.SensorType) (string, bool) {
	var (
		label     string
		boolValue bool
	)

	switch s {
	case "te":
		label = "temperature"
		boolValue = false
	case "hu":
		label = "humidity"
		boolValue = false
	case "il":
		label = "illumination"
		boolValue = false
	case "mo":
		label = "motion"
		boolValue = true
	default:
		label = string(s)
		boolValue = false
	}
	return label, boolValue
}
