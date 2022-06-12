module sametemiroglu.com/user-management-system

go 1.17

replace sametemiroglu.com/dataloaders => ../dataloaders

replace sametemiroglu.com/handlers => ../handlers

require sametemiroglu.com/handlers v0.0.0-00010101000000-000000000000

require (
	sametemiroglu.com/dataloaders v0.0.0-00010101000000-000000000000 // indirect
	sametemiroglu.com/models v0.0.0-00010101000000-000000000000 // indirect
	sametemiroglu.com/utils v0.0.0-00010101000000-000000000000 // indirect
)

replace sametemiroglu.com/models => ../models

replace sametemiroglu.com/utils => ../utils
