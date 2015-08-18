// ConfigController.go
package ConfigObject

import (
	"fmt"
	"encoding/json"
	"os"

//	"net/http"
)

type configObj struct {
	GlobalName string
	LocalGroupName string
	Node nodeObj
	RootPath string
	Admin adminObj
	Data string
	Events string
	Index string
	Security string
	Log string
	Schema string
	Replication string
	Shards string 
}

type serverObj struct {
	Port string
	IPAddreses []string 
}

type nodeObj struct {
	Name string
	NodeID int 
	ActorType string
}

type adminObj struct {
	serverObj 
}

type dataObj struct {
	serverObj
	Paths []string
	AutoID bool 
	AutoIndex bool
	AutoReplicate bool
}

type eventsObj struct {
	SleepTime int
	Path string 
}

type indexObj struct {
	serverObj
	Paths []string
}

type securityObj struct {
	Users []userObj
	Communications communicationsObj
	Encryption encryptionObj
	Path string
}

type userObj struct {
	Name string 
	Salt string
	Password string
}

type communicationsObj struct {
	Name string 
	Salt string
	Password string
}

type encryptionObj struct {
	Enabled bool 
	User encryptKeyObj
	Data encryptKeyObj
	Path string
}

type encryptKeyObj struct {
	Enabled bool 
	Key string
	Path string
}

type logicObj struct {
	Errors bool
	UserAccess bool
	LogicChange bool
	DataChange bool
	Path string
}

type logObj struct {
	Path string
}
type schemaObj struct {
	Path string
}
type replicationObj struct {
	AcceptData bool
	Replicas int
	Level string
	EnableDiscovery bool
}
type shardsObj struct {
	serverObj
	EnableDiscovery bool
	
}

//func handelConfig(w http.ResponseWriter, r *http.Request) {
//	http.Handle("/Config/save", SaveConfig)
//}

//func SaveConfig(w http.ResponseWriter, r *http.Request) {
//	decoder := json.NewDecoder(r.Body)
//	var c config_struct
//	err := decoder.Decode(&c)
//	if err != nil {
//		//cfmt.Fprint(w, "{'Save': '"+err.Error()+"'")
//	}
//	http.ServeFile(w, r, "Admin/Index.htm")
//}

func LoadConfig()(config configObj){
	configFile, err := os.Open("config.json")
    if err != nil {
        fmt.Println("Error opening config file ")
    }

    jsonParser := json.NewDecoder(configFile)
    if err = jsonParser.Decode(&config); err != nil {
        fmt.Println("Error parsing config file " + err.Error())
    }
	configFile.Close()
	return 
}
