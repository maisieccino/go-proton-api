package proton

import (
	"context"

	"github.com/go-resty/resty/v2"
)

func (c *Client) GetAllFilters(ctx context.Context) ([]Filter, error) {
	var res struct {
		Filters []Filter
	}

	if err := c.do(ctx, func(r *resty.Request) (*resty.Response, error) {
		return r.SetResult(&res).Get("/mail/v4/filters")
	}); err != nil {
		return nil, err
	}
	return res.Filters, nil
}

func (c *Client) GetFilter(ctx context.Context, id string) (Filter, error) {
	var res struct {
		Filter Filter
	}
	if err := c.do(ctx, func(r *resty.Request) (*resty.Response, error) {
		return r.SetResult(&res).Get("/mail/v4/filters/" + id)
	}); err != nil {
		return Filter{}, err
	}
	return res.Filter, nil
}

func (c *Client) AddFilter(ctx context.Context, name string, sieve string) error {
	body := struct {
		Name    string
		Sieve   string
		Version int
	}{
		Name:    name,
		Sieve:   sieve,
		Version: 2,
	}
	if err := c.do(ctx, func(r *resty.Request) (*resty.Response, error) {
		return r.SetBody(body).Post("/mail/v4/filters")
	}); err != nil {
		return err
	}
	return nil
}

func (c *Client) ClearFilters(ctx context.Context) error {
	return c.do(ctx, func(r *resty.Request) (*resty.Response, error) {
		return r.Delete("/mail/v4/filters")
	})
}

func (c *Client) UpdateFilter(ctx context.Context, id string, params FilterUpdateReq) (Filter, error) {
	var res struct {
		Filter Filter
	}
	if err := c.do(ctx, func(r *resty.Request) (*resty.Response, error) {
		return r.SetBody(params).SetResult(&res).Put("/mail/v4/filters/" + id)
	}); err != nil {
		return Filter{}, err
	}
	return res.Filter, nil
}

func (c *Client) CheckFilter(ctx context.Context, sieve string) ([]FilterIssue, error) {
	var res struct {
		Issues []FilterIssue
	}

	body := struct {
		Sieve   string
		Version int
	}{
		Sieve:   sieve,
		Version: 2,
	}

	if err := c.do(ctx, func(r *resty.Request) (*resty.Response, error) {
		return r.SetBody(body).SetResult(&res).Put("/mail/v4/filters/check")
	}); err != nil {
		return nil, err
	}
	return res.Issues, nil
}

func (c *Client) EnableFilter(ctx context.Context, id string) error {
	return c.do(ctx, func(r *resty.Request) (*resty.Response, error) {
		return r.Put("/mail/v4/filters/" + id + "/enable")
	})
}

func (c *Client) DisableFilter(ctx context.Context, id string) error {
	return c.do(ctx, func(r *resty.Request) (*resty.Response, error) {
		return r.Put("/mail/v4/filters/" + id + "/disable")
	})
}

func (c *Client) DeleteFilter(ctx context.Context, id string) error {
	return c.do(ctx, func(r *resty.Request) (*resty.Response, error) {
		return r.Delete("/mail/v4/filters/" + id)
	})
}

func (c *Client) UpdateFilterOrder(ctx context.Context, ids []string) error {
	return c.do(ctx, func(r *resty.Request) (*resty.Response, error) {
		return r.SetBody(ids).Put("/mail/v4/filters/order")
	})
}
