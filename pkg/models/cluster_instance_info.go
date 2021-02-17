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
	"encoding/json"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// ClusterInstanceInfo Information about each Kibana instance and APM Server in the Elasticsearch cluster.
//
// swagger:model ClusterInstanceInfo
type ClusterInstanceInfo struct {

	// The id of the allocator on which this instance is running (if the container is started or starting)
	AllocatorID string `json:"allocator_id,omitempty"`

	// Whether the container has started (does not tell you anything about the service -ie Elasticsearch- running inside the container)
	// Required: true
	ContainerStarted *bool `json:"container_started"`

	// disk
	Disk *ClusterInstanceDiskInfo `json:"disk,omitempty"`

	// Whether the instance is healthy (ie started and running)
	// Required: true
	Healthy *bool `json:"healthy"`

	// instance configuration
	InstanceConfiguration *ClusterInstanceConfigurationInfo `json:"instance_configuration,omitempty"`

	// Whether the instance is healthy (ie started and running)
	// Required: true
	InstanceName *string `json:"instance_name"`

	// Whether the service is is maintenance mode (meaning that the proxy is not routing external traffic to it)
	// Required: true
	MaintenanceMode *bool `json:"maintenance_mode"`

	// memory
	Memory *ClusterInstanceMemoryInfo `json:"memory,omitempty"`

	// A list of the node roles assigned to the service running in the instance. Currently populated only for Elasticsearch.
	NodeRoles []string `json:"node_roles"`

	// The service-specific (eg Elasticsearch) id of the node, if available
	ServiceID string `json:"service_id,omitempty"`

	// List of roles assigned to the service running in the instance. Currently only populated for Elasticsearch, with possible values: master,data,ingest,ml
	ServiceRoles []string `json:"service_roles"`

	// Whether the service launched inside the container -ie Elasticsearch- is actually running
	// Required: true
	ServiceRunning *bool `json:"service_running"`

	// The version of the service that the instance is running (eg Elasticsearch or Kibana), if available
	ServiceVersion string `json:"service_version,omitempty"`

	// The zone in which this instance is being allocated
	Zone string `json:"zone,omitempty"`
}

// Validate validates this cluster instance info
func (m *ClusterInstanceInfo) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateContainerStarted(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDisk(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateHealthy(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateInstanceConfiguration(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateInstanceName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMaintenanceMode(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMemory(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNodeRoles(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateServiceRunning(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ClusterInstanceInfo) validateContainerStarted(formats strfmt.Registry) error {

	if err := validate.Required("container_started", "body", m.ContainerStarted); err != nil {
		return err
	}

	return nil
}

func (m *ClusterInstanceInfo) validateDisk(formats strfmt.Registry) error {

	if swag.IsZero(m.Disk) { // not required
		return nil
	}

	if m.Disk != nil {
		if err := m.Disk.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("disk")
			}
			return err
		}
	}

	return nil
}

func (m *ClusterInstanceInfo) validateHealthy(formats strfmt.Registry) error {

	if err := validate.Required("healthy", "body", m.Healthy); err != nil {
		return err
	}

	return nil
}

func (m *ClusterInstanceInfo) validateInstanceConfiguration(formats strfmt.Registry) error {

	if swag.IsZero(m.InstanceConfiguration) { // not required
		return nil
	}

	if m.InstanceConfiguration != nil {
		if err := m.InstanceConfiguration.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("instance_configuration")
			}
			return err
		}
	}

	return nil
}

func (m *ClusterInstanceInfo) validateInstanceName(formats strfmt.Registry) error {

	if err := validate.Required("instance_name", "body", m.InstanceName); err != nil {
		return err
	}

	return nil
}

func (m *ClusterInstanceInfo) validateMaintenanceMode(formats strfmt.Registry) error {

	if err := validate.Required("maintenance_mode", "body", m.MaintenanceMode); err != nil {
		return err
	}

	return nil
}

func (m *ClusterInstanceInfo) validateMemory(formats strfmt.Registry) error {

	if swag.IsZero(m.Memory) { // not required
		return nil
	}

	if m.Memory != nil {
		if err := m.Memory.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("memory")
			}
			return err
		}
	}

	return nil
}

var clusterInstanceInfoNodeRolesItemsEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["master","ingest","ml","data_hot","data_content","data_warm","data_cold","remote_cluster_client","transform","voting_only"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		clusterInstanceInfoNodeRolesItemsEnum = append(clusterInstanceInfoNodeRolesItemsEnum, v)
	}
}

func (m *ClusterInstanceInfo) validateNodeRolesItemsEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, clusterInstanceInfoNodeRolesItemsEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *ClusterInstanceInfo) validateNodeRoles(formats strfmt.Registry) error {

	if swag.IsZero(m.NodeRoles) { // not required
		return nil
	}

	for i := 0; i < len(m.NodeRoles); i++ {

		// value enum
		if err := m.validateNodeRolesItemsEnum("node_roles"+"."+strconv.Itoa(i), "body", m.NodeRoles[i]); err != nil {
			return err
		}

	}

	return nil
}

func (m *ClusterInstanceInfo) validateServiceRunning(formats strfmt.Registry) error {

	if err := validate.Required("service_running", "body", m.ServiceRunning); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ClusterInstanceInfo) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ClusterInstanceInfo) UnmarshalBinary(b []byte) error {
	var res ClusterInstanceInfo
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
