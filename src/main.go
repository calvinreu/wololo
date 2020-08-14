package main

func main() {
	gameData := CreateGame()
	running := true

	go gameData.Graphics.RunOutput(60, &running)
}
