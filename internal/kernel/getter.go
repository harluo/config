package kernel

type Getter interface {
	Get(key string) string
}
