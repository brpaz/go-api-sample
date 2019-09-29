package context

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httputil"
	"os"
	"strconv"
	"strings"

	"github.com/DATA-DOG/godog"
	"github.com/DATA-DOG/godog/gherkin"
	"github.com/oliveagle/jsonpath"
	"github.com/xeipuuv/gojsonschema"
)

type apiContext struct {
	BaseURL         string
	JSONSchemasPath string
	client          *http.Client
	headers         map[string]string
	queryParams     map[string]string
	debug           bool
	resp            *Response
}

type Response struct {
	StatusCode  int
	Body        string
	ResponseObj *http.Response
}

type ApiResponse struct {
	Response []interface{}
	Success  bool
}

// NewAPIContext Creates a new API context for testing
func NewAPIContext(s *godog.Suite, baseURL string) *apiContext {
	ctx := &apiContext{
		BaseURL:     baseURL,
		client:      &http.Client{},
		headers:     make(map[string]string, 0),
		queryParams: make(map[string]string, 0),
		debug:       false,
	}

	ctx.registerSteps(s)

	return ctx
}

// SetDebug sets debug mode
func (ctx *apiContext) SetDebug(debug bool) {
	ctx.debug = debug
}

// SetDebug sets debug mode
func (ctx *apiContext) SetJSONSchemaPath(path string) {
	ctx.JSONSchemasPath = path
}

// Register steps into the suite
func (ctx *apiContext) registerSteps(s *godog.Suite) {
	s.BeforeScenario(ctx.Reset)

	s.Step(`^I send "([^"]*)" request to "([^"]*)"$`, ctx.ISendRequestTo)
	s.Step(`^The response code should be (\d+)$`, ctx.TheResponseCodeShouldBe)
	s.Step(`^The response should match json:$`, ctx.TheResponseShouldMatchJson)
	s.Step(`^The response should have count (\d+)$`, ctx.TheResponseShouldHaveCount)
	s.Step(`^I set header "([^"]*)" with value "([^"]*)"$`, ctx.ISetHeaderWithValue)
	s.Step(`^The json path "([^"]*)" should have value "([^"]*)"$`, ctx.TheJsonPathShouldHaveValue)
	s.Step(`^The response should match json schema "([^"]*)"$`, ctx.TheResponseShouldMatchJsonSchema)
	s.Step(`^I send "([^"]*)" request to "([^"]*)" with body:$`, ctx.ISendRequestToWithBody)
	s.Step(`^The response should be a valid json$`, ctx.TheResponseShouldBeAValidJson)
	s.Step(`^I set query param "([^"]*)" with value "([^"]*)"$`, ctx.ISetQueryParamWithValue)
	s.Step(`^I set headers to:$`, ctx.ISetHeadersTo)
	s.Step(`^I set query params to:$`, ctx.ISetQueryParamsTo)
}

// Reset Reset the internal stored data
func (ctx *apiContext) Reset(interface{}) {
	ctx.headers = make(map[string]string, 0)
	ctx.queryParams = make(map[string]string, 0)
	ctx.resp = nil
}

// ISetHeadersTo Set headers from a Data Table
func (ctx *apiContext) ISetHeadersTo(data *gherkin.DataTable) error {
	for i := 0; i < len(data.Rows); i++ {
		ctx.headers[data.Rows[i].Cells[0].Value] = data.Rows[i].Cells[1].Value
	}

	return nil
}

// IAddHeaderWithValue Step that add a new header to the current request.
func (ctx *apiContext) ISetHeaderWithValue(name string, value string) error {
	ctx.headers[name] = value
	return nil
}

// ISetQueryParamWithValue Adds a new query param to the request
func (ctx *apiContext) ISetQueryParamWithValue(name string, value string) error {
	ctx.queryParams[name] = value
	return nil
}

// ISetQueryParamsTo Set query params from a Data Table
func (ctx *apiContext) ISetQueryParamsTo(data *gherkin.DataTable) error {
	for i := 0; i < len(data.Rows); i++ {
		ctx.queryParams[data.Rows[i].Cells[0].Value] = data.Rows[i].Cells[1].Value
	}

	return nil
}

// ISendRequestTo Sends a request to the specified endpoint using the specified method.
func (ctx *apiContext) ISendRequestTo(method, uri string) error {
	reqURL := fmt.Sprintf("%s%s", ctx.BaseURL, uri)

	req, _ := http.NewRequest(method, reqURL, nil)

	// Add headers to request
	for name, value := range ctx.headers {
		req.Header.Set(name, value)
	}

	// Add query string to request
	q := req.URL.Query()
	for name, value := range ctx.queryParams {
		q.Add(name, value)
	}

	req.URL.RawQuery = q.Encode()

	if ctx.debug {
		requestDump, _ := httputil.DumpRequestOut(req, false)
		log.Printf("New Request:\n%q", requestDump)
	}

	resp, err := ctx.client.Do(req)

	if err != nil {
		return err
	}

	if ctx.debug {
		dump, _ := httputil.DumpResponse(resp, true)
		log.Printf("Received response:\n%q", dump)
	}

	body, err2 := ioutil.ReadAll(resp.Body)

	if err2 != nil {
		return err2
	}

	ctx.resp = &Response{
		StatusCode:  resp.StatusCode,
		ResponseObj: resp,
		Body:        string(body),
	}

	return nil
}

