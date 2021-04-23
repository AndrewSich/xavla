package system

import "net/http"

// This a Route configuration
const staticRouteTarget = "/static/"
const staticFileTarget = "public"




// This a routing base for all handler route
func routingBase() {
  HandlerStatic(staticRouteTarget, staticFileTarget)
  http.HandleFunc("/", HandlerIndex)
}