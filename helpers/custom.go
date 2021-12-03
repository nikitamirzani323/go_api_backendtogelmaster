package helpers

import "github.com/nleeper/goment"

func GetEndRangeDate(month string) (string, string, string) {
	tglnow, _ := goment.New()
	startmonthyear := ""
	endmonthyear := ""
	end := ""
	switch month {
	case "01":
		end = "31"
		startmonthyear = tglnow.Format("YYYY") + "-01-01"
		endmonthyear = tglnow.Format("YYYY") + "-01-" + end
	case "02":
		end = "28"
		startmonthyear = tglnow.Format("YYYY") + "-02-01"
		endmonthyear = tglnow.Format("YYYY") + "-02-" + end
	case "03":
		end = "31"
		startmonthyear = tglnow.Format("YYYY") + "-03-01"
		endmonthyear = tglnow.Format("YYYY") + "-03-" + end
	case "04":
		end = "30"
		startmonthyear = tglnow.Format("YYYY") + "-04-01"
		endmonthyear = tglnow.Format("YYYY") + "-04-" + end
	case "05":
		end = "31"
		startmonthyear = tglnow.Format("YYYY") + "-05-01"
		endmonthyear = tglnow.Format("YYYY") + "-05-" + end
	case "06":
		end = "30"
		startmonthyear = tglnow.Format("YYYY") + "-06-01"
		endmonthyear = tglnow.Format("YYYY") + "-06-" + end
	case "07":
		end = "31"
		startmonthyear = tglnow.Format("YYYY") + "-07-01"
		endmonthyear = tglnow.Format("YYYY") + "-07-" + end
	case "08":
		end = "31"
		startmonthyear = tglnow.Format("YYYY") + "-08-01"
		endmonthyear = tglnow.Format("YYYY") + "-08-" + end
	case "09":
		end = "30"
		startmonthyear = tglnow.Format("YYYY") + "-09-01"
		endmonthyear = tglnow.Format("YYYY") + "-09-" + end
	case "10":
		end = "31"
		startmonthyear = tglnow.Format("YYYY") + "-10-01"
		endmonthyear = tglnow.Format("YYYY") + "-10-" + end
	case "11":
		end = "30"
		startmonthyear = tglnow.Format("YYYY") + "-11-01"
		endmonthyear = tglnow.Format("YYYY") + "-11-" + end
	case "12":
		end = "31"
		startmonthyear = tglnow.Format("YYYY") + "-12-01"
		endmonthyear = tglnow.Format("YYYY") + "-12-" + end
	}
	return end, startmonthyear, endmonthyear
}
