package entities

type Model_domain struct {
	Domain_iddomain int    `json:"domain_iddomain"`
	Domain_name     string `json:"domain_name"`
	Domain_status   string `json:"domain_status"`
	Domain_create   string `json:"domain_create"`
	Domain_update   string `json:"domain_update"`
}

type Controller_domainsave struct {
	Sdata         string `json:"sdata" validate:"required"`
	Master        string `json:"master" validate:"required"`
	Page          string `json:"page" validate:"required"`
	Domain_id     int    `json:"domain_id"`
	Domain_name   string `json:"domain_name" validate:"required"`
	Domain_status string `json:"domain_status" validate:"required"`
}
