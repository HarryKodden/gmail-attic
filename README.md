# gmail-attic

Helps to offload GMail messages from Google platform to storage of your choice.

## Contents

In te *test* folder we collect different methods of interactiving with GMail.
* **gmail-pop3**

This methods uses USERNAME + PASSWORD to connect with GMail. This only works if the GMail account is not setup with 2FA.

* **gmail-oauth**

This method uses a OAuth Authentication flow. This will als work for accounts setup with 2FA

This method requires that the application is first provided with Application Credentials to allow initiating the Google OAuth flow.
## Run samples

Either run:

```
$ go run test/gmail-pop3/main.go
```

or

```
$ go run test/gmail-oauth/main.go
```

