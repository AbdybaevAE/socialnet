package pgtestcontainer

import (
	"database/sql"
	_ "database/sql"
	"fmt"
	"github.com/abdybaevae/socialnet/migrations"
	entities2 "github.com/abdybaevae/socialnet/pkg/entities"
	"github.com/docker/go-connections/nat"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	"github.com/golang-migrate/migrate/v4/source/iofs"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/stretchr/testify/suite"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/wait"
	"log"
	"time"
)

import (
	"context"
)

type PostgresqlSuite struct {
	suite.Suite
	DB                  *sqlx.DB
	shutdown            func(ctx context.Context) error
	AlreadyExistedLangs []entities2.LanguageEntity
	ctx                 context.Context
	DbUrl               string
	migrator            *migrate.Migrate
}

func (s *PostgresqlSuite) SetupSuite() {
	// up container
	s.ctx = context.Background()
	s.DB, s.shutdown = s.startDb()
}

func (s *PostgresqlSuite) TearDownSuite() {
	// down container
	s.T().Log("Shutting down DB container")
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()
	if err := s.shutdown(ctx); err != nil {
		s.T().Logf("failed to shutdown container %d", err)
	}
}

func (s *PostgresqlSuite) SetupTest() {
	s.Require().NoError(s.migrator.Up())
	s.Require().NoError(s.DB.SelectContext(s.ctx, &s.AlreadyExistedLangs, "select * from lang"))
}
func (s *PostgresqlSuite) TearDownTest() {
	s.Require().NoError(s.migrator.Down())
}
func (s *PostgresqlSuite) startDb() (*sqlx.DB, func(ctx context.Context) error) {
	const dbName = "socialnet"
	const dbUser = "postgres"
	const dbPassword = "password"
	const port = "5432/tcp"
	var env = map[string]string{
		"POSTGRES_PASSWORD": dbPassword,
		"POSTGRES_USER":     dbUser,
		"POSTGRES_DB":       dbName,
	}
	dbURL := func(port nat.Port) string {
		return fmt.Sprintf("postgres://postgres:password@localhost:%s/%s?sslmode=disable", port.Port(), dbName)
	}

	request := testcontainers.GenericContainerRequest{
		ContainerRequest: testcontainers.ContainerRequest{
			Image:        "postgres:latest",
			ExposedPorts: []string{port},
			Cmd:          []string{"postgres", "-c", "fsync=off"},
			Env:          env,
			WaitingFor:   wait.ForSQL(port, "postgres", dbURL).Timeout(time.Second * 5),
		},
		Started: true,
	}

	container, err := testcontainers.GenericContainer(context.Background(), request)
	s.Require().NoError(err)

	mappedPort, err := container.MappedPort(context.Background(), port)
	s.Require().NoError(err)

	s.DbUrl = fmt.Sprintf("postgres://%s:%s@localhost:%s/%s?sslmode=disable", dbUser, dbPassword, mappedPort.Port(), dbName)
	s.Require().NoError(err)

	s.migrator = s.createMigrator()

	dbConnection, err := sqlx.ConnectContext(context.Background(), "postgres", s.DbUrl)
	s.Require().NoError(err)

	return dbConnection, container.Terminate
}

func (s *PostgresqlSuite) createMigrator() *migrate.Migrate {
	d, err := iofs.New(migrations.MigrationFS, "postgresql")
	if err != nil {
		log.Fatal(err)
	}
	db, err := sql.Open("postgres", s.DbUrl)
	s.Require().NoError(err)

	driver, err := postgres.WithInstance(db, &postgres.Config{})
	s.Require().NoError(err)

	s.migrator, err = migrate.NewWithInstance("iofs", d, s.DbUrl, driver)
	s.Require().NoError(err)

	// run migrations
	return s.migrator
}
