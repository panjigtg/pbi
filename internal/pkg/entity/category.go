package entity

import "time"

type Category struct {
	ID   			int    		`gorm:"primaryKey"`	
	Nama 			string 		
	CreatedAt    	time.Time 	
	UpdatedAt    	time.Time 	
}