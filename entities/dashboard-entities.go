package entities

type Model_dashboardwinlose_parent struct {
	Dashboardwinlose_nmcompany string      `json:"dashboardwinlose_nmcompany"`
	Dashboardwinlose_detail    interface{} `json:"dashboardwinlose_detail"`
}
type Model_dashboardwinlose_child struct {
	Dashboardwinlose_winlose int `json:"dashboardwinlose_winlose"`
}
type Model_companypasaran_parent struct {
	Dashboardcompanypasaran_nmpasaran string      `json:"dashboardcompanypasaran_nmpasaran"`
	Dashboardcompanypasaran_detail    interface{} `json:"dashboardcompanypasaran_detail"`
}
type Model_companypasaran_child struct {
	Dashboardcompanypasaran_winlose int `json:"dashboardcompanypasaran_winlose"`
}
type Controller_dashboard struct {
	Master string `json:"master" validate:"required"`
	Year   string `json:"year" validate:"required"`
}
type Controller_dashboardcompanypasaran struct {
	Master  string `json:"master" validate:"required"`
	Company string `json:"company" validate:"required"`
	Year    string `json:"year" validate:"required"`
}
