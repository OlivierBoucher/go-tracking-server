package utilities

import (
  "net"
  "net/http"
  "encoding/json"
  "io/ioutil"
)

//GetIP retrieves the IP address from an http.Request
func GetIP(r *http.Request) string {
    if ipProxy := r.Header.Get("X-FORWARDED-FOR"); len(ipProxy) > 0 {
        return ipProxy
    }
    ip, _, _ := net.SplitHostPort(r.RemoteAddr)
    return ip
}
//SrvConfiguration represents a container with the server's configuration
type SrvConfiguration struct {
  AuthDbConnectionString string `json:"authDb"`
  QueueConnectionUrl string `json:"queueUrl"`
}
//LoadJSONConfig returns a SrvConfiguration struct from a json file
func LoadJSONConfig() (*SrvConfiguration, error) {
  file, err := ioutil.ReadFile("./config.json")
  if err != nil {
    return nil, err
  }

  var config SrvConfiguration
  json.Unmarshal(file, &config)

  return &config, nil
}
