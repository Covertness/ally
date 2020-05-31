package scraper

import (
	"fmt"
	"github.com/Covertness/ally/pkg/config"
	"github.com/imroc/req"
)

var ins *Scraper

// Init create the FTX connection
func Init() *Scraper {
	ins = New(config.GetScraperEndpoint())
	return ins
}

// GetInstance get the Scraper instance
func GetInstance() *Scraper {
	return ins
}

// New create a new instance
func New(endpoint string) *Scraper {
	return &Scraper{
		Endpoint:          endpoint,
	}
}

func (s *Scraper) get(path string) (*req.Resp, error) {
	url := fmt.Sprintf("%s%s",
		s.Endpoint, path,
	)

	r, err := req.Get(url)
	if err != nil {
		return nil, err
	}

	if r.Response().StatusCode != 200 {
		return nil, fmt.Errorf("http code %d", r.Response().StatusCode)
	}
	return r, nil
}

// Scraper instance connect to the internal service scraper
type Scraper struct {
	Endpoint          string
}
