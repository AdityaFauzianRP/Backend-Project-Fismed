package service

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

type RequestData struct {
	Headers map[string]string      `json:"headers"`
	Body    map[string]interface{} `json:"body"`
}

var GlobalRequestData RequestData

func ParseRequest(c *gin.Context) (RequestData, error) {
	allHeaders := c.Request.Header

	var body map[string]interface{}
	if err := c.ShouldBindJSON(&body); err != nil {
		return RequestData{}, err
	}

	formattedHeaders := formatHeaders(allHeaders)

	return RequestData{
		Headers: formattedHeaders,
		Body:    body,
	}, nil
}

func formatHeaders(headers map[string][]string) map[string]string {
	formattedHeaders := make(map[string]string)
	for key, values := range headers {
		formattedHeaders[key] = values[0] // Mengambil nilai pertama dari slice
	}
	return formattedHeaders
}

func InterfaceToMapOrSliceOfMap(data interface{}) (interface{}, error) {
	switch v := data.(type) {
	case map[string]interface{}:
		result := make(map[string]string)
		for key, value := range v {
			if strValue, ok := value.(string); ok {
				result[key] = strValue
			} else {
				return nil, fmt.Errorf("value for key '%s' is not a string", key)
			}
		}
		return result, nil
	case []interface{}:
		var result []map[string]string
		for _, item := range v {
			if mapItem, ok := item.(map[string]interface{}); ok {
				tempMap := make(map[string]string)
				for key, value := range mapItem {
					if strValue, ok := value.(string); ok {
						tempMap[key] = strValue
					} else {
						return nil, fmt.Errorf("value for key '%s' is not a string", key)
					}
				}
				result = append(result, tempMap)
			} else {
				return nil, fmt.Errorf("item in array is not of type map[string]interface{}")
			}
		}
		return result, nil
	default:
		return nil, fmt.Errorf("data is neither map[string]interface{} nor []interface{}")
	}
}
