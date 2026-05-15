package dto

// This file contains the Data Transfer Object (DTO) definitions for the Role entity.
type RoleDTO struct {
	Name string `json:"name" binding:"required"`
}
