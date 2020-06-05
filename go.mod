module video-server

go 1.14

replace (
	video-server/api => ./api
	video-server/api/database => ./api/database
	video-server/api/handlers => ./api/handlers
	video-server/api/model => ./api/model
	video-server/api/session => ./api/session
	video-server/api/utils => ./api/utils
	video-server/streame_server => ./streame_server
	video-server/streame_server/config => ./streame_server/config
	video-server/streame_server/handlers => ./streame_server/handlers
	video-server/streame_server/token_bucket => ./streame_server/token_bucket
	video-server/streame_server/util => ./streame_server/util
	video-server/scheduler => ./scheduler
	video-server/scheduler/database => ./scheduler/database
	video-server/scheduler/handlers => ./scheduler/handlers
	video-server/scheduler/taskrunner => ./scheduler/taskrunner
	video-server/scheduler/util => ./scheduler/util
)

require (
	github.com/go-sql-driver/mysql v1.5.0
	github.com/julienschmidt/httprouter v1.3.0
)
