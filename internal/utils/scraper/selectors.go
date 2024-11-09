package scraper

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
	"HasStorage":  "#app > div.container--has-footer-d86a9.kt-container > div > main > article > div > div.kt-col-5 > section:nth-child(1) > div.post-page__section--padded > table:nth-child(10) > tbody > tr > td:nth-child(3)",
	"HasParking":  "#app > div.container--has-footer-d86a9.kt-container > div > main > article > div > div.kt-col-5 > section:nth-child(1) > div.post-page__section--padded > table:nth-child(10) > tbody > tr > td.kt-group-row-item.kt-group-row-item__value.kt-body.kt-body--stable.kt-group-row-item--disabled",
	"Images":      "#app > div.container--has-footer-d86a9.kt-container > div > main > article > div > div.kt-col-6.kt-offset-1 > section:nth-child(1) > div > div > div.keen-slider.kt-base-carousel__slides.slides-d6304 > div > figure > div > picture > img",
}

var selectorDivarRent = selector{}

var Selectors = map[string]map[string]selector{
	"Divar": {
		"Buy":  selectorDivarBuy,
		"Rent": selectorDivarRent,
	},
	"Sheypoor": {
		"Buy":  nil,
		"Rent": nil,
	},
}
