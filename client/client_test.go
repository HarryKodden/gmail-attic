package client

import (
	"context"
	"strings"
	"testing"
	"time"
)

func TestNew(t *testing.T) {
	for _, test := range []struct {
		tokenFile, credentialsFile string
		timeout                    time.Duration
		wantError                  string
	}{
		{
			// No tokenfile, no creds file
			timeout:   time.Second,
			wantError: "token and credentials files must be given",
		},
		{
			// Negative timeout
			tokenFile:       "a",
			credentialsFile: "a",
			timeout:         time.Second * -1,
			wantError:       "timeout may not be negative",
		},
		{
			// Credentials file must exist
			tokenFile:       "a",
			credentialsFile: "/no/where/to/be/found",
			wantError:       "failed to read credentials file",
		},
	} {
		opts := &Opts{
			TokenFile:       test.tokenFile,
			CredentialsFile: test.credentialsFile,
			Timeout:         test.timeout,
		}
		_, err := New(context.TODO(), opts)
		switch {
		case test.wantError != "" && err == nil:
			t.Errorf("New(%v) = _,nil, want error with %q", opts, test.wantError)
		case test.wantError == "" && err != nil:
			t.Errorf("New(%v) = _,%v, want nil error", opts, err)
		case test.wantError == "" && err != nil && !strings.Contains(err.Error(), test.wantError):
			t.Errorf("New(%v) = _,%v, want error with %q", opts, err, test.wantError)
		}
	}
}

func TestSaveToken(t *testing.T) {
	name := "/no/where/to/be/found"
	err := saveToken(name, nil)
	if err == nil {
		t.Errorf("saveToken(%q, _) = nil, want error", name)
	}
	errStr := "cannot cache oauth token"
	if !strings.Contains(err.Error(), errStr) {
		t.Errorf("saveToken(%q, _) = %v, want error with %q", name, err, errStr)
	}
}
