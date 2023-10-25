package agent

// selectCountry determines country by probability.
func selectCountry(countries []country) string {
	var res country
	var max float32 = 0
	for _, v := range countries {
		if max < v.Probability {
			max = v.Probability
			res = v
		}
	}

	return res.CountryID
}
