package api

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"log/slog"
	"net/http"
	"net/url"
	"strings"
)

func Login(credentials Credentials) error {
	response, err := RawLogin(credentials.Username(), credentials.Password(), credentials.ApiVersion())
	if err != nil {
		return err
	}

	credentials.SetAccessToken(response.AccessToken)

	return nil
}

func RawLogin(username string, password string, version string) (*RawAccessTokenResponse, error) {
	data := url.Values{
		"grant_type": {"password"},
		"username":   {username},
		"password":   {password},
		"version":    {version},
	}
	resp, err := http.Post(BaseUrl+"/token", "application/x-www-form-urlencoded", strings.NewReader(data.Encode()))
	if err != nil {
		slog.Error("Couldn't perform API login request", "error", err)
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		slog.Error("Login response returned error", "status_code", resp.StatusCode, "body", body)
		return nil, errors.New("error response")
	}

	slog.Info("Successfully logged in")

	var token RawAccessTokenResponse
	err = json.NewDecoder(resp.Body).Decode(&token)
	if err != nil {
		return nil, err
	}

	return &token, nil
}

func CallWithToken(credentials Credentials, fn func(Credentials) (*http.Response, error)) (*http.Response, error) {
	if credentials.AccessToken() == nil {
		_ = Login(credentials)
	}

	resp, err := fn(credentials)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode == http.StatusUnauthorized {
		slog.Info("Performing authorization retry")

		err = Login(credentials)
		if err != nil {
			slog.Error("Authorization retry failed")
			return nil, err
		}

		resp, err = fn(credentials)
		if err != nil {
			return nil, err
		}

		if resp.StatusCode == http.StatusUnauthorized {
			slog.Error("Authorization retry failed")
			return nil, errors.New("authorization retry failed")
		}
	}

	return resp, nil
}

func RawGetSchedule(credentials Credentials, request *RawScheduleRequest) ([]RawScheduleEntry, error) {
	resp, err := CallWithToken(credentials, func(accessToken Credentials) (*http.Response, error) {
		client := http.Client{}
		body, err := json.Marshal(request)
		if err != nil {
			return nil, err
		}

		req, err := http.NewRequest("POST", Url+"/schedule/get", bytes.NewBuffer(body))
		if err != nil {
			return nil, err
		}
		req.Header.Set("Authorization", "Bearer "+*accessToken.AccessToken())
		req.Header.Set("Content-Type", "application/json")

		resp, err := client.Do(req)
		if err != nil {
			return nil, err
		}

		slog.Info("Successfully fetched schedule")

		return resp, err
	})
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		slog.Error("Couldn't get schedule", "status_code", resp.StatusCode, "body", body)
		return nil, errors.New("error response")
	}

	var schedule []RawScheduleEntry
	err = json.NewDecoder(resp.Body).Decode(&schedule)
	if err != nil {
		return nil, err
	}

	return schedule, nil
}
