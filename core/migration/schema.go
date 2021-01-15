// Copyright 2020 Clivern. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package migration

import (
	"encoding/json"

	"github.com/jinzhu/gorm"
)

// AuthenticationType type
type AuthenticationType int

const (
	// KeyAuthentication const
	KeyAuthentication AuthenticationType = iota

	// BasicAuthentication const
	BasicAuthentication

	// OAuthAuthentication const
	OAuthAuthentication
)

// Option struct
type Option struct {
	gorm.Model

	Key   string `json:"key"`
	Value string `json:"value"`
}

// LoadFromJSON update object from json
func (o *Option) LoadFromJSON(data []byte) error {
	err := json.Unmarshal(data, &o)

	if err != nil {
		return err
	}

	return nil
}

// ConvertToJSON convert object to json
func (o *Option) ConvertToJSON() (string, error) {
	data, err := json.Marshal(&o)

	if err != nil {
		return "", err
	}

	return string(data), nil
}

// AuthMethod struct
type AuthMethod struct {
	gorm.Model

	Type string `json:"type"`
	UUID string `json:"uuid"`
}

// LoadFromJSON update object from json
func (a *AuthMethod) LoadFromJSON(data []byte) error {
	err := json.Unmarshal(data, &a)

	if err != nil {
		return err
	}

	return nil
}

// ConvertToJSON convert object to json
func (a *AuthMethod) ConvertToJSON() (string, error) {
	data, err := json.Marshal(&a)

	if err != nil {
		return "", err
	}

	return string(data), nil
}

// AuthData struct
type AuthData struct {
	gorm.Model

	APIKey     string `json:"api_key"`
	MethodUUID string `json:"method_uuid"`
}

// LoadFromJSON update object from json
func (a *AuthData) LoadFromJSON(data []byte) error {
	err := json.Unmarshal(data, &a)

	if err != nil {
		return err
	}

	return nil
}

// ConvertToJSON convert object to json
func (a *AuthData) ConvertToJSON() (string, error) {
	data, err := json.Marshal(&a)

	if err != nil {
		return "", err
	}

	return string(data), nil
}