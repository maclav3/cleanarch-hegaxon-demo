package service

import (
	"context"
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestNewService checks that NewService returns a service that has all the fields initialized.
func TestNewService(t *testing.T) {
	svc := NewService(context.Background())

	v := reflect.ValueOf(svc).Elem()
	for i := 0; i < v.NumField(); i++ {
		fld := v.Field(i)
		assert.Truef(t, fld.IsValid(), "expected a non-zero value for %s", fld.String())
	}
}
