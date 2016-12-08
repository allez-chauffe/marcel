##Setup
```go
go get github.com/briandowns/openweathermap
go get github.com/GwennaelBuchet/openweathermap
go get github.com/rs/cors
go get -u github.com/gorilla/mux
go get -u google.golang.org/api/calendar/v3
go get -u golang.org/x/oauth2/...
```

```shell
export OWM_API_KEY="your_owm_api_key"
export GOOGLE_API_KEY="your_google_api_key"
export GOOGLE_API_SECRET="your_google_api_secret"
export MARCEL_AGENDA_ID="id_of_your_google_agenda"
```

Build cross architecture :
``` shell
env GOOS=linux GOARCH=arm go build -o ./bin/MARCEL
```

In order to use Realize to manage your local builds :
```shell
go get github.com/tockins/realize
```
(https://tockins.github.io/realize/)

Then, from project(s) root, execute :
```shell
realize add
```



##Credits
 - OpenWeatherMap
 - OpenWeatherMap Go API by briandowns (http://briandowns.github.io/openweathermap/)