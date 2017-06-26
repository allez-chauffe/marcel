# Backend

The main purpose of this part is to serve the plugin list, configured by the back-office.

## Setup

```go
go get github.com/briandowns/openweathermap
go get github.com/GwennaelBuchet/openweathermap
go get github.com/dghubble/go-twitter/twitter
go get github.com/dghubble/oauth1
go get github.com/rs/cors
go get -u github.com/gorilla/mux
go get -u google.golang.org/api/calendar/v3
go get -u golang.org/x/oauth2/...
go get github.com/mitchellh/mapstructure
```

```shell
docker pull quay.io/goswagger/swagger
alias swagger="docker run --rm -it -v $HOME:$HOME -w $(pwd) quay.io/goswagger/swagger"

export OWM_API_KEY="your_owm_api_key"
export GOOGLE_API_KEY_FILE="your_google_api_key_file"
export MARCEL_AGENDA_ID="id_of_your_google_agenda"
export TWITTER_CONSUMER_KEY="your_twitter_consumer_key"
export TWITTER_CONSUMER_SECRET="your_twitter_consummer_secret"
export TWITTER_ACCESS_TOKEN="your_twitter_access_token"
export TWITTER_ACCESS_SECRET="your_twitter_access_secret"
export MARCEL_LOG_FILE="path_to_log_file" # defaults to $PWD/marcel.log
```

Build cross architecture :

``` shell
env GOOS=linux GOARCH=arm go build -o ./bin/MARCEL
```

In order to use Realize to manage your local builds :

```shell
go get github.com/tockins/realize
```

[Realize](https://tockins.github.io/realize/)

Then, from project(s) root, execute :

```shell
realize add
```

Once every is done, you can launch the server with :

```shell
go run main.go
```

## Credits

* OpenWeatherMap
* [OpenWeatherMap](http://briandowns.github.io/openweathermap/) Go API by briandowns
* [Twitter](https://github.com/dghubble/go-twitter) Go API by Dalton Hubble
* [MapStructure](https://github.com/mitchellh/mapstructure) from Mitchellh
* [Go-Swagger.io](https://goswagger.io)