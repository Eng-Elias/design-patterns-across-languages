package configuration_manager

import (
	"fmt"
	"sync"
)

// ConfigurationManager holds the configuration data.
// It's kept private to enforce the singleton pattern.
type ConfigurationManager struct {
	configData map[string]interface{}
	lock       sync.RWMutex // RWMutex allows multiple readers or one writer
}

var (
	instance *ConfigurationManager
	once     sync.Once // Ensures initialization happens only once
)

// loadConfig simulates loading configuration data.
// This is called internally during the first GetInstance call.
func (cm *ConfigurationManager) loadConfig() {
	fmt.Println("Loading configuration...")
	// In a real app, load from file, env vars, config service, etc.
	cm.configData = map[string]interface{}{
		"dbHost":     "db.example.go.com",
		"dbPort":     1521,
		"featureFlagZ": true,
	}
	fmt.Println("Configuration loaded.")
}

// GetInstance returns the single instance of ConfigurationManager.
// It uses sync.Once to ensure thread-safe initialization.
func GetInstance() *ConfigurationManager {
	once.Do(func() {
		instance = &ConfigurationManager{
			configData: make(map[string]interface{}),
		}
		instance.loadConfig()
	})
	return instance
}

// GetSetting retrieves a configuration setting by key.
// Uses RLock for safe concurrent reads.
func (cm *ConfigurationManager) GetSetting(key string) (interface{}, bool) {
	cm.lock.RLock()         // Acquire read lock
	defer cm.lock.RUnlock() // Release read lock when function returns
	val, ok := cm.configData[key]
	return val, ok
}

// SetSetting sets a configuration value.
// Uses Lock for safe concurrent writes.
func (cm *ConfigurationManager) SetSetting(key string, value interface{}) {
	cm.lock.Lock()         // Acquire write lock
	defer cm.lock.Unlock() // Release write lock when function returns
	cm.configData[key] = value
}

// GetAllSettings returns a copy of all configuration settings.
// Uses RLock for safe concurrent reads.
func (cm *ConfigurationManager) GetAllSettings() map[string]interface{} {
	cm.lock.RLock()
	defer cm.lock.RUnlock()

	// Return a copy to prevent external modification
	copiedSettings := make(map[string]interface{}, len(cm.configData))
	for k, v := range cm.configData {
		copiedSettings[k] = v
	}
	return copiedSettings
}
