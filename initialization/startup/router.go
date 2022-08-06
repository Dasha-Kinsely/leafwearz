package startup

import (
	"leafwearz/routers"
)

type RouterGroup struct {
	// accessible by all
	routers.ProductRoutes
	routers.InventoryRoutes
	routers.ServiceRoutes
	routers.AvailabilityRoutes
	routers.PubInfoRoutes
	routers.PubPromoRoutes
	routers.AuthRoutes
	// requires authorization
	routers.UserProfileRoutes
	routers.PaymentGatewayRoutes
	routers.C2MCommunicationRoutes
	routers.ConsumerIntelligenceRoutes
	// requires admin privileges
	routers.AdminRoutes
	routers.RecommendationSettingRoutes
	routers.StatRoutes
}

var Router *RouterGroup

func InitRouterGroup() {
	publicGroup := DefaultServer.Group("public")
	{
		Router.AuthRoutes.InitRoutes(publicGroup)
		Router.PubInfoRoutes.InitRoutes(publicGroup)
		Router.PubPromoRoutes.InitRoutes(publicGroup)
		Router.ProductRoutes.InitRoutes(publicGroup)
		Router.InventoryRoutes.InitRoutes(publicGroup)
		Router.ServiceRoutes.InitRoutes(publicGroup)
		Router.AvailabilityRoutes.InitRoutes(publicGroup)
	}
	authorizedGroup := DefaultServer.Group("authorized")
	{
		Router.PaymentGatewayRoutes.InitRoutes(authorizedGroup)
		Router.UserProfileRoutes.InitRoutes(authorizedGroup)
		Router.C2MCommunicationRoutes.InitRoutes(authorizedGroup)
		Router.ConsumerIntelligenceRoutes.InitRoutes(authorizedGroup)
	}
	adminGroup := DefaultServer.Group("admin")
	{
		Router.AdminRoutes.InitRoutes(adminGroup)
		Router.StatRoutes.InitRoutes(adminGroup)
		Router.RecommendationSettingRoutes.InitRoutes(adminGroup)
	}
	
}