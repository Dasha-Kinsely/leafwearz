package routers

import "github.com/gin-gonic/gin"

type RouterGroup struct {
	ProductRouterGroup
	ServiceRouterGroup
	AdminRouterGroup
	ShareResponsibilityRouterGroup
	ProfilerRouterGroup
}

type ProductRouterGroup struct {
	ProductRoutes
	InventoryRoutes
}

type ServiceRouterGroup struct {
	ServiceRoutes
	AvailabilityRoutes
}

type AdminRouterGroup struct {
	AdminRoutes
	PubInfoRoutes
	RecommendationSettingRoutes
	StatRoutes
}

type ShareResponsibilityRouterGroup struct {
	PaymentGatewayRoutes
	PromoRoutes
	C2MCommunicationRoutes
	AuthRoutes
}

type ProfilerRouterGroup struct {
	UserRoutes
	ConsumerIntelligenceRoutes
}

func InitRouterGroup() *gin.Engine {
	engine := gin.Default()
	return engine
}