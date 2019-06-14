package http

import (
	"fmt"
	"github.com/gorilla/feeds"
	"github.com/labstack/echo/v4"
	"github.com/mradile/rssfeeder"
	"github.com/sirupsen/logrus"
	"net/http"
	"strconv"
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

	contentType := echo.MIMEApplicationXML
	var feedError error
	var content string
	switch feedType {
	case "json":
		content, feedError = feed.ToJSON()
		contentType = echo.MIMEApplicationJSON
	case "rss":
		content, feedError = feed.ToRss()
	default:
		content, feedError = feed.ToAtom()
	}

	if feedError != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "could fetch feed")
	}

	return c.Blob(http.StatusOK, contentType, []byte(content))
}

func (h *Handler) makeFeed(entries []*rssfeeder.FeedEntry, feedName string) *feeds.Feed {
	updated := entries[0].CreateDate
	created := entries[len(entries)-1].CreateDate

	f := &feeds.Feed{
		Title:       fmt.Sprintf("RSS Feeder - %s", feedName),
		Link:        &feeds.Link{Href: h.cfg.Hostname + "/"},
		Description: feedName,
		Author:      &feeds.Author{Name: "rssfeeder"},
		Created:     created,
		Updated:     updated,
	}

	items := make([]*feeds.Item, 0, len(entries))
	for _, entry := range entries {
		fi := &feeds.Item{
			Title: entry.URI,
			Link:  &feeds.Link{Href: entry.URI},
			//Description: desc,
			//Author:      &feeds.Author{Name: "", Email: ""},
			Created: entry.CreateDate,
			Id:      strconv.Itoa(entry.ID),
		}
		items = append(items, fi)
	}
	f.Items = items

	return f
}
