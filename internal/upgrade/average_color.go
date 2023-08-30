package upgrade

import (
	"fmt"
	"regexp"
	"strconv"
)

func AverageColor() {
	var coefficient float64

	color1 := "#ff0000"
	color2 := "#ffffff"
	coefficient = 0.8

	averageColor, err := calculateAverageColor(color1, color2, &coefficient)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Average color:", averageColor)
}

func calculateAverageColor(color1, color2 string, c *float64) (string, error) {
	// Устанавливаем коэффициент, если не передан
	coef := 0.5
	if c != nil {
		coef = *c
	}

	r, _ := regexp.Compile("^#([A-Fa-f0-9]{6}|[A-Fa-f0-9]{3})$")

	// Валидируем переданные строки
	match1 := r.MatchString(color1)
	if !match1 {
		return "", fmt.Errorf("parse color1 error")
	}
	match2 := r.MatchString(color2)
	if !match2 {
		return "", fmt.Errorf("parse color2 error")
	}

	// Первый цвет по компонентам (string)
	rs1 := color1[1:3]
	gs1 := color1[3:5]
	bs1 := color1[5:7]

	// Второй цвет по компонентам (string)
	rs2 := color2[1:3]
	gs2 := color2[3:5]
	bs2 := color2[5:7]

	// Парсим компоненты цветов в int
	r1, _ := strconv.ParseInt(rs1, 16, 0)
	g1, _ := strconv.ParseInt(gs1, 16, 0)
	b1, _ := strconv.ParseInt(bs1, 16, 0)

	r2, _ := strconv.ParseInt(rs2, 16, 0)
	g2, _ := strconv.ParseInt(gs2, 16, 0)
	b2, _ := strconv.ParseInt(bs2, 16, 0)

	// Вычисляем средние значения компонентов цветов
	avgR := float64(r1)*(1-coef) + float64(r2)*coef
	avgG := float64(g1)*(1-coef) + float64(g2)*coef
	avgB := float64(b1)*(1-coef) + float64(b2)*coef

	// Преобразуем средние значения обратно в шестнадцатеричное представление цвета
	averageColor := fmt.Sprintf("#%02x%02x%02x", int(avgR), int(avgG), int(avgB))

	return averageColor, nil
}
