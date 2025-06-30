package dndcharacter

import (
    "math/rand"
    "time"
    "sort"
)

func init() {
    rand.Seed(time.Now().UnixNano())
}

type Character struct {
    Strength     int
    Dexterity    int
    Constitution int
    Intelligence int
    Wisdom       int
    Charisma     int
    Hitpoints    int
}

func Modifier(score int) int {
    if score >= 10 {
        return (score - 10) / 2
    }
    return (score - 11) / 2
}

func Ability() int {
    rolls := make([]int, 4)
    
    for i := 0; i < 4; i++ {
        rolls[i] = rand.Intn(6) + 1
    }
    
    sort.Slice(rolls, func(i, j int) bool {
        return rolls[i] > rolls[j]
    })
    
    return rolls[0] + rolls[1] + rolls[2]
}

func GenerateCharacter() Character {
    constitution := Ability()
    
    character := Character{
        Strength:     Ability(),
        Dexterity:    Ability(),
        Constitution: constitution,
        Intelligence: Ability(),
        Wisdom:       Ability(),
        Charisma:     Ability(),
        Hitpoints:    10 + Modifier(constitution), 
    }
    
    return character
}