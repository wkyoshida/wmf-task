package mediawiki

import (
	"errors"
	"regexp"
	"strings"

	"shortdesc-api/internal/config"

	"cgt.name/pkg/go-mwclient"
)

// PageShortDesc is a struct to hold the short description of a MediaWiki page.
type PageShortDesc struct {
	PageID       string `json:"pageid"`
	Title        string `json:"title"`
	ShortDesc    string `json:"shortdescription"`
	ShortDescRaw string `json:"shortdescriptionraw"`
	Timestamp    string `json:"timestamp"`
}

var shortDescRawRegex = "(\\{\\{\\s*[sS]hort description\\|(?:1=)?)([^}|]+)([^}]*\\}\\})"

var shortDescRegex = "\\{\\{\\s*[sS]hort description\\|(.*?)\\}\\}"

var client *mwclient.Client

var userAgent = "shortdesc-api (https://github.com/wkyoshida/wmf-task) go-mwclient/v1.2.0"

// ErrClientExists is returned by Init() when there already is an existing client.
var ErrClientExists = errors.New("existing MediaWiki client")

// Client gets the MediaWiki client.
func Client() *mwclient.Client {
	return client
}

// Init creates a new MediaWiki client.
func Init(config config.MWInstanceConfig) error {
	if client != nil {
		return ErrClientExists
	}

	var err error
	client, err = mwclient.New(config.ApiUrl, userAgent)
	if err != nil {
		return err
	}

	if config.User != "" {
		err = client.Login(config.User, config.Pass)
		if err != nil {
			return err
		}
	}

	return nil
}

// GetShortDescs gets the short descriptions for MediaWiki pages.
func GetShortDescs(titles []string) ([]PageShortDesc, error) {
	titlesParam := constructTitlesParam(titles)

	pages, err := client.GetPagesByName(titlesParam)
	if err != nil {
		return nil, err
	}

	pageShortDescs := make([]PageShortDesc, len(pages))

	i := 0
	for title, page := range pages {
		shortDescRaw, err := parseShortDescRaw(page.Content)
		if err != nil {
			return nil, err
		}

		shortDesc, err := parseShortDesc(shortDescRaw)
		if err != nil {
			return nil, err
		}

		pageShortDescs[i] = PageShortDesc{
			PageID:       page.PageID,
			Title:        title,
			ShortDesc:    shortDesc,
			ShortDescRaw: shortDescRaw,
			Timestamp:    page.Timestamp,
		}

		i++
	}

	return pageShortDescs, nil
}

func constructTitlesParam(titles []string) string {
	return strings.Join(titles, "|")
}

func parseShortDescRaw(content string) (string, error) {
	r, err := regexp.Compile(shortDescRawRegex)
	if err != nil {
		return "", err
	}

	return r.FindString(content), nil
}

func parseShortDesc(shortDescRaw string) (string, error) {
	r, err := regexp.Compile(shortDescRegex)
	if err != nil {
		return "", err
	}

	// Capture the text inside the raw short-description annotation.
	matches := r.FindStringSubmatch(shortDescRaw)
	if len(matches) > 1 {
		return matches[len(matches)-1], nil
	}

	return "", nil
}
