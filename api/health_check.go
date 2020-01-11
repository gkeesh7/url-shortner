package api

import (
	"net/http"
	"strings"
	"sync"
)

var (
	serviceState    = "BIR"
	serviceStateOOR = "OOR"
	serviceStateBIR = "BIR"
	mutex           = &sync.Mutex{}
)

// HealthCheck defines normal health check call.
func HealthCheck(resp http.ResponseWriter, req *http.Request) {
	if strings.Compare(serviceState, serviceStateBIR) == 0 {
		resp.WriteHeader(http.StatusOK)
		_, _ = resp.Write([]byte(`health check is fine`))
	} else {
		resp.WriteHeader(http.StatusInternalServerError)
		_, _ = resp.Write([]byte(`Sorry I am not healthy.`))
	}
}

// OutOfRotation takes service out of rotation. Service won't serve further request
func OutOfRotation(resp http.ResponseWriter, req *http.Request) {
	mutex.Lock()
	serviceState = serviceStateOOR
	mutex.Unlock()
	resp.WriteHeader(http.StatusOK)
	_, _ = resp.Write([]byte(`Service Is Out Of Rotation`))
}

// BackInRotation gets service in rotation. Service is now healthy to serve request
func BackInRotation(resp http.ResponseWriter, req *http.Request) {
	mutex.Lock()
	serviceState = serviceStateBIR
	mutex.Unlock()
	resp.WriteHeader(http.StatusOK)
	_, _ = resp.Write([]byte(`Service Is In Rotation`))
}
