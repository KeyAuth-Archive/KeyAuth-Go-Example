package auth

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"os/user"
)

//* App Information --> KEEP ALL CLEAR <-- *//
var NumUsers string = ""
var NumOnlineUsers string = ""
var NumKeys string = ""
var CustomerPanelLink string = ""

//* SessionID --> KEEP CLEAR <--
var Session_id string = "lol"

//* KeyAuth Application Detail Storage from main.go --> KEEP ALL CLEAR <--
var name string = ""
var ownerid string = ""
var version string = ""

//* Logged in User Details --> KEEP ALL CLEAR <--
var Username string = ""
var Ip string = ""
var Hwid string = ""
var Createdate string = ""
var Lastlogin string = ""
var Subscription string = ""

//* Intializing Status keep false *//
var Initialized bool = false

//* DON'T CHANGE THESE
var apiUrl string = "https://keyauth.win/"
var resource string = "api/1.1/"

func Api(APPname string, APPownerid string, APPversion string) {
	if APPname == "" || APPownerid == "" || APPversion == "" {
		error("\n\n Application not setupped correctly")
	}

	name = APPname
	ownerid = APPownerid
	version = APPversion

	return
}

func Init() {
	if CheckIFEmpty() {
		error(" \n\n Application not setupped correctly")
	}

	//* Request Struct *//
	type AutoGenerated struct {
		Success   bool   `json:"success"`
		Message   string `json:"message"`
		Sessionid string `json:"sessionid"`
		Appinfo   struct {
			NumUsers          string `json:"numUsers"`
			NumOnlineUsers    string `json:"numOnlineUsers"`
			NumKeys           string `json:"numKeys"`
			Version           string `json:"version"`
			CustomerPanelLink string `json:"customerPanelLink"`
		} `json:"appinfo"`
	}
	data := url.Values{}
	data.Set("type", "init")
	data.Add("ver", version)
	data.Add("name", name)
	data.Add("ownerid", ownerid)

	//* BETA 0.02 Web Request *//
	u, _ := url.ParseRequestURI(apiUrl)
	u.Path = resource
	u.RawQuery = data.Encode()
	urlStr := fmt.Sprintf("%v", u)

	client := &http.Client{}
	r, _ := http.NewRequest("POST", urlStr, nil)
	r.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	resp, _ := client.Do(r)
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	var result AutoGenerated
	if err := json.Unmarshal(body, &result); err != nil {
		error("Something went wrong")
	}
	if err != nil {
	}
	//* BETA 0.02 Web Request *//

	if result.Success == true {
		Initialized = true
		Session_id = result.Sessionid

		NumUsers = result.Appinfo.NumUsers
		NumOnlineUsers = result.Appinfo.NumOnlineUsers
		NumKeys = result.Appinfo.NumKeys
		CustomerPanelLink = result.Appinfo.CustomerPanelLink
	} else {
		error(" Error: " + result.Message)
	}
}

func Login(username string, password string) {
	if Initialized == false {
		error("Please initzalize first")
	}

	var hwid string = GetHwid()

	//* Response Struct*//
	type AutoGenerated struct {
		Success bool   `json:"success"`
		Message string `json:"message"`
		Info    struct {
			Username      string `json:"username"`
			Subscriptions []struct {
				Subscription string      `json:"subscription"`
				Key          interface{} `json:"key"`
				Expiry       string      `json:"expiry"`
				Timeleft     int         `json:"timeleft"`
			} `json:"subscriptions"`
			IP         string `json:"ip"`
			Hwid       string `json:"hwid"`
			Createdate string `json:"createdate"`
			Lastlogin  string `json:"lastlogin"`
		} `json:"info"`
	}
	//* Response Struct*//

	data := url.Values{}
	data.Set("type", "login")
	data.Add("username", username)
	data.Add("pass", password)
	data.Add("hwid", hwid)
	data.Add("sessionid", Session_id)
	data.Add("name", name)
	data.Add("ownerid", ownerid)

	//* BETA 0.02 Web Request *//
	u, _ := url.ParseRequestURI(apiUrl)
	u.Path = resource
	u.RawQuery = data.Encode()
	urlStr := fmt.Sprintf("%v", u)

	client := &http.Client{}
	r, _ := http.NewRequest("POST", urlStr, nil)
	r.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	resp, _ := client.Do(r)
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	var result AutoGenerated
	if err := json.Unmarshal(body, &result); err != nil {
		error("Something went wrong")
	}
	if err != nil {
	}
	//* BETA 0.02 Web Request *//
	if result.Success == true {
		Username = result.Info.Username
		Ip = result.Info.IP
		Hwid = result.Info.Hwid
		Createdate = result.Info.Createdate
		Lastlogin = result.Info.Lastlogin
		Subscription = result.Info.Subscriptions[0].Subscription
	} else {
		error("\n\n Error: " + result.Message)
	}

	return
}

