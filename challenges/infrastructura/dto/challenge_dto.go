package dto

type Challenge struct {
	Id_Challenge        string `json:"id_Challenge"`
	Id_Video            string `json:"id_Video"`
	Name_Challenge      string `json:"name_Challenge"`
	Number_Participants int    `json:"number_Participants"`
}
