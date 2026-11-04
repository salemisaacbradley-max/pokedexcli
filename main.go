package main
import ("time"
		"github.com/salemisaacbradley-max/pokedexcli/internal/pokeapi")


func main() {
	cfg := config{commands: commands, 
				pokecache: pokeapi.NewCache(5*time.Second),
				caughtPokemon: make(map[string]Pokemon) }
	REPL(&cfg)
}
