package ip

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetLocalMachineIp(t *testing.T) {
	ip, err := GetLocalMachineIp()
	assert.Nil(t, err)
	assert.NotEqual(t, ip, "")

	expectedIpRegex := "^(([0-9]|[1-9][0-9]|1[0-9]{2}|2[0-4][0-9]|25[0-5]).){3}([0-9]|[1-9][0-9]|1[0-9]{2}|2[0-4][0-9]|25[0-5])$"
	assert.Regexp(t, expectedIpRegex, ip)
}

func TestGetPublicMachineIp(t *testing.T) {
	ip, err := GetPublicMachineIp()
	assert.Nil(t, err)
	assert.NotEqual(t, ip, "")

	expectedIpRegex := "^(([0-9]|[1-9][0-9]|1[0-9]{2}|2[0-4][0-9]|25[0-5]).){3}([0-9]|[1-9][0-9]|1[0-9]{2}|2[0-4][0-9]|25[0-5])$"
	assert.Regexp(t, expectedIpRegex, ip)

	originalApiUri := API_URI

	API_URI = "https://foo_api.ipify.org"

	_, err = GetPublicMachineIp()
	if assert.NotNil(t, err) {
		expectedErrorMsg := "Request failed because it wasn't able to reach the ipify service"
		assert.EqualError(t, err, expectedErrorMsg)
	}

	API_URI = originalApiUri
}
