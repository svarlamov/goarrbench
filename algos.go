package main

func main() {
	return
}

func isLocalSpotifyTrackId(tId string) bool {
	if len(tId) > 14 && tId[:14] == "spotify:local:" {
		return true
	}
	return false
}

func MakeNewIdsArraySameCap(ids []string) []string {
	validTrackIds := make([]string, len(ids))
	n := 0
	for _, tId := range ids {
		if isLocalSpotifyTrackId(tId) == false {
			validTrackIds[n] = tId
			n++
		}
	}
	return validTrackIds[0:n]
}

func MakeNewIdsArrayDiffCapWithAppend(ids []string) []string {
	validTrackIds := make([]string, 0, len(ids))
	for _, tId := range ids {
		if isLocalSpotifyTrackId(tId) == false {
			validTrackIds = append(validTrackIds, tId)
		}
	}
	return validTrackIds
}

func PluckFromBaseArrayWithCounterAndAppend(ids []string) []string {
	cnt := 0
	l := len(ids)
	for i := 0; i < l; i++ {
		if isLocalSpotifyTrackId(ids[i-cnt]) == true {
			ids = append(ids[:i-cnt], ids[i-cnt+1:]...)
			cnt++
		}
	}
	return ids
}
