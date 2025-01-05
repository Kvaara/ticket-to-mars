package main

import (
	"fmt"
	"math"
	"math/rand"
	"strings"
)

const (
	DEPARTURE_DATE       string  = "October 13, 2020"
	NUMBER_OF_TICKETS    int     = 10
	MARS_DISTANCE_IN_KM  int     = 62_100_100
	MIN_COST_IN_MILLIONS float64 = 36
	MAX_COST_IN_MILLIONS float64 = 50
	MIN_SPEED_IN_KMS     float64 = 16.0
	MAX_SPEED_IN_KMS     float64 = 30.0
)

func main() {
	// This would be an array:
	// spacelines := [3]string{"Virgin Galactic", "SpaceX", "Space Adventures"}
	// This is a dynamic slice:
	spacelines := []string{"Virgin Galactic", "SpaceX", "Space Adventures"}

	chosen_spacelines := generate_spacelines(spacelines, NUMBER_OF_TICKETS)
	header := fmt.Sprintf("%-20s %-5s %-20s %s", "Spaceline", "Days", "Trip type", "Price")
	fmt.Println(header)
	// We have to type cast the header into a slice of runes before calculating its length using the len() function.
	// Why? Because len(string) returns the number of bytes not the number of characters (runes).
	// So len("a") will always be 1, but len("ä") will always be 2 even though it's a single character.
	// ASCII characters normally take 1 byte (8 bits) and Unicode characters take 2 bytes (16 bits) of memory.
	// A string is an immutable sequence of bytes (1 byte equals 8 bits) and not a sequence of runes.
	// A rune is an alias to int32 and denotes a unicode code point (e.g., the decimal unicode code point 83 equals 'S').
	// So for the accurate numbers of "=" characters in the separator we have to type cast into a slice of runes.
	// See https://programming.guide/go/rune.html
	// We could have used utf8.RuneCountInString function but the less packages, the better.
	fmt.Println(strings.Repeat("=", len([]rune(header))))
	for _, spaceline := range chosen_spacelines {
		chosen_speed, chosen_cost := generate_speed_and_calculate_cost()
		trip_type := generate_trip_type()
		days := calculate_days(chosen_speed)

		fmt.Printf("%-20s %-5d %-20s $%dM \n", spaceline, days, trip_type, int(math.Round(chosen_cost)))
	}
}

func generate_spacelines(spacelines []string, number_of_tickets int) []string {
	chosen_spacelines := make([]string, number_of_tickets)

	for i := 0; i < number_of_tickets; i++ {
		random_index := rand.Intn(3)
		chosen_spacelines[i] = spacelines[random_index]
	}

	return chosen_spacelines
}

func generate_speed_and_calculate_cost() (float64, float64) {
	randomFactor := rand.Float64()

	speed_interval := MAX_SPEED_IN_KMS - MIN_SPEED_IN_KMS
	chosen_speed := (randomFactor * speed_interval) + MIN_SPEED_IN_KMS

	cost_interval := MAX_COST_IN_MILLIONS - MIN_COST_IN_MILLIONS
	chosen_cost := (randomFactor * cost_interval) + MIN_COST_IN_MILLIONS

	return chosen_speed, chosen_cost
}

func generate_trip_type() string {
	ship_types := [2]string{"Round-trip", "One-way"}
	chosen_ship_type := ship_types[rand.Int31n(2)]
	return chosen_ship_type
}

func calculate_days(speed_in_kms float64) int {
	seconds := int(float64(MARS_DISTANCE_IN_KM) / speed_in_kms)
	minutes := seconds / 60
	hours := minutes / 60
	days := hours / 24
	return days
}
