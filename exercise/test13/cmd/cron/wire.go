package cron

// NewServices NewServices
func NewServices() (*services, error) {
	panic(wire.Build(
		wire.Struct(new(services), "*"),
		set,
	))
}
