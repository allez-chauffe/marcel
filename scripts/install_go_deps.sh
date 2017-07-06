echo "Installing go dependencies..."
echo "    - github.com/briandowns/openweathermap"
go get github.com/briandowns/openweathermap
echo "    - github.com/GwennaelBuchet/openweathermap"
go get github.com/GwennaelBuchet/openweathermap
echo "    - github.com/dghubble/go-twitter/twitter"
go get github.com/dghubble/go-twitter/twitter
echo "    - github.com/dghubble/oauth1"
go get github.com/dghubble/oauth1
echo "    - github.com/rs/cors"
go get github.com/rs/cors
echo "    - github.com/gorilla/mux"
go get -u github.com/gorilla/mux
echo "    - google.golang.org/api/calendar/v3"
go get -u google.golang.org/api/calendar/v3
echo "    - golang.org/x/oauth2/..."
go get -u golang.org/x/oauth2/...
echo "    - github.com/mitchellh/mapstructure"
go get github.com/mitchellh/mapstructure
echo "Go dependencies installation finished"