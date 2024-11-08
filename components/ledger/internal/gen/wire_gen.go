// Code generated by Wire. DO NOT EDIT.

//go:generate go run -mod=mod github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package gen

import (
	"fmt"
	"github.com/LerianStudio/midaz/common"
	"github.com/LerianStudio/midaz/common/mcasdoor"
	"github.com/LerianStudio/midaz/common/mlog"
	"github.com/LerianStudio/midaz/common/mmongo"
	"github.com/LerianStudio/midaz/common/mopentelemetry"
	"github.com/LerianStudio/midaz/common/mpostgres"
	mrabbitmq2 "github.com/LerianStudio/midaz/common/mrabbitmq"
	"github.com/LerianStudio/midaz/common/mzap"
	"github.com/LerianStudio/midaz/components/ledger/internal/adapters/database/mongodb"
	"github.com/LerianStudio/midaz/components/ledger/internal/adapters/database/postgres"
	"github.com/LerianStudio/midaz/components/ledger/internal/adapters/rabbitmq"
	"github.com/LerianStudio/midaz/components/ledger/internal/app/command"
	"github.com/LerianStudio/midaz/components/ledger/internal/app/query"
	"github.com/LerianStudio/midaz/components/ledger/internal/domain/metadata"
	"github.com/LerianStudio/midaz/components/ledger/internal/domain/onboarding/ledger"
	"github.com/LerianStudio/midaz/components/ledger/internal/domain/onboarding/organization"
	"github.com/LerianStudio/midaz/components/ledger/internal/domain/portfolio/account"
	"github.com/LerianStudio/midaz/components/ledger/internal/domain/portfolio/asset"
	"github.com/LerianStudio/midaz/components/ledger/internal/domain/portfolio/portfolio"
	"github.com/LerianStudio/midaz/components/ledger/internal/domain/portfolio/product"
	"github.com/LerianStudio/midaz/components/ledger/internal/domain/rabbitmq"
	"github.com/LerianStudio/midaz/components/ledger/internal/ports/grpc"
	"github.com/LerianStudio/midaz/components/ledger/internal/ports/http"
	"github.com/LerianStudio/midaz/components/ledger/internal/service"
	"github.com/google/wire"
	"sync"
)

// Injectors from inject.go:

