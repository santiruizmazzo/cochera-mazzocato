package domain

type EntityID uint32

func NewEntityID(rawValue int) (EntityID, error) {
	return EntityID(rawValue), nil
}
