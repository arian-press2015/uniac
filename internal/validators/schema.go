package validators

import "github.com/mitchellh/mapstructure"

type Config struct {
	VM               []VM             `yaml:"vms" validate:"omitempty,dive"`
	Disk             []Disk           `yaml:"disks" validate:"omitempty,dive"`
	Database         []Database       `yaml:"databases" validate:"omitempty,dive"`
	FileStorage      []FileStorage    `yaml:"file_storage" validate:"omitempty,dive"`
	ObjectStorage    []ObjectStorage  `yaml:"object_storage" validate:"omitempty,dive"`
	CDN              []CDN            `yaml:"cdn" validate:"omitempty,dive"`
	DNS              []DNS            `yaml:"dns" validate:"omitempty,dive"`
	Network          []Network        `yaml:"network" validate:"omitempty,dive"`
	LoadBalancer     []LoadBalancer   `yaml:"load_balancer" validate:"omitempty,dive"`
	Firewall         []Firewall       `yaml:"firewall" validate:"omitempty,dive"`
	AccessManagement AccessManagement `yaml:"access_management" validate:"omitempty"`
}

type VM struct {
	Name    string            `yaml:"name" validate:"required"`
	Size    string            `yaml:"size" validate:"required,oneof=xs s m l xl"`
	Type    string            `yaml:"type" validate:"required,oneof=general-purpose gpu-optimized ram-optimized compute-optimized"`
	Region  string            `yaml:"region" validate:"required"`
	Network string            `yaml:"network" validate:"required"`
	Image   string            `yaml:"image" validate:"required"`
	Disks   []string          `yaml:"disks" validate:"dive,required"`
	Tags    map[string]string `yaml:"tags" validate:"-"`
}

type Disk struct {
	Name string            `yaml:"name" validate:"required"`
	Size string            `yaml:"size" validate:"required,storage_size"`
	Type string            `yaml:"type" validate:"required,oneof=hdd ssd nvme ultra-ssd"`
	Tags map[string]string `yaml:"tags" validate:"-"`
}

type Database struct {
	Name    string            `yaml:"name" validate:"required"`
	Storage string            `yaml:"storage" validate:"required,storage_size"`
	Region  string            `yaml:"region" validate:"required"`
	Engine  string            `yaml:"engine" validate:"required"`
	Version float64           `yaml:"version" validate:"required"`
	Tags    map[string]string `yaml:"tags" validate:"-"`
}

type FileStorage struct {
	Name     string            `yaml:"name" validate:"required"`
	Size     string            `yaml:"size" validate:"required,storage_size"`
	Type     string            `yaml:"type" validate:"required,oneof=standard premium"`
	Region   string            `yaml:"region" validate:"required"`
	Protocol string            `yaml:"protocol" validate:"required,oneof=NFS SMB"`
	Tags     map[string]string `yaml:"tags" validate:"-"`
}

type ObjectStorage struct {
	Name   string            `yaml:"name" validate:"required"`
	Size   string            `yaml:"size" validate:"required,storage_size"`
	Type   string            `yaml:"type" validate:"required,oneof=standard infrequent archive"`
	Region string            `yaml:"region" validate:"required"`
	Access string            `yaml:"access" validate:"required,oneof=public private"`
	Tags   map[string]string `yaml:"tags" validate:"-"`
}

type CDN struct {
	Name   string            `yaml:"name" validate:"required"`
	Origin string            `yaml:"origin" validate:"required"`
	Region string            `yaml:"region" validate:"required"`
	TTL    int               `yaml:"ttl" validate:"required,min=0"`
	SSL    bool              `yaml:"ssl"`
	Tags   map[string]string `yaml:"tags" validate:"-"`
}

type DNS struct {
	Name   string            `yaml:"name" validate:"required"`
	Type   string            `yaml:"type" validate:"required,oneof=A CNAME MX"`
	Value  string            `yaml:"value" validate:"required"`
	TTL    int               `yaml:"ttl" validate:"required,min=0"`
	Region string            `yaml:"region" validate:"required"`
	Tags   map[string]string `yaml:"tags" validate:"-"`
}

type Network struct {
	Name    string   `yaml:"name" validate:"required"`
	CIDR    string   `yaml:"cidr" validate:"required,cidr"`
	Region  string   `yaml:"region" validate:"required"`
	Subnets []Subnet `yaml:"subnets" validate:"dive"`
}

type Subnet struct {
	Name string `yaml:"name" validate:"required"`
	CIDR string `yaml:"cidr" validate:"required,cidr"`
}

type LoadBalancer struct {
	Name    string            `yaml:"name" validate:"required"`
	Type    string            `yaml:"type" validate:"required,oneof=application network"`
	Region  string            `yaml:"region" validate:"required"`
	Port    int               `yaml:"port" validate:"required,gt=0,lte=65535"`
	Targets []Target          `yaml:"targets" validate:"dive"`
	Tags    map[string]string `yaml:"tags" validate:"-"`
}

type Target struct {
	Name string `yaml:"name" validate:"required"`
	Port int    `yaml:"port" validate:"required,gt=0,lte=65535"`
}

type Firewall struct {
	Name    string         `yaml:"name" validate:"required"`
	Type    string         `yaml:"type" validate:"required,oneof=security-group waf network-acl"`
	Network string         `yaml:"network" validate:"required"`
	Rules   []FirewallRule `yaml:"rules" validate:"dive"`
}

type FirewallRule struct {
	Type      string `yaml:"type" validate:"required,oneof=ingress egress block"`
	Protocol  string `yaml:"protocol" validate:"omitempty,oneof=tcp udp"`
	Port      int    `yaml:"port" validate:"omitempty,gt=0,lte=65535"`
	Source    string `yaml:"source" validate:"omitempty,cidr"`
	Condition string `yaml:"condition" validate:"omitempty"` // For WAF
}

type AccessManagement struct {
	Roles []Role `yaml:"roles" validate:"required,dive"`
	Users []User `yaml:"users" validate:"required,dive"`
}

type Role struct {
	Name        string   `yaml:"name" validate:"required"`
	Permissions []string `yaml:"permissions" validate:"dive,required"`
}

type User struct {
	Name string            `yaml:"name" validate:"required"`
	Role string            `yaml:"role" validate:"required"`
	Tags map[string]string `yaml:"tags" validate:"-"` // Exempt from validation
}

func NewConfig(data map[string]interface{}) (*Config, error) {
	var config Config
	decoder, err := mapstructure.NewDecoder(&mapstructure.DecoderConfig{
		TagName: "yaml",
		Result:  &config,
	})
	if err != nil {
		return nil, err
	}

	err = decoder.Decode(data)
	if err != nil {
		return nil, err
	}

	return &config, nil
}
