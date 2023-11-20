package medicine

import (
	"strings"

	"github.com/sahilm/fuzzy"
)

type MedicineItem struct {
	Name       string
	Type       string
	TimePeriod int
}

// MedicineList is a list of medicine
type MedicineList []MedicineItem

// Why not filter in database and not application code?

// Filter filters our medicine list by a pattern
func (m MedicineList) Filter(f FilterFunc) MedicineList {
	return f(m)
}

func (m MedicineList) String() string {
	b := strings.Builder{}
	for i, entry := range m {
		b.WriteString(entry.Name)
		// write newline up to second to last entry
		if i < len(m)-1 {
			b.WriteString("\n")
		}
	}
	return b.String()
}

type FilterFunc func(m MedicineList) MedicineList

func FilterByName(name string) FilterFunc {
	return func(medicineList MedicineList) MedicineList {
		matches := filterFunc(name, MedicineListNameSource(medicineList))
		return findMatches(medicineList, matches)
	}
}

func filterFunc(pattern string, source fuzzy.Source) fuzzy.Matches {
	return fuzzy.FindFrom(pattern, source)
}

func findMatches(medicineList MedicineList, matches fuzzy.Matches) MedicineList {
	result := make(MedicineList, matches.Len())
	for i := range result {
		itemIdx := matches[i].Index
		result[i] = medicineList[itemIdx]
	}
	return result
}

func FilterByType(pattern string) FilterFunc {
	return func(medicineList MedicineList) MedicineList {
		matches := filterFunc(pattern, MedicineListTypeSource(medicineList))
		return findMatches(medicineList, matches)
	}
}

type MedicineListNameSource MedicineList

func (m MedicineListNameSource) String(i int) string {
	return m[i].Name
}

func (m MedicineListNameSource) Len() int {
	return len(m)
}

type MedicineListTypeSource MedicineList

func (m MedicineListTypeSource) String(i int) string {
	return m[i].Type
}

func (m MedicineListTypeSource) Len() int {
	return len(m)
}
