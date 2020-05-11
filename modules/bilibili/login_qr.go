package bilibili

import (
	"errors"
	"net/url"
	"reflect"
	"strings"
	"time"

	"github.com/Yesterday17/bili-archive/bilibili"
	"github.com/Yesterday17/pug/modules/base"
	"github.com/Yesterday17/pug/utils/log"

	"github.com/Yesterday17/pug/api"
	"github.com/Yesterday17/pug/utils/types"
)

type loginQrPipe struct {
	timeout int
}

var loginQrPipeBuilder = base.NewPipeBuilder(map[string]reflect.Kind{"timeout": reflect.Int}, nil, newLoginQrPipe)

func newLoginQrPipe(m map[string]interface{}) (api.Pipe, api.PipeBuildError) {
	timeout, err := types.Int(m, "timeout")
	if err != api.PipeNoError {
		timeout = 30
	}

	if timeout < 30 {
		timeout = 30
	}
	return &loginQrPipe{timeout: timeout}, api.PipeNoError
}

func (l *loginQrPipe) Validate() map[string]reflect.Kind {
	return map[string]reflect.Kind{
		"+bili_cookies": reflect.Map,
	}
}

func (l *loginQrPipe) Execute(work api.State) error {
	code := bilibili.GetLoginQRCode()
	log.Infof(code.Image)

	response := ""
	cookies := url.Values{}
	for times := 0; times < l.timeout; times++ {
		ok, ret, err := code.Check()
		response = ret

		if err != nil {
			log.Fatal(err.Error())
		}

		if ok {
			response = response[42 : len(response)-72]
			for _, value := range strings.Split(response, "&") {
				ans := strings.Split(value, "=")
				cookies.Set(ans[0], ans[1])
			}
			break
		} else {
			response = ""
			time.Sleep(3 * time.Second)
		}
	}

	if response == "" {
		return errors.New("no login") // FIXME
	}

	work.Set("bili_cookies", cookies)
	return nil
}

func (l *loginQrPipe) Clone() api.Pipe {
	return &loginQrPipe{}
}
