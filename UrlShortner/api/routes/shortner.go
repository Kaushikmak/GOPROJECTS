package routes

import "time"



type request struct{
    URL             string          `json:"url"`
    CustomShortner  string          `json:"customshortner"`
    Expiry          time.Duration   `json:"expiry"`

}

// rate limiting
type response struct{
    URL             string          `json:"url"`
    CustomShortner  string          `json:"customshort"`    
    Expiry          time.Duration   `json:"expiry"`
    XRateRemaining  int             `json:"rate_limit"`
    XRateLimitReset time.Duration   `json:"rate_limit_reset"`
}


