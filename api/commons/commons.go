package commons

import (
	"encoding/json"
	"io"
	"net/http"
	"os"
	"reflect"

	log "github.com/sirupsen/logrus"
)

func IsInArray(val interface{}, array interface{}) (exists bool, index int) {
	return FindIndexInArray(
		func(curr interface{}) bool {
			return reflect.DeepEqual(val, curr)
		},
		array,
	)
}

func FindIndexInArray(predicate func(val interface{}) bool, array interface{}) (exists bool, index int) {
	exists = false
	index = -1

	switch reflect.TypeOf(array).Kind() {
	case reflect.Slice:
		s := reflect.ValueOf(array)

		for i := 0; i < s.Len(); i++ {
			if predicate(s.Index(i).Interface()) {
				index = i
				exists = true
				return
			}
		}
	}

	return
}

// Thx to https://www.socketloop.com/tutorials/golang-copy-directory-including-sub-directories-files
func CopyFile(source string, dest string) (err error) {
	sourcefile, err := os.Open(source)
	if err != nil {
		return err
	}

	defer sourcefile.Close()

	destfile, err := os.Create(dest)
	if err != nil {
		return err
	}

	defer destfile.Close()

	_, err = io.Copy(destfile, sourcefile)
	if err == nil {
		si, err := os.Stat(source)
		if err == nil {
			err = os.Chmod(dest, si.Mode())
		}
	}

	return
}

// Thx to https://www.socketloop.com/tutorials/golang-copy-directory-including-sub-directories-files
func CopyDir(source string, dest string) (err error) {

	// get properties of source dir
	sourceinfo, err := os.Stat(source)
	if err != nil {
		return err
	}

	// create dest dir
	err = os.MkdirAll(dest, sourceinfo.Mode())
	if err != nil {
		return err
	}

	directory, _ := os.Open(source)

	objects, err := directory.Readdir(-1)

	for _, obj := range objects {
		sourcefilepointer := source + "/" + obj.Name()
		destinationfilepointer := dest + "/" + obj.Name()

		if obj.IsDir() {
			// create sub-directories - recursively
			err = CopyDir(sourcefilepointer, destinationfilepointer)
			if err != nil {
				log.Errorln(err)
			}
		} else {
			// perform copy
			err = CopyFile(sourcefilepointer, destinationfilepointer)
			if err != nil {
				log.Errorln(err)
			}
		}

	}
	return
}

func WriteResponse(w http.ResponseWriter, statusCode int, message string) {
	w.WriteHeader(statusCode)
	w.Write([]byte(" " + message))
}

func WriteJsonResponse(w http.ResponseWriter, body interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(body); err != nil {
		log.Errorf("Error while send JSON data (%s)", err.Error())
	}
}
