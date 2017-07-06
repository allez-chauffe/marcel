package medias

import (
	"errors"
	"fmt"
	"reflect"
)

// MediasConfig encapsulates all configuration data
//
// swagger:model
type MediasConfiguration struct {
	LastID   int     `json:"last_id"`
	Medias   []Media `json:"medias"`
}

// Media represents a media configuration
//
// swagger:model
type Media struct {
	// the id for this media
	//
	// required: true
	// unique: true
	// min: 1
	ID          string                 `json:"id"`
	Name        string                 `json:"name"`
	Description string                 `json:"description"`
	Rows        int                    `json:"rows"`
	Cols        int                    `json:"cols"`
	Stylesvar   map[string]interface{} `json:"stylesvar"`
	Plugins     []MediaPlugin          `json:"plugins"`
}

// MediaPlugin represents a plugin configuration for the media
//
// Properties and configuration for a plugin used in the media
//
// swagger:model
type MediaPlugin struct {
	InstanceId string              `json:"instanceId"`
	Name       string              `json:"name"`
	FrontEnd   MediaPluginFrontEnd `json:"frontend"`
	BackEnd    MediaPluginBackEnd  `json:"backend"`
}

type MediaPluginFrontEnd struct {
	Files   []string               `json:"files"`
	EltName string                 `json:"eltName"`
	X       int                    `json:"x"`
	Y       int                    `json:"y"`
	Rows    int                    `json:"rows"`
	Cols    int                    `json:"cols"`
	Props   map[string]interface{} `json:"props"`
}

type MediaPluginBackEnd struct {
	Ports []int                  `json:"ports"`
	Props map[string]interface{} `json:"props"`
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
