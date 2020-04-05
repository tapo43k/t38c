package t38c

import (
	"strconv"
	"strings"
)

// SearchOption ...
type SearchOption string

// Sparse will distribute the results of a search evenly across the requested area.
func Sparse(n int) SearchOption {
	return SearchOption(
		"SPARSE " + strconv.Itoa(n),
	)
}

// Where allows for filtering out results based on field values.
func Where(field string, min, max float64) SearchOption {
	return SearchOption(
		"WHERE " + field + " " + floatToString(min) + " " + floatToString(max),
	)
}

// Wherein is similar to Where except that it checks whether the object’s field value is in a given list.
func Wherein(field string, values ...float64) SearchOption {
	var sb strings.Builder
	sb.WriteString("WHEREIN " + field + " " + strconv.Itoa(len(values)))
	for _, val := range values {
		sb.WriteString(" " + floatToString(val))
	}
	return SearchOption(sb.String())
}

// Match is similar to WHERE except that it works on the object id instead of fields.
func Match(pattern string) SearchOption {
	return SearchOption(
		"MATCH " + pattern,
	)
}

// SetOption ...
type SetOption string

// SetField ...
func SetField(name string, value float64) SetOption {
	return SetOption(
		"FIELD " + name + " " + floatToString(value),
	)
}

// func SetEX(d time.Duration) SetOption {
// 	return SetOption(
// 		"EX ",
// 	)
// }

// SetNX ...
func SetNX() SetOption {
	return SetOption("NX")
}

// SetXX ...
func SetXX() SetOption {
	return SetOption("XX")
}
