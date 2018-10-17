### Authentication

Basic flow of authentication implemented using golang. Database used MongoDB.

### Instructions
1. Make sure to have go dependencies installed specified in vendor.json file under vendor directory.
2. Until dependencies get installed, make sure to start mongodb server. By application will try to listen on port `27017`.
3. Create `config.json` file in root directory of the application and fill up with details in following way.
```
{
  "ClientID": "YOUR_CLIENT_ID_HERE",
  "ClientSecret": "YOUR_CLIENT_SECRET_HERE",
  "SendGridAPIKey": "YOUR_SENDGRID_API_KEY_HERE",
  "MONGODB_URI": "YOUR_MONGODB_URI_HERE",
  "RedirectURI": "YOUR_REDIRECT_URI_HERE" // For e.g [http://localhost:8000/google/callback]
}
```
4. Once dependencies and server has been started, run `go run main.go` which will start server on [localhost:8000](http://localhost:8000).

### Description

Application will get connected to database named `go`. All user data will be stored under collection named `first`. Sessions will be stored under collection `session` and for password reset functionality, token and every other required data will be stored under collection `reset`.

```
HAPPY LEARNING
HAPPY CODING
```
