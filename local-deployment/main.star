zond_module = import_module("github.com/theQRL/zond-package/main.star")

POSTGRES_PORT_ID = "postgres"
POSTGRES_DB = "db"
POSTGRES_USER = "postgres"
POSTGRES_PASSWORD = "pass"

POSTGRESADM_PORT_ID = "postgresadm"

REDIS_PORT_ID = "redis"

REDISINSIGHT_PORT_ID = "redisinsight"

REDASH_PORT_ID = "redash"

LITTLE_BIGTABLE_PORT_ID = "littlebigtable"

def run(plan, args):
	db_services = plan.add_services(
		configs={
			# Add a Postgres server
			"postgres": ServiceConfig(
				image = "postgres:15.2-alpine",
				ports = {
					POSTGRES_PORT_ID: PortSpec(5432, application_protocol = "postgresql"),
				},
				env_vars = {
					"POSTGRES_DB": POSTGRES_DB,
					"POSTGRES_USER": POSTGRES_USER,
					"POSTGRES_PASSWORD": POSTGRES_PASSWORD,
				},
			),
			# Add Postgres Admin
			"postgresadm": ServiceConfig(
			 	image = "dpage/pgadmin4",
				ports = {
					POSTGRESADM_PORT_ID: PortSpec(80, application_protocol = "tcp"),
				},
				env_vars = {
					"PGADMIN_DEFAULT_EMAIL": "test@test.com",
					"PGADMIN_DEFAULT_PASSWORD": "pass",
				},
			),
			# Add a Redis server
			"redis": ServiceConfig(
				image = "redis:7",
				ports = {
					REDIS_PORT_ID: PortSpec(6379, application_protocol = "tcp"),
				},
			),
			# Add RedisInsight
			"redisinsight": ServiceConfig(
				image = "redis/redisinsight:latest",
				ports = {
					REDISINSIGHT_PORT_ID: PortSpec(5540, application_protocol = "tcp"),
				},
			),
			# Add a Bigtable Emulator server
			"littlebigtable": ServiceConfig(
				image = "gobitfly/little_bigtable:latest",
				ports = {
					LITTLE_BIGTABLE_PORT_ID: PortSpec(9000, application_protocol = "tcp"),
				},
			),
			# Add Redash to debug redis and postgres
			#"redash": ServiceConfig(
			#	image = "redash/redash:latest",
			#	ports = {
			#		REDASH_PORT_ID: PortSpec(5000, application_protocol = "tcp"),
			#	},
			#	env_vars = {
			#		"PYTHONUNBUFFERED": "0",
			#		"REDASH_LOG_LEVEL": "INFO",
			#		"REDASH_REDIS_URL": "redis://redis:6379/0",
			#		"REDASH_DATABASE_URL": "postgresql://postgres:pass@postgres/db",
			#		"REDASH_COOKIE_SECRET": "veryverysecret",
			#		"REDASH_WEB_WORKERS": "4",
			#	},
			#),
			#"redashcreatedb": ServiceConfig(
			#	image = "redash/redash:latest",
			#	cmd=["create_db"],
			#	env_vars = {
			#		"REDASH_LOG_LEVEL": "INFO",
			#		"REDASH_REDIS_URL": "redis://redis:6379/0",
			#		"REDASH_DATABASE_URL": "postgresql://postgres:pass@postgres/db",
			#		"REDASH_COOKIE_SECRET": "veryverysecret",
			#		"REDASH_WEB_WORKERS": "4",
			#	},
			#),
		}
	)
	
	# Spin up a local zond testnet
	zond_module.run(plan, args)