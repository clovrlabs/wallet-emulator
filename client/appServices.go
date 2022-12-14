package client

import "errors"

type AppServices struct{}

func (services AppServices) BackupProviderName() string {
	return "gdrive"
}

func (services AppServices) BackupProviderSignIn() (string, error) {
	return "", errors.New("no token provided")
}

func (services AppServices) Notify(notificationEvent []byte) {

}
