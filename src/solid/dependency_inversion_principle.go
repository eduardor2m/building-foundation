package solid

import "fmt"

type UserRepository interface {
    Save(user User)
}

type MemoryUserRepository struct {
    users []User
}

func (repo *MemoryUserRepository) Save(user User) {
    repo.users = append(repo.users, user)
    fmt.Printf("Salvando usuário %s no repositório de memória.\n", user.Name)
}

type User struct {
    Name string
}

type UserService struct {
    userRepository UserRepository
}

func (service *UserService) CreateUser(name string) {
    user := User{Name: name}
    service.userRepository.Save(user)
}

func DIP() {
    userRepository := &MemoryUserRepository{}
    userService := &UserService{userRepository: userRepository}

    userService.CreateUser("Eduardo")
}