// InitializeService the setup the dependencies and returns a new *service.Service instance
func InitializeService() *service.Service {
	config := service.NewConfig()
	logger := mzap.InitializeLogger()
	telemetry := setupTelemetryProviders(config, logger)
	casdoorConnection := setupCasdoorConnection(config, logger)
	postgresConnection := setupPostgreSQLConnection(config, logger)
	organizationPostgreSQLRepository := postgres.NewOrganizationPostgreSQLRepository(postgresConnection)
	ledgerPostgreSQLRepository := postgres.NewLedgerPostgreSQLRepository(postgresConnection)
	productPostgreSQLRepository := postgres.NewProductPostgreSQLRepository(postgresConnection)
	portfolioPostgreSQLRepository := postgres.NewPortfolioPostgreSQLRepository(postgresConnection)
	accountPostgreSQLRepository := postgres.NewAccountPostgreSQLRepository(postgresConnection)
	assetPostgreSQLRepository := postgres.NewAssetPostgreSQLRepository(postgresConnection)
	mongoConnection := setupMongoDBConnection(config, logger)
	metadataMongoDBRepository := mongodb.NewMetadataMongoDBRepository(mongoConnection)
	rabbitMQConnection := setupRabbitMQConnection(config, logger)
	producerRabbitMQRepository := mrabbitmq.NewProducerRabbitMQ(rabbitMQConnection)
	useCase := &command.UseCase{
		OrganizationRepo: organizationPostgreSQLRepository,
		LedgerRepo:       ledgerPostgreSQLRepository,
		ProductRepo:      productPostgreSQLRepository,
		PortfolioRepo:    portfolioPostgreSQLRepository,
		AccountRepo:      accountPostgreSQLRepository,
		AssetRepo:        assetPostgreSQLRepository,
		MetadataRepo:     metadataMongoDBRepository,
		RabbitMQRepo:     producerRabbitMQRepository,
	}
	consumerRabbitMQRepository := mrabbitmq.NewConsumerRabbitMQ(rabbitMQConnection)
	queryUseCase := &query.UseCase{
		OrganizationRepo: organizationPostgreSQLRepository,
		LedgerRepo:       ledgerPostgreSQLRepository,
		ProductRepo:      productPostgreSQLRepository,
		PortfolioRepo:    portfolioPostgreSQLRepository,
		AccountRepo:      accountPostgreSQLRepository,
		AssetRepo:        assetPostgreSQLRepository,
		MetadataRepo:     metadataMongoDBRepository,
		RabbitMQRepo:     consumerRabbitMQRepository,
	}
	accountHandler := &http.AccountHandler{
		Command: useCase,
		Query:   queryUseCase,
	}
	portfolioHandler := &http.PortfolioHandler{
		Command: useCase,
		Query:   queryUseCase,
	}
	ledgerHandler := &http.LedgerHandler{
		Command: useCase,
		Query:   queryUseCase,
	}
	assetHandler := &http.AssetHandler{
		Command: useCase,
		Query:   queryUseCase,
	}
	organizationHandler := &http.OrganizationHandler{
		Command: useCase,
		Query:   queryUseCase,
	}
	productHandler := &http.ProductHandler{
		Command: useCase,
		Query:   queryUseCase,
	}
	app := http.NewRouter(logger, telemetry, casdoorConnection, accountHandler, portfolioHandler, ledgerHandler, assetHandler, organizationHandler, productHandler)
	server := service.NewServer(config, app, logger, telemetry)
	grpcServer := grpc.NewRouterGRPC(logger, casdoorConnection, useCase, queryUseCase)
	serverGRPC := service.NewServerGRPC(config, grpcServer, logger)
	serviceService := &service.Service{
		Server:     server,
		ServerGRPC: serverGRPC,
		Logger:     logger,
	}
	return serviceService
}

// inject.go:

var onceConfig sync.Once

const prdEnvName = "production"

func setupPostgreSQLConnection(cfg *service.Config, log mlog.Logger) *mpostgres.PostgresConnection {
	connStrPrimary := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		cfg.PrimaryDBHost, cfg.PrimaryDBUser, cfg.PrimaryDBPassword, cfg.PrimaryDBName, cfg.PrimaryDBPort)

	connStrReplica := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		cfg.ReplicaDBHost, cfg.ReplicaDBUser, cfg.ReplicaDBPassword, cfg.ReplicaDBName, cfg.ReplicaDBPort)

	return &mpostgres.PostgresConnection{
		ConnectionStringPrimary: connStrPrimary,
		ConnectionStringReplica: connStrReplica,
		PrimaryDBName:           cfg.PrimaryDBName,
		ReplicaDBName:           cfg.ReplicaDBName,
		Component:               "ledger",
		Logger:                  log,
	}
}

func setupMongoDBConnection(cfg *service.Config, log mlog.Logger) *mmongo.MongoConnection {
	connStrSource := fmt.Sprintf("mongodb://%s:%s@%s:%s",
		cfg.MongoDBUser, cfg.MongoDBPassword, cfg.MongoDBHost, cfg.MongoDBPort)

	return &mmongo.MongoConnection{
		ConnectionStringSource: connStrSource,
		Database:               cfg.MongoDBName,
		Logger:                 log,
	}
}

func setupCasdoorConnection(cfg *service.Config, log mlog.Logger) *mcasdoor.CasdoorConnection {
	casdoor := &mcasdoor.CasdoorConnection{
		JWKUri:           cfg.JWKAddress,
		Endpoint:         cfg.CasdoorAddress,
		ClientID:         cfg.CasdoorClientID,
		ClientSecret:     cfg.CasdoorClientSecret,
		OrganizationName: cfg.CasdoorOrganizationName,
		ApplicationName:  cfg.CasdoorApplicationName,
		EnforcerName:     cfg.CasdoorEnforcerName,
		Logger:           log,
	}

	return casdoor
}

