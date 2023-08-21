package problem

type ApplicationError interface {
	Error() string
	Kind() int
}
