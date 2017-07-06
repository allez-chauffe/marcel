package commons

import (
	"math/rand"
	"strconv"
	"time"
	"net/http"
	"log"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func GetUID() string {
	s := 6
	return randomString(s)
}

func randomString(l int) string {
	r := strconv.Itoa(rand.Intn(10000))

	bytes := make([]byte, l)
	for i := 0; i < l; i++ {
		bytes[i] = byte(randInt(64, 90))
	}
	return string(bytes) + r
}

func randInt(min int, max int) int {
	return min + rand.Intn(max-min)
}


func WriteResponseWithError(w http.ResponseWriter, errorCode int) {
	w.WriteHeader(errorCode)
}

func Check(e error) {
	if e != nil {
		log.Fatal(e)
		panic(e)
	}
}