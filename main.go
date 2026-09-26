package main
import ("time"
		"github.com/salemisaacbradley-max/pokedexcli/internal/pokecache")


func main() {
	cfg := config{commands: commands, 
				pokecache: pokecache.NewCache(5*time.Second),}
	REPL(&cfg)
}
