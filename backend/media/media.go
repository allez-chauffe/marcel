package media

import (
	"errors"
	"fmt"
	"reflect"
)
/**
The global attributes for a Media
 */
type Media struct {
	ID      int    `json:"id"`
	Name    string `json:"name"`
	Config  MediaConfig `json:"config"`
	Components []MediaPlugin `json:"components"`
}

type MediaConfig struct {
	Styles []string `json:"styles"`
}

/**
Properties and configuration for a plugin used in the media
 */
type MediaPlugin struct {
	ComponentName       string `json:"componentName"`
	EltName    string `json:"eltName"`
	Files      []string `json:"files"`
	PropValues map[string]interface{} `json:"propValues"` //MediaPluginProps `json:"propValues"`
}

/**
Because we don't know what will compounds the props for a plugin, we use a map[string] interface{}
 */
type MediaPluginProps struct {
	X map[string]interface{} `json:"-"` //map[string]string
}

func SetField(obj interface{}, name string, value interface{}) error {
	structValue := reflect.ValueOf(obj).Elem()
	structFieldValue := structValue.FieldByName(name)

	if !structFieldValue.IsValid() {
		return fmt.Errorf("No such field: %s in obj", name)
	}

	if !structFieldValue.CanSet() {
		return fmt.Errorf("Cannot set %s field value", name)
	}

	structFieldType := structFieldValue.Type()
	val := reflect.ValueOf(value)
	if structFieldType != val.Type() {
		return errors.New("Provided value type didn't match obj field type")
	}

	structFieldValue.Set(val)
	return nil
}


func (s *Media) FillStruct(m map[string]interface{}) error {
	for k, v := range m {
		err := SetField(s, k, v)
		if err != nil {
			return err
		}
	}
	return nil
}