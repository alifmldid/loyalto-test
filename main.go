package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

func main() {
	var player_amount int
	var dice_amount int
	
	player_point := map[int]int{}
	player_dice := map[int][]int{}
	player_chance := map[int]int{}

	fmt.Print("Jumlah pemain: ")
	fmt.Scan(&player_amount)
	fmt.Print("Jumlah dadu: ")
	fmt.Scan(&dice_amount)

	for player_id:=1; player_id<=player_amount; player_id++{
		player_chance[player_id] = dice_amount
		player_point[player_id] = 0
	}
	
	var num_one = 0
	var check_dice = player_amount

	x := 0
	for check_dice > 0{
		x++
		for player_id:=1; player_id<=player_amount; player_id++{
			player_dice[player_id] = []int{}

			if num_one != 0 {
				for x := 1; x <= num_one; x++{
					player_dice[player_id] = append(player_dice[player_id], 1)
					num_one -= 1
				}
			}
			
			for dice_order:=1; dice_order<=player_chance[player_id]; dice_order++{
				rand.Seed(time.Now().UnixNano() + int64(dice_order) + int64(player_id))
				dice_numer := []rune("123456")

				rand.Shuffle(len(dice_numer), func(i, j int) {
					dice_numer[i], dice_numer[j] = dice_numer[j], dice_numer[i]
				})

				dice_num, _ := strconv.Atoi(string(dice_numer[dice_order]))

				if dice_num == 6 {
					player_point[player_id] += 1
				}else if dice_num == 1{
					if player_id != player_amount{
						num_one += 1
					}else{
						player_dice[1] = append(player_dice[1], dice_num)
						player_chance[1] = len(player_dice[1])
					}
				}else{
					player_dice[player_id] = append(player_dice[player_id], dice_num)
				}
			}
			player_chance[player_id] = len(player_dice[player_id])
		}

		fmt.Println("\nEVALUASI ", x)
		fmt.Println("Poin")
		for x, pp := range player_point {
			fmt.Println("Player ", x, ": ", pp)
		}
		fmt.Println("\nDadu")
		for x, pd := range player_dice {
			fmt.Println("Player ", x, ": ", pd)
			if len(pd) == 0 {
				check_dice -= 1
			}
		}
	}

	var winner_point = 0
	var winner int

	for player_id:=1; player_id<=player_amount; player_id++{
		if player_point[player_id] > winner_point {
			winner_point = player_point[player_id]
			winner = player_id
		}
	}

	fmt.Println("Pemenang adalah Player ", winner)
}