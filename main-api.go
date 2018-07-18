package main

// func main() {
// 	// url := "https://swapi.co/api/people/1"
// 	// res, err := http.Get(url)
// 	// if err != nil {
// 	// 	fmt.Println(err)
// 	// }

// 	// var body People
// 	// blob, err := ioutil.ReadAll(res.Body)
// 	// json.Unmarshal(blob, &body)

// 	// blob2, _ := json.Marshal(body)
// 	// fmt.Println(string(blob2))

// 	type People struct {
// 		Name      string   `json:"name"`
// 		Height    string   `json:"height"`
// 		Mass      string   `json:"mass"`
// 		HairColor string   `json:"hair_color"`
// 		SkinColor string   `json:"skin_color"`
// 		EyeColor  string   `json:"eye_color"`
// 		BirthYear string   `json:"birth_year"`
// 		Gender    string   `json:"gender"`
// 		Homeworld string   `json:"homeworld"`
// 		Films     []string `json:"films"`
// 		Species   []string `json:"species"`
// 		Vehicles  []string `json:"vehicles"`
// 		Starships []string `json:"starships"`
// 		Created   string   `json:"created"`
// 		Edited    string   `json:"edited"`
// 		URL       string   `json:"url"`
// 	}

// 	bodyData := struct {
// 		Hello string
// 	}{
// 		"world",
// 	}

// 	blobBody, _ := json.Marshal(bodyData)
// 	body := bytes.NewReader(blobBody)
// 	res, err := http.Post("http://192.168.0.244:5001/", "application/json", body)
// 	blob, err := ioutil.ReadAll(res.Body)
// 	if err != nil {
// 		fmt.Println(err)
// 	}
// 	fmt.Println(string(blob))
// }
