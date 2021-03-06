package types

import (
	"testing"
	"time"

	"github.com/bookingcom/carbonapi/pkg/types"
)

func TestMarshalCSVInUTC(t *testing.T) {
	results := []*MetricData{
		&MetricData{
			Metric: types.Metric{
				Name:      "foo",
				StartTime: 0,
				StopTime:  1,
				StepTime:  1,
				Values:    []float64{2},
				IsAbsent:  []bool{false},
			},
		},
	}

	blob := MarshalCSV(results, time.UTC)
	got := string(blob)

	exp := "\"foo\",1970-01-01 00:00:00,2\n"

	if got != exp {
		t.Errorf("Expected '%s', got '%s'", exp, got)
	}
}

func TestMarshalCSVNotInUTC(t *testing.T) {
	results := []*MetricData{
		&MetricData{
			Metric: types.Metric{
				Name:      "foo",
				StartTime: 0,
				StopTime:  1,
				StepTime:  1,
				Values:    []float64{2},
				IsAbsent:  []bool{false},
			},
		},
	}

	tz := time.FixedZone("UTC+1", int(time.Hour/time.Second))
	blob := MarshalCSV(results, tz)
	got := string(blob)

	exp := "\"foo\",1970-01-01 01:00:00,2\n"

	if got != exp {
		t.Errorf("Expected '%s', got '%s'", exp, got)
	}
}