func setupRabbitMQConnection(cfg *service.Config, log mlog.Logger) *mrabbitmq2.RabbitMQConnection {
	connStrSource := fmt.Sprintf("amqp://%s:%s@%s:%s",
		cfg.RabbitMQUser, cfg.RabbitMQPass, cfg.RabbitMQHost, cfg.RabbitMQPortHost)

	return &mrabbitmq2.RabbitMQConnection{
		ConnectionStringSource: connStrSource,
		Host:                   cfg.RabbitMQHost,
		Port:                   cfg.RabbitMQPortAMQP,
		User:                   cfg.RabbitMQUser,
		Pass:                   cfg.RabbitMQPass,
		Exchange:               cfg.RabbitMQExchange,
		Key:                    cfg.RabbitMQKey,
		Queue:                  cfg.RabbitMQQueue,
		Logger:                 log,
	}
}

func setupTelemetryProviders(cfg *service.Config, log mlog.Logger) *mopentelemetry.Telemetry {
	t := &mopentelemetry.Telemetry{
		LibraryName:               cfg.OtelLibraryName,
		ServiceName:               cfg.OtelServiceName,
		ServiceVersion:            cfg.OtelServiceVersion,
		DeploymentEnv:             cfg.OtelDeploymentEnv,
		CollectorExporterEndpoint: cfg.OtelColExporterEndpoint,
		Logger:                    log,
	}

	return t
}

var (
	serviceSet = wire.NewSet(common.InitLocalEnvConfig, mzap.InitializeLogger, setupTelemetryProviders,
		setupPostgreSQLConnection,
		setupMongoDBConnection,
		setupCasdoorConnection,
		setupRabbitMQConnection, grpc.NewRouterGRPC, service.NewServerGRPC, http.NewRouter, service.NewConfig, service.NewServer, postgres.NewOrganizationPostgreSQLRepository, postgres.NewLedgerPostgreSQLRepository, postgres.NewAssetPostgreSQLRepository, postgres.NewPortfolioPostgreSQLRepository, postgres.NewProductPostgreSQLRepository, postgres.NewAccountPostgreSQLRepository, mongodb.NewMetadataMongoDBRepository, mrabbitmq.NewProducerRabbitMQ, mrabbitmq.NewConsumerRabbitMQ, wire.Struct(new(http.OrganizationHandler), "*"), wire.Struct(new(http.LedgerHandler), "*"), wire.Struct(new(http.AssetHandler), "*"), wire.Struct(new(http.PortfolioHandler), "*"), wire.Struct(new(http.ProductHandler), "*"), wire.Struct(new(http.AccountHandler), "*"), wire.Struct(new(command.UseCase), "*"), wire.Struct(new(query.UseCase), "*"), wire.Bind(new(organization.Repository), new(*postgres.OrganizationPostgreSQLRepository)), wire.Bind(new(ledger.Repository), new(*postgres.LedgerPostgreSQLRepository)), wire.Bind(new(asset.Repository), new(*postgres.AssetPostgreSQLRepository)), wire.Bind(new(portfolio.Repository), new(*postgres.PortfolioPostgreSQLRepository)), wire.Bind(new(product.Repository), new(*postgres.ProductPostgreSQLRepository)), wire.Bind(new(account.Repository), new(*postgres.AccountPostgreSQLRepository)), wire.Bind(new(metadata.Repository), new(*mongodb.MetadataMongoDBRepository)), wire.Bind(new(rabbitmq.ConsumerRepository), new(*mrabbitmq.ConsumerRabbitMQRepository)), wire.Bind(new(rabbitmq.ProducerRepository), new(*mrabbitmq.ProducerRabbitMQRepository)),
	)

	svcSet = wire.NewSet(wire.Struct(new(service.Service), "Server", "ServerGRPC", "Logger"))
)
