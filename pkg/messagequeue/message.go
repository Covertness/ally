package messagequeue

import (
	"encoding/json"
	"time"
)

// Message message in the queue
type Message struct {
	Package string
	Action  string
	Params  map[string]interface{}
}

// SendMessage post message into the queue
func SendMessage(m *Message) (uint64, error) {
	data, err := json.Marshal(m)
	if err != nil {
		return 0, err
	}
	return conn.Put(data, 1, 0, time.Minute)
}

// GetMessage get message from the queue
func GetMessage(timeout time.Duration) (uint64, *Message, error) {
	id, data, err := conn.Reserve(timeout)
	if err != nil {
		return 0, nil, err
	}

	var m Message
	err = json.Unmarshal(data, &m)
	if err != nil {
		return 0, nil, err
	}
	return id, &m, nil
}

// PutbackMessage put message back into the queue
func PutbackMessage(id uint64) error {
	return conn.Release(id, 1, 0)
}

// DeleteMessage remove message from the queue
func DeleteMessage(id uint64) error {
	return conn.Delete(id)
}
