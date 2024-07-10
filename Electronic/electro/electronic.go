package electronic

type Phone interface {
	Brand() string
	Model() string
	Type() string
}
type StationPhone interface {
	ButtonsCount() int
}
type Smartphone interface {
	OS() string
}
type applePhone struct {
	brand string
	model string
	os    string
}

func (ap applePhone) Brand() string {
	return ap.brand
}
func (ap applePhone) Model() string {
	return ap.model
}
func (ap applePhone) Type() string {
	return "smartphone"
}
func (ap applePhone) OS() string {
	return ap.os
}
func NewApplePhone(model, os string) Phone {
	return applePhone{
		brand: "Apple",
		model: model,
		os:    os,
	}
}

type androidPhone struct {
	brand string
	model string
	os    string
}

func (ap androidPhone) Brand() string {
	return ap.brand
}
func (ap androidPhone) Model() string {
	return ap.model
}
func (ap androidPhone) Type() string {
	return "smartphone"
}
func (ap androidPhone) OS() string {
	return ap.os
}
func NewAndroidPhone(brand, model, os string) Phone {
	return androidPhone{
		brand: brand,
		model: model,
		os:    os,
	}
}

type radioPhone struct {
	brand        string
	model        string
	buttonsCount int
}

func (rp radioPhone) Brand() string {
	return rp.brand
}
func (rp radioPhone) Model() string {
	return rp.model
}
func (rp radioPhone) Type() string {
	return "station"
}
func (rp radioPhone) ButtonsCount() int {
	return rp.buttonsCount
}
func NewRadioPhone(brand, model string, buttonsCount int) Phone {
	return radioPhone{
		brand:        brand,
		model:        model,
		buttonsCount: buttonsCount,
	}
}
