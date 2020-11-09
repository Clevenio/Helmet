// Copyright 2021 Clivern. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package module

import (
	"github.com/spacewalkio/helmet/core/migration"
	"github.com/spacewalkio/helmet/core/model"
)

// CreateOAuthAccessData creates a new entity
func (db *Database) CreateOAuthAccessData(oauthAccessData *model.OAuthAccessData) *model.OAuthAccessData {
	db.Connection.Create(oauthAccessData)

	return oauthAccessData
}

// UpdateOAuthAccessDataByID updates an entity by ID
func (db *Database) UpdateOAuthAccessDataByID(oauthAccessData *model.OAuthAccessData) *model.OAuthAccessData {
	db.Connection.Save(&oauthAccessData)

	return oauthAccessData
}

// GetOAuthAccessDataByID gets an entity by id
func (db *Database) GetOAuthAccessDataByID(id int) model.OAuthAccessData {
	oauthAccessData := model.OAuthAccessData{}

	db.Connection.Where("id = ?", id).First(&oauthAccessData)

	return oauthAccessData
}

// GetOAuthAccessDataByKeys gets an entity by keys
func (db *Database) GetOAuthAccessDataByKeys(accessToken string) model.OAuthAccessData {
	oauthAccessData := model.OAuthAccessData{}

	db.Connection.Where(
		"access_token = ?",
		accessToken,
	).First(&oauthAccessData)

	return oauthAccessData
}

// DeleteOAuthAccessDataByID deletes an entity by id
func (db *Database) DeleteOAuthAccessDataByID(id int) {
	db.Connection.Unscoped().Where("id = ?", id).Delete(&migration.OAuthAccessData{})
}

// GetOAuthAccessDataItems gets oauth data items
func (db *Database) GetOAuthAccessDataItems() []model.OAuthAccessData {
	keys := []model.OAuthAccessData{}

	db.Connection.Select("*").Find(&keys)

	return keys
}
