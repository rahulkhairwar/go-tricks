package factoryPattern

import (
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
)

// the super type for all region-wise compliance objects.
type Compliance interface {
	// can add compliance-specific functions here.
	String() string
}

// the base struct.
type Component struct {
	Id          string     `json:"id" bson:"_id,omitempty"`
	Name        string     `json:"name" bson:"name"`
	Code        string     `json:"code" bson:"code"`
	CountryCode string     `json:"countryCode" bson:"countryCode"`
	Compliance  Compliance `json:"compliance" bson:"compliance"`
}

// UnmarshalBSON is the custom unmarshaller for Component.
func (c *Component) UnmarshalBSON(b []byte) error {
	base := &struct {
		Id          string `json:"id" bson:"_id,omitempty"`
		Name        string `json:"name" bson:"name"`
		Code        string `json:"code" bson:"code"`
		CountryCode string `bson:"countryCode"`
	}{}
	if err := bson.Unmarshal(b, base); err != nil {
		return err
	}

	c.Id = base.Id
	c.Name = base.Name
	c.Code = base.Code
	c.CountryCode = base.CountryCode

	switch base.CountryCode {
	case "SG":
		internal := &struct {
			*SingaporeCompliance `bson:"compliance"`
		}{}
		if err := bson.Unmarshal(b, internal); err != nil {
			return err
		}
		c.Compliance = internal.SingaporeCompliance
	case "IN":
		internal := &struct {
			*IndiaCompliance `bson:"compliance"`
		}{}
		if err := bson.Unmarshal(b, internal); err != nil {
			return err
		}
		c.Compliance = internal.IndiaCompliance
	}
	return nil
}

type SingaporeCompliance struct {
	CountryCode string `json:"countryCode" bson:"countryCode"`
	SomeInfo    string `json:"someInfo" bson:"someInfo"`
	ValidForCPF bool   `json:"validForCPF" bson:"validForCPF"`
	SDL         bool   `json:"sdl" bson:"sdl"`
	SHG         bool   `json:"shg" bson:"shg"`
}

func (s *SingaporeCompliance) String() string {
	return fmt.Sprintf("CountryCode : %s, SomeInfo : %s, ValidForCPF : %t, SDL : %t, SHG : %t\n", s.CountryCode, s.SomeInfo, s.ValidForCPF, s.SDL, s.SHG)
}

type IndiaCompliance struct {
	CountryCode       string `json:"countryCode" bson:"countryCode"`
	IndiaSpecificInfo string `json:"indiaSpecificInfo" bson:"indiaSpecificInfo"`
}

func (i *IndiaCompliance) String() string {
	return fmt.Sprintf("CountryCode : %s, IndiaSpecificInfo : %s", i.CountryCode, i.IndiaSpecificInfo)
}
