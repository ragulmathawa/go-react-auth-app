package main

import (
	"context"
	"log"
	"log/slog"
	"net/http"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/ragulmathawa/go-react-auth-app/pkg/db"
	"github.com/ragulmathawa/go-react-auth-app/pkg/utils"
	"github.com/supertokens/supertokens-golang/recipe/emailpassword"
	"github.com/supertokens/supertokens-golang/recipe/session"
	"github.com/supertokens/supertokens-golang/recipe/thirdparty"
	"github.com/supertokens/supertokens-golang/recipe/thirdparty/tpmodels"
	"github.com/supertokens/supertokens-golang/supertokens"
)

func main() {

	appConfig := utils.GetAppConfig()
	utils.InitLogging(appConfig)

	initSuperTokens(appConfig)
	db.InitSqlite(appConfig)
	defer db.CloseDB()
	initHttpServer(appConfig) // with graceful shutdown

}

func initHttpServer(appConfig utils.AppConfig) {
	// Create context that listens for the interrupt signal from the OS.
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	if appConfig.Mode == "production" {
		gin.SetMode(gin.ReleaseMode)
	}

	router := gin.New()
	router.SetTrustedProxies(nil)
	// Adding the SuperTokens middleware
	router.Use(cors.New(cors.Config{
		AllowOrigins: []string{appConfig.WebsiteDomain},
		AllowMethods: []string{"GET", "POST", "DELETE", "PUT", "OPTIONS"},
		AllowHeaders: append([]string{"content-type"},
			supertokens.GetAllCORSHeaders()...),
		AllowCredentials: true,
	}))
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	router.Use(func(c *gin.Context) {
		supertokens.Middleware(http.HandlerFunc(
			func(rw http.ResponseWriter, r *http.Request) {
				c.Next()
			})).ServeHTTP(c.Writer, c.Request)
		// we call Abort so that the next handler in the chain is not called, unless we call Next explicitly
		c.Abort()
	})

	srv := &http.Server{
		Addr:    ":8080",
		Handler: router.Handler(),
	}

	go func() {
		slog.Info("Starting server", "port", appConfig.Port)
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	// Listen for the interrupt signal.
	<-ctx.Done()

	// Restore default behavior on the interrupt signal and notify user of shutdown.
	stop()
	log.Println("shutting down gracefully, press Ctrl+C again to force")

	// The context is used to inform the server it has 5 seconds to finish
	// the request it is currently handling
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server forced to shutdown: ", err)
	}

	log.Println("Server exiting")
}

func initSuperTokens(appConfig utils.AppConfig) {
	// These are the connection details of the app you created on supertokens.com
	// We have provided you with development keys which you can use for testing.
	// IMPORTANT: Please replace them with your own OAuth keys for production use.

	apiBasePath := "/api/auth"
	websiteBasePath := "/auth"
	err := supertokens.Init(supertokens.TypeInput{
		Supertokens: &supertokens.ConnectionInfo{

			ConnectionURI: appConfig.STAuthURI,
			APIKey:        appConfig.STAuthAPIKey,
		},
		Debug: appConfig.AuthDebug,
		AppInfo: supertokens.AppInfo{
			AppName:         appConfig.AppName,
			APIDomain:       appConfig.APIDomain,
			WebsiteDomain:   appConfig.WebsiteDomain,
			APIBasePath:     &apiBasePath,
			WebsiteBasePath: &websiteBasePath,
		},
		RecipeList: []supertokens.Recipe{
			thirdparty.Init(&tpmodels.TypeInput{
				SignInAndUpFeature: tpmodels.TypeInputSignInAndUp{
					Providers: []tpmodels.ProviderInput{

						{
							Config: tpmodels.ProviderConfig{
								ThirdPartyId: "google",
								Clients: []tpmodels.ProviderClientConfig{
									{
										ClientID:     "1060725074195-kmeum4crr01uirfl2op9kd5acmi9jutn.apps.googleusercontent.com",
										ClientSecret: "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",
									},
								},
							},
						},
						{
							Config: tpmodels.ProviderConfig{
								ThirdPartyId: "github",
								Clients: []tpmodels.ProviderClientConfig{
									{
										ClientID:     "467101b197249757c71f",
										ClientSecret: "e97051221f4b6426e8fe8d51486396703012f5bd",
									},
								},
							},
						},
						{
							Config: tpmodels.ProviderConfig{
								ThirdPartyId: "apple",
								Clients: []tpmodels.ProviderClientConfig{
									{
										ClientID: "4398792-io.supertokens.example.service",
										AdditionalConfig: map[string]interface{}{
											"keyId":      "7M48Y4RYDL",
											"privateKey": "-----BEGIN PRIVATE KEY-----\nMIGTAgEAMBMGByqGSM49AgEGCCqGSM49AwEHBHkwdwIBAQQgu8gXs+XYkqXD6Ala9Sf/iJXzhbwcoG5dMh1OonpdJUmgCgYIKoZIzj0DAQehRANCAASfrvlFbFCYqn3I2zeknYXLwtH30JuOKestDbSfZYxZNMqhF/OzdZFTV0zc5u5s3eN+oCWbnvl0hM+9IW0UlkdA\n-----END PRIVATE KEY-----",
											"teamId":     "YWQCXGJRJL",
										},
									},
								},
							},
						},
					},
				},
			}),
			emailpassword.Init(nil),
			session.Init(nil), // initializes session features
		},
	})

	if err != nil {
		panic(err.Error())
	}
}