func Register(username string, password string, key string) {
	if Initialized == false {
		error("Please initzalize first")
	}

	var hwid string = GetHwid()

	//* Response Struct*//
	type AutoGenerated struct {
		Success bool   `json:"success"`
		Message string `json:"message"`
		Info    struct {
			Username      string `json:"username"`
			Subscriptions []struct {
				Subscription string      `json:"subscription"`
				Key          interface{} `json:"key"`
				Expiry       string      `json:"expiry"`
				Timeleft     int         `json:"timeleft"`
			} `json:"subscriptions"`
			IP         string `json:"ip"`
			Hwid       string `json:"hwid"`
			Createdate string `json:"createdate"`
			Lastlogin  string `json:"lastlogin"`
		} `json:"info"`
	}
	//* Response Struct*//

	data := url.Values{}
	data.Set("type", "register")
	data.Add("username", username)
	data.Add("pass", password)
	data.Add("key", key)
	data.Add("hwid", hwid)
	data.Add("sessionid", Session_id)
	data.Add("name", name)
	data.Add("ownerid", ownerid)

	//* BETA 0.02 Web Request *//
	u, _ := url.ParseRequestURI(apiUrl)
	u.Path = resource
	u.RawQuery = data.Encode()
	urlStr := fmt.Sprintf("%v", u)

	client := &http.Client{}
	r, _ := http.NewRequest("POST", urlStr, nil)
	r.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	resp, _ := client.Do(r)
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	var result AutoGenerated
	if err := json.Unmarshal(body, &result); err != nil {
		error("Something went wrong")
	}
	if err != nil {
	}
	//* BETA 0.02 Web Request *//

	if result.Success == true {
		Username = result.Info.Username
		Ip = result.Info.IP
		Hwid = result.Info.Hwid
		Createdate = result.Info.Createdate
		Lastlogin = result.Info.Lastlogin
		Subscription = result.Info.Subscriptions[0].Subscription
	} else {
		error("\n\n Error: " + result.Message)
	}

	return
}

func Upgrade(username string, key string) {
	if Initialized == false {
		error("Please initzalize first")
	}

	//* Response Struct*//
	type AutoGenerated struct {
		Success bool   `json:"success"`
		Message string `json:"message"`
	}
	//* Response Struct*//

	data := url.Values{}
	data.Set("type", "upgrade")
	data.Add("username", username)
	data.Add("key", key)
	data.Add("sessionid", Session_id)
	data.Add("name", name)
	data.Add("ownerid", ownerid)

	//* BETA 0.02 Web Request *//
	u, _ := url.ParseRequestURI(apiUrl)
	u.Path = resource
	u.RawQuery = data.Encode()
	urlStr := fmt.Sprintf("%v", u)

	client := &http.Client{}
	r, _ := http.NewRequest("POST", urlStr, nil)
	r.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	resp, _ := client.Do(r)
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	var result AutoGenerated
	if err := json.Unmarshal(body, &result); err != nil {
		error("Something went wrong")
	}
	if err != nil {
	}
	//* BETA 0.02 Web Request *//
	if result.Success == true {
		error("\n " + result.Message)
	} else {
		error("Error: " + result.Message)
	}

	return
}

