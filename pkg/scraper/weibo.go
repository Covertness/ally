package scraper

import "fmt"

// WeiBoSearch search keyword on WeiBo
func(s *Scraper) WeiBoSearch(keyword string) ([]*WeiBoSearchItem, error) {
	path := fmt.Sprintf("/weibo/search/%s", keyword)
	r, err := s.get(path)
	if err != nil {
		return nil, err
	}

	var resp []*WeiBoSearchItem
	err = r.ToJSON(&resp)

	return resp, err
}

// WeiBoSearchItem the item of search result
type WeiBoSearchItem struct {
	ID string
	Name string
}

// WeiBoTimeline get news of a specified user on WeiBo
func(s *Scraper) WeiBoTimeline(userID string) ([]*WeiBoTimelineItem, error) {
	path := fmt.Sprintf("/weibo/%s", userID)
	r, err := s.get(path)
	if err != nil {
		return nil, err
	}

	var resp []*WeiBoTimelineItem
	err = r.ToJSON(&resp)

	return resp, err
}

// WeiBoTimelineItem the item of timeline
type WeiBoTimelineItem struct {
	Name string
	Text string
	Timestamp string
	ID string
	Link string
}
