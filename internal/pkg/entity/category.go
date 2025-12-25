package entity

import "time"

type Category struct {
	ID   			int    		
	Nama 			string 		
	CreatedAt    	time.Time 	
	UpdatedAt    	time.Time 	
}