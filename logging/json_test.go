package logging

import (
	"bufio"
	"bytes"
	"testing"

	"github.com/mailru/easyjson"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
)

func TestNewJSONFormatter(t *testing.T) {
	var log = logrus.New()
	var b bytes.Buffer
	output := bufio.NewWriter(&b)
	log.SetOutput(output)

	// When
	log.SetFormatter(newJSONFormatter())

	// Then
	log.Info("This is a test message")

	output.Flush()
	//fmt.Println(b.String())

	rec := LogRecord{}
	err := easyjson.Unmarshal(b.Bytes(), &rec)
	assert.NoError(t, err)
	assert.Equal(t, "This is a test message", rec.Message)
	assert.Equal(t, "INFO", rec.LogLevel)
	assert.Contains(t, rec.FuncName, "TestNewJSONFormatter")
	assert.Equal(t, "logrus", rec.Category)
}