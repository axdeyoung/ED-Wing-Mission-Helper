package journalparser

import (
	//"json"
)


/*
	Resets file pointer and returns all relevant data since previous call
*/
func DumpTradeJson() string {
	return "{ \"placeholder\": \"this is where we will dump all the trade-relevant data from the Elite Dangerous pilot journal\" }"
}

/*
	Moves file pointer down and returns new relevant data since previous call
*/
func UpdateTradeJson() string {
	return "{ \"placeholder\": \"this is where we will put updates to the trade-relevant data from the Elite Dangerous pilot journal\" }"
}