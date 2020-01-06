package prediction

import "testing"

// TestAnalysis ...
func TestAnalysis(t *testing.T) {
	var closeData = []float32{7.80, 8.10, 8.10, 8.20, 8.25, 8.19, 8.03, 8.22, 8.31, 8.36, 8.30, 8.25, 8.20, 8.19, 7.89, 8.02, 8.01, 8.03, 7.90, 7.75, 7.97, 7.86, 8.00, 8.15, 8.30, 8.35, 8.39, 8.11, 8.30, 8.26, 8.48, 8.55, 8.61, 8.21, 8.10, 8.16, 8.04, 8.04, 7.94, 8.12, 8.18, 7.83, 7.52, 7.45, 7.65, 7.76, 7.64, 7.87, 7.71, 7.80, 7.42, 7.03, 7.11, 7.27, 7.45, 7.66, 7.52, 7.77, 7.91, 8.14, 8.10, 8.27, 8.40, 8.53, 8.59, 8.81, 8.90, 8.93, 8.96, 9.01, 9.09, 9.25, 8.92, 8.93, 8.93, 9.03, 9.30, 9.21, 9.32, 9.47, 9.35, 9.31, 9.34, 9.37, 9.30, 9.44, 9.08, 9.08, 9.15, 9.22, 9.20, 9.19, 9.05, 9.11, 9.19, 9.16, 9.20, 9.31, 9.11, 9.06, 9.25, 9.16, 9.25, 9.30, 9.02, 8.91, 9.03, 9.21, 9.16, 9.19, 9.37, 9.27, 9.55, 9.61, 9.86, 10.10, 10.05, 10.15, 10.10, 10.35, 10.55, 10.75, 10.70, 10.90, 10.85, 11.05, 10.85, 10.85, 11.10, 11.10, 11.25, 11.20, 11.25, 11.30, 11.40, 11.30, 11.40, 11.40, 11.45, 11.40, 11.55, 11.50, 11.30, 11.25, 11.30, 11.35, 11.20, 11.40, 11.55, 11.50, 11.60, 11.40, 11.45, 11.40, 11.55, 11.45, 11.40, 11.60, 11.65, 11.50, 11.55, 11.60, 11.70, 11.40, 11.40, 11.35, 11.10, 11.05, 11.05, 10.95, 11.10, 11.00, 11.00, 10.90, 11.05, 10.95, 10.95, 11.25, 11.90, 12.75, 12.60, 12.65, 12.45, 12.40, 12.30, 12.45, 12.45, 12.45, 12.40, 12.55, 12.45, 12.45, 12.45, 12.45, 12.15, 12.10, 11.90, 11.90, 11.95, 12.40, 12.50, 12.35, 12.35, 12.45, 12.45, 12.30, 12.35, 12.35, 12.50, 12.60, 12.60, 12.60, 12.55, 12.70, 12.50, 12.40, 12.60, 12.60, 12.55, 12.55, 12.75, 13.05, 13.15, 13.05, 13.30, 12.95, 12.90, 12.70, 12.65, 12.45, 12.40, 12.25, 12.25, 12.20, 12.30, 12.40, 12.40, 12.50, 12.55, 12.50, 12.35, 12.45, 12.35, 12.35, 12.35, 12.55, 12.45, 12.25, 12.35, 12.60, 12.55, 12.25, 12.25, 12.30, 12.20, 12.10, 12.20, 12.35, 12.35, 12.15, 12.15, 12.15, 12.15, 12.05, 12.20, 12.15, 12.30, 12.45, 12.65, 12.70, 12.60, 13.05, 12.45, 12.40, 12.40, 12.30, 12.30, 12.35, 12.20, 12.25, 12.55}
	var inputData InputData
	inputData.DayOp = 7.95
	inputData.DayCL = 7.80
	for index, count := 0, len(closeData); index < count; index++ {
		if len(inputData.Day3) < 3 {
			inputData.Day3 = append(inputData.Day3, closeData[index])
		}
		if len(inputData.Day7) < 7 {
			inputData.Day7 = append(inputData.Day7, closeData[index])
		}
		if len(inputData.Day14) < 14 {
			inputData.Day14 = append(inputData.Day14, closeData[index])
		}
		if len(inputData.Day30) < 30 {
			inputData.Day30 = append(inputData.Day30, closeData[index])
		}
		if len(inputData.Day60) < 60 {
			inputData.Day60 = append(inputData.Day60, closeData[index])
		}
		if len(inputData.Day90) < 90 {
			inputData.Day90 = append(inputData.Day90, closeData[index])
		}
		if len(inputData.Day120) < 120 {
			inputData.Day120 = append(inputData.Day120, closeData[index])
		}
		if len(inputData.Day180) < 180 {
			inputData.Day180 = append(inputData.Day180, closeData[index])
		}
		if len(inputData.Day240) < 240 {
			inputData.Day240 = append(inputData.Day240, closeData[index])
		}
		if len(inputData.Year1) < 365 {
			inputData.Year1 = append(inputData.Year1, closeData[index])
		}

	}

	reverse(inputData.Day3)
	reverse(inputData.Day7)
	reverse(inputData.Day14)
	reverse(inputData.Day30)
	reverse(inputData.Day60)
	reverse(inputData.Day90)
	reverse(inputData.Day120)
	reverse(inputData.Day180)
	reverse(inputData.Day240)
	reverse(inputData.Year1)

	Analysis(inputData)
}
func reverse(s []float32) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}
