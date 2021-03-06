package reconcileutils

// Code generated by github.com/launchdarkly/go-options.  DO NOT EDIT.

type ApplyGetActionOptionFunc func(c *getActionOptions) error

func (f ApplyGetActionOptionFunc) apply(c *getActionOptions) error {
	return f(c)
}

func newGetActionOptions(options ...GetActionOption) (getActionOptions, error) {
	var c getActionOptions
	err := applyGetActionOptionsOptions(&c, options...)
	return c, err
}

func applyGetActionOptionsOptions(c *getActionOptions, options ...GetActionOption) error {
	for _, o := range options {
		if err := o.apply(c); err != nil {
			return err
		}
	}
	return nil
}

type GetActionOption interface {
	apply(*getActionOptions) error
}
