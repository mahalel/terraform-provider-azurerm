package clusters

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	"github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ClustersClient struct {
	Client *resourcemanager.Client
}

func NewClustersClientWithBaseURI(api environments.Api) (*ClustersClient, error) {
	client, err := resourcemanager.NewResourceManagerClient(api, "clusters", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating ClustersClient: %+v", err)
	}

	return &ClustersClient{
		Client: client,
	}, nil
}
