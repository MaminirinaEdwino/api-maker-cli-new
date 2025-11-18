package goapi

import (
	"fmt"
	"github/mameinirinaedwino/api-maker-cli/module/utils"
)

func WriteBodyDecodeur(endPoint string) string {
	return fmt.Sprintf("var body %sbodyType \ndecoder := json.NewDecoder(r.Body) \nerr := decoder.Decode(&body)\n%s", endPoint, utils.WriteErrorCheker("Parsing Error"))
}

