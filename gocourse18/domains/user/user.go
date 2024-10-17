package user

type User struct {
	ID         uint   `json:"id"`
	Name       string `json:"Імʼя"`
	Surname    string `json:"Прізвище"`
	Email      string `json:"Email"`
	Age        int64  `json:"Вік"`
	Sex        string `json:"Стать"`
	City       string `json:"Місто проживання"`
	TaxiCount  int64  `json:"Кількість поїздок таксі" db:"taxi_count"`
	Profession string `json:"Професія"`
}
