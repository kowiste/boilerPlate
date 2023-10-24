package other

func (s *Stuff) InjectAPI() {
	api := s.controller.GetAPI().Group("api")
	{
		stuff := api.Group("other")
		{
			stuff.POST("create", s.Create)
			stuff.GET("list", s.List)
			stuff.GET("find/:id", s.Find)
			stuff.PUT("update/:id", s.Update)
			stuff.DELETE("delete/:id", s.Delete)
		}
	}
}
