package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"path"

	"github.com/Brightwater/ironmanager/groupIron"
	"github.com/Brightwater/ironmanager/htmlgrabber"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/cors"
	"github.com/go-chi/render"
)

func SetupRoutes(client *groupIron.ApiClient) *chi.Mux {
	r := chi.NewRouter()

	cors := cors.New(
		cors.Options{
			AllowedOrigins:   []string{"*"},
			AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
			AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token", "Access-Control-Allow-Origin"},
			AllowCredentials: true,
		},
	)
	r.Use(cors.Handler)

	r.Use(middleware.Logger) // add log middleware

	// Serve the index.html file as the root path
	r.Get("/*", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, path.Join("static", "index.html"))
	})

	// Serve static files from the "public" directory
	r.Handle("/_app/*", http.FileServer(http.Dir("static")))
	r.Handle("/ui/*", http.FileServer(http.Dir("static")))
	r.Handle("/fonts/*", http.FileServer(http.Dir("static")))
 
	r.Get("/skillsScreenshot/{member}", getIronManScreenShotForMember())
	r.Get("/getIronData/{member}", getDataForIronMan(client))
	r.Get("/getIronXpYear", getYearXpDataForIronMan(client))
	r.Get("/graphScreenshot", getIronManYearXpGraph())

	return r
}

func getIronManYearXpGraph() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		buf, err := htmlgrabber.GrabGroupIronGraphPage()
		if err != nil {
			render.Status(r, http.StatusInternalServerError)
			render.PlainText(w, r, "Something went wrong")
			return
		}

		w.Header().Set("Content-Type", "image/png")
		w.Header().Set("Content-Disposition", "inline; filename=elementScreenshot.png")

		_, err = w.Write(*buf)
		if err != nil {
			render.Status(r, http.StatusInternalServerError)
			render.PlainText(w, r, "Something went wrong")
			return
		}
		// fmt.Fprintf(w, "Hello world")
	}
}

func getIronManScreenShotForMember() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		memberName := chi.URLParam(r, "member")

		buf, err := htmlgrabber.GrabGroupIronSkillsPage(memberName)
		if err != nil {
			render.Status(r, http.StatusInternalServerError)
			render.PlainText(w, r, "Something went wrong")
			return
		}

		w.Header().Set("Content-Type", "image/png")
		w.Header().Set("Content-Disposition", "inline; filename=elementScreenshot.png")

		_, err = w.Write(*buf)
		if err != nil {
			render.Status(r, http.StatusInternalServerError)
			render.PlainText(w, r, "Something went wrong")
			return
		}
		// fmt.Fprintf(w, "Hello world")
	}
}

func getDataForIronMan(client *groupIron.ApiClient) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		name := chi.URLParam(r, "member")

		players := groupIron.GetAllPlayersCurrentStatus(client)
		if players == nil {
			http.Error(w, "Failed to get players", http.StatusInternalServerError)
			return
		}

		player, err := groupIron.GetPlayerCurrentStats(players, name)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		jsonPlayers, err := json.Marshal(*player)
		if err != nil {
			render.Status(r, http.StatusInternalServerError)
			render.PlainText(w, r, "Something went wrong")
		}

		fmt.Fprintf(w, "%s", string(jsonPlayers))
	}
}

func getYearXpDataForIronMan(client *groupIron.ApiClient) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// name := chi.URLParam(r, "member")

		xpOverTimeJson, err := client.GetXpAllTime()
		if err != nil {
			http.Error(w, "Failed to get players", http.StatusInternalServerError)
			return
		}

		fmt.Fprintf(w, "%s", string(xpOverTimeJson))
	}
}
