package common

func StartUp() {
	/* Init configs structs */
	initConfig()
	/* Init private/public keys */
	initKeys()
	/* Init db */
	createDbSession()
	/* Add indexes */
	addIndexes()
}
