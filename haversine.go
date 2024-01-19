// https://en.wikipedia.org/wiki/Haversine_formula
package haversine

import "math"

const (
	// earthRadiusKm is the radius of the earth in kilometers.
	earthRadiusKm = 6371.0
	// earthRadiusMiles is the radius of earth in miles.
	earthRadiusMiles = 3958.0
)

// degreesToRadians converts from degrees to radians.
func degreesToRadians(d float64) float64 {
	return d * math.Pi / 180
}

// distance calculates the shortest path between two coordinates on the surface
// of the earth.
func distance(pLat, pLon, qLat, qLon float64) float64 {
	lat1 := degreesToRadians(pLat)
	lon1 := degreesToRadians(pLon)
	lat2 := degreesToRadians(qLat)
	lon2 := degreesToRadians(qLon)

	diffLatitude := lat2 - lat1
	diffLongitude := lon2 - lon1

	a := math.Pow(math.Sin(diffLatitude/2), 2) + math.Cos(lat1)*math.Cos(lat2)*math.Pow(math.Sin(diffLongitude/2), 2)
	c := 2 * math.Atan2(math.Sqrt(a), math.Sqrt(1-a))

	return c
}

// convertToKm converts given distance to kilometers.
func convertToKm(c float64) float64 {
	km := c * earthRadiusKm
	return km
}

// convertToMiles converts given distance to miles.
func convertToMiles(c float64) float64 {
	miles := c * earthRadiusMiles
	return miles
}

// Distance calculates the shortest path between two coordinates on the surface
// of the earth. Returns the result in kilometers and miles.
func Distance(pLat, pLon, qLat, qLon float64) (km, miles float64) {
	c := distance(pLat, pLon, qLat, qLon)
	km = convertToKm(c)
	miles = convertToMiles(c)
	return km, miles
}

// DistanceKm calculates the shortest path between two coordinates on the
// surface of the earth in kilometers.
func DistanceKm(pLat, pLon, qLat, qLon float64) float64 {
	c := distance(pLat, pLon, qLat, qLon)
	km := convertToKm(c)
	return km
}

// DistanceMiles calculates the shortest path between two coordinates on the
// surface of the earth in miles.
func DistanceMiles(pLat, pLon, qLat, qLon float64) float64 {
	c := distance(pLat, pLon, qLat, qLon)
	miles := convertToMiles(c)
	return miles
}
