package cache

import (
	"encoding/json"
	"golang_in_memory_key_value/constants"
	"golang_in_memory_key_value/models"
	"io/ioutil"
	"time"
)

type Cache struct {
	Items            map[string]models.Item
	CleaningInterval time.Duration
	BackupInterval   time.Duration
}

var instance *Cache

func GetCache() *Cache {
	items := make(map[string]models.Item)
	if instance == nil {
		instance = &Cache{
			Items:            items,
			CleaningInterval: constants.CLEANING_INTERVAL,
			BackupInterval:   constants.BACKUP_INTERVAL,
		}
		instance.ReadBackup()
		go instance.BatchProcessExecuter()
	}
	return instance
}

func (c *Cache) BatchProcessExecuter() {
	cleaningIntervalticker := time.NewTicker(c.CleaningInterval)
	backupIntervalticker := time.NewTicker(c.BackupInterval)

	for {
		select {
		case <-cleaningIntervalticker.C:
			c.DeleteExpired()
		case <-backupIntervalticker.C:
			c.GetBackup()
		}
	}
}

func (c *Cache) DeleteExpired() {
	now := time.Now().UnixNano()
	for key, value := range c.Items {
		if value.TTL > 0 && now > value.TTL {
			delete(c.Items, key)
		}
	}
}

func (c *Cache) GetBackup() {
	file, _ := json.MarshalIndent(c.Items, "", " ")
	ioutil.WriteFile("/tmp/TIMESTAMP-data.json", file, 0644)
}

func (c *Cache) ReadBackup() {
	dat, _ := ioutil.ReadFile("/tmp/TIMESTAMP-data.json")
	var oldData map[string]models.Item
	json.Unmarshal(dat, &oldData)

	c.Items = oldData
}

func (c *Cache) Set(key string, value string, ttl time.Duration) {
	c.Items[key] = models.Item{
		Value: value,
		TTL:   time.Now().Add(ttl).UnixNano(),
	}
}

func (c *Cache) Get(k string) (string, bool) {
	item, found := c.Items[k]
	if !found {
		return "not found", false
	}
	if time.Now().UnixNano() > item.TTL {
		return "ttl is expired", false
	}
	return item.Value, true
}

func (c *Cache) Flush() {
	c.Items = map[string]models.Item{}
}
