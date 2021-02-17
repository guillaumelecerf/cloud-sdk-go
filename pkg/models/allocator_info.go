// Licensed to Elasticsearch B.V. under one or more contributor
// license agreements. See the NOTICE file distributed with
// this work for additional information regarding copyright
// ownership. Elasticsearch B.V. licenses this file to you under
// the Apache License, Version 2.0 (the "License"); you may
// not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing,
// software distributed under the License is distributed on an
// "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
// KIND, either express or implied.  See the License for the
// specific language governing permissions and limitations
// under the License.

// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// AllocatorInfo The overview information for the allocator.
//
// swagger:model AllocatorInfo
type AllocatorInfo struct {

	// Identifier for this allocator
	// Required: true
	AllocatorID *string `json:"allocator_id"`

	// Build Info of the artifact
	BuildInfo *AllocatorBuildInfo `json:"build_info,omitempty"`

	// capacity
	// Required: true
	Capacity *AllocatorCapacity `json:"capacity"`

	// External resources related to this allocator.
	// Required: true
	// Unique: true
	ExternalLinks []*ExternalHyperlink `json:"external_links"`

	// List of features associated with this allocator. Note this is only present for backwards compatibility purposes and is scheduled for removal in the next major version release.
	// Required: true
	Features []string `json:"features"`

	// Host IP of this allocator
	// Required: true
	HostIP *string `json:"host_ip"`

	// instances
	// Required: true
	Instances []*AllocatedInstanceStatus `json:"instances"`

	// Arbitrary metadata associated with this allocator
	// Required: true
	Metadata []*MetadataItem `json:"metadata"`

	// Public hostname of this allocator
	// Required: true
	PublicHostname *string `json:"public_hostname"`

	// The region that this allocator belongs to. Only populated in SaaS or federated ECE.
	Region string `json:"region,omitempty"`

	// settings
	// Required: true
	Settings *AllocatorSettings `json:"settings"`

	// status
	// Required: true
	Status *AllocatorHealthStatus `json:"status"`

	// Identifier of the zone
	// Required: true
	ZoneID *string `json:"zone_id"`
}

// Validate validates this allocator info
func (m *AllocatorInfo) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAllocatorID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateBuildInfo(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCapacity(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateExternalLinks(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFeatures(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateHostIP(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateInstances(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMetadata(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePublicHostname(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSettings(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStatus(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateZoneID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AllocatorInfo) validateAllocatorID(formats strfmt.Registry) error {

	if err := validate.Required("allocator_id", "body", m.AllocatorID); err != nil {
		return err
	}

	return nil
}

func (m *AllocatorInfo) validateBuildInfo(formats strfmt.Registry) error {

	if swag.IsZero(m.BuildInfo) { // not required
		return nil
	}

	if m.BuildInfo != nil {
		if err := m.BuildInfo.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("build_info")
			}
			return err
		}
	}

	return nil
}

func (m *AllocatorInfo) validateCapacity(formats strfmt.Registry) error {

	if err := validate.Required("capacity", "body", m.Capacity); err != nil {
		return err
	}

	if m.Capacity != nil {
		if err := m.Capacity.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("capacity")
			}
			return err
		}
	}

	return nil
}

func (m *AllocatorInfo) validateExternalLinks(formats strfmt.Registry) error {

	if err := validate.Required("external_links", "body", m.ExternalLinks); err != nil {
		return err
	}

	if err := validate.UniqueItems("external_links", "body", m.ExternalLinks); err != nil {
		return err
	}

	for i := 0; i < len(m.ExternalLinks); i++ {
		if swag.IsZero(m.ExternalLinks[i]) { // not required
			continue
		}

		if m.ExternalLinks[i] != nil {
			if err := m.ExternalLinks[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("external_links" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *AllocatorInfo) validateFeatures(formats strfmt.Registry) error {

	if err := validate.Required("features", "body", m.Features); err != nil {
		return err
	}

	return nil
}

func (m *AllocatorInfo) validateHostIP(formats strfmt.Registry) error {

	if err := validate.Required("host_ip", "body", m.HostIP); err != nil {
		return err
	}

	return nil
}

func (m *AllocatorInfo) validateInstances(formats strfmt.Registry) error {

	if err := validate.Required("instances", "body", m.Instances); err != nil {
		return err
	}

	for i := 0; i < len(m.Instances); i++ {
		if swag.IsZero(m.Instances[i]) { // not required
			continue
		}

		if m.Instances[i] != nil {
			if err := m.Instances[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("instances" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *AllocatorInfo) validateMetadata(formats strfmt.Registry) error {

	if err := validate.Required("metadata", "body", m.Metadata); err != nil {
		return err
	}

	for i := 0; i < len(m.Metadata); i++ {
		if swag.IsZero(m.Metadata[i]) { // not required
			continue
		}

		if m.Metadata[i] != nil {
			if err := m.Metadata[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("metadata" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *AllocatorInfo) validatePublicHostname(formats strfmt.Registry) error {

	if err := validate.Required("public_hostname", "body", m.PublicHostname); err != nil {
		return err
	}

	return nil
}

func (m *AllocatorInfo) validateSettings(formats strfmt.Registry) error {

	if err := validate.Required("settings", "body", m.Settings); err != nil {
		return err
	}

	if m.Settings != nil {
		if err := m.Settings.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("settings")
			}
			return err
		}
	}

	return nil
}

func (m *AllocatorInfo) validateStatus(formats strfmt.Registry) error {

	if err := validate.Required("status", "body", m.Status); err != nil {
		return err
	}

	if m.Status != nil {
		if err := m.Status.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("status")
			}
			return err
		}
	}

	return nil
}

func (m *AllocatorInfo) validateZoneID(formats strfmt.Registry) error {

	if err := validate.Required("zone_id", "body", m.ZoneID); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *AllocatorInfo) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AllocatorInfo) UnmarshalBinary(b []byte) error {
	var res AllocatorInfo
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