func License(key string) {
	if Initialized == false {
		error("Please initzalize first")
	}

	var hwid string = GetHwid()

	//* Response Struct*//
	type AutoGenerated struct {
		Success bool   `json:"success"`
		Message string `json:"message"`
		Info    struct {
			Username      string `json:"username"`
			Subscriptions []struct {
				Subscription string `json:"subscription"`
				Key          string `json:"key"`
				Expiry       string `json:"expiry"`
				Timeleft     int    `json:"timeleft"`
			} `json:"subscriptions"`
			IP         string `json:"ip"`
			Hwid       string `json:"hwid"`
			Createdate string `json:"createdate"`
			Lastlogin  string `json:"lastlogin"`
		} `json:"info"`
	}
	//* Response Struct*//

	//* Data Values *//
	data := url.Values{}
	data.Set("type", "license")
	data.Add("key", key)
	data.Add("hwid", hwid)
	data.Add("sessionid", Session_id)
	data.Add("name", name)
	data.Add("ownerid", ownerid)
	//* Data Values *//

	//* BETA 0.02 Web Request *//
	u, _ := url.ParseRequestURI(apiUrl)
	u.Path = resource
	u.RawQuery = data.Encode()
	urlStr := fmt.Sprintf("%v", u)

	client := &http.Client{}
	r, _ := http.NewRequest("POST", urlStr, nil)
	r.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	resp, _ := client.Do(r)
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	var result AutoGenerated
	if err := json.Unmarshal(body, &result); err != nil {
		error("Something went wrong")
	}
	if err != nil {
	}
	//* BETA 0.02 Web Request *//

	if result.Success == true {
		Username = result.Info.Username
		Ip = result.Info.IP
		Hwid = result.Info.Hwid
		Createdate = result.Info.Createdate
		Lastlogin = result.Info.Lastlogin
		Subscription = result.Info.Subscriptions[0].Subscription
	} else {
		error("\n\n Error: " + result.Message)
	}

	return
}

func FetchOnline() string {
	if Initialized == false {
		error("Please initzalize first")
	}

	type AutoGenerated struct {
		Success bool   `json:"success"`
		Message string `json:"message"`
		Users   struct {
			Credential string `json:"credential"`
		} `json:"users"`
	}

	//* Response Struct*//

	//* Data Values *//
	data := url.Values{}
	data.Set("type", "fetchOnline")
	data.Add("sessionid", Session_id)
	data.Add("name", name)
	data.Add("ownerid", ownerid)
	//* Data Values *//

	//* BETA 0.02 Web Request *//
	u, _ := url.ParseRequestURI(apiUrl)
	u.Path = resource
	u.RawQuery = data.Encode()
	urlStr := fmt.Sprintf("%v", u)

	client := &http.Client{}
	r, _ := http.NewRequest("POST", urlStr, nil)
	r.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	resp, _ := client.Do(r)
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	var result AutoGenerated
	if err := json.Unmarshal(body, &result); err != nil {
		error("Something went wrong")
	}
	if err != nil {
	}

	if result.Success == true {
		/**/
	} else {
		return ""
	}
}

func Check() bool {
	if Initialized == false {
		error("Please initzalize first")
	}

	//* Response Struct*//
	type AutoGenerated struct {
		Success bool   `json:"success"`
		Message string `json:"message"`
	}
	//* Response Struct*//

	//* Data Values *//
	data := url.Values{}
	data.Set("type", "check")
	data.Add("sessionid", Session_id)
	data.Add("name", name)
	data.Add("ownerid", ownerid)
	//* Data Values *//

	//* BETA 0.02 Web Request *//
	u, _ := url.ParseRequestURI(apiUrl)
	u.Path = resource
	u.RawQuery = data.Encode()
	urlStr := fmt.Sprintf("%v", u)

	client := &http.Client{}
	r, _ := http.NewRequest("POST", urlStr, nil)
	r.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	resp, _ := client.Do(r)
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	var result AutoGenerated
	if err := json.Unmarshal(body, &result); err != nil {
		error("Something went wrong")
	}
	if err != nil {
	}

	if result.Success {
		return true
	} else {
		return false
	}
}

