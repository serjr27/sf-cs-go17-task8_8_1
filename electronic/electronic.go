// electronic - пакет для создания объектов, содержащих информацию
// о телефонах и вывода этой информации на экран.
package electronic

type Phone interface {
	Brand() string
	Model() string
	Type() string
}

type StationPhone interface {
	ButtonCounts() int
}

type Smartphone interface {
	OS() string
}

type applePhone struct {
	model string
	os    string
}

// NewApplePhone - конструктор. Создает объект типа applePhone.
// Параметры: model : string - название модели телефона,
// os : string - наименование операционной системы.
// Параметры: brand и type не передаются для создания объекта.
// Подразумевается: brand = "Apple", type = "smartphone".
func NewApplePhone(model, os string) *applePhone {
	return &applePhone{model, os}
}

// Brand - реализация метода Brand() интерфейса Phone для объекта applePhone.
// Возвращает информацию о бренде, т.е. "Apple"
func (a applePhone) Brand() string { return "Apple" }

// Model - реализация метода Model() интерфейса Phone для объекта applePhone.
// Возвращает информацию о моделе телефона.
func (a applePhone) Model() string { return a.model }

// Type - реализация метода Type() интерфейса Phone для объекта applePhone.
// Возвращает информацию о типе телефона, т.е. "smartphone".
func (a applePhone) Type() string { return "smartphone" }

// OS - реализация метода OS() интерфейса Smarthone для объекта applePhone.
// Возвращает информацию о операционной системе.
func (a applePhone) OS() string { return a.os }

type androidPhone struct {
	brand string
	model string
	os    string
}

// NewAndroidPhone - конструктор. Создает объект типа androidPhone.
// Параметры: brand : string - наименование бренда, model : string -
// название модели телефона, os : string - наименование операционной системы.
// Параметр: type не передаются для создания объекта.
// Подразумевается: type = "smartphone".
func NewAndroidPhone(brand, model, os string) *androidPhone {
	return &androidPhone{brand, model, os}
}

// Brand - реализация метода Brand() интерфейса Phone для объекта androidPhone.
// Возвращает информацию о бренде."
func (a androidPhone) Brand() string { return a.brand }

// Model - реализация метода Model() интерфейса Phone для объекта androidPhone.
// Возвращает информацию о моделе телефона.
func (a androidPhone) Model() string { return a.model }

// Type - реализация метода Type() интерфейса Phone для объекта androidPhone.
// Возвращает информацию о типе телефона, т.е. "smartphone".
func (a androidPhone) Type() string { return "smartphone" }

// OS - реализация метода OS() интерфейса Smarthone для объекта androidPhone.
// Возвращает информацию о операционной системе.
func (a androidPhone) OS() string { return a.os }

type radioPhone struct {
	brand   string
	model   string
	buttons int
}

// NewRadioPhone - конструктор. Создает объект типа radioPhone.
// Параметры: brand : string - наименование бренда, model : string -
// название модели телефона, buttons : int - количество кнопок.
// Параметр: type не передаются для создания объекта.
// Подразумевается: type = "station".
func NewRadioPhone(brand, model string, buttons int) *radioPhone {
	return &radioPhone{brand, model, buttons}
}

// Brand - реализация метода Brand() интерфейса Phone для объекта radiodPhone.
// Возвращает информацию о бренде."
func (r radioPhone) Brand() string { return r.brand }

// Model - реализация метода Model() интерфейса Phone для объекта radioPhone.
// Возвращает информацию о моделе телефона.
func (r radioPhone) Model() string { return r.model }

// Type - реализация метода Type() интерфейса Phone для объекта radioPhone.
// Возвращает информацию о типе телефона, т.е. "station".
func (r radioPhone) Type() string { return "station" }

// ButtonCounts - реализация метода  интерфейса StationPhone для объекта radioPhone.
// Возвращает информацию о количестве кнопок в телефоне.
func (r radioPhone) ButtonCounts() int { return r.buttons }
