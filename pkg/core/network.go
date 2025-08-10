package core

import "github.com/arian-press2015/uniac/internal/validators"

type Network struct {
	Name    string
	CIDR    string
	Region  string
	Subnets []Subnet
}

type Subnet struct {
	Name string
	CIDR string
}

func NewSubnet(configSubnet *validators.Subnet) *Subnet {
	return &Subnet{
		Name: configSubnet.Name,
		CIDR: configSubnet.CIDR,
	}
}

func NewNetwork(configNetwork *validators.Network) *Network {
	subnets := make([]Subnet, len(configNetwork.Subnets))
	for i, sub := range configNetwork.Subnets {
		subnets[i] = *NewSubnet(&sub)
	}

	return &Network{
		Name:    configNetwork.Name,
		CIDR:    configNetwork.CIDR,
		Region:  configNetwork.Region,
		Subnets: subnets,
	}
}
