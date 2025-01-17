package environment

import (
	"github.com/xavidop/dialogflow-cx-cli/internal/global"
	cxpkg "github.com/xavidop/dialogflow-cx-cli/pkg/cx"
)

func Update(name, description, locationID, projectID, agentName string, flowVersions []string) error {
	agentClient, err := cxpkg.CreateAgentRESTClient(locationID)
	if err != nil {
		return err
	}
	defer agentClient.Close()

	environmentClient, err := cxpkg.CreateEnvironmentGRPCClient(locationID)
	if err != nil {
		return err
	}
	defer environmentClient.Close()

	versionClient, err := cxpkg.CreateVersionRESTClient(locationID)
	if err != nil {
		return err
	}
	defer versionClient.Close()

	flowClient, err := cxpkg.CreateFlowRESTClient(locationID)
	if err != nil {
		return err
	}
	defer flowClient.Close()

	agent, err := cxpkg.GetAgentIdByName(agentClient, agentName, projectID, locationID)
	if err != nil {
		return err
	}

	_, err = cxpkg.UpdateEnvironment(environmentClient, versionClient, flowClient, agent, name, description, flowVersions)
	if err != nil {
		return err
	}

	global.Log.Infof("Environment updated:")

	return nil
}
