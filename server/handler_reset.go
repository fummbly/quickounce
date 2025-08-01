package main

import "net/http"

func (cfg *apiConfig) handlerReset(w http.ResponseWriter, r *http.Request) {
	// only allowing reseting of database in an dev enviroment
	if cfg.platform != "dev" {
		w.WriteHeader(http.StatusForbidden)
		w.Write([]byte("Reset is only allowed in dev environment"))
		return
	}

	// sending the reset query to the database
	err := cfg.db.Reset(r.Context())
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Failed to reset the database: " + err.Error()))
		return
	}

	// respond with an okay
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Databases reset"))
}
