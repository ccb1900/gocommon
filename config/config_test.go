package config

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestDefaultLog(t *testing.T) {
	Init("test")
	s, _ := json.Marshal(GetLog())
	fmt.Println(string(s))
}
