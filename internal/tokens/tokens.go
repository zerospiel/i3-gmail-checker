package tokens

import (
	"io/ioutil"
	"net/http"
	"os"

	"context"

	"github.com/pkg/errors"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/gmail/v1"
)

const (
	oAuth2CredentialsFile = "credentials.json"
	defaultTokenFile      = "token.json"
)

type Tokenizer struct {
	config *oauth2.Config
	ctx    context.Context
}

func NewTokenizer() (*Tokenizer, error) {
	// var data []byte
	// for _, o := range opts {
	// 	data = append(data, o()...)
	// }
	// if len(data) == 0 {
	// 	return NewTokenizerWithCredentials(OAuth2CredentialsFile)
	// }

	// temp := os.TempDir() + string(os.PathSeparator) + randomString(10)
	// err := ioutil.WriteFile(temp, []byte(storage.Cred), 0400)
	// if err != nil {
	// 	return nil, err
	// }
	return NewTokenizerWithCredentials(oAuth2CredentialsFile)
}

func NewTokenizerWithCredentials(credentialsFile string) (*Tokenizer, error) {
	if _, err := os.Stat(credentialsFile); err == nil {
		jsonKey, err := ioutil.ReadFile(credentialsFile)
		if err != nil {
			return nil, err
		}

		cfg, err := google.ConfigFromJSON(jsonKey, gmail.GmailReadonlyScope)
		if err != nil {
			return nil, err
		}

		return &Tokenizer{
			config: cfg,
			ctx:    context.Background(),
		}, nil
	} else if os.IsNotExist(err) {
		return nil, errors.New("provide path to file with OAuth2.0 credentials")
	} else {
		return nil, errors.Wrap(err, "unexpected error")
	}
}

func (u *Tokenizer) GetClient() (*http.Client, error) {
	token, err := getTokenFromFile(defaultTokenFile)
	if err != nil {
		token, err = getTokenFromWeb(u.ctx, u.config)
		if err != nil {
			return nil, err
		}
		err = saveToken(defaultTokenFile, token)
		if err != nil {
			return nil, err
		}
	}
	return u.config.Client(u.ctx, token), nil
}
