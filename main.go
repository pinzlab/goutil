package main

import (
	"fmt"

	"github.com/pinzlab/goutil/pg/track"
)

type ProfileType string

const (
	ProfilePersonal  ProfileType = "Personal"
	ProfileAgency    ProfileType = "Agency"
	ProfileDeveloper ProfileType = "Developer"
	ProfileBroker    ProfileType = "Broker"
)

type GQLProfileType string

const (
	GQLProfilePersonal  GQLProfileType = "Personal"
	GQLProfileAgency    GQLProfileType = "Agency"
	GQLProfileDeveloper GQLProfileType = "Developer"
	GQLProfileBroker    GQLProfileType = "Broker"
)

type Profile struct {
	UserID      int64
	ProfileType ProfileType
	Document    string
}

type GQLProfile struct {
	UserID      string
	ProfileType GQLProfileType
	Document    string
}

func main() {

	var profile Profile
	gql := GQLProfile{
		UserID:      "1",
		ProfileType: GQLProfilePersonal,
		Document:    "0604059741",
	}

	track.ToCreate(&gql, &profile, nil)

	fmt.Printf("%+v\n\n", profile)

}
