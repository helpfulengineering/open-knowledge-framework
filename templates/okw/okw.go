package okw

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"path/filepath"
	"strings"

	"net/url"

	"reflect"

	daml "github.com/psprings/go-daml"
	"gopkg.in/go-playground/validator.v9"
	"gopkg.in/yaml.v2"
)

type Student struct {
	Fname  string
	Lname  string
	City   string
	Mobile int64
}

const (
	// Active :
	Active FacilityStatus = "Active"
	// Planned :
	Planned FacilityStatus = "Planned"
	// TemporaryClosure :
	TemporaryClosure FacilityStatus = "Temporary Closure"
	// Closed :
	Closed FacilityStatus = "Closed"
	// Restricted : only certain people (e.g. staff members) can use the equipment
	Restricted AccessType = "Restricted"
	// RestrictedWithPublicHours : the equipment can be used by the public during limited hours
	RestrictedWithPublicHours AccessType = "Restricted with public hours"
	// SharedSpace : the facility is a shared workspace where access is by qualifying criteria (e.g. rental of a desk or workspace)
	SharedSpace AccessType = "Shared space"
	// Public : anyone may use the equipment (e.g. training may be required and other restructions may apply)
	Public AccessType = "Public"
	// Membership : access requires membership, which is available to the public or a certain demographic
	Membership AccessType = "Membership"
)

func (fs FacilityStatus) IsEnum() bool {
	return true
}

func (fs FacilityStatus) Enum() []FacilityStatus {
	return []FacilityStatus{
		Active,
		Planned,
		TemporaryClosure,
		Closed,
	}
}

func (fs FacilityStatus) EnumOptions() []string {
	return []string{
		string(Active),
		string(Planned),
		string(TemporaryClosure),
		string(Closed),
	}
}

func (at AccessType) IsEnum() bool {
	return true
}

func (at AccessType) Enum() []AccessType {
	return []AccessType{
		Restricted,
		RestrictedWithPublicHours,
		SharedSpace,
		Public,
		Membership,
	}
}

func (at AccessType) EnumOptions() []string {
	return []string{
		string(Restricted),
		string(RestrictedWithPublicHours),
		string(SharedSpace),
		string(Public),
		string(Membership),
	}
}

