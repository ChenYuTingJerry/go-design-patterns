package builder

//A Builder of type bus
type BusBuilder struct {
	v VehicleProduct
}

func (b *BusBuilder) SetWheels() BuildProcess {
	b.v.Wheels = 8
	return b
}

func (b *BusBuilder) SetSeats() BuildProcess {
	b.v.Seats = 30
	return b
}

func (b *BusBuilder) SetStructure() BuildProcess {
	b.v.Structure = "Bus"
	return b
}

func (b *BusBuilder) GetVehicle() VehicleProduct {
	return b.v
}
