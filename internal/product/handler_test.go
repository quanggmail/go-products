package product

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

type mockProductCreator struct {
	createFn func(p *Product) error
}

func (m *mockProductCreator) Create(p *Product) error {
	if m.createFn != nil {
		return m.createFn(p)
	}
	return nil
}

func TestCreateProduct_Success(t *testing.T) {
	handler := NewHandler(&mockProductCreator{
		createFn: func(p *Product) error {
			p.ID = 42
			return nil
		},
	})

	body := `{"title":"T-Shirt","description":"Cotton tee","price":29.99,"size":"M"}`
	req := httptest.NewRequest(http.MethodPost, "/products", strings.NewReader(body))
	rec := httptest.NewRecorder()

	handler.CreateProduct(rec, req)

	if rec.Code != http.StatusOK {
		t.Fatalf("status = %d, want %d; body = %s", rec.Code, http.StatusOK, rec.Body.String())
	}

	var got Product
	if err := json.NewDecoder(rec.Body).Decode(&got); err != nil {
		t.Fatalf("decode response: %v", err)
	}

	if got.ID != 42 {
		t.Errorf("id = %d, want 42", got.ID)
	}
	if got.Title != "T-Shirt" {
		t.Errorf("title = %q, want %q", got.Title, "T-Shirt")
	}
	if got.Description != "Cotton tee" {
		t.Errorf("description = %q, want %q", got.Description, "Cotton tee")
	}
	if got.Price != 29.99 {
		t.Errorf("price = %v, want 29.99", got.Price)
	}
	if got.Size != "M" {
		t.Errorf("size = %q, want %q", got.Size, "M")
	}
}

func TestCreateProduct_InvalidJSON(t *testing.T) {
	handler := NewHandler(&mockProductCreator{})

	req := httptest.NewRequest(http.MethodPost, "/products", strings.NewReader(`{invalid`))
	rec := httptest.NewRecorder()

	handler.CreateProduct(rec, req)

	if rec.Code != http.StatusBadRequest {
		t.Fatalf("status = %d, want %d; body = %s", rec.Code, http.StatusBadRequest, rec.Body.String())
	}
}

func TestCreateProduct_ServiceError(t *testing.T) {
	handler := NewHandler(&mockProductCreator{
		createFn: func(p *Product) error {
			return errors.New("database unavailable")
		},
	})

	body := `{"title":"T-Shirt","description":"Cotton tee","price":29.99,"size":"M"}`
	req := httptest.NewRequest(http.MethodPost, "/products", strings.NewReader(body))
	rec := httptest.NewRecorder()

	handler.CreateProduct(rec, req)

	if rec.Code != http.StatusInternalServerError {
		t.Fatalf("status = %d, want %d; body = %s", rec.Code, http.StatusInternalServerError, rec.Body.String())
	}

	if !strings.Contains(rec.Body.String(), "database unavailable") {
		t.Errorf("body = %q, want error message containing %q", rec.Body.String(), "database unavailable")
	}
}
