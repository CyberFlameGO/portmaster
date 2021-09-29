package geoip

// FillMissingInfo tries to fill missing location information based on the
// available existing information.
func (l *Location) FillMissingInfo() {
	// Get coordinates from country.
	if l.Coordinates.Latitude == 0 &&
		l.Coordinates.Longitude == 0 &&
		l.Country.ISOCode != "" {
		if c, ok := countryCoordinates[l.Country.ISOCode]; ok {
			l.Coordinates = c
		}
	}
}

var countryCoordinates = map[string]Coordinates{
	"AD": {500, 42, 1},
	"AE": {500, 23, 53},
	"AF": {500, 33, 67},
	"AG": {500, 17, -61},
	"AI": {500, 18, -63},
	"AL": {500, 41, 20},
	"AM": {500, 40, 45},
	"AN": {500, 12, -69},
	"AO": {500, -11, 17},
	"AQ": {500, -75, -0},
	"AR": {500, -38, -63},
	"AS": {500, -14, -170},
	"AT": {500, 47, 14},
	"AU": {500, -25, 133},
	"AW": {500, 12, -69},
	"AZ": {500, 40, 47},
	"BA": {500, 43, 17},
	"BB": {500, 13, -59},
	"BD": {500, 23, 90},
	"BE": {500, 50, 4},
	"BF": {500, 12, -1},
	"BG": {500, 42, 25},
	"BH": {500, 25, 50},
	"BI": {500, -3, 29},
	"BJ": {500, 9, 2},
	"BM": {500, 32, -64},
	"BN": {500, 4, 114},
	"BO": {500, -16, -63},
	"BR": {500, -14, -51},
	"BS": {500, 25, -77},
	"BT": {500, 27, 90},
	"BV": {500, -54, 3},
	"BW": {500, -22, 24},
	"BY": {500, 53, 27},
	"BZ": {500, 17, -88},
	"CA": {500, 56, -106},
	"CC": {500, -12, 96},
	"CD": {500, -4, 21},
	"CF": {500, 6, 20},
	"CG": {500, -0, 15},
	"CH": {500, 46, 8},
	"CI": {500, 7, -5},
	"CK": {500, -21, -159},
	"CL": {500, -35, -71},
	"CM": {500, 7, 12},
	"CN": {500, 35, 104},
	"CO": {500, 4, -74},
	"CR": {500, 9, -83},
	"CU": {500, 21, -77},
	"CV": {500, 16, -24},
	"CX": {500, -10, 105},
	"CY": {500, 35, 33},
	"CZ": {500, 49, 15},
	"DE": {500, 51, 10},
	"DJ": {500, 11, 42},
	"DK": {500, 56, 9},
	"DM": {500, 15, -61},
	"DO": {500, 18, -70},
	"DZ": {500, 28, 1},
	"EC": {500, -1, -78},
	"EE": {500, 58, 25},
	"EG": {500, 26, 30},
	"EH": {500, 24, -12},
	"ER": {500, 15, 39},
	"ES": {500, 40, -3},
	"ET": {500, 9, 40},
	"FI": {500, 61, 25},
	"FJ": {500, -16, 179},
	"FK": {500, -51, -59},
	"FM": {500, 7, 150},
	"FO": {500, 61, -6},
	"FR": {500, 46, 2},
	"GA": {500, -0, 11},
	"GB": {500, 55, -3},
	"GD": {500, 12, -61},
	"GE": {500, 42, 43},
	"GF": {500, 3, -53},
	"GG": {500, 49, -2},
	"GH": {500, 7, -1},
	"GI": {500, 36, -5},
	"GL": {500, 71, -42},
	"GM": {500, 13, -15},
	"GN": {500, 9, -9},
	"GP": {500, 16, -62},
	"GQ": {500, 1, 10},
	"GR": {500, 39, 21},
	"GS": {500, -54, -36},
	"GT": {500, 15, -90},
	"GU": {500, 13, 144},
	"GW": {500, 11, -15},
	"GY": {500, 4, -58},
	"GZ": {500, 31, 34},
	"HK": {500, 22, 114},
	"HM": {500, -53, 73},
	"HN": {500, 15, -86},
	"HR": {500, 45, 15},
	"HT": {500, 18, -72},
	"HU": {500, 47, 19},
	"ID": {500, -0, 113},
	"IE": {500, 53, -8},
	"IL": {500, 31, 34},
	"IM": {500, 54, -4},
	"IN": {500, 20, 78},
	"IO": {500, -6, 71},
	"IQ": {500, 33, 43},
	"IR": {500, 32, 53},
	"IS": {500, 64, -19},
	"IT": {500, 41, 12},
	"JE": {500, 49, -2},
	"JM": {500, 18, -77},
	"JO": {500, 30, 36},
	"JP": {500, 36, 138},
	"KE": {500, -0, 37},
	"KG": {500, 41, 74},
	"KH": {500, 12, 104},
	"KI": {500, -3, -168},
	"KM": {500, -11, 43},
	"KN": {500, 17, -62},
	"KP": {500, 40, 127},
	"KR": {500, 35, 127},
	"KW": {500, 29, 47},
	"KY": {500, 19, -80},
	"KZ": {500, 48, 66},
	"LA": {500, 19, 102},
	"LB": {500, 33, 35},
	"LC": {500, 13, -60},
	"LI": {500, 47, 9},
	"LK": {500, 7, 80},
	"LR": {500, 6, -9},
	"LS": {500, -29, 28},
	"LT": {500, 55, 23},
	"LU": {500, 49, 6},
	"LV": {500, 56, 24},
	"LY": {500, 26, 17},
	"MA": {500, 31, -7},
	"MC": {500, 43, 7},
	"MD": {500, 47, 28},
	"ME": {500, 42, 19},
	"MG": {500, -18, 46},
	"MH": {500, 7, 171},
	"MK": {500, 41, 21},
	"ML": {500, 17, -3},
	"MM": {500, 21, 95},
	"MN": {500, 46, 103},
	"MO": {500, 22, 113},
	"MP": {500, 17, 145},
	"MQ": {500, 14, -61},
	"MR": {500, 21, -10},
	"MS": {500, 16, -62},
	"MT": {500, 35, 14},
	"MU": {500, -20, 57},
	"MV": {500, 3, 73},
	"MW": {500, -13, 34},
	"MX": {500, 23, -102},
	"MY": {500, 4, 101},
	"MZ": {500, -18, 35},
	"NA": {500, -22, 18},
	"NC": {500, -20, 165},
	"NE": {500, 17, 8},
	"NF": {500, -29, 167},
	"NG": {500, 9, 8},
	"NI": {500, 12, -85},
	"NL": {500, 52, 5},
	"NO": {500, 60, 8},
	"NP": {500, 28, 84},
	"NR": {500, -0, 166},
	"NU": {500, -19, -169},
	"NZ": {500, -40, 174},
	"OM": {500, 21, 55},
	"PA": {500, 8, -80},
	"PE": {500, -9, -75},
	"PF": {500, -17, -149},
	"PG": {500, -6, 143},
	"PH": {500, 12, 121},
	"PK": {500, 30, 69},
	"PL": {500, 51, 19},
	"PM": {500, 46, -56},
	"PN": {500, -24, -127},
	"PR": {500, 18, -66},
	"PS": {500, 31, 35},
	"PT": {500, 39, -8},
	"PW": {500, 7, 134},
	"PY": {500, -23, -58},
	"QA": {500, 25, 51},
	"RE": {500, -21, 55},
	"RO": {500, 45, 24},
	"RS": {500, 44, 21},
	"RU": {500, 61, 105},
	"RW": {500, -1, 29},
	"SA": {500, 23, 45},
	"SB": {500, -9, 160},
	"SC": {500, -4, 55},
	"SD": {500, 12, 30},
	"SE": {500, 60, 18},
	"SG": {500, 1, 103},
	"SH": {500, -24, -10},
	"SI": {500, 46, 14},
	"SJ": {500, 77, 23},
	"SK": {500, 48, 19},
	"SL": {500, 8, -11},
	"SM": {500, 43, 12},
	"SN": {500, 14, -14},
	"SO": {500, 5, 46},
	"SR": {500, 3, -56},
	"ST": {500, 0, 6},
	"SV": {500, 13, -88},
	"SY": {500, 34, 38},
	"SZ": {500, -26, 31},
	"TC": {500, 21, -71},
	"TD": {500, 15, 18},
	"TF": {500, -49, 69},
	"TG": {500, 8, 0},
	"TH": {500, 15, 100},
	"TJ": {500, 38, 71},
	"TK": {500, -8, -171},
	"TL": {500, -8, 125},
	"TM": {500, 38, 59},
	"TN": {500, 33, 9},
	"TO": {500, -21, -175},
	"TR": {500, 38, 35},
	"TT": {500, 10, -61},
	"TV": {500, -7, 177},
	"TW": {500, 23, 120},
	"TZ": {500, -6, 34},
	"UA": {500, 48, 31},
	"UG": {500, 1, 32},
	"US": {500, 37, -95},
	"UY": {500, -32, -55},
	"UZ": {500, 41, 64},
	"VA": {500, 41, 12},
	"VC": {500, 12, -61},
	"VE": {500, 6, -66},
	"VG": {500, 18, -64},
	"VI": {500, 18, -64},
	"VN": {500, 14, 108},
	"VU": {500, -15, 166},
	"WF": {500, -13, -177},
	"WS": {500, -13, -172},
	"XK": {500, 42, 20},
	"YE": {500, 15, 48},
	"YT": {500, -12, 45},
	"ZA": {500, -30, 22},
	"ZM": {500, -13, 27},
	"ZW": {500, -19, 29},
}
