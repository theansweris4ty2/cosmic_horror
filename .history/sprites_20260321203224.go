package main 

import (
	
)



playerImg, _, err := ebitenutil.NewImageFromFile("Walk.png")
	if err != nil {
		fmt.Println(err)
	}
	playerImg2, _, err := ebitenutil.NewImageFromFile("Walk_left.png")
	if err != nil {
		fmt.Println(err)
	}
	playerImg3, _, err := ebitenutil.NewImageFromFile("Jump.png")
	if err != nil {
		fmt.Println(err)
	}
	playerImg4, _, err := ebitenutil.NewImageFromFile("Attack_1.png")
	if err != nil {
		fmt.Println(err)
	}
	playerImg5, _, err := ebitenutil.NewImageFromFile("Attack_1_back.png")
	if err != nil {
		fmt.Println(err)
	}