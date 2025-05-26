package tests

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type graphiteDataPoints []struct {
	Target string `json:"target"`
	Tags   struct {
		Name string `json:"name"`
	} `json:"tags"`
	Datapoints [][2]any `json:"datapoints"`
}

func isMetricEmitted(client http.Client, req *http.Request, metricName string, expectedNumberOfTimes int) (bool, error) {
	res, err := client.Do(req)
	if err != nil {
		return false, err
	}
	defer res.Body.Close()
	if res.StatusCode != http.StatusOK {
		return false, fmt.Errorf("expected status code to be 200OK, got %d", res.StatusCode)
	}
	var graphiteDataPoints graphiteDataPoints
	err = json.NewDecoder(res.Body).Decode(&graphiteDataPoints)
	if err != nil {
		return false, err
	}
	actualNumberOfTimes := 0
	for _, v := range graphiteDataPoints {
		if v.Tags.Name == metricName {
			for _, vv := range v.Datapoints {
				if vv[0] != nil {
					actualNumberOfTimes++
				}
			}
		}
	}
	if actualNumberOfTimes >= expectedNumberOfTimes {
		return true, nil
	}
	return false, fmt.Errorf("metric: %s emitted %d time(s) out of %d", metricName, actualNumberOfTimes, expectedNumberOfTimes)
}
