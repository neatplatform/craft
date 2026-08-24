package ui

type nopUI struct{}

// NewNop creates a nop user interface for testing purposes.
func NewNop() UI {
	return &nopUI{}
}

func (u *nopUI) Printf(string, ...any) {}

func (u *nopUI) GetLevel() Level {
	return None
}

func (u *nopUI) SetLevel(Level) {}

func (u *nopUI) Tracef(Style, string, ...any) {}

func (u *nopUI) Debugf(Style, string, ...any) {}

func (u *nopUI) Infof(Style, string, ...any) {}

func (u *nopUI) Warnf(Style, string, ...any) {}

func (u *nopUI) Errorf(Style, string, ...any) {}
