package parser

type selector map[string]string

var selectorDivarBuy = selector{
	"Title":       "#app > div.container--has-footer-d86a9.kt-container > div > main > article > div > div.kt-col-5 > section:nth-child(1) > div.kt-page-title > div > h1",
	"Description": "#app > div.container--has-footer-d86a9.kt-container > div > main > article > div > div.kt-col-5 > section.post-page__section--padded > div > div.kt-base-row.kt-base-row--large.kt-description-row > div > p",
	"BuyPrice":    "#app > div.container--has-footer-d86a9.kt-container > div > main > article > div > div.kt-col-5 > section:nth-child(1) > div.post-page__section--padded > div:nth-child(3) > div.kt-base-row__end.kt-unexpandable-row__value-box > p",
	"Area":        "#app > div.container--has-footer-d86a9.kt-container > div > main > article > div > div.kt-col-5 > section:nth-child(1) > div.post-page__section--padded > table:nth-child(1) > tbody > tr > td:nth-child(1)",
	"Rooms":       "#app > div.container--has-footer-d86a9.kt-container > div > main > article > div > div.kt-col-5 > section:nth-child(1) > div.post-page__section--padded > table:nth-child(1) > tbody > tr > td:nth-child(3)",
	"City":        "#app > div.container--has-footer-d86a9.kt-container > div > main > article > div > div.kt-col-5 > section:nth-child(1) > div.kt-page-title > div > div",
	"District":    "#app > div.container--has-footer-d86a9.kt-container > div > main > article > div > div.kt-col-5 > section:nth-child(1) > div.kt-page-title > div > div",
	"BuildYear":   "#app > div.container--has-footer-d86a9.kt-container > div > main > article > div > div.kt-col-5 > section:nth-child(1) > div.post-page__section--padded > table:nth-child(1) > tbody > tr > td:nth-child(2)",
	"Floor":       "#app > div.container--has-footer-d86a9.kt-container > div > main > article > div > div.kt-col-5 > section:nth-child(1) > div.post-page__section--padded > div:nth-child(7) > div.kt-base-row__end.kt-unexpandable-row__value-box > p",
	"HasElevator": "#app > div.container--has-footer-d86a9.kt-container > div > main > article > div > div.kt-col-5 > section:nth-child(1) > div.post-page__section--padded > table:nth-child(10) > tbody > tr > td:nth-child(1)",
	"HasParking":  "#app > div.container--has-footer-d86a9.kt-container > div > main > article > div > div.kt-col-5 > section:nth-child(1) > div.post-page__section--padded > table:nth-child(10) > tbody > tr > td:nth-child(2)",
	"HasStorage":  "#app > div.container--has-footer-d86a9.kt-container > div > main > article > div > div.kt-col-5 > section:nth-child(1) > div.post-page__section--padded > table:nth-child(10) > tbody > tr > td:nth-child(3)",
	"Images":      "div.keen-slider__slide:nth-child(2) > figure:nth-child(1) > div:nth-child(1) > picture:nth-child(1) > img:nth-child(1)",
}

var selectorSheypoorBuy = selector{
	"Title":       "#listing-title",
	"Description": "#zXihl > div > div._6WBnL.YPVjY._8zJiw._6ySN8 > div > div.Pc66i > div > div.FyLl0 > div.VNOCj.p3lU- > div.MQJ5W",
	"BuyPrice":    "#zXihl > div > div._6WBnL.YPVjY._8zJiw._6ySN8 > div > div.Pc66i > div > div.Xg4Vc > div.pKgFC > div.tOq3m > span > strong",
	"Area":        "#zXihl > div > div._6WBnL.YPVjY._8zJiw._6ySN8 > div > div.Pc66i > div > div.bWPjU > div:nth-child(1) > div:nth-child(1) > p._874-x",
	"Rooms":       "#zXihl > div > div._6WBnL.YPVjY._8zJiw._6ySN8 > div > div.Pc66i > div > div.bWPjU > div:nth-child(1) > div:nth-child(3) > p._874-x",
	"City":        "#zXihl > div > div._6WBnL.YPVjY._8zJiw._6ySN8 > div > div.Pc66i > div > div.Xg4Vc > div.pKgFC > div._3oBho > span:nth-child(2)",
	"District":    "#zXihl > div > div._6WBnL.YPVjY._8zJiw._6ySN8 > div > div.Pc66i > div > div.Xg4Vc > div.pKgFC > div._3oBho > span:nth-child(2)",
	"BuildYear":   "#zXihl > div > div._6WBnL.YPVjY._8zJiw._6ySN8 > div > div.Pc66i > div > div.bWPjU > div:nth-child(2) > div:nth-child(3) > p._874-x",
	"Floor":       "",
	"HasElevator": "#zXihl > div > div._6WBnL.YPVjY._8zJiw._6ySN8 > div > div.Pc66i > div > div.bWPjU > div:nth-child(2) > div:nth-child(2) > p._874-x",
	"HasParking":  "#zXihl > div > div._6WBnL.YPVjY._8zJiw._6ySN8 > div > div.Pc66i > div > div.bWPjU > div:nth-child(1) > div:nth-child(4) > p._874-x",
	"HasStorage":  "#zXihl > div > div._6WBnL.YPVjY._8zJiw._6ySN8 > div > div.Pc66i > div > div.bWPjU > div:nth-child(2) > div:nth-child(1) > p._874-x",
	"Images":      "#zXihl > div > div.FPa3Y.YPVjY.B3sgi > div > div > div.swiper-wrapper > div:nth-child(6) > img",
}

var selectorDivarRent = selector{
	"Title":       "#app > div.container--has-footer-d86a9.kt-container > div > main > article > div > div.kt-col-5 > section:nth-child(1) > div.kt-page-title > div > h1",
	"Description": "#app > div.container--has-footer-d86a9.kt-container > div > main > article > div > div.kt-col-5 > section.post-page__section--padded > div > div.kt-base-row.kt-base-row--large.kt-description-row > div > p",
	"Area":        "#app > div.container--has-footer-d86a9.kt-container > div > main > article > div > div.kt-col-5 > section:nth-child(1) > div.post-page__section--padded > table:nth-child(1) > tbody > tr > td:nth-child(1)",
	"Rooms":       "#app > div.container--has-footer-d86a9.kt-container > div > main > article > div > div.kt-col-5 > section:nth-child(1) > div.post-page__section--padded > table:nth-child(1) > tbody > tr > td:nth-child(3)",
	"City":        "#app > div.container--has-footer-d86a9.kt-container > div > main > article > div > div.kt-col-5 > section:nth-child(1) > div.kt-page-title > div > div",
	"District":    "#app > div.container--has-footer-d86a9.kt-container > div > main > article > div > div.kt-col-5 > section:nth-child(1) > div.kt-page-title > div > div",
	"BuildYear":   "#app > div.container--has-footer-d86a9.kt-container > div > main > article > div > div.kt-col-5 > section:nth-child(1) > div.post-page__section--padded > table:nth-child(1) > tbody > tr > td:nth-child(2)",
	"Convertable": "#app > div.container--has-footer-d86a9.kt-container > div > main > article > div > div.kt-col-5 > section:nth-child(1) > div.post-page__section--padded > div.kt-base-row.kt-base-row--large.kt-base-row--has-icon.kt-feature-row > div > p",
	"Images":      "div.keen-slider__slide:nth-child(2) > figure:nth-child(1) > div:nth-child(1) > picture:nth-child(1) > img:nth-child(1)",
}

