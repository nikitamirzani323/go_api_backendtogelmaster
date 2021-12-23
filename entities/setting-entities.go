package entities

type Model_setting struct {
	StartMaintenance string `json:"maintenance_start"`
	EndMaintenance   string `json:"maintenance_end"`
}

type Controller_settinghome struct {
	Master string `json:"master" validate:"required"`
}
type Controller_settingsave struct {
	Maintenance_start string `json:"maintenance_start" validate:"required"`
	Maintenance_end   string `json:"maintenance_end" validate:"required"`
}
