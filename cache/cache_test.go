package cache

import (
	"encoding/json"
	"golang_in_memory_key_value/models"
	"io/ioutil"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func Test_GetValueOK(t *testing.T) {
	cache := GetCache()
	cache.Set("testKey", "testValue", time.Second*time.Duration(1))
	actual, _ := cache.Get("testKey")
	expected := "testValue"
	assert.Equal(t, expected, actual, "It should get value")
}

func Test_GetValueTTL(t *testing.T) {
	cache := GetCache()
	cache.Set("testKey2", "testValue", 1)
	time.Sleep(10)
	actual, _ := cache.Get("testKey2")
	expected := "ttl is expired"
	assert.Equal(t, expected, actual, "It should get value")
}
func Test_GetValueNotFound(t *testing.T) {
	cache := GetCache()
	cache.Set("testKey", "testValue", time.Second*time.Duration(1))
	actual, _ := cache.Get("noKey")
	expected := "not found"
	assert.Equal(t, expected, actual, "It should get value")
}
func Test_Flush(t *testing.T) {
	cache := GetCache()
	cache.Set("testKey", "testValue", time.Second*time.Duration(1))
	cache.Flush()
	actual, _ := cache.Get("testKey")
	expected := "not found"
	assert.Equal(t, expected, actual, "It should get value")
}
func Test_GetBackup(t *testing.T) {
	cache := GetCache()
	cache.Set("testKey", "testValue", time.Second*time.Duration(1))
	cache.GetBackup()
	dat, _ := ioutil.ReadFile("/tmp/TIMESTAMP-data.json")
	var oldData map[string]models.Item
	json.Unmarshal(dat, &oldData)
	actual := oldData["testKey"].Value
	expected := "testValue"
	assert.Equal(t, expected, actual, "It should get value")
}
func Test_DeleteExpire(t *testing.T) {
	cache := GetCache()
	cache.Flush()
	cache.Set("testKey3", "testValue3", 1)
	time.Sleep(10)
	cache.DeleteExpired()
	actual, _ := cache.Get("testKey3")
	expected := "not found"
	assert.Equal(t, expected, actual, "It should get value")
}
