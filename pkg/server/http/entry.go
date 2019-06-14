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

func (h *Handler) AddEntry(c echo.Context) error {
	var req rest.AddEntryRequest
	if err := c.Bind(&req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "request body invalid")
	}

	login := getLoginFromContext(c)

	entry := &rssfeeder.FeedEntry{
		ID:         0,
		Login:      login,
		FeedName:   req.FeedName,
		URI:        req.URI,
		CreateDate: time.Now(),
	}
	if err := h.adder.AddFeedEntry(entry); err != nil {
		logrus.WithFields(logrus.Fields{
			"uri":       req.URI,
			"feed_name": req.FeedName,
			"login":     login,
			"error":     err,
		}).Error("could not add entry")
		return echo.NewHTTPError(http.StatusInternalServerError, "could not add entry")
	}

	res := &rest.AddEntryResponse{
		ID:       entry.ID,
		FeedName: entry.FeedName,
	}

	return c.JSONPretty(http.StatusOK, res, "  ")
}

func (h *Handler) DeleteEntry(c echo.Context) error {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "id must be of type integer")
	}

	login := getLoginFromContext(c)

	if err := h.deleter.DeleteFeedEntry(id, login); err != nil {
		logrus.WithFields(logrus.Fields{
			"id":    id,
			"login": login,
			"error": err,
		}).Error("could not delete entry")
		return echo.NewHTTPError(http.StatusInternalServerError, "could not delete entry")
	}

	return c.NoContent(http.StatusNoContent)
}
