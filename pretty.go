// Godoc PrettyPrint
func PrettyPrint(target interface{}) string {
	byteJson, err := json.MarshalIndent(target, "", " ")
	if err != nil {
		log.Println(aurora.Red(err))
		return ""
	}

	return string(byteJson)
}
