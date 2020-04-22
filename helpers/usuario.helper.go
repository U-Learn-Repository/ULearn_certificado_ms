package helpers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/graphql-go/graphql"
)

type usuario struct {
	ID        int    `json:"id"`
	Name      string `json:"names"`
	Documento int    `json:"idDocumment"`
	Surname   string `json:"surnames"`
	Username  string `json:"username"`
}

var data map[string]usuario

var userType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "User",
		Fields: graphql.Fields{
			"id": &graphql.Field{
				Type: graphql.Int,
			},
			"names": &graphql.Field{
				Type: graphql.String,
			},
			"idDocumment": &graphql.Field{
				Type: graphql.Int,
			},
			"surnames": &graphql.Field{
				Type: graphql.String,
			},
			"username": &graphql.Field{
				Type: graphql.String,
			},
		},
	},
)

var queryType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Query",
		Fields: graphql.Fields{
			"usuario": &graphql.Field{
				Type: userType,
				Args: graphql.FieldConfigArgument{
					"id": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
				},
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					idQuery, isOK := p.Args["id"].(string)
					if isOK {
						return data[idQuery], nil
					}
					return nil, nil
				},
			},
		},
	})

var schema, _ = graphql.NewSchema(
	graphql.SchemaConfig{
		Query: queryType,
	},
)

func executeQuery(query string, schema graphql.Schema) *graphql.Result {
	result := graphql.Do(graphql.Params{
		Schema:        schema,
		RequestString: query,
	})
	if len(result.Errors) > 0 {
		fmt.Printf("wrong result, unexpected errors: %v", result.Errors)
	}
	return result
}

func importJSONDataFromFile(fileName string, result interface{}) (isOK bool) {
	isOK = true
	content, err := ioutil.ReadFile(fileName)
	if err != nil {
		fmt.Print("Error1:", err)
		isOK = false
	}
	err = json.Unmarshal(content, result)
	if err != nil {
		isOK = false
		fmt.Print("Error:", err)
	}
	return
}

func Principal() {
	_ = importJSONDataFromFile("data.json", &data)

	http.HandleFunc("/graphql", func(w http.ResponseWriter, r *http.Request) {
		result := executeQuery("http://localhost:5000/graphql?query={buscarUsuario(userId:\"1\"){names}}", schema)
		json.NewEncoder(w).Encode(result)
		fmt.Println("-------------------------------------------------------------------------------------------")
		fmt.Println(w)
		fmt.Println("-------------------------------------------------------------------------------------------")
	})
	fmt.Println("-------------------------------------------------------------------------------------------")
	fmt.Println("Now server is running on port 8080")
	fmt.Println("Test with Get      : curl -g 'http://172.17.0.1:6665/graphql?query={user(id:\"1\"){name}}'")
	http.ListenAndServe(":5000", nil)

	return
}

func ObtenerUsuario(usu int) map[string]interface{} {
	/*	var (
		url_crud  string
		respuesta map[string]interface{}
	)*/

	/*	if _, err := request.getJson(URLCrud,respuesta); error = nil{
			return respuesta,  nil
		}else {
			nil, err
		}	*/
	respuestaTemp := map[string]interface{}{
		"user_id":  usu,
		"username": "diagutierrezro",
		"password": "12345",
		"names":    "Diego Alejandro",
		"surnames": "Gutierrez Rojas",
	}
	return respuestaTemp
}
