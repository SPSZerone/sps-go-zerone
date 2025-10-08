package glfw

func NewOptions(opts ...Option) Options {
	o := Options{}
	o.Update(opts...)
	return o
}

type (
	Option func(o *Options)

	NewWindow  func() Window
	OnInitPre  func() error
	OnGLFWInit func() error
	OnGLInit   func() error
	OnInitPost func() error
	OnStop     func()
	OnLoop     func(win Window)
)

type Options struct {
	NewWindow  NewWindow
	OnInitPre  OnInitPre
	OnGLFWInit OnGLFWInit
	OnGLInit   OnGLInit
	OnInitPost OnInitPost
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

func OptOnInitPre(value OnInitPre) Option {
	return Option(func(o *Options) {
		o.OnInitPre = value
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

func OptOnInitPost(value OnInitPost) Option {
	return Option(func(o *Options) {
		o.OnInitPost = value
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
