echo "Installing go dependencies..."
echo ""
echo "    - github.com/rs/cors"
go get -u github.com/rs/cors
echo "    - github.com/gorilla/mux"
go get -u github.com/gorilla/mux
echo "    - github.com/gorilla/websocket"
go get github.com/gorilla/websocket
echo "    - github.com/mitchellh/mapstructure"
go get -u github.com/mitchellh/mapstructure
echo "    - github.com/satori/go.uuid"
go get -u github.com/satori/go.uuid
echo "    - github.com/dgrijalva/jwt-go"
go get -u github.com/dgrijalva/jwt-go
echo "    - github.com/Pallinder/go-randomdata"
go get -u github.com/Pallinder/go-randomdata
echo "    - github.com/gorilla/handlers"
go get -u github.com/gorilla/handlers
echo ""
echo "Go dependencies installation finished"