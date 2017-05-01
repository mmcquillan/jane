package commands

import (
	"log"
	"reflect"

	"github.com/hexbotio/hex/models"
)

// Input interface
type Action interface {
	Act(message *models.Message, config *models.Config)
}

// List of Inputs
var List = make(map[string]reflect.Type)

func init() {
	List["help*"] = reflect.TypeOf(Help{})
	List["passwd"] = reflect.TypeOf(Passwd{})
	List["ping"] = reflect.TypeOf(Ping{})
	List["uptime"] = reflect.TypeOf(Uptime{})
	List["version"] = reflect.TypeOf(Version{})
	List["whoami"] = reflect.TypeOf(Whoami{})
}

// commandHelp function
func CommandHelp(config *models.Config) (command []string) {
	command = make([]string, 6)
	command[0] = "help <filter> - This help"
	command[1] = "passwd - Password generator"
	command[2] = "ping - Simple ping response for the bot"
	command[3] = "uptime - Number of seconds process has been running"
	command[4] = "version - Compiled version number/sha"
	command[5] = "whoami - Your user name"
	return command
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