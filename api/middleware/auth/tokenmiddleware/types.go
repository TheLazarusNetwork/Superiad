package tokenmiddleware

type GenericAuthHeaders struct {
	Authorization string `binding:"required"`
}
