package direct

import (
	"strings"
	"testing"
)

var data = []Data{
	{
		Name: strings.Repeat("b", 1000),
	},
	{
		Name: strings.Repeat("b", 10000),
	},
	{
		Name: strings.Repeat("b", 100000),
	},
}

func BenchmarkConvertData(b *testing.B) {
	converter := Converter{}
	for _, d := range data {
		f := func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				converter.ConvertData(&d)
			}
		}

		b.Run("ConvertData", f)
	}
}

func BenchmarkConvertInterface(b *testing.B) {
	converter := Converter{}
	for _, d := range data {
		f := func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				converter.ConvertInterface(&d)
			}
		}

		b.Run("ConvertData", f)
	}
}

// func BenchmarkGetString(b *testing.B) {
// 	for _, v := range table {
// 		f := func(b *testing.B) {
// 			for i := 0; i < b.N; i++ {
// 				GetString(v)
// 			}
// 		}

// 		b.Run(v.Name, f)
// 	}
// }

// func BenchmarkGetStringPointer(b *testing.B) {
// 	for _, v := range table {
// 		f := func(b *testing.B) {
// 			for i := 0; i < b.N; i++ {
// 				GetStringPointer(&v)
// 			}
// 		}

// 		b.Run(v.Name, f)
// 	}
// }

// func BenchmarkSetString(b *testing.B) {
// 	for _, v := range table {
// 		f := func(b *testing.B) {
// 			for i := 0; i < b.N; i++ {
// 				v.String = SetString(v)
// 			}
// 		}

// 		b.Run(v.Name, f)
// 	}
// }

// func BenchmarkSetStringPointer(b *testing.B) {
// 	for _, v := range table {
// 		f := func(b *testing.B) {
// 			for i := 0; i < b.N; i++ {
// 				SetStringPointer(&v)
// 			}
// 		}

// 		b.Run(v.Name, f)
// 	}
// }
