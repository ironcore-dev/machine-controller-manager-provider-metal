// SPDX-FileCopyrightText: 2024 SAP SE or an SAP affiliate company and IronCore contributors
// SPDX-License-Identifier: Apache-2.0

package metal

import (
	"context"
	"fmt"

	"github.com/gardener/machine-controller-manager/pkg/util/provider/driver"
	"github.com/gardener/machine-controller-manager/pkg/util/provider/machinecodes/codes"
	"github.com/gardener/machine-controller-manager/pkg/util/provider/machinecodes/status"
	apiv1alpha1 "github.com/ironcore-dev/machine-controller-manager-provider-metal/pkg/api/v1alpha1"
	metalv1alpha1 "github.com/ironcore-dev/metal-operator/api/v1alpha1"
	"k8s.io/apimachinery/pkg/runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

const (
	defaultIgnitionKey     = "ignition.json"
	ShootNameLabelKey      = "shoot-name"
	ShootNamespaceLabelKey = "shoot-namespace"
)

var (
	fieldOwner = client.FieldOwner("mcm.ironcore.dev/field-owner")
)

type metalDriver struct {
	Schema         *runtime.Scheme
	metalClient    client.Client
	metalNamespace string
	csiDriverName  string
}

func (d *metalDriver) InitializeMachine(ctx context.Context, request *driver.InitializeMachineRequest) (*driver.InitializeMachineResponse, error) {
	return nil, status.Error(codes.Unimplemented, "IronCore Provider does not yet implement InitializeMachine")
}

func (d *metalDriver) GetVolumeIDs(_ context.Context, req *driver.GetVolumeIDsRequest) (*driver.GetVolumeIDsResponse, error) {
	return nil, status.Error(codes.Unimplemented, "IronCore Provider does not yet implement GetVolumeIDs")
}

// NewDriver returns a new Gardener ironcore driver object
func NewDriver(c client.Client, namespace, csiDriverName string) driver.Driver {
	return &metalDriver{
		metalClient:    c,
		metalNamespace: namespace,
		csiDriverName:  csiDriverName,
	}
}

func (d *metalDriver) GenerateMachineClassForMigration(_ context.Context, _ *driver.GenerateMachineClassForMigrationRequest) (*driver.GenerateMachineClassForMigrationResponse, error) {
	return &driver.GenerateMachineClassForMigrationResponse{}, nil
}

func getIgnitionNameForMachine(machineName string) string {
	return fmt.Sprintf("%s-%s", machineName, "ignition")
}

func getProviderIDForServerClaim(serverClaim *metalv1alpha1.ServerClaim) string {
	return fmt.Sprintf("%s://%s/%s", apiv1alpha1.ProviderName, serverClaim.Namespace, serverClaim.Name)
}
