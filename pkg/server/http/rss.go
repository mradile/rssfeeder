package http

import (
	"fmt"
	"github.com/gorilla/feeds"
	"github.com/labstack/echo/v4"
	"github.com/mradile/rssfeeder"
	"github.com/sirupsen/logrus"
	"net/http"
)

func (h *Handler) RSSFeed(c echo.Context) error {
	login := c.Param("login")
	feedName := c.Param("feed")
	entries, err := h.viewer.GetFeed(feedName, login)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"error": err,
		}).Error("rss feed")

		return echo.NewHTTPError(http.StatusInternalServerError, "could fetch feed")
	}

	feedType := c.Param("type")

	if feedType == "json" {
		return c.JSONPretty(http.StatusOK, entries, " ")
	}

	feed := h.makeFeed(entries, feedName)

	var contentType string
	var feedError error
	var content string
	switch feedType {
	case "json":
		content, feedError = feed.ToJSON()
		contentType = echo.MIMEApplicationJSON
	case "rss":
		content, feedError = feed.ToRss()
		contentType = "application/rss+xml"
	default:
		content, feedError = feed.ToAtom()
		contentType = "application/atom+xml"
	}

	if feedError != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "could fetch feed")
	}

	return c.Blob(http.StatusOK, contentType, []byte(content))
}

func (h *Handler) makeFeed(entries []*rssfeeder.FeedEntry, feedName string) *feeds.Feed {
	f := &feeds.Feed{
		Title:       fmt.Sprintf("RSS Feeder - %s", feedName),
		Link:        &feeds.Link{Href: h.cfg.Hostname + "/"},
		Description: feedName,
		Author:      &feeds.Author{Name: "rssfeeder"},
	}

	if len(entries) < 1 {
		return f
	}

	f.Updated = entries[0].CreateDate
	f.Created = entries[len(entries)-1].CreateDate

	items := make([]*feeds.Item, 0, len(entries))
	for _, entry := range entries {
		fi := &feeds.Item{
			Title: entry.URI,
			Link:  &feeds.Link{Href: entry.URI},
			//Description: desc,
			//Author:      &feeds.Author{Name: "", Email: ""},
			Created: entry.CreateDate,
			Id:      fmt.Sprintf("%s/%d", h.cfg.Hostname, entry.ID),
		}
		items = append(items, fi)
	}
	f.Items = items

	return f
}
