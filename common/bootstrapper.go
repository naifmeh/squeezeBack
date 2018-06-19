package common

import "log"

func StartUp() {
	/* Init configs structs */
	log.Print("Initializing configuration")
	initConfig()
	/* Init private/public keys */
	log.Print("Initializing keys")
	initKeys()
	/* Init db */
	log.Print("Creating session")
	createDbSession()
	/* Add indexes
	log.Print("Creating indexes")
	addIndexes()*/
}
