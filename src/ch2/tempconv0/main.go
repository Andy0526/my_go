package tempconv0

type Celsius float64    // 摄⽒温度
type Fahrenheit float64 // 华⽒温度
const (
	AbsoluteZeroC Celsius = -273.15 // 绝对零度
	FreezingC     Celsius = 0       // 结冰点温度
	BoilingC      Celsius = 100     // 沸⽔温度
)

func CToF(celsius Celsius) Fahrenheit {
	return Fahrenheit(celsius*9/5 + 32)
}

func FToC(fahrenheit Fahrenheit) Celsius {
	return Celsius((fahrenheit - 32) * 5 / 9)
}
