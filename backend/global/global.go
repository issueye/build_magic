package global

import (
	"net/http"

	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
	"github.com/issueye/build_magic/backend/code_engine"
	"go.uber.org/zap"
	"gopkg.in/antage/eventsource.v1"
	"gorm.io/gorm"
)

var (
	DB         *gorm.DB
	Log        *zap.SugaredLogger
	Logger     *zap.Logger
	Router     *gin.Engine
	HttpServer *http.Server
	Auth       *jwt.GinJWTMiddleware
	SSE        eventsource.EventSource
	JsVMCore   *code_engine.Core
)
