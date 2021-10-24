package databackend

import (
	"github.com/sers-dev/kubetables/internal/databackend/kubernetes"
	"github.com/sers-dev/kubetables/internal/databackend/types"
	"os"
)

type DataBackend interface {
	List () (types.Ktbans, error)
	Watch (chan types.Event, chan os.Signal) ()
}

func CreateDataBackend () (dbe DataBackend, err error) {
	dbe, err = kubernetes.Initialize()
	if err != nil {
		return nil, err
	}
	return dbe, nil
}