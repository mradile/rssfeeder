package http

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"net/http"

	"github.com/mradile/rssfeeder/pkg/rest"

	"github.com/labstack/echo/v4"
)

var feedTypes = [3]string{"atom", "rss", "json"}

func (h *Handler) ListFeeds(c echo.Context) error {
	login := getLoginFromContext(c)
	feeds, err := h.viewer.GetFeeds(login)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"error": err,
		}).Error("list feeds")

		return echo.NewHTTPError(http.StatusInternalServerError, "could not fetch feeds")
	}

	feedList := make([]*rest.Feed, 0, len(feeds))
	for _, feed := range feeds {
		f := &rest.Feed{
			Name: feed.Name,
			URIs: make([]string, 0, 3),
		}
		for _, feedType := range feedTypes {
			uri := fmt.Sprintf("%s/feeds/%s/%s/%s/%s/.rss",
				h.cfg.Hostname,
				login,
				feed.Token,
				feed.Name,
				feedType,
			)
			f.URIs = append(f.URIs, uri)
		}
		feedList = append(feedList, f)
	}
	feedRes := &rest.FeedListResponse{
		Feeds: feedList,
	}

	return c.JSONPretty(http.StatusOK, feedRes, "  ")
}
