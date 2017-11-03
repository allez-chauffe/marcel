package screenshotService

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	"github.com/gorilla/mux"

	"github.com/Zenika/MARCEL/backend/commons"
)

const sep = string(os.PathSeparator)

//TakeScreenshot take a screen shot of the given media.
//It return the name of the newly created screenshot
func TakeScreenshot(mediaID int, screenshotDir string) string {
	media := strconv.Itoa(mediaID)
	containerName := "screenshot-" + media
	absoluteDirPath, _ := filepath.Abs(screenshotDir)

	cmd := exec.Command(
		"docker", "container", "run", "--rm", "--shm-size", "200m",
		"--name", containerName,
		"-v", absoluteDirPath+":/usr/src/app/screenshots/",
		"marcel-screenshot-service",
		"192.168.1.52", media,
	)

	var out []byte
	var err error

	log.Printf("Launching screenshot container %s", containerName)
	ok := commons.RunWithTimeout(10*time.Second, func() {
		out, err = cmd.Output()
	})

	if err != nil {
		switch e := err.(type) {
		case *exec.ExitError:
			log.Printf("Screenshot command error : %s", e.Stderr)
		case error:
			log.Printf("Error while taking screenshot : %s", e)
		}
		return ""
	}

	if !ok {
		log.Printf("Screenshot timeout ! Stopping container %s", containerName)

		cmd = exec.Command("docker", "rm", "-f", containerName)
		commons.RunWithTimeout(5*time.Second, func() {
			cmd.Run()
		})
		log.Printf("Screenshot container stoped")
		return ""
	}

	newScreenshot := string(out)
	log.Printf("New screenshot saved : %s", newScreenshot)

	return strings.Trim(newScreenshot, " \n")
}

func GetScreenshotHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	mediaID, ok := vars["idMedia"]

	if !ok {
		commons.WriteResponse(w, http.StatusBadRequest, "Media id not specified")
		return
	}

	screenshot, ok := vars["screenshot"]

	if !ok {
		commons.WriteResponse(w, http.StatusBadRequest, "Screenshot not specified")
		return
	}

	http.ServeFile(w, r, fmt.Sprintf("medias/%s/screenshot/%s", mediaID, screenshot))
}
