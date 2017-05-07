package web

import (
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"mycap/libs/stat/duration"
	"mycap/libs/stat/rps"
	"time"
)

type PlotItem [2]interface{}
type PlotData []PlotItem
type PlotDataMap map[string]interface{}
type PlotDataMaps []PlotDataMap

func RenderPlotRps(items []int64) template.JS {
	result := make(PlotData, len(items))
	_, offset := time.Now().In(time.Local).Zone()

	for key, val := range items {
		result[key] = PlotItem{1000 * (time.Now().Unix() + int64(offset) + (int64(key))), val}
	}

	if result_js, err := json.Marshal(result); err == nil {
		return template.JS(result_js)
	} else {
		log.Println(err)
		return template.JS("")
	}
}

func RenderPlotRpsAvg(items rps.Items, stepSize int64) template.JS {
	min := make(PlotData, len(items))
	max := make(PlotData, len(items))
	avg := make(PlotData, len(items))

	result := PlotDataMaps{
		make(PlotDataMap),
		make(PlotDataMap),
		make(PlotDataMap),
	}

	_, offset := time.Now().In(time.Local).Zone()

	zeroTime := time.Now().Unix() - int64(len(items))*stepSize + int64(offset)

	for key, val := range items {
		min[key] = PlotItem{1000 * (zeroTime + (int64(key) * stepSize)), val.Min}
		max[key] = PlotItem{1000 * (zeroTime + (int64(key) * stepSize)), val.Max}
		avg[key] = PlotItem{1000 * (zeroTime + (int64(key) * stepSize)), val.Avg}
	}

	result[0]["data"] = min
	result[0]["label"] = "min"
	result[0]["lines"] = map[string]interface{}{"show": "true"}

	result[1]["data"] = max
	result[1]["label"] = "max"
	result[1]["lines"] = map[string]interface{}{"show": "true"}

	result[2]["data"] = avg
	result[2]["label"] = "avg"
	result[2]["lines"] = map[string]interface{}{"show": "true"}

	if result_js, err := json.Marshal(result); err == nil {
		return template.JS(result_js)
	} else {
		log.Println(err)
		return template.JS("")
	}
}

func RenderPlotDuration(items []float64) template.JS {
	result := make(PlotData, len(items))
	_, offset := time.Now().In(time.Local).Zone()

	for key, val := range items {
		result[key] = PlotItem{
			1000 * (time.Now().Unix() + int64(offset) + (int64(key))),
			fmt.Sprintf("%.3f", val),
		}
	}

	if result_js, err := json.Marshal(result); err == nil {
		return template.JS(result_js)
	} else {
		log.Println(err)
		return template.JS("")
	}
}

func RenderPlotDurationAvg(items duration.Items, stepSize int64) template.JS {
	min := make(PlotData, len(items))
	max := make(PlotData, len(items))
	avg := make(PlotData, len(items))

	result := PlotDataMaps{
		make(PlotDataMap),
		make(PlotDataMap),
		make(PlotDataMap),
	}

	_, offset := time.Now().In(time.Local).Zone()

	zeroTime := time.Now().Unix() - int64(len(items))*stepSize + int64(offset)

	for key, val := range items {
		min[key] = PlotItem{1000 * (zeroTime + (int64(key) * stepSize)), fmt.Sprintf("%.3f", val.Min)}
		max[key] = PlotItem{1000 * (zeroTime + (int64(key) * stepSize)), fmt.Sprintf("%.3f", val.Max)}
		avg[key] = PlotItem{1000 * (zeroTime + (int64(key) * stepSize)), fmt.Sprintf("%.3f", val.Avg)}
	}

	result[0]["data"] = min
	result[0]["label"] = "min"
	result[0]["lines"] = map[string]interface{}{"show": "true"}

	result[1]["data"] = max
	result[1]["label"] = "max"
	result[1]["lines"] = map[string]interface{}{"show": "true"}

	result[2]["data"] = avg
	result[2]["label"] = "avg"
	result[2]["lines"] = map[string]interface{}{"show": "true"}

	if result_js, err := json.Marshal(result); err == nil {
		return template.JS(result_js)
	} else {
		log.Println(err)
		return template.JS("")
	}
}
