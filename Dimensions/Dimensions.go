//Я не со всем понял, что требовалось сделать по этому просто делал как было написано.
//Опишите 2 интерфейса: Auto и Dimensions, которые будут иметь следующий вид:
//Создайте 2 реализации интерфейса Dimensions: одна будет для дюймов, другая — для сантиметров.
//Создайте по 1 реализации интерфейса Auto для автомобилей BMW, Mercedes и Dodge. Для реализации Dodge, dimensions должны возвращаться в дюймах.
type InchesDimensions struct {
	length Unit
	width  Unit
	height Unit
  }
  
  func (id InchesDimensions) Length() Unit {
	return id.length
  }
  
  func (id InchesDimensions) Width() Unit {
	return id.width
  }
  
  func (id InchesDimensions) Height() Unit {
	return id.height
  }

  
  type CmDimensions struct {
	length Unit
	width  Unit
	height Unit
  }
  
  func (cd CmDimensions) Length() Unit {
	return cd.length
  }
  
  func (cd CmDimensions) Width() Unit {
	return cd.width
  }
  
  func (cd CmDimensions) Height() Unit {
	return cd.height
  }
  
  type BMW struct {
	brand        string
	model        string
	dimensions   Dimensions
	maxSpeed     int
	enginePower  int
  }
  
  func (b BMW) Brand() string {
	return b.brand
  }
  
  func (b BMW) Model() string {
	return b.model
  }
  
  func (b BMW) Dimensions() Dimensions {
	return b.dimensions
  }
  
  func (b BMW) MaxSpeed() int {
	return b.maxSpeed
  }
  
  func (b BMW) EnginePower() int {
	return b.enginePower
  }
  
  type Mercedes struct {
	brand        string
	model        string
	dimensions   Dimensions
	maxSpeed     int
	enginePower  int
  }
  
  func (m Mercedes) Brand() string {
	return m.brand
  }
  
  func (m Mercedes) Model() string {
	return m.model
  }
  
  func (m Mercedes) Dimensions() Dimensions {
	return m.dimensions
  }
  
  func (m Mercedes) MaxSpeed() int {
	return m.maxSpeed
  }
  
  func (m Mercedes) EnginePower() int {
	return m.enginePower
  }
  
  type Dodge struct {
	brand        string
	model        string
	dimensions   Dimensions
	maxSpeed     int
	enginePower  int
  }
  
  func (d Dodge) Brand() string {
	return d.brand
  }
  
  func (d Dodge) Model() string {
	return d.model
  }
  
  func (d Dodge) Dimensions() Dimensions {
	return d.dimensions
  }
  
  func (d Dodge) MaxSpeed() int {
	return d.maxSpeed
  }
  
  func (d Dodge) EnginePower() int {
	return d.enginePower
  }
  
  main() {
	inchesDimensions := InchesDimensions{
	  length: Unit{Value: 72, T: Inch},
	  width:  Unit{Value: 30, T: Inch},
	  height: Unit{Value: 48, T: Inch},
	}
  
	cmDimensions := CmDimensions{
	  length: Unit{Value: 183, T: CM},
	  width:  Unit{Value: 76, T: CM},
	  height: Unit{Value: 122, T: CM},
	}
	bmw := BMW{
	  brand:       "BMW",
	  model:       "X5",
	  dimensions:  inchesDimensions,
	  maxSpeed:    250,
	  enginePower: 300,
	}
	mercedes := Mercedes{
	  brand:       "Mercedes",
	  model:       "E-Class",
	  dimensions:  cmDimensions,
	  maxSpeed:    220,
	  enginePower: 280,
	}
	dodge := Dodge{
	  brand:       "Dodge",
	  model:       "Charger",
	  dimensions:  inchesDimensions,
	  maxSpeed:    240,
	  enginePower: 350,
	}
  }