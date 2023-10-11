package handler

import "errors"

var (
	idNotFound = errors.New("param `id` not found")

	idMustBeEmpty = errors.New("param `id` must be empty")

	SpectrumCannotBeEmpty = errors.New("spectrum name cannot be empty")

	headerNotFound = errors.New("no file uploaded")

	SatelliteIDOrSpectrumIDIsEmpty = errors.New("satellite or spectrum cannot be empty")

	SatelliteNumberCannotBeEmpty = errors.New("param `satellite_number` cannot be empty")
)
