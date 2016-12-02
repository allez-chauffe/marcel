package agenda

import (
	"encoding/json"
	"fmt"
	"google.golang.org/api/calendar/v3"
	"log"
	"github.com/Zenika/MARCEL/backend/auth"
	"net/http"
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

	//var nbEvents int64 = 2
	//we want events for today only
	var startTime time.Time = time.Now() //today
	var endTime time.Time = time.Date(
		startTime.Year(),
		startTime.Month(),
		startTime.Day(),
		23, 59, 59, 0,
		startTime.Location()) //end of today

	calendarEvents, err := calendarService.Events.List("zenika.com_h6tsmhv6iqs0e3l0vrugi6tbfg@group.calendar.google.com").
		TimeMin(startTime.Format(time.RFC3339)).
		TimeMax(endTime.Format(time.RFC3339)).
		//MaxResults(nbEvents).
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
			fmt.Fprintln(w, i.Summary, " ", i.Start.DateTime)
		}
	} else {
		fmt.Printf("No upcoming events found.\n")
	}

	j, err := json.Marshal(calendarEvents)
	if err == nil {
		w.Write([]byte(j))
	}
}
