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
	"HasElevator":  "",
	"HasParking":   "",
	"HasStorage":   "",
	"Floor":        "",
	"RentPrice":    "",
	"RentPriceMin": "",
	"RentPriceMax": "",
	"RahnPriceMin": "",
	"RahnPriceMax": "",
}

var Selectors = map[string]map[string]selector{
	"Divar": {
		"Buy":                    selectorDivarBuy,
		"Rent":                   selectorDivarRent,
		"RentTypeConvertable":    seletorDivarRentTypeConvertable,
		"RentTypeNonConvertable": seletorDivarRentTypeNonConvertable,
	},
	"Sheypoor": {
		"Buy":  nil,
		"Rent": nil,
	},
}