// OKW :
type OKW struct {
	Name     string   `yaml:"name" daml:"name"  json:"name" validate:"required"`
	Location Location `yaml:"location" daml:"location"  json:"location" validate:"required"`
	// Owner : Definition: An Agent who owns or manages the facility. | Format: Uses the Agent class.
	Owner Agent `yaml:"owner" daml:"owner"  json:"owner"`
	// Contact : Definition: An Agent who is the contact for enquiries about making. | Format: Uses the Agent class.
	Contact Agent `yaml:"contact" daml:"contact"  json:"contact" validate:"required"`
	// Affiliations : Definition: The Agent(s) who the manufacturing facility is affiliated with. | Format: Uses the Agent class.
	Affiliations []Agent `yaml:"affiliations" daml:"affiliations"  json:"affiliations"`
	// FacilityStatus : Definition: Status of the facility. | Format: Use of one the following:
	FacilityStatus FacilityStatus `yaml:"facility_status" daml:"facility_status"  json:"facility_status"`
	// OpeningHours : Definition: Hours in which the facility operates. | Format: Free text.
	OpeningHours string `yaml:"opening_hours" daml:"opening_hours"  json:"opening_hours"`
	// Description : Definition: Description of the facility. | Format: Free text.
	Description string `yaml:"description" daml:"description"  json:"description"`
	// DateFounded : Definition: Date the facility was founded. | Format: Recommended practice is to use ISO 8601, i.e. the format YYYY-MM-DD. | Note: It is acceptable to include only the Year (YYYY) or year and month (YYYY-MM).
	// DateFounded time.Time `yaml:"date_founded" daml:"date_founded"  json:"date_founded"`
	DateFounded string `yaml:"date_founded" daml:"date_founded"  json:"date_founded"`
	// AccessType : Definition: How the manufacturing equipment is accessed.
	// Format: Use one of the following:
	// Restricted (only certain people (e.g. staff members) can use the equipment)
	// Restricted with public hours (the equipment can be used by the public during limited hours)
	// Shared space (the facility is a shared workspace where access is by qualifying criteria (e.g. rental of a desk or workspace))
	// Public (anyone may use the equipment (e.g. training may be required and other restructions may apply))
	// Membership (access requires membership, which is available to the public or a certain demographic)
	// Note: For facilities, use this field on a general-terms basis (i.e. if most equipment is available to members, but certain equipment requires staff to operate use Membership). This field can also be used as a property of individual equipment where a facility has different aspect types for different equipment.
	AccessType AccessType `yaml:"access_type" daml:"access_type"  json:"access_type"`
	// WheelchairAcessibility : Definition: Whether the manufacturing facility is wheelchair accessible. | Format: Free text.
	WheelchairAcessibility bool `yaml:"wheelchair_acessibility" daml:"wheelchair_acessibility"  json:"wheelchair_acessibility"`
	// Equipment : Definition: The equipment available for use at the manufacturing facility. | Format: List the equipment available using the Equipment class.
	Equipment Equipment `yaml:"equipment" daml:"equipment"  json:"equipment"`
	// ManufacturingProcesses : Definition: Manufacturing process the Equipment is capable of. | Format: Provide the Wikipedia URL for the relevant manufacturing process. | Note: For instructions how to do this, please see section 3.5.
	// ManufacturingProcesses URL `yaml:"manufacturing_process" daml:"manufacturing_process"  json:"manufacturing_process"`
	ManufacturingProcesses string `yaml:"manufacturing_processes" daml:"manufacturing_processes"  json:"manufacturing_processes"`
	// TypicalBatchSize : Definition: Typical batch size output. | Format:  Use one of the following:
	TypicalBatchSize TypicalBatchSize `yaml:"typical_batch_size" daml:"typical_batch_size"  json:"typical_batch_size"`
	// SizeFloorSize : Definition: The size or floor size of a manufacturing facility. | Format: Integer. Unit: square metres (sqm). | Note: This helps a prospective user gauge the scale of a manufacturing facility.
	SizeFloorSize int `yaml:"size_floor_size" daml:"size_floor_size"  json:"size_floor_size"`
	// StorageCapacity :
	StorageCapacity string `yaml:"storage_capacity" daml:"storage_capacity"  json:"storage_capacity"`
	// TypicalMaterials : Definition: Typical materials used by the facility. | Format: Uses the Materials class.
	TypicalMaterials []Material `yaml:"typical_materials" daml:"typical_materials"  json:"typical_materials"`
	// Certifications : Definition: Certifications obtained by the facility. | Format: List the certifications. | Note: Knowledge of these is imperative informal manufacturing and procurement. For example, aid agencies would be able to see which manufacturing facilities have particular manufacturing licenses, such as medical manufacturing.
	Certifications []Certification `yaml:"certifications" daml:"certifications"  json:"certifications"`
	// BackupGenerator : Definition: Whether a manufacturing facility has a backup generator. | Format: TRUE / FALSE | Note: Knowledge of this is particiularly useful in places where there are frequent power outages.
	BackupGenerator bool `yaml:"backup_generator" daml:"backup_generator"  json:"backup_generator"`
	// UninterruptedPowerSupply : Definition: Whether a manufacturing facility has an uninterrupted power supply. | Format: TRUE / FALSE
	UninterruptedPowerSupply bool `yaml:"uninterrupted_power_supply" daml:"uninterrupted_power_supply"  json:"uninterrupted_power_supply"`
	// RoadAccess : Definition: Whether a manufacturing facility has road access. | Format: TRUE / FALSE
	RoadAccess bool `yaml:"road_access" daml:"road_access"  json:"road_access"`
	// LoadingDock : Definition: Whether a manufacturing facility has a loading dock. | Format: TRUE / FALSE
	LoadingDock bool `yaml:"loading_dock" daml:"loading_dock"  json:"loading_dock"`
	// MaintenanceSchedule : Definition: The maintenance schedule of a manufacturing facility. | Format: Free text.
	MaintenanceSchedule string `yaml:"maintenance_schedule" daml:"maintenance_schedule"  json:"maintenance_schedule"`
	// TypicalProducts : Definition: Typical products produced by the facility. | Format: List the typical products produced.
	TypicalProducts []string `yaml:"typical_products" daml:"typical_products"  json:"typical_products"`
	// PartnerFunder : Definition: The Agent which partners or funds the facility. | Format: Uses the Agent class.
	PartnerFunder Agent `yaml:"partner_funder" daml:"partner_funder"  json:"partner_funder"`
	// CustomerReviews : Definition: Customer reviews of the facility. | Format: Free text.
	CustomerReviews []CustomerReview `yaml:"customer_reviews" daml:"customer_reviews"  json:"customer_reviews"`
}

// FacilityStatus : Definition: Status of the facility. | Format: Use of one the following: Active, Planned, Temporary Closure, Closed
type FacilityStatus string

// AccessType : How the manufacturing equipment is accessed. | Format: Use one of the following:
type AccessType string

type TypicalBatchSize string

// Enum : return enumeration options as slice of type
func (tbs TypicalBatchSize) Enum() []TypicalBatchSize {
	return []TypicalBatchSize{}
}

// EnumOptions : return enumeration options as slice of string
func (tbs TypicalBatchSize) EnumOptions() []string {
	return []string{}
}

// Location :  Definition: Location of the facility. | Format: Uses the Location class.
type Location struct {
	Address Address `yaml:"address" daml:"address"  json:"address"`
	GPS     GPS     `yaml:"gps" daml:"gps"  json:"gps"`
	// Directions : Definition: Directions to manufacturing facility, person or organisation. | Format: Free text. | Note: This qualitative data field may be helpful for a difficult to find location, or in an area where the standard address format is irrelevant.
	Directions string `yaml:"directions" daml:"directions"  json:"directions"`
	What3Words string `yaml:"what_3_words" daml:"what_3_words"  json:"what_3_words"`
}

// What3Words : Definition: What 3 Words phrase for location. | Format: State the What 3 Words phrase. | Note: Often informal settlements, or developing countries do not have street addresses, and communicating GPS coordinates can be tricky and error-prone. What 3 Words is an alternative geospatial address system.
type What3Words struct {
	Coordinates string `yaml:"coordinates" daml:"coordinates"  json:"coordinates" validate:"what3words"`
	Language    string `yaml:"language" daml:"language"  json:"language"`
}

// Address : Definition: Address relating to a manufacturing facility, person or organisation. | Format: Use the defined Address sub-properties
type Address struct {
	Number   string `yaml:"number" daml:"number"  json:"number"`
	Street   string `yaml:"street" daml:"street"  json:"street"`
	District string `yaml:"district" daml:"district"  json:"district"`
	City     string `yaml:"city" daml:"city"  json:"city"`
	Region   string `yaml:"region" daml:"region"  json:"region"`
	Country  string `yaml:"country" daml:"country"  json:"country"`
	Postcode string `yaml:"postcode" daml:"postcode"  json:"postcode"`
}

// GPS : Definition: The relevant GPS coordinates. | Format: Provide the relevant GPS coordinates, using Decimal Degrees.
type GPS struct {
	Latitude  float64 `yaml:"latitude" daml:"latitude"  json:"latitude"`
	Longitude float64 `yaml:"longitude" daml:"longitude"  json:"logitude"`
}

// Agent :
type Agent struct {
	Name     string   `yaml:"name" daml:"name"  json:"name" validate:"required"`
	Location Location `yaml:"location" daml:"location"  json:"location"`
	// ContactPerson : Definition: An Agent who is the key point of contact for a manufacturing facility or organisation. | Format: Provide the name of the Agent.
	ContactPerson string `yaml:"contact_person" daml:"contact_person"  json:"contact_person"`
	// Contact :
	Contact Contact `yaml:"contact" daml:"contact"  json:"contact"`
	// Would be cool to implement if Marshall to JSON and YAML would work the right way.
	// See type URL and MarhsallJSON() below
	// Website       URL         `yaml:"website" daml:"website"  json:"website"`
	Website     string      `yaml:"website" daml:"website"  json:"website"`
	SocialMedia SocialMedia `yaml:"social_media" daml:"social_media"  json:"social_media"`
}

// URL :
type URL url.URL

// type URL struct {
// 	url.URL
// }

func MarshallJSON(u URL) ([]byte, error) {
	// x := u.(url.URL)
	x := interface{}(u).(url.URL)
	return []byte(x.String()), nil
}

// Contact :
type Contact struct {
	// Landline : Definition: A landline telephone number to contact the facility, person or organisation. | Format: Provide the telephone number.
	Landline string `yaml:"landline" daml:"landline"  json:"landline"`
	// Mobile : Definition: A mobile telephone number to contact the facility, person or organisation. | Format: Provide the telephone number.
	Mobile string `yaml:"mobile" daml:"mobile"  json:"mobile"`
	// Fax : Definition: A fax number to contact the facility, person or organisation. | Format: Provide the fax number.
	Fax      string `yaml:"fax" daml:"fax"  json:"fax"`
	Email    string `yaml:"email" daml:"email"  json:"email"`
	WhatsApp string `yaml:"whatsapp" daml:"whatsapp"  json:"whatsapp"`
}

// SocialMedia :
type SocialMedia struct {
	Facebook  string   `yaml:"landline" daml:"landline"  json:"landline"`
	Twitter   string   `yaml:"twitter" daml:"twitter"  json:"twitter"`
	Instagram string   `yaml:"instagram" daml:"instagram"  json:"instagram"`
	OtherURLs []string `yaml:"other_urls" daml:"other_urls"  json:"other_urls"`
}

// Skill :
type Skill string

func (s Skill) Enum() []Skill {
	return []Skill{}
}

func (s Skill) EnumOptions() []string {
	return []string{}
}

// Equipment : Definition: The equipment available for use at the manufacturing facility. | Format: List the equipment available using the Equipment class.
type Equipment struct {
	// EquipmentType : Definition: Classification of Equipment. | Format: Provide the Wikipedia URL for the relevant Equipment Type. | Note: For instructions how to do this, please see section 3.5.
	// EquipmentType URL `yaml:"equipment_type" daml:"equipment_type"  json:"equipment_type"`
	EquipmentType string `yaml:"equipment_type" daml:"equipment_type"  json:"equipment_type"`
	// ManufacturingProcess : Definition: Manufacturing process the Equipment is capable of. | Format: Provide the Wikipedia URL for the relevant manufacturing process. | Note: For instructions how to do this, please see section 3.5.
	// ManufacturingProcess URL `yaml:"manufacturing_process" daml:"manufacturing_process"  json:"manufacturing_process"`
	ManufacturingProcess string `yaml:"manufacturing_process" daml:"manufacturing_process"  json:"manufacturing_process"`
	// Make : Definition: Make of the piece of equipment. | Format: Provide the make of the model. | Note: Provides detailed information about a piece of equipment/tool. For example, you can design generically for a 3D printer, or you can design for a specific make or model of 3D printer.
	Make string `yaml:"make" daml:"make"  json:"make"`
	// Model : Definition: Model of the piece of Equipment. | Format: Provide the name of the model.
	Model string `yaml:"model" daml:"model"  json:"model"`
	// SerialNumber : Definition: Serial number of the piece of Equipment. | Format: Provide the serial number of the Equipment.
	SerialNumber string `yaml:"serial_number" daml:"serial_number"  json:"serial_number"`
	// Location : Definition: Location of the equipment. | Format: Uses Location class.
	Location Location `yaml:"location" daml:"location"  json:"location"`
	// SkillsRequired : Identified as future work.
	SkillsRequired []Skill `yaml:"skills_required" daml:"skills_required"  json:"skills_required"`
	// Condition : Definition: The condition of the piece of equipment. | Format: State the condition of the piece of equipment. | Note: This provides a user with information surrounding the quality of a piece of equipment/tool, and whether it can complete the task they need it for.
	Condition string `yaml:"condition" daml:"condition"  json:"condition"`
}

func TypeMap() map[string]interface{} {
	// var intr interface{}
	return map[string]interface{}{
		"FacilityStatus": FacilityStatus(""),
		// "FacilityStatus": "",
		"AccessType": AccessType(""),
		// "AccessType": "",
		"TypicalBatchSize": TypicalBatchSize(""),
		// "TypicalBatchSize":    "",
		"Location":            Location{},
		"What3Words":          What3Words{},
		"Address":             Address{},
		"GPS":                 GPS{},
		"Agent":               Agent{},
		"URL":                 URL(url.URL{}),
		"Contact":             Contact{},
		"SocialMedia":         SocialMedia{},
		"Skill":               Skill(""),
		"Equipment":           Equipment{},
		"EquipmentProperties": EquipmentProperties{},
		"Material":            Material{},
		"MaterialType":        MaterialType(""),
		"CircularEconomy":     CircularEconomy{},
		"HumanCapacity":       HumanCapacity{},
		"InnovationSpace":     InnovationSpace{},
		"LearningResource":    LearningResource{},
		"CustomerReview":      CustomerReview{},
		"Certification":       Certification(""),
	}
}

type EquipmentProperties struct {
	// Axes : Definition: The bed size of a piece of equipment. | Format: Integer. Unit: mm.
	Axes int `yaml:"axes" daml:"axes"  json:"axes"`
	// BedSize : Definition: The bed size of a piece of equipment. | Format: Integer. Unit: mm.
	BedSize                int `yaml:"bed_size" daml:"bed_size"  json:"bed_size"`
	BendingLength          int
	BuildVolume            int
	ChuckJawDiameter       int
	ColletSize             int
	ComputerControlled     bool
	CrossSlideTravel       int
	DaylightOpening        int
	DistanceBetweenCentres int
	EjectorThreads         int
	ExtractionSystem       bool
	GantryMaterial         Material
}

type Material struct {
	// MaterialType : Definition: Type of material. | Format: Provide the Wikiepedia URL for the relevant material type. | Note: For instructions how to do this, please see section 3.5.
	// MaterialType URL `yaml:"material_type" daml:"material_type"  json:"material_type"`
	MaterialType string `yaml:"material_type" daml:"material_type"  json:"material_type"`
	// Manufacturer :
	Manufacturer        string
	Brand               string
	SupplierLocation    Location
	DefinedMaterialType MaterialType
}

// type MaterialType interface{}
type MaterialType string

type CircularEconomy struct {
	// CircularEconomy : Definition: Whether a manufacturing facility applies Circular Economy principles. | Format: TRUE / FALSE
	CircularEconomy bool `yaml:"circular_economy" daml:"circular_economy"  json:"circular_economy"`
	// Description : Definition: Definition of how Circular Economy principles are applied. | Format: Free text.
	Description string `yaml:"description" daml:"description"  json:"description"`
	// ByProducts : Definition: List of the by-products produced. | Format: Uses the Materials class.
	ByProducts []Material `yaml:"material" daml:"material"  json:"material"`
}

// HumanCapacity : Definition: The human capacity of the facility sub-properties.
type HumanCapacity struct {
	// Headcount : Definition: The headcount of the facility in FTE, using definition provided here. | Format: Integer.
	Headcount int    `yaml:"headcount" daml:"headcount"  json:"headcount"`
	Maker     string `yaml:"maker" daml:"maker"  json:"maker"`
}

type InnovationSpace struct {
	Staff             int                `yaml:"staff" daml:"staff"  json:"staff"`
	LearningResources []LearningResource `yaml:"learning_resources" daml:"learning_resources"  json:"learning_resources"`
	Services          []Service          `yaml:"services" daml:"services"  json:"services"`
	// Footfall : Definition: The footfall at a manufacturing facility. | Format: Integer. | Note: It is useful to help determine the scale of the manufacturing facility.
	Footfall    int  `yaml:"footfall" daml:"footfall"  json:"footfall"`
	Residencies bool `yaml:"residencies" daml:"residencies"  json:"residencies"`
}

type LearningResource struct {
}

type Service interface{}

type Certification string

// Enum : return enumeration options as slice of type
func (c Certification) Enum() []Certification {
	return []Certification{}
}

// EnumOptions : return enumeration options as slice of string
func (c Certification) EnumOptions() []string {
	return []string{}
}

type CustomerReview struct {
	Identifier string `yaml:"indentifier" daml:"indentifier" json:"indentifier"`
	Rating     int    `yaml:"rating" daml:"rating"  json:"rating" validate:"gte=1,lte=5"`
	Body       string `yaml:"body" daml:"body"  json:"body"`
}

func writeFile(filename string, content []byte) error {
	return ioutil.WriteFile(filename, content, 0644)
}

