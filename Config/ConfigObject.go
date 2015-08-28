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
	Port string
	IPAddreses []string 
}

type dataObj struct {
	Port string
	IPAddreses []string 
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
	Port string
	IPAddreses []string 
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
	RequreSSL string 
	CertLocation string
}

type encryptionObj struct {
	Enabled bool 
	UserKey encryptKeyObj
	DataKey encryptKeyObj
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
	Errors bool
    UserAccess bool
    LogicChange bool
    DataChange bool
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
	Port string
	IPAddreses []string 
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
	if _, err := os.Stat(cnfgName); err != nil {
    	fmt.Printf("no file exists; processing...")
    	setDefalultConfig()
	}
	
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

func setDefalultConfig()(configObj){
	
	var cnfg = configObj{
	    GlobalName : "AsyncLogic",
	    LocalGroupName : "LocalLogic",
	    Node : nodeObj{
	        Name : "Node1",
	        NodeID : 0,
	        ActorType : "Master",
	    },
	    RootPath: "here",
	    Admin : adminObj{
	        Port: "6868",
			IPAddreses : []string{"127.0.0.1" },
	    },
	    Data : dataObj{
	        Port: "6869",
	        IPAddreses: []string{ "127.0.0.1" },
	        Paths: []string{ "C:\\AsyncData","C:\\AsyncData1" },
	        AutoID: true,
	        AutoIndex: false,
	        AutoReplicate: true,
	    },
	    Events : eventsObj{
	        SleepTime: 1000,
	        Path: "[root]/Events/",
	    },
	    Index : indexObj{
	        Port: "6870",
	        IPAddreses: []string{ "127.0.0.1" },
	        Paths: []string{ "C:\\AsyncData\\Index","C:\\AsyncData\\Index1" },
	    },
	    Security : securityObj{
	        Users: []userObj{
	            {
	                Name: "Admin",
	                Salt: "ndk#ts32mx87",
	                Password: "",
	            },
	        },
	        Communications : communicationsObj{
	            RequreSSL: "",
	            CertLocation: "",
	        },
	        Encryption: encryptionObj{
	            Enabled: true,
	            UserKey: encryptKeyObj{
	                Enabled: false,
	                Key: "kdsfghufhv",
	            },
	            DataKey: encryptKeyObj{
	                Enabled: false,
	                Key: "kdsfghufhv",
	            },
	        },
	        Path: "[root]/Security/",
	    },
	    Logic: logicObj{
	        Path: "[root]/Logic/",
	    },
	    Log: logObj{
	        Errors: true,
	        UserAccess: true,
	        LogicChange: false,
	        DataChange: false,
	        Path: "[root]/Log/",
	    },
	    Schema: schemaObj{ 
			Path: "[root]/Schema/" ,
		},
	    Replication: replicationObj{
	        AcceptData: true,
	        Replicas: 2,
	        Level: "Global",
	        EnableDiscovery: true,
	    },
	    Shards: shardsObj{
	        Port: "6886",
	        EnableDiscovery: true,
	    },
	}
	
	return cnfg
}
