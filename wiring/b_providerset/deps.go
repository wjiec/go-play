package main

import "github.com/google/wire"

type MySQLConnectString string

type MySQLConfig struct {
	dsn MySQLConnectString
}

func NewMySQLConfig(dsn MySQLConnectString) (*MySQLConfig, error) {
	return &MySQLConfig{dsn: dsn}, nil
}

type MySQL struct {
	cfg *MySQLConfig
}

func NewMySQL(cfg *MySQLConfig) (*MySQL, error) {
	return &MySQL{cfg: cfg}, nil
}

// MySQLProviderSet 由于 MySQLConfig 和 MySQL 经常一起出现
// 所以可以直接用 ProviderSet 把这些相关性非常强的依赖打包在一起
// 后再统一进行依赖注入
var MySQLProviderSet = wire.NewSet(
	NewMySQLConfig,
	NewMySQL,
)

type UserRepository struct {
	db *MySQL
}

func NewUserRepository(db *MySQL) (*UserRepository, error) {
	return &UserRepository{db: db}, nil
}

type ArticleRepository struct {
	db *MySQL
}

func NewArticleRepository(db *MySQL) (*ArticleRepository, error) {
	return &ArticleRepository{db: db}, nil
}

type BlogService struct {
	ur *UserRepository
	ar *ArticleRepository
}

func NewBlogService(ur *UserRepository, ar *ArticleRepository) (*BlogService, error) {
	return &BlogService{ur: ur, ar: ar}, nil
}
