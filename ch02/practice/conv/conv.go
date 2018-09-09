package conv


func CToF(c Celsius) Fahrenheit { return Fahrenheit(c*9/5 + 32) }
func FToC(f Fahrenheit) Celsius { return Celsius((f - 32) * 5 / 9) }
func CToK(c Celsius) Kelvin { return Kelvin(c + 273.15) }
func KToC(k Kelvin) Celsius { return Celsius(k - 273.15) }
func FToK(f Fahrenheit) Kelvin { return Kelvin(CToK(FToC(f))) }
func KToF(k Kelvin) Fahrenheit { return Fahrenheit(CToF(KToC(k))) }

func KToP(k Kilogram) Pound { return Pound(k * 0.454) }
func PToK(p Pound) Kilogram { return Kilogram(p / 0.454) }

func MToF(m Meter) Feet { return Feet(m / 0.3048) }
func FToM(f Feet) Meter { return Meter(f * 0.3048) }

