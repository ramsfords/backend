package writer

import (
	"io/ioutil"
	"log"
)

func WriteRapidAuthToken(byts []byte) {
	err := ioutil.WriteFile("/Users/surenl/projects/carriers/business/test/auth.txt", byts, 0644)
	if err != nil {
		log.Fatal(err)
	}
}
func WriteToAdminTxt(byts []byte) {
	//write user id to a file so admin can delete user in admin api. for now its logs/admin.txt
	// err := ioutil.WriteFile(configs.GetConfig().LogPath+"admin.txt", byts, 0644)
	// if err != nil {
	// 	log.Fatal(err)
	// }
}
func WriteToTokenTxt(byts []byte) {
	//write user id to a file so admin can delete user in admin api. for now its logs/admin.txt
	// err := ioutil.WriteFile(configs.GetConfig().LogPath+"token.txt", byts, 0644)
	// if err != nil {
	// 	log.Fatal(err)
	// }

}
