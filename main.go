package main

import (
	"encoding/json"
	"fmt"
	"os/exec"
)

type Subscription struct {
	ID         string `json:"id"`
	Name       string `json:"name"`
	State      string `json:"state"`
	TenantID   string `json:"tenantId"`
	IsDefault  bool   `json:"isDefault"`
	CloudName  string `json:"cloudName"`
	HomeTenant string `json:"homeTenantId"`
}

func listSubscriptionsTable() error {
	cmd := exec.Command("az", "account", "list", "--output", "table")
	out, err := cmd.Output()
	if err != nil {
		return err
	}
	fmt.Println(string(out))
	return nil
}

func getSubscriptions() ([]Subscription, error) {
	cmd := exec.Command("az", "account", "list", "--output", "json")
	out, err := cmd.Output()
	if err != nil {
		return nil, err
	}
	var subs []Subscription
	err = json.Unmarshal(out, &subs)
	if err != nil {
		return nil, err
	}
	// for _, sub := range subs {
	// 	fmt.Printf("ID: %s, Name: %s, State: %s, TenantID: %s, IsDefault: %t, CloudName: %s, HomeTenant: %s\n",
	// 		sub.ID, sub.Name, sub.State, sub.TenantID, sub.IsDefault, sub.CloudName, sub.HomeTenant)
	// }
	return subs, nil
}

type AKSObject struct {
	Name               string `json:"name"`
	AutoUpgradeProfile struct {
		UpgradeChannel string `json:"upgradeChannel"`
	} `json:"autoUpgradeProfile"`
	CurrentKubernetesVersion string `json:"currentKubernetesVersion"`
	AgentPoolProfiles        []struct {
		Name             string `json:"name"`
		NodeImageVersion string `json:"nodeImageVersion"`
	} `json:"agentPoolProfiles"`
}

func listAKSObjects(sub Subscription) error {
	cmd := exec.Command("az", "aks", "list", "--subscription", sub.ID, "--output", "json")
	out, err := cmd.Output()
	if err != nil {
		return err
	}
	var aksObjs []AKSObject
	err = json.Unmarshal(out, &aksObjs)
	if err != nil {
		return err
	}
	fmt.Println("==== " + sub.Name + " ==== " + sub.ID + " ====")
	for _, aksObj := range aksObjs {
		fmt.Printf("Name: %s, AutoUpgradeProfile: %s, CurrentKubernetesVersion: %s\n",
			aksObj.Name, aksObj.AutoUpgradeProfile.UpgradeChannel, aksObj.CurrentKubernetesVersion)
		for _, agentPool := range aksObj.AgentPoolProfiles {
			fmt.Printf("Agent Pool Name: %s, NodeImageVersion: %s\n", agentPool.Name, agentPool.NodeImageVersion)
		}
	}
	return nil
}

func main() {
	err := listSubscriptionsTable()
	if err != nil {
		fmt.Println(err)
	}
	subs, err2 := getSubscriptions()

	if err2 != nil {
		fmt.Println(err2)
	}

	for _, sub := range subs {
		fmt.Printf("Subscription ID: %s\n", sub.ID)
		err := listAKSObjects(sub)
		if err != nil {
			fmt.Println(err)
		}
	}
}
