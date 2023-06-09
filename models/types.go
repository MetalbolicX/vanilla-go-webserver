package models

type BioMechanicalMovement struct {
	BioMechanicalMovementName string
}

type GymExercise struct {
	GymExerciseName string `json:"gym_exercise_name"`
	ResistanceCurve string `json:"resistance_curve"`
}

type IndexPageInfo struct {
	ListBioMechanicalMovements []BioMechanicalMovement
	ListGymExercises           []GymExercise
}
