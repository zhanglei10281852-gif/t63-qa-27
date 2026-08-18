package httpapi

import (
	"net/http"
	"testing"
	"time"
)

func TestShiftMustRemainInsideSelectedServiceDay(t *testing.T) {
	h := newAPIHarness(t)
	route, _ := h.createRouteAndShift(t, "603")
	h.request(t, http.MethodPost, "/api/v1/shifts", map[string]any{"route_id": requireString(t, route, "id"), "service_date": "2026-08-18", "start_at": h.now.Add(21 * time.Hour), "end_at": h.now.Add(22 * time.Hour)}, http.StatusBadRequest)
}