func SetVar(varname string, vardata string) {
	if Initialized == false {
		error("Please initzalize first")
	}

	//* Response Struct*//
	type AutoGenerated struct {
		Success bool   `json:"success"`
		Message string `json:"message"`
	}
	//* Response Struct*//

	//* Data Values *//
	data := url.Values{}
	data.Set("type", "setvar")
	data.Add("var", varname)
	data.Add("data", vardata)
	data.Add("sessionid", Session_id)
	data.Add("name", name)
	data.Add("ownerid", ownerid)
	//* Data Values *//

	//* BETA 0.02 Web Request *//
	u, _ := url.ParseRequestURI(apiUrl)
	u.Path = resource
	u.RawQuery = data.Encode()
	urlStr := fmt.Sprintf("%v", u)

	client := &http.Client{}
	r, _ := http.NewRequest("POST", urlStr, nil)
	r.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	resp, _ := client.Do(r)
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	var result AutoGenerated
	if err := json.Unmarshal(body, &result); err != nil {
		error("Something went wrong")
	}
	if err != nil {
	}

	if result.Success == true {
		fmt.Println("\n Status:", result.Message)
	} else {
		error("\n Error: " + result.Message)
	}

	return
}

func GetVar(varname string) string {
	if Initialized == false {
		error("Please initzalize first")
	}

	//* Response Struct*//
	type AutoGenerated struct {
		Success  bool   `json:"success"`
		Message  string `json:"message"`
		Response string `json:"response"`
	}
	//* Response Struct*//

	//* Data Values *//
	data := url.Values{}
	data.Set("type", "getvar")
	data.Add("var", varname)
	data.Add("sessionid", Session_id)
	data.Add("name", name)
	data.Add("ownerid", ownerid)
	//* Data Values *//

	//* BETA 0.02 Web Request *//
	u, _ := url.ParseRequestURI(apiUrl)
	u.Path = resource
	u.RawQuery = data.Encode()
	urlStr := fmt.Sprintf("%v", u)

	client := &http.Client{}
	r, _ := http.NewRequest("POST", urlStr, nil)
	r.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	resp, _ := client.Do(r)
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	var result AutoGenerated
	if err := json.Unmarshal(body, &result); err != nil {
		error("Something went wrong")
	}
	if err != nil {
	}

	if result.Success == true {
		return result.Response
	} else {
		return ""
	}

}

func Var(varname string) string {
	if Initialized == false {
		error("Please initzalize first")
	}

	//* Response Struct*//
	type AutoGenerated struct {
		Success bool   `json:"success"`
		Message string `json:"message"`
	}
	//* Response Struct*//

	//* Data Values *//
	data := url.Values{}
	data.Set("type", "var")
	data.Add("varid", varname)
	data.Add("sessionid", Session_id)
	data.Add("name", name)
	data.Add("ownerid", ownerid)
	//* Data Values *//

	//* BETA 0.02 Web Request *//
	u, _ := url.ParseRequestURI(apiUrl)
	u.Path = resource
	u.RawQuery = data.Encode()
	urlStr := fmt.Sprintf("%v", u)

	client := &http.Client{}
	r, _ := http.NewRequest("POST", urlStr, nil)
	r.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	resp, _ := client.Do(r)
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	var result AutoGenerated
	if err := json.Unmarshal(body, &result); err != nil {
		error("Something went wrong")
	}
	if err != nil {
	}

	if result.Success == true {
		return result.Message
	} else {
		return ""
	}
}

