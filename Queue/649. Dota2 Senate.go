package Queue

func predictPartyVictory(senate string) string {
	// Menyimpan indeks senator dari Radiant dan Dire
	radiantQueue := []int{}
	direQueue := []int{}

	// Mengisi antrian dengan indeks senator masing-masing
	for i, char := range senate {
		if char == 'R' {
			radiantQueue = append(radiantQueue, i)
		} else if char == 'D' {
			direQueue = append(direQueue, i)
		}
	}

	// Simulasi perputaran pemilihan
	n := len(senate)
	for len(radiantQueue) > 0 && len(direQueue) > 0 {
		radiantIndex := radiantQueue[0]
		direIndex := direQueue[0]
		radiantQueue = radiantQueue[1:]
		direQueue = direQueue[1:]

		if radiantIndex < direIndex {
			// Radiant menang putaran ini, tambahkan senator Radiant ke akhir antrian
			radiantQueue = append(radiantQueue, radiantIndex+n)
		} else {
			// Dire menang putaran ini, tambahkan senator Dire ke akhir antrian
			direQueue = append(direQueue, direIndex+n)
		}
	}

	// Menentukan pemenang berdasarkan antrian yang tersisa
	if len(radiantQueue) > 0 {
		return "Radiant"
	} else {
		return "Dire"
	}
}
