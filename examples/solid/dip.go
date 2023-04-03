package main

import "gorm.io/gorm"

/*
High-level modules should not depend on low-level modules.
Both should depend on abstractions. Abstractions should not depend on details.
Details should depend on abstractions.
*/

// infrastructure layer

// type UserRepository struct {
// 	db *gorm.DB
// }

// func NewUserRepository(db *gorm.DB) *UserRepository {
// 	return &UserRepository{
// 		db: db,
// 	}
// }

// func (r *UserRepository) GetByID(id uint) (*Users, error) {
// 	user := Users{}
// 	err := r.db.Where("id = ?", id).First(&user).Error
// 	if err != nil {
// 		return nil, err
// 	}

// 	return &user, nil
// }

// ///Refactoring

// infrastructure layer

type UserGorm struct {
	// some fields
}

func (g UserGorm) ToUser() *UserDomain {
	return &UserDomain{
		// some fields
	}
}

type UserDatabaseRepository struct {
	db *gorm.DB
}

var _ UserRepository = &UserDatabaseRepository{}

/*
type UserRedisRepository struct {

}
type UserCassandraRepository struct {
}
*/

func NewUserDatabaseRepository(db *gorm.DB) UserRepository {
	return &UserDatabaseRepository{
		db: db,
	}
}

func (r *UserDatabaseRepository) GetByID(id uint) (*UserDomain, error) {
	user := UserGorm{}
	err := r.db.Where("id = ?", id).First(&user).Error
	if err != nil {
		return nil, err
	}

	return user.ToUser(), nil
}

// domain layer

type UserDomain struct {
	// some fields
}

type UserRepository interface {
	GetByID(id uint) (*UserDomain, error)
}

// application layer

type EmailService struct {
	repository UserRepository
	// some email sender
}

func NewEmailService(repository UserRepository) *EmailService {
	return &EmailService{
		repository: repository,
	}
}

func (s *EmailService) SendRegistrationEmail(userID uint) error {
	_, err := s.repository.GetByID(userID)
	if err != nil {
		return err
	}
	// send email
	return nil
}
