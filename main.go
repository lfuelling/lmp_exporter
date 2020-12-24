package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"
)

type Config struct {
	Servers               []string
	ListenAddress         string
	DetailedVesselMetrics bool
}

type Vessel struct {
	Id                       string  `json:"Id"`
	Name                     string  `json:"Name"`
	Type                     string  `json:"Type"`
	DistanceTravelled        float64 `json:"DistanceTravelled"`
	Situation                string  `json:"Situation"`
	Lat                      float64 `json:"Lat"`
	Lon                      float64 `json:"Lon"`
	Alt                      float64 `json:"Alt"`
	SemiMajorAxis            float64 `json:"SemiMajorAxis"`
	Eccentricity             float64 `json:"Eccentricity"`
	Inclination              float64 `json:"Inclination"`
	ArgumentOfPeriapsis      float64 `json:"ArgumentOfPeriapsis"`
	LongitudeOfAscendingNode float64 `json:"LongitudeOfAscendingNode"`
	MeanAnomaly              float64 `json:"MeanAnomaly"`
	Epoch                    float64 `json:"Epoch"`
	ReferenceBody            int64   `json:"ReferenceBody"`
}

type Subspace struct {
	Id      int     `json:"Id"`
	Time    float64 `json:"Time"`
	Creator string  `json:"Creator"`
}

type GeneralSettings struct {
	ServerName        string  `json:"ServerName"`
	Description       string  `json:"Description"`
	CountryCode       string  `json:"CountryCode"`
	WebsiteText       string  `json:"WebsiteText"`
	Website           string  `json:"Website"`
	HasPassword       bool    `json:"HasPassword"`
	HasAdminPassword  bool    `json:"HasAdminPassword"`
	MaxPlayers        int     `json:"MaxPlayers"`
	MaxUsernameLength int     `json:"MaxUsernameLength"`
	AutoDekessler     float64 `json:"AutoDekessler"`
	AutoNuke          float64 `json:"AutoNuke"`
	Cheats            bool    `json:"Cheats"`
	AllowSackKerbals  bool    `json:"AllowSackKerbals"`
	NumberOfAsteroids int     `json:"NumberOfAsteroids"`
	NumberOfComets    int     `json:"NumberOfComets"`
}

type Metrics struct {
	Server          string
	StartTime       string          `json:"StartTime"`
	Vessels         []Vessel        `json:"CurrentVessels"`
	Subspaces       []Subspace      `json:"Subspaces"`
	GeneralSettings GeneralSettings `json:"GeneralSettings"`
	CurrentPlayers  []string        `json:"CurrentPlayers"`
	Success         bool
	Duration        int64
}

var config Config

func loadMetrics() []Metrics {

	var serverMetrics []Metrics

	for _, d := range config.Servers {
		startTime := time.Now().UnixNano()
		log.Println("Connecting to '" + d + "'...")
		resp, err := http.Get("http://" + d + "/")
		if err != nil {
			log.Println("Error fetching '"+d+"'!", err)
			serverMetrics = append(serverMetrics, Metrics{Success: false, Server: d, Duration: time.Now().UnixNano() - startTime})
		} else {
			var metrics Metrics
			err := json.NewDecoder(resp.Body).Decode(&metrics)
			if err != nil {
				log.Println("Error parsing '"+d+"'!", err)
				serverMetrics = append(serverMetrics, Metrics{Success: false, Server: d, Duration: time.Now().UnixNano() - startTime})
			} else {
				serverMetrics = append(serverMetrics, Metrics{
					Server:          d,
					StartTime:       metrics.StartTime,
					Vessels:         metrics.Vessels,
					Subspaces:       metrics.Subspaces,
					GeneralSettings: metrics.GeneralSettings,
					CurrentPlayers:  metrics.CurrentPlayers,
					Success:         true,
					Duration:        time.Now().UnixNano() - startTime,
				})
			}
		}
		err1 := resp.Body.Close()
		if err1 != nil {
			log.Println("Error closing connection '"+d+"'!", err)
		}
	}
	return serverMetrics
}

