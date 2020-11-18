package commands

import (
	"github.com/vojapet/golang-tplink-hs100/internal/connector"
	"time"
	"encoding/json"
	"github.com/pkg/errors"
)

type CommandExecutor struct {
	Ip string
	Timeout time.Duration
}

func (self *CommandExecutor) GetIp() string {
	return self.Ip
}

func (self *CommandExecutor) GetTimeout() time.Duration {
	return self.Timeout
}

func (self *CommandExecutor)  Execute(aQuery Query) (Response, error) {
	b, err := json.Marshal(aQuery)

	if err != nil {
		return nil, err
	}

	resp, err := connector.SendCommand(self.Ip, string(b), self.Timeout)
	if err != nil {
		return nil, errors.Wrap(err, "Error on sending sysinfo command.")
	}

	response := make(Response)
	err = json.Unmarshal([]byte(resp), &response)
	if err != nil {
		return nil, err
	}

	return response, nil
}
