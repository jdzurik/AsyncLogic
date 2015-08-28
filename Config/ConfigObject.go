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
	Data dataObj
	Events eventsObj
	Index indexObj
	Security securityObj
	Logic logicObj
	Log logObj
	Schema schemaObj
	Replication replicationObj
	Shards shardsObj
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
	IPAddreses []string 
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
	var cnfgName = "config.json"
	/*if _, err := os.Stat(cnfgName); err != nil {
    	fmt.Printf("no file exists; processing...")
    	setDefalultConfig()
	}
	*/
	configFile, err := os.Open(cnfgName)
    if err != nil {
        fmt.Println("Error opening config file ")
    }

    jsonParser := json.NewDecoder(configFile)
    if err = jsonParser.Decode(&config); err != nil {
		fmt.Println(configFile);
        fmt.Println("Error parsing config file " + err.Error() )
    }
	configFile.Close()
	return 
}

func setDefalultConfig(){
	
	var cnfg = configObj{
	    GlobalName : "AsyncLogic",
	    LocalGroupName : "LocalLogic",
	    Node : {
	        Name : "Node1",
	        NodeID : 0,
	        ActorType : "Master",
	    },
	    RootPath: "here",
	    Admin : {
	        Port: "6868",
			IPAddreses : []string{"127.0.0.1" },
	    },
	    Data : {
	        Port: "6869",
	        IPAddreses: []string{ "127.0.0.1" },
	        Paths: []string{ "C:\\AsyncData","C:\\AsyncData" },
	        AutoID: true,
	        AutoIndex: false,
	        AutoReplicate: true,
	    },
	    Events : {
	        SleepTime: 1000,
	        Path: "[root]/Events/",
	    },
	    Index : {
	        Port: "6870",
	        IPAddreses: []string{ "127.0.0.1" },
	        Path: "[root]/Index/",
	    },
	    Security : {
	        Users: []userObj{
	            {
	                Name: "Admin",
	                Salt: "ndk#ts32mx87",
	                Password: "",
	            },
	        },
	        Communications : {
	            RequreSSL: "",
	            CertLocation: "",
	        },
	        Encryption: {
	            Enabled: true,
	            UserData: {
	                Enabled: false,
	                Key: "kdsfghufhv",
	            },
	            Data: {
	                Enabled: false,
	                Key: "kdsfghufhv",
	            },
	        },
	        Path: "[root]/Security/",
	    },
	    Logic: {
	        Path: "[root]/Logic/",
	    },
	    Log: {
	        Errors: true,
	        UserAccess: true,
	        LogicChange: false,
	        DataChange: false,
	        UserAccess: false,
	        Path: "[root]/Log/",
	    },
	    Schema: { 
			Path: "[root]/Schema/" ,
		},
	    Replication: {
	        AcceptData: true,
	        Replicas: 2,
	        Level: "Global",
	        EnableDiscovery: true,
	    },
	    Shards: {
	        Port: "6886",
	        EnableDiscovery: true,
	    },
	}
}
