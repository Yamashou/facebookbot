# go-cabocha

cabocha binding for golang.

## Installation


```bash
brew install cabocha #for mac osx
sudo apt-get install cabocha #for debian/ubuntu
go install "github.com/ledyba/go-cabocha"
```

## Usage

```go
import (
	cabocha "github.com/ledyba/go-cabocha"
	"log"
)

func main() {
	cabo = cabocha.MakeCabocha()
	sentence, err := cabocha.Parse("あなたとJava")
	if err != nil {
		log.Error(err)
	}
	for _,chunk := range sentence.Chunks {
		log.Printf("Id: %d",chunk.Id)
		log.Printf("Link: %d",chunk.Link)
		log.Printf("Head: %d",chunk.Head)
		log.Printf("Func: %d",chunk.Func)
		// and so on...
		for _,token := range chunk.Tokens {
			log.Printf("    Id: %d",token.Id)
			log.Printf("    Features: %s",chunk.Features)
			log.Printf("    %s",chunk.Body)
		}
	}
}
```
