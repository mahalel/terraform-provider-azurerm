package caches

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	"github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CachesClient struct {
	Client *resourcemanager.Client
}

func NewCachesClientWithBaseURI(api environments.Api) (*CachesClient, error) {
	client, err := resourcemanager.NewResourceManagerClient(api, "caches", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating CachesClient: %+v", err)
	}

	return &CachesClient{
		Client: client,
	}, nil
}
