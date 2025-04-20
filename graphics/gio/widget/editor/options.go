package editor

type Option func(*Editor)

func OptHint(value string) Option {
	return func(editor *Editor) {
		editor.Hint = value
	}
}

func OptSingleLine(value bool) Option {
	return func(editor *Editor) {
		editor.SingleLine = value
	}
}

func OptSubmit(value bool) Option {
	return func(editor *Editor) {
		editor.Submit = value
	}
}

func OptReadOnly(value bool) Option {
	return func(editor *Editor) {
		editor.ReadOnly = value
	}
}
