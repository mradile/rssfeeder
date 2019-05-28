package http

import (
	"net/http"
	"strconv"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"

	"github.com/mradile/rssfeeder"
	"github.com/mradile/rssfeeder/pkg/rest"
)

func (s *Server) AddEntry(c echo.Context) error {
	var req rest.AddEntryRequest
	if err := c.Bind(&req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "request body invalid")
	}

	login := getLoginFromContext(c)

	entry := &rssfeeder.FeedEntry{
		ID:         0,
		Login:      login,
		Category:   req.Category,
		URI:        req.URI,
		CreateDate: time.Now(),
	}
	if err := s.adder.AddFeedEntry(entry); err != nil {
		logrus.WithFields(logrus.Fields{
			"uri":      req.URI,
			"category": req.Category,
			"login":    login,
			"error":    err,
		}).Error("could not add entry")
		return echo.NewHTTPError(http.StatusInternalServerError, "could not add entry")
	}

	res := &rest.AddEntryResponse{
		ID:       entry.ID,
		Category: entry.Category,
	}

	return c.JSONPretty(http.StatusOK, res, "  ")
}

func (s *Server) DeleteEntry(c echo.Context) error {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "id must be of type integer")
	}

	login := getLoginFromContext(c)

	if err := s.deleter.DeleteFeedEntry(id, login); err != nil {
		logrus.WithFields(logrus.Fields{
			"id":    id,
			"login": login,
			"error": err,
		}).Error("could not delete entry")
		return echo.NewHTTPError(http.StatusInternalServerError, "could not delete entry")
	}

	return c.NoContent(http.StatusNoContent)
}
