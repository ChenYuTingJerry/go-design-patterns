package builder

//Director
type ManufacturingDirector struct{
	builder BuildProcess
}

func (f *ManufacturingDirector) Construct() {
	f.builder.SetStructure().SetSeats().SetWheels()
}

func (f *ManufacturingDirector) SetBuilder(b BuildProcess) {
	f.builder = b
}
