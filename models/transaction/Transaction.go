package transaction

type Transaction struct {
	UserId         string `gorm:"not null"`
	TrasactionHash string `gorm:"primary_key"`
}
