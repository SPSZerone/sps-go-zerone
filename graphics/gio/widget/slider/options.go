package slider

type Option func(s *Slider)

func OptMin(value float32) Option {
	return func(s *Slider) {
		s.Min = value
	}
}

func OptMax(value float32) Option {
	return func(s *Slider) {
		s.Max = value
	}
}

func OptMinMax(min, max float32) Option {
	return func(s *Slider) {
		s.Min = min
		s.Max = max
	}
}
