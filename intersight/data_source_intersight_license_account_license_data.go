package intersight

import (
	"context"
	"encoding/json"
	"log"
	"reflect"
	"time"

	models "github.com/CiscoDevNet/terraform-provider-intersight/intersight_gosdk"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceLicenseAccountLicenseData() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceLicenseAccountLicenseDataRead,
		Schema: map[string]*schema.Schema{
			"account": {
				Description: "A reference to a iamAccount resource.\nWhen the $expand query parameter is specified, the referenced resource is returned inline.",
				Type:        schema.TypeList,
				MaxItems:    1,
				Optional:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"additional_properties": {
							Type:             schema.TypeString,
							Optional:         true,
							DiffSuppressFunc: SuppressDiffAdditionProps,
						},
						"class_id": {
							Description: "The fully-qualified name of the instantiated, concrete type.\nThis property is used as a discriminator to identify the type of the payload\nwhen marshaling and unmarshaling data.",
							Type:        schema.TypeString,
							Optional:    true,
						},
						"moid": {
							Description: "The Moid of the referenced REST resource.",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
						},
						"object_type": {
							Description: "The fully-qualified name of the remote type referred by this relationship.",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
						},
						"selector": {
							Description: "An OData $filter expression which describes the REST resource to be referenced. This field may\nbe set instead of 'moid' by clients.\n1. If 'moid' is set this field is ignored.\n1. If 'selector' is set and 'moid' is empty/absent from the request, Intersight determines the Moid of the\nresource matching the filter expression and populates it in the MoRef that is part of the object\ninstance being inserted/updated to fulfill the REST request.\nAn error is returned if the filter matches zero or more than one REST resource.\nAn example filter string is: Serial eq '3AA8B7T11'.",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
						},
					},
				},
				Computed: true,
			},
			"account_id": {
				Description: "Root user's ID of the account.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"additional_properties": {
				Type:             schema.TypeString,
				Optional:         true,
				DiffSuppressFunc: SuppressDiffAdditionProps,
			},
			"agent_data": {
				Description: "Agent trusted store data.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"auth_expire_time": {
				Description: "Authorization expiration time.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"auth_initial_time": {
				Description: "Intial authorization time.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"auth_next_time": {
				Description: "Next time for the authorization.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"category": {
				Description: "Account license data category name.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"class_id": {
				Description: "The fully-qualified name of the instantiated, concrete type.\nThis property is used as a discriminator to identify the type of the payload\nwhen marshaling and unmarshaling data.\nThe enum values provides the list of concrete types that can be instantiated from this abstract type.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"customer_op": {
				Description: "A reference to a licenseCustomerOp resource.\nWhen the $expand query parameter is specified, the referenced resource is returned inline.",
				Type:        schema.TypeList,
				MaxItems:    1,
				Optional:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"additional_properties": {
							Type:             schema.TypeString,
							Optional:         true,
							DiffSuppressFunc: SuppressDiffAdditionProps,
						},
						"class_id": {
							Description: "The fully-qualified name of the instantiated, concrete type.\nThis property is used as a discriminator to identify the type of the payload\nwhen marshaling and unmarshaling data.",
							Type:        schema.TypeString,
							Optional:    true,
						},
						"moid": {
							Description: "The Moid of the referenced REST resource.",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
						},
						"object_type": {
							Description: "The fully-qualified name of the remote type referred by this relationship.",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
						},
						"selector": {
							Description: "An OData $filter expression which describes the REST resource to be referenced. This field may\nbe set instead of 'moid' by clients.\n1. If 'moid' is set this field is ignored.\n1. If 'selector' is set and 'moid' is empty/absent from the request, Intersight determines the Moid of the\nresource matching the filter expression and populates it in the MoRef that is part of the object\ninstance being inserted/updated to fulfill the REST request.\nAn error is returned if the filter matches zero or more than one REST resource.\nAn example filter string is: Serial eq '3AA8B7T11'.",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
						},
					},
				},
				Computed: true,
			},
			"default_license_type": {
				Description: "Default license tier set by user.\n* `Base` - Base as a License type. It is default license type.\n* `Essential` - Essential as a License type.\n* `Standard` - Standard as a License type.\n* `Advantage` - Advantage as a License type.\n* `Premier` - Premier as a License type.\n* `IWO-Essential` - IWO-Essential as a License type.\n* `IWO-Advantage` - IWO-Advantage as a License type.\n* `IWO-Premier` - IWO-Premier as a License type.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"error_desc": {
				Description: "The detailed error message when there is any error related to license sync of this account.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"group": {
				Description: "Account license data group name.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"highest_compliant_license_tier": {
				Description: "The highest license tier which is in compliant of this account.\n* `Base` - Base as a License type. It is default license type.\n* `Essential` - Essential as a License type.\n* `Standard` - Standard as a License type.\n* `Advantage` - Advantage as a License type.\n* `Premier` - Premier as a License type.\n* `IWO-Essential` - IWO-Essential as a License type.\n* `IWO-Advantage` - IWO-Advantage as a License type.\n* `IWO-Premier` - IWO-Premier as a License type.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"iwo_customer_op": {
				Description: "A reference to a licenseIwoCustomerOp resource.\nWhen the $expand query parameter is specified, the referenced resource is returned inline.",
				Type:        schema.TypeList,
				MaxItems:    1,
				Optional:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"additional_properties": {
							Type:             schema.TypeString,
							Optional:         true,
							DiffSuppressFunc: SuppressDiffAdditionProps,
						},
						"class_id": {
							Description: "The fully-qualified name of the instantiated, concrete type.\nThis property is used as a discriminator to identify the type of the payload\nwhen marshaling and unmarshaling data.",
							Type:        schema.TypeString,
							Optional:    true,
						},
						"moid": {
							Description: "The Moid of the referenced REST resource.",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
						},
						"object_type": {
							Description: "The fully-qualified name of the remote type referred by this relationship.",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
						},
						"selector": {
							Description: "An OData $filter expression which describes the REST resource to be referenced. This field may\nbe set instead of 'moid' by clients.\n1. If 'moid' is set this field is ignored.\n1. If 'selector' is set and 'moid' is empty/absent from the request, Intersight determines the Moid of the\nresource matching the filter expression and populates it in the MoRef that is part of the object\ninstance being inserted/updated to fulfill the REST request.\nAn error is returned if the filter matches zero or more than one REST resource.\nAn example filter string is: Serial eq '3AA8B7T11'.",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
						},
					},
				},
				Computed: true,
			},
			"iwo_license_count": {
				Description: "A reference to a licenseIwoLicenseCount resource.\nWhen the $expand query parameter is specified, the referenced resource is returned inline.",
				Type:        schema.TypeList,
				MaxItems:    1,
				Optional:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"additional_properties": {
							Type:             schema.TypeString,
							Optional:         true,
							DiffSuppressFunc: SuppressDiffAdditionProps,
						},
						"class_id": {
							Description: "The fully-qualified name of the instantiated, concrete type.\nThis property is used as a discriminator to identify the type of the payload\nwhen marshaling and unmarshaling data.",
							Type:        schema.TypeString,
							Optional:    true,
						},
						"moid": {
							Description: "The Moid of the referenced REST resource.",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
						},
						"object_type": {
							Description: "The fully-qualified name of the remote type referred by this relationship.",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
						},
						"selector": {
							Description: "An OData $filter expression which describes the REST resource to be referenced. This field may\nbe set instead of 'moid' by clients.\n1. If 'moid' is set this field is ignored.\n1. If 'selector' is set and 'moid' is empty/absent from the request, Intersight determines the Moid of the\nresource matching the filter expression and populates it in the MoRef that is part of the object\ninstance being inserted/updated to fulfill the REST request.\nAn error is returned if the filter matches zero or more than one REST resource.\nAn example filter string is: Serial eq '3AA8B7T11'.",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
						},
					},
				},
				Computed: true,
			},
			"last_sync": {
				Description: "Specifies last sync time with SA.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"last_updated_time": {
				Description: "Record's last update datetime.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"license_state": {
				Description: "Aggregrated mode for the agent.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"license_tech_support_info": {
				Description: "Tech-support info of a smart-agent.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"licenseinfos": {
				Description: "An array of relationships to licenseLicenseInfo resources.",
				Type:        schema.TypeList,
				Optional:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"additional_properties": {
							Type:             schema.TypeString,
							Optional:         true,
							DiffSuppressFunc: SuppressDiffAdditionProps,
						},
						"class_id": {
							Description: "The fully-qualified name of the instantiated, concrete type.\nThis property is used as a discriminator to identify the type of the payload\nwhen marshaling and unmarshaling data.",
							Type:        schema.TypeString,
							Optional:    true,
						},
						"moid": {
							Description: "The Moid of the referenced REST resource.",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
						},
						"object_type": {
							Description: "The fully-qualified name of the remote type referred by this relationship.",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
						},
						"selector": {
							Description: "An OData $filter expression which describes the REST resource to be referenced. This field may\nbe set instead of 'moid' by clients.\n1. If 'moid' is set this field is ignored.\n1. If 'selector' is set and 'moid' is empty/absent from the request, Intersight determines the Moid of the\nresource matching the filter expression and populates it in the MoRef that is part of the object\ninstance being inserted/updated to fulfill the REST request.\nAn error is returned if the filter matches zero or more than one REST resource.\nAn example filter string is: Serial eq '3AA8B7T11'.",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
						},
					},
				},
				Computed: true,
			},
			"moid": {
				Description: "The unique identifier of this Managed Object instance.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"object_type": {
				Description: "The fully-qualified name of the instantiated, concrete type.\nThe value should be the same as the 'ClassId' property.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"register_expire_time": {
				Description: "Registration exipiration time.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"register_initial_time": {
				Description: "Initial time of registration.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"register_next_time": {
				Description: "Next time for the license registration.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"registration_status": {
				Description: "Registration status of a smart-agent.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"renew_failure_string": {
				Description: "License renewal failure message.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"smart_account": {
				Description: "Name of the smart account.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"smartlicense_token": {
				Description: "A reference to a licenseSmartlicenseToken resource.\nWhen the $expand query parameter is specified, the referenced resource is returned inline.",
				Type:        schema.TypeList,
				MaxItems:    1,
				Optional:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"additional_properties": {
							Type:             schema.TypeString,
							Optional:         true,
							DiffSuppressFunc: SuppressDiffAdditionProps,
						},
						"class_id": {
							Description: "The fully-qualified name of the instantiated, concrete type.\nThis property is used as a discriminator to identify the type of the payload\nwhen marshaling and unmarshaling data.",
							Type:        schema.TypeString,
							Optional:    true,
						},
						"moid": {
							Description: "The Moid of the referenced REST resource.",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
						},
						"object_type": {
							Description: "The fully-qualified name of the remote type referred by this relationship.",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
						},
						"selector": {
							Description: "An OData $filter expression which describes the REST resource to be referenced. This field may\nbe set instead of 'moid' by clients.\n1. If 'moid' is set this field is ignored.\n1. If 'selector' is set and 'moid' is empty/absent from the request, Intersight determines the Moid of the\nresource matching the filter expression and populates it in the MoRef that is part of the object\ninstance being inserted/updated to fulfill the REST request.\nAn error is returned if the filter matches zero or more than one REST resource.\nAn example filter string is: Serial eq '3AA8B7T11'.",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
						},
					},
				},
				Computed: true,
			},
			"sync_status": {
				Description: "Current sync status for the account.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"tags": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"additional_properties": {
							Type:             schema.TypeString,
							Optional:         true,
							DiffSuppressFunc: SuppressDiffAdditionProps,
						},
						"key": {
							Description: "The string representation of a tag key.",
							Type:        schema.TypeString,
							Optional:    true,
						},
						"value": {
							Description: "The string representation of a tag value.",
							Type:        schema.TypeString,
							Optional:    true,
						},
					},
				},
			},
			"virtual_account": {
				Description: "Name of the virtual account.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
		},
	}
}

func dataSourceLicenseAccountLicenseDataRead(c context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Printf("%v", meta)
	conn := meta.(*Config)
	var de diag.Diagnostics
	var o = &models.LicenseAccountLicenseData{}
	if v, ok := d.GetOk("account_id"); ok {
		x := (v.(string))
		o.SetAccountId(x)
	}
	if v, ok := d.GetOk("agent_data"); ok {
		x := (v.(string))
		o.SetAgentData(x)
	}
	if v, ok := d.GetOk("auth_expire_time"); ok {
		x := (v.(string))
		o.SetAuthExpireTime(x)
	}
	if v, ok := d.GetOk("auth_initial_time"); ok {
		x := (v.(string))
		o.SetAuthInitialTime(x)
	}
	if v, ok := d.GetOk("auth_next_time"); ok {
		x := (v.(string))
		o.SetAuthNextTime(x)
	}
	if v, ok := d.GetOk("category"); ok {
		x := (v.(string))
		o.SetCategory(x)
	}
	if v, ok := d.GetOk("class_id"); ok {
		x := (v.(string))
		o.SetClassId(x)
	}
	if v, ok := d.GetOk("default_license_type"); ok {
		x := (v.(string))
		o.SetDefaultLicenseType(x)
	}
	if v, ok := d.GetOk("error_desc"); ok {
		x := (v.(string))
		o.SetErrorDesc(x)
	}
	if v, ok := d.GetOk("group"); ok {
		x := (v.(string))
		o.SetGroup(x)
	}
	if v, ok := d.GetOk("highest_compliant_license_tier"); ok {
		x := (v.(string))
		o.SetHighestCompliantLicenseTier(x)
	}
	if v, ok := d.GetOk("last_sync"); ok {
		x, _ := time.Parse(v.(string), time.RFC1123)
		o.SetLastSync(x)
	}
	if v, ok := d.GetOk("last_updated_time"); ok {
		x, _ := time.Parse(v.(string), time.RFC1123)
		o.SetLastUpdatedTime(x)
	}
	if v, ok := d.GetOk("license_state"); ok {
		x := (v.(string))
		o.SetLicenseState(x)
	}
	if v, ok := d.GetOk("license_tech_support_info"); ok {
		x := (v.(string))
		o.SetLicenseTechSupportInfo(x)
	}
	if v, ok := d.GetOk("moid"); ok {
		x := (v.(string))
		o.SetMoid(x)
	}
	if v, ok := d.GetOk("object_type"); ok {
		x := (v.(string))
		o.SetObjectType(x)
	}
	if v, ok := d.GetOk("register_expire_time"); ok {
		x := (v.(string))
		o.SetRegisterExpireTime(x)
	}
	if v, ok := d.GetOk("register_initial_time"); ok {
		x := (v.(string))
		o.SetRegisterInitialTime(x)
	}
	if v, ok := d.GetOk("register_next_time"); ok {
		x := (v.(string))
		o.SetRegisterNextTime(x)
	}
	if v, ok := d.GetOk("registration_status"); ok {
		x := (v.(string))
		o.SetRegistrationStatus(x)
	}
	if v, ok := d.GetOk("renew_failure_string"); ok {
		x := (v.(string))
		o.SetRenewFailureString(x)
	}
	if v, ok := d.GetOk("smart_account"); ok {
		x := (v.(string))
		o.SetSmartAccount(x)
	}
	if v, ok := d.GetOk("sync_status"); ok {
		x := (v.(string))
		o.SetSyncStatus(x)
	}
	if v, ok := d.GetOk("virtual_account"); ok {
		x := (v.(string))
		o.SetVirtualAccount(x)
	}

	data, err := o.MarshalJSON()
	if err != nil {
		return diag.Errorf("json marshal of LicenseAccountLicenseData object failed with error : %s", err.Error())
	}
	resMo, _, responseErr := conn.ApiClient.LicenseApi.GetLicenseAccountLicenseDataList(conn.ctx).Filter(getRequestParams(data)).Execute()
	if responseErr.Error() != "" {
		return diag.Errorf("error occurred while fetching LicenseAccountLicenseData: %s Response from endpoint: %s", responseErr.Error(), string(responseErr.Body()))
	}

	x, err := resMo.MarshalJSON()
	if err != nil {
		return diag.Errorf("error occurred while marshalling response for LicenseAccountLicenseData list: %s", err.Error())
	}
	var s = &models.LicenseAccountLicenseDataList{}
	err = json.Unmarshal(x, s)
	if err != nil {
		return diag.Errorf("error occurred while unmarshalling response to LicenseAccountLicenseData list: %s", err.Error())
	}
	result := s.GetResults()
	if result == nil {
		return diag.Errorf("your query for LicenseAccountLicenseData did not return results. Please change your search criteria and try again")
	}
	switch reflect.TypeOf(result).Kind() {
	case reflect.Slice:
		r := reflect.ValueOf(result)
		for i := 0; i < r.Len(); i++ {
			var s = &models.LicenseAccountLicenseData{}
			oo, _ := json.Marshal(r.Index(i).Interface())
			if err = json.Unmarshal(oo, s); err != nil {
				return diag.Errorf("error occurred while unmarshalling result at index %+v: %s", i, err.Error())
			}

			if err := d.Set("account", flattenMapIamAccountRelationship(s.GetAccount(), d)); err != nil {
				return diag.Errorf("error occurred while setting property Account: %s", err.Error())
			}
			if err := d.Set("account_id", (s.GetAccountId())); err != nil {
				return diag.Errorf("error occurred while setting property AccountId: %s", err.Error())
			}
			if err := d.Set("additional_properties", flattenAdditionalProperties(s.AdditionalProperties)); err != nil {
				return diag.Errorf("error occurred while setting property AdditionalProperties: %s", err.Error())
			}
			if err := d.Set("agent_data", (s.GetAgentData())); err != nil {
				return diag.Errorf("error occurred while setting property AgentData: %s", err.Error())
			}
			if err := d.Set("auth_expire_time", (s.GetAuthExpireTime())); err != nil {
				return diag.Errorf("error occurred while setting property AuthExpireTime: %s", err.Error())
			}
			if err := d.Set("auth_initial_time", (s.GetAuthInitialTime())); err != nil {
				return diag.Errorf("error occurred while setting property AuthInitialTime: %s", err.Error())
			}
			if err := d.Set("auth_next_time", (s.GetAuthNextTime())); err != nil {
				return diag.Errorf("error occurred while setting property AuthNextTime: %s", err.Error())
			}
			if err := d.Set("category", (s.GetCategory())); err != nil {
				return diag.Errorf("error occurred while setting property Category: %s", err.Error())
			}
			if err := d.Set("class_id", (s.GetClassId())); err != nil {
				return diag.Errorf("error occurred while setting property ClassId: %s", err.Error())
			}

			if err := d.Set("customer_op", flattenMapLicenseCustomerOpRelationship(s.GetCustomerOp(), d)); err != nil {
				return diag.Errorf("error occurred while setting property CustomerOp: %s", err.Error())
			}
			if err := d.Set("default_license_type", (s.GetDefaultLicenseType())); err != nil {
				return diag.Errorf("error occurred while setting property DefaultLicenseType: %s", err.Error())
			}
			if err := d.Set("error_desc", (s.GetErrorDesc())); err != nil {
				return diag.Errorf("error occurred while setting property ErrorDesc: %s", err.Error())
			}
			if err := d.Set("group", (s.GetGroup())); err != nil {
				return diag.Errorf("error occurred while setting property Group: %s", err.Error())
			}
			if err := d.Set("highest_compliant_license_tier", (s.GetHighestCompliantLicenseTier())); err != nil {
				return diag.Errorf("error occurred while setting property HighestCompliantLicenseTier: %s", err.Error())
			}

			if err := d.Set("iwo_customer_op", flattenMapLicenseIwoCustomerOpRelationship(s.GetIwoCustomerOp(), d)); err != nil {
				return diag.Errorf("error occurred while setting property IwoCustomerOp: %s", err.Error())
			}

			if err := d.Set("iwo_license_count", flattenMapLicenseIwoLicenseCountRelationship(s.GetIwoLicenseCount(), d)); err != nil {
				return diag.Errorf("error occurred while setting property IwoLicenseCount: %s", err.Error())
			}

			if err := d.Set("last_sync", (s.GetLastSync()).String()); err != nil {
				return diag.Errorf("error occurred while setting property LastSync: %s", err.Error())
			}

			if err := d.Set("last_updated_time", (s.GetLastUpdatedTime()).String()); err != nil {
				return diag.Errorf("error occurred while setting property LastUpdatedTime: %s", err.Error())
			}
			if err := d.Set("license_state", (s.GetLicenseState())); err != nil {
				return diag.Errorf("error occurred while setting property LicenseState: %s", err.Error())
			}
			if err := d.Set("license_tech_support_info", (s.GetLicenseTechSupportInfo())); err != nil {
				return diag.Errorf("error occurred while setting property LicenseTechSupportInfo: %s", err.Error())
			}

			if err := d.Set("licenseinfos", flattenListLicenseLicenseInfoRelationship(s.GetLicenseinfos(), d)); err != nil {
				return diag.Errorf("error occurred while setting property Licenseinfos: %s", err.Error())
			}
			if err := d.Set("moid", (s.GetMoid())); err != nil {
				return diag.Errorf("error occurred while setting property Moid: %s", err.Error())
			}
			if err := d.Set("object_type", (s.GetObjectType())); err != nil {
				return diag.Errorf("error occurred while setting property ObjectType: %s", err.Error())
			}
			if err := d.Set("register_expire_time", (s.GetRegisterExpireTime())); err != nil {
				return diag.Errorf("error occurred while setting property RegisterExpireTime: %s", err.Error())
			}
			if err := d.Set("register_initial_time", (s.GetRegisterInitialTime())); err != nil {
				return diag.Errorf("error occurred while setting property RegisterInitialTime: %s", err.Error())
			}
			if err := d.Set("register_next_time", (s.GetRegisterNextTime())); err != nil {
				return diag.Errorf("error occurred while setting property RegisterNextTime: %s", err.Error())
			}
			if err := d.Set("registration_status", (s.GetRegistrationStatus())); err != nil {
				return diag.Errorf("error occurred while setting property RegistrationStatus: %s", err.Error())
			}
			if err := d.Set("renew_failure_string", (s.GetRenewFailureString())); err != nil {
				return diag.Errorf("error occurred while setting property RenewFailureString: %s", err.Error())
			}
			if err := d.Set("smart_account", (s.GetSmartAccount())); err != nil {
				return diag.Errorf("error occurred while setting property SmartAccount: %s", err.Error())
			}

			if err := d.Set("smartlicense_token", flattenMapLicenseSmartlicenseTokenRelationship(s.GetSmartlicenseToken(), d)); err != nil {
				return diag.Errorf("error occurred while setting property SmartlicenseToken: %s", err.Error())
			}
			if err := d.Set("sync_status", (s.GetSyncStatus())); err != nil {
				return diag.Errorf("error occurred while setting property SyncStatus: %s", err.Error())
			}

			if err := d.Set("tags", flattenListMoTag(s.GetTags(), d)); err != nil {
				return diag.Errorf("error occurred while setting property Tags: %s", err.Error())
			}
			if err := d.Set("virtual_account", (s.GetVirtualAccount())); err != nil {
				return diag.Errorf("error occurred while setting property VirtualAccount: %s", err.Error())
			}
			d.SetId(s.GetMoid())
		}
	}
	return de
}
