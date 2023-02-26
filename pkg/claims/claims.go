package claims

import (
	"os"
	"strconv"
	"time"

	log "github.com/sirupsen/logrus"
)

type CustomClaims struct {
	UUID       string    `json:"uuid"`
	Email      string    `json:"email"`
	Expiration time.Time `json:"expiryTime"`
}

func New(claimsValue CustomClaims) CustomClaims {
	pasetoExpirationInHours, ok := os.LookupEnv("PASETO_EXPIRATION_IN_HOURS")
	pasetoExpirationInHoursInt := time.Duration(24)

	if ok {
		res, err := strconv.Atoi(pasetoExpirationInHours)
		if err != nil {
			log.WithFields(log.Fields{
				"err": err,
			}).Error("failed to bind")
		} else {
			pasetoExpirationInHoursInt = time.Duration(res)
		}
	}
	pasetoExpirationHours := pasetoExpirationInHoursInt * time.Hour
	expiration := time.Now().Add(pasetoExpirationHours)

	return CustomClaims{
		claimsValue.Email,
		claimsValue.UUID,
		expiration,
	}
}
