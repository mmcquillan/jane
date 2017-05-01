package actions

import (
	"log"
	"reflect"

	"github.com/hexbotio/hex/models"
	"github.com/hexbotio/hex/parse"
)

// Input interface
type Action interface {
	Act(action models.Action, message *models.Message, config *models.Config)
}

// List of Inputs
var List = make(map[string]reflect.Type)

func init() {
	List["format"] = reflect.TypeOf(Format{})
	List["ssh"] = reflect.TypeOf(Ssh{})
	List["winrm"] = reflect.TypeOf(WinRM{})
}

// Exists function
func Exists(connType string) (exists bool) {
	_, exists = List[connType]
	return exists
}

// MakeService
func Make(connType string) interface{} {
	if ct, ok := List[connType]; ok {
		c := (reflect.New(ct).Elem().Interface())
		return c
	} else {
		return nil
	}
}

// Recovery
func Recovery(service models.Service) {
	msg := "Panic - " + service.Name + " " + service.Type
	if r := recover(); r != nil {
		log.Print(msg, r)
	}
}

// ActionSuccess
func ActionSuccess(results string, success string, failure string) (state bool) {
	if results != "" && success != "" {
		if parse.Match(success, results) {
			state = true
		}
	}
	if results != "" && failure != "" {
		if parse.Match(failure, results) {
			state = false
		}
	}
	return state
}