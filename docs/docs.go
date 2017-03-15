package docs

import (
	"encoding/json"
	"strings"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/swagger"
)

const (
    Rootinfo string = `{"apiVersion":"1.0.0","swaggerVersion":"1.2","apis":[{"path":"/user","description":"Operations about Users\n"}],"info":{"title":"beego Test API","description":"beego has a very cool tools to autogenerate documents for your API","contact":"astaxie@gmail.com","termsOfServiceUrl":"http://beego.me/","license":"Url http://www.apache.org/licenses/LICENSE-2.0.html"}}`
    Subapi string = `{"/user":{"apiVersion":"1.0.0","swaggerVersion":"1.2","basePath":"","resourcePath":"/user","produces":["application/json","application/xml","text/plain","text/html"],"apis":[{"path":"/signin","description":"","operations":[{"httpMethod":"GET","nickname":"signin","type":"","summary":"user signin","parameters":[{"paramType":"query","name":"username","description":"\"The username for signin\"","dataType":"string","type":"","format":"","allowMultiple":false,"required":true,"minimum":0,"maximum":0},{"paramType":"query","name":"password","description":"\"The password for signin\"","dataType":"string","type":"","format":"","allowMultiple":false,"required":true,"minimum":0,"maximum":0}],"responseMessages":[{"code":100,"message":"signin success","responseModel":""},{"code":101,"message":"signin failed","responseModel":""}]}]},{"path":"/signup","description":"","operations":[{"httpMethod":"POST","nickname":"signup","type":"","summary":"user signup","parameters":[{"paramType":"query","name":"username","description":"\"The username for signup\"","dataType":"string","type":"","format":"","allowMultiple":false,"required":true,"minimum":0,"maximum":0},{"paramType":"query","name":"password","description":"\"The password for signup\"","dataType":"string","type":"","format":"","allowMultiple":false,"required":true,"minimum":0,"maximum":0}],"responseMessages":[{"code":200,"message":"signup success","responseModel":""},{"code":201,"message":"signup failed","responseModel":""}]}]},{"path":"/signout","description":"","operations":[{"httpMethod":"GET","nickname":"signout","type":"","summary":"user signout","responseMessages":[{"code":300,"message":"signout success","responseModel":""},{"code":301,"message":"signout failed","responseModel":""}]}]},{"path":"/:username","description":"","operations":[{"httpMethod":"GET","nickname":"getUser","type":"","summary":"user getUser","parameters":[{"paramType":"query","name":"username","description":"\"The username for getUser\"","dataType":"string","type":"","format":"","allowMultiple":false,"required":true,"minimum":0,"maximum":0}],"responseMessages":[{"code":400,"message":"getUser success, show username","responseModel":""},{"code":401,"message":"getUser failed","responseModel":""}]}]},{"path":"/:username/contacts","description":"","operations":[{"httpMethod":"GET","nickname":"getContact","type":"","summary":"user getContact","parameters":[{"paramType":"query","name":"username","description":"\"The username for getContact\"","dataType":"string","type":"","format":"","allowMultiple":false,"required":true,"minimum":0,"maximum":0}],"responseMessages":[{"code":500,"message":"getContact success, show contacts","responseModel":""},{"code":501,"message":"getContact failed","responseModel":""}]}]},{"path":"/:username/contacts/:contact_username","description":"","operations":[{"httpMethod":"GET","nickname":"addContact","type":"","summary":"user addContact","parameters":[{"paramType":"query","name":"username","description":"\"The username for me\"","dataType":"string","type":"","format":"","allowMultiple":false,"required":true,"minimum":0,"maximum":0},{"paramType":"query","name":"contact","description":"\"The username for contact\"","dataType":"string","type":"","format":"","allowMultiple":false,"required":true,"minimum":0,"maximum":0}],"responseMessages":[{"code":600,"message":"addContact success , show id","responseModel":""},{"code":601,"message":"addContact failed","responseModel":""}]}]},{"path":"/:username/contacts/:contact_username","description":"","operations":[{"httpMethod":"DELETE","nickname":"delContact","type":"","summary":"user delContact","parameters":[{"paramType":"query","name":"username","description":"\"The username for me\"","dataType":"string","type":"","format":"","allowMultiple":false,"required":true,"minimum":0,"maximum":0},{"paramType":"query","name":"contact","description":"\"The username for contact\"","dataType":"string","type":"","format":"","allowMultiple":false,"required":true,"minimum":0,"maximum":0}],"responseMessages":[{"code":700,"message":"delContact success","responseModel":""},{"code":701,"message":"delContact failed","responseModel":""}]}]},{"path":"/:username/contacts/:contact_username/chats","description":"","operations":[{"httpMethod":"GET","nickname":"getChat","type":"","summary":"user getChat","parameters":[{"paramType":"query","name":"username","description":"\"The username for me\"","dataType":"string","type":"","format":"","allowMultiple":false,"required":true,"minimum":0,"maximum":0},{"paramType":"query","name":"contact","description":"\"The username for contact\"","dataType":"string","type":"","format":"","allowMultiple":false,"required":true,"minimum":0,"maximum":0}],"responseMessages":[{"code":800,"message":"getChat success, show chats","responseModel":""}]}]},{"path":"/:username/contacts/:contact_username/chats/:id","description":"","operations":[{"httpMethod":"DELETE","nickname":"delChat","type":"","summary":"user delContact","parameters":[{"paramType":"query","name":"id","description":"\"The id for chat\"","dataType":"string","type":"","format":"","allowMultiple":false,"required":true,"minimum":0,"maximum":0}],"responseMessages":[{"code":900,"message":"delChat success","responseModel":""},{"code":901,"message":"delChat failed","responseModel":""}]}]},{"path":"/:username/contacts/:contact_username/chats/:id","description":"","operations":[{"httpMethod":"GET","nickname":"updateChat","type":"","summary":"user updateChat","parameters":[{"paramType":"query","name":"id","description":"\"The id for chat\"","dataType":"string","type":"","format":"","allowMultiple":false,"required":true,"minimum":0,"maximum":0}],"responseMessages":[{"code":1000,"message":"updateChat success","responseModel":""},{"code":1001,"message":"updateChat failed","responseModel":""}]}]}]}}`
    BasePath string= "/v1"
)

var rootapi swagger.ResourceListing
var apilist map[string]*swagger.APIDeclaration

func init() {
	if beego.BConfig.WebConfig.EnableDocs {
		err := json.Unmarshal([]byte(Rootinfo), &rootapi)
		if err != nil {
			beego.Error(err)
		}
		err = json.Unmarshal([]byte(Subapi), &apilist)
		if err != nil {
			beego.Error(err)
		}
		beego.GlobalDocAPI["Root"] = rootapi
		for k, v := range apilist {
			for i, a := range v.APIs {
				a.Path = urlReplace(k + a.Path)
				v.APIs[i] = a
			}
			v.BasePath = BasePath
			beego.GlobalDocAPI[strings.Trim(k, "/")] = v
		}
	}
}


func urlReplace(src string) string {
	pt := strings.Split(src, "/")
	for i, p := range pt {
		if len(p) > 0 {
			if p[0] == ':' {
				pt[i] = "{" + p[1:] + "}"
			} else if p[0] == '?' && p[1] == ':' {
				pt[i] = "{" + p[2:] + "}"
			}
		}
	}
	return strings.Join(pt, "/")
}
