package messagequeue

import (
	"fmt"

	"github.com/Covertness/ally/pkg/config"

	"github.com/beanstalkd/go-beanstalk"
)

var conn *beanstalk.Conn

// InitQueue init message queue
func InitQueue() error {
	c, err := beanstalk.Dial("tcp", fmt.Sprintf("%s:%d", config.GetMQHost(), config.GetMQPort()))
	conn = c
	return err
}