var seletorDivarRentTypeConvertable = selector{
	"HasElevator": "#app > div.container--has-footer-d86a9.kt-container > div > main > article > div > div.kt-col-5 > section:nth-child(1) > div.post-page__section--padded > table:nth-child(9) > tbody > tr > td:nth-child(1)",
	"HasParking":  "#app > div.container--has-footer-d86a9.kt-container > div > main > article > div > div.kt-col-5 > section:nth-child(1) > div.post-page__section--padded > table:nth-child(9) > tbody > tr > td:nth-child(2)",
	"HasStorage":  "#app > div.container--has-footer-d86a9.kt-container > div > main > article > div > div.kt-col-5 > section:nth-child(1) > div.post-page__section--padded > table:nth-child(9) > tbody > tr > td:nth-child(3)",
	"Floor":       "#app > div.container--has-footer-d86a9.kt-container > div > main > article > div > div.kt-col-5 > section:nth-child(1) > div.post-page__section--padded > div.kt-base-row.kt-base-row--large.kt-unexpandable-row > div.kt-base-row__end.kt-unexpandable-row__value-box > p",
	//"RentPrice":    "",
	"RahnPriceMax":  "#app > div.container--has-footer-d86a9.kt-container > div > main > article > div > div.kt-col-5 > section:nth-child(1) > div.post-page__section--padded > div.convert-slider > div:nth-child(3) > div.convert-slider__info-right.kt-col-6 > span",
	"RahnPriceMin":  "#app > div.container--has-footer-d86a9.kt-container > div > main > article > div > div.kt-col-5 > section:nth-child(1) > div.post-page__section--padded > div.convert-slider > div:nth-child(3) > div.convert-slider__info-left.kt-col-6 > span",
	"RentPriceMax":  "#app > div.container--has-footer-d86a9.kt-container > div > main > article > div > div.kt-col-5 > section:nth-child(1) > div.post-page__section--padded > div.convert-slider > div:nth-child(4) > div.convert-slider__info-left.kt-col-6 > span",
	"RentPriceMin":  "#app > div.container--has-footer-d86a9.kt-container > div > main > article > div > div.kt-col-5 > section:nth-child(1) > div.post-page__section--padded > div.convert-slider > div:nth-child(4) > div.convert-slider__info-right.kt-col-6",
	"RentPriceMin2": "#app > div.container--has-footer-d86a9.kt-container > div > main > article > div > div.kt-col-5 > section:nth-child(1) > div.post-page__section--padded > div.convert-slider > div:nth-child(4) > div.convert-slider__info-right.kt-col-6 > span",
}
var seletorDivarRentTypeNonConvertable = selector{
	"HasElevator":             "#app > div.container--has-footer-d86a9.kt-container > div > main > article > div > div.kt-col-5 > section:nth-child(1) > div.post-page__section--padded > table:nth-child(12) > tbody > tr > td:nth-child(1)",
	"HasParking":              "#app > div.container--has-footer-d86a9.kt-container > div > main > article > div > div.kt-col-5 > section:nth-child(1) > div.post-page__section--padded > table:nth-child(12) > tbody > tr > td:nth-child(2)",
	"HasStorage":              "#app > div.container--has-footer-d86a9.kt-container > div > main > article > div > div.kt-col-5 > section:nth-child(1) > div.post-page__section--padded > table:nth-child(12) > tbody > tr > td:nth-child(3)",
	"Floor":                   "#app > div.container--has-footer-d86a9.kt-container > div > main > article > div > div.kt-col-5 > section:nth-child(1) > div.post-page__section--padded > div:nth-child(9) > div.kt-base-row__end.kt-unexpandable-row__value-box > p",
	"RentPriceNonConvertable": "#app > div.container--has-footer-d86a9.kt-container > div > main > article > div > div.kt-col-5 > section:nth-child(1) > div.post-page__section--padded > div:nth-child(5) > div.kt-base-row__end.kt-unexpandable-row__value-box > p",
	"RahnPriceNonConvertable": "#app > div.container--has-footer-d86a9.kt-container > div > main > article > div > div.kt-col-5 > section:nth-child(1) > div.post-page__section--padded > div:nth-child(3) > div.kt-base-row__end.kt-unexpandable-row__value-box > p",
}

var Selectors = map[string]map[string]selector{
	"Divar": {
		"Buy":                    selectorDivarBuy,
		"Rent":                   selectorDivarRent,
		"RentTypeConvertable":    seletorDivarRentTypeConvertable,
		"RentTypeNonConvertable": seletorDivarRentTypeNonConvertable,
	},
	"Sheypoor": {
		"Buy":  selectorSheypoorBuy,
		"Rent": nil,
	},
}
