package config

import (
    log "github.com/siruspen/logrus"
    "github.com/spf13/viper"
)

//Config is used to define the config struct which will be used through app.
type Config struct {
    config *viper.Viper
}

//NewConfig is used to instantiate a config
func NewConfig() *Config {
    c := new(Config)
    c.config = readConfig()
    return c
}

// Get is used to retrieve the current config of project.
func (c *Config) Get() *viper.Viper {
    if c.config == nil {
        log.Fatal("config not initialized")
    }
    return c.config
}
func readConfig() *viper.Viper {
    log.Info("Reading environment variables.")
    v := viper.New()
    v.AutomaticEnv()
    env := v.GetString("ENVIRONMENT")
    if env == "" {
        env = "local"
    }
    log.Infof("Environment: %s", env)
    v.SetConfigName(env)
    v.SetConfigType("yaml")
    v.AddConfigPath("config")
    err := v.ReadInConfig()
    if err != nil {
        log.Fatalf("error reading config file or env variable '%s' ", err.Error())
    }
    return v
}
