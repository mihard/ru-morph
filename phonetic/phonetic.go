package phonetic

func IsVowel(l string) bool {
	if l == "а" || l == "е" || l == "ё" || l == "и" || l == "о" || l == "у" || l == "ы" || l == "ю" || l == "я" {
		return true
	}

	return false
}

func IsVoiceless(l string) bool {
	if l == "п" || l == "ф" || l == "к" || l == "т" || l == "ш" || l == "с" || l == "х" || l == "ц" || l == "ч" || l == "щ" {
		return true
	}

	return false
}
//
//func IsVoiced(l string) bool {
//	if l == "б" || l == "в" || l == "г" || l == "д" || l == "ж" || l == "з" || l == "й" || l == "л" || l == "м" || l == "н" || l == "р" {
//		return true
//	}
//
//	return false
//}
