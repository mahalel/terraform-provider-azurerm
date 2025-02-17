package simpolicy

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	"github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SIMPolicyClient struct {
	Client *resourcemanager.Client
}

func NewSIMPolicyClientWithBaseURI(api environments.Api) (*SIMPolicyClient, error) {
	client, err := resourcemanager.NewResourceManagerClient(api, "simpolicy", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating SIMPolicyClient: %+v", err)
	}

	return &SIMPolicyClient{
		Client: client,
	}, nil
}