func renderMetricsResponse() string {
	metrics := loadMetrics()

	res := "# HELP lmp_current_players The number of currently active players.\n" +
		"# TYPE lmp_current_players gauge\n" +
		"# HELP lmp_subspaces_total The number of subspaces.\n" +
		"# TYPE lmp_subspaces_total gauge\n" +
		"# HELP lmp_vessels_total The number of vessels.\n" +
		"# TYPE lmp_vessels_total gauge\n" +
		"# HELP lmp_subspaces_time The time of subspaces.\n" +
		"# TYPE lmp_subspaces_time gauge\n" +
		"# HELP lmp_vessels_distance_travelled The total distance travelled by a vessel.\n" +
		"# TYPE lmp_vessels_distance_travelled counter\n" +
		"# HELP lmp_vessels_lat The latitude of a vessel.\n" +
		"# TYPE lmp_vessels_lat gauge\n" +
		"# HELP lmp_vessels_lon The longitude of a vessel.\n" +
		"# TYPE lmp_vessels_lon gauge\n" +
		"# HELP lmp_vessels_alt The altitude of a vessel.\n" +
		"# TYPE lmp_vessels_alt gauge\n" +
		"# HELP lmp_vessels_epoch The epoch of a vessel.\n" +
		"# TYPE lmp_vessels_epoch gauge\n" +
		"# HELP lmp_vessels_reference_body The reference body of a vessel.\n" +
		"# TYPE lmp_vessels_reference_body gauge\n" +
		"# HELP lmp_scrapes_duration The duration of a the metrics collection in nanoseconds.\n" +
		"# TYPE lmp_scrapes_duration gauge\n"

	if config.DetailedVesselMetrics {
		res += "# HELP lmp_vessel_semi_major_axis The semi major axis of a vessel.\n" +
			"# TYPE lmp_vessel_semi_major_axis gauge\n" +
			"# HELP lmp_vessel_eccentricity The eccentricity of a vessel.\n" +
			"# TYPE lmp_vessel_eccentricity gauge\n" +
			"# HELP lmp_vessel_inclination The inclination of a vessel.\n" +
			"# TYPE lmp_vessel_inclination gauge\n" +
			"# HELP lmp_vessel_argument_of_periapsis A vessels argument of periapsis.\n" +
			"# TYPE lmp_vessel_argument_of_periapsis gauge\n" +
			"# HELP lmp_vessel_lon_of_ascending_node A vessels longitude of ascending node.\n" +
			"# TYPE lmp_vessel_lon_of_ascending_node gauge\n" +
			"# HELP lmp_vessel_mean_anomaly A vessels mean anomaly.\n" +
			"# TYPE lmp_vessel_mean_anomaly gauge\n"
	}

	for _, m := range metrics {
		res += "# Start of metrics for " + m.Server + " \n"
		if m.Success {
			res += `lmp_current_players{server="` + m.Server + `"} ` + strconv.FormatInt(int64(len(m.CurrentPlayers)), 10) + "\n"
			res += `lmp_subspaces_total{server="` + m.Server + `"} ` + strconv.FormatInt(int64(len(m.Subspaces)), 10) + "\n"

			for _, s := range m.Subspaces {
				res += `lmp_subspaces_time{server="` + m.Server + `",subspace="` + strconv.FormatInt(int64(s.Id), 10) + `"} ` + strconv.FormatFloat(s.Time, 'f', 6, 64) + "\n"
			}

			for _, v := range m.Vessels {
				res += `lmp_vessel_distance_travelled{server="` + m.Server + `",vessel="` + v.Name + `",type="` + v.Type + `",vesselId="` + v.Id + `"} ` + strconv.FormatFloat(v.DistanceTravelled, 'f', 6, 64) + "\n"

				res += `lmp_vessel_lat{server="` + m.Server + `",vessel="` + v.Name + `",type="` + v.Type + `",vesselId="` + v.Id + `"} ` + strconv.FormatFloat(v.Lat, 'f', 6, 64) + "\n"
				res += `lmp_vessel_lon{server="` + m.Server + `",vessel="` + v.Name + `",type="` + v.Type + `",vesselId="` + v.Id + `"} ` + strconv.FormatFloat(v.Lon, 'f', 6, 64) + "\n"
				res += `lmp_vessel_alt{server="` + m.Server + `",vessel="` + v.Name + `",type="` + v.Type + `",vesselId="` + v.Id + `"} ` + strconv.FormatFloat(v.Alt, 'f', 6, 64) + "\n"

				if config.DetailedVesselMetrics {
					res += `lmp_vessel_semi_major_axis{server="` + m.Server + `",vessel="` + v.Name + `",type="` + v.Type + `",vesselId="` + v.Id + `"} ` + strconv.FormatFloat(v.SemiMajorAxis, 'f', 6, 64) + "\n"
					res += `lmp_vessel_eccentricity{server="` + m.Server + `",vessel="` + v.Name + `",type="` + v.Type + `",vesselId="` + v.Id + `"} ` + strconv.FormatFloat(v.Eccentricity, 'f', 6, 64) + "\n"
					res += `lmp_vessel_inclination{server="` + m.Server + `",vessel="` + v.Name + `",type="` + v.Type + `",vesselId="` + v.Id + `"} ` + strconv.FormatFloat(v.Inclination, 'f', 6, 64) + "\n"
					res += `lmp_vessel_argument_of_periapsis{server="` + m.Server + `",vessel="` + v.Name + `",type="` + v.Type + `",vesselId="` + v.Id + `"} ` + strconv.FormatFloat(v.ArgumentOfPeriapsis, 'f', 6, 64) + "\n"
					res += `lmp_vessel_lon_of_ascending_node{server="` + m.Server + `",vessel="` + v.Name + `",type="` + v.Type + `",vesselId="` + v.Id + `"} ` + strconv.FormatFloat(v.LongitudeOfAscendingNode, 'f', 6, 64) + "\n"
					res += `lmp_vessel_mean_anomaly{server="` + m.Server + `",vessel="` + v.Name + `",type="` + v.Type + `",vesselId="` + v.Id + `"} ` + strconv.FormatFloat(v.MeanAnomaly, 'f', 6, 64) + "\n"
				}

				res += `lmp_vessel_epoch{server="` + m.Server + `",vessel="` + v.Name + `",type="` + v.Type + `",vesselId="` + v.Id + `"} ` + strconv.FormatFloat(v.Epoch, 'f', 6, 64) + "\n"
				res += `lmp_vessel_reference_body{server="` + m.Server + `",vessel="` + v.Name + `",type="` + v.Type + `",vesselId="` + v.Id + `"} ` + strconv.FormatInt(v.ReferenceBody, 10) + "\n"
			}
			res += `lmp_scrape_success{server="` + m.Server + `"} 1` + "\n"
		} else {
			res += `lmp_scrape_success{server="` + m.Server + `"} 0` + "\n"
		}
		res += `lmp_scrape_duration{server="` + m.Server + `"} ` + strconv.FormatInt(m.Duration, 10) + "\n"
	}

	return res
}

func handleMetrics(w http.ResponseWriter, r *http.Request) {
	if r.RequestURI == "/metrics" {
		response := renderMetricsResponse()
		_, _ = fmt.Fprint(w, response)
	} else {
		log.Println("Not found: '" + r.RequestURI)
		w.WriteHeader(404)
	}
}

func main() {
	var configPath = flag.String("config", "config.json", "path to the config file")
	flag.Parse()
	file, err := os.Open(*configPath)
	defer file.Close()
	if err != nil {
		log.Fatalln("Unable to open config file!", err)
		return
	}
	decoder := json.NewDecoder(file)
	config = Config{}
	err1 := decoder.Decode(&config)
	if err1 != nil {
		log.Fatalln("Unable to read config!", err1)
		return
	}

	server := &http.Server{
		Addr:         config.ListenAddress,
		Handler:      http.HandlerFunc(handleMetrics),
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}
	log.Println("Starting server...")
	err2 := server.ListenAndServe()
	if err2 != nil {
		log.Fatalln("Unable to start server!", err2)
	}
}
