package utils

import (
	"bytes"
	"encoding/gob"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/utils"
)

func GetJsonTestRequestResponse(app *fiber.App, method string, url string, reqBody any) (code int, respBody map[string]any, err error) {
	bodyJson := []byte("")
	if reqBody != nil {
		bodyJson, _ = json.Marshal(reqBody)
	}
	req := httptest.NewRequest(method, url, bytes.NewReader(bodyJson))
	resp, err := app.Test(req, 10)
	code = resp.StatusCode
	// If error we're done
	if err != nil {
		return
	}
	// If no body content, we're done
	if resp.ContentLength == 0 {
		return
	}
	bodyData := make([]byte, resp.ContentLength)
	_, _ = resp.Body.Read(bodyData)
	err = json.Unmarshal(bodyData, &respBody)
	return
}

func TestEmptyHandler(_ *fiber.Ctx) error {
	return nil
}

func TestStatus200(t *testing.T, app *fiber.App, url, method string) {
	t.Helper()

	req := httptest.NewRequest(method, url, nil)

	resp, err := app.Test(req)
	utils.AssertEqual(t, nil, err, "app.Test(req)")
	utils.AssertEqual(t, 200, resp.StatusCode, "Status code")
}

func TestErrorResponse(t *testing.T, err error, resp *http.Response, expectedBodyError string) {
	t.Helper()

	utils.AssertEqual(t, nil, err, "app.Test(req)")
	utils.AssertEqual(t, 500, resp.StatusCode, "Status code")

	body, err := io.ReadAll(resp.Body)
	utils.AssertEqual(t, nil, err)
	utils.AssertEqual(t, expectedBodyError, string(body), "Response body")
}

func EncodeToBytes(p interface{}) []byte {
	buf := bytes.Buffer{}
	enc := gob.NewEncoder(&buf)
	err := enc.Encode(p)
	if err != nil {
		log.Fatal(err)
	}
	return buf.Bytes()
}
