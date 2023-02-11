package utils

func MakeRapidQuoteSk(firstShipperQuoteId string) string {
	return "quote#" + firstShipperQuoteId + "#rapid_quote"
}
