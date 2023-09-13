// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: BUSL-1.1

package types

import (
	"github.com/hashicorp/go-multierror"

	"github.com/hashicorp/consul/internal/resource"
	pbcatalog "github.com/hashicorp/consul/proto-public/pbcatalog/v2beta1"
	"github.com/hashicorp/consul/proto-public/pbresource"
)

const (
	VirtualIPsKind = "VirtualIPs"
)

var (
	VirtualIPsV2Beta1Type = &pbresource.Type{
		Group:        GroupName,
		GroupVersion: VersionV1Alpha1,
		Kind:         VirtualIPsKind,
	}

	VirtualIPsType = VirtualIPsV2Beta1Type
)

func RegisterVirtualIPs(r resource.Registry) {
	r.Register(resource.Registration{
		Type:     VirtualIPsV2Beta1Type,
		Proto:    &pbcatalog.VirtualIPs{},
		Scope:    resource.ScopeNamespace,
		Validate: ValidateVirtualIPs,
	})
}

func ValidateVirtualIPs(res *pbresource.Resource) error {
	var vips pbcatalog.VirtualIPs

	if err := res.Data.UnmarshalTo(&vips); err != nil {
		return resource.NewErrDataParse(&vips, err)
	}

	var err error
	for idx, ip := range vips.Ips {
		if vipErr := validateIPAddress(ip.Address); vipErr != nil {
			err = multierror.Append(err, resource.ErrInvalidListElement{
				Name:  "ips",
				Index: idx,
				Wrapped: resource.ErrInvalidField{
					Name:    "address",
					Wrapped: vipErr,
				},
			})
		}
	}
	return err
}
