package fallbackSeries

import (
	"github.com/go-graphite/carbonapi/expr/helper"
	"github.com/go-graphite/carbonapi/expr/interfaces"
	"github.com/go-graphite/carbonapi/expr/metadata"
	"github.com/go-graphite/carbonapi/expr/types"
	"github.com/go-graphite/carbonapi/pkg/parser"
)

func init() {
	metadata.RegisterFunction("fallbackSeries", &fallbackSeries{})
}

type fallbackSeries struct {
	interfaces.FunctionBase
}

// fallbackSeries( seriesList, fallback )
func (f *fallbackSeries) Do(e parser.Expr, from, until int32, values map[parser.MetricRequest][]*types.MetricData) ([]*types.MetricData, error) {
	/*
		Takes a wildcard seriesList, and a second fallback metric.
		If the wildcard does not match any series, draws the fallback metric.
	*/
	seriesList, err := helper.GetSeriesArg(e.Args()[0], from, until, values)
	fallback, errFallback := helper.GetSeriesArg(e.Args()[1], from, until, values)
	if errFallback != nil && err != nil {
		return nil, errFallback
	}

	if seriesList != nil && len(seriesList) > 0 {
		return seriesList, nil
	}
	return fallback, nil
}

// Description is auto-generated description, based on output of https://github.com/graphite-project/graphite-web
func (f *fallbackSeries) Description() map[string]*types.FunctionDescription {
	return map[string]*types.FunctionDescription{
		"fallbackSeries": {
			Description: "Takes a wildcard seriesList, and a second fallback metric.\nIf the wildcard does not match any series, draws the fallback metric.\n\nExample:\n\n.. code-block:: none\n\n  &target=fallbackSeries(server*.requests_per_second, constantLine(0))\n\nDraws a 0 line when server metric does not exist.",
			Function:    "fallbackSeries(seriesList, fallback)",
			Group:       "Special",
			Module:      "graphite.render.functions",
			Name:        "fallbackSeries",
			Params: []types.FunctionParam{
				{
					Name:     "seriesList",
					Required: true,
					Type:     types.SeriesList,
				},
				{
					Name:     "fallback",
					Required: true,
					Type:     types.SeriesList,
				},
			},
		},
	}
}
