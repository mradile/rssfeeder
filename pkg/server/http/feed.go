package http

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"net/http"

	"github.com/mradile/rssfeeder/pkg/rest"

	"github.com/labstack/echo/v4"
)

var feedTypes = [3]string{"atom", "rss", "json"}

func (s *Server) ListFeeds(c echo.Context) error {
	login := getLoginFromContext(c)
	categories, err := s.viewer.GetCategories(login)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"error": err,
		}).Error("list feeds")

		return echo.NewHTTPError(http.StatusInternalServerError, "could not fetch feeds")
	}

	feedList := make([]*rest.Feed, 0, len(categories))
	for _, cat := range categories {
		f := &rest.Feed{
			Name: cat,
			URIs: make([]string, 0, 3),
		}
		for _, feedType := range feedTypes {
			uri := fmt.Sprintf("%s/feeds/%s/%s/%s/.rss",
				s.cfg.Hostname,
				login,
				cat,
				feedType,
			)
			f.URIs = append(f.URIs, uri)
		}
		feedList = append(feedList, f)
	}
	feeds := &rest.FeedListResponse{
		Feeds: feedList,
	}

	return c.JSONPretty(http.StatusOK, feeds, "  ")
}
