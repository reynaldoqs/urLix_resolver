package server

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	service "github.com/reynaldoqs/urLix_resolver/internal/core/services"
	"github.com/reynaldoqs/urLix_resolver/internal/infrastructure/controller"
	ffirebase "github.com/reynaldoqs/urLix_resolver/internal/infrastructure/google"
	mongodb "github.com/reynaldoqs/urLix_resolver/internal/infrastructure/repositories/mongo"
)

func RegisterRouter(port string) {
	chiDispatcher := chi.NewRouter()
	chiDispatcher.Use(middleware.RequestID)
	chiDispatcher.Use(middleware.RealIP)
	chiDispatcher.Use(middleware.Logger)
	chiDispatcher.Use(middleware.Recoverer)

	mrepo := mongodb.NewRechargesRepository("mongodb://localhost:27017", "project-x", 30)
	gservice := ffirebase.NewFirebaseApp("./gu-project.json")

	rservice := service.NewService(gservice, gservice, mrepo, mrepo)
	rcontroller := controller.NewRechargesController(rservice)

	chiDispatcher.Get("/recharges", rcontroller.GetRecharges)
	chiDispatcher.Post("/recharges", rcontroller.AddRecharge)

	aservice := service.NewAdminService(gservice, gservice, mrepo)
	acontroller := controller.NewAdminController(aservice)

	chiDispatcher.Post("/admin/exec", acontroller.PostExecution)
	chiDispatcher.Post("/admin/ussd", acontroller.PostUssdAction)
	chiDispatcher.Get("/admin/ussd", acontroller.GetUssdActions)
	chiDispatcher.Patch("/admin/ussd", acontroller.PatchUssdAction)

	reportservice := service.NewReportsRepository(mrepo)
	reportscontroller := controller.NewReportsController(reportservice)

	chiDispatcher.Post("/reports/admin", reportscontroller.PostAdminReport)
	chiDispatcher.Post("/reports/recharge", reportscontroller.PostRechargeReport)

	fmt.Printf("Chi HTTP server running on port %v\n", port)
	http.ListenAndServe(port, chiDispatcher)

}