// ISendRequestToWithBody Send a request with json body. Ex: a POST request.
func (ctx *apiContext) ISendRequestToWithBody(method, uri string, requestBody *gherkin.DocString) error {

	reqURL := fmt.Sprintf("%s%s", ctx.BaseURL, uri)

	var jsonStr = []byte(requestBody.Content)
	req, err := http.NewRequest(method, reqURL, bytes.NewBuffer(jsonStr))

	for name, value := range ctx.headers {
		req.Header.Set(name, value)
	}

	if err != nil {
		return err
	}

	if ctx.debug {
		requestDump, _ := httputil.DumpRequestOut(req, false)
		log.Printf("New Request:\n%q", requestDump)
	}

	resp, err := ctx.client.Do(req)

	if err != nil {
		return err
	}

	if ctx.debug {
		dump, _ := httputil.DumpResponse(resp, true)
		log.Printf("Received response:\n%q", dump)
	}

	body, err2 := ioutil.ReadAll(resp.Body)

	if err2 != nil {
		return err2
	}

	ctx.resp = &Response{
		StatusCode:  resp.StatusCode,
		ResponseObj: resp,
		Body:        string(body),
	}

	return nil
}

// TheResponseCodeShouldBe Check if the http status code of the response matches the specified value.
func (ctx *apiContext) TheResponseCodeShouldBe(code int) error {
	if code != ctx.resp.StatusCode {
		if ctx.resp.StatusCode >= 400 {
			return fmt.Errorf("expected Response code to be: %d, but actual is: %d, Response message: %s", code, ctx.resp.StatusCode, ctx.resp.Body)
		}
		return fmt.Errorf("expected Response code to be: %d, but actual is: %d", code, ctx.resp.StatusCode)
	}
	return nil
}

// TheResponseShouldBeAValidJson checks if the response is a valid JSON.
func (ctx *apiContext) TheResponseShouldBeAValidJson() error {
	var data interface{}
	if err := json.Unmarshal([]byte(ctx.resp.Body), &data); err != nil {
		return err
	}

	return nil
}

// TheJsonPathShouldHaveValue Validates if the json object have the expected value at the specified path.
func (ctx *apiContext) TheJsonPathShouldHaveValue(path string, value string) error {

	var jsonData interface{}
	json.Unmarshal([]byte(ctx.resp.Body), &jsonData)

	res, _ := jsonpath.JsonPathLookup(jsonData, path)

	// TODO handle other variable types
	if resV, ok := res.(bool); ok == true {
		v2, _ := strconv.ParseBool(value)

		if resV != v2 {
			return fmt.Errorf("Property value does not match the expected value. Expected %s | Actual %s", value, res)
		}
	}

	if resV, ok := res.(string); ok == true {
		if resV != value {
			return fmt.Errorf("Property value does not match the expected value. Expected %s | Actual %s", value, res)
		}
	}

	return nil
}

// TheResponseShouldMatchJson Check that response matches the expected JSON.
func (ctx *apiContext) TheResponseShouldMatchJson(body *gherkin.DocString) error {
	var expected, actual []byte
	var data interface{}
	var err error
	if err = json.Unmarshal([]byte(body.Content), &data); err != nil {
		return err
	}
	if expected, err = json.Marshal(data); err != nil {
		return err
	}

	if ctx.resp.Body != string(expected) {
		return fmt.Errorf("expected json %s, does not match actual: %s", string(expected), string(actual))
	}
	return nil
}

// TheResponseShouldHaveCount Check if the response have the expected number of items.
func (ctx *apiContext) TheResponseShouldHaveCount(count int) error {

	var data ApiResponse
	var err error

	if err != nil {
		return err
	}

	if err = json.Unmarshal([]byte(ctx.resp.Body), &data); err != nil {
		return err
	}

	if len(data.Response) != count {
		return fmt.Errorf("expected %d items, but found: %d", count, len(data.Response))
	}

	return nil
}

// TheResponseShouldMatchJsonSchema Checks if the response matches the specified JSON schemctx.
func (ctx *apiContext) TheResponseShouldMatchJsonSchema(path string) error {

	path = strings.Trim(path, "/")

	schemaPath := fmt.Sprintf("%s/%s", ctx.JSONSchemasPath, path)

	if _, err := os.Stat(schemaPath); os.IsNotExist(err) {
		return fmt.Errorf("JSON schema file does not exist: %s", schemaPath)
	}

	schemaContents, err := ioutil.ReadFile(schemaPath)
	if err != nil {
		return fmt.Errorf("Cannot open json schema file: %s", err)
	}

	schemaLoader := gojsonschema.NewStringLoader(string(schemaContents))
	documentLoader := gojsonschema.NewStringLoader(ctx.resp.Body)
	result, err := gojsonschema.Validate(schemaLoader, documentLoader)

	if err != nil {
		return err
	}

	if !result.Valid() {
		fmt.Printf("The document is not valid according to the specified schema %s:", path)
		for _, desc := range result.Errors() {
			fmt.Printf("- %s\n", desc)
		}

		return errors.New("The document is not valid according to the specified schema")
	}

	return nil
}
