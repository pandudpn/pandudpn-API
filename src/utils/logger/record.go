package logger

import (
	"bytes"
	"context"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"
	
	"pandudpn/api/src/utils/config"
	
	"github.com/google/uuid"
)

// Initialize for init first context from http rest api
func Initialize(req *http.Request) (*http.Request, DataLogger) {
	var (
		timezone = config.Timezone()
		start    = time.Now().In(timezone)
		lock     = new(Locker)
		dl       DataLogger
	)

	dl.RequestId = uuid.New().String()
	dl.Type = rest
	dl.Service = getServiceName()
	dl.Host = req.Host
	dl.Endpoint = req.URL.Path
	dl.RequestMethod = req.Method
	dl.TimeStart = start
	dl.RequestHeader = dumpHeaderFromRequest(req)
	dl.RequestBody = dumpBodyFromRequest(req)

	ctx := context.WithValue(req.Context(), logKey, lock)

	return req.WithContext(ctx), dl
}

// Store for storing data context to third parties
func (th ThirdParty) Store(ctx context.Context) {
	var (
		data []ThirdParty
		val  Values
	)

	val, ok := extract(ctx)
	if !ok {
		return
	}

	tmp, ok := val.LoadAndDelete(_ThirdParties)
	if ok {
		data = tmp.([]ThirdParty)
	}

	data = append(data, th)

	val.Set(_ThirdParties, data)
}

// Response is record data response to context
func Response(ctx context.Context, status int, res interface{}, err error) {
	value, ok := extract(ctx)
	if !ok {
		Log.Error(ctx, "error extract")
		return
	}

	if err != nil {
		value.Set(_ErrorMessage, err.Error())
	}
	// check total string for response
	if res != nil {
		buf, err := json.Marshal(res)
		if err != nil {
			Log.Error(ctx, err)
			return
		}

		if len(buf) > 1000 {
			value.Set(_Response, success)
		} else {
			value.Set(_Response, res)
		}
	}
	value.Set(_StatusCode, status)
}

// dumpHeaderFromRequest for getting all request from header http_rest_api
func dumpHeaderFromRequest(req *http.Request) map[string]interface{} {
	var reqHeader = make(map[string]interface{})

	for key, value := range req.Header {
		reqHeader[key] = strings.Join(value, ",")
	}

	return reqHeader
}

// dumpBodyFromRequest for getting all request from payload body http_rest_api
func dumpBodyFromRequest(req *http.Request) map[string]interface{} {
	var reqBody = make(map[string]interface{})
	// extract all payload
	buf, err := ioutil.ReadAll(req.Body)
	if err != nil {
		return reqBody
	}
	// put again body to payload request
	req.Body = ioutil.NopCloser(bytes.NewBuffer(buf))

	// check total len of request body
	// this is for handle when payload has sending image (binary)
	if len(buf) > 1000 {
		reqBody = map[string]interface{}{"body": string(buf[:1000])}
	} else {
		err = json.Unmarshal(buf, &reqBody)
		if err != nil {
			return reqBody
		}
	}

	return reqBody
}

// dumpBodyFromGrpc for getting all request from payload body grpc server
func dumpBodyFromGrpc(i interface{}) map[string]interface{} {
	var reqBody = make(map[string]interface{})
	// extract all payload
	buf, err := json.Marshal(i)
	if err != nil {
		return reqBody
	}

	// set data to map
	err = json.Unmarshal(buf, &reqBody)
	if err != nil {
		return reqBody
	}

	return reqBody
}

func getServiceName() string {
	return filepath.Base(os.Args[0])
}
