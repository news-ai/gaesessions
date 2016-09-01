gaesessions
===========

Session stores for Google App Engine's datastore and memcahe.

Remove expired sessions in the datastore. you can call this function from a cron job.

Rample handler config in `app.yaml`: handlers:

```yaml
- url: /tasks/removeExpiredSessions
  script: _go_app
  login: admin
- url: /.*
  script: _go_app
```

Handler registration code:

```go
http.HandleFunc("/tasks/removeExpiredSessions", removeExpiredSessionsHandler)
```

Sample handler:

```go
func removeExpiredSessionsHandler(w http.ResponseWriter, r *http.Request) {
	c := appengine.NewContext(r)
	gaesessions.RemoveExpiredDatastoreSessions(c, "")
}
```

sample ```cron.yaml`:

```yaml
cron:
- description: expired session removal job
  url: /tasks/removeExpiredSessions
  schedule: every 1 minutes
```
