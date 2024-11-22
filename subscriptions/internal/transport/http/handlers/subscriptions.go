package handlers

import (
	"mzhn/subscriptions-service/internal/domain"
	"mzhn/subscriptions-service/internal/services/subscriptionservice"
	"mzhn/subscriptions-service/internal/transport/http/middleware"
	"mzhn/subscriptions-service/pb/authpb"
	"net/http"

	"github.com/labstack/echo/v4"
)

func SubscribeToSport(ss *subscriptionservice.Service) echo.HandlerFunc {
	return func(c echo.Context) error {
		user := c.Get(middleware.USER).(*authpb.UserInfo)

		sub := new(domain.SportSubscription)

		if err := c.Bind(sub); err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		}

		if err := c.Validate(sub); err != nil {
			return err
		}

		sub.UserId = user.Id
		result, err := ss.SubscribeToSport(sub)
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}

		return c.JSON(http.StatusOK, echo.Map{"subscription": result})
	}
}

func UnsubscribeFromSport(ss *subscriptionservice.Service) echo.HandlerFunc {
	return func(c echo.Context) error {
		user := c.Get(middleware.USER).(*authpb.UserInfo)

		sub := new(domain.SportSubscription)

		if err := c.Bind(sub); err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		}

		if err := c.Validate(sub); err != nil {
			return err
		}

		sub.UserId = user.Id
		if err := ss.UnsubscribeFromSport(sub); err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}

		return c.JSON(http.StatusOK, echo.Map{"success": true})
	}
}

func SubscribeToEvent(ss *subscriptionservice.Service) echo.HandlerFunc {
	return func(c echo.Context) error {
		user := c.Get(middleware.USER).(*authpb.UserInfo)

		sub := new(domain.EventSubscription)

		if err := c.Bind(sub); err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		}

		if err := c.Validate(sub); err != nil {
			return err
		}

		sub.UserId = user.Id
		result, err := ss.SubscribeToEvent(sub)
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}

		return c.JSON(http.StatusOK, echo.Map{"subscription": result})
	}
}

func UnsubscribeFromEvent(ss *subscriptionservice.Service) echo.HandlerFunc {
	return func(c echo.Context) error {
		user := c.Get(middleware.USER).(*authpb.UserInfo)

		sub := new(domain.EventSubscription)

		if err := c.Bind(sub); err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		}

		if err := c.Validate(sub); err != nil {
			return err
		}

		sub.UserId = user.Id
		if err := ss.UnsubscribeFromEvent(sub); err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}

		return c.JSON(http.StatusOK, echo.Map{"success": true})
	}
}
