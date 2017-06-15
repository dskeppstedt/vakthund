package vakthund

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha1"
	"encoding/hex"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

var key = []byte(os.Getenv("SCIKEY"))

//Source https://stackoverflow.com/questions/39426096/go-lang-generate-hmac/39426143
func verify_hmac(msg []byte, message_MAC []byte) bool {
	mac := hmac.New(sha1.New, key)
	mac.Write(msg)
	expected_mac := "sha1=" + hex.EncodeToString(mac.Sum(nil))
	log.Println("computed hmac:", expected_mac)

	return hmac.Equal(message_MAC, []byte(expected_mac))
}

func HmacHandler(next http.Handler) http.Handler {
	if devMode {
		log.Println("In dev mode, skipping hmac check")
		fn := func (w http.ResponseWriter, r *http.Request) {
			next.ServeHTTP(w,r)
		}
		return http.HandlerFunc(fn)
	}



	fn := func(w http.ResponseWriter, req *http.Request) {
		hmac_signature := req.Header.Get("X-Hub-Signature")
		log.Println("HMAC:", hmac_signature)

		body, err := ioutil.ReadAll(req.Body)
		if err != nil {
			panic(err)
		}

		// Restore the io.ReadCloser to its original state
		//src:https://medium.com/@xoen/golang-read-from-an-io-readwriter-without-loosing-its-content-2c6911805361
		req.Body = ioutil.NopCloser(bytes.NewBuffer(body))

		//verify the hmac to make sure that the request came from
		//github.com TODO: move this to a middleware
		if !verify_hmac(body, []byte(hmac_signature)) {
			w.WriteHeader(401)
			w.Write([]byte("Wrong hmac"))
			return
		}

		//if we made it this far, the hmac is ok
		next.ServeHTTP(w, req)

	}

	return http.HandlerFunc(fn)
}
