package okt

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"path/filepath"
	"strings"

	"net/url"

	daml "github.com/psprings/go-daml"
	"gopkg.in/go-playground/validator.v9"
	"gopkg.in/yaml.v2"
)

const ()

func (fs FacilityStatus) IsEnum() bool {
	return true
}

func (at AccessType) IsEnum() bool {
	return true
}

// OKT :
type OKT struct {
	Name string `yaml:"name" daml:"name"  json:"name" validate:"required"`
	// Description : Definition: Description of the facility. | Format: Free text.
	Description string   `yaml:"description" daml:"description"  json:"description"`
	Location    Location `yaml:"location" daml:"location"  json:"location" validate:"required"`
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
	// Vehicles : the vehicles which are offered by the represented carrier. | Format : []Vehicle
	Vehicles []Vehicle `yaml:"vehicles" daml:"vehicles"  json:"vehicles"`
	// Services : the transportation services or offerings provided by the represented carrier. | Format : []Service
	Services []Service `yaml:"services" daml:"services"  json:"services"`
	// AreaOfService : a list of regions defined by geographical bounding where vehicles and services are available | Format : []GeoShape
	AreasOfService []GeoShape `yaml:"areasOfService" daml:"areasOfService"  json:"areasOfService"`
	// Permits : a list of permits or endorsements held which allow this carrier to operate vehicles or services in a given locality or country | Format : []Permit
	Permits []Permit `yaml:"permits" daml:"permits"  json:"permits"`
	// DateFounded : Definition: Date the facility was founded. | Format: Recommended practice is to use ISO 8601, i.e. the format YYYY-MM-DD. | Note: It is acceptable to include only the Year (YYYY) or year and month (YYYY-MM).
	// DateFounded time.Time `yaml:"date_founded" daml:"date_founded"  json:"date_founded"`
	DateFounded string `yaml:"date_founded" daml:"date_founded"  json:"date_founded"`
	// Equipment : Definition: The equipment available for use at the manufacturing facility. | Format: List the equipment available using the Equipment class.
	Equipment Equipment `yaml:"equipment" daml:"equipment"  json:"equipment"`
	// TypicalMaterials : Definition: Typical materials used by the facility. | Format: Uses the Materials class.
	TypicalMaterials []Material `yaml:"typical_materials" daml:"typical_materials"  json:"typical_materials"`
	// Certifications : Definition: Certifications obtained by the facility. | Format: List the certifications. | Note: Knowledge of these is imperative informal manufacturing and procurement. For example, aid agencies would be able to see which manufacturing facilities have particular manufacturing licenses, such as medical manufacturing.
	Certifications []Certification `yaml:"certifications" daml:"certifications"  json:"certifications"`
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
		"Location":       Location{},
		"What3Words":     What3Words{},
		"Address":        Address{},
		"GPS":            GPS{},
		"Agent":          Agent{},
		"URL":            URL(url.URL{}),
		"Contact":        Contact{},
		"SocialMedia":    SocialMedia{},
		"Skill":          Skill(""),
		"Equipment":      Equipment{},
		"Vehicles":       []Vehicle{},
		"Material":       Material{},
		"Services":       []Service{},
		"Permits":        []Permit{},
		"AreasOfService": []GeoShape{},
		"Certification":  Certification(""),
	}
}

type Material struct {
	// MaterialType : Definition: Type of material. | Format: Provide the Wikiepedia URL for the relevant material type. | Note: For instructions how to do this, please see section 3.5.
	// MaterialType URL `yaml:"material_type" daml:"material_type"  json:"material_type"`
	MaterialType string `yaml:"material_type" daml:"material_type"  json:"material_type"`
	// Manufacturer :
	Manufacturer     string
	Brand            string
	SupplierLocation Location
	// DefinedMaterialType MaterialType
}

type Service interface{}

type GeoShape interface{}

type Permit interface{}

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

type QuantitativeValue struct {
	MaxValue int    `yaml:"maxValue" daml:"maxValue" json:"maxValue"` // MaxValue : The upper value of some characteristic or property.
	MinValue int    `yaml:"minValue" daml:"minValue" json:"minValue"` // MinValue : The upper value of some characteristic or property.
	UnitCode string `yaml:"unitCode" daml:"unitCode" json:"unitCode"` // UnitCode : The unit of measurement given using the UN/CEFACT Common Code (3 characters) or a URL. Other codes than the UN/CEFACT Common Code may be used with a prefix followed by a colon.
	UnitText string `yaml:"unitText" daml:"unitText" json:"unitText"` // UnitText : A string or text indicating the unit of measurement. Useful if you cannot provide a standard unit code for unitCode.
	Value    string `yaml:"value" daml:"value" json:"value"`          // Value : The value of the quantitative value or property value node.
	// For QuantitativeValue and MonetaryAmount, the recommended type for values is 'Number'.
	// For PropertyValue, it can be 'Text;', 'Number', 'Boolean', or 'StructuredValue'.
	// Use values from 0123456789 (Unicode 'DIGIT ZERO' (U+0030) to 'DIGIT NINE' (U+0039)) rather than superficially similiar Unicode symbols.
	// Use '.' (Unicode 'FULL STOP' (U+002E)) rather than ',' to indicate a decimal point. Avoid using these symbols as a readability separator.
	ValueReference string `yaml:"valueReference" daml:"valueReference" json:"valueReference"` // ValueReference : A secondary value that provides additional information on the original value, e.g. a reference temperature or a type of measurement.
}

type Vehicle struct {
	AcclerationTime QuantitativeValue `yaml:"accelerationTime" daml:"accelerationTime" json:"accelerationTime"` // AcclerationTime : The time needed to accelerate the vehicle from a given start velocity to a given target velocity.
	BodyType        string            `yaml:"bodyType" daml:"bodyType" json:"bodyType"`                         // BodyType : Indicates the design and body style of the vehicle (e.g. station wagon, hatchback, etc.)
	CargoVolume     QuantitativeValue `yaml:"cargoVolume" daml:"cargoVolume" json:"cargoVolume"`                // CargoVolume : The available volume for cargo or luggage. For automobiles, this is usually the trunk volume.; Typical unit code(s): LTR for liters, FTQ for cubic foot/feet

}

func writeFile(filename string, content []byte) error {
	return ioutil.WriteFile(filename, content, 0644)
}

func Sample(outputDir string) {
	v := validator.New()
	_ = v.RegisterValidation("what3words", func(fl validator.FieldLevel) bool {
		return len(strings.Split(fl.Field().String(), ".")) == 3
	})

	okt := OKT{
		Contact: Agent{
			Name:    "Some Person",
			Website: "https://example.com",
		},
	}
	yamlSample, err := yaml.Marshal(&okt)
	if err != nil {
		log.Fatalf("error: %v", err)
	}
	writeFile(filepath.Join(outputDir, "okt.yaml"), yamlSample)
	jsonSample, err := json.Marshal(&okt)
	if err != nil {
		log.Fatalf("error: %v", err)
	}
	writeFile(filepath.Join(outputDir, "okt.json"), jsonSample)
	damlSample, err := daml.Marshal(okt, TypeMap)
	if err != nil {
		log.Fatalf("error: %v", err)
	}
	writeFile(filepath.Join(outputDir, "src/main/daml/OKT.daml"), damlSample)
}
