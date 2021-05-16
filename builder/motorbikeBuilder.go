package builder

//A Builder of type motorbike
type MotorbikeBuilder struct {
	v VehicleProduct
}

func (b *MotorbikeBuilder) SetWheels() BuildProcess {
	b.v.Wheels = 2
	return b
}

func (b *MotorbikeBuilder) SetSeats() BuildProcess {
	b.v.Seats = 2
	return b
}

func (b *MotorbikeBuilder) SetStructure() BuildProcess {
	b.v.Structure = "Motorbike"
	return b
}

func (b *MotorbikeBuilder) GetVehicle() VehicleProduct {
	return b.v
}
