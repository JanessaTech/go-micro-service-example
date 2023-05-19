package divekoanf

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/knadh/koanf"
	"github.com/knadh/koanf/parsers/json"
	"github.com/knadh/koanf/parsers/yaml"
	"github.com/knadh/koanf/providers/confmap"
	"github.com/knadh/koanf/providers/env"
	"github.com/knadh/koanf/providers/file"
	"github.com/knadh/koanf/providers/rawbytes"
)

// see: https://pkg.go.dev/github.com/knadh/koanf

func loadFromJsonFile() (*Config, error) {
	configPath := "./dive-koanf/properties.json"
	// Global koanf instance. Use . as the key path delimiter. This can be / or anything.
	k := koanf.New(".")
	if err := k.Load(file.Provider(configPath), json.Parser()); err != nil {
		fmt.Println("Failed to load configration from json file ", configPath)
		return nil, err
	}
	fmt.Println("[From json]server.port = ", k.String("server.port"))

	var config Config
	if err := k.UnmarshalWithConf("", &config, koanf.UnmarshalConf{Tag: "json", FlatPaths: false}); err != nil {
		fmt.Printf("failed to unmarshal with conf. err: %v", err)
	}

	return &config, nil
}

func loadFromYmlFile() (*Config, error) {
	configPath := "./dive-koanf/properties.yml"
	// Global koanf instance. Use . as the key path delimiter. This can be / or anything.
	k := koanf.New(".")
	if err := k.Load(file.Provider(configPath), yaml.Parser()); err != nil {
		fmt.Println("Failed to load configration from json file ", configPath)
		return nil, err
	}
	fmt.Println("[From yml]server.port = ", k.String("server.port"))

	var config Config
	return &config, nil
}

func loadFromMemory() (*Config, error) {
	// load from default config
	k := koanf.New(".")
	err := k.Load(confmap.Provider(defaultConfig, "."), nil)
	if err != nil {
		log.Printf("failed to load default config. err: %v", err)
		return nil, err
	}
	fmt.Println("[From memory]server.port = ", k.String("server.port"))
	var config Config
	if err := k.UnmarshalWithConf("", &config, koanf.UnmarshalConf{Tag: "json", FlatPaths: false}); err != nil {
		fmt.Printf("failed to unmarshal with conf. err: %v", err)
	}

	return &config, nil
}

// MYDEMO_SERVER_PORT=8080
// It works in debug mode
func loadFromEnv() (*Config, error) {
	// Global koanf instance. Use . as the key path delimiter. This can be / or anything.
	v := os.Getenv("MYDEMO_SERVER_PORT")
	fmt.Println("MYDEMO_SERVER_PORT=", v)
	var k = koanf.New(".")
	k.Load(env.Provider("MYDEMO_", ".", func(s string) string {
		return strings.Replace(strings.ToLower(
			strings.TrimPrefix(s, "MYDEMO_")), "_", ".", -1)
	}), nil)
	fmt.Println("[From env]server.port = ", k.String("server.port"))
	var config Config
	if err := k.UnmarshalWithConf("", &config, koanf.UnmarshalConf{Tag: "json", FlatPaths: false}); err != nil {
		fmt.Printf("failed to unmarshal with conf. err: %v", err)
	}

	return &config, nil
}
func loadFromBytes() (*Config, error) {
	// Global koanf instance. Use . as the key path delimiter. This can be / or anything.
	var k = koanf.New(".")
	b := []byte(`{"server":{"port":8080,"readTimeout":"5s"},"logging":{"level":-1},"db":{"dataSourceName":"test","migrate":{"enable":false}},"cache":{"enabled":false,"redis":{"cluster":false}}}`)
	k.Load(rawbytes.Provider(b), json.Parser())
	fmt.Println("[From bytes]server.port = ", k.String("server.port"))

	var config Config
	if err := k.UnmarshalWithConf("", &config, koanf.UnmarshalConf{Tag: "json", FlatPaths: false}); err != nil {
		fmt.Printf("failed to unmarshal with conf. err: %v", err)
	}

	return &config, nil

}

func Main() {

	loadFromJsonFile()
	loadFromYmlFile()
	loadFromMemory()
	loadFromEnv()
	loadFromBytes()
}
