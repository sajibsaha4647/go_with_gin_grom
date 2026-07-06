package dto

type UserUpdateRequest struct {
	Name     *string `json:"name" binding:"omitempty"`
    // Can be sent as "avatar.png", sent as null, or omitted entirely.
    Image    *string `json:"image"`
}