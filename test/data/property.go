package data

import "creepy/internal/models"

var Properties []models.Property

func init() {
	Properties = []models.Property{
		{
			Title: "Spacious 2-Bedroom Apartment",
			Description: "A modern 2-bedroom apartment in the heart of the city with an elevator and parking.",
			BuyPrice: 450000,
			RentPrice: 2500,
			RentPriceMin: 2300,
			RentPriceMax: 2700,
			RahnPriceMin: 2000,
			RahnPriceMax: 3000,
			Area: 80,
			Rooms: 2,
			DealingType: "buy",
			Type: "apartment",
			City: "New York",
			District: "Manhattan",
			Address: "123 Main St, Apt 45",
			BuildYear: 2018,
			Floor: 5,
			HasElevator: true,
			HasStorage: true,
			HasParking: true,
			Latitude: 40.7128,
			Longitude: -74.0060,
			Source: "realtor.com",
			URL: "https://realtor.com/123-main-apt45",
			Image: "https://example.com/images/apt45.jpg",
		},
		{
			Title: "Cozy Studio Apartment for Rent",
			Description: "A charming studio apartment ideal for singles or students, in a quiet neighborhood with easy access to public transport.",
			BuyPrice: 180000,
			RentPrice: 1200,
			RentPriceMin: 1100,
			RentPriceMax: 1300,
			RahnPriceMin: 1500,
			RahnPriceMax: 2000,
			Area: 40,
			Rooms: 1,
			DealingType: "rent",
			Type: "studio",
			City: "Los Angeles",
			District: "Downtown",
			Address: "456 Elm St, Apt 7",
			BuildYear: 1995,
			Floor: 2,
			HasElevator: false,
			HasStorage: true,
			HasParking: false,
			Latitude: 34.0522,
			Longitude: -118.2437,
			Source: "zillow.com",
			URL: "https://zillow.com/456-elm-apt7",
			Image: "https://example.com/images/studio7.jpg",
		},
		{
			Title: "Luxury 4-Bedroom House with Pool",
			Description: "A stunning 4-bedroom house with a private pool, large backyard, and modern amenities. Perfect for families.",
			BuyPrice: 650000,
			RentPrice: 3500,
			RentPriceMin: 3300,
			RentPriceMax: 3800,
			RahnPriceMin: 2500,
			RahnPriceMax: 4000,
			Area: 200,
			Rooms: 4,
			DealingType: "buy",
			Type: "house",
			City: "Los Angeles",
			District: "Beverly Hills",
			Address: "789 Oakwood Dr",
			BuildYear: 2020,
			Floor: 1,
			HasElevator: true,
			HasStorage: true,
			HasParking: true,
			Latitude: 34.0696,
			Longitude: -118.4053,
			Source: "zillow.com",
			URL: "https://zillow.com/789-oakwood-dr",
			Image: "https://example.com/images/oakwooddr.jpg",
		},
		{
			Title: "Modern 1-Bedroom Apartment in London",
			Description: "A sleek 1-bedroom apartment located near the city center, ideal for professionals.",
			BuyPrice: 450000,
			RentPrice: 2000,
			RentPriceMin: 1800,
			RentPriceMax: 2200,
			RahnPriceMin: 2200,
			RahnPriceMax: 2500,
			Area: 60,
			Rooms: 1,
			DealingType: "rent",
			Type: "apartment",
			City: "London",
			District: "Westminster",
			Address: "101 Regent St, Apt 4",
			BuildYear: 2019,
			Floor: 3,
			HasElevator: true,
			HasStorage: false,
			HasParking: false,
			Latitude: 51.5074,
			Longitude: -0.1278,
			Source: "propertyfinder.com",
			URL: "https://propertyfinder.com/101-regent-apt4",
			Image: "https://example.com/images/regentst.jpg",
		},
	}	
}
