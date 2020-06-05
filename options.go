package conrev

type Option func(c *Config)

func WithCumulativeFunc(cumulativeFn func(main, join, root interface{}) interface{}) Option {
	return func(c *Config) {
		c.cumulativeFunc = cumulativeFn
	}
}
