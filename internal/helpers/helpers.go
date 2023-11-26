package helpers

import (
    "os"
    "strings"
    "log"
    "path"
)

func ParsePath(p string) string {
    after, found := strings.CutPrefix(p, "~")
    if !found {
        return p
    }
    home, err := os.UserHomeDir()
    if err != nil {
        log.Fatal(err)
    }
    return path.Join(home, after)
}
