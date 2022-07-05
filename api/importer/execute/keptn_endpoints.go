package execute

type StaticKeptnEndpointProvider struct{}

func (_ StaticKeptnEndpointProvider) GetControlPlaneEndpoint() string {
	return "http://shipyard-controller:8080"
}
