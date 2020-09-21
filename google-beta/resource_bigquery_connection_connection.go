// ----------------------------------------------------------------------------
//
//     ***     AUTO GENERATED CODE    ***    AUTO GENERATED CODE     ***
//
// ----------------------------------------------------------------------------
//
//     This file is automatically generated by Magic Modules and manual
//     changes will be clobbered when the file is regenerated.
//
//     Please read more about how to change this file in
//     .github/CONTRIBUTING.md.
//
// ----------------------------------------------------------------------------

package google

import (
	"fmt"
	"log"
	"reflect"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceBigqueryConnectionConnection() *schema.Resource {
	return &schema.Resource{
		Create: resourceBigqueryConnectionConnectionCreate,
		Read:   resourceBigqueryConnectionConnectionRead,
		Update: resourceBigqueryConnectionConnectionUpdate,
		Delete: resourceBigqueryConnectionConnectionDelete,

		Importer: &schema.ResourceImporter{
			State: resourceBigqueryConnectionConnectionImport,
		},

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(4 * time.Minute),
			Update: schema.DefaultTimeout(4 * time.Minute),
			Delete: schema.DefaultTimeout(4 * time.Minute),
		},

		Schema: map[string]*schema.Schema{
			"cloud_sql": {
				Type:        schema.TypeList,
				Required:    true,
				Description: `Cloud SQL properties.`,
				MaxItems:    1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"credential": {
							Type:        schema.TypeList,
							Required:    true,
							Description: `Cloud SQL properties.`,
							MaxItems:    1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"password": {
										Type:        schema.TypeString,
										Required:    true,
										Description: `Password for database.`,
										Sensitive:   true,
									},
									"username": {
										Type:        schema.TypeString,
										Required:    true,
										Description: `Username for database.`,
									},
								},
							},
						},
						"database": {
							Type:        schema.TypeString,
							Required:    true,
							Description: `Database name.`,
						},
						"instance_id": {
							Type:        schema.TypeString,
							Required:    true,
							Description: `Cloud SQL instance ID in the form project:location:instance.`,
						},
						"type": {
							Type:         schema.TypeString,
							Required:     true,
							ValidateFunc: validation.StringInSlice([]string{"DATABASE_TYPE_UNSPECIFIED", "POSTGRES", "MYSQL"}, false),
							Description:  `Type of the Cloud SQL database. Possible values: ["DATABASE_TYPE_UNSPECIFIED", "POSTGRES", "MYSQL"]`,
						},
					},
				},
			},
			"connection_id": {
				Type:        schema.TypeString,
				Optional:    true,
				ForceNew:    true,
				Description: `Optional connection id that should be assigned to the created connection.`,
			},
			"description": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: `A descriptive description for the connection`,
			},
			"friendly_name": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: `A descriptive name for the connection`,
			},
			"location": {
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
				Description: `The geographic location where the connection should reside.
Cloud SQL instance must be in the same location as the connection
with following exceptions: Cloud SQL us-central1 maps to BigQuery US, Cloud SQL europe-west1 maps to BigQuery EU.
Examples: US, EU, asia-northeast1, us-central1, europe-west1. The default value is US.`,
				Default: "US",
			},
			"has_credential": {
				Type:        schema.TypeBool,
				Computed:    true,
				Description: `True if the connection has credential assigned.`,
			},
			"name": {
				Type:     schema.TypeString,
				Computed: true,
				Description: `The resource name of the connection in the form of: 
"projects/{project_id}/locations/{location_id}/connections/{connectionId}"`,
			},
			"project": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
		},
	}
}

func resourceBigqueryConnectionConnectionCreate(d *schema.ResourceData, meta interface{}) error {
	var m providerMeta

	err := d.GetProviderMeta(&m)
	if err != nil {
		return err
	}

	config := meta.(*Config)
	config.userAgent = fmt.Sprintf("%s %s", config.userAgent, m.ModuleName)

	obj := make(map[string]interface{})
	friendlyNameProp, err := expandBigqueryConnectionConnectionFriendlyName(d.Get("friendly_name"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("friendly_name"); !isEmptyValue(reflect.ValueOf(friendlyNameProp)) && (ok || !reflect.DeepEqual(v, friendlyNameProp)) {
		obj["friendlyName"] = friendlyNameProp
	}
	descriptionProp, err := expandBigqueryConnectionConnectionDescription(d.Get("description"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("description"); !isEmptyValue(reflect.ValueOf(descriptionProp)) && (ok || !reflect.DeepEqual(v, descriptionProp)) {
		obj["description"] = descriptionProp
	}
	cloudSqlProp, err := expandBigqueryConnectionConnectionCloudSql(d.Get("cloud_sql"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("cloud_sql"); !isEmptyValue(reflect.ValueOf(cloudSqlProp)) && (ok || !reflect.DeepEqual(v, cloudSqlProp)) {
		obj["cloudSql"] = cloudSqlProp
	}

	url, err := replaceVars(d, config, "{{BigqueryConnectionBasePath}}projects/{{project}}/locations/{{location}}/connections?connectionId={{connection_id}}")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Creating new Connection: %#v", obj)
	billingProject := ""

	project, err := getProject(d, config)
	if err != nil {
		return err
	}
	billingProject = project

	// err == nil indicates that the billing_project value was found
	if bp, err := getBillingProject(d, config); err == nil {
		billingProject = bp
	}

	res, err := sendRequestWithTimeout(config, "POST", billingProject, url, obj, d.Timeout(schema.TimeoutCreate))
	if err != nil {
		return fmt.Errorf("Error creating Connection: %s", err)
	}
	if err := d.Set("name", flattenBigqueryConnectionConnectionName(res["name"], d, config)); err != nil {
		return fmt.Errorf(`Error setting computed identity field "name": %s`, err)
	}

	// Store the ID now
	id, err := replaceVars(d, config, "{{name}}")
	if err != nil {
		return fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	log.Printf("[DEBUG] Finished creating Connection %q: %#v", d.Id(), res)

	return resourceBigqueryConnectionConnectionRead(d, meta)
}

func resourceBigqueryConnectionConnectionRead(d *schema.ResourceData, meta interface{}) error {
	var m providerMeta

	err := d.GetProviderMeta(&m)
	if err != nil {
		return err
	}

	config := meta.(*Config)
	config.userAgent = fmt.Sprintf("%s %s", config.userAgent, m.ModuleName)

	url, err := replaceVars(d, config, "{{BigqueryConnectionBasePath}}{{name}}")
	if err != nil {
		return err
	}

	billingProject := ""

	project, err := getProject(d, config)
	if err != nil {
		return err
	}
	billingProject = project

	// err == nil indicates that the billing_project value was found
	if bp, err := getBillingProject(d, config); err == nil {
		billingProject = bp
	}

	res, err := sendRequest(config, "GET", billingProject, url, nil)
	if err != nil {
		return handleNotFoundError(err, d, fmt.Sprintf("BigqueryConnectionConnection %q", d.Id()))
	}

	if err := d.Set("project", project); err != nil {
		return fmt.Errorf("Error reading Connection: %s", err)
	}

	if err := d.Set("name", flattenBigqueryConnectionConnectionName(res["name"], d, config)); err != nil {
		return fmt.Errorf("Error reading Connection: %s", err)
	}
	if err := d.Set("friendly_name", flattenBigqueryConnectionConnectionFriendlyName(res["friendlyName"], d, config)); err != nil {
		return fmt.Errorf("Error reading Connection: %s", err)
	}
	if err := d.Set("description", flattenBigqueryConnectionConnectionDescription(res["description"], d, config)); err != nil {
		return fmt.Errorf("Error reading Connection: %s", err)
	}
	if err := d.Set("has_credential", flattenBigqueryConnectionConnectionHasCredential(res["hasCredential"], d, config)); err != nil {
		return fmt.Errorf("Error reading Connection: %s", err)
	}
	if err := d.Set("cloud_sql", flattenBigqueryConnectionConnectionCloudSql(res["cloudSql"], d, config)); err != nil {
		return fmt.Errorf("Error reading Connection: %s", err)
	}

	return nil
}

func resourceBigqueryConnectionConnectionUpdate(d *schema.ResourceData, meta interface{}) error {
	var m providerMeta

	err := d.GetProviderMeta(&m)
	if err != nil {
		return err
	}

	config := meta.(*Config)
	config.userAgent = fmt.Sprintf("%s %s", config.userAgent, m.ModuleName)

	billingProject := ""

	project, err := getProject(d, config)
	if err != nil {
		return err
	}
	billingProject = project

	obj := make(map[string]interface{})
	friendlyNameProp, err := expandBigqueryConnectionConnectionFriendlyName(d.Get("friendly_name"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("friendly_name"); !isEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, friendlyNameProp)) {
		obj["friendlyName"] = friendlyNameProp
	}
	descriptionProp, err := expandBigqueryConnectionConnectionDescription(d.Get("description"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("description"); !isEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, descriptionProp)) {
		obj["description"] = descriptionProp
	}
	cloudSqlProp, err := expandBigqueryConnectionConnectionCloudSql(d.Get("cloud_sql"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("cloud_sql"); !isEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, cloudSqlProp)) {
		obj["cloudSql"] = cloudSqlProp
	}

	url, err := replaceVars(d, config, "{{BigqueryConnectionBasePath}}{{name}}")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Updating Connection %q: %#v", d.Id(), obj)
	updateMask := []string{}

	if d.HasChange("friendly_name") {
		updateMask = append(updateMask, "friendlyName")
	}

	if d.HasChange("description") {
		updateMask = append(updateMask, "description")
	}

	if d.HasChange("cloud_sql") {
		updateMask = append(updateMask, "cloudSql")
	}
	// updateMask is a URL parameter but not present in the schema, so replaceVars
	// won't set it
	url, err = addQueryParams(url, map[string]string{"updateMask": strings.Join(updateMask, ",")})
	if err != nil {
		return err
	}

	// err == nil indicates that the billing_project value was found
	if bp, err := getBillingProject(d, config); err == nil {
		billingProject = bp
	}

	res, err := sendRequestWithTimeout(config, "PATCH", billingProject, url, obj, d.Timeout(schema.TimeoutUpdate))

	if err != nil {
		return fmt.Errorf("Error updating Connection %q: %s", d.Id(), err)
	} else {
		log.Printf("[DEBUG] Finished updating Connection %q: %#v", d.Id(), res)
	}

	return resourceBigqueryConnectionConnectionRead(d, meta)
}

func resourceBigqueryConnectionConnectionDelete(d *schema.ResourceData, meta interface{}) error {
	var m providerMeta

	err := d.GetProviderMeta(&m)
	if err != nil {
		return err
	}

	config := meta.(*Config)
	config.userAgent = fmt.Sprintf("%s %s", config.userAgent, m.ModuleName)

	billingProject := ""

	project, err := getProject(d, config)
	if err != nil {
		return err
	}
	billingProject = project

	url, err := replaceVars(d, config, "{{BigqueryConnectionBasePath}}{{name}}")
	if err != nil {
		return err
	}

	var obj map[string]interface{}
	log.Printf("[DEBUG] Deleting Connection %q", d.Id())

	// err == nil indicates that the billing_project value was found
	if bp, err := getBillingProject(d, config); err == nil {
		billingProject = bp
	}

	res, err := sendRequestWithTimeout(config, "DELETE", billingProject, url, obj, d.Timeout(schema.TimeoutDelete))
	if err != nil {
		return handleNotFoundError(err, d, "Connection")
	}

	log.Printf("[DEBUG] Finished deleting Connection %q: %#v", d.Id(), res)
	return nil
}

func resourceBigqueryConnectionConnectionImport(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	config := meta.(*Config)
	if err := parseImportId([]string{
		"(?P<name>[^/]+)",
	}, d, config); err != nil {
		return nil, err
	}

	// Replace import id for the resource id
	id, err := replaceVars(d, config, "{{name}}")
	if err != nil {
		return nil, fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	return []*schema.ResourceData{d}, nil
}

func flattenBigqueryConnectionConnectionName(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func flattenBigqueryConnectionConnectionFriendlyName(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func flattenBigqueryConnectionConnectionDescription(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func flattenBigqueryConnectionConnectionHasCredential(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func flattenBigqueryConnectionConnectionCloudSql(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	if v == nil {
		return nil
	}
	original := v.(map[string]interface{})
	if len(original) == 0 {
		return nil
	}
	transformed := make(map[string]interface{})
	transformed["instance_id"] =
		flattenBigqueryConnectionConnectionCloudSqlInstanceId(original["instanceId"], d, config)
	transformed["database"] =
		flattenBigqueryConnectionConnectionCloudSqlDatabase(original["database"], d, config)
	transformed["credential"] =
		flattenBigqueryConnectionConnectionCloudSqlCredential(original["credential"], d, config)
	transformed["type"] =
		flattenBigqueryConnectionConnectionCloudSqlType(original["type"], d, config)
	return []interface{}{transformed}
}
func flattenBigqueryConnectionConnectionCloudSqlInstanceId(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func flattenBigqueryConnectionConnectionCloudSqlDatabase(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func flattenBigqueryConnectionConnectionCloudSqlCredential(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return []interface{}{
		map[string]interface{}{
			"username": d.Get("cloud_sql.0.credential.0.username"),
			"password": d.Get("cloud_sql.0.credential.0.password"),
		},
	}
}

func flattenBigqueryConnectionConnectionCloudSqlType(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func expandBigqueryConnectionConnectionFriendlyName(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandBigqueryConnectionConnectionDescription(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandBigqueryConnectionConnectionCloudSql(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedInstanceId, err := expandBigqueryConnectionConnectionCloudSqlInstanceId(original["instance_id"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedInstanceId); val.IsValid() && !isEmptyValue(val) {
		transformed["instanceId"] = transformedInstanceId
	}

	transformedDatabase, err := expandBigqueryConnectionConnectionCloudSqlDatabase(original["database"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedDatabase); val.IsValid() && !isEmptyValue(val) {
		transformed["database"] = transformedDatabase
	}

	transformedCredential, err := expandBigqueryConnectionConnectionCloudSqlCredential(original["credential"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedCredential); val.IsValid() && !isEmptyValue(val) {
		transformed["credential"] = transformedCredential
	}

	transformedType, err := expandBigqueryConnectionConnectionCloudSqlType(original["type"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedType); val.IsValid() && !isEmptyValue(val) {
		transformed["type"] = transformedType
	}

	return transformed, nil
}

func expandBigqueryConnectionConnectionCloudSqlInstanceId(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandBigqueryConnectionConnectionCloudSqlDatabase(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandBigqueryConnectionConnectionCloudSqlCredential(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedUsername, err := expandBigqueryConnectionConnectionCloudSqlCredentialUsername(original["username"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedUsername); val.IsValid() && !isEmptyValue(val) {
		transformed["username"] = transformedUsername
	}

	transformedPassword, err := expandBigqueryConnectionConnectionCloudSqlCredentialPassword(original["password"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedPassword); val.IsValid() && !isEmptyValue(val) {
		transformed["password"] = transformedPassword
	}

	return transformed, nil
}

func expandBigqueryConnectionConnectionCloudSqlCredentialUsername(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandBigqueryConnectionConnectionCloudSqlCredentialPassword(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandBigqueryConnectionConnectionCloudSqlType(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}
