package device

type Repository interface {
    findAll(limit int, startingAfter int) (*ResponseMany, error)
    getById(deviceId string) (*ResponseOne, error)
}