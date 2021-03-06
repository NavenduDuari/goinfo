package utils

func IsCmdValid(argsMap map[string]string) bool {
	if len(argsMap) == 0 {
		return true
	}

	for _, validArg := range CovidArgs {
		for arg := range argsMap {
			if validArg == arg {
				return true
			}
		}
	}
	return false
}

var CovidArgs = []string{"--state=", "--suggest", "--help"}

var States = map[string]string{
	"MH": "Maharashtra",
	"TN": "Tamil Nadu",
	"DL": "Delhi",
	"TG": "Telangana",
	"RJ": "Rajasthan",
	"KL": "Kerala",
	"UP": "Uttar Pradesh",
	"AP": "Andhra Pradesh",
	"MP": "Madhya Pradesh",
	"KA": "Karnataka",
	"GJ": "Gujarat",
	"HR": "Haryana",
	"JK": "Jammu and Kashmir",
	"PB": "Punjab",
	"WB": "West Bengal",
	"OR": "Odisha",
	"BR": "Bihar",
	"UT": "Uttarakhand",
	"AS": "Assam",
	"CH": "Chandigarh",
	"HP": "Himachal Pradesh",
	"LA": "Ladakh",
	"AN": "Andaman and Nicobar Islands",
	"CT": "Chhattisgarh",
	"GA": "Goa",
	"PY": "Puducherry",
	"JH": "Jharkhand",
	"MN": "Manipur",
	"MZ": "Mizoram",
	"AR": "Arunachal Pradesh",
	"DN": "Dadra and Nagar Haveli",
	"TR": "Tripura",
	"DD": "Daman and Diu",
	"LD": "Lakshadweep",
	"ML": "Meghalaya",
	"NL": "Nagaland",
	"SK": "Sikkim",
}
