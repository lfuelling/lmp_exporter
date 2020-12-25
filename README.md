# lmp-exporter

A prometheus exporter that returns the metrics visible on the "website" of an [LunaMultiplayer](https://github.com/LunaMultiplayer/LunaMultiplayer) server.

## Usage

1. Clone repo
2. Build (`go build -o exporter main.go`)
3. Configure
    - See `config-example.json` for default values
    - Save changed version as `config.json` in working directory
4. Run (`./exporter`)
    - Or `./exporter -config ~/config.json` if the config is somewhere else.
5. (Optional) You can use the `example-systemd.service` file to create the service.
    - Make sure you edit the placeholder values to fit your setup!

### Queries

- Total distance traveled for all user made vessels
   - `lmp_vessel_distance_travelled{type!="SpaceObject"}`
- Altitude of all planes named "Mallard"
   - `lmp_vessel_alt{type="Plane",vessel="Mallard"}`

### Example response

```
# HELP lmp_current_players The number of currently active players.
# TYPE lmp_current_players gauge
# HELP lmp_subspaces_total The number of subspaces.
# TYPE lmp_subspaces_total gauge
# HELP lmp_vessels_total The number of vessels.
# TYPE lmp_vessels_total gauge
# HELP lmp_subspaces_time The time of subspaces.
# TYPE lmp_subspaces_time gauge
# HELP lmp_vessels_distance_travelled The total distance travelled by a vessel.
# TYPE lmp_vessels_distance_travelled counter
# HELP lmp_vessels_lat The latitude of a vessel.
# TYPE lmp_vessels_lat gauge
# HELP lmp_vessels_lon The longitude of a vessel.
# TYPE lmp_vessels_lon gauge
# HELP lmp_vessels_alt The altitude of a vessel.
# TYPE lmp_vessels_alt gauge
# HELP lmp_vessels_epoch The epoch of a vessel.
# TYPE lmp_vessels_epoch gauge
# HELP lmp_vessels_reference_body The reference body of a vessel.
# TYPE lmp_vessels_reference_body gauge
# HELP lmp_scrapes_duration The duration of a the metrics collection in nanoseconds.
# TYPE lmp_scrapes_duration gauge
# HELP lmp_vessel_semi_major_axis The semi major axis of a vessel.
# TYPE lmp_vessel_semi_major_axis gauge
# HELP lmp_vessel_eccentricity The eccentricity of a vessel.
# TYPE lmp_vessel_eccentricity gauge
# HELP lmp_vessel_inclination The inclination of a vessel.
# TYPE lmp_vessel_inclination gauge
# HELP lmp_vessel_argument_of_periapsis A vessels argument of periapsis.
# TYPE lmp_vessel_argument_of_periapsis gauge
# HELP lmp_vessel_lon_of_ascending_node A vessels longitude of ascending node.
# TYPE lmp_vessel_lon_of_ascending_node gauge
# HELP lmp_vessel_mean_anomaly A vessels mean anomaly.
# TYPE lmp_vessel_mean_anomaly gauge
# Start of metrics for 127.0.0.1:8900
lmp_current_players{server="127.0.0.1:8900"} 1
lmp_subspaces_total{server="127.0.0.1:8900"} 1
lmp_subspaces_time{server="127.0.0.1:8900",subspace="6"} 1048.672573
lmp_vessel_distance_travelled{server="127.0.0.1:8900",vessel="Mallard",type="Plane",vesselId="3594dc63-a212-44ef-814f-cba4b2684405"} 0.236290
lmp_vessel_lat{server="127.0.0.1:8900",vessel="Mallard",type="Plane",vesselId="3594dc63-a212-44ef-814f-cba4b2684405"} 0.236290
lmp_vessel_lon{server="127.0.0.1:8900",vessel="Mallard",type="Plane",vesselId="3594dc63-a212-44ef-814f-cba4b2684405"} -74.616618
lmp_vessel_alt{server="127.0.0.1:8900",vessel="Mallard",type="Plane",vesselId="3594dc63-a212-44ef-814f-cba4b2684405"} 0.236290
lmp_vessel_semi_major_axis{server="127.0.0.1:8900",vessel="Mallard",type="Plane",vesselId="3594dc63-a212-44ef-814f-cba4b2684405"} 300817.710977
lmp_vessel_eccentricity{server="127.0.0.1:8900",vessel="Mallard",type="Plane",vesselId="3594dc63-a212-44ef-814f-cba4b2684405"} 0.994798
lmp_vessel_inclination{server="127.0.0.1:8900",vessel="Mallard",type="Plane",vesselId="3594dc63-a212-44ef-814f-cba4b2684405"} 0.058288
lmp_vessel_argument_of_periapsis{server="127.0.0.1:8900",vessel="Mallard",type="Plane",vesselId="3594dc63-a212-44ef-814f-cba4b2684405"} 89.694956
lmp_vessel_lon_of_ascending_node{server="127.0.0.1:8900",vessel="Mallard",type="Plane",vesselId="3594dc63-a212-44ef-814f-cba4b2684405"} 304.264365
lmp_vessel_mean_anomaly{server="127.0.0.1:8900",vessel="Mallard",type="Plane",vesselId="3594dc63-a212-44ef-814f-cba4b2684405"} 3.141593
lmp_vessel_epoch{server="127.0.0.1:8900",vessel="Mallard",type="Plane",vesselId="3594dc63-a212-44ef-814f-cba4b2684405"} 33436.084887
lmp_vessel_reference_body{server="127.0.0.1:8900",vessel="Mallard",type="Plane",vesselId="3594dc63-a212-44ef-814f-cba4b2684405"} 1
lmp_vessel_distance_travelled{server="127.0.0.1:8900",vessel="Ast. YRJ-552",type="SpaceObject",vesselId="a63703b2-ba7f-4e27-a926-f603d9f786ce"} 0.000000
lmp_vessel_lat{server="127.0.0.1:8900",vessel="Ast. YRJ-552",type="SpaceObject",vesselId="a63703b2-ba7f-4e27-a926-f603d9f786ce"} 0.000000
lmp_vessel_lon{server="127.0.0.1:8900",vessel="Ast. YRJ-552",type="SpaceObject",vesselId="a63703b2-ba7f-4e27-a926-f603d9f786ce"} -172.679371
lmp_vessel_alt{server="127.0.0.1:8900",vessel="Ast. YRJ-552",type="SpaceObject",vesselId="a63703b2-ba7f-4e27-a926-f603d9f786ce"} 0.000000
lmp_vessel_semi_major_axis{server="127.0.0.1:8900",vessel="Ast. YRJ-552",type="SpaceObject",vesselId="a63703b2-ba7f-4e27-a926-f603d9f786ce"} 14299805007.977070
lmp_vessel_eccentricity{server="127.0.0.1:8900",vessel="Ast. YRJ-552",type="SpaceObject",vesselId="a63703b2-ba7f-4e27-a926-f603d9f786ce"} 0.049691
lmp_vessel_inclination{server="127.0.0.1:8900",vessel="Ast. YRJ-552",type="SpaceObject",vesselId="a63703b2-ba7f-4e27-a926-f603d9f786ce"} 0.636341
lmp_vessel_argument_of_periapsis{server="127.0.0.1:8900",vessel="Ast. YRJ-552",type="SpaceObject",vesselId="a63703b2-ba7f-4e27-a926-f603d9f786ce"} 8.971307
lmp_vessel_lon_of_ascending_node{server="127.0.0.1:8900",vessel="Ast. YRJ-552",type="SpaceObject",vesselId="a63703b2-ba7f-4e27-a926-f603d9f786ce"} 323.743558
lmp_vessel_mean_anomaly{server="127.0.0.1:8900",vessel="Ast. YRJ-552",type="SpaceObject",vesselId="a63703b2-ba7f-4e27-a926-f603d9f786ce"} -0.176219
lmp_vessel_epoch{server="127.0.0.1:8900",vessel="Ast. YRJ-552",type="SpaceObject",vesselId="a63703b2-ba7f-4e27-a926-f603d9f786ce"} 3623553.428192
lmp_vessel_reference_body{server="127.0.0.1:8900",vessel="Ast. YRJ-552",type="SpaceObject",vesselId="a63703b2-ba7f-4e27-a926-f603d9f786ce"} 0
lmp_vessel_distance_travelled{server="127.0.0.1:8900",vessel="Ast. LVB-070",type="SpaceObject",vesselId="dce1a466-e45c-4331-b538-daae24eebfa7"} 0.000000
lmp_vessel_lat{server="127.0.0.1:8900",vessel="Ast. LVB-070",type="SpaceObject",vesselId="dce1a466-e45c-4331-b538-daae24eebfa7"} 0.000000
lmp_vessel_lon{server="127.0.0.1:8900",vessel="Ast. LVB-070",type="SpaceObject",vesselId="dce1a466-e45c-4331-b538-daae24eebfa7"} -169.494839
lmp_vessel_alt{server="127.0.0.1:8900",vessel="Ast. LVB-070",type="SpaceObject",vesselId="dce1a466-e45c-4331-b538-daae24eebfa7"} 0.000000
lmp_vessel_semi_major_axis{server="127.0.0.1:8900",vessel="Ast. LVB-070",type="SpaceObject",vesselId="dce1a466-e45c-4331-b538-daae24eebfa7"} 14503029014.023754
lmp_vessel_eccentricity{server="127.0.0.1:8900",vessel="Ast. LVB-070",type="SpaceObject",vesselId="dce1a466-e45c-4331-b538-daae24eebfa7"} 0.065090
lmp_vessel_inclination{server="127.0.0.1:8900",vessel="Ast. LVB-070",type="SpaceObject",vesselId="dce1a466-e45c-4331-b538-daae24eebfa7"} 0.154954
lmp_vessel_argument_of_periapsis{server="127.0.0.1:8900",vessel="Ast. LVB-070",type="SpaceObject",vesselId="dce1a466-e45c-4331-b538-daae24eebfa7"} 234.511164
lmp_vessel_lon_of_ascending_node{server="127.0.0.1:8900",vessel="Ast. LVB-070",type="SpaceObject",vesselId="dce1a466-e45c-4331-b538-daae24eebfa7"} 107.414807
lmp_vessel_mean_anomaly{server="127.0.0.1:8900",vessel="Ast. LVB-070",type="SpaceObject",vesselId="dce1a466-e45c-4331-b538-daae24eebfa7"} 0.099768
lmp_vessel_epoch{server="127.0.0.1:8900",vessel="Ast. LVB-070",type="SpaceObject",vesselId="dce1a466-e45c-4331-b538-daae24eebfa7"} 4303778.489203
lmp_vessel_reference_body{server="127.0.0.1:8900",vessel="Ast. LVB-070",type="SpaceObject",vesselId="dce1a466-e45c-4331-b538-daae24eebfa7"} 0
lmp_vessel_distance_travelled{server="127.0.0.1:8900",vessel="Ast. GAH-580",type="SpaceObject",vesselId="1ac02957-3af7-40a3-a7c2-edff232a294d"} 0.000000
lmp_vessel_lat{server="127.0.0.1:8900",vessel="Ast. GAH-580",type="SpaceObject",vesselId="1ac02957-3af7-40a3-a7c2-edff232a294d"} 0.000000
lmp_vessel_lon{server="127.0.0.1:8900",vessel="Ast. GAH-580",type="SpaceObject",vesselId="1ac02957-3af7-40a3-a7c2-edff232a294d"} 177.228928
lmp_vessel_alt{server="127.0.0.1:8900",vessel="Ast. GAH-580",type="SpaceObject",vesselId="1ac02957-3af7-40a3-a7c2-edff232a294d"} 0.000000
lmp_vessel_semi_major_axis{server="127.0.0.1:8900",vessel="Ast. GAH-580",type="SpaceObject",vesselId="1ac02957-3af7-40a3-a7c2-edff232a294d"} 16313628493.522465
lmp_vessel_eccentricity{server="127.0.0.1:8900",vessel="Ast. GAH-580",type="SpaceObject",vesselId="1ac02957-3af7-40a3-a7c2-edff232a294d"} 0.169958
lmp_vessel_inclination{server="127.0.0.1:8900",vessel="Ast. GAH-580",type="SpaceObject",vesselId="1ac02957-3af7-40a3-a7c2-edff232a294d"} 1.788519
lmp_vessel_argument_of_periapsis{server="127.0.0.1:8900",vessel="Ast. GAH-580",type="SpaceObject",vesselId="1ac02957-3af7-40a3-a7c2-edff232a294d"} 355.796559
lmp_vessel_lon_of_ascending_node{server="127.0.0.1:8900",vessel="Ast. GAH-580",type="SpaceObject",vesselId="1ac02957-3af7-40a3-a7c2-edff232a294d"} 244.032485
lmp_vessel_mean_anomaly{server="127.0.0.1:8900",vessel="Ast. GAH-580",type="SpaceObject",vesselId="1ac02957-3af7-40a3-a7c2-edff232a294d"} 0.040269
lmp_vessel_epoch{server="127.0.0.1:8900",vessel="Ast. GAH-580",type="SpaceObject",vesselId="1ac02957-3af7-40a3-a7c2-edff232a294d"} 1619682.961080
lmp_vessel_reference_body{server="127.0.0.1:8900",vessel="Ast. GAH-580",type="SpaceObject",vesselId="1ac02957-3af7-40a3-a7c2-edff232a294d"} 0
lmp_vessel_distance_travelled{server="127.0.0.1:8900",vessel="Butterfly Rover",type="Lander",vesselId="3658c51c-4aed-469a-8085-9b31bcbdead5"} 8.043919
lmp_vessel_lat{server="127.0.0.1:8900",vessel="Butterfly Rover",type="Lander",vesselId="3658c51c-4aed-469a-8085-9b31bcbdead5"} 8.043919
lmp_vessel_lon{server="127.0.0.1:8900",vessel="Butterfly Rover",type="Lander",vesselId="3658c51c-4aed-469a-8085-9b31bcbdead5"} -74.685304
lmp_vessel_alt{server="127.0.0.1:8900",vessel="Butterfly Rover",type="Lander",vesselId="3658c51c-4aed-469a-8085-9b31bcbdead5"} 8.043919
lmp_vessel_semi_major_axis{server="127.0.0.1:8900",vessel="Butterfly Rover",type="Lander",vesselId="3658c51c-4aed-469a-8085-9b31bcbdead5"} 300815.140074
lmp_vessel_eccentricity{server="127.0.0.1:8900",vessel="Butterfly Rover",type="Lander",vesselId="3658c51c-4aed-469a-8085-9b31bcbdead5"} 0.994799
lmp_vessel_inclination{server="127.0.0.1:8900",vessel="Butterfly Rover",type="Lander",vesselId="3658c51c-4aed-469a-8085-9b31bcbdead5"} 0.061259
lmp_vessel_argument_of_periapsis{server="127.0.0.1:8900",vessel="Butterfly Rover",type="Lander",vesselId="3658c51c-4aed-469a-8085-9b31bcbdead5"} 90.000000
lmp_vessel_lon_of_ascending_node{server="127.0.0.1:8900",vessel="Butterfly Rover",type="Lander",vesselId="3658c51c-4aed-469a-8085-9b31bcbdead5"} 303.894639
lmp_vessel_mean_anomaly{server="127.0.0.1:8900",vessel="Butterfly Rover",type="Lander",vesselId="3658c51c-4aed-469a-8085-9b31bcbdead5"} 3.141593
lmp_vessel_epoch{server="127.0.0.1:8900",vessel="Butterfly Rover",type="Lander",vesselId="3658c51c-4aed-469a-8085-9b31bcbdead5"} 33436.304887
lmp_vessel_reference_body{server="127.0.0.1:8900",vessel="Butterfly Rover",type="Lander",vesselId="3658c51c-4aed-469a-8085-9b31bcbdead5"} 1
lmp_vessel_distance_travelled{server="127.0.0.1:8900",vessel="Chungus One",type="Ship",vesselId="7d7186f4-f326-45f8-b0cc-8849eecfb70f"} 1101489.535562
lmp_vessel_lat{server="127.0.0.1:8900",vessel="Chungus One",type="Ship",vesselId="7d7186f4-f326-45f8-b0cc-8849eecfb70f"} 1101489.535562
lmp_vessel_lon{server="127.0.0.1:8900",vessel="Chungus One",type="Ship",vesselId="7d7186f4-f326-45f8-b0cc-8849eecfb70f"} -119.048252
lmp_vessel_alt{server="127.0.0.1:8900",vessel="Chungus One",type="Ship",vesselId="7d7186f4-f326-45f8-b0cc-8849eecfb70f"} 1101489.535562
lmp_vessel_semi_major_axis{server="127.0.0.1:8900",vessel="Chungus One",type="Ship",vesselId="7d7186f4-f326-45f8-b0cc-8849eecfb70f"} 300782.140681
lmp_vessel_eccentricity{server="127.0.0.1:8900",vessel="Chungus One",type="Ship",vesselId="7d7186f4-f326-45f8-b0cc-8849eecfb70f"} 0.994801
lmp_vessel_inclination{server="127.0.0.1:8900",vessel="Chungus One",type="Ship",vesselId="7d7186f4-f326-45f8-b0cc-8849eecfb70f"} 0.586146
lmp_vessel_argument_of_periapsis{server="127.0.0.1:8900",vessel="Chungus One",type="Ship",vesselId="7d7186f4-f326-45f8-b0cc-8849eecfb70f"} 90.009576
lmp_vessel_lon_of_ascending_node{server="127.0.0.1:8900",vessel="Chungus One",type="Ship",vesselId="7d7186f4-f326-45f8-b0cc-8849eecfb70f"} 20.030914
lmp_vessel_mean_anomaly{server="127.0.0.1:8900",vessel="Chungus One",type="Ship",vesselId="7d7186f4-f326-45f8-b0cc-8849eecfb70f"} 3.141590
lmp_vessel_epoch{server="127.0.0.1:8900",vessel="Chungus One",type="Ship",vesselId="7d7186f4-f326-45f8-b0cc-8849eecfb70f"} 105298.197743
lmp_vessel_reference_body{server="127.0.0.1:8900",vessel="Chungus One",type="Ship",vesselId="7d7186f4-f326-45f8-b0cc-8849eecfb70f"} 1
lmp_vessel_distance_travelled{server="127.0.0.1:8900",vessel="Ast. NBB-041",type="SpaceObject",vesselId="d482b204-1773-484d-aaec-e7fd14075098"} 0.000000
lmp_vessel_lat{server="127.0.0.1:8900",vessel="Ast. NBB-041",type="SpaceObject",vesselId="d482b204-1773-484d-aaec-e7fd14075098"} 0.000000
lmp_vessel_lon{server="127.0.0.1:8900",vessel="Ast. NBB-041",type="SpaceObject",vesselId="d482b204-1773-484d-aaec-e7fd14075098"} -177.251777
lmp_vessel_alt{server="127.0.0.1:8900",vessel="Ast. NBB-041",type="SpaceObject",vesselId="d482b204-1773-484d-aaec-e7fd14075098"} 0.000000
lmp_vessel_semi_major_axis{server="127.0.0.1:8900",vessel="Ast. NBB-041",type="SpaceObject",vesselId="d482b204-1773-484d-aaec-e7fd14075098"} 15410250575.839464
lmp_vessel_eccentricity{server="127.0.0.1:8900",vessel="Ast. NBB-041",type="SpaceObject",vesselId="d482b204-1773-484d-aaec-e7fd14075098"} 0.116716
lmp_vessel_inclination{server="127.0.0.1:8900",vessel="Ast. NBB-041",type="SpaceObject",vesselId="d482b204-1773-484d-aaec-e7fd14075098"} 0.837490
lmp_vessel_argument_of_periapsis{server="127.0.0.1:8900",vessel="Ast. NBB-041",type="SpaceObject",vesselId="d482b204-1773-484d-aaec-e7fd14075098"} 177.625435
lmp_vessel_lon_of_ascending_node{server="127.0.0.1:8900",vessel="Ast. NBB-041",type="SpaceObject",vesselId="d482b204-1773-484d-aaec-e7fd14075098"} 101.775970
lmp_vessel_mean_anomaly{server="127.0.0.1:8900",vessel="Ast. NBB-041",type="SpaceObject",vesselId="d482b204-1773-484d-aaec-e7fd14075098"} 0.049894
lmp_vessel_epoch{server="127.0.0.1:8900",vessel="Ast. NBB-041",type="SpaceObject",vesselId="d482b204-1773-484d-aaec-e7fd14075098"} 2640098.690836
lmp_vessel_reference_body{server="127.0.0.1:8900",vessel="Ast. NBB-041",type="SpaceObject",vesselId="d482b204-1773-484d-aaec-e7fd14075098"} 0
lmp_vessel_distance_travelled{server="127.0.0.1:8900",vessel="Ast. GEZ-472",type="SpaceObject",vesselId="c749bf18-13fc-4819-a88b-f5ff2d5364c9"} 0.000000
lmp_vessel_lat{server="127.0.0.1:8900",vessel="Ast. GEZ-472",type="SpaceObject",vesselId="c749bf18-13fc-4819-a88b-f5ff2d5364c9"} 0.000000
lmp_vessel_lon{server="127.0.0.1:8900",vessel="Ast. GEZ-472",type="SpaceObject",vesselId="c749bf18-13fc-4819-a88b-f5ff2d5364c9"} -168.074686
lmp_vessel_alt{server="127.0.0.1:8900",vessel="Ast. GEZ-472",type="SpaceObject",vesselId="c749bf18-13fc-4819-a88b-f5ff2d5364c9"} 0.000000
lmp_vessel_semi_major_axis{server="127.0.0.1:8900",vessel="Ast. GEZ-472",type="SpaceObject",vesselId="c749bf18-13fc-4819-a88b-f5ff2d5364c9"} 14242418943.475710
lmp_vessel_eccentricity{server="127.0.0.1:8900",vessel="Ast. GEZ-472",type="SpaceObject",vesselId="c749bf18-13fc-4819-a88b-f5ff2d5364c9"} 0.045531
lmp_vessel_inclination{server="127.0.0.1:8900",vessel="Ast. GEZ-472",type="SpaceObject",vesselId="c749bf18-13fc-4819-a88b-f5ff2d5364c9"} 0.312098
lmp_vessel_argument_of_periapsis{server="127.0.0.1:8900",vessel="Ast. GEZ-472",type="SpaceObject",vesselId="c749bf18-13fc-4819-a88b-f5ff2d5364c9"} 202.463913
lmp_vessel_lon_of_ascending_node{server="127.0.0.1:8900",vessel="Ast. GEZ-472",type="SpaceObject",vesselId="c749bf18-13fc-4819-a88b-f5ff2d5364c9"} 166.613522
lmp_vessel_mean_anomaly{server="127.0.0.1:8900",vessel="Ast. GEZ-472",type="SpaceObject",vesselId="c749bf18-13fc-4819-a88b-f5ff2d5364c9"} -0.172272
lmp_vessel_epoch{server="127.0.0.1:8900",vessel="Ast. GEZ-472",type="SpaceObject",vesselId="c749bf18-13fc-4819-a88b-f5ff2d5364c9"} 4559566.068967
lmp_vessel_reference_body{server="127.0.0.1:8900",vessel="Ast. GEZ-472",type="SpaceObject",vesselId="c749bf18-13fc-4819-a88b-f5ff2d5364c9"} 0
lmp_scrape_success{server="127.0.0.1:8900"} 1
lmp_scrape_duration{server="127.0.0.1:8900"} 3334229
```
