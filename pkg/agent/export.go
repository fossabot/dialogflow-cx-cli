package agent

import (
	"os"

	"github.com/xavidop/dialogflow-cx-test-runner/internal/global"
	cxpkg "github.com/xavidop/dialogflow-cx-test-runner/pkg/cx"
)

func Export(locationID, projectID, agentName, output string) error {

	agentClient, err := cxpkg.CreateAgentGRPCClient(locationID)
	if err != nil {
		return err
	}
	defer agentClient.Close()

	agent, err := cxpkg.GetAgentIdByName(agentClient, agentName, projectID, locationID)
	if err != nil {
		return err
	}

	responseData, err := cxpkg.ExportAgentByFullName(agentClient, agent.GetName(), projectID, locationID)
	if err != nil {
		return err
	}

	err = os.WriteFile(output, responseData.GetAgentContent(), 0644)
	if err != nil {
		return err
	}
	global.Log.Infof("Agent exported to file: %v\n", output)

	return nil
}