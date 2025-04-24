package document_processor

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

// --- Product Interface ---
type Document interface {
	// Save writes the document content to a file.
	// It constructs the filename from outputDir and filenameBase.
	Save(outputDir string, filenameBase string) error
	// GetTitle retrieves the document title.
	GetTitle() string
	// GetContent retrieves the document content lines.
	GetContent() []string
}

// --- Concrete Products ---

// TextDocument implements Document
type TextDocument struct {
	Title   string
	Content []string
}

func NewTextDocument(title string, content []string) *TextDocument {
	return &TextDocument{Title: title, Content: content}
}

func (t *TextDocument) GetTitle() string {
	return t.Title
}

func (t *TextDocument) GetContent() []string {
	return t.Content
}

func (t *TextDocument) Save(outputDir string, filenameBase string) error {
	filePath := filepath.Join(outputDir, filenameBase+".txt")
	file, err := os.Create(filePath)
	if err != nil {
		return fmt.Errorf("failed to create text file %s: %w", filePath, err)
	}
	defer file.Close()

	_, err = fmt.Fprintf(file, "Title: %s\n\n", t.Title)
	if err != nil {
		return fmt.Errorf("failed to write title to text file %s: %w", filePath, err)
	}

	for _, line := range t.Content {
		_, err = fmt.Fprintln(file, line)
		if err != nil {
			return fmt.Errorf("failed to write content line to text file %s: %w", filePath, err)
		}
	}
	fmt.Printf("Text document saved to: %s\n", filePath)
	return nil
}

// JSONDocument implements Document
type JSONDocument struct {
	Title   string   `json:"title"`
	Content []string `json:"content"`
}

func NewJSONDocument(title string, content []string) *JSONDocument {
	return &JSONDocument{Title: title, Content: content}
}

func (j *JSONDocument) GetTitle() string {
	return j.Title
}

func (j *JSONDocument) GetContent() []string {
	return j.Content
}

func (j *JSONDocument) Save(outputDir string, filenameBase string) error {
	filePath := filepath.Join(outputDir, filenameBase+".json")
	data, err := json.MarshalIndent(j, "", "  ") // Use indentation for readability
	if err != nil {
		return fmt.Errorf("failed to marshal JSON data: %w", err)
	}

	err = os.WriteFile(filePath, data, 0644) // 0644 are standard file permissions
	if err != nil {
		return fmt.Errorf("failed to write JSON file %s: %w", filePath, err)
	}
	fmt.Printf("JSON document saved to: %s\n", filePath)
	return nil
}

// HTMLDocument implements Document
type HTMLDocument struct {
	Title   string
	Content []string
}

func NewHTMLDocument(title string, content []string) *HTMLDocument {
	return &HTMLDocument{Title: title, Content: content}
}

func (h *HTMLDocument) GetTitle() string {
	return h.Title
}

func (h *HTMLDocument) GetContent() []string {
	return h.Content
}

func (h *HTMLDocument) Save(outputDir string, filenameBase string) error {
	filePath := filepath.Join(outputDir, filenameBase+".html")
	file, err := os.Create(filePath)
	if err != nil {
		return fmt.Errorf("failed to create HTML file %s: %w", filePath, err)
	}
	defer file.Close()

	var htmlBuilder strings.Builder
	htmlBuilder.WriteString("<!DOCTYPE html>\n")
	htmlBuilder.WriteString("<html>\n")
	htmlBuilder.WriteString("<head>\n")
	htmlBuilder.WriteString(fmt.Sprintf("  <title>%s</title>\n", h.Title))
	htmlBuilder.WriteString("</head>\n")
	htmlBuilder.WriteString("<body>\n")
	htmlBuilder.WriteString(fmt.Sprintf("  <h1>%s</h1>\n", h.Title))
	for _, line := range h.Content {
		htmlBuilder.WriteString(fmt.Sprintf("  <p>%s</p>\n", line))
	}
	htmlBuilder.WriteString("</body>\n")
	htmlBuilder.WriteString("</html>\n")

	_, err = file.WriteString(htmlBuilder.String())
	if err != nil {
		return fmt.Errorf("failed to write HTML content to file %s: %w", filePath, err)
	}

	fmt.Printf("HTML document saved to: %s\n", filePath)
	return nil
}

// --- Creator Interface (Factory) ---
type DocumentProcessor interface {
	// CreateDocument is the Factory Method
	CreateDocument(title string, content []string) Document
	// ProcessDocument uses the factory method to create and save a document.
	ProcessDocument(title string, content []string, outputDir string, filenameBase string) (Document, error)
}

// --- Concrete Creators (Concrete Factories) ---

// TextProcessor implements DocumentProcessor
type TextProcessor struct{}

func (p *TextProcessor) CreateDocument(title string, content []string) Document {
	fmt.Println("TextProcessor: Creating a Text document.")
	return NewTextDocument(title, content)
}

func (p *TextProcessor) ProcessDocument(title string, content []string, outputDir string, filenameBase string) (Document, error) {
	fmt.Println("--- Processing with TextProcessor ---")
	// Ensure output directory exists
	err := os.MkdirAll(outputDir, 0755) // 0755 are standard directory permissions
	if err != nil {
		return nil, fmt.Errorf("failed to create output directory %s: %w", outputDir, err)
	}

	doc := p.CreateDocument(title, content)
	err = doc.Save(outputDir, filenameBase)
	if err != nil {
		return nil, err // Error already formatted by Save method
	}
	fmt.Println("--- Text processing complete ---")
	return doc, nil
}

// JSONProcessor implements DocumentProcessor
type JSONProcessor struct{}

func (p *JSONProcessor) CreateDocument(title string, content []string) Document {
	fmt.Println("JSONProcessor: Creating a JSON document.")
	return NewJSONDocument(title, content)
}

func (p *JSONProcessor) ProcessDocument(title string, content []string, outputDir string, filenameBase string) (Document, error) {
	fmt.Println("--- Processing with JSONProcessor ---")
	err := os.MkdirAll(outputDir, 0755)
	if err != nil {
		return nil, fmt.Errorf("failed to create output directory %s: %w", outputDir, err)
	}

	doc := p.CreateDocument(title, content)
	err = doc.Save(outputDir, filenameBase)
	if err != nil {
		return nil, err
	}
	fmt.Println("--- JSON processing complete ---")
	return doc, nil
}

// HTMLProcessor implements DocumentProcessor
type HTMLProcessor struct{}

func (p *HTMLProcessor) CreateDocument(title string, content []string) Document {
	fmt.Println("HTMLProcessor: Creating an HTML document.")
	return NewHTMLDocument(title, content)
}

func (p *HTMLProcessor) ProcessDocument(title string, content []string, outputDir string, filenameBase string) (Document, error) {
	fmt.Println("--- Processing with HTMLProcessor ---")
	err := os.MkdirAll(outputDir, 0755)
	if err != nil {
		return nil, fmt.Errorf("failed to create output directory %s: %w", outputDir, err)
	}

	doc := p.CreateDocument(title, content)
	err = doc.Save(outputDir, filenameBase)
	if err != nil {
		return nil, err
	}
	fmt.Println("--- HTML processing complete ---")
	return doc, nil
}
