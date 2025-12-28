package generator

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"go-scaffold/internal/parser"
)

// GenerateRequests génère les fichiers de validation des requêtes
func (g *Generator) GenerateRequests() error {
	modelName := g.Schema.Model
	filename := filepath.Join("app", "requests", toSnakeCase(modelName)+"_request.go")

	// Créer le répertoire s'il n'existe pas
	if err := os.MkdirAll(filepath.Dir(filename), 0755); err != nil {
		return err
	}

	content := g.generateRequestsContent()
	return os.WriteFile(filename, []byte(content), 0644)
}

func (g *Generator) generateRequestsContent() string {
	var sb strings.Builder
	modelName := g.Schema.Model

	sb.WriteString("package requests\n\n")
	sb.WriteString("import (\n")
	
	needsTime := false
	for _, col := range g.Schema.Columns {
		if strings.Contains(col.GetGoType(), "time.Time") {
			needsTime = true
			break
		}
	}
	
	if needsTime {
		sb.WriteString("\t\"time\"\n\n")
	}
	
	sb.WriteString("\t\"app/models\"\n")
	sb.WriteString(")\n\n")

	// Create Request
	sb.WriteString(fmt.Sprintf("// Create%sRequest représente les données pour créer un %s\n", 
		modelName, toCamelCase(modelName)))
	sb.WriteString(fmt.Sprintf("type Create%sRequest struct {\n", modelName))

	for _, col := range g.Schema.Columns {
		// Exclure les colonnes auto-générées
		if col.Name == "id" || col.Name == "created_at" || col.Name == "updated_at" {
			continue
		}

		fieldName := toPascalCase(col.Name)
		goType := col.GetGoType()
		jsonTag := col.GetJSONTag()

		// Construire les tags de validation
		validationTags := g.buildValidationTags(col, true)

		sb.WriteString(fmt.Sprintf("\t%s %s `%s %s`\n", 
			fieldName, 
			goType, 
			jsonTag,
			validationTags,
		))
	}

	sb.WriteString("}\n\n")

	// Méthode ToModel pour Create
	sb.WriteString(fmt.Sprintf("// ToModel convertit la requête en model\n"))
	sb.WriteString(fmt.Sprintf("func (r *Create%sRequest) ToModel() models.%s {\n", modelName, modelName))
	sb.WriteString(fmt.Sprintf("\treturn models.%s{\n", modelName))

	for _, col := range g.Schema.Columns {
		if col.Name == "id" || col.Name == "created_at" || col.Name == "updated_at" {
			continue
		}
		fieldName := toPascalCase(col.Name)
		sb.WriteString(fmt.Sprintf("\t\t%s: r.%s,\n", fieldName, fieldName))
	}

	sb.WriteString("\t}\n")
	sb.WriteString("}\n\n")

	// Update Request
	sb.WriteString(fmt.Sprintf("// Update%sRequest représente les données pour mettre à jour un %s\n", 
		modelName, toCamelCase(modelName)))
	sb.WriteString(fmt.Sprintf("type Update%sRequest struct {\n", modelName))

	for _, col := range g.Schema.Columns {
		// Exclure les colonnes qui ne peuvent pas être mises à jour
		if col.Name == "id" || col.Name == "created_at" || col.Name == "updated_at" {
			continue
		}

		fieldName := toPascalCase(col.Name)
		goType := col.GetGoType()
		
		// Pour l'update, rendre les champs optionnels
		if !strings.HasPrefix(goType, "*") && goType != "interface{}" {
			goType = "*" + goType
		}
		
		jsonTag := "json:\"" + col.Name + ",omitempty\""

		// Construire les tags de validation (plus souples pour update)
		validationTags := g.buildValidationTags(col, false)

		sb.WriteString(fmt.Sprintf("\t%s %s `%s %s`\n", 
			fieldName, 
			goType, 
			jsonTag,
			validationTags,
		))
	}

	sb.WriteString("}\n\n")

	// Méthode UpdateModel pour Update
	sb.WriteString(fmt.Sprintf("// UpdateModel met à jour le model avec les données de la requête\n"))
	sb.WriteString(fmt.Sprintf("func (r *Update%sRequest) UpdateModel(m *models.%s) {\n", modelName, modelName))

	for _, col := range g.Schema.Columns {
		if col.Name == "id" || col.Name == "created_at" || col.Name == "updated_at" {
			continue
		}
		fieldName := toPascalCase(col.Name)
		sb.WriteString(fmt.Sprintf("\tif r.%s != nil {\n", fieldName))
		sb.WriteString(fmt.Sprintf("\t\tm.%s = *r.%s\n", fieldName, fieldName))
		sb.WriteString("\t}\n")
	}

	sb.WriteString("}\n")

	return sb.String()
}

func (g *Generator) buildValidationTags(col parser.Column, isCreate bool) string {
	var tags []string

	// Chercher les validations personnalisées dans le schéma
	for _, val := range g.Schema.Validations {
		if val.Field == col.Name {
			for rule, value := range val.Rules {
				switch rule {
				case "required":
					if isCreate && value == true {
						tags = append(tags, "required")
					}
				case "min":
					tags = append(tags, fmt.Sprintf("min=%v", value))
				case "max":
					tags = append(tags, fmt.Sprintf("max=%v", value))
				case "email":
					if value == true {
						tags = append(tags, "email")
					}
				case "url":
					if value == true {
						tags = append(tags, "url")
					}
				case "in":
					if arr, ok := value.([]interface{}); ok {
						var values []string
						for _, v := range arr {
							values = append(values, fmt.Sprintf("%v", v))
						}
						tags = append(tags, fmt.Sprintf("oneof=%s", strings.Join(values, " ")))
					}
				case "regex":
					tags = append(tags, fmt.Sprintf("regex=%v", value))
				}
			}
		}
	}

	// Ajouter des validations par défaut basées sur le type
	if len(tags) == 0 {
		if !col.Nullable && isCreate && col.Name != "id" {
			tags = append(tags, "required")
		}
		if col.Size > 0 && col.Type == "string" {
			tags = append(tags, fmt.Sprintf("max=%d", col.Size))
		}
		if col.Type == "email" {
			tags = append(tags, "email")
		}
	}

	if len(tags) == 0 {
		return ""
	}

	return fmt.Sprintf("validate:\"%s\"", strings.Join(tags, ","))
}
