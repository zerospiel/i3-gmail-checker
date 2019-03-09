package fetcher

import (
	"github.com/zerospiel/i3-gmail-checker/internal/tokens"
	"google.golang.org/api/gmail/v1"
)

const (
	user        = "me"
	unreadLabel = "UNREAD"
)

type GmailService struct {
	s *gmail.Service
}

// GenerateAuthURL return Google OAuth2 URL to auth the app
func NewGmailService() (*GmailService, error) {
	u, err := tokens.NewTokenizer()
	if err != nil {
		return nil, err
	}

	gmailClient, err := u.GetClient()
	if err != nil {
		return nil, err
	}

	service, err := gmail.New(gmailClient)
	if err != nil {
		return nil, err
	}

	return &GmailService{
		s: service,
	}, nil
}

func (gs *GmailService) FetchUnread() (int64, error) {
	resp, err := gs.s.Users.Labels.Get(user, unreadLabel).Do()
	if err != nil {
		return 0, err
	}
	return resp.MessagesTotal, nil
}
