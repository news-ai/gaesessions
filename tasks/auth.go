package tasks

import (
	"net/http"

	"google.golang.org/appengine"
	"google.golang.org/appengine/log"

	"github.com/news-ai/gaesessions"
)

func RemoveExpiredSessionsHandler(w http.ResponseWriter, r *http.Request) {
	c := appengine.NewContext(r)
	err := gaesessions.RemoveExpiredDatastoreSessions(c, "")
	if err != nil {
		log.Errorf(c, "%v", err)
	}
}
