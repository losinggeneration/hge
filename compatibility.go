package hge

func (h *HGE) Initiate() error {
	return h.Initialize()
}

func (h *HGE) Start() error {
	return h.Run()
}