func Sample(outputDir string) {
	v := validator.New()
	_ = v.RegisterValidation("what3words", func(fl validator.FieldLevel) bool {
		return len(strings.Split(fl.Field().String(), ".")) == 3
	})

	foo := Active
	r := reflect.ValueOf(foo)
	enum := r.MethodByName("Enum")
	log.Printf("Is Enum?: %s", enum.IsValid())

	// u, err := url.Parse("https://google.com")
	// println(u.String() + " WOAH!")
	// var okw OKW
	// okw := OKW{
	// 	Owner: Agent{
	// 		Website: url.URL{
	// 			Scheme: "https",
	// 			Host:   "google.com",
	// 		},
	// 	},
	// }
	okw := OKW{
		Contact: Agent{
			Name:    "Some Person",
			Website: "https://example.com",
		},
	}
	yamlSample, err := yaml.Marshal(&okw)
	if err != nil {
		log.Fatalf("error: %v", err)
	}
	writeFile(filepath.Join(outputDir, "okw.yaml"), yamlSample)
	jsonSample, err := json.Marshal(&okw)
	if err != nil {
		log.Fatalf("error: %v", err)
	}
	writeFile(filepath.Join(outputDir, "okw.json"), jsonSample)
	damlSample, err := daml.Marshal(okw, TypeMap)
	if err != nil {
		log.Fatalf("error: %v", err)
	}
	// writeFile("okw.daml", damlSample)
	writeFile(filepath.Join(outputDir, "src/main/daml/OKW.daml"), damlSample)
}

// # Open Know Where
// ---
// name: # string | Definition: Name of the facility. | Format: Provide the name of the facility.
// location: # Location | Definition: Location of the facility. | Format: Uses the Location class.
//   address: # Address | Definition: Address relating to a manufacturing facility, person or organisation. | Format: Use the defined Address sub-properties
//     number: [value]
//     street: [value]
//     district: [value]
//     city: [value]
//     region: [value]
//     country: [value]
//     postcode: [value]
//   gps: [value] # GPS Coordinates | Definition: The relevant GPS coordinates. | Format: Provide the relevant GPS coordinates, using Decimal Degrees.
//   directions: [value] # text | Definition: Directions to manufacturing facility, person or organisation. | Format: Free text. | Note: This qualitative data field may be helpful for a difficult to find location, or in an area where the standard address format is irrelevant.
//   what_3_words: # What 3 Words | Definition: What 3 Words phrase for location. | Format: State the What 3 Words phrase. | Note: Often informal settlements, or developing countries do not have street addresses, and communicating GPS coordinates can be tricky and error-prone. What 3 Words is an alternative geospatial address system.
//     language: [value] # text | Definition: Language What 3 Words has been recorded in. | Format: ISO 639-2 or ISO 639-3, for example “en-gb”. | Note: What 3 Words is available in 43 different languages and the words for an address are not direct translations of each other.
// owner: [value] # Agent | Definition: An Agent who owns or manages the facility. | Format: Uses the Agent class.
// contact: [value] # Agent | Definition: An Agent who is the contact for enquiries about making. | Format: Uses the Agent class.
// agents: [value] # []Agent | Definition: The Agent(s) who the manufacturing facility is affiliated with. | Format: Uses the Agent class.

// ---
// # Location
// location: # Location | Definition: Location of the facility. | Format: Uses the Location class.
//   address: # Address | Definition: Address relating to a manufacturing facility, person or organisation. | Format: Use the defined Address sub-properties
//     number: [value]
//     street: [value]
//     district: [value]
//     city: [value]
//     region: [value]
//     country: [value]
//     postcode: [value]
// ---
// # Agent
// agent:
//   name: [value]
//   location: [value] # Location |
//   contact_person: [value] # Agent |
//   website: [value] # URL |
//   social_media: # Social Media |
// ---
// # Contact
// contact:
//   landline: [value]
//   mobile: [value]
//   fax: [value]
//   email: [value]
//   whatsapp: [value]
// ---
// # Social Media
// social_media:
//   facebook: [value]
//   twitter: [value]
//   instagram: [value]
//   other_urls: # [URL] |
//     - [value] # URL |
// ---
