package check

import (
	"fmt"
	"net/http"
)

type BodyEqualChecker struct {
	Value string
}

func (c *BodyEqualChecker) Check(response *http.Response) error {
	body, err := readResponseBody(response)
	if err != nil {
		return fmt.Errorf("could not read response body: %s", err)
	}

	if exp, got := c.Value, string(body); exp != got {
		return fmt.Errorf("expected response body of `%s`, got `%s`", exp, got)
	}

	return nil
}