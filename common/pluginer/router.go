package pluginer

import (
	"net/http"
	"path"
	"strings"

	"github.com/charmbracelet/log"
	"github.com/gin-gonic/gin"
)

type Router struct {
	RouterGroup
}

type Context = gin.Context

const (
	GET    = http.MethodGet
	POST   = http.MethodPost
	PUT    = http.MethodPut
	DELETE = http.MethodDelete
	PATCH  = http.MethodPatch
)

type RouterGroup struct {
	Handlers     HandlersChain
	BasePath     string
	RelativePath string
	SubGroups    map[string]*RouterGroup
}

type HandlerFunction func(*Context)
type HandlersChain map[string][]HandlerFunction

func getElementsFromPath(path string) []string {
	components := strings.Split(path, "/")

	var cleaned []string
	for _, comp := range components {
		if comp != "" {
			cleaned = append(cleaned, comp)
		}
	}

	return cleaned
}

func (group *RouterGroup) combineHandlers(handlers HandlersChain) HandlersChain {
	finalSize := len(group.Handlers) + len(handlers)
	mergedHandlers := make(HandlersChain, finalSize)

	i := 0
	for key, groupFuncs := range group.Handlers {
		if funcs, ok := handlers[key]; ok {
			mergedHandlers[key] = append(groupFuncs, funcs...)
		} else {
			mergedHandlers[key] = groupFuncs
		}
		i++
	}

	for key, funcs := range handlers {
		if _, ok := mergedHandlers[key]; !ok {
			mergedHandlers[key] = funcs
		}
	}

	return mergedHandlers
}

func (group *RouterGroup) calculateAbsolutePath(relativePath string) string {
	return path.Join(group.BasePath, relativePath)
}

func NewRouter() *Router {
	return &Router{}
}

func (group *RouterGroup) GetGroup(relativePath string) *RouterGroup {
	elements := getElementsFromPath(relativePath)
	currentGroup := group

	for _, element := range elements {
		if subGroup, ok := currentGroup.SubGroups[element]; ok {
			currentGroup = subGroup
			continue
		}

		for _, subGroup := range currentGroup.SubGroups {
			if strings.HasPrefix(subGroup.RelativePath, ":") {
				currentGroup = subGroup
				break
			}
		}
	}

	if len(getElementsFromPath(currentGroup.BasePath)) != len(elements) {
		return nil
	}

	return currentGroup
}

func (group *RouterGroup) Group(relativePath string, handlers HandlersChain) *RouterGroup {
	// Lazy load sub groups
	if group.SubGroups == nil {
		group.SubGroups = make(map[string]*RouterGroup)
	}

	newHandlers := handlers
	if handlers != nil {
		newHandlers = group.combineHandlers(handlers)
	}

	newGroup := &RouterGroup{
		Handlers:     newHandlers,
		BasePath:     group.calculateAbsolutePath(relativePath),
		RelativePath: relativePath,
		SubGroups:    make(map[string]*RouterGroup),
	}

	group.SubGroups[relativePath] = newGroup

	log.Debug("created new group", "relative path", relativePath, "base path", newGroup.BasePath, "handlers", newGroup.Handlers, "own handlers", handlers)

	return newGroup
}

func (group *RouterGroup) Execute(httpMethod string, context *Context) {
	if handlers, ok := group.Handlers[httpMethod]; ok {
		for _, handler := range handlers {
			handler(context)
		}
	} else {
		log.Warn("no handlers found for the specified HTTP method")
		context.AbortWithStatus(http.StatusNotFound)
	}
}

func (group *RouterGroup) Handle(httpMethod, relativePath string, handlers ...HandlerFunction) *RouterGroup {
	if group.Handlers == nil {
		group.Handlers = make(HandlersChain)
	}

	newHandlers := HandlersChain{}
	newHandlers[httpMethod] = append(newHandlers[httpMethod], handlers...)

	return group.Group(relativePath, newHandlers)
}

func (group *RouterGroup) POST(relativePath string, handlers ...HandlerFunction) *RouterGroup {
	return group.Handle(http.MethodPost, relativePath, handlers...)
}

func (group *RouterGroup) GET(relativePath string, handlers ...HandlerFunction) *RouterGroup {
	return group.Handle(http.MethodGet, relativePath, handlers...)
}
