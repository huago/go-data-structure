package tboxs

import (
	"math/rand"
	"time"
)

func Distribute(totalAmount float64, count int) []float64{
	totalAmountCent := int(totalAmount * 100)

	packets := make([]float64, count)

	// 随机分成count份
	total, every := divide(count)

	// 前count-1个红包金额
	var notIncludeLastOneAmount int
	for i := 0; i < count-1; i++ {
		packet := float64(every[i]) / float64(total) * float64(totalAmountCent)
		packetInt := int(packet)
		packets[i] = float64(packetInt)
		notIncludeLastOneAmount += packetInt
	}

	// 最后一个红包金额
	packets[count-1] = float64(totalAmountCent - notIncludeLastOneAmount)

	for i := 0; i < count; i++ {
		packets[i] /= 100
	}

	return packets
}

func divide(count int) (total int, every []int) {
	for i := 0; i < count; i++ {
		var random int
		for {
			source := rand.NewSource(time.Now().UnixNano())
			random = rand.New(source).Intn(100)
			if random != 0 {
				break
			}
		}

		total += random
		every = append(every, random)
	}

	return
}
