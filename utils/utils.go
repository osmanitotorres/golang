package utils

import (
	"regexp"
	"strconv"
)

// 0001 ===========  VALIDA CNPJ  =================//
func ValidaCNPJ(cnpj string) bool {
	// Remove caracteres não numéricos
	re := regexp.MustCompile(`[^\d]`)
	cnpj = re.ReplaceAllString(cnpj, "")

	// Verifica se o CNPJ tem 14 dígitos
	if len(cnpj) != 14 {
		return false
	}

	// CNPJs conhecidos como inválidos
	if cnpj == "00000000000000" || cnpj == "11111111111111" ||
		cnpj == "22222222222222" || cnpj == "33333333333333" ||
		cnpj == "44444444444444" || cnpj == "55555555555555" ||
		cnpj == "66666666666666" || cnpj == "77777777777777" ||
		cnpj == "88888888888888" || cnpj == "99999999999999" {
		return false
	}

	// Cálculo do primeiro dígito verificador
	weights := []int{5, 4, 3, 2, 9, 8, 7, 6, 5, 4, 3, 2}
	sum := 0
	for i, weight := range weights {
		num, _ := strconv.Atoi(string(cnpj[i]))
		sum += num * weight
	}
	d1 := 11 - (sum % 11)
	if d1 >= 10 {
		d1 = 0
	}

	// Verifica o primeiro dígito verificador
	if d1 != int(cnpj[12]-'0') {
		return false
	}

	// Cálculo do segundo dígito verificador
	weights = append([]int{6}, weights...)
	sum = 0
	for i, weight := range weights {
		num, _ := strconv.Atoi(string(cnpj[i]))
		sum += num * weight
	}
	d2 := 11 - (sum % 11)
	if d2 >= 10 {
		d2 = 0
	}

	// Verifica o segundo dígito verificador
	return d2 == int(cnpj[13]-'0')

}
