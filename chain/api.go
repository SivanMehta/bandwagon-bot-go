package chain

// concatenate some strings
func join(strs ...string) string {
	var ret string
	for _, str := range strs {
		ret += str
	}
	return ret
}

//
// TweetFromTrend will try and generate a tweet from the given trend
//
func TweetFromTrend(trend string) string {
	if generator, ok := currentBandwagons[trend]; ok {
		return join(generator(), "\n")
	}
	return join("\"", trend, "\" is a not a currently available bandwagon\n")
}
