package controller

import (
	"encoding/json"
	"net/http"
	"serviceX/src/model"
	"serviceX/src/model/other"
	"serviceX/src/model/stuff"
)

const (
	create string = "create"
	update string = "update"
	delete string = "delete"

	modelOther string = "other"
	modelStuff string = "stuff"
)

// AsyncMessage
func (c Controller) AsyncMessage(topic string, data []byte) (err error) {
	msg := new(model.Message)
	json.Unmarshal(data, msg)
	//Select model
	var m model.ModelI
	switch msg.Model {
	case modelOther:
		m = &other.Other{}
	case modelStuff:
		m = &stuff.Stuff{}
	}
	status := 0
	switch msg.Event {
	case create:
		json.Unmarshal(msg.Data, m)
		status, err = c.AsyncCreateCore(m)
	case update:
		dataMap := map[string]any{}
		json.Unmarshal(msg.Data, &dataMap)
		status, err = c.AsyncUpdateCore(dataMap, m)
	case delete:
		json.Unmarshal(msg.Data, m)
		status, err = c.AsyncDeleteCore(m)
	}
	if err != nil {
		//send notification
		
		//Only repeat if is a server problem
		if status != http.StatusInternalServerError {
			return nil
		}
	}
	return nil
}
