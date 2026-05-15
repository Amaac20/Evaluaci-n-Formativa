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
		return nil, errors.New("El nombre no puede estar vacío")
	}
	if manufacturer == "" {
		return nil, errors.New("El fabricante no puede estar vacío")
	}
	if shelfLife <= 0 {
		return nil, errors.New("La duración de la vida útil debe ser positiva")
	}

	return &Medicine{
		name:            name,
		manufacturer:    manufacturer,
		manufactureDate: manufactureDate,
		shelfLife:       shelfLife,
	}, nil
}

// Getters
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

// Setters
func (m *Medicine) SetName(name string) error {
	if name == "" {
		return errors.New("El nombre no puede estar vacío")
	}
	m.name = name
	return nil
}

func (m *Medicine) SetManufacturer(manufacturer string) error {
	if manufacturer == "" {
		return errors.New("El fabricante no puede estar vacío")
	}
	m.manufacturer = manufacturer
	return nil
}

func (m *Medicine) SetManufactureDate(manufactureDate time.Time) {
	m.manufactureDate = manufactureDate
}

func (m *Medicine) SetShelfLife(shelfLife int) error {
	if shelfLife <= 0 {
		return errors.New("La duración de la vida útil debe ser positiva")
	}
	m.shelfLife = shelfLife
	return nil
}

// Método para calcular fecha de caducidad
func (m *Medicine) ExpirationDate() time.Time {
	return m.manufactureDate.AddDate(0, m.shelfLife, 0)
}

// ================= TABLET =================
type Tablet struct {
	*Medicine
	dosePerTablet  float64
	isPrescription bool
}

func NewTablet(name string, manufacturer string, manufactureDate time.Time, shelfLife int, dosePerTablet float64, isPrescription bool) (*Tablet, error) {
	if dosePerTablet <= 0 {
		return nil, errors.New("La dosis debe ser mayor a 0")
	}

	med, err := NewMedicine(name, manufacturer, manufactureDate, shelfLife)
	if err != nil {
		return nil, err
	}

	return &Tablet{
		Medicine:       med,
		dosePerTablet:  dosePerTablet,
		isPrescription: isPrescription,
	}, nil
}

func (t *Tablet) ShowDetails() {
	fmt.Println("=== TABLET ===")
	fmt.Println("Nombre:", t.GetName())
	fmt.Println("Fabricante:", t.GetManufacturer())
	fmt.Println("Fecha de fabricación:", t.GetManufactureDate().Format("2006-01-02"))
	fmt.Println("Vida útil:", t.GetShelfLife(), "meses")
	fmt.Println("Dosis:", t.dosePerTablet, "mg")
	fmt.Println("Requiere receta:", t.isPrescription)
	fmt.Println("Caduca:", t.ExpirationDate().Format("2006-01-02"))
	fmt.Println()
}

// ================= SYRUP =================
type Syrup struct {
	*Medicine
	volume float64
	flavor string
}

func NewSyrup(name string, manufacturer string, manufactureDate time.Time, shelfLife int, volume float64, flavor string) (*Syrup, error) {
	if volume <= 0 {
		return nil, errors.New("El volumen debe ser mayor a 0")
	}

	med, err := NewMedicine(name, manufacturer, manufactureDate, shelfLife)
	if err != nil {
		return nil, err
	}

	return &Syrup{
		Medicine: med,
		volume:   volume,
		flavor:   flavor,
	}, nil
}

func (s *Syrup) ShowDetails() {
	fmt.Println("=== SYRUP ===")
	fmt.Println("Nombre:", s.GetName())
	fmt.Println("Fabricante:", s.GetManufacturer())
	fmt.Println("Fecha de fabricación:", s.GetManufactureDate().Format("2006-01-02"))
	fmt.Println("Vida útil:", s.GetShelfLife(), "meses")
	fmt.Println("Volumen:", s.volume, "ml")
	fmt.Println("Sabor:", s.flavor)
	fmt.Println("Caduca:", s.ExpirationDate().Format("2006-01-02"))
	fmt.Println()
}

// ================= MAIN =================
func main() {
	date1, _ := time.Parse("2006-01-02", "2025-01-10")
	date2, _ := time.Parse("2006-01-02", "2025-03-15")

	// Inventario de tabletas
	tablet1, _ := NewTablet("Paracetamol", "Bayer", date1, 24, 500, false)
	tablet2, _ := NewTablet("Ibuprofeno", "Pfizer", date2, 18, 400, true)

	// Inventario de jarabes
	syrup1, _ := NewSyrup("Jarabe para la Tos", "MK", date1, 12, 120, "Cereza")
	syrup2, _ := NewSyrup("Vitamina C", "Genfar", date2, 10, 150, "Naranja")

	tablets := []*Tablet{tablet1, tablet2}
	syrups := []*Syrup{syrup1, syrup2}

	/
	fmt.Println("===== INVENTARIO DE TABLETAS =====")
	for _, t := range tablets {
		t.ShowDetails()
	}

	fmt.Println("===== INVENTARIO DE JARABES =====")
	for _, s := range syrups {
		s.ShowDetails()
	}

	fmt.Println("=== ACTUALIZACIÓN ===")
	tablet1.SetName("Paracetamol Extra Forte")
	fmt.Println("Nuevo nombre:", tablet1.GetName())

	// Mostrar fechas de caducidad
	fmt.Println("\n=== FECHAS DE CADUCIDAD ===")
	for _, t := range tablets {
		fmt.Printf("%s caduca el: %s\n", t.GetName(), t.ExpirationDate().Format("2006-01-02"))
	}

	for _, s := range syrups {
		fmt.Printf("%s caduca el: %s\n", s.GetName(), s.ExpirationDate().Format("2006-01-02"))
	}
}
