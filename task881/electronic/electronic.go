package electronic

type Phone interface {
	Brand() string
	Model() string
	Type() string
}

type StationPhone interface {
	ButtonsCount() int
}

type SmartPhone interface {
	OS() string
}

type phone struct {
	brand     string
	model     string
	phoneType string
}

type stationPhone struct {
	buttonsCount int
}

type smartPhone struct {
	os string
}

type applePhone struct {
	phone
	smartPhone
}

type androidPhone struct {
	phone
	smartPhone
}

type radioPhone struct {
	phone
	stationPhone
}

func (p phone) Brand() string {
	return p.brand
}

func (p phone) Model() string {
	return p.model
}

func (p phone) Type() string {
	return p.phoneType
}

func (s smartPhone) OS() string {
	return s.os
}

func (s stationPhone) ButtonsCount() int {
	return s.buttonsCount
}

func NewApplePhone(model, os string) applePhone {
	return applePhone{
		phone:      phone{brand: "Apple", model: model, phoneType: "smartphone"},
		smartPhone: smartPhone{os: os},
	}
}

func NewAndroidPhone(brand, model, os string) androidPhone {
	return androidPhone{
		phone:      phone{brand: brand, model: model, phoneType: "smartphone"},
		smartPhone: smartPhone{os: os},
	}
}

func NewRadioPhone(brand, model string, buttons int) radioPhone {
	return radioPhone{
		phone:        phone{brand: brand, model: model, phoneType: "station"},
		stationPhone: stationPhone{buttonsCount: buttons},
	}
}