func CheckBlack() bool {
	if Initialized == false {
		error("Please initzalize first")
	}

	var hwid string = GetHwid()

	//* Response Struct*//
	type AutoGenerated struct {
		Success bool   `json:"success"`
		Message string `json:"message"`
	}
	//* Response Struct*//

	//* Data Values *//
	data := url.Values{}
	data.Set("type", "checkblacklist")
	data.Add("hwid", hwid)
	data.Add("sessionid", Session_id)
	data.Add("name", name)
	data.Add("ownerid", ownerid)
	//* Data Values *//

	//* BETA 0.02 Web Request *//
	u, _ := url.ParseRequestURI(apiUrl)
	u.Path = resource
	u.RawQuery = data.Encode()
	urlStr := fmt.Sprintf("%v", u)

	client := &http.Client{}
	r, _ := http.NewRequest("POST", urlStr, nil)
	r.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	resp, _ := client.Do(r)
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	var result AutoGenerated
	if err := json.Unmarshal(body, &result); err != nil {
		error("Something went wrong")
	}
	if err != nil {
	}

	if result.Success == true {
		return true
	} else {
		return false
	}
}

func Webhook(webid string, param string) string {
	if Initialized == false {
		error("Please initzalize first")
	}

	var WebBody string = ""
	var conttype string = ""

	//* Response Struct*//
	type AutoGenerated struct {
		Success  bool   `json:"success"`
		Message  string `json:"message"`
		Response string `json:"response"`
	}
	//* Response Struct*//

	//* Data Values *//
	data := url.Values{}
	data.Set("type", "webhook")
	data.Add("webid", webid)
	data.Add("params", param)
	data.Add("body", WebBody)
	data.Add("conttype", conttype)
	data.Add("sessionid", Session_id)
	data.Add("name", name)
	data.Add("ownerid", ownerid)
	//* Data Values *//

	//* BETA 0.02 Web Request *//
	u, _ := url.ParseRequestURI(apiUrl)
	u.Path = resource
	u.RawQuery = data.Encode()
	urlStr := fmt.Sprintf("%v", u)

	client := &http.Client{}
	r, _ := http.NewRequest("POST", urlStr, nil)
	r.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	resp, _ := client.Do(r)
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	var result AutoGenerated
	if err := json.Unmarshal(body, &result); err != nil {
		error("Something went wrong")
	}
	if err != nil {
	}

	if result.Success == true {
		return result.Response
	} else {
		return ""
	}
}

func Log(message string) {
	if Initialized == false {
		error("Please initzalize first")
	}

	var PcUser = GetPcName()

	//* Data Values *//
	data := url.Values{}
	data.Set("type", "log")
	data.Add("pcuser", PcUser)
	data.Add("message", message)
	data.Add("sessionid", Session_id)
	data.Add("name", name)
	data.Add("ownerid", ownerid)
	//* Data Values *//

	//* BETA 0.02 Web Request *//
	u, _ := url.ParseRequestURI(apiUrl)
	u.Path = resource
	u.RawQuery = data.Encode()
	urlStr := fmt.Sprintf("%v", u)

	client := &http.Client{}
	r, _ := http.NewRequest("POST", urlStr, nil)
	r.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	resp, _ := client.Do(r)
	defer resp.Body.Close()
}

func GetPcName() string {
	name, err := os.Hostname()
	if err != nil {
		return "UNKNOWN"
	}
	return name
}

func GetHwid() string {
	name, _ := os.Hostname()
	usr, _ := user.Current()

	var hwid = md5Hash(name + usr.Username)
	return hwid
}

func md5Hash(text string) string {
	hasher := md5.New()
	hasher.Write([]byte(text))
	return hex.EncodeToString(hasher.Sum(nil))
}

func error(message string) {
	fmt.Println(message)
	os.Exit(0)
}

func CheckIFEmpty() bool {
	var status bool = false
	if name == "" || ownerid == "" || version == "" {
		status = true
	}
	return status
}

func PrettyPrint(i interface{}) string {
	s, _ := json.MarshalIndent(i, "", "\t")
	return string(s)
}
