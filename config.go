package conrev

type Config struct {
	cumulativeFunc func(main, join, root interface{}) interface{}
}

var config *Config

var defaultCumulativeFunc = func(main, join, root interface{}) interface{} {
	return join
}

func init() {
	config = &Config{
		cumulativeFunc: defaultCumulativeFunc,
	}
}
