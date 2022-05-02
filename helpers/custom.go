package helpers

import "github.com/nleeper/goment"

func GetEndRangeDate(month string) (string, string, string, string) {
	tglnow, _ := goment.New()
	startmonthyear := ""
	endmonthyear := ""
	end := ""
	namemonth := ""
	switch month {
	case "01":
		namemonth = "JANUARY"
		end = "31"
		startmonthyear = tglnow.Format("YYYY") + "-01-01"
		endmonthyear = tglnow.Format("YYYY") + "-01-" + end
	case "02":
		namemonth = "FEBUARY"
		end = "28"
		startmonthyear = tglnow.Format("YYYY") + "-02-01"
		endmonthyear = tglnow.Format("YYYY") + "-02-" + end
	case "03":
		namemonth = "MARET"
		end = "31"
		startmonthyear = tglnow.Format("YYYY") + "-03-01"
		endmonthyear = tglnow.Format("YYYY") + "-03-" + end
	case "04":
		namemonth = "APRIL"
		end = "30"
		startmonthyear = tglnow.Format("YYYY") + "-04-01"
		endmonthyear = tglnow.Format("YYYY") + "-04-" + end
	case "05":
		namemonth = "MAY"
		end = "31"
		startmonthyear = tglnow.Format("YYYY") + "-05-01"
		endmonthyear = tglnow.Format("YYYY") + "-05-" + end
	case "06":
		namemonth = "JUNE"
		end = "30"
		startmonthyear = tglnow.Format("YYYY") + "-06-01"
		endmonthyear = tglnow.Format("YYYY") + "-06-" + end
	case "07":
		namemonth = "JULY"
		end = "31"
		startmonthyear = tglnow.Format("YYYY") + "-07-01"
		endmonthyear = tglnow.Format("YYYY") + "-07-" + end
	case "08":
		namemonth = "AUGUSTUS"
		end = "31"
		startmonthyear = tglnow.Format("YYYY") + "-08-01"
		endmonthyear = tglnow.Format("YYYY") + "-08-" + end
	case "09":
		namemonth = "SEPTEMBER"
		end = "30"
		startmonthyear = tglnow.Format("YYYY") + "-09-01"
		endmonthyear = tglnow.Format("YYYY") + "-09-" + end
	case "10":
		namemonth = "OCTOBER"
		end = "31"
		startmonthyear = tglnow.Format("YYYY") + "-10-01"
		endmonthyear = tglnow.Format("YYYY") + "-10-" + end
	case "11":
		namemonth = "NOVEMBER"
		end = "30"
		startmonthyear = tglnow.Format("YYYY") + "-11-01"
		endmonthyear = tglnow.Format("YYYY") + "-11-" + end
	case "12":
		namemonth = "DECEMBER"
		end = "31"
		startmonthyear = tglnow.Format("YYYY") + "-12-01"
		endmonthyear = tglnow.Format("YYYY") + "-12-" + end
	}
	return end, startmonthyear, endmonthyear, namemonth
}
func GetEndRangeDate_year(month, year string) (string, string, string, string) {
	startmonthyear := ""
	endmonthyear := ""
	end := ""
	namemonth := ""
	switch month {
	case "01":
		namemonth = "JANUARY"
		end = "31"
		startmonthyear = year + "-01-01"
		endmonthyear = year + "-01-" + end
	case "02":
		namemonth = "FEBUARY"
		end = "28"
		startmonthyear = year + "-02-01"
		endmonthyear = year + "-02-" + end
	case "03":
		namemonth = "MARET"
		end = "31"
		startmonthyear = year + "-03-01"
		endmonthyear = year + "-03-" + end
	case "04":
		namemonth = "APRIL"
		end = "30"
		startmonthyear = year + "-04-01"
		endmonthyear = year + "-04-" + end
	case "05":
		namemonth = "MAY"
		end = "31"
		startmonthyear = year + "-05-01"
		endmonthyear = year + "-05-" + end
	case "06":
		namemonth = "JUNE"
		end = "30"
		startmonthyear = year + "-06-01"
		endmonthyear = year + "-06-" + end
	case "07":
		namemonth = "JULY"
		end = "31"
		startmonthyear = year + "-07-01"
		endmonthyear = year + "-07-" + end
	case "08":
		namemonth = "AUGUSTUS"
		end = "31"
		startmonthyear = year + "-08-01"
		endmonthyear = year + "-08-" + end
	case "09":
		namemonth = "SEPTEMBER"
		end = "30"
		startmonthyear = year + "-09-01"
		endmonthyear = year + "-09-" + end
	case "10":
		namemonth = "OCTOBER"
		end = "31"
		startmonthyear = year + "-10-01"
		endmonthyear = year + "-10-" + end
	case "11":
		namemonth = "NOVEMBER"
		end = "30"
		startmonthyear = year + "-11-01"
		endmonthyear = year + "-11-" + end
	case "12":
		namemonth = "DECEMBER"
		end = "31"
		startmonthyear = year + "-12-01"
		endmonthyear = year + "-12-" + end
	}
	return end, startmonthyear, endmonthyear, namemonth
}
