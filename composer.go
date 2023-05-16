package main

type ComposerAuthor struct {
	Name     string `json:"name,omitempty"`
	Email    string `json:"email,omitempty"`
	Homepage string `json:"homepage,omitempty"`
}

type ComposerFunding struct {
	Type string `json:"type,omitempty"`
	URL  string `json:"url,omitempty"`
}

type ComposerJson struct {
	Name        string              `json:"name,omitempty" validate:"regexp=^[a-z0-9]([_.-]?[a-z0-9]+)*/[a-z0-9](([_.]|-{1\\,2})?[a-z0-9]+)*$"`
	Description string              `json:"description,omitempty"`
	Keywords    []string            `json:"keywords,omitempty"`
	Homepage    string              `json:"homepage,omitempty"`
	Type        string              `json:"type,omitempty" validate:"regexp=^[a-z0-9-]+$"`
	License     string              `json:"license,omitempty"`
	Authors     []ComposerAuthor    `json:"authors,omitempty"`
	Support     []map[string]string `json:"support,omitempty"`
	Funding     []ComposerFunding   `json:"funding,omitempty"`
	Require     []map[string]string `json:"require,omitempty"`
	RequireDev  []map[string]string `json:"require-dev,omitempty"`

	// Config struct {
	// 	Platform struct {
	// 		Php string `json:"php,omitempty"`
	// 	} `json:"platform,omitempty"`
	// } `json:"config,omitempty"`
	// Suggest struct {
	// 	ExtZip     string `json:"ext-zip,omitempty"`
	// 	ExtOpenssl string `json:"ext-openssl,omitempty"`
	// } `json:"suggest,omitempty"`
	// Autoload struct {
	// 	Psr0 struct {
	// 		Composer string `json:"Composer,omitempty"`
	// 	} `json:"psr-0,omitempty"`
	// } `json:"autoload,omitempty"`
	// AutoloadDev struct {
	// 	Psr0 struct {
	// 		ComposerTest string `json:"Composer\Test,omitempty"`
	// 	} `json:"psr-0,omitempty"`
	// } `json:"autoload-dev,omitempty"`
	// Bin   string `json:"bin,omitempty"`
	// Extra struct {
	// 	BranchAlias struct {
	// 		DevMaster string `json:"dev-master,omitempty"`
	// 	} `json:"branch-alias,omitempty"`
	// } `json:"extra,omitempty"`
	// Scripts struct {
	// 	Test string `json:"test,omitempty"`
	// } `json:"scripts,omitempty"`
}
