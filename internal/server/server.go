package server

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/wirequery/semver-server/pkg/store"
	"github.com/wirequery/semver-server/pkg/versioning"
	"io"
)

type Server struct {
	Port         int
	VersionStore store.Store
}

type newProjectRequestBody struct {
	Group string
	Name  string
}

func (server *Server) Serve() {
	r := gin.Default()
	err := r.SetTrustedProxies(nil)
	if err != nil {
		return
	}

	v1 := r.Group("/v1")
	{
		v1.POST("/projects", func(c *gin.Context) {
			bodyAsByteArray, _ := io.ReadAll(c.Request.Body)
			jsonMap := newProjectRequestBody{}
			err := json.Unmarshal(bodyAsByteArray, &jsonMap)
			if err != nil {
				panic(err)
			}
			key := store.ProjectKey{
				Group: jsonMap.Group,
				Name:  jsonMap.Name,
			}
			value := store.VersionValue{Version: versioning.New()}
			server.VersionStore.AddVersion(key, value)
			c.IndentedJSON(200, value)
		})

		v1.GET("/projects/:group/:name", func(c *gin.Context) {
			key := store.ProjectKey{
				Group: c.Param("group"),
				Name:  c.Param("name"),
			}
			value := server.VersionStore.GetLatestVersion(key)
			c.IndentedJSON(200, value)
		})

		v1.POST("/projects/:group/:name/major", func(c *gin.Context) {
			server.handleIncrement(c, versioning.IncrementMajor)
		})

		v1.POST("/projects/:group/:name/minor", func(c *gin.Context) {
			server.handleIncrement(c, versioning.IncrementMinor)
		})

		v1.POST("/projects/:group/:name/patch", func(c *gin.Context) {
			server.handleIncrement(c, versioning.IncrementPatch)
		})
	}

	err = r.Run(fmt.Sprintf("%s%d", ":", server.Port))
	if err != nil {
		panic(err)
	}
}

func (server *Server) handleIncrement(c *gin.Context, incrementFunc func(s versioning.SemanticVersion) versioning.SemanticVersion) {
	key := store.ProjectKey{
		Group: c.Param("group"),
		Name:  c.Param("name"),
	}
	value := server.VersionStore.GetLatestVersion(key)
	updatedValue := incrementFunc(value.Version)
	server.VersionStore.AddVersion(key, store.VersionValue{Version: updatedValue})
	c.IndentedJSON(200, &updatedValue)
}
