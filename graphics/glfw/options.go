package glfw

func NewOptions(opts ...Option) Options {
	o := Options{}
	o.Update(opts...)
	return o
}

type (
	Option func(o *Options)

	NewWindow  func() Window
	OnGLFWInit func()
	OnGLInit   func()
	OnStop     func()
	OnLoop     func(win Window)
)

type Options struct {
	NewWindow  NewWindow
	OnGLFWInit OnGLFWInit
	OnGLInit   OnGLInit
	OnStop     OnStop
	OnLoop     OnLoop
}

func (o *Options) Update(opts ...Option) {
	for _, opt := range opts {
		opt(o)
	}
}

func OptNewWindow(value NewWindow) Option {
	return Option(func(o *Options) {
		o.NewWindow = value
	})
}

func OptOnGLFWInit(value OnGLFWInit) Option {
	return Option(func(o *Options) {
		o.OnGLFWInit = value
	})
}

func OptOnGLInit(value OnGLInit) Option {
	return Option(func(o *Options) {
		o.OnGLInit = value
	})
}

func OptOnStop(value OnStop) Option {
	return Option(func(o *Options) {
		o.OnStop = value
	})
}

func OptOnLoop(value OnLoop) Option {
	return Option(func(o *Options) {
		o.OnLoop = value
	})
}
