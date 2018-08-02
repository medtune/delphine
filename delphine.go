package delphine

import (
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
)

var c = client.Client{}

func a() {
	t := types.ImagePushOptions{}
}
