package Func

import "strconv"

func Click(pseudo string, Id string, style string) {

	switch style {
	case "like":
		islike, _ := IsLike(pseudo, Id)
		if islike {
			DelLike(pseudo, Id)
			num, _ := strconv.Atoi(Id)
			DecrementLike(num)
		} else {
			num, _ := strconv.Atoi(Id)
			IncrementLike(num)
			AddLike(pseudo, Id)
		}

	case "dislike":
		Isdislike, _ := IsDislike(pseudo, Id)
		if Isdislike {
			DelDislike(pseudo, Id)
			num, _ := strconv.Atoi(Id)
			DecrementDislike(num)
		} else {
			num, _ := strconv.Atoi(Id)
			IncrementDislike(num)
			AddDislike(pseudo, Id)
		}

	case "lier":
		Islier, _ := IsLier(pseudo, Id)
		if Islier {
			DelLier(pseudo, Id)
			num, _ := strconv.Atoi(Id)
			DecrementLier(num)
		} else {
			num, _ := strconv.Atoi(Id)
			IncrementLier(num)
			Addlier(pseudo, Id)
		}
	default:
	}
}
