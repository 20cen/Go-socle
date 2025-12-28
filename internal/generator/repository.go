package generator

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

// GenerateRepository génère le fichier repository
func (g *Generator) GenerateRepository() error {
	modelName := g.Schema.Model
	filename := filepath.Join("app", "repositories", toSnakeCase(modelName)+"_repository.go")

	// Créer le répertoire s'il n'existe pas
	if err := os.MkdirAll(filepath.Dir(filename), 0755); err != nil {
		return err
	}

	content := g.generateRepositoryContent()
	return os.WriteFile(filename, []byte(content), 0644)
}

func (g *Generator) generateRepositoryContent() string {
	var sb strings.Builder
	modelName := g.Schema.Model
	repoName := modelName + "Repository"
	varName := toCamelCase(modelName)

	sb.WriteString("package repositories\n\n")
	sb.WriteString("import (\n")
	sb.WriteString("\t\"errors\"\n")
	sb.WriteString("\t\"app/models\"\n")
	sb.WriteString("\t\"config\"\n\n")
	sb.WriteString("\t\"gorm.io/gorm\"\n")
	sb.WriteString(")\n\n")

	// Interface du repository
	sb.WriteString(fmt.Sprintf("// %sInterface définit les méthodes du repository\n", modelName))
	sb.WriteString(fmt.Sprintf("type %sInterface interface {\n", modelName))
	sb.WriteString(fmt.Sprintf("\tCreate(%s *models.%s) error\n", varName, modelName))
	sb.WriteString(fmt.Sprintf("\tFindByID(id string) (*models.%s, error)\n", modelName))
	sb.WriteString(fmt.Sprintf("\tFindAll(page, pageSize int) ([]models.%s, int64, error)\n", modelName))
	sb.WriteString(fmt.Sprintf("\tUpdate(%s *models.%s) error\n", varName, modelName))
	sb.WriteString(fmt.Sprintf("\tDelete(id string) error\n"))
	
	// Ajouter des méthodes de recherche personnalisées basées sur les colonnes
	for _, col := range g.Schema.Columns {
		if col.Unique && col.Name != "id" {
			fieldName := toPascalCase(col.Name)
			sb.WriteString(fmt.Sprintf("\tFindBy%s(%s %s) (*models.%s, error)\n", 
				fieldName, 
				col.Name, 
				col.GetGoType(), 
				modelName))
		}
	}
	
	sb.WriteString("}\n\n")

	// Structure du repository
	sb.WriteString(fmt.Sprintf("// %s implémente %sInterface\n", repoName, modelName))
	sb.WriteString(fmt.Sprintf("type %s struct {\n", repoName))
	sb.WriteString("\tdb *gorm.DB\n")
	sb.WriteString("}\n\n")

	// Constructeur
	sb.WriteString(fmt.Sprintf("// New%s crée une nouvelle instance du repository\n", repoName))
	sb.WriteString(fmt.Sprintf("func New%s() %sInterface {\n", repoName, modelName))
	sb.WriteString(fmt.Sprintf("\treturn &%s{\n", repoName))
	sb.WriteString("\t\tdb: config.GetDB(),\n")
	sb.WriteString("\t}\n")
	sb.WriteString("}\n\n")

	// Méthode Create
	sb.WriteString(fmt.Sprintf("// Create crée un nouveau %s\n", varName))
	sb.WriteString(fmt.Sprintf("func (r *%s) Create(%s *models.%s) error {\n", repoName, varName, modelName))
	sb.WriteString(fmt.Sprintf("\treturn r.db.Create(%s).Error\n", varName))
	sb.WriteString("}\n\n")

	// Méthode FindByID
	sb.WriteString(fmt.Sprintf("// FindByID trouve un %s par son ID\n", varName))
	sb.WriteString(fmt.Sprintf("func (r *%s) FindByID(id string) (*models.%s, error) {\n", repoName, modelName))
	sb.WriteString(fmt.Sprintf("\tvar %s models.%s\n", varName, modelName))
	
	// Ajouter les préchargements des relations
	preloads := []string{}
	for _, rel := range g.Schema.Relations {
		preloads = append(preloads, rel.Model)
	}
	
	query := "r.db"
	for _, preload := range preloads {
		query += fmt.Sprintf(".Preload(\"%s\")", preload)
	}
	
	sb.WriteString(fmt.Sprintf("\terr := %s.First(&%s, \"id = ?\", id).Error\n", query, varName))
	sb.WriteString("\tif err != nil {\n")
	sb.WriteString("\t\tif errors.Is(err, gorm.ErrRecordNotFound) {\n")
	sb.WriteString("\t\t\treturn nil, errors.New(\"enregistrement non trouvé\")\n")
	sb.WriteString("\t\t}\n")
	sb.WriteString("\t\treturn nil, err\n")
	sb.WriteString("\t}\n")
	sb.WriteString(fmt.Sprintf("\treturn &%s, nil\n", varName))
	sb.WriteString("}\n\n")

	// Méthode FindAll
	sb.WriteString(fmt.Sprintf("// FindAll récupère tous les %ss avec pagination\n", varName))
	sb.WriteString(fmt.Sprintf("func (r *%s) FindAll(page, pageSize int) ([]models.%s, int64, error) {\n", repoName, modelName))
	sb.WriteString(fmt.Sprintf("\tvar %ss []models.%s\n", varName, modelName))
	sb.WriteString("\tvar total int64\n\n")
	sb.WriteString(fmt.Sprintf("\t// Compter le total\n"))
	sb.WriteString(fmt.Sprintf("\tif err := r.db.Model(&models.%s{}).Count(&total).Error; err != nil {\n", modelName))
	sb.WriteString("\t\treturn nil, 0, err\n")
	sb.WriteString("\t}\n\n")
	sb.WriteString("\t// Calculer l'offset\n")
	sb.WriteString("\toffset := (page - 1) * pageSize\n\n")
	sb.WriteString("\t// Récupérer les données avec pagination\n")
	
	query = "r.db"
	for _, preload := range preloads {
		query += fmt.Sprintf(".Preload(\"%s\")", preload)
	}
	
	sb.WriteString(fmt.Sprintf("\terr := %s.\n", query))
	sb.WriteString("\t\tOffset(offset).\n")
	sb.WriteString("\t\tLimit(pageSize).\n")
	sb.WriteString(fmt.Sprintf("\t\tFind(&%ss).Error\n\n", varName))
	sb.WriteString("\tif err != nil {\n")
	sb.WriteString("\t\treturn nil, 0, err\n")
	sb.WriteString("\t}\n\n")
	sb.WriteString(fmt.Sprintf("\treturn %ss, total, nil\n", varName))
	sb.WriteString("}\n\n")

	// Méthode Update
	sb.WriteString(fmt.Sprintf("// Update met à jour un %s\n", varName))
	sb.WriteString(fmt.Sprintf("func (r *%s) Update(%s *models.%s) error {\n", repoName, varName, modelName))
	sb.WriteString(fmt.Sprintf("\treturn r.db.Save(%s).Error\n", varName))
	sb.WriteString("}\n\n")

	// Méthode Delete
	sb.WriteString(fmt.Sprintf("// Delete supprime un %s\n", varName))
	sb.WriteString(fmt.Sprintf("func (r *%s) Delete(id string) error {\n", repoName))
	sb.WriteString(fmt.Sprintf("\treturn r.db.Delete(&models.%s{}, \"id = ?\", id).Error\n", modelName))
	sb.WriteString("}\n\n")

	// Méthodes de recherche personnalisées
	for _, col := range g.Schema.Columns {
		if col.Unique && col.Name != "id" {
			fieldName := toPascalCase(col.Name)
			sb.WriteString(fmt.Sprintf("// FindBy%s trouve un %s par son %s\n", fieldName, varName, col.Name))
			sb.WriteString(fmt.Sprintf("func (r *%s) FindBy%s(%s %s) (*models.%s, error) {\n", 
				repoName, 
				fieldName, 
				col.Name, 
				col.GetGoType(), 
				modelName))
			sb.WriteString(fmt.Sprintf("\tvar %s models.%s\n", varName, modelName))
			
			query = "r.db"
			for _, preload := range preloads {
				query += fmt.Sprintf(".Preload(\"%s\")", preload)
			}
			
			sb.WriteString(fmt.Sprintf("\terr := %s.Where(\"%s = ?\", %s).First(&%s).Error\n", 
				query, 
				col.Name, 
				col.Name, 
				varName))
			sb.WriteString("\tif err != nil {\n")
			sb.WriteString("\t\tif errors.Is(err, gorm.ErrRecordNotFound) {\n")
			sb.WriteString("\t\t\treturn nil, errors.New(\"enregistrement non trouvé\")\n")
			sb.WriteString("\t\t}\n")
			sb.WriteString("\t\treturn nil, err\n")
			sb.WriteString("\t}\n")
			sb.WriteString(fmt.Sprintf("\treturn &%s, nil\n", varName))
			sb.WriteString("}\n\n")
		}
	}

	return sb.String()
}
