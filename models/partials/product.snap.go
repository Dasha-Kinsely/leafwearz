package partials

type ProductSnap struct {
	ImageID uint `json:"image_id"`
	ImageAddress string `json:"image_address"`
	Title string `json:"title"`
	RenderMethod string `json:"render_method"` // dictates what other pieces of info are needed
}