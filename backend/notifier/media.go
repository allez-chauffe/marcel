package notifier

type clients map[*Client]bool

//Media represent a Media goroutine, waiting for new clients and notify them when nedeed.
type Media struct {
	mediaID int
}

func newMedia(mediaID int) *Media {
	return &Media{mediaID}
}

//Close clean up channels and close all connected clients
func (m *Media) Close() {

}
