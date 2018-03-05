package google

import (
	"github.com/banzaicloud/banzai-types/constants"
	"github.com/pkg/errors"
	"strings"
)

var versionPrefix = "1.7."

type CreateClusterGoogle struct {
	Project string        `json:"project"`
	Node    *GoogleNode   `json:"node,omitempty"`
	Master  *GoogleMaster `json:"master,omitempty"`
}

type GoogleMaster struct {
	Version string `json:"version"`
}

type GoogleNode struct {
	Count          int    `json:"count"`
	Version        string `json:"version"`
	ServiceAccount string `json:"serviceAccount"`
}

type UpdateClusterGoogle struct {
	*GoogleNode   `json:"node,omitempty"`
	*GoogleMaster `json:"master,omitempty"`
}

func (g *CreateClusterGoogle) Validate() error {

	if g == nil {
		return errors.New("Google is <nil>")
	}

	if len(g.Project) == 0 {
		msg := "Project id is empty"
		return errors.New(msg)
	}

	if g.Node == nil {
		g.Node = &GoogleNode{
			Count: 1,
		}
	}

	if g.Master == nil {
		g.Master = &GoogleMaster{}
	}

	if strings.HasPrefix(g.Node.Version, versionPrefix) || strings.HasPrefix(g.Master.Version, versionPrefix) {
		return constants.ErrorWrongKubernetesVersion
	}

	if g.Master.Version != g.Node.Version {
		return constants.ErrorDifferentKubernetesVersion
	}

	if g.Node.Count == 0 {
		g.Node.Count = constants.GoogleDefaultNodeCount
	}

	return nil
}

// Validate validates the update request (only google part). If any of the fields is missing, the method fills
// with stored data.
func (a *UpdateClusterGoogle) Validate() error {

	// ---- [ Google field check ] ---- //
	if a == nil {
		return errors.New("'google' field is empty")
	}

	// check version prefix
	if (a.GoogleMaster != nil && strings.HasPrefix(a.GoogleMaster.Version, versionPrefix)) ||
		(a.GoogleNode != nil && strings.HasPrefix(a.GoogleNode.Version, versionPrefix)) {
		return constants.ErrorWrongKubernetesVersion
	}

	// check version equality
	if a.GoogleMaster != nil && a.GoogleNode != nil && a.GoogleMaster.Version != a.GoogleNode.Version {
		return constants.ErrorDifferentKubernetesVersion
	}

	return nil
}

type ClusterProfileGoogle struct {
	Master *GoogleMaster `json:"master,omitempty"`
	Node   *GoogleNode   `json:"node,omitempty"`
}
