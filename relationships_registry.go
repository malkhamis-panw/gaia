package gaia

import "go.aporeto.io/elemental"

var relationshipsRegistry elemental.RelationshipsRegistry

func init() {

	relationshipsRegistry = elemental.RelationshipsRegistry{}

	relationshipsRegistry[APIAuthorizationPolicyIdentity] = &elemental.Relationship{
		Create: map[string]*elemental.RelationshipInfo{
			"root": {},
		},
		Update: map[string]*elemental.RelationshipInfo{
			"root": {},
		},
		Patch: map[string]*elemental.RelationshipInfo{
			"root": {},
		},
		Delete: map[string]*elemental.RelationshipInfo{
			"root": {
				Parameters: []elemental.ParameterDefinition{
					{
						Name:     "q",
						Type:     "string",
						Multiple: true,
					},
				},
			},
		},
		Retrieve: map[string]*elemental.RelationshipInfo{
			"root": {},
		},
		RetrieveMany: map[string]*elemental.RelationshipInfo{
			"root": {
				Parameters: []elemental.ParameterDefinition{
					{
						Name:     "q",
						Type:     "string",
						Multiple: true,
					},
					{
						Name: "propagated",
						Type: "boolean",
					},
				},
			},
		},
		Info: map[string]*elemental.RelationshipInfo{
			"root": {
				Parameters: []elemental.ParameterDefinition{
					{
						Name:     "q",
						Type:     "string",
						Multiple: true,
					},
					{
						Name: "propagated",
						Type: "boolean",
					},
				},
			},
		},
	}

	relationshipsRegistry[APICheckIdentity] = &elemental.Relationship{
		Create: map[string]*elemental.RelationshipInfo{
			"root": {},
		},
	}

	relationshipsRegistry[AccessReportIdentity] = &elemental.Relationship{
		Create: map[string]*elemental.RelationshipInfo{
			"root": {},
		},
	}

	relationshipsRegistry[AccessibleNamespaceIdentity] = &elemental.Relationship{
		RetrieveMany: map[string]*elemental.RelationshipInfo{
			"root": {
				Parameters: []elemental.ParameterDefinition{
					{
						Name:     "q",
						Type:     "string",
						Multiple: true,
					},
				},
			},
		},
		Info: map[string]*elemental.RelationshipInfo{
			"root": {
				Parameters: []elemental.ParameterDefinition{
					{
						Name:     "q",
						Type:     "string",
						Multiple: true,
					},
				},
			},
		},
	}

	relationshipsRegistry[AccountIdentity] = &elemental.Relationship{
		Create: map[string]*elemental.RelationshipInfo{
			"root": {},
		},
		Update: map[string]*elemental.RelationshipInfo{
			"root": {},
		},
		Patch: map[string]*elemental.RelationshipInfo{
			"root": {},
		},
		Delete: map[string]*elemental.RelationshipInfo{
			"root": {
				Parameters: []elemental.ParameterDefinition{
					{
						Name:     "q",
						Type:     "string",
						Multiple: true,
					},
				},
			},
		},
		Retrieve: map[string]*elemental.RelationshipInfo{
			"root": {},
		},
		RetrieveMany: map[string]*elemental.RelationshipInfo{
			"root": {
				Parameters: []elemental.ParameterDefinition{
					{
						Name: "name",
						Type: "string",
					},
					{
						Name: "status",
						Type: "string",
					},
					{
						Name:     "q",
						Type:     "string",
						Multiple: true,
					},
				},
			},
		},
		Info: map[string]*elemental.RelationshipInfo{
			"root": {
				Parameters: []elemental.ParameterDefinition{
					{
						Name: "name",
						Type: "string",
					},
					{
						Name: "status",
						Type: "string",
					},
					{
						Name:     "q",
						Type:     "string",
						Multiple: true,
					},
				},
			},
		},
	}

	relationshipsRegistry[AccountCheckIdentity] = &elemental.Relationship{
		Create: map[string]*elemental.RelationshipInfo{
			"root": {},
		},
	}

	relationshipsRegistry[ActivateIdentity] = &elemental.Relationship{
		RetrieveMany: map[string]*elemental.RelationshipInfo{
			"root": {
				RequiredParameters: elemental.NewParametersRequirement(
					[][][]string{
						{
							{
								"token",
							},
						},
					},
				),
				Parameters: []elemental.ParameterDefinition{
					{
						Name: "noRedirect",
						Type: "boolean",
					},
					{
						Name: "token",
						Type: "string",
					},
				},
			},
		},
		Info: map[string]*elemental.RelationshipInfo{
			"root": {
				RequiredParameters: elemental.NewParametersRequirement(
					[][][]string{
						{
							{
								"token",
							},
						},
					},
				),
				Parameters: []elemental.ParameterDefinition{
					{
						Name: "noRedirect",
						Type: "boolean",
					},
					{
						Name: "token",
						Type: "string",
					},
				},
			},
		},
	}

	relationshipsRegistry[ActivityIdentity] = &elemental.Relationship{
		Retrieve: map[string]*elemental.RelationshipInfo{
			"root": {},
		},
		RetrieveMany: map[string]*elemental.RelationshipInfo{
			"root": {
				Parameters: []elemental.ParameterDefinition{
					{
						Name:     "q",
						Type:     "string",
						Multiple: true,
					},
				},
			},
		},
		Info: map[string]*elemental.RelationshipInfo{
			"root": {
				Parameters: []elemental.ParameterDefinition{
					{
						Name:     "q",
						Type:     "string",
						Multiple: true,
					},
				},
			},
		},
	}

	relationshipsRegistry[AlarmIdentity] = &elemental.Relationship{
		Create: map[string]*elemental.RelationshipInfo{
			"root": {},
		},
		Update: map[string]*elemental.RelationshipInfo{
			"root": {},
		},
		Patch: map[string]*elemental.RelationshipInfo{
			"root": {},
		},
		Delete: map[string]*elemental.RelationshipInfo{
			"root": {
				Parameters: []elemental.ParameterDefinition{
					{
						Name:     "q",
						Type:     "string",
						Multiple: true,
					},
				},
			},
		},
		Retrieve: map[string]*elemental.RelationshipInfo{
			"root": {},
		},
		RetrieveMany: map[string]*elemental.RelationshipInfo{
			"root": {
				Parameters: []elemental.ParameterDefinition{
					{
						Name:     "q",
						Type:     "string",
						Multiple: true,
					},
				},
			},
		},
		Info: map[string]*elemental.RelationshipInfo{
			"root": {
				Parameters: []elemental.ParameterDefinition{
					{
						Name:     "q",
						Type:     "string",
						Multiple: true,
					},
				},
			},
		},
	}

	relationshipsRegistry[AppIdentity] = &elemental.Relationship{
		RetrieveMany: map[string]*elemental.RelationshipInfo{
			"root": {
				Parameters: []elemental.ParameterDefinition{
					{
						Name: "name",
						Type: "string",
					},
					{
						Name:     "q",
						Type:     "string",
						Multiple: true,
					},
				},
			},
		},
		Info: map[string]*elemental.RelationshipInfo{
			"root": {
				Parameters: []elemental.ParameterDefinition{
					{
						Name: "name",
						Type: "string",
					},
					{
						Name:     "q",
						Type:     "string",
						Multiple: true,
					},
				},
			},
		},
	}

	relationshipsRegistry[AppCredentialIdentity] = &elemental.Relationship{
		Create: map[string]*elemental.RelationshipInfo{
			"root": {},
		},
		Update: map[string]*elemental.RelationshipInfo{
			"root": {},
		},
		Patch: map[string]*elemental.RelationshipInfo{
			"root": {},
		},
		Delete: map[string]*elemental.RelationshipInfo{
			"root": {
				Parameters: []elemental.ParameterDefinition{
					{
						Name:     "q",
						Type:     "string",
						Multiple: true,
					},
				},
			},
		},
		Retrieve: map[string]*elemental.RelationshipInfo{
			"root": {},
		},
		RetrieveMany: map[string]*elemental.RelationshipInfo{
			"root": {
				Parameters: []elemental.ParameterDefinition{
					{
						Name:     "q",
						Type:     "string",
						Multiple: true,
					},
				},
			},
		},
		Info: map[string]*elemental.RelationshipInfo{
			"root": {
				Parameters: []elemental.ParameterDefinition{
					{
						Name:     "q",
						Type:     "string",
						Multiple: true,
					},
				},
			},
		},
	}

	relationshipsRegistry[AuditProfileIdentity] = &elemental.Relationship{
		Create: map[string]*elemental.RelationshipInfo{
			"root": {},
		},
		Update: map[string]*elemental.RelationshipInfo{
			"root": {},
		},
		Patch: map[string]*elemental.RelationshipInfo{
			"root": {},
		},
		Delete: map[string]*elemental.RelationshipInfo{
			"root": {
				Parameters: []elemental.ParameterDefinition{
					{
						Name:     "q",
						Type:     "string",
						Multiple: true,
					},
				},
			},
		},
		Retrieve: map[string]*elemental.RelationshipInfo{
			"root": {
				Parameters: []elemental.ParameterDefinition{
					{
						Name: "propagated",
						Type: "boolean",
					},
				},
			},
		},
		RetrieveMany: map[string]*elemental.RelationshipInfo{
			"auditprofilemappingpolicy": {},
			"enforcer":                  {},
			"root": {
				Parameters: []elemental.ParameterDefinition{
					{
						Name:     "q",
						Type:     "string",
						Multiple: true,
					},
					{
						Name: "propagated",
						Type: "boolean",
					},
				},
			},
		},
		Info: map[string]*elemental.RelationshipInfo{
			"auditprofilemappingpolicy": {},
			"enforcer":                  {},
			"root": {
				Parameters: []elemental.ParameterDefinition{
					{
						Name:     "q",
						Type:     "string",
						Multiple: true,
					},
					{
						Name: "propagated",
						Type: "boolean",
					},
				},
			},
		},
	}

	relationshipsRegistry[AuditProfileMappingPolicyIdentity] = &elemental.Relationship{
		Create: map[string]*elemental.RelationshipInfo{
			"root": {},
		},
		Update: map[string]*elemental.RelationshipInfo{
			"root": {},
		},
		Patch: map[string]*elemental.RelationshipInfo{
			"root": {},
		},
		Delete: map[string]*elemental.RelationshipInfo{
			"root": {
				Parameters: []elemental.ParameterDefinition{
					{
						Name:     "q",
						Type:     "string",
						Multiple: true,
					},
				},
			},
		},
		Retrieve: map[string]*elemental.RelationshipInfo{
			"root": {
				Parameters: []elemental.ParameterDefinition{
					{
						Name: "propagated",
						Type: "boolean",
					},
				},
			},
		},
		RetrieveMany: map[string]*elemental.RelationshipInfo{
			"root": {
				Parameters: []elemental.ParameterDefinition{
					{
						Name:     "q",
						Type:     "string",
						Multiple: true,
					},
					{
						Name: "propagated",
						Type: "boolean",
					},
				},
			},
		},
		Info: map[string]*elemental.RelationshipInfo{
			"root": {
				Parameters: []elemental.ParameterDefinition{
					{
						Name:     "q",
						Type:     "string",
						Multiple: true,
					},
					{
						Name: "propagated",
						Type: "boolean",
					},
				},
			},
		},
	}

	relationshipsRegistry[AuditReportIdentity] = &elemental.Relationship{
		Create: map[string]*elemental.RelationshipInfo{
			"root": {},
		},
	}

	relationshipsRegistry[AuthnIdentity] = &elemental.Relationship{
		Create: map[string]*elemental.RelationshipInfo{
			"root": {},
		},
		RetrieveMany: map[string]*elemental.RelationshipInfo{
			"root": {
				Parameters: []elemental.ParameterDefinition{
					{
						Name: "token",
						Type: "string",
					},
				},
			},
		},
		Info: map[string]*elemental.RelationshipInfo{
			"root": {
				Parameters: []elemental.ParameterDefinition{
					{
						Name: "token",
						Type: "string",
					},
				},
			},
		},
	}

	relationshipsRegistry[AuthorityIdentity] = &elemental.Relationship{
		Create: map[string]*elemental.RelationshipInfo{
			"root": {
				Parameters: []elemental.ParameterDefinition{
					{
						Name:     "q",
						Type:     "string",
						Multiple: true,
					},
				},
			},
		},
		Delete: map[string]*elemental.RelationshipInfo{
			"root": {
				Parameters: []elemental.ParameterDefinition{
					{
						Name:     "q",
						Type:     "string",
						Multiple: true,
					},
				},
			},
		},
	}

	relationshipsRegistry[AuthzIdentity] = &elemental.Relationship{
		Create: map[string]*elemental.RelationshipInfo{
			"root": {
				Parameters: []elemental.ParameterDefinition{
					{
						Name: "forwardpolicyrules",
						Type: "boolean",
					},
				},
			},
		},
	}

	relationshipsRegistry[AutomationIdentity] = &elemental.Relationship{
		Create: map[string]*elemental.RelationshipInfo{
			"root": {},
		},
		Update: map[string]*elemental.RelationshipInfo{
			"root": {},
		},
		Patch: map[string]*elemental.RelationshipInfo{
			"root": {},
		},
		Delete: map[string]*elemental.RelationshipInfo{
			"root": {
				Parameters: []elemental.ParameterDefinition{
					{
						Name:     "q",
						Type:     "string",
						Multiple: true,
					},
				},
			},
		},
		Retrieve: map[string]*elemental.RelationshipInfo{
			"root": {},
		},
		RetrieveMany: map[string]*elemental.RelationshipInfo{
			"root": {
				Parameters: []elemental.ParameterDefinition{
					{
						Name:     "q",
						Type:     "string",
						Multiple: true,
					},
				},
			},
		},
		Info: map[string]*elemental.RelationshipInfo{
			"root": {
				Parameters: []elemental.ParameterDefinition{
					{
						Name:     "q",
						Type:     "string",
						Multiple: true,
					},
				},
			},
		},
	}

	relationshipsRegistry[AutomationTemplateIdentity] = &elemental.Relationship{
		Retrieve: map[string]*elemental.RelationshipInfo{
			"root": {},
		},
		RetrieveMany: map[string]*elemental.RelationshipInfo{
			"root": {},
		},
		Info: map[string]*elemental.RelationshipInfo{
			"root": {},
		},
	}

	relationshipsRegistry[CNSConfigIdentity] = &elemental.Relationship{
		Create: map[string]*elemental.RelationshipInfo{
			"root": {},
		},
		Update: map[string]*elemental.RelationshipInfo{
			"root": {},
		},
		Patch: map[string]*elemental.RelationshipInfo{
			"root": {},
		},
		Delete: map[string]*elemental.RelationshipInfo{
			"root": {},
		},
		Retrieve: map[string]*elemental.RelationshipInfo{
			"root": {},
		},
		RetrieveMany: map[string]*elemental.RelationshipInfo{
			"root": {
				Parameters: []elemental.ParameterDefinition{
					{
						Name:     "q",
						Type:     "string",
						Multiple: true,
					},
				},
			},
		},
		Info: map[string]*elemental.RelationshipInfo{
			"root": {
				Parameters: []elemental.ParameterDefinition{
					{
						Name:     "q",
						Type:     "string",
						Multiple: true,
					},
				},
			},
		},
	}

	relationshipsRegistry[CNSSearchIdentity] = &elemental.Relationship{
		Create: map[string]*elemental.RelationshipInfo{
			"root": {},
		},
	}

	relationshipsRegistry[CNSSuggestionIdentity] = &elemental.Relationship{
		Create: map[string]*elemental.RelationshipInfo{
			"root": {},
		},
	}

	relationshipsRegistry[CachedFlowReportIdentity] = &elemental.Relationship{
		Create: map[string]*elemental.RelationshipInfo{
			"root": {},
		},
	}

	relationshipsRegistry[CategoryIdentity] = &elemental.Relationship{}

	relationshipsRegistry[ClaimsIdentity] = &elemental.Relationship{
		Create: map[string]*elemental.RelationshipInfo{
			"root": {},
		},
		Retrieve: map[string]*elemental.RelationshipInfo{
			"root": {},
		},
		RetrieveMany: map[string]*elemental.RelationshipInfo{
			"root": {
				Parameters: []elemental.ParameterDefinition{
					{
						Name:     "q",
						Type:     "string",
						Multiple: true,
					},
				},
			},
		},
		Info: map[string]*elemental.RelationshipInfo{
			"root": {
				Parameters: []elemental.ParameterDefinition{
					{
						Name:     "q",
						Type:     "string",
						Multiple: true,
					},
				},
			},
		},
	}

	relationshipsRegistry[ClauseMatchIdentity] = &elemental.Relationship{
		Create: map[string]*elemental.RelationshipInfo{
			"root": {},
		},
	}

	relationshipsRegistry[CloudAccountCleanerIdentity] = &elemental.Relationship{
		Create: map[string]*elemental.RelationshipInfo{
			"root": {},
		},
	}

	relationshipsRegistry[CloudAlertIdentity] = &elemental.Relationship{
		Update: map[string]*elemental.RelationshipInfo{
			"root": {},
		},
		Patch: map[string]*elemental.RelationshipInfo{
			"root": {},
		},
		Delete: map[string]*elemental.RelationshipInfo{
			"root": {},
		},
		Retrieve: map[string]*elemental.RelationshipInfo{
			"root": {
				Parameters: []elemental.ParameterDefinition{
					{
						Name:     "q",
						Type:     "string",
						Multiple: true,
					},
				},
			},
		},
	}

	relationshipsRegistry[CloudEndpointIdentity] = &elemental.Relationship{
		Create: map[string]*elemental.RelationshipInfo{
			"root": {},
		},
		Update: map[string]*elemental.RelationshipInfo{
			"root": {},
		},
		Patch: map[string]*elemental.RelationshipInfo{
			"root": {},
		},
		Delete: map[string]*elemental.RelationshipInfo{
			"root": {},
		},
		Retrieve: map[string]*elemental.RelationshipInfo{
			"root": {},
		},
		RetrieveMany: map[string]*elemental.RelationshipInfo{
			"root": {
				Parameters: []elemental.ParameterDefinition{
					{
						Name:     "q",
						Type:     "string",
						Multiple: true,
					},
				},
			},
		},
		Info: map[string]*elemental.RelationshipInfo{
			"root": {
				Parameters: []elemental.ParameterDefinition{
					{
						Name:     "q",
						Type:     "string",
						Multiple: true,
					},
				},
			},
		},
	}

	relationshipsRegistry[CloudGraphIdentity] = &elemental.Relationship{
		Create: map[string]*elemental.RelationshipInfo{
			"root": {},
		},
		RetrieveMany: map[string]*elemental.RelationshipInfo{
			"cloudnetworkquery": {},
		},
		Info: map[string]*elemental.RelationshipInfo{
			"cloudnetworkquery": {},
		},
	}

	relationshipsRegistry[CloudManagedNetworkIdentity] = &elemental.Relationship{
		Create: map[string]*elemental.RelationshipInfo{
			"root": {},
		},
		Update: map[string]*elemental.RelationshipInfo{
			"root": {},
		},
		Patch: map[string]*elemental.RelationshipInfo{
			"root": {},
		},
		Delete: map[string]*elemental.RelationshipInfo{
			"root": {},
		},
		Retrieve: map[string]*elemental.RelationshipInfo{
			"root": {
				Parameters: []elemental.ParameterDefinition{
					{
						Name: "archived",
						Type: "boolean",
					},
					{
						Name: "propagated",
						Type: "boolean",
					},
				},
			},
		},
		RetrieveMany: map[string]*elemental.RelationshipInfo{
			"root": {
				Parameters: []elemental.ParameterDefinition{
					{
						Name:     "q",
						Type:     "string",
						Multiple: true,
					},
				},
			},
		},
		Info: map[string]*elemental.RelationshipInfo{
			"root": {
				Parameters: []elemental.ParameterDefinition{
					{
						Name:     "q",
						Type:     "string",
						Multiple: true,
					},
				},
			},
		},
	}

	relationshipsRegistry[CloudNetworkInterfaceIdentity] = &elemental.Relationship{
		Create: map[string]*elemental.RelationshipInfo{
			"root": {},
		},
		Update: map[string]*elemental.RelationshipInfo{
			"root": {},
		},
		Patch: map[string]*elemental.RelationshipInfo{
			"root": {},
		},
		Delete: map[string]*elemental.RelationshipInfo{
			"root": {},
		},
		Retrieve: map[string]*elemental.RelationshipInfo{
			"root": {
				Parameters: []elemental.ParameterDefinition{
					{
						Name:     "q",
						Type:     "string",
						Multiple: true,
					},
				},
			},
		},
		RetrieveMany: map[string]*elemental.RelationshipInfo{
			"root": {
				Parameters: []elemental.ParameterDefinition{
					{
						Name:     "q",
						Type:     "string",
						Multiple: true,
					},
				},
			},
		},
		Info: map[string]*elemental.RelationshipInfo{
			"root": {
				Parameters: []elemental.ParameterDefinition{
					{
						Name:     "q",
						Type:     "string",
						Multiple: true,
					},
				},
			},
		},
	}

	relationshipsRegistry[CloudNetworkQueryIdentity] = &elemental.Relationship{
		Create: map[string]*elemental.RelationshipInfo{
			"root": {},
		},
		Update: map[string]*elemental.RelationshipInfo{
			"root": {},
		},
		Patch: map[string]*elemental.RelationshipInfo{
			"root": {},
		},
		Delete: map[string]*elemental.RelationshipInfo{
			"root": {},
		},
		Retrieve: map[string]*elemental.RelationshipInfo{
			"root": {
				Parameters: []elemental.ParameterDefinition{
					{
						Name:     "q",
						Type:     "string",
						Multiple: true,
					},
				},
			},
		},
		RetrieveMany: map[string]*elemental.RelationshipInfo{
			"root": {
				Parameters: []elemental.ParameterDefinition{
					{
						Name:     "q",
						Type:     "string",
						Multiple: true,
					},
				},
			},
		},
		Info: map[string]*elemental.RelationshipInfo{
			"root": {
				Parameters: []elemental.ParameterDefinition{
					{
						Name:     "q",
						Type:     "string",
						Multiple: true,
					},
				},
			},
		},
	}

	relationshipsRegistry[CloudNetworkRuleSetIdentity] = &elemental.Relationship{
		Create: map[string]*elemental.RelationshipInfo{
			"root": {},
		},
		Update: map[string]*elemental.RelationshipInfo{
			"root": {},
		},
		Patch: map[string]*elemental.RelationshipInfo{
			"root": {},
		},
		Delete: map[string]*elemental.RelationshipInfo{
			"root": {},
		},
		Retrieve: map[string]*elemental.RelationshipInfo{
			"root": {},
		},
		RetrieveMany: map[string]*elemental.RelationshipInfo{
			"root": {
				Parameters: []elemental.ParameterDefinition{
					{
						Name:     "q",
						Type:     "string",
						Multiple: true,
					},
					{
						Name: "propagated",
						Type: "boolean",
					},
				},
			},
		},
		Info: map[string]*elemental.RelationshipInfo{
			"root": {
				Parameters: []elemental.ParameterDefinition{
					{
						Name:     "q",
						Type:     "string",
						Multiple: true,
					},
					{
						Name: "propagated",
						Type: "boolean",
					},
				},
			},
		},
	}

	relationshipsRegistry[CloudNodeIdentity] = &elemental.Relationship{
		Delete: map[string]*elemental.RelationshipInfo{
			"root": {},
		},
		Retrieve: map[string]*elemental.RelationshipInfo{
			"root": {},
		},
		RetrieveMany: map[string]*elemental.RelationshipInfo{
			"root": {
				Parameters: []elemental.ParameterDefinition{
					{
						Name:     "q",
						Type:     "string",
						Multiple: true,
					},
				},
			},
		},
		Info: map[string]*elemental.RelationshipInfo{
			"root": {
				Parameters: []elemental.ParameterDefinition{
					{
						Name:     "q",
						Type:     "string",
						Multiple: true,
					},
				},
			},
		},
	}

	relationshipsRegistry[CloudPolicyIdentity] = &elemental.Relationship{
		Create: map[string]*elemental.RelationshipInfo{
			"cloudnetworkquery": {},
		},
		Update: map[string]*elemental.RelationshipInfo{
			"root": {},
		},
		Patch: map[string]*elemental.RelationshipInfo{
			"root": {},
		},
		Delete: map[string]*elemental.RelationshipInfo{
			"root": {},
		},
		Retrieve: map[string]*elemental.RelationshipInfo{
			"root": {
				Parameters: []elemental.ParameterDefinition{
					{
						Name:     "q",
						Type:     "string",
						Multiple: true,
					},
				},
			},
		},
		RetrieveMany: map[string]*elemental.RelationshipInfo{
			"cloudnetworkquery": {},
		},
		Info: map[string]*elemental.RelationshipInfo{
			"cloudnetworkquery": {},
		},
	}

	relationshipsRegistry[CloudRouteTableIdentity] = &elemental.Relationship{
		Create: map[string]*elemental.RelationshipInfo{
			"root": {},
		},
		Update: map[string]*elemental.RelationshipInfo{
			"root": {},
		},
		Patch: map[string]*elemental.RelationshipInfo{
			"root": {},
		},
		Delete: map[string]*elemental.RelationshipInfo{
			"root": {
				Parameters: []elemental.ParameterDefinition{
					{
						Name:     "q",
						Type:     "string",
						Multiple: true,
					},
				},
			},
		},
		Retrieve: map[string]*elemental.RelationshipInfo{
			"root": {},
		},
		RetrieveMany: map[string]*elemental.RelationshipInfo{
			"root": {
				Parameters: []elemental.ParameterDefinition{
					{
						Name:     "q",
						Type:     "string",
						Multiple: true,
					},
				},
			},
		},
		Info: map[string]*elemental.RelationshipInfo{
			"root": {
				Parameters: []elemental.ParameterDefinition{
					{
						Name:     "q",
						Type:     "string",
						Multiple: true,
					},
				},
			},
		},
	}

	relationshipsRegistry[CloudSnapshotAccountIdentity] = &elemental.Relationship{
		Create: map[string]*elemental.RelationshipInfo{
			"root": {},
		},
	}

	relationshipsRegistry[CloudSubnetIdentity] = &elemental.Relationship{
		Create: map[string]*elemental.RelationshipInfo{
			"root": {},
		},
		Update: map[string]*elemental.RelationshipInfo{
			"root": {},
		},
		Patch: map[string]*elemental.RelationshipInfo{
			"root": {},
		},
		Delete: map[string]*elemental.RelationshipInfo{
			"root": {},
		},
		Retrieve: map[string]*elemental.RelationshipInfo{
			"root": {
				Parameters: []elemental.ParameterDefinition{
					{
						Name:     "q",
						Type:     "string",
						Multiple: true,
					},
				},
			},
		},
		RetrieveMany: map[string]*elemental.RelationshipInfo{
			"root": {
				Parameters: []elemental.ParameterDefinition{
					{
						Name:     "q",
						Type:     "string",
						Multiple: true,
					},
				},
			},
		},
		Info: map[string]*elemental.RelationshipInfo{
			"root": {
				Parameters: []elemental.ParameterDefinition{
					{
						Name:     "q",
						Type:     "string",
						Multiple: true,
					},
				},
			},
		},
	}

	relationshipsRegistry[CloudVPCIdentity] = &elemental.Relationship{
		Create: map[string]*elemental.RelationshipInfo{
			"root": {},
		},
		Update: map[string]*elemental.RelationshipInfo{
			"root": {},
		},
		Patch: map[string]*elemental.RelationshipInfo{
			"root": {},
		},
		Delete: map[string]*elemental.RelationshipInfo{
			"root": {},
		},
		Retrieve: map[string]*elemental.RelationshipInfo{
			"root": {},
		},
		RetrieveMany: map[string]*elemental.RelationshipInfo{
			"root": {
				Parameters: []elemental.ParameterDefinition{
					{
						Name:     "q",
						Type:     "string",
						Multiple: true,
					},
				},
			},
		},
		Info: map[string]*elemental.RelationshipInfo{
			"root": {
				Parameters: []elemental.ParameterDefinition{
					{
						Name:     "q",
						Type:     "string",
						Multiple: true,
					},
				},
			},
		},
	}

	relationshipsRegistry[ConnectionExceptionReportIdentity] = &elemental.Relationship{
		Create: map[string]*elemental.RelationshipInfo{
			"root": {},
		},
	}

	relationshipsRegistry[CounterReportIdentity] = &elemental.Relationship{
		Create: map[string]*elemental.RelationshipInfo{
			"root": {},
		},
	}

	relationshipsRegistry[DNSLookupReportIdentity] = &elemental.Relationship{
		Create: map[string]*elemental.RelationshipInfo{
			"root": {},
		},
	}

	relationshipsRegistry[DataPathCertificateIdentity] = &elemental.Relationship{
		Create: map[string]*elemental.RelationshipInfo{
			"root": {},
		},
	}

	relationshipsRegistry[DebugBundleIdentity] = &elemental.Relationship{
		Create: map[string]*elemental.RelationshipInfo{
			"enforcer": {},
		},
		RetrieveMany: map[string]*elemental.RelationshipInfo{
			"enforcer": {},
		},
		Info: map[string]*elemental.RelationshipInfo{
			"enforcer": {},
		},
	}

	relationshipsRegistry[DefaultEnforcerVersionIdentity] = &elemental.Relationship{
		Create: map[string]*elemental.RelationshipInfo{
			"root": {},
		},
		RetrieveMany: map[string]*elemental.RelationshipInfo{
			"root": {},
		},
		Info: map[string]*elemental.RelationshipInfo{
			"root": {},
		},
	}

	relationshipsRegistry[DependencyMapIdentity] = &elemental.Relationship{
		RetrieveMany: map[string]*elemental.RelationshipInfo{
			"root": {
				RequiredParameters: elemental.NewParametersRequirement(
					[][][]string{
						{
							{
								"endRelative",
							},
							{
								"startRelative",
							},
							{
								"startRelative",
								"endRelative",
							},
							{
								"startRelative",
								"endAbsolute",
							},
							{
								"startAbsolute",
								"endRelative",
							},
							{
								"startAbsolute",
								"endAbsolute",
							},
						},
					},
				),
				Parameters: []elemental.ParameterDefinition{
					{
						Name:     "tag",
						Type:     "string",
						Multiple: true,
					},
					{
						Name: "view",
						Type: "string",
					},
					{
						Name: "viewSuggestions",
						Type: "boolean",
					},
					{
						Name: "endAbsolute",
						Type: "time",
					},
					{
						Name: "endRelative",
						Type: "duration",
					},
					{
						Name: "startAbsolute",
						Type: "time",
					},
					{
						Name: "startRelative",
						Type: "duration",
					},
					{
						Name: "flowOffset",
						Type: "duration",
					},
					{
						Name:     "q",
						Type:     "string",
						Multiple: true,
					},
				},
			},
		},
		Info: map[string]*elemental.RelationshipInfo{
			"root": {
				RequiredParameters: elemental.NewParametersRequirement(
					[][][]string{
						{
							{
								"endRelative",
							},
							{
								"startRelative",
							},
							{
								"startRelative",
								"endRelative",
							},
							{
								"startRelative",
								"endAbsolute",
							},
							{
								"startAbsolute",
								"endRelative",
							},
							{
								"startAbsolute",
								"endAbsolute",
							},
						},
					},
				),
				Parameters: []elemental.ParameterDefinition{
					{
						Name:     "tag",
						Type:     "string",
						Multiple: true,
					},
					{
						Name: "view",
						Type: "string",
					},
					{
						Name: "viewSuggestions",
						Type: "boolean",
					},
					{
						Name: "endAbsolute",
						Type: "time",
					},
					{
						Name: "endRelative",
						Type: "duration",
					},
					{
						Name: "startAbsolute",
						Type: "time",
					},
					{
						Name: "startRelative",
						Type: "duration",
					},
					{
						Name: "flowOffset",
						Type: "duration",
					},
					{
						Name:     "q",
						Type:     "string",
						Multiple: true,
					},
				},
			},
		},
	}

	relationshipsRegistry[DiscoveryModeIdentity] = &elemental.Relationship{
		Create: map[string]*elemental.RelationshipInfo{
			"root": {
				Deprecated: true,
			},
		},
		Delete: map[string]*elemental.RelationshipInfo{
			"root": {
				Deprecated: true,
			},
		},
		Retrieve: map[string]*elemental.RelationshipInfo{
			"root": {
				Deprecated: true,
			},
		},
		RetrieveMany: map[string]*elemental.RelationshipInfo{
			"root": {
				Deprecated: true,
			},
		},
		Info: map[string]*elemental.RelationshipInfo{
			"root": {
				Deprecated: true,
			},
		},
	}

	relationshipsRegistry[EmailIdentity] = &elemental.Relationship{
		Create: map[string]*elemental.RelationshipInfo{
			"root": {},
		},
	}

	relationshipsRegistry[EnforcerIdentity] = &elemental.Relationship{
		Create: map[string]*elemental.RelationshipInfo{
			"root": {},
		},
		Update: map[string]*elemental.RelationshipInfo{
			"root": {},
		},
		Patch: map[string]*elemental.RelationshipInfo{
			"root": {},
		},
		Delete: map[string]*elemental.RelationshipInfo{
			"root": {
				Parameters: []elemental.ParameterDefinition{
					{
						Name:     "q",
						Type:     "string",
						Multiple: true,
					},
				},
			},
		},
		Retrieve: map[string]*elemental.RelationshipInfo{
			"root": {},
		},
		RetrieveMany: map[string]*elemental.RelationshipInfo{
			"auditprofilemappingpolicy":    {},
			"enforcerprofilemappingpolicy": {},
			"hostservicemappingpolicy":     {},
			"root": {
				Parameters: []elemental.ParameterDefinition{
					{
						Name:     "q",
						Type:     "string",
						Multiple: true,
					},
				},
			},
		},
		Info: map[string]*elemental.RelationshipInfo{
			"auditprofilemappingpolicy":    {},
			"enforcerprofilemappingpolicy": {},
			"hostservicemappingpolicy":     {},
			"root": {
				Parameters: []elemental.ParameterDefinition{
					{
						Name:     "q",
						Type:     "string",
						Multiple: true,
					},
				},
			},
		},
	}

	relationshipsRegistry[EnforcerLogIdentity] = &elemental.Relationship{
		Create: map[string]*elemental.RelationshipInfo{
			"root": {},
		},
		Retrieve: map[string]*elemental.RelationshipInfo{
			"root": {},
		},
		RetrieveMany: map[string]*elemental.RelationshipInfo{
			"root": {
				Parameters: []elemental.ParameterDefinition{
					{
						Name:     "q",
						Type:     "string",
						Multiple: true,
					},
				},
			},
		},
		Info: map[string]*elemental.RelationshipInfo{
			"root": {
				Parameters: []elemental.ParameterDefinition{
					{
						Name:     "q",
						Type:     "string",
						Multiple: true,
					},
				},
			},
		},
	}

	relationshipsRegistry[EnforcerProfileIdentity] = &elemental.Relationship{
		Create: map[string]*elemental.RelationshipInfo{
			"root": {},
		},
		Update: map[string]*elemental.RelationshipInfo{
			"root": {},
		},
		Patch: map[string]*elemental.RelationshipInfo{
			"root": {},
		},
		Delete: map[string]*elemental.RelationshipInfo{
			"root": {
				Parameters: []elemental.ParameterDefinition{
					{
						Name:     "q",
						Type:     "string",
						Multiple: true,
					},
				},
			},
		},
		Retrieve: map[string]*elemental.RelationshipInfo{
			"root": {
				Parameters: []elemental.ParameterDefinition{
					{
						Name: "propagated",
						Type: "boolean",
					},
				},
			},
		},
		RetrieveMany: map[string]*elemental.RelationshipInfo{
			"enforcer":                     {},
			"enforcerprofilemappingpolicy": {},
			"root": {
				Parameters: []elemental.ParameterDefinition{
					{
						Name:     "q",
						Type:     "string",
						Multiple: true,
					},
					{
						Name: "propagated",
						Type: "boolean",
					},
				},
			},
		},
		Info: map[string]*elemental.RelationshipInfo{
			"enforcer":                     {},
			"enforcerprofilemappingpolicy": {},
			"root": {
				Parameters: []elemental.ParameterDefinition{
					{
						Name:     "q",
						Type:     "string",
						Multiple: true,
					},
					{
						Name: "propagated",
						Type: "boolean",
					},
				},
			},
		},
	}

	relationshipsRegistry[EnforcerProfileMappingPolicyIdentity] = &elemental.Relationship{
		Create: map[string]*elemental.RelationshipInfo{
			"root": {},
		},
		Update: map[string]*elemental.RelationshipInfo{
			"root": {},
		},
		Patch: map[string]*elemental.RelationshipInfo{
			"root": {},
		},
		Delete: map[string]*elemental.RelationshipInfo{
			"root": {
				Parameters: []elemental.ParameterDefinition{
					{
						Name:     "q",
						Type:     "string",
						Multiple: true,
					},
				},
			},
		},
		Retrieve: map[string]*elemental.RelationshipInfo{
			"root": {},
		},
		RetrieveMany: map[string]*elemental.RelationshipInfo{
			"root": {
				Parameters: []elemental.ParameterDefinition{
					{
						Name:     "q",
						Type:     "string",
						Multiple: true,
					},
					{
						Name: "propagated",
						Type: "boolean",
					},
				},
			},
		},
		Info: map[string]*elemental.RelationshipInfo{
			"root": {
				Parameters: []elemental.ParameterDefinition{
					{
						Name:     "q",
						Type:     "string",
						Multiple: true,
					},
					{
						Name: "propagated",
						Type: "boolean",
					},
				},
			},
		},
	}

	relationshipsRegistry[EnforcerRefreshIdentity] = &elemental.Relationship{
		Create: map[string]*elemental.RelationshipInfo{
			"enforcer": {},
			"root":     {},
		},
	}

	relationshipsRegistry[EnforcerReportIdentity] = &elemental.Relationship{
		Create: map[string]*elemental.RelationshipInfo{
			"root": {},
		},
	}

	relationshipsRegistry[EnforcerTraceReportIdentity] = &elemental.Relationship{
		Create: map[string]*elemental.RelationshipInfo{
			"root": {},
		},
	}

	relationshipsRegistry[EventLogIdentity] = &elemental.Relationship{
		Create: map[string]*elemental.RelationshipInfo{
			"root": {},
		},
	}

	relationshipsRegistry[ExportIdentity] = &elemental.Relationship{
		Create: map[string]*elemental.RelationshipInfo{
			"root": {
				Parameters: []elemental.ParameterDefinition{
					{
						Name:     "q",
						Type:     "string",
						Multiple: true,
					},
				},
			},
		},
	}

	relationshipsRegistry[ExternalNetworkIdentity] = &elemental.Relationship{
		Create: map[string]*elemental.RelationshipInfo{
			"root": {},
		},
		Update: map[string]*elemental.RelationshipInfo{
			"root": {},
		},
		Patch: map[string]*elemental.RelationshipInfo{
			"root": {},
		},
		Delete: map[string]*elemental.RelationshipInfo{
			"root": {
				Parameters: []elemental.ParameterDefinition{
					{
						Name:     "q",
						Type:     "string",
						Multiple: true,
					},
				},
			},
		},
		Retrieve: map[string]*elemental.RelationshipInfo{
			"root": {
				Parameters: []elemental.ParameterDefinition{
					{
						Name: "archived",
						Type: "boolean",
					},
					{
						Name: "propagated",
						Type: "boolean",
					},
				},
			},
		},
		RetrieveMany: map[string]*elemental.RelationshipInfo{
			"infrastructurepolicy": {
				Parameters: []elemental.ParameterDefinition{
					{
						Name:         "mode",
						Type:         "enum",
						DefaultValue: "object",
						AllowedChoices: []string{
							"subject",
							"object",
						},
					},
				},
			},
			"networkaccesspolicy": {
				Parameters: []elemental.ParameterDefinition{
					{
						Name:         "mode",
						Type:         "enum",
						DefaultValue: "object",
						AllowedChoices: []string{
							"subject",
							"object",
						},
					},
				},
			},
			"networkrulesetpolicy": {
				Parameters: []elemental.ParameterDefinition{
					{
						Name:         "mode",
						Type:         "enum",
						DefaultValue: "object",
						AllowedChoices: []string{
							"subject",
							"object",
						},
					},
				},
			},
			"root": {
				Parameters: []elemental.ParameterDefinition{
					{
						Name:     "q",
						Type:     "string",
						Multiple: true,
					},
					{
						Name: "archived",
						Type: "boolean",
					},
					{
						Name: "propagated",
						Type: "boolean",
					},
				},
			},
		},
		Info: map[string]*elemental.RelationshipInfo{
			"infrastructurepolicy": {
				Parameters: []elemental.ParameterDefinition{
					{
						Name:         "mode",
						Type:         "enum",
						DefaultValue: "object",
						AllowedChoices: []string{
							"subject",
							"object",
						},
					},
				},
			},
			"networkaccesspolicy": {
				Parameters: []elemental.ParameterDefinition{
					{
						Name:         "mode",
						Type:         "enum",
						DefaultValue: "object",
						AllowedChoices: []string{
							"subject",
							"object",
						},
					},
				},
			},
			"networkrulesetpolicy": {
				Parameters: []elemental.ParameterDefinition{
					{
						Name:         "mode",
						Type:         "enum",
						DefaultValue: "object",
						AllowedChoices: []string{
							"subject",
							"object",
						},
					},
				},
			},
			"root": {
				Parameters: []elemental.ParameterDefinition{
					{
						Name:     "q",
						Type:     "string",
						Multiple: true,
					},
					{
						Name: "archived",
						Type: "boolean",
					},
					{
						Name: "propagated",
						Type: "boolean",
					},
				},
			},
		},
	}

	relationshipsRegistry[FileAccessPolicyIdentity] = &elemental.Relationship{
		Create: map[string]*elemental.RelationshipInfo{
			"root": {},
		},
		Update: map[string]*elemental.RelationshipInfo{
			"root": {},
		},
		Patch: map[string]*elemental.RelationshipInfo{
			"root": {},
		},
		Delete: map[string]*elemental.RelationshipInfo{
			"root": {
				Parameters: []elemental.ParameterDefinition{
					{
						Name:     "q",
						Type:     "string",
						Multiple: true,
					},
				},
			},
		},
		Retrieve: map[string]*elemental.RelationshipInfo{
			"root": {},
		},
		RetrieveMany: map[string]*elemental.RelationshipInfo{
			"root": {
				Parameters: []elemental.ParameterDefinition{
					{
						Name:     "q",
						Type:     "string",
						Multiple: true,
					},
					{
						Name: "propagated",
						Type: "boolean",
					},
				},
			},
		},
		Info: map[string]*elemental.RelationshipInfo{
			"root": {
				Parameters: []elemental.ParameterDefinition{
					{
						Name:     "q",
						Type:     "string",
						Multiple: true,
					},
					{
						Name: "propagated",
						Type: "boolean",
					},
				},
			},
		},
	}

	relationshipsRegistry[FileAccessReportIdentity] = &elemental.Relationship{
		Create: map[string]*elemental.RelationshipInfo{
			"root": {},
		},
	}

	relationshipsRegistry[FilePathIdentity] = &elemental.Relationship{
		Create: map[string]*elemental.RelationshipInfo{
			"root": {},
		},
		Update: map[string]*elemental.RelationshipInfo{
			"root": {},
		},
		Patch: map[string]*elemental.RelationshipInfo{
			"root": {},
		},
		Delete: map[string]*elemental.RelationshipInfo{
			"root": {
				Parameters: []elemental.ParameterDefinition{
					{
						Name:     "q",
						Type:     "string",
						Multiple: true,
					},
				},
			},
		},
		Retrieve: map[string]*elemental.RelationshipInfo{
			"root": {
				Parameters: []elemental.ParameterDefinition{
					{
						Name: "archived",
						Type: "boolean",
					},
					{
						Name: "propagated",
						Type: "boolean",
					},
				},
			},
		},
		RetrieveMany: map[string]*elemental.RelationshipInfo{
			"fileaccesspolicy": {},
			"root": {
				Parameters: []elemental.ParameterDefinition{
					{
						Name:     "q",
						Type:     "string",
						Multiple: true,
					},
					{
						Name: "archived",
						Type: "boolean",
					},
					{
						Name: "propagated",
						Type: "boolean",
					},
				},
			},
		},
		Info: map[string]*elemental.RelationshipInfo{
			"fileaccesspolicy": {},
			"root": {
				Parameters: []elemental.ParameterDefinition{
					{
						Name:     "q",
						Type:     "string",
						Multiple: true,
					},
					{
						Name: "archived",
						Type: "boolean",
					},
					{
						Name: "propagated",
						Type: "boolean",
					},
				},
			},
		},
	}

	relationshipsRegistry[FlowReportIdentity] = &elemental.Relationship{
		Create: map[string]*elemental.RelationshipInfo{
			"root": {},
		},
	}

	relationshipsRegistry[GraphEdgeIdentity] = &elemental.Relationship{
		RetrieveMany: map[string]*elemental.RelationshipInfo{
			"root": {
				RequiredParameters: elemental.NewParametersRequirement(
					[][][]string{
						{
							{
								"endRelative",
							},
							{
								"startRelative",
							},
							{
								"startRelative",
								"endRelative",
							},
							{
								"startRelative",
								"endAbsolute",
							},
							{
								"startAbsolute",
								"endRelative",
							},
							{
								"startAbsolute",
								"endAbsolute",
							},
						},
					},
				),
				Parameters: []elemental.ParameterDefinition{
					{
						Name: "endAbsolute",
						Type: "time",
					},
					{
						Name: "endRelative",
						Type: "duration",
					},
					{
						Name: "startAbsolute",
						Type: "time",
					},
					{
						Name: "startRelative",
						Type: "duration",
					},
					{
						Name:     "q",
						Type:     "string",
						Multiple: true,
					},
				},
			},
		},
		Info: map[string]*elemental.RelationshipInfo{
			"root": {
				RequiredParameters: elemental.NewParametersRequirement(
					[][][]string{
						{
							{
								"endRelative",
							},
							{
								"startRelative",
							},
							{
								"startRelative",
								"endRelative",
							},
							{
								"startRelative",
								"endAbsolute",
							},
							{
								"startAbsolute",
								"endRelative",
							},
							{
								"startAbsolute",
								"endAbsolute",
							},
						},
					},
				),
				Parameters: []elemental.ParameterDefinition{
					{
						Name: "endAbsolute",
						Type: "time",
					},
					{
						Name: "endRelative",
						Type: "duration",
					},
					{
						Name: "startAbsolute",
						Type: "time",
					},
					{
						Name: "startRelative",
						Type: "duration",
					},
					{
						Name:     "q",
						Type:     "string",
						Multiple: true,
					},
				},
			},
		},
	}

	relationshipsRegistry[GraphNodeIdentity] = &elemental.Relationship{
		RetrieveMany: map[string]*elemental.RelationshipInfo{
			"root": {
				RequiredParameters: elemental.NewParametersRequirement(
					[][][]string{
						{
							{
								"endRelative",
							},
							{
								"startRelative",
							},
							{
								"startRelative",
								"endRelative",
							},
							{
								"startRelative",
								"endAbsolute",
							},
							{
								"startAbsolute",
								"endRelative",
							},
							{
								"startAbsolute",
								"endAbsolute",
							},
						},
					},
				),
				Parameters: []elemental.ParameterDefinition{
					{
						Name: "endAbsolute",
						Type: "time",
					},
					{
						Name: "endRelative",
						Type: "duration",
					},
					{
						Name: "startAbsolute",
						Type: "time",
					},
					{
						Name: "startRelative",
						Type: "duration",
					},
					{
						Name: "archived",
						Type: "boolean",
					},
					{
						Name:     "q",
						Type:     "string",
						Multiple: true,
					},
				},
			},
		},
		Info: map[string]*elemental.RelationshipInfo{
			"root": {
				RequiredParameters: elemental.NewParametersRequirement(
					[][][]string{
						{
							{
								"endRelative",
							},
							{
								"startRelative",
							},
							{
								"startRelative",
								"endRelative",
							},
							{
								"startRelative",
								"endAbsolute",
							},
							{
								"startAbsolute",
								"endRelative",
							},
							{
								"startAbsolute",
								"endAbsolute",
							},
						},
					},
				),
				Parameters: []elemental.ParameterDefinition{
					{
						Name: "endAbsolute",
						Type: "time",
					},
					{
						Name: "endRelative",
						Type: "duration",
					},
					{
						Name: "startAbsolute",
						Type: "time",
					},
					{
						Name: "startRelative",
						Type: "duration",
					},
					{
						Name: "archived",
						Type: "boolean",
					},
					{
						Name:     "q",
						Type:     "string",
						Multiple: true,
					},
				},
			},
		},
	}

	relationshipsRegistry[HTTPResourceSpecIdentity] = &elemental.Relationship{
		Create: map[string]*elemental.RelationshipInfo{
			"root": {},
		},
		Update: map[string]*elemental.RelationshipInfo{
			"root": {},
		},
		Patch: map[string]*elemental.RelationshipInfo{
			"root": {},
		},
		Delete: map[string]*elemental.RelationshipInfo{
			"root": {
				Parameters: []elemental.ParameterDefinition{
					{
						Name:     "q",
						Type:     "string",
						Multiple: true,
					},
				},
			},
		},
		Retrieve: map[string]*elemental.RelationshipInfo{
			"root": {
				Parameters: []elemental.ParameterDefinition{
					{
						Name: "archived",
						Type: "boolean",
					},
				},
			},
		},
		RetrieveMany: map[string]*elemental.RelationshipInfo{
			"root": {
				Parameters: []elemental.ParameterDefinition{
					{
						Name:     "q",
						Type:     "string",
						Multiple: true,
					},
					{
						Name: "propagated",
						Type: "boolean",
					},
					{
						Name: "archived",
						Type: "boolean",
					},
				},
			},
			"service": {},
		},
		Info: map[string]*elemental.RelationshipInfo{
			"root": {
				Parameters: []elemental.ParameterDefinition{
					{
						Name:     "q",
						Type:     "string",
						Multiple: true,
					},
					{
						Name: "propagated",
						Type: "boolean",
					},
					{
						Name: "archived",
						Type: "boolean",
					},
				},
			},
			"service": {},
		},
	}

	relationshipsRegistry[HealthCheckIdentity] = &elemental.Relationship{
		RetrieveMany: map[string]*elemental.RelationshipInfo{
			"root": {
				Parameters: []elemental.ParameterDefinition{
					{
						Name: "quiet",
						Type: "boolean",
					},
				},
			},
		},
		Info: map[string]*elemental.RelationshipInfo{
			"root": {
				Parameters: []elemental.ParameterDefinition{
					{
						Name: "quiet",
						Type: "boolean",
					},
				},
			},
		},
	}

	relationshipsRegistry[HitIdentity] = &elemental.Relationship{
		Create: map[string]*elemental.RelationshipInfo{
			"root": {
				Parameters: []elemental.ParameterDefinition{
					{
						Name: "reset",
						Type: "boolean",
					},
				},
			},
		},
		RetrieveMany: map[string]*elemental.RelationshipInfo{
			"root": {
				RequiredParameters: elemental.NewParametersRequirement(
					[][][]string{
						{
							{
								"name",
								"targetID",
								"targetIdentity",
							},
							{
								"targetID",
								"targetIdentity",
							},
						},
					},
				),
				Parameters: []elemental.ParameterDefinition{
					{
						Name:         "name",
						Type:         "string",
						DefaultValue: "counter",
					},
					{
						Name: "targetID",
						Type: "string",
					},
					{
						Name: "targetIdentity",
						Type: "string",
					},
				},
			},
		},
		Info: map[string]*elemental.RelationshipInfo{
			"root": {
				RequiredParameters: elemental.NewParametersRequirement(
					[][][]string{
						{
							{
								"name",
								"targetID",
								"targetIdentity",
							},
							{
								"targetID",
								"targetIdentity",
							},
						},
					},
				),
				Parameters: []elemental.ParameterDefinition{
					{
						Name:         "name",
						Type:         "string",
						DefaultValue: "counter",
					},
					{
						Name: "targetID",
						Type: "string",
					},
					{
						Name: "targetIdentity",
						Type: "string",
					},
				},
			},
		},
	}

	relationshipsRegistry[HookPolicyIdentity] = &elemental.Relationship{
		Create: map[string]*elemental.RelationshipInfo{
			"root": {},
		},
		Update: map[string]*elemental.RelationshipInfo{
			"root": {},
		},
		Patch: map[string]*elemental.RelationshipInfo{
			"root": {},
		},
		Delete: map[string]*elemental.RelationshipInfo{
			"root": {
				Parameters: []elemental.ParameterDefinition{
					{
						Name:     "q",
						Type:     "string",
						Multiple: true,
					},
				},
			},
		},
		Retrieve: map[string]*elemental.RelationshipInfo{
			"root": {},
		},
		RetrieveMany: map[string]*elemental.RelationshipInfo{
			"root": {
				Parameters: []elemental.ParameterDefinition{
					{
						Name:     "q",
						Type:     "string",
						Multiple: true,
					},
					{
						Name: "propagated",
						Type: "boolean",
					},
				},
			},
		},
		Info: map[string]*elemental.RelationshipInfo{
			"root": {
				Parameters: []elemental.ParameterDefinition{
					{
						Name:     "q",
						Type:     "string",
						Multiple: true,
					},
					{
						Name: "propagated",
						Type: "boolean",
					},
				},
			},
		},
	}

	relationshipsRegistry[HostServiceIdentity] = &elemental.Relationship{
		Create: map[string]*elemental.RelationshipInfo{
			"root": {},
		},
		Update: map[string]*elemental.RelationshipInfo{
			"root": {},
		},
		Patch: map[string]*elemental.RelationshipInfo{
			"root": {},
		},
		Delete: map[string]*elemental.RelationshipInfo{
			"root": {
				Parameters: []elemental.ParameterDefinition{
					{
						Name:     "q",
						Type:     "string",
						Multiple: true,
					},
				},
			},
		},
		Retrieve: map[string]*elemental.RelationshipInfo{
			"root": {
				Parameters: []elemental.ParameterDefinition{
					{
						Name: "archived",
						Type: "boolean",
					},
					{
						Name: "propagated",
						Type: "boolean",
					},
				},
			},
		},
		RetrieveMany: map[string]*elemental.RelationshipInfo{
			"enforcer": {
				Parameters: []elemental.ParameterDefinition{
					{
						Name: "appliedServices",
						Type: "boolean",
					},
					{
						Name: "setServices",
						Type: "boolean",
					},
				},
			},
			"hostservicemappingpolicy": {},
			"root": {
				Parameters: []elemental.ParameterDefinition{
					{
						Name:     "q",
						Type:     "string",
						Multiple: true,
					},
					{
						Name: "propagated",
						Type: "boolean",
					},
					{
						Name: "propagated",
						Type: "boolean",
					},
				},
			},
		},
		Info: map[string]*elemental.RelationshipInfo{
			"enforcer": {
				Parameters: []elemental.ParameterDefinition{
					{
						Name: "appliedServices",
						Type: "boolean",
					},
					{
						Name: "setServices",
						Type: "boolean",
					},
				},
			},
			"hostservicemappingpolicy": {},
			"root": {
				Parameters: []elemental.ParameterDefinition{
					{
						Name:     "q",
						Type:     "string",
						Multiple: true,
					},
					{
						Name: "propagated",
						Type: "boolean",
					},
					{
						Name: "propagated",
						Type: "boolean",
					},
				},
			},
		},
	}

	relationshipsRegistry[HostServiceMappingPolicyIdentity] = &elemental.Relationship{
		Create: map[string]*elemental.RelationshipInfo{
			"root": {},
		},
		Update: map[string]*elemental.RelationshipInfo{
			"root": {},
		},
		Patch: map[string]*elemental.RelationshipInfo{
			"root": {},
		},
		Delete: map[string]*elemental.RelationshipInfo{
			"root": {
				Parameters: []elemental.ParameterDefinition{
					{
						Name:     "q",
						Type:     "string",
						Multiple: true,
					},
				},
			},
		},
		Retrieve: map[string]*elemental.RelationshipInfo{
			"root": {
				Parameters: []elemental.ParameterDefinition{
					{
						Name: "propagated",
						Type: "boolean",
					},
				},
			},
		},
		RetrieveMany: map[string]*elemental.RelationshipInfo{
			"root": {
				Parameters: []elemental.ParameterDefinition{
					{
						Name:     "q",
						Type:     "string",
						Multiple: true,
					},
					{
						Name: "propagated",
						Type: "boolean",
					},
				},
			},
		},
		Info: map[string]*elemental.RelationshipInfo{
			"root": {
				Parameters: []elemental.ParameterDefinition{
					{
						Name:     "q",
						Type:     "string",
						Multiple: true,
					},
					{
						Name: "propagated",
						Type: "boolean",
					},
				},
			},
		},
	}

	relationshipsRegistry[IPInfoIdentity] = &elemental.Relationship{
		RetrieveMany: map[string]*elemental.RelationshipInfo{
			"root": {
				RequiredParameters: elemental.NewParametersRequirement(
					[][][]string{
						{
							{
								"ip",
							},
						},
					},
				),
				Parameters: []elemental.ParameterDefinition{
					{
						Name:     "ip",
						Type:     "string",
						Multiple: true,
					},
				},
			},
		},
		Info: map[string]*elemental.RelationshipInfo{
			"root": {
				RequiredParameters: elemental.NewParametersRequirement(
					[][][]string{
						{
							{
								"ip",
							},
						},
					},
				),
				Parameters: []elemental.ParameterDefinition{
					{
						Name:     "ip",
						Type:     "string",
						Multiple: true,
					},
				},
			},
		},
	}

	relationshipsRegistry[ImportIdentity] = &elemental.Relationship{
		Create: map[string]*elemental.RelationshipInfo{
			"root": {},
		},
	}

	relationshipsRegistry[ImportReferenceIdentity] = &elemental.Relationship{
		Create: map[string]*elemental.RelationshipInfo{
			"recipe": {},
			"root":   {},
		},
		Delete: map[string]*elemental.RelationshipInfo{
			"root": {
				Parameters: []elemental.ParameterDefinition{
					{
						Name:     "q",
						Type:     "string",
						Multiple: true,
					},
				},
			},
		},
		Retrieve: map[string]*elemental.RelationshipInfo{
			"root": {},
		},
		RetrieveMany: map[string]*elemental.RelationshipInfo{
			"recipe": {},
			"root": {
				Parameters: []elemental.ParameterDefinition{
					{
						Name:     "q",
						Type:     "string",
						Multiple: true,
					},
				},
			},
		},
		Info: map[string]*elemental.RelationshipInfo{
			"recipe": {},
			"root": {
				Parameters: []elemental.ParameterDefinition{
					{
						Name:     "q",
						Type:     "string",
						Multiple: true,
					},
				},
			},
		},
	}

	relationshipsRegistry[ImportRequestIdentity] = &elemental.Relationship{
		Create: map[string]*elemental.RelationshipInfo{
			"root": {},
		},
		Update: map[string]*elemental.RelationshipInfo{
			"root": {},
		},
		Patch: map[string]*elemental.RelationshipInfo{
			"root": {},
		},
		Delete: map[string]*elemental.RelationshipInfo{
			"root": {},
		},
		Retrieve: map[string]*elemental.RelationshipInfo{
			"root": {},
		},
		RetrieveMany: map[string]*elemental.RelationshipInfo{
			"root": {
				Parameters: []elemental.ParameterDefinition{
					{
						Name:     "q",
						Type:     "string",
						Multiple: true,
					},
				},
			},
		},
		Info: map[string]*elemental.RelationshipInfo{
			"root": {
				Parameters: []elemental.ParameterDefinition{
					{
						Name:     "q",
						Type:     "string",
						Multiple: true,
					},
				},
			},
		},
	}

	relationshipsRegistry[InfrastructurePolicyIdentity] = &elemental.Relationship{
		Create: map[string]*elemental.RelationshipInfo{
			"root": {},
		},
		Update: map[string]*elemental.RelationshipInfo{
			"root": {},
		},
		Patch: map[string]*elemental.RelationshipInfo{
			"root": {},
		},
		Delete: map[string]*elemental.RelationshipInfo{
			"root": {
				Parameters: []elemental.ParameterDefinition{
					{
						Name:     "q",
						Type:     "string",
						Multiple: true,
					},
				},
			},
		},
		Retrieve: map[string]*elemental.RelationshipInfo{
			"root": {},
		},
		RetrieveMany: map[string]*elemental.RelationshipInfo{
			"root": {
				Parameters: []elemental.ParameterDefinition{
					{
						Name:     "q",
						Type:     "string",
						Multiple: true,
					},
				},
			},
		},
		Info: map[string]*elemental.RelationshipInfo{
			"root": {
				Parameters: []elemental.ParameterDefinition{
					{
						Name:     "q",
						Type:     "string",
						Multiple: true,
					},
				},
			},
		},
	}

	relationshipsRegistry[InstalledAppIdentity] = &elemental.Relationship{
		Create: map[string]*elemental.RelationshipInfo{
			"root": {},
		},
		Update: map[string]*elemental.RelationshipInfo{
			"root": {},
		},
		Patch: map[string]*elemental.RelationshipInfo{
			"root": {},
		},
		Delete: map[string]*elemental.RelationshipInfo{
			"root": {
				Parameters: []elemental.ParameterDefinition{
					{
						Name:     "q",
						Type:     "string",
						Multiple: true,
					},
				},
			},
		},
		Retrieve: map[string]*elemental.RelationshipInfo{
			"root": {},
		},
		RetrieveMany: map[string]*elemental.RelationshipInfo{
			"root": {
				Parameters: []elemental.ParameterDefinition{
					{
						Name:     "tag",
						Type:     "string",
						Multiple: true,
					},
					{
						Name:     "q",
						Type:     "string",
						Multiple: true,
					},
				},
			},
		},
		Info: map[string]*elemental.RelationshipInfo{
			"root": {
				Parameters: []elemental.ParameterDefinition{
					{
						Name:     "tag",
						Type:     "string",
						Multiple: true,
					},
					{
						Name:     "q",
						Type:     "string",
						Multiple: true,
					},
				},
			},
		},
	}

	relationshipsRegistry[IsolationProfileIdentity] = &elemental.Relationship{
		Create: map[string]*elemental.RelationshipInfo{
			"root": {},
		},
		Update: map[string]*elemental.RelationshipInfo{
			"root": {},
		},
		Patch: map[string]*elemental.RelationshipInfo{
			"root": {},
		},
		Delete: map[string]*elemental.RelationshipInfo{
			"root": {},
		},
		Retrieve: map[string]*elemental.RelationshipInfo{
			"root": {
				Parameters: []elemental.ParameterDefinition{
					{
						Name: "propagated",
						Type: "boolean",
					},
				},
			},
		},
		RetrieveMany: map[string]*elemental.RelationshipInfo{
			"processingunitpolicy": {},
			"root": {
				Parameters: []elemental.ParameterDefinition{
					{
						Name:     "q",
						Type:     "string",
						Multiple: true,
					},
					{
						Name: "propagated",
						Type: "boolean",
					},
				},
			},
		},
		Info: map[string]*elemental.RelationshipInfo{
			"processingunitpolicy": {},
			"root": {
				Parameters: []elemental.ParameterDefinition{
					{
						Name:     "q",
						Type:     "string",
						Multiple: true,
					},
					{
						Name: "propagated",
						Type: "boolean",
					},
				},
			},
		},
	}

	relationshipsRegistry[IssueIdentity] = &elemental.Relationship{
		Create: map[string]*elemental.RelationshipInfo{
			"root": {
				Parameters: []elemental.ParameterDefinition{
					{
						Name: "asCookie",
						Type: "boolean",
					},
					{
						Name: "token",
						Type: "string",
					},
				},
			},
		},
	}

	relationshipsRegistry[IssueServiceTokenIdentity] = &elemental.Relationship{
		Create: map[string]*elemental.RelationshipInfo{
			"root": {},
		},
	}

	relationshipsRegistry[LDAPProviderIdentity] = &elemental.Relationship{
		Create: map[string]*elemental.RelationshipInfo{
			"root": {},
		},
		Update: map[string]*elemental.RelationshipInfo{
			"root": {},
		},
		Patch: map[string]*elemental.RelationshipInfo{
			"root": {},
		},
		Delete: map[string]*elemental.RelationshipInfo{
			"root": {
				Parameters: []elemental.ParameterDefinition{
					{
						Name:     "q",
						Type:     "string",
						Multiple: true,
					},
				},
			},
		},
		Retrieve: map[string]*elemental.RelationshipInfo{
			"root": {},
		},
		RetrieveMany: map[string]*elemental.RelationshipInfo{
			"root": {
				Parameters: []elemental.ParameterDefinition{
					{
						Name:     "q",
						Type:     "string",
						Multiple: true,
					},
				},
			},
		},
		Info: map[string]*elemental.RelationshipInfo{
			"root": {
				Parameters: []elemental.ParameterDefinition{
					{
						Name:     "q",
						Type:     "string",
						Multiple: true,
					},
				},
			},
		},
	}

	relationshipsRegistry[LocalCAIdentity] = &elemental.Relationship{
		Create: map[string]*elemental.RelationshipInfo{
			"root": {},
		},
		RetrieveMany: map[string]*elemental.RelationshipInfo{
			"root": {},
		},
		Info: map[string]*elemental.RelationshipInfo{
			"root": {},
		},
	}

	relationshipsRegistry[LogIdentity] = &elemental.Relationship{
		RetrieveMany: map[string]*elemental.RelationshipInfo{
			"installedapp": {},
		},
		Info: map[string]*elemental.RelationshipInfo{
			"installedapp": {},
		},
	}

	relationshipsRegistry[LogoutIdentity] = &elemental.Relationship{
		RetrieveMany: map[string]*elemental.RelationshipInfo{
			"root": {},
		},
		Info: map[string]*elemental.RelationshipInfo{
			"root": {},
		},
	}

	relationshipsRegistry[MessageIdentity] = &elemental.Relationship{
		Create: map[string]*elemental.RelationshipInfo{
			"root": {},
		},
		Update: map[string]*elemental.RelationshipInfo{
			"root": {},
		},
		Patch: map[string]*elemental.RelationshipInfo{
			"root": {},
		},
		Delete: map[string]*elemental.RelationshipInfo{
			"root": {
				Parameters: []elemental.ParameterDefinition{
					{
						Name:     "q",
						Type:     "string",
						Multiple: true,
					},
				},
			},
		},
		Retrieve: map[string]*elemental.RelationshipInfo{
			"root": {
				Parameters: []elemental.ParameterDefinition{
					{
						Name: "propagated",
						Type: "boolean",
					},
				},
			},
		},
		RetrieveMany: map[string]*elemental.RelationshipInfo{
			"root": {
				Parameters: []elemental.ParameterDefinition{
					{
						Name:     "q",
						Type:     "string",
						Multiple: true,
					},
					{
						Name: "propagated",
						Type: "boolean",
					},
				},
			},
		},
		Info: map[string]*elemental.RelationshipInfo{
			"root": {
				Parameters: []elemental.ParameterDefinition{
					{
						Name:     "q",
						Type:     "string",
						Multiple: true,
					},
					{
						Name: "propagated",
						Type: "boolean",
					},
				},
			},
		},
	}

	relationshipsRegistry[MetricsQueryIdentity] = &elemental.Relationship{
		Create: map[string]*elemental.RelationshipInfo{
			"root": {},
		},
		RetrieveMany: map[string]*elemental.RelationshipInfo{
			"root": {
				RequiredParameters: elemental.NewParametersRequirement(
					[][][]string{
						{
							{
								"query",
							},
						},
					},
				),
				Parameters: []elemental.ParameterDefinition{
					{
						Name: "query",
						Type: "string",
					},
					{
						Name: "time",
						Type: "string",
					},
				},
			},
		},
		Info: map[string]*elemental.RelationshipInfo{
			"root": {
				RequiredParameters: elemental.NewParametersRequirement(
					[][][]string{
						{
							{
								"query",
							},
						},
					},
				),
				Parameters: []elemental.ParameterDefinition{
					{
						Name: "query",
						Type: "string",
					},
					{
						Name: "time",
						Type: "string",
					},
				},
			},
		},
	}

	relationshipsRegistry[MetricsQueryRangeIdentity] = &elemental.Relationship{
		Create: map[string]*elemental.RelationshipInfo{
			"root": {},
		},
		RetrieveMany: map[string]*elemental.RelationshipInfo{
			"root": {
				RequiredParameters: elemental.NewParametersRequirement(
					[][][]string{
						{
							{
								"query",
							},
						},
					},
				),
				Parameters: []elemental.ParameterDefinition{
					{
						Name: "end",
						Type: "string",
					},
					{
						Name: "query",
						Type: "string",
					},
					{
						Name: "start",
						Type: "string",
					},
					{
						Name: "step",
						Type: "string",
					},
				},
			},
		},
		Info: map[string]*elemental.RelationshipInfo{
			"root": {
				RequiredParameters: elemental.NewParametersRequirement(
					[][][]string{
						{
							{
								"query",
							},
						},
					},
				),
				Parameters: []elemental.ParameterDefinition{
					{
						Name: "end",
						Type: "string",
					},
					{
						Name: "query",
						Type: "string",
					},
					{
						Name: "start",
						Type: "string",
					},
					{
						Name: "step",
						Type: "string",
					},
				},
			},
		},
	}

	relationshipsRegistry[NamespaceIdentity] = &elemental.Relationship{
		Create: map[string]*elemental.RelationshipInfo{
			"root": {},
		},
		Update: map[string]*elemental.RelationshipInfo{
			"root": {},
		},
		Patch: map[string]*elemental.RelationshipInfo{
			"root": {},
		},
		Delete: map[string]*elemental.RelationshipInfo{
			"root": {
				Parameters: []elemental.ParameterDefinition{
					{
						Name:     "q",
						Type:     "string",
						Multiple: true,
					},
				},
			},
		},
		Retrieve: map[string]*elemental.RelationshipInfo{
			"root": {},
		},
		RetrieveMany: map[string]*elemental.RelationshipInfo{
			"root": {
				Parameters: []elemental.ParameterDefinition{
					{
						Name:     "q",
						Type:     "string",
						Multiple: true,
					},
				},
			},
		},
		Info: map[string]*elemental.RelationshipInfo{
			"root": {
				Parameters: []elemental.ParameterDefinition{
					{
						Name:     "q",
						Type:     "string",
						Multiple: true,
					},
				},
			},
		},
	}

	relationshipsRegistry[NamespaceMappingPolicyIdentity] = &elemental.Relationship{
		Create: map[string]*elemental.RelationshipInfo{
			"root": {},
		},
		Update: map[string]*elemental.RelationshipInfo{
			"root": {},
		},
		Patch: map[string]*elemental.RelationshipInfo{
			"root": {},
		},
		Delete: map[string]*elemental.RelationshipInfo{
			"root": {
				Parameters: []elemental.ParameterDefinition{
					{
						Name:     "q",
						Type:     "string",
						Multiple: true,
					},
				},
			},
		},
		Retrieve: map[string]*elemental.RelationshipInfo{
			"root": {},
		},
		RetrieveMany: map[string]*elemental.RelationshipInfo{
			"root": {
				Parameters: []elemental.ParameterDefinition{
					{
						Name:     "q",
						Type:     "string",
						Multiple: true,
					},
				},
			},
		},
		Info: map[string]*elemental.RelationshipInfo{
			"root": {
				Parameters: []elemental.ParameterDefinition{
					{
						Name:     "q",
						Type:     "string",
						Multiple: true,
					},
				},
			},
		},
	}

	relationshipsRegistry[NamespacePolicyInfoIdentity] = &elemental.Relationship{
		RetrieveMany: map[string]*elemental.RelationshipInfo{
			"root": {},
		},
		Info: map[string]*elemental.RelationshipInfo{
			"root": {},
		},
	}

	relationshipsRegistry[NamespaceRendererIdentity] = &elemental.Relationship{
		Create: map[string]*elemental.RelationshipInfo{
			"root": {},
		},
	}

	relationshipsRegistry[NamespaceTypeIdentity] = &elemental.Relationship{
		RetrieveMany: map[string]*elemental.RelationshipInfo{
			"root": {},
		},
		Info: map[string]*elemental.RelationshipInfo{
			"root": {},
		},
	}

	relationshipsRegistry[NetworkAccessPolicyIdentity] = &elemental.Relationship{
		Create: map[string]*elemental.RelationshipInfo{
			"root": {},
		},
		Update: map[string]*elemental.RelationshipInfo{
			"root": {},
		},
		Patch: map[string]*elemental.RelationshipInfo{
			"root": {},
		},
		Delete: map[string]*elemental.RelationshipInfo{
			"root": {
				Parameters: []elemental.ParameterDefinition{
					{
						Name:     "q",
						Type:     "string",
						Multiple: true,
					},
				},
			},
		},
		Retrieve: map[string]*elemental.RelationshipInfo{
			"root": {
				Parameters: []elemental.ParameterDefinition{
					{
						Name: "propagated",
						Type: "boolean",
					},
				},
			},
		},
		RetrieveMany: map[string]*elemental.RelationshipInfo{
			"root": {
				Parameters: []elemental.ParameterDefinition{
					{
						Name:     "q",
						Type:     "string",
						Multiple: true,
					},
					{
						Name: "propagated",
						Type: "boolean",
					},
				},
			},
		},
		Info: map[string]*elemental.RelationshipInfo{
			"root": {
				Parameters: []elemental.ParameterDefinition{
					{
						Name:     "q",
						Type:     "string",
						Multiple: true,
					},
					{
						Name: "propagated",
						Type: "boolean",
					},
				},
			},
		},
	}

	relationshipsRegistry[NetworkRuleSetPolicyIdentity] = &elemental.Relationship{
		Create: map[string]*elemental.RelationshipInfo{
			"root": {},
		},
		Update: map[string]*elemental.RelationshipInfo{
			"root": {},
		},
		Patch: map[string]*elemental.RelationshipInfo{
			"root": {},
		},
		Delete: map[string]*elemental.RelationshipInfo{
			"root": {
				Parameters: []elemental.ParameterDefinition{
					{
						Name:     "q",
						Type:     "string",
						Multiple: true,
					},
				},
			},
		},
		Retrieve: map[string]*elemental.RelationshipInfo{
			"root": {
				Parameters: []elemental.ParameterDefinition{
					{
						Name: "propagated",
						Type: "boolean",
					},
				},
			},
		},
		RetrieveMany: map[string]*elemental.RelationshipInfo{
			"root": {
				Parameters: []elemental.ParameterDefinition{
					{
						Name:     "q",
						Type:     "string",
						Multiple: true,
					},
					{
						Name: "propagated",
						Type: "boolean",
					},
				},
			},
		},
		Info: map[string]*elemental.RelationshipInfo{
			"root": {
				Parameters: []elemental.ParameterDefinition{
					{
						Name:     "q",
						Type:     "string",
						Multiple: true,
					},
					{
						Name: "propagated",
						Type: "boolean",
					},
				},
			},
		},
	}

	relationshipsRegistry[OAUTHInfoIdentity] = &elemental.Relationship{
		Retrieve: map[string]*elemental.RelationshipInfo{
			"root": {},
		},
		RetrieveMany: map[string]*elemental.RelationshipInfo{
			"namespace": {
				Parameters: []elemental.ParameterDefinition{
					{
						Name: "mode",
						Type: "enum",
						AllowedChoices: []string{
							"oidc",
						},
					},
				},
			},
		},
		Info: map[string]*elemental.RelationshipInfo{
			"namespace": {
				Parameters: []elemental.ParameterDefinition{
					{
						Name: "mode",
						Type: "enum",
						AllowedChoices: []string{
							"oidc",
						},
					},
				},
			},
		},
	}

	relationshipsRegistry[OAUTHKeyIdentity] = &elemental.Relationship{
		Retrieve: map[string]*elemental.RelationshipInfo{
			"root": {},
		},
		RetrieveMany: map[string]*elemental.RelationshipInfo{
			"namespace": {
				Parameters: []elemental.ParameterDefinition{
					{
						Name: "mode",
						Type: "enum",
						AllowedChoices: []string{
							"oidc",
						},
					},
				},
			},
		},
		Info: map[string]*elemental.RelationshipInfo{
			"namespace": {
				Parameters: []elemental.ParameterDefinition{
					{
						Name: "mode",
						Type: "enum",
						AllowedChoices: []string{
							"oidc",
						},
					},
				},
			},
		},
	}

	relationshipsRegistry[OIDCProviderIdentity] = &elemental.Relationship{
		Create: map[string]*elemental.RelationshipInfo{
			"root": {},
		},
		Update: map[string]*elemental.RelationshipInfo{
			"root": {},
		},
		Patch: map[string]*elemental.RelationshipInfo{
			"root": {},
		},
		Delete: map[string]*elemental.RelationshipInfo{
			"root": {
				Parameters: []elemental.ParameterDefinition{
					{
						Name:     "q",
						Type:     "string",
						Multiple: true,
					},
				},
			},
		},
		Retrieve: map[string]*elemental.RelationshipInfo{
			"root": {},
		},
		RetrieveMany: map[string]*elemental.RelationshipInfo{
			"root": {
				Parameters: []elemental.ParameterDefinition{
					{
						Name:     "q",
						Type:     "string",
						Multiple: true,
					},
				},
			},
		},
		Info: map[string]*elemental.RelationshipInfo{
			"root": {
				Parameters: []elemental.ParameterDefinition{
					{
						Name:     "q",
						Type:     "string",
						Multiple: true,
					},
				},
			},
		},
	}

	relationshipsRegistry[OrganizationalMetadataIdentity] = &elemental.Relationship{
		RetrieveMany: map[string]*elemental.RelationshipInfo{
			"root": {},
		},
		Info: map[string]*elemental.RelationshipInfo{
			"root": {},
		},
	}

	relationshipsRegistry[PCCProviderIdentity] = &elemental.Relationship{
		Create: map[string]*elemental.RelationshipInfo{
			"root": {},
		},
		Update: map[string]*elemental.RelationshipInfo{
			"root": {},
		},
		Patch: map[string]*elemental.RelationshipInfo{
			"root": {},
		},
		Delete: map[string]*elemental.RelationshipInfo{
			"root": {
				Parameters: []elemental.ParameterDefinition{
					{
						Name:     "q",
						Type:     "string",
						Multiple: true,
					},
				},
			},
		},
		Retrieve: map[string]*elemental.RelationshipInfo{
			"root": {},
		},
		RetrieveMany: map[string]*elemental.RelationshipInfo{
			"root": {
				Parameters: []elemental.ParameterDefinition{
					{
						Name:     "q",
						Type:     "string",
						Multiple: true,
					},
				},
			},
		},
		Info: map[string]*elemental.RelationshipInfo{
			"root": {
				Parameters: []elemental.ParameterDefinition{
					{
						Name:     "q",
						Type:     "string",
						Multiple: true,
					},
				},
			},
		},
	}

	relationshipsRegistry[PCSearchResultIdentity] = &elemental.Relationship{}

	relationshipsRegistry[PCTimeRangeIdentity] = &elemental.Relationship{}

	relationshipsRegistry[PUTrafficActionIdentity] = &elemental.Relationship{
		Create: map[string]*elemental.RelationshipInfo{
			"root": {},
		},
		RetrieveMany: map[string]*elemental.RelationshipInfo{
			"root": {},
		},
		Info: map[string]*elemental.RelationshipInfo{
			"root": {},
		},
	}

	relationshipsRegistry[PacketReportIdentity] = &elemental.Relationship{
		Create: map[string]*elemental.RelationshipInfo{
			"root": {},
		},
	}

	relationshipsRegistry[PasswordResetIdentity] = &elemental.Relationship{
		Create: map[string]*elemental.RelationshipInfo{
			"root": {},
		},
		RetrieveMany: map[string]*elemental.RelationshipInfo{
			"root": {
				RequiredParameters: elemental.NewParametersRequirement(
					[][][]string{
						{
							{
								"email",
							},
						},
					},
				),
				Parameters: []elemental.ParameterDefinition{
					{
						Name: "email",
						Type: "string",
					},
				},
			},
		},
		Info: map[string]*elemental.RelationshipInfo{
			"root": {
				RequiredParameters: elemental.NewParametersRequirement(
					[][][]string{
						{
							{
								"email",
							},
						},
					},
				),
				Parameters: []elemental.ParameterDefinition{
					{
						Name: "email",
						Type: "string",
					},
				},
			},
		},
	}

	relationshipsRegistry[PingProbeIdentity] = &elemental.Relationship{
		Create: map[string]*elemental.RelationshipInfo{
			"processingunit": {},
		},
		Retrieve: map[string]*elemental.RelationshipInfo{
			"root": {},
		},
	}

	relationshipsRegistry[PingRequestIdentity] = &elemental.Relationship{
		Create: map[string]*elemental.RelationshipInfo{
			"root": {},
		},
	}

	relationshipsRegistry[PingResultIdentity] = &elemental.Relationship{
		RetrieveMany: map[string]*elemental.RelationshipInfo{
			"root": {
				Parameters: []elemental.ParameterDefinition{
					{
						Name:     "q",
						Type:     "string",
						Multiple: true,
					},
				},
			},
		},
		Info: map[string]*elemental.RelationshipInfo{
			"root": {
				Parameters: []elemental.ParameterDefinition{
					{
						Name:     "q",
						Type:     "string",
						Multiple: true,
					},
				},
			},
		},
	}

	relationshipsRegistry[PlanIdentity] = &elemental.Relationship{
		Retrieve: map[string]*elemental.RelationshipInfo{
			"root": {},
		},
		RetrieveMany: map[string]*elemental.RelationshipInfo{
			"root": {},
		},
		Info: map[string]*elemental.RelationshipInfo{
			"root": {},
		},
	}

	relationshipsRegistry[PokeIdentity] = &elemental.Relationship{
		RetrieveMany: map[string]*elemental.RelationshipInfo{
			"enforcer": {
				Parameters: []elemental.ParameterDefinition{
					{
						Name: "cpuload",
						Type: "float",
					},
					{
						Name: "enforcementStatus",
						Type: "enum",
						AllowedChoices: []string{
							"Failed",
							"Inactive",
							"Active",
						},
					},
					{
						Name: "forceFullPoke",
						Type: "boolean",
					},
					{
						Name: "memory",
						Type: "integer",
					},
					{
						Name: "processes",
						Type: "integer",
					},
					{
						Name: "sessionClose",
						Type: "boolean",
					},
					{
						Name: "sessionID",
						Type: "string",
					},
					{
						Name: "status",
						Type: "enum",
						AllowedChoices: []string{
							"Registered",
							"Connected",
							"Disconnected",
						},
					},
					{
						Name: "ts",
						Type: "time",
					},
					{
						Name: "version",
						Type: "string",
					},
					{
						Name: "zhash",
						Type: "integer",
					},
				},
			},
			"processingunit": {
				Parameters: []elemental.ParameterDefinition{
					{
						Name: "enforcementStatus",
						Type: "enum",
						AllowedChoices: []string{
							"Failed",
							"Inactive",
							"Active",
						},
					},
					{
						Name: "forceFullPoke",
						Type: "boolean",
					},
					{
						Name: "notify",
						Type: "boolean",
					},
					{
						Name: "status",
						Type: "enum",
						AllowedChoices: []string{
							"Initialized",
							"Paused",
							"Running",
							"Stopped",
						},
					},
					{
						Name: "ts",
						Type: "time",
					},
					{
						Name: "zhash",
						Type: "integer",
					},
				},
			},
		},
		Info: map[string]*elemental.RelationshipInfo{
			"enforcer": {
				Parameters: []elemental.ParameterDefinition{
					{
						Name: "cpuload",
						Type: "float",
					},
					{
						Name: "enforcementStatus",
						Type: "enum",
						AllowedChoices: []string{
							"Failed",
							"Inactive",
							"Active",
						},
					},
					{
						Name: "forceFullPoke",
						Type: "boolean",
					},
					{
						Name: "memory",
						Type: "integer",
					},
					{
						Name: "processes",
						Type: "integer",
					},
					{
						Name: "sessionClose",
						Type: "boolean",
					},
					{
						Name: "sessionID",
						Type: "string",
					},
					{
						Name: "status",
						Type: "enum",
						AllowedChoices: []string{
							"Registered",
							"Connected",
							"Disconnected",
						},
					},
					{
						Name: "ts",
						Type: "time",
					},
					{
						Name: "version",
						Type: "string",
					},
					{
						Name: "zhash",
						Type: "integer",
					},
				},
			},
			"processingunit": {
				Parameters: []elemental.ParameterDefinition{
					{
						Name: "enforcementStatus",
						Type: "enum",
						AllowedChoices: []string{
							"Failed",
							"Inactive",
							"Active",
						},
					},
					{
						Name: "forceFullPoke",
						Type: "boolean",
					},
					{
						Name: "notify",
						Type: "boolean",
					},
					{
						Name: "status",
						Type: "enum",
						AllowedChoices: []string{
							"Initialized",
							"Paused",
							"Running",
							"Stopped",
						},
					},
					{
						Name: "ts",
						Type: "time",
					},
					{
						Name: "zhash",
						Type: "integer",
					},
				},
			},
		},
	}

	relationshipsRegistry[PolicyIdentity] = &elemental.Relationship{
		Delete: map[string]*elemental.RelationshipInfo{
			"root": {
				Parameters: []elemental.ParameterDefinition{
					{
						Name:     "q",
						Type:     "string",
						Multiple: true,
					},
				},
			},
		},
		Retrieve: map[string]*elemental.RelationshipInfo{
			"root": {},
		},
		RetrieveMany: map[string]*elemental.RelationshipInfo{
			"root": {
				Parameters: []elemental.ParameterDefinition{
					{
						Name:     "q",
						Type:     "string",
						Multiple: true,
					},
					{
						Name: "propagated",
						Type: "boolean",
					},
				},
			},
		},
		Info: map[string]*elemental.RelationshipInfo{
			"root": {
				Parameters: []elemental.ParameterDefinition{
					{
						Name:     "q",
						Type:     "string",
						Multiple: true,
					},
					{
						Name: "propagated",
						Type: "boolean",
					},
				},
			},
		},
	}

	relationshipsRegistry[PolicyGraphIdentity] = &elemental.Relationship{
		Create: map[string]*elemental.RelationshipInfo{
			"root": {
				Parameters: []elemental.ParameterDefinition{
					{
						Name: "view",
						Type: "string",
					},
				},
			},
		},
	}

	relationshipsRegistry[PolicyRefreshIdentity] = &elemental.Relationship{}

	relationshipsRegistry[PolicyRendererIdentity] = &elemental.Relationship{
		Create: map[string]*elemental.RelationshipInfo{
			"root": {},
		},
	}

	relationshipsRegistry[PolicyRuleIdentity] = &elemental.Relationship{
		Retrieve: map[string]*elemental.RelationshipInfo{
			"root": {},
		},
	}

	relationshipsRegistry[PolicyTTLIdentity] = &elemental.Relationship{}

	relationshipsRegistry[PollAccountIdentity] = &elemental.Relationship{
		Create: map[string]*elemental.RelationshipInfo{
			"root": {},
		},
	}

	relationshipsRegistry[ProcessingUnitIdentity] = &elemental.Relationship{
		Create: map[string]*elemental.RelationshipInfo{
			"root": {},
		},
		Update: map[string]*elemental.RelationshipInfo{
			"root": {},
		},
		Patch: map[string]*elemental.RelationshipInfo{
			"root": {},
		},
		Delete: map[string]*elemental.RelationshipInfo{
			"root": {
				Parameters: []elemental.ParameterDefinition{
					{
						Name:     "q",
						Type:     "string",
						Multiple: true,
					},
				},
			},
		},
		Retrieve: map[string]*elemental.RelationshipInfo{
			"root": {
				Parameters: []elemental.ParameterDefinition{
					{
						Name: "archived",
						Type: "boolean",
					},
				},
			},
		},
		RetrieveMany: map[string]*elemental.RelationshipInfo{
			"fileaccesspolicy": {},
			"infrastructurepolicy": {
				Parameters: []elemental.ParameterDefinition{
					{
						Name:         "mode",
						Type:         "enum",
						DefaultValue: "object",
						AllowedChoices: []string{
							"subject",
							"object",
						},
					},
				},
			},
			"networkaccesspolicy": {
				Parameters: []elemental.ParameterDefinition{
					{
						Name:         "mode",
						Type:         "enum",
						DefaultValue: "object",
						AllowedChoices: []string{
							"subject",
							"object",
						},
					},
				},
			},
			"networkrulesetpolicy": {
				Parameters: []elemental.ParameterDefinition{
					{
						Name:         "mode",
						Type:         "enum",
						DefaultValue: "object",
						AllowedChoices: []string{
							"subject",
							"object",
						},
					},
				},
			},
			"processingunitpolicy": {},
			"root": {
				Parameters: []elemental.ParameterDefinition{
					{
						Name:     "q",
						Type:     "string",
						Multiple: true,
					},
					{
						Name: "archived",
						Type: "boolean",
					},
				},
			},
			"service":                 {},
			"servicedependencypolicy": {},
			"vulnerability":           {},
		},
		Info: map[string]*elemental.RelationshipInfo{
			"fileaccesspolicy": {},
			"infrastructurepolicy": {
				Parameters: []elemental.ParameterDefinition{
					{
						Name:         "mode",
						Type:         "enum",
						DefaultValue: "object",
						AllowedChoices: []string{
							"subject",
							"object",
						},
					},
				},
			},
			"networkaccesspolicy": {
				Parameters: []elemental.ParameterDefinition{
					{
						Name:         "mode",
						Type:         "enum",
						DefaultValue: "object",
						AllowedChoices: []string{
							"subject",
							"object",
						},
					},
				},
			},
			"networkrulesetpolicy": {
				Parameters: []elemental.ParameterDefinition{
					{
						Name:         "mode",
						Type:         "enum",
						DefaultValue: "object",
						AllowedChoices: []string{
							"subject",
							"object",
						},
					},
				},
			},
			"processingunitpolicy": {},
			"root": {
				Parameters: []elemental.ParameterDefinition{
					{
						Name:     "q",
						Type:     "string",
						Multiple: true,
					},
					{
						Name: "archived",
						Type: "boolean",
					},
				},
			},
			"service":                 {},
			"servicedependencypolicy": {},
			"vulnerability":           {},
		},
	}

	relationshipsRegistry[ProcessingUnitPolicyIdentity] = &elemental.Relationship{
		Create: map[string]*elemental.RelationshipInfo{
			"root": {},
		},
		Update: map[string]*elemental.RelationshipInfo{
			"root": {},
		},
		Patch: map[string]*elemental.RelationshipInfo{
			"root": {},
		},
		Delete: map[string]*elemental.RelationshipInfo{
			"root": {
				Parameters: []elemental.ParameterDefinition{
					{
						Name:     "q",
						Type:     "string",
						Multiple: true,
					},
				},
			},
		},
		Retrieve: map[string]*elemental.RelationshipInfo{
			"root": {},
		},
		RetrieveMany: map[string]*elemental.RelationshipInfo{
			"root": {
				Parameters: []elemental.ParameterDefinition{
					{
						Name:     "q",
						Type:     "string",
						Multiple: true,
					},
					{
						Name: "propagated",
						Type: "boolean",
					},
				},
			},
		},
		Info: map[string]*elemental.RelationshipInfo{
			"root": {
				Parameters: []elemental.ParameterDefinition{
					{
						Name:     "q",
						Type:     "string",
						Multiple: true,
					},
					{
						Name: "propagated",
						Type: "boolean",
					},
				},
			},
		},
	}

	relationshipsRegistry[ProcessingUnitRefreshIdentity] = &elemental.Relationship{
		Create: map[string]*elemental.RelationshipInfo{
			"processingunit": {},
		},
	}

	relationshipsRegistry[QuotaCheckIdentity] = &elemental.Relationship{
		Create: map[string]*elemental.RelationshipInfo{
			"root": {
				Parameters: []elemental.ParameterDefinition{
					{
						Name: "remaining",
						Type: "boolean",
					},
				},
			},
		},
	}

	relationshipsRegistry[QuotaPolicyIdentity] = &elemental.Relationship{
		Create: map[string]*elemental.RelationshipInfo{
			"root": {},
		},
		Update: map[string]*elemental.RelationshipInfo{
			"root": {},
		},
		Patch: map[string]*elemental.RelationshipInfo{
			"root": {},
		},
		Delete: map[string]*elemental.RelationshipInfo{
			"root": {
				Parameters: []elemental.ParameterDefinition{
					{
						Name:     "q",
						Type:     "string",
						Multiple: true,
					},
				},
			},
		},
		Retrieve: map[string]*elemental.RelationshipInfo{
			"root": {},
		},
		RetrieveMany: map[string]*elemental.RelationshipInfo{
			"root": {
				Parameters: []elemental.ParameterDefinition{
					{
						Name:     "q",
						Type:     "string",
						Multiple: true,
					},
					{
						Name: "propagated",
						Type: "boolean",
					},
				},
			},
		},
		Info: map[string]*elemental.RelationshipInfo{
			"root": {
				Parameters: []elemental.ParameterDefinition{
					{
						Name:     "q",
						Type:     "string",
						Multiple: true,
					},
					{
						Name: "propagated",
						Type: "boolean",
					},
				},
			},
		},
	}

	relationshipsRegistry[RecipeIdentity] = &elemental.Relationship{
		Create: map[string]*elemental.RelationshipInfo{
			"root": {},
		},
		Update: map[string]*elemental.RelationshipInfo{
			"root": {},
		},
		Patch: map[string]*elemental.RelationshipInfo{
			"root": {},
		},
		Delete: map[string]*elemental.RelationshipInfo{
			"root": {
				Parameters: []elemental.ParameterDefinition{
					{
						Name:     "q",
						Type:     "string",
						Multiple: true,
					},
				},
			},
		},
		Retrieve: map[string]*elemental.RelationshipInfo{
			"root": {
				Parameters: []elemental.ParameterDefinition{
					{
						Name: "propagated",
						Type: "boolean",
					},
				},
			},
		},
		RetrieveMany: map[string]*elemental.RelationshipInfo{
			"root": {
				Parameters: []elemental.ParameterDefinition{
					{
						Name:     "q",
						Type:     "string",
						Multiple: true,
					},
					{
						Name: "propagated",
						Type: "boolean",
					},
				},
			},
		},
		Info: map[string]*elemental.RelationshipInfo{
			"root": {
				Parameters: []elemental.ParameterDefinition{
					{
						Name:     "q",
						Type:     "string",
						Multiple: true,
					},
					{
						Name: "propagated",
						Type: "boolean",
					},
				},
			},
		},
	}

	relationshipsRegistry[RemoteProcessorIdentity] = &elemental.Relationship{
		Create: map[string]*elemental.RelationshipInfo{
			"root": {},
		},
	}

	relationshipsRegistry[RenderTemplateIdentity] = &elemental.Relationship{
		Create: map[string]*elemental.RelationshipInfo{
			"root": {},
		},
	}

	relationshipsRegistry[RenderedPolicyIdentity] = &elemental.Relationship{
		Create: map[string]*elemental.RelationshipInfo{
			"root": {
				Parameters: []elemental.ParameterDefinition{
					{
						Name: "renderer",
						Type: "enum",
						AllowedChoices: []string{
							"v1",
							"v2",
						},
					},
				},
			},
		},
		RetrieveMany: map[string]*elemental.RelationshipInfo{
			"processingunit": {
				Parameters: []elemental.ParameterDefinition{
					{
						Name: "renderer",
						Type: "enum",
						AllowedChoices: []string{
							"v1",
							"v2",
						},
					},
				},
			},
		},
		Info: map[string]*elemental.RelationshipInfo{
			"processingunit": {
				Parameters: []elemental.ParameterDefinition{
					{
						Name: "renderer",
						Type: "enum",
						AllowedChoices: []string{
							"v1",
							"v2",
						},
					},
				},
			},
		},
	}

	relationshipsRegistry[ReportIdentity] = &elemental.Relationship{
		Create: map[string]*elemental.RelationshipInfo{
			"root": {},
		},
	}

	relationshipsRegistry[ReportsQueryIdentity] = &elemental.Relationship{
		Create: map[string]*elemental.RelationshipInfo{
			"root": {
				RequiredParameters: elemental.NewParametersRequirement(
					[][][]string{
						{
							{
								"endRelative",
							},
							{
								"startRelative",
							},
							{
								"startRelative",
								"endRelative",
							},
							{
								"startRelative",
								"endAbsolute",
							},
							{
								"startAbsolute",
								"endRelative",
							},
							{
								"startAbsolute",
								"endAbsolute",
							},
						},
					},
				),
				Parameters: []elemental.ParameterDefinition{
					{
						Name:     "q",
						Type:     "string",
						Multiple: true,
					},
					{
						Name: "endAbsolute",
						Type: "time",
					},
					{
						Name: "endRelative",
						Type: "duration",
					},
					{
						Name: "startAbsolute",
						Type: "time",
					},
					{
						Name: "startRelative",
						Type: "duration",
					},
				},
			},
		},
	}

	relationshipsRegistry[RevocationIdentity] = &elemental.Relationship{
		Update: map[string]*elemental.RelationshipInfo{
			"root": {},
		},
		Patch: map[string]*elemental.RelationshipInfo{
			"root": {},
		},
		RetrieveMany: map[string]*elemental.RelationshipInfo{
			"root": {
				Parameters: []elemental.ParameterDefinition{
					{
						Name:     "q",
						Type:     "string",
						Multiple: true,
					},
				},
			},
		},
		Info: map[string]*elemental.RelationshipInfo{
			"root": {
				Parameters: []elemental.ParameterDefinition{
					{
						Name:     "q",
						Type:     "string",
						Multiple: true,
					},
				},
			},
		},
	}

	relationshipsRegistry[RoleIdentity] = &elemental.Relationship{
		RetrieveMany: map[string]*elemental.RelationshipInfo{
			"root": {},
		},
		Info: map[string]*elemental.RelationshipInfo{
			"root": {},
		},
	}

	relationshipsRegistry[RootIdentity] = &elemental.Relationship{}

	relationshipsRegistry[SAMLProviderIdentity] = &elemental.Relationship{
		Create: map[string]*elemental.RelationshipInfo{
			"root": {},
		},
		Update: map[string]*elemental.RelationshipInfo{
			"root": {},
		},
		Patch: map[string]*elemental.RelationshipInfo{
			"root": {},
		},
		Delete: map[string]*elemental.RelationshipInfo{
			"root": {
				Parameters: []elemental.ParameterDefinition{
					{
						Name:     "q",
						Type:     "string",
						Multiple: true,
					},
				},
			},
		},
		Retrieve: map[string]*elemental.RelationshipInfo{
			"root": {},
		},
		RetrieveMany: map[string]*elemental.RelationshipInfo{
			"root": {
				Parameters: []elemental.ParameterDefinition{
					{
						Name:     "q",
						Type:     "string",
						Multiple: true,
					},
				},
			},
		},
		Info: map[string]*elemental.RelationshipInfo{
			"root": {
				Parameters: []elemental.ParameterDefinition{
					{
						Name:     "q",
						Type:     "string",
						Multiple: true,
					},
				},
			},
		},
	}

	relationshipsRegistry[SSHAuthorityIdentity] = &elemental.Relationship{
		Create: map[string]*elemental.RelationshipInfo{
			"root": {},
		},
		Delete: map[string]*elemental.RelationshipInfo{
			"root": {},
		},
	}

	relationshipsRegistry[SSHAuthorizationPolicyIdentity] = &elemental.Relationship{
		Create: map[string]*elemental.RelationshipInfo{
			"root": {},
		},
		Update: map[string]*elemental.RelationshipInfo{
			"root": {},
		},
		Patch: map[string]*elemental.RelationshipInfo{
			"root": {},
		},
		Delete: map[string]*elemental.RelationshipInfo{
			"root": {
				Parameters: []elemental.ParameterDefinition{
					{
						Name:     "q",
						Type:     "string",
						Multiple: true,
					},
				},
			},
		},
		Retrieve: map[string]*elemental.RelationshipInfo{
			"root": {
				Parameters: []elemental.ParameterDefinition{
					{
						Name: "propagated",
						Type: "boolean",
					},
				},
			},
		},
		RetrieveMany: map[string]*elemental.RelationshipInfo{
			"root": {
				Parameters: []elemental.ParameterDefinition{
					{
						Name:     "q",
						Type:     "string",
						Multiple: true,
					},
					{
						Name: "propagated",
						Type: "boolean",
					},
				},
			},
		},
		Info: map[string]*elemental.RelationshipInfo{
			"root": {
				Parameters: []elemental.ParameterDefinition{
					{
						Name:     "q",
						Type:     "string",
						Multiple: true,
					},
					{
						Name: "propagated",
						Type: "boolean",
					},
				},
			},
		},
	}

	relationshipsRegistry[SSHCertificateIdentity] = &elemental.Relationship{
		Create: map[string]*elemental.RelationshipInfo{
			"root": {},
		},
	}

	relationshipsRegistry[SSHIdentityIdentity] = &elemental.Relationship{
		Create: map[string]*elemental.RelationshipInfo{
			"root": {},
		},
	}

	relationshipsRegistry[SandboxIdentity] = &elemental.Relationship{
		Create: map[string]*elemental.RelationshipInfo{
			"root": {},
		},
	}

	relationshipsRegistry[SearchIdentity] = &elemental.Relationship{
		RetrieveMany: map[string]*elemental.RelationshipInfo{
			"root": {
				RequiredParameters: elemental.NewParametersRequirement(
					[][][]string{
						{
							{
								"q",
							},
						},
					},
				),
				Parameters: []elemental.ParameterDefinition{
					{
						Name: "q",
						Type: "string",
					},
				},
			},
		},
		Info: map[string]*elemental.RelationshipInfo{
			"root": {
				RequiredParameters: elemental.NewParametersRequirement(
					[][][]string{
						{
							{
								"q",
							},
						},
					},
				),
				Parameters: []elemental.ParameterDefinition{
					{
						Name: "q",
						Type: "string",
					},
				},
			},
		},
	}

	relationshipsRegistry[ServiceIdentity] = &elemental.Relationship{
		Create: map[string]*elemental.RelationshipInfo{
			"root": {},
		},
		Update: map[string]*elemental.RelationshipInfo{
			"root": {},
		},
		Patch: map[string]*elemental.RelationshipInfo{
			"root": {},
		},
		Delete: map[string]*elemental.RelationshipInfo{
			"root": {
				Parameters: []elemental.ParameterDefinition{
					{
						Name:     "q",
						Type:     "string",
						Multiple: true,
					},
				},
			},
		},
		Retrieve: map[string]*elemental.RelationshipInfo{
			"root": {
				Parameters: []elemental.ParameterDefinition{
					{
						Name: "archived",
						Type: "boolean",
					},
					{
						Name: "propagated",
						Type: "boolean",
					},
				},
			},
		},
		RetrieveMany: map[string]*elemental.RelationshipInfo{
			"infrastructurepolicy": {
				Parameters: []elemental.ParameterDefinition{
					{
						Name:         "mode",
						Type:         "enum",
						DefaultValue: "object",
						AllowedChoices: []string{
							"subject",
							"object",
						},
					},
				},
			},
			"networkaccesspolicy": {
				Parameters: []elemental.ParameterDefinition{
					{
						Name:         "mode",
						Type:         "enum",
						DefaultValue: "object",
						AllowedChoices: []string{
							"subject",
							"object",
						},
					},
				},
			},
			"networkrulesetpolicy": {
				Parameters: []elemental.ParameterDefinition{
					{
						Name:         "mode",
						Type:         "enum",
						DefaultValue: "object",
						AllowedChoices: []string{
							"subject",
							"object",
						},
					},
				},
			},
			"processingunit": {},
			"root": {
				Parameters: []elemental.ParameterDefinition{
					{
						Name:     "q",
						Type:     "string",
						Multiple: true,
					},
					{
						Name: "archived",
						Type: "boolean",
					},
					{
						Name: "propagated",
						Type: "boolean",
					},
				},
			},
			"servicedependencypolicy": {},
		},
		Info: map[string]*elemental.RelationshipInfo{
			"infrastructurepolicy": {
				Parameters: []elemental.ParameterDefinition{
					{
						Name:         "mode",
						Type:         "enum",
						DefaultValue: "object",
						AllowedChoices: []string{
							"subject",
							"object",
						},
					},
				},
			},
			"networkaccesspolicy": {
				Parameters: []elemental.ParameterDefinition{
					{
						Name:         "mode",
						Type:         "enum",
						DefaultValue: "object",
						AllowedChoices: []string{
							"subject",
							"object",
						},
					},
				},
			},
			"networkrulesetpolicy": {
				Parameters: []elemental.ParameterDefinition{
					{
						Name:         "mode",
						Type:         "enum",
						DefaultValue: "object",
						AllowedChoices: []string{
							"subject",
							"object",
						},
					},
				},
			},
			"processingunit": {},
			"root": {
				Parameters: []elemental.ParameterDefinition{
					{
						Name:     "q",
						Type:     "string",
						Multiple: true,
					},
					{
						Name: "archived",
						Type: "boolean",
					},
					{
						Name: "propagated",
						Type: "boolean",
					},
				},
			},
			"servicedependencypolicy": {},
		},
	}

	relationshipsRegistry[ServiceDependencyPolicyIdentity] = &elemental.Relationship{
		Create: map[string]*elemental.RelationshipInfo{
			"root": {},
		},
		Update: map[string]*elemental.RelationshipInfo{
			"root": {},
		},
		Patch: map[string]*elemental.RelationshipInfo{
			"root": {},
		},
		Delete: map[string]*elemental.RelationshipInfo{
			"root": {
				Parameters: []elemental.ParameterDefinition{
					{
						Name:     "q",
						Type:     "string",
						Multiple: true,
					},
				},
			},
		},
		Retrieve: map[string]*elemental.RelationshipInfo{
			"root": {},
		},
		RetrieveMany: map[string]*elemental.RelationshipInfo{
			"root": {
				Parameters: []elemental.ParameterDefinition{
					{
						Name:     "q",
						Type:     "string",
						Multiple: true,
					},
					{
						Name: "propagated",
						Type: "boolean",
					},
				},
			},
		},
		Info: map[string]*elemental.RelationshipInfo{
			"root": {
				Parameters: []elemental.ParameterDefinition{
					{
						Name:     "q",
						Type:     "string",
						Multiple: true,
					},
					{
						Name: "propagated",
						Type: "boolean",
					},
				},
			},
		},
	}

	relationshipsRegistry[ServicePublicationIdentity] = &elemental.Relationship{
		Create: map[string]*elemental.RelationshipInfo{
			"root": {},
		},
	}

	relationshipsRegistry[ServiceTokenIdentity] = &elemental.Relationship{
		Create: map[string]*elemental.RelationshipInfo{
			"root": {},
		},
	}

	relationshipsRegistry[SquallTagIdentity] = &elemental.Relationship{
		RetrieveMany: map[string]*elemental.RelationshipInfo{
			"root": {
				RequiredParameters: elemental.NewParametersRequirement(
					[][][]string{
						{
							{
								"identity",
							},
						},
					},
				),
				Parameters: []elemental.ParameterDefinition{
					{
						Name: "identity",
						Type: "string",
					},
				},
			},
		},
		Info: map[string]*elemental.RelationshipInfo{
			"root": {
				RequiredParameters: elemental.NewParametersRequirement(
					[][][]string{
						{
							{
								"identity",
							},
						},
					},
				),
				Parameters: []elemental.ParameterDefinition{
					{
						Name: "identity",
						Type: "string",
					},
				},
			},
		},
	}

	relationshipsRegistry[StatsInfoIdentity] = &elemental.Relationship{
		Create: map[string]*elemental.RelationshipInfo{
			"root": {},
		},
	}

	relationshipsRegistry[StatsQueryIdentity] = &elemental.Relationship{
		Create: map[string]*elemental.RelationshipInfo{
			"root": {
				RequiredParameters: elemental.NewParametersRequirement(
					[][][]string{
						{
							{
								"endRelative",
							},
							{
								"startRelative",
							},
							{
								"startRelative",
								"endRelative",
							},
							{
								"startRelative",
								"endAbsolute",
							},
							{
								"startAbsolute",
								"endRelative",
							},
							{
								"startAbsolute",
								"endAbsolute",
							},
						},
					},
				),
				Parameters: []elemental.ParameterDefinition{
					{
						Name: "endAbsolute",
						Type: "time",
					},
					{
						Name: "endRelative",
						Type: "duration",
					},
					{
						Name: "startAbsolute",
						Type: "time",
					},
					{
						Name: "startRelative",
						Type: "duration",
					},
					{
						Name:     "q",
						Type:     "string",
						Multiple: true,
					},
				},
			},
		},
	}

	relationshipsRegistry[SuggestedPolicyIdentity] = &elemental.Relationship{
		RetrieveMany: map[string]*elemental.RelationshipInfo{
			"root": {
				RequiredParameters: elemental.NewParametersRequirement(
					[][][]string{
						{
							{
								"endRelative",
							},
							{
								"startRelative",
							},
							{
								"startRelative",
								"endRelative",
							},
							{
								"startRelative",
								"endAbsolute",
							},
							{
								"startAbsolute",
								"endRelative",
							},
							{
								"startAbsolute",
								"endAbsolute",
							},
						},
					},
				),
				Parameters: []elemental.ParameterDefinition{
					{
						Name: "filterAction",
						Type: "enum",
						AllowedChoices: []string{
							"include",
							"exclude",
						},
					},
					{
						Name:     "filterTags",
						Type:     "string",
						Multiple: true,
					},
					{
						Name: "endAbsolute",
						Type: "time",
					},
					{
						Name: "endRelative",
						Type: "duration",
					},
					{
						Name: "startAbsolute",
						Type: "time",
					},
					{
						Name: "startRelative",
						Type: "duration",
					},
					{
						Name: "flowOffset",
						Type: "duration",
					},
				},
			},
		},
		Info: map[string]*elemental.RelationshipInfo{
			"root": {
				RequiredParameters: elemental.NewParametersRequirement(
					[][][]string{
						{
							{
								"endRelative",
							},
							{
								"startRelative",
							},
							{
								"startRelative",
								"endRelative",
							},
							{
								"startRelative",
								"endAbsolute",
							},
							{
								"startAbsolute",
								"endRelative",
							},
							{
								"startAbsolute",
								"endAbsolute",
							},
						},
					},
				),
				Parameters: []elemental.ParameterDefinition{
					{
						Name: "filterAction",
						Type: "enum",
						AllowedChoices: []string{
							"include",
							"exclude",
						},
					},
					{
						Name:     "filterTags",
						Type:     "string",
						Multiple: true,
					},
					{
						Name: "endAbsolute",
						Type: "time",
					},
					{
						Name: "endRelative",
						Type: "duration",
					},
					{
						Name: "startAbsolute",
						Type: "time",
					},
					{
						Name: "startRelative",
						Type: "duration",
					},
					{
						Name: "flowOffset",
						Type: "duration",
					},
				},
			},
		},
	}

	relationshipsRegistry[TagIdentity] = &elemental.Relationship{
		RetrieveMany: map[string]*elemental.RelationshipInfo{
			"root": {
				Parameters: []elemental.ParameterDefinition{
					{
						Name: "onlyPolicyTags",
						Type: "boolean",
					},
					{
						Name:     "q",
						Type:     "string",
						Multiple: true,
					},
				},
			},
		},
		Info: map[string]*elemental.RelationshipInfo{
			"root": {
				Parameters: []elemental.ParameterDefinition{
					{
						Name: "onlyPolicyTags",
						Type: "boolean",
					},
					{
						Name:     "q",
						Type:     "string",
						Multiple: true,
					},
				},
			},
		},
	}

	relationshipsRegistry[TagInjectIdentity] = &elemental.Relationship{
		Create: map[string]*elemental.RelationshipInfo{
			"root": {},
		},
	}

	relationshipsRegistry[TagPrefixIdentity] = &elemental.Relationship{
		Create: map[string]*elemental.RelationshipInfo{
			"root": {},
		},
		RetrieveMany: map[string]*elemental.RelationshipInfo{
			"root": {},
		},
		Info: map[string]*elemental.RelationshipInfo{
			"root": {},
		},
	}

	relationshipsRegistry[TagValueIdentity] = &elemental.Relationship{
		RetrieveMany: map[string]*elemental.RelationshipInfo{
			"root": {
				RequiredParameters: elemental.NewParametersRequirement(
					[][][]string{
						{
							{
								"key",
							},
						},
					},
				),
				Parameters: []elemental.ParameterDefinition{
					{
						Name:     "key",
						Type:     "string",
						Multiple: true,
					},
				},
			},
		},
		Info: map[string]*elemental.RelationshipInfo{
			"root": {
				RequiredParameters: elemental.NewParametersRequirement(
					[][][]string{
						{
							{
								"key",
							},
						},
					},
				),
				Parameters: []elemental.ParameterDefinition{
					{
						Name:     "key",
						Type:     "string",
						Multiple: true,
					},
				},
			},
		},
	}

	relationshipsRegistry[TenantIdentity] = &elemental.Relationship{
		Create: map[string]*elemental.RelationshipInfo{
			"root": {},
		},
		Delete: map[string]*elemental.RelationshipInfo{
			"root": {},
		},
		Retrieve: map[string]*elemental.RelationshipInfo{
			"root": {},
		},
	}

	relationshipsRegistry[TextIndexIdentity] = &elemental.Relationship{}

	relationshipsRegistry[TokenIdentity] = &elemental.Relationship{
		Create: map[string]*elemental.RelationshipInfo{
			"root": {},
		},
	}

	relationshipsRegistry[TokenScopePolicyIdentity] = &elemental.Relationship{
		Create: map[string]*elemental.RelationshipInfo{
			"root": {},
		},
		Update: map[string]*elemental.RelationshipInfo{
			"root": {},
		},
		Patch: map[string]*elemental.RelationshipInfo{
			"root": {},
		},
		Delete: map[string]*elemental.RelationshipInfo{
			"root": {
				Parameters: []elemental.ParameterDefinition{
					{
						Name:     "q",
						Type:     "string",
						Multiple: true,
					},
				},
			},
		},
		Retrieve: map[string]*elemental.RelationshipInfo{
			"root": {},
		},
		RetrieveMany: map[string]*elemental.RelationshipInfo{
			"root": {
				Parameters: []elemental.ParameterDefinition{
					{
						Name:     "q",
						Type:     "string",
						Multiple: true,
					},
					{
						Name: "propagated",
						Type: "boolean",
					},
				},
			},
		},
		Info: map[string]*elemental.RelationshipInfo{
			"root": {
				Parameters: []elemental.ParameterDefinition{
					{
						Name:     "q",
						Type:     "string",
						Multiple: true,
					},
					{
						Name: "propagated",
						Type: "boolean",
					},
				},
			},
		},
	}

	relationshipsRegistry[TriggerIdentity] = &elemental.Relationship{
		Create: map[string]*elemental.RelationshipInfo{
			"automation": {},
		},
		RetrieveMany: map[string]*elemental.RelationshipInfo{
			"automation": {},
		},
		Info: map[string]*elemental.RelationshipInfo{
			"automation": {},
		},
	}

	relationshipsRegistry[TrustedCAIdentity] = &elemental.Relationship{
		RetrieveMany: map[string]*elemental.RelationshipInfo{
			"enforcer": {
				Parameters: []elemental.ParameterDefinition{
					{
						Name:         "type",
						Type:         "enum",
						DefaultValue: "Any",
						AllowedChoices: []string{
							"Any",
							"X509",
							"SSH",
						},
					},
				},
			},
			"namespace": {
				Parameters: []elemental.ParameterDefinition{
					{
						Name:         "type",
						Type:         "enum",
						DefaultValue: "Any",
						AllowedChoices: []string{
							"Any",
							"X509",
							"SSH",
							"JWT",
						},
					},
				},
			},
			"root": {
				Parameters: []elemental.ParameterDefinition{
					{
						Name:         "type",
						Type:         "enum",
						DefaultValue: "Any",
						AllowedChoices: []string{
							"Any",
							"X509",
							"SSH",
							"JWT",
						},
					},
				},
			},
		},
		Info: map[string]*elemental.RelationshipInfo{
			"enforcer": {
				Parameters: []elemental.ParameterDefinition{
					{
						Name:         "type",
						Type:         "enum",
						DefaultValue: "Any",
						AllowedChoices: []string{
							"Any",
							"X509",
							"SSH",
						},
					},
				},
			},
			"namespace": {
				Parameters: []elemental.ParameterDefinition{
					{
						Name:         "type",
						Type:         "enum",
						DefaultValue: "Any",
						AllowedChoices: []string{
							"Any",
							"X509",
							"SSH",
							"JWT",
						},
					},
				},
			},
			"root": {
				Parameters: []elemental.ParameterDefinition{
					{
						Name:         "type",
						Type:         "enum",
						DefaultValue: "Any",
						AllowedChoices: []string{
							"Any",
							"X509",
							"SSH",
							"JWT",
						},
					},
				},
			},
		},
	}

	relationshipsRegistry[TrustedNamespaceIdentity] = &elemental.Relationship{
		Create: map[string]*elemental.RelationshipInfo{
			"root": {},
		},
		Update: map[string]*elemental.RelationshipInfo{
			"root": {},
		},
		Patch: map[string]*elemental.RelationshipInfo{
			"root": {},
		},
		Delete: map[string]*elemental.RelationshipInfo{
			"root": {},
		},
		Retrieve: map[string]*elemental.RelationshipInfo{
			"root": {},
		},
		RetrieveMany: map[string]*elemental.RelationshipInfo{
			"root": {
				Parameters: []elemental.ParameterDefinition{
					{
						Name:     "q",
						Type:     "string",
						Multiple: true,
					},
					{
						Name: "propagated",
						Type: "boolean",
					},
				},
			},
		},
		Info: map[string]*elemental.RelationshipInfo{
			"root": {
				Parameters: []elemental.ParameterDefinition{
					{
						Name:     "q",
						Type:     "string",
						Multiple: true,
					},
					{
						Name: "propagated",
						Type: "boolean",
					},
				},
			},
		},
	}

	relationshipsRegistry[UserAccessPolicyIdentity] = &elemental.Relationship{
		Create: map[string]*elemental.RelationshipInfo{
			"root": {},
		},
		Update: map[string]*elemental.RelationshipInfo{
			"root": {},
		},
		Patch: map[string]*elemental.RelationshipInfo{
			"root": {},
		},
		Delete: map[string]*elemental.RelationshipInfo{
			"root": {
				Parameters: []elemental.ParameterDefinition{
					{
						Name:     "q",
						Type:     "string",
						Multiple: true,
					},
				},
			},
		},
		Retrieve: map[string]*elemental.RelationshipInfo{
			"root": {
				Parameters: []elemental.ParameterDefinition{
					{
						Name: "propagated",
						Type: "boolean",
					},
				},
			},
		},
		RetrieveMany: map[string]*elemental.RelationshipInfo{
			"root": {
				Parameters: []elemental.ParameterDefinition{
					{
						Name:     "q",
						Type:     "string",
						Multiple: true,
					},
					{
						Name: "propagated",
						Type: "boolean",
					},
				},
			},
		},
		Info: map[string]*elemental.RelationshipInfo{
			"root": {
				Parameters: []elemental.ParameterDefinition{
					{
						Name:     "q",
						Type:     "string",
						Multiple: true,
					},
					{
						Name: "propagated",
						Type: "boolean",
					},
				},
			},
		},
	}

	relationshipsRegistry[ValidateUIParameterIdentity] = &elemental.Relationship{
		Create: map[string]*elemental.RelationshipInfo{
			"root": {},
		},
	}

	relationshipsRegistry[VulnerabilityIdentity] = &elemental.Relationship{
		Create: map[string]*elemental.RelationshipInfo{
			"root": {},
		},
		Retrieve: map[string]*elemental.RelationshipInfo{
			"root": {},
		},
		RetrieveMany: map[string]*elemental.RelationshipInfo{
			"processingunit": {
				Parameters: []elemental.ParameterDefinition{
					{
						Name: "propagated",
						Type: "boolean",
					},
				},
			},
			"root": {
				Parameters: []elemental.ParameterDefinition{
					{
						Name:     "q",
						Type:     "string",
						Multiple: true,
					},
					{
						Name: "propagated",
						Type: "boolean",
					},
				},
			},
		},
		Info: map[string]*elemental.RelationshipInfo{
			"processingunit": {
				Parameters: []elemental.ParameterDefinition{
					{
						Name: "propagated",
						Type: "boolean",
					},
				},
			},
			"root": {
				Parameters: []elemental.ParameterDefinition{
					{
						Name:     "q",
						Type:     "string",
						Multiple: true,
					},
					{
						Name: "propagated",
						Type: "boolean",
					},
				},
			},
		},
	}

	relationshipsRegistry[X509CertificateIdentity] = &elemental.Relationship{
		Create: map[string]*elemental.RelationshipInfo{
			"root": {},
		},
	}

	relationshipsRegistry[X509CertificateCheckIdentity] = &elemental.Relationship{
		Retrieve: map[string]*elemental.RelationshipInfo{
			"root": {},
		},
		RetrieveMany: map[string]*elemental.RelationshipInfo{
			"root": {
				Parameters: []elemental.ParameterDefinition{
					{
						Name:     "q",
						Type:     "string",
						Multiple: true,
					},
				},
			},
		},
		Info: map[string]*elemental.RelationshipInfo{
			"root": {
				Parameters: []elemental.ParameterDefinition{
					{
						Name:     "q",
						Type:     "string",
						Multiple: true,
					},
				},
			},
		},
	}

}
