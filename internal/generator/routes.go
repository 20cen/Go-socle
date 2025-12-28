package generator

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

// GenerateRoutes génère ou met à jour le fichier de routes
func (g *Generator) GenerateRoutes() error {
	modelName := g.Schema.Model
	routeFilename := filepath.Join("routes", toSnakeCase(modelName)+"_routes.go")

	// Créer le répertoire s'il n'existe pas
	if err := os.MkdirAll(filepath.Dir(routeFilename), 0755); err != nil {
		return err
	}

	// Générer le fichier de routes spécifique
	content := g.generateRouteContent()
	if err := os.WriteFile(routeFilename, []byte(content), 0644); err != nil {
		return err
	}

	// Mettre à jour le fichier routes.go principal
	return g.updateMainRoutesFile()
}

func (g *Generator) generateRouteContent() string {
	var sb strings.Builder
	modelName := g.Schema.Model
	varName := toCamelCase(modelName)
	resourceName := toSnakeCase(modelName) + "s"

	sb.WriteString("package routes\n\n")
	sb.WriteString("import (\n")
	sb.WriteString("\t\"app/controllers\"\n\n")
	sb.WriteString("\t\"github.com/gin-gonic/gin\"\n")
	sb.WriteString(")\n\n")

	sb.WriteString(fmt.Sprintf("// Register%sRoutes enregistre les routes pour %s\n", 
		modelName, modelName))
	sb.WriteString(fmt.Sprintf("func Register%sRoutes(router *gin.RouterGroup) {\n", modelName))
	sb.WriteString(fmt.Sprintf("\tctrl := controllers.New%sController()\n\n", modelName))
	
	sb.WriteString(fmt.Sprintf("\t// Routes RESTful pour %s\n", varName))
	sb.WriteString(fmt.Sprintf("\t%sGroup := router.Group(\"/%s\")\n", varName, resourceName))
	sb.WriteString("\t{\n")
	sb.WriteString(fmt.Sprintf("\t\t%sGroup.GET(\"\", ctrl.Index)        // GET /%s\n", 
		varName, resourceName))
	sb.WriteString(fmt.Sprintf("\t\t%sGroup.POST(\"\", ctrl.Store)       // POST /%s\n", 
		varName, resourceName))
	sb.WriteString(fmt.Sprintf("\t\t%sGroup.GET(\"/:id\", ctrl.Show)     // GET /%s/:id\n", 
		varName, resourceName))
	sb.WriteString(fmt.Sprintf("\t\t%sGroup.PUT(\"/:id\", ctrl.Update)   // PUT /%s/:id\n", 
		varName, resourceName))
	sb.WriteString(fmt.Sprintf("\t\t%sGroup.DELETE(\"/:id\", ctrl.Delete) // DELETE /%s/:id\n", 
		varName, resourceName))
	sb.WriteString("\t}\n")
	sb.WriteString("}\n")

	return sb.String()
}

func (g *Generator) updateMainRoutesFile() error {
	mainRoutesFile := filepath.Join("routes", "routes.go")
	modelName := g.Schema.Model

	// Lire le fichier existant
	content, err := os.ReadFile(mainRoutesFile)
	if err != nil {
		// Si le fichier n'existe pas, créer un nouveau
		return g.createMainRoutesFile()
	}

	contentStr := string(content)

	// Vérifier si la route est déjà enregistrée
	registerCall := fmt.Sprintf("Register%sRoutes(api)", modelName)
	if strings.Contains(contentStr, registerCall) {
		// La route est déjà enregistrée
		return nil
	}

	// Trouver où insérer le nouvel appel
	// Chercher la position avant la fermeture du bloc api
	insertPos := strings.LastIndex(contentStr, "}")
	if insertPos == -1 {
		return fmt.Errorf("impossible de trouver le point d'insertion dans routes.go")
	}

	// Trouver la dernière ligne avant la fermeture
	lastNewline := strings.LastIndex(contentStr[:insertPos], "\n")
	
	// Insérer le nouvel appel
	newCall := fmt.Sprintf("\n\t\tRegister%sRoutes(api)", modelName)
	newContent := contentStr[:lastNewline] + newCall + contentStr[lastNewline:]

	return os.WriteFile(mainRoutesFile, []byte(newContent), 0644)
}

func (g *Generator) createMainRoutesFile() error {
	mainRoutesFile := filepath.Join("routes", "routes.go")
	modelName := g.Schema.Model

	content := `package routes

import (
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.Engine) {
	// Routes de l'API
	api := router.Group("/api")
	{
		// Route de santé
		api.GET("/health", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"status": "ok",
				"message": "Service en cours d'exécution",
			})
		})

		// Routes générées
		Register` + modelName + `Routes(api)
	}
}
`

	return os.WriteFile(mainRoutesFile, []byte(content), 0644)
}
