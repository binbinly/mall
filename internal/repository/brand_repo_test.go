package repository

import (
	"context"
	"testing"
)

func TestRepo_GetBrandByID(t *testing.T) {
	got, err := r.GetBrandByID(context.Background(), 4)
	if err != nil {
		t.Errorf("GetBrandByID() error = %v", err)
		return
	}
	t.Logf("got: %v", got)
}

func TestRepo_GetBrandsByIds(t *testing.T) {
	got, err := r.GetBrandsByIds(context.Background(), []int{166, 182, 268, 3})
	if err != nil {
		t.Errorf("GetBrandsByIds() error = %v", err)
		return
	}
	t.Logf("got: %v", got)
}
