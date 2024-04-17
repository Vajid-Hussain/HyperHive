package middlewire_auth_svc

import (
	"context"
	"fmt"
	"net/http"
	"strings"

	"github.com/Vajid-Hussain/HiperHive/api-gateway/pkg/auth-svc/pb"
	"github.com/labstack/echo/v4"
)

type Middlewire struct {
	Clind pb.AuthServiceClient
}

func NewAuthMiddlewire(clind pb.AuthServiceClient) *Middlewire {
	return &Middlewire{Clind: clind}
}

func (m *Middlewire) UserAuthMiddlewire(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		accessToken := c.Request().Header.Get("AccessToken")
		refreshToken := c.Request().Header.Get("RefreshToken")
		// fmt.Println("---------", accessToken, refreshToken)
		fmt.Println("user middlwire called")

		if accessToken == "" || len(accessToken) < 20 {
			return c.JSON(http.StatusUnauthorized, map[string]string{"error": "there is no access token"})
		}

		result, err := m.Clind.ValidateUserToken(context.Background(), &pb.ValidateTokenRequest{
			AccessToken:  accessToken,
			RefreshToken: refreshToken,
		})
		if err != nil {
			return c.JSON(http.StatusUnauthorized, map[string]string{"error": strings.TrimPrefix(err.Error(), "rpc error: code = Unknown desc =  ")})
		}

		// Set the "userID" in the context for downstream handlers to access
		c.Set("userID", result.UserID)
		// fmt.Println("----", result.UserID)
		// Call the next handler in the chain
		return next(c)
	}
}

func (m *Middlewire) AdminAuthMiddlewire(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		token := c.Request().Header.Get("Token")
		// fmt.Println("==========")
		_, err := m.Clind.ValidateAdminToken(context.Background(), &pb.ValidateAdminTokenRequest{
			Token: token,
		})

		if err != nil {
			return c.JSON(http.StatusUnauthorized, map[string]string{"error": strings.TrimPrefix(err.Error(), "rpc error: code = Unknown desc =  ")})
		}

		return next(c)
	}
}
