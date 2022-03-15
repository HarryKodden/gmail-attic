/*

Client test.

Steps to make this work:
- Create a Google project and enable the gmail API, see
  https://developers.google.com/workspace/guides/create-project
- Create access credentials, see
  https://developers.google.com/workspace/guides/create-credentials
  - Make sure you do this for the gmail API. Enable all `gmail-api` scopes (also the restricted ones).
  - Create an OAuth Client ID. Choose for the application type: `desktop-app`. You may need to
    configure an OAuth consent screen, if so, use type "internal" (means: valid for your org).
  - Download your credentials in JSON format. The file is something like
    `client_secret_$ID.apps.googleusercontent.com.json`. Store this in a safe place. I chose
	`~/.gmail-attic/$ORG-credentials.json` (and I applied `chmod 700 ~/.gmail-attic` to lock down
	that directory).

Once all preparations are in place, try to create a client. Supply 2 arguments: the token file and
the credentials file. At the first run the token file won't exist and you will be sent to the OAuth
screen to allow access - and then the token file will be created.

Next runs will see a token- and credentials file and won't require human intervention.

--- Work in progress ---
The above should work, but I think that the OAuth screen needs more info because I am using
read/write/privileged access to the gmail API:

```
There was an error while loading /apis/credentials/consent/edit;newAppInternalUser=true?authuser=1&project=kubat-nl-gmail-api&pli=1&rapt=AEjHL4NMR1qNeaE9SyXTf9JHNuE-YX8NklZwCKtSZakazf7NUvwGCyeAXwi4XjJsVOx0fJLsWYyqOlc5moA0wex4NjW0yq0UnA

You are missing the following required permissions:
Project
oauthconfig.testusers.get
oauthconfig.verification.get
resourcemanager.projects.get
serviceusage.services.list
And you already have the following required permissions:
Project
clientauthconfig.brands.get
```
`
*/
package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/HarryKodden/gmail-attic/client"
)

func main() {
	// Require the token file and the credentials file as args
	if len(os.Args) != 3 {
		log.Fatalln("this requires 2 commandline arguments: the token file and the credentials file")
	}

	_, err := client.New(context.Background(), &client.Opts{
		TokenFile:       os.Args[1],
		CredentialsFile: os.Args[2],
		Timeout:         time.Second * 30,
	})
	if err != nil {
		log.Fatalln("failed to create client:", err)
	}
	fmt.Println("client successfully created")
}
