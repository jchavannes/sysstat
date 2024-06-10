package config

type Influx struct {
	Source string `yaml:"SOURCE"`
	Url    string `yaml:"URL"`
	Token  string `yaml:"TOKEN"`
	Org    string `yaml:"ORG"`
	Bucket string `yaml:"BUCKET"`
}
