package app

import "auditfox/app/routes"

func StartUp() {
	r := routes.SetupRouters()
	r.Run()
}
