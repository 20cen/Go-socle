package parser

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

// Schema représente la structure complète d'un schéma de table
type Schema struct {
	Table       string       `yaml:"table"`
	Model       string       `yaml:"model"`
	Columns     []Column     `yaml:"columns"`
	Relations   []Relation   `yaml:"relations"`
	Indexes     []Index      `yaml:"indexes"`
	Validations []Validation `yaml:"validations"`
}

// Column représente une colonne de table
type Column struct {
	Name          string      `yaml:"name"`
	Type          string      `yaml:"type"`
	Size          int         `yaml:"size"`
	Primary       bool        `yaml:"primary"`
	AutoIncrement bool        `yaml:"auto_increment"`
	Nullable      bool        `yaml:"nullable"`
	Unique        bool        `yaml:"unique"`
	Default       interface{} `yaml:"default"`
	Comment       string      `yaml:"comment"`
}

// Relation représente une relation entre tables
type Relation struct {
	Type        string `yaml:"type"` // belongs_to, has_many, has_one, many_to_many
	Model       string `yaml:"model"`
	ForeignKey  string `yaml:"foreign_key"`
	References  string `yaml:"references"`
	PivotTable  string `yaml:"pivot_table"`
	RelatedKey  string `yaml:"related_key"`
}

// Index représente un index de base de données
type Index struct {
	Name    string   `yaml:"name"`
	Columns []string `yaml:"columns"`
	Unique  bool     `yaml:"unique"`
}

// Validation représente les règles de validation
type Validation struct {
	Field string                 `yaml:"field"`
	Rules map[string]interface{} `yaml:"rules"`
}

// ParseSchema lit et parse un fichier de schéma YAML
func ParseSchema(filename string) (*Schema, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		return nil, fmt.Errorf("impossible de lire le fichier: %w", err)
	}

	var schema Schema
	if err := yaml.Unmarshal(data, &schema); err != nil {
		return nil, fmt.Errorf("erreur de parsing YAML: %w", err)
	}

	// Validation du schéma
	if err := validateSchema(&schema); err != nil {
		return nil, err
	}

	return &schema, nil
}

// validateSchema valide la structure du schéma
func validateSchema(schema *Schema) error {
	if schema.Table == "" {
		return fmt.Errorf("le nom de la table est requis")
	}
	if schema.Model == "" {
		return fmt.Errorf("le nom du model est requis")
	}
	if len(schema.Columns) == 0 {
		return fmt.Errorf("au moins une colonne est requise")
	}
	return nil
}

// GetGoType convertit un type de base de données en type Go
func (c *Column) GetGoType() string {
	typeMap := map[string]string{
		"string":    "string",
		"text":      "string",
		"int":       "int",
		"integer":   "int",
		"bigint":    "int64",
		"smallint":  "int16",
		"float":     "float64",
		"double":    "float64",
		"decimal":   "float64",
		"boolean":   "bool",
		"bool":      "bool",
		"date":      "time.Time",
		"datetime":  "time.Time",
		"timestamp": "time.Time",
		"time":      "time.Time",
		"uuid":      "string",
		"json":      "string",
		"jsonb":     "string",
	}

	goType, exists := typeMap[c.Type]
	if !exists {
		goType = "interface{}"
	}

	if c.Nullable && goType != "interface{}" {
		return "*" + goType
	}

	return goType
}

// GetDBType convertit le type en type de base de données
func (c *Column) GetDBType() string {
	if c.Type == "uuid" {
		return "uuid"
	}
	if c.Type == "string" && c.Size > 0 {
		return fmt.Sprintf("varchar(%d)", c.Size)
	}
	return c.Type
}

// GetValidationTag retourne les tags de validation pour un champ
func (c *Column) GetValidationTag() string {
	tags := []string{}

	if !c.Nullable && c.Name != "id" && c.Name != "created_at" && c.Name != "updated_at" {
		tags = append(tags, "required")
	}

	if c.Size > 0 {
		tags = append(tags, fmt.Sprintf("max=%d", c.Size))
	}

	if c.Unique {
		tags = append(tags, "unique")
	}

	if len(tags) == 0 {
		return ""
	}

	result := "validate:\""
	for i, tag := range tags {
		if i > 0 {
			result += ","
		}
		result += tag
	}
	result += "\""

	return result
}

// GetJSONTag retourne le tag JSON pour un champ
func (c *Column) GetJSONTag() string {
	tag := "json:\"" + c.Name
	if c.Nullable {
		tag += ",omitempty"
	}
	tag += "\""
	return tag
}
