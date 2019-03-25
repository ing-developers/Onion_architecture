package dto

// i18n struct for internacionalization of app
type i18n struct {
	//TYPE ACTION / DESCRIPTION
	ErrReadFile          string `json:"err_read_file"`
	ErrConnectDB         string `json:"err_connect_db"`
	ErrNoImplementEngine string `json:"err_no_implement_engine"`
	ErrCreateRegisterDB  string `json:"err_create_register_db"`
	ErrBadEntity         string `json:"err_bad_entity"`
	CrrCreateRegisterDB string `json:"crr_create_register_db"`
}
