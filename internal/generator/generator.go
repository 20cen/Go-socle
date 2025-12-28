package generator

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"go-scaffold/internal/parser"
)

// Generator gère la génération de code
type Generator struct {
	Schema *parser.Schema
}

// NewGenerator crée une nouvelle instance de Generator
func NewGenerator(schema *parser.Schema) *Generator {
	return &Generator{
		Schema: schema,
	}
}

// GenerateModel génère le fichier model
func (g *Generator) GenerateModel() error {
	modelName := g.Schema.Model
	filename := filepath.Join("app", "models", toSnakeCase(modelName)+".go")

	// Créer le répertoire s'il n'existe pas
	if err := os.MkdirAll(filepath.Dir(filename), 0755); err != nil {
		return err
	}

	content := g.generateModelContent()
	return os.WriteFile(filename, []byte(content), 0644)
}

func (g *Generator) generateModelContent() string {
	var sb strings.Builder

	sb.WriteString("package models\n\n")
	sb.WriteString("import (\n")
	
	needsTime := false
	for _, col := range g.Schema.Columns {
		if strings.Contains(col.GetGoType(), "time.Time") {
			needsTime = true
			break
		}
	}
	
	if needsTime {
		sb.WriteString("\t\"time\"\n")
	}
	
	sb.WriteString("\t\"gorm.io/gorm\"\n")
	sb.WriteString(")\n\n")

	// Structure principale
	sb.WriteString(fmt.Sprintf("// %s représente la table %s\n", g.Schema.Model, g.Schema.Table))
	sb.WriteString(fmt.Sprintf("type %s struct {\n", g.Schema.Model))

	// Générer les champs
	for _, col := range g.Schema.Columns {
		fieldName := toPascalCase(col.Name)
		goType := col.GetGoType()
		jsonTag := col.GetJSONTag()
		
		// Construire les tags GORM
		gormTags := []string{}
		if col.Primary {
			gormTags = append(gormTags, "primaryKey")
		}
		if col.AutoIncrement {
			gormTags = append(gormTags, "autoIncrement")
		}
		if col.Unique {
			gormTags = append(gormTags, "unique")
		}
		if !col.Nullable {
			gormTags = append(gormTags, "not null")
		}
		if col.Size > 0 && col.Type == "string" {
			gormTags = append(gormTags, fmt.Sprintf("size:%d", col.Size))
		}
		if col.Default != nil {
			gormTags = append(gormTags, fmt.Sprintf("default:%v", col.Default))
		}
		if col.Name != "" {
			gormTags = append(gormTags, fmt.Sprintf("column:%s", col.Name))
		}
		
		gormTag := ""
		if len(gormTags) > 0 {
			gormTag = fmt.Sprintf(" gorm:\"%s\"", strings.Join(gormTags, ";"))
		}

		validTag := col.GetValidationTag()
		if validTag != "" {
			validTag = " " + validTag
		}

		sb.WriteString(fmt.Sprintf("\t%s %s `%s%s%s`\n", 
			fieldName, 
			goType, 
			jsonTag,
			gormTag,
			validTag,
		))
	}

	// Ajouter les relations
	for _, rel := range g.Schema.Relations {
		sb.WriteString(fmt.Sprintf("\t// Relation: %s\n", rel.Type))
		relModel := rel.Model
		
		switch rel.Type {
		case "belongs_to":
			sb.WriteString(fmt.Sprintf("\t%s *%s `json:\"%s,omitempty\" gorm:\"foreignKey:%s\"`\n",
				relModel,
				relModel,
				toSnakeCase(relModel),
				rel.ForeignKey,
			))
		case "has_many":
			sb.WriteString(fmt.Sprintf("\t%ss []%s `json:\"%ss,omitempty\" gorm:\"foreignKey:%s\"`\n",
				relModel,
				relModel,
				toSnakeCase(relModel),
				rel.ForeignKey,
			))
		case "has_one":
			sb.WriteString(fmt.Sprintf("\t%s *%s `json:\"%s,omitempty\" gorm:\"foreignKey:%s\"`\n",
				relModel,
				relModel,
				toSnakeCase(relModel),
				rel.ForeignKey,
			))
		case "many_to_many":
			sb.WriteString(fmt.Sprintf("\t%ss []%s `json:\"%ss,omitempty\" gorm:\"many2many:%s\"`\n",
				relModel,
				relModel,
				toSnakeCase(relModel),
				rel.PivotTable,
			))
		}
	}

	sb.WriteString("}\n\n")

	// TableName
	sb.WriteString(fmt.Sprintf("// TableName retourne le nom de la table\n"))
	sb.WriteString(fmt.Sprintf("func (%s) TableName() string {\n", g.Schema.Model))
	sb.WriteString(fmt.Sprintf("\treturn \"%s\"\n", g.Schema.Table))
	sb.WriteString("}\n\n")

	// Hooks GORM (optionnel)
	sb.WriteString(fmt.Sprintf("// BeforeCreate hook GORM\n"))
	sb.WriteString(fmt.Sprintf("func (m *%s) BeforeCreate(tx *gorm.DB) error {\n", g.Schema.Model))
	sb.WriteString("\t// Logique avant création\n")
	sb.WriteString("\treturn nil\n")
	sb.WriteString("}\n\n")

	sb.WriteString(fmt.Sprintf("// BeforeUpdate hook GORM\n"))
	sb.WriteString(fmt.Sprintf("func (m *%s) BeforeUpdate(tx *gorm.DB) error {\n", g.Schema.Model))
	sb.WriteString("\t// Logique avant mise à jour\n")
	sb.WriteString("\treturn nil\n")
	sb.WriteString("}\n")

	return sb.String()
}

// Fonctions utilitaires
func toSnakeCase(s string) string {
	var result strings.Builder
	for i, r := range s {
		if i > 0 && r >= 'A' && r <= 'Z' {
			result.WriteRune('_')
		}
		result.WriteRune(r)
	}
	return strings.ToLower(result.String())
}

func toPascalCase(s string) string {
	words := strings.FieldsFunc(s, func(r rune) bool {
		return r == '_' || r == '-' || r == ' '
	})
	for i, word := range words {
		if len(word) > 0 {
			words[i] = strings.ToUpper(word[:1]) + strings.ToLower(word[1:])
		}
	}
	return strings.Join(words, "")
}

func toCamelCase(s string) string {
	pascal := toPascalCase(s)
	if len(pascal) == 0 {
		return pascal
	}
	return strings.ToLower(pascal[:1]) + pascal[1:]
}
