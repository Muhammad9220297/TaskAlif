package Senders

import (
	"AlifTack/database"
)

type Client struct {
	Id    int    `gorm:"id"`
	Name  string `gorm:"name"`
	Tel   string `gorm:"tel"`
	Email string `gorm:"email"`
}

type Pakupka struct {
	Id           int     `gorm:"id"`
	Id_client    int     `gorm:"id_client"`
	Tovar        string  `gorm:"tovar"`
	Sena_pokupka float64 `gorm:"sena_pokupka"`
}

type Pakupkas []Pakupka
type Clients []Client

func (p *Pakupkas) ScanTable() error {
	base := database.DB_J()
	base = base.Table("pakupka").Scan(&p)
	return base.Error
}

func (c *Client) GetClients(id int) error {
	base := database.DB_J()
	base = base.Raw("SELECT  * FROM  `client` c INNER JOIN  pakupka p ON c.id=p.id_client AND  c.id=?", id).Scan(&c)
	return base.Error
}

func (p *Pakupka) ScanPakupkaBY_Id_client(id_client int) error {
	base := database.DB_J()
	base = base.Raw("SELECT p.id_client,p.tovar,p.sena_pokupka    FROM  `client` c INNER JOIN  pakupka p ON c.id=p.id_client where  c.id=? ORDER BY p.id_client ", id_client).Scan(&p)
	return base.Error
}
