package config

import "encoding/json"

func ConvertTo(request,category interface{}) (err error)  {
	dataByte,err :=json.Marshal(request)
	if err !=nil {
		return
	}
	return json.Unmarshal(dataByte,category)
}
