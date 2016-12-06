package agenda

import (
	"encoding/json"
	"fmt"
	"github.com/Zenika/MARCEL/backend/auth"
	"github.com/gorilla/mux"
	"google.golang.org/api/calendar/v3"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"
)

var calendarService *calendar.Service

func GetNextEvents(w http.ResponseWriter, r *http.Request) {

	var err error = nil
	calendarService, err = calendar.New(auth.Google_API_client)

	if err != nil {
		log.Printf(err.Error() + "\n Requesting a new Client ID.")
		http.Redirect(w, r, "/api/v1/GoogleLogin", http.StatusTemporaryRedirect)
		return
	}

	var agendaId = os.Getenv("MARCEL_AGENDA_ID")

	//we want 'nbEvents' next events from today
	vars := mux.Vars(r)
	e := vars["nbEvents"]
	if e != "" {
		nbEvents, _ := strconv.Atoi(e)

		var startTime time.Time = time.Now() //today

		calendarEvents, err := calendarService.Events.List(agendaId).
			TimeMin(startTime.Format(time.RFC3339)).
			MaxResults(int64(nbEvents)).
			Do()

		if err != nil {
			fmt.Fprintln(w, err)
			return
		}

		if len(calendarEvents.Items) > 0 {
			for _, i := range calendarEvents.Items {
				var when string
				// If the DateTime is an empty string the Event is an all-day Event.
				// So only Date is available.
				if i.Start.DateTime != "" {
					when = i.Start.DateTime
				} else {
					when = i.Start.Date
				}
				fmt.Printf("%s (%s)\n", i.Summary, when)
			}
		} else {
			fmt.Printf("No upcoming events found.\n")
		}

		j, err := json.Marshal(calendarEvents)
		if err == nil {
			w.Write([]byte(j))
		}
	}
}
