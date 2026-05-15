package main

import (
	"errors"
	"fmt"
	"time"
)

type Medicine struct {
	name            string
	manufacturer    string
	manufactureDate time.Time
	shelfLife       int
}

func NewMedicine(name string, manufacturer string, manufactureDate time.Time, shelfLife int) (*Medicine, error) {
	if name == "" {
		return nil, errors.New("El nombre no puede estar vacio")
	}
	if manufacturer == "" {
		return nil, errors.New("El fabricante no puede estar vacio")
	}
	if shelfLife <= 0 {
		return nil, errors.New("La duracion de la vida util en meses debe ser un numero positivo")
	}

	return &Medicine{
		name:            name,
		manufacturer:    manufacturer,
		manufactureDate: manufactureDate,
		shelfLife:       shelfLife,
	}, nil
}

func (m *Medicine) GetName() string {
	return m.name
}

func (m *Medicine) GetManufacturer() string {
	return m.manufacturer
}

func (m *Medicine) GetManufactureDate() time.Time {
	return m.manufactureDate
}
func (m *Medicine) GetShelfLife() int {
	return m.shelfLife
}

func (m *Medicine) SetName(name string) error {
	if name == "" {
		return errors.New("El nombre no puede estar vacio")
	}
	m.name = name
	return nil
}

func (m *Medicine) SetManufacturer(manufacturer string) error {
	if manufacturer == "" {
		return errors.New("El fabricante no puede estar vacio")
	}
	m.manufacturer = manufacturer
	return nil
}

func (m *Medicine) SetManufactureDate(manufactureDate time.Time) {
	m.manufactureDate = manufactureDate
}

func (m *Medicine) SetShelfLife(shelfLife int) error {
	if shelfLife <= 0 {
		return errors.New("La duracion de la vida util en meses debe ser un numero positivo")
	}
	m.shelfLife = shelfLife
	return nil
}

func (m *Medicine) ExpirationDate() time.Time {
	return m.manufactureDate.AddDate(0, m.shelfLife, 0)
}

type Tablet struct {
	*Medicine
	dosePerTablet  float64
	isPrescription bool
}

func NewTablet(name string, manufacturer string, manufactureDate time.Time, shelfLife int, dosePerTablet float64, isPrescription bool) (*Tablet, error) {
	if dose <= 0 {
		return nil, errors.New("la dosis debe ser mayor a 0")
	}
	med, err := NewMedicine(name, manufacturer, manufactureDate, shelfLife)
	if err != nil {
		return nil, err
	}
	return &Tablet{
		Medicine:        med,
		dosePerTablet:   dose,
		isPrescription: prescription,
	}, nil
}

func (t *Tablet) ShowDetails() {
	fmt.Println("=== TABLET ===")
	fmt.Println("Nombre:", t.GetName())
	fmt.Println("Fabricante:", t.GetManufacturer())
	fmt.Println("Fecha de fabricacion:", t.GetManufactureDate().Format("2015-01-02"))
	fmt.Println("Vida util:", t.GetShelfLife(), "meses")
	fmt.Println("Dosis:", t.dosePerTablet, "mg")
	fmt.Println("Requiere receta:", t.isPrescription)
	fmt.Println("Caduca:", t.ExpirationDate().Format("2024-01-02"))
	fmt.Println()
}

type Syrup struct {
	*Medicine
	volumen float64
	flavor  string
}

func NewSyrup(name string, manufacturer string, manufactureDate time.Time, shelfLife int, volume float64, flavor string) (*Syrup, error){
	if volume <=0{
		return nil, errors.New("El volumen debe ser mayor a 0")
	}
	med, err := NewMedicine(name, manufacturer, manufactureDate, shelfLife)
	if err != nil {
		return nil, err
	}
	return &Syrup{
		Medicine: med,
		volume: volume,
		flavor: flavor,
	}, nil
}
func (s *Syrup) ShowDetails(){
	fmt.Println("=== Syrup ===")
	fmt.Println("Nombre:", s.GetName())
	fmt.Println("Fabricante:", s.GetManufacturer())
	fmt.Println("Fecha de fabricacion:", s.GetManufactureDate().Format("2015-01-02"))
	fmt.Println("Vida util:", s.GetShelfLife(), "meses")
	fmt.Println("Volumen: ", s.volume, "ml")
	fmt.Println("Sabor: ", s.flavor)
	fmt.Println("Caduca:", t.ExpirationDate().Format("2024-01-02"))
	fmt.Println()
}
func main() {
	date1, 

}
