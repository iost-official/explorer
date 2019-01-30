package boot

type Config struct {
	Net map[string]Network `mapstructure:"net"`
}

type Network struct {
	EndPoint     string
	AdminAccount string
	AdminKey     string
	Initial      Initial
	ChainID      uint32
}

type Initial struct {
	GasPledge int64
	Ram       int64
	Coin      int64
}

var C Config