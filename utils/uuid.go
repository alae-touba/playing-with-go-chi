package utils

import (
    "github.com/google/uuid"
)

// ParseUUID validates and parses a string into UUID
// Returns uuid.Nil if string is empty
func ParseUUID(id string) (uuid.UUID, error) {
    if id == "" {
        return uuid.Nil, nil
    }
    
    parsedID, err := uuid.Parse(id)
    if err != nil {
        return uuid.Nil, err
    }
    
    return parsedID, nil
}