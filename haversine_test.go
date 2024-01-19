package haversine_test

import (
	"testing"

	"kkn.fi/haversine"
)

var (
	pLat, pLon    = 60.174289, 24.951477
	qLat, qLon    = 60.192706, 24.961082
	expectedKm    = 2.1156116642288487
	expectedMiles = 1.3143291425235886
)

func TestHaversineDistanceKm(t *testing.T) {
	if d := haversine.DistanceKm(pLat, pLon, qLat, qLon); d != expectedKm {
		t.Errorf("expected %v km, got %v km", expectedKm, d)
	}
}
func TestHaversineDistanceMiles(t *testing.T) {
	if d := haversine.DistanceMiles(pLat, pLon, qLat, qLon); d != expectedMiles {
		t.Errorf("expected %v miles, got %v miles", expectedMiles, d)
	}
}

func TestHaversineDistance(t *testing.T) {
	km, miles := haversine.Distance(pLat, pLon, qLat, qLon)
	if km != expectedKm {
		t.Errorf("expected %v km, got %v km", expectedKm, km)
	}
	if miles != expectedMiles {
		t.Errorf("expected %v miles, got %v miles", expectedMiles, miles)
	}
}
