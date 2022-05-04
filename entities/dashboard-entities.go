package entities

type Model_dashboardwinlose struct {
	Dashboardwinlose_idcompany string `json:"dashboardwinlose_idcompany"`
	Dashboardwinlose_nmcompany string `json:"dashboardwinlose_nmcompany"`
	Dashboardwinlose_periode   string `json:"dashboardwinlose_periode"`
	Dashboardwinlose_winlose   int    `json:"dashboardwinlose_winlose"`
}
type Model_dashboardwinlose_parent struct {
	Dashboardwinlose_nmcompany string      `json:"dashboardwinlose_nmcompany"`
	Dashboardwinlose_detail    interface{} `json:"dashboardwinlose_detail"`
}
type Model_dashboardwinlose_child struct {
	Dashboardwinlose_winlose int `json:"dashboardwinlose_winlose"`
}
type Controller_dashboard struct {
	Master string `json:"master" validate:"required"`
	Year   string `json:"year" validate:"required"`
}
