package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	birdJson := `{"birds":{"pigeon":"likes to perch on rocks","eagle":"bird of prey"},"animals":"none"}`
	var result map[string]interface{}
	json.Unmarshal([]byte(birdJson), &result)
	fmt.Println(result)
	birds := result["birds"].(map[string]interface{})
	fmt.Println(birds)
for key, value := range birds{
	fmt.Println(key, " - ", value.(string))
}

}

