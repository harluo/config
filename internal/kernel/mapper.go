package kernel

type Mapper interface {
	Get(key string) (value string)
}
