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
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceComputeRegionNetworkEndpointGroup() *schema.Resource {
	return &schema.Resource{
		Create: resourceComputeRegionNetworkEndpointGroupCreate,
		Read:   resourceComputeRegionNetworkEndpointGroupRead,
		Delete: resourceComputeRegionNetworkEndpointGroupDelete,

		Importer: &schema.ResourceImporter{
			State: resourceComputeRegionNetworkEndpointGroupImport,
		},

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(4 * time.Minute),
			Delete: schema.DefaultTimeout(4 * time.Minute),
		},

		Schema: map[string]*schema.Schema{
			"name": {
				Type:         schema.TypeString,
				Required:     true,
				ForceNew:     true,
				ValidateFunc: validateGCPName,
				Description: `Name of the resource; provided by the client when the resource is
created. The name must be 1-63 characters long, and comply with
RFC1035. Specifically, the name must be 1-63 characters long and match
the regular expression '[a-z]([-a-z0-9]*[a-z0-9])?' which means the
first character must be a lowercase letter, and all following
characters must be a dash, lowercase letter, or digit, except the last
character, which cannot be a dash.`,
			},
			"region": {
				Type:             schema.TypeString,
				Required:         true,
				ForceNew:         true,
				DiffSuppressFunc: compareSelfLinkOrResourceName,
				Description:      `A reference to the region where the Serverless NEGs Reside.`,
			},
			"app_engine": {
				Type:     schema.TypeList,
				Optional: true,
				ForceNew: true,
				Description: `Only valid when networkEndpointType is "SERVERLESS".
Only one of cloud_run, app_engine or cloud_function may be set.`,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"service": {
							Type:     schema.TypeString,
							Optional: true,
							ForceNew: true,
							Description: `Optional serving service.
The service name must be 1-63 characters long, and comply with RFC1035.
Example value: "default", "my-service".`,
						},
						"url_mask": {
							Type:     schema.TypeString,
							Optional: true,
							ForceNew: true,
							Description: `A template to parse service and version fields from a request URL.
URL mask allows for routing to multiple App Engine services without
having to create multiple Network Endpoint Groups and backend services.

For example, the request URLs "foo1-dot-appname.appspot.com/v1" and
"foo1-dot-appname.appspot.com/v2" can be backed by the same Serverless NEG with
URL mask "-dot-appname.appspot.com/". The URL mask will parse
them to { service = "foo1", version = "v1" } and { service = "foo1", version = "v2" } respectively.`,
						},
						"version": {
							Type:     schema.TypeString,
							Optional: true,
							ForceNew: true,
							Description: `Optional serving version.
The version must be 1-63 characters long, and comply with RFC1035.
Example value: "v1", "v2".`,
						},
					},
				},
				ExactlyOneOf: []string{"app_engine", "cloud_function", "cloud_run"},
			},
			"cloud_function": {
				Type:     schema.TypeList,
				Optional: true,
				ForceNew: true,
				Description: `Only valid when networkEndpointType is "SERVERLESS".
Only one of cloud_run, app_engine or cloud_function may be set.`,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"function": {
							Type:     schema.TypeString,
							Optional: true,
							ForceNew: true,
							Description: `A user-defined name of the Cloud Function.
The function name is case-sensitive and must be 1-63 characters long.
Example value: "func1".`,
							AtLeastOneOf: []string{"cloud_function.0.function", "cloud_function.0.url_mask"},
						},
						"url_mask": {
							Type:     schema.TypeString,
							Optional: true,
							ForceNew: true,
							Description: `A template to parse function field from a request URL. URL mask allows
for routing to multiple Cloud Functions without having to create
multiple Network Endpoint Groups and backend services.

For example, request URLs "mydomain.com/function1" and "mydomain.com/function2"
can be backed by the same Serverless NEG with URL mask "/". The URL mask
will parse them to { function = "function1" } and { function = "function2" } respectively.`,
							AtLeastOneOf: []string{"cloud_function.0.function", "cloud_function.0.url_mask"},
						},
					},
				},
				ExactlyOneOf: []string{"app_engine", "cloud_function", "cloud_run"},
			},
			"cloud_run": {
				Type:     schema.TypeList,
				Optional: true,
				ForceNew: true,
				Description: `Only valid when networkEndpointType is "SERVERLESS".
Only one of cloud_run, app_engine or cloud_function may be set.`,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"service": {
							Type:     schema.TypeString,
							Optional: true,
							ForceNew: true,
							Description: `Cloud Run service is the main resource of Cloud Run.
The service must be 1-63 characters long, and comply with RFC1035.
Example value: "run-service".`,
							AtLeastOneOf: []string{"cloud_run.0.service", "cloud_run.0.url_mask"},
						},
						"tag": {
							Type:     schema.TypeString,
							Optional: true,
							ForceNew: true,
							Description: `Cloud Run tag represents the "named-revision" to provide
additional fine-grained traffic routing information.
The tag must be 1-63 characters long, and comply with RFC1035.
Example value: "revision-0010".`,
						},
						"url_mask": {
							Type:     schema.TypeString,
							Optional: true,
							ForceNew: true,
							Description: `A template to parse service and tag fields from a request URL. 
URL mask allows for routing to multiple Run services without having 
to create multiple network endpoint groups and backend services.

For example, request URLs "foo1.domain.com/bar1" and "foo1.domain.com/bar2" 
an be backed by the same Serverless Network Endpoint Group (NEG) with 
URL mask ".domain.com/". The URL mask will parse them to { service="bar1", tag="foo1" } 
and { service="bar2", tag="foo2" } respectively.`,
							AtLeastOneOf: []string{"cloud_run.0.service", "cloud_run.0.url_mask"},
						},
					},
				},
				ExactlyOneOf: []string{"cloud_run", "cloud_function", "app_engine"},
			},
			"description": {
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
				Description: `An optional description of this resource. Provide this property when
you create the resource.`,
			},
			"network_endpoint_type": {
				Type:         schema.TypeString,
				Optional:     true,
				ForceNew:     true,
				ValidateFunc: validation.StringInSlice([]string{"SERVERLESS", ""}, false),
				Description:  `Type of network endpoints in this network endpoint group. Defaults to SERVERLESS Default value: "SERVERLESS" Possible values: ["SERVERLESS"]`,
				Default:      "SERVERLESS",
			},
			"project": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"self_link": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func resourceComputeRegionNetworkEndpointGroupCreate(d *schema.ResourceData, meta interface{}) error {
	var m providerMeta

	err := d.GetProviderMeta(&m)
	if err != nil {
		return err
	}

	config := meta.(*Config)
	config.userAgent = fmt.Sprintf("%s %s", config.userAgent, m.ModuleName)

	obj := make(map[string]interface{})
	nameProp, err := expandComputeRegionNetworkEndpointGroupName(d.Get("name"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("name"); !isEmptyValue(reflect.ValueOf(nameProp)) && (ok || !reflect.DeepEqual(v, nameProp)) {
		obj["name"] = nameProp
	}
	descriptionProp, err := expandComputeRegionNetworkEndpointGroupDescription(d.Get("description"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("description"); !isEmptyValue(reflect.ValueOf(descriptionProp)) && (ok || !reflect.DeepEqual(v, descriptionProp)) {
		obj["description"] = descriptionProp
	}
	networkEndpointTypeProp, err := expandComputeRegionNetworkEndpointGroupNetworkEndpointType(d.Get("network_endpoint_type"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("network_endpoint_type"); !isEmptyValue(reflect.ValueOf(networkEndpointTypeProp)) && (ok || !reflect.DeepEqual(v, networkEndpointTypeProp)) {
		obj["networkEndpointType"] = networkEndpointTypeProp
	}
	cloudRunProp, err := expandComputeRegionNetworkEndpointGroupCloudRun(d.Get("cloud_run"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("cloud_run"); !isEmptyValue(reflect.ValueOf(cloudRunProp)) && (ok || !reflect.DeepEqual(v, cloudRunProp)) {
		obj["cloudRun"] = cloudRunProp
	}
	appEngineProp, err := expandComputeRegionNetworkEndpointGroupAppEngine(d.Get("app_engine"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("app_engine"); !isEmptyValue(reflect.ValueOf(appEngineProp)) && (ok || !reflect.DeepEqual(v, appEngineProp)) {
		obj["appEngine"] = appEngineProp
	}
	cloudFunctionProp, err := expandComputeRegionNetworkEndpointGroupCloudFunction(d.Get("cloud_function"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("cloud_function"); !isEmptyValue(reflect.ValueOf(cloudFunctionProp)) && (ok || !reflect.DeepEqual(v, cloudFunctionProp)) {
		obj["cloudFunction"] = cloudFunctionProp
	}
	regionProp, err := expandComputeRegionNetworkEndpointGroupRegion(d.Get("region"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("region"); !isEmptyValue(reflect.ValueOf(regionProp)) && (ok || !reflect.DeepEqual(v, regionProp)) {
		obj["region"] = regionProp
	}

	url, err := replaceVars(d, config, "{{ComputeBasePath}}projects/{{project}}/regions/{{region}}/networkEndpointGroups")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Creating new RegionNetworkEndpointGroup: %#v", obj)
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
		return fmt.Errorf("Error creating RegionNetworkEndpointGroup: %s", err)
	}

	// Store the ID now
	id, err := replaceVars(d, config, "projects/{{project}}/regions/{{region}}/networkEndpointGroups/{{name}}")
	if err != nil {
		return fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	err = computeOperationWaitTime(
		config, res, project, "Creating RegionNetworkEndpointGroup",
		d.Timeout(schema.TimeoutCreate))

	if err != nil {
		// The resource didn't actually create
		d.SetId("")
		return fmt.Errorf("Error waiting to create RegionNetworkEndpointGroup: %s", err)
	}

	log.Printf("[DEBUG] Finished creating RegionNetworkEndpointGroup %q: %#v", d.Id(), res)

	return resourceComputeRegionNetworkEndpointGroupRead(d, meta)
}

func resourceComputeRegionNetworkEndpointGroupRead(d *schema.ResourceData, meta interface{}) error {
	var m providerMeta

	err := d.GetProviderMeta(&m)
	if err != nil {
		return err
	}

	config := meta.(*Config)
	config.userAgent = fmt.Sprintf("%s %s", config.userAgent, m.ModuleName)

	url, err := replaceVars(d, config, "{{ComputeBasePath}}projects/{{project}}/regions/{{region}}/networkEndpointGroups/{{name}}")
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
		return handleNotFoundError(err, d, fmt.Sprintf("ComputeRegionNetworkEndpointGroup %q", d.Id()))
	}

	if err := d.Set("project", project); err != nil {
		return fmt.Errorf("Error reading RegionNetworkEndpointGroup: %s", err)
	}

	if err := d.Set("name", flattenComputeRegionNetworkEndpointGroupName(res["name"], d, config)); err != nil {
		return fmt.Errorf("Error reading RegionNetworkEndpointGroup: %s", err)
	}
	if err := d.Set("description", flattenComputeRegionNetworkEndpointGroupDescription(res["description"], d, config)); err != nil {
		return fmt.Errorf("Error reading RegionNetworkEndpointGroup: %s", err)
	}
	if err := d.Set("network_endpoint_type", flattenComputeRegionNetworkEndpointGroupNetworkEndpointType(res["networkEndpointType"], d, config)); err != nil {
		return fmt.Errorf("Error reading RegionNetworkEndpointGroup: %s", err)
	}
	if err := d.Set("cloud_run", flattenComputeRegionNetworkEndpointGroupCloudRun(res["cloudRun"], d, config)); err != nil {
		return fmt.Errorf("Error reading RegionNetworkEndpointGroup: %s", err)
	}
	if err := d.Set("app_engine", flattenComputeRegionNetworkEndpointGroupAppEngine(res["appEngine"], d, config)); err != nil {
		return fmt.Errorf("Error reading RegionNetworkEndpointGroup: %s", err)
	}
	if err := d.Set("cloud_function", flattenComputeRegionNetworkEndpointGroupCloudFunction(res["cloudFunction"], d, config)); err != nil {
		return fmt.Errorf("Error reading RegionNetworkEndpointGroup: %s", err)
	}
	if err := d.Set("region", flattenComputeRegionNetworkEndpointGroupRegion(res["region"], d, config)); err != nil {
		return fmt.Errorf("Error reading RegionNetworkEndpointGroup: %s", err)
	}
	if err := d.Set("self_link", ConvertSelfLinkToV1(res["selfLink"].(string))); err != nil {
		return fmt.Errorf("Error reading RegionNetworkEndpointGroup: %s", err)
	}

	return nil
}

func resourceComputeRegionNetworkEndpointGroupDelete(d *schema.ResourceData, meta interface{}) error {
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

	url, err := replaceVars(d, config, "{{ComputeBasePath}}projects/{{project}}/regions/{{region}}/networkEndpointGroups/{{name}}")
	if err != nil {
		return err
	}

	var obj map[string]interface{}
	log.Printf("[DEBUG] Deleting RegionNetworkEndpointGroup %q", d.Id())

	// err == nil indicates that the billing_project value was found
	if bp, err := getBillingProject(d, config); err == nil {
		billingProject = bp
	}

	res, err := sendRequestWithTimeout(config, "DELETE", billingProject, url, obj, d.Timeout(schema.TimeoutDelete))
	if err != nil {
		return handleNotFoundError(err, d, "RegionNetworkEndpointGroup")
	}

	err = computeOperationWaitTime(
		config, res, project, "Deleting RegionNetworkEndpointGroup",
		d.Timeout(schema.TimeoutDelete))

	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Finished deleting RegionNetworkEndpointGroup %q: %#v", d.Id(), res)
	return nil
}

func resourceComputeRegionNetworkEndpointGroupImport(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	config := meta.(*Config)
	if err := parseImportId([]string{
		"projects/(?P<project>[^/]+)/regions/(?P<region>[^/]+)/networkEndpointGroups/(?P<name>[^/]+)",
		"(?P<project>[^/]+)/(?P<region>[^/]+)/(?P<name>[^/]+)",
		"(?P<region>[^/]+)/(?P<name>[^/]+)",
		"(?P<name>[^/]+)",
	}, d, config); err != nil {
		return nil, err
	}

	// Replace import id for the resource id
	id, err := replaceVars(d, config, "projects/{{project}}/regions/{{region}}/networkEndpointGroups/{{name}}")
	if err != nil {
		return nil, fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	return []*schema.ResourceData{d}, nil
}

func flattenComputeRegionNetworkEndpointGroupName(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func flattenComputeRegionNetworkEndpointGroupDescription(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func flattenComputeRegionNetworkEndpointGroupNetworkEndpointType(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func flattenComputeRegionNetworkEndpointGroupCloudRun(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	if v == nil {
		return nil
	}
	original := v.(map[string]interface{})
	if len(original) == 0 {
		return nil
	}
	transformed := make(map[string]interface{})
	transformed["service"] =
		flattenComputeRegionNetworkEndpointGroupCloudRunService(original["service"], d, config)
	transformed["tag"] =
		flattenComputeRegionNetworkEndpointGroupCloudRunTag(original["tag"], d, config)
	transformed["url_mask"] =
		flattenComputeRegionNetworkEndpointGroupCloudRunUrlMask(original["urlMask"], d, config)
	return []interface{}{transformed}
}
func flattenComputeRegionNetworkEndpointGroupCloudRunService(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func flattenComputeRegionNetworkEndpointGroupCloudRunTag(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func flattenComputeRegionNetworkEndpointGroupCloudRunUrlMask(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func flattenComputeRegionNetworkEndpointGroupAppEngine(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	if v == nil {
		return nil
	}
	original := v.(map[string]interface{})
	transformed := make(map[string]interface{})
	transformed["service"] =
		flattenComputeRegionNetworkEndpointGroupAppEngineService(original["service"], d, config)
	transformed["version"] =
		flattenComputeRegionNetworkEndpointGroupAppEngineVersion(original["version"], d, config)
	transformed["url_mask"] =
		flattenComputeRegionNetworkEndpointGroupAppEngineUrlMask(original["urlMask"], d, config)
	return []interface{}{transformed}
}
func flattenComputeRegionNetworkEndpointGroupAppEngineService(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func flattenComputeRegionNetworkEndpointGroupAppEngineVersion(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func flattenComputeRegionNetworkEndpointGroupAppEngineUrlMask(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func flattenComputeRegionNetworkEndpointGroupCloudFunction(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	if v == nil {
		return nil
	}
	original := v.(map[string]interface{})
	if len(original) == 0 {
		return nil
	}
	transformed := make(map[string]interface{})
	transformed["function"] =
		flattenComputeRegionNetworkEndpointGroupCloudFunctionFunction(original["function"], d, config)
	transformed["url_mask"] =
		flattenComputeRegionNetworkEndpointGroupCloudFunctionUrlMask(original["urlMask"], d, config)
	return []interface{}{transformed}
}
func flattenComputeRegionNetworkEndpointGroupCloudFunctionFunction(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func flattenComputeRegionNetworkEndpointGroupCloudFunctionUrlMask(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func flattenComputeRegionNetworkEndpointGroupRegion(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	if v == nil {
		return v
	}
	return ConvertSelfLinkToV1(v.(string))
}

func expandComputeRegionNetworkEndpointGroupName(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandComputeRegionNetworkEndpointGroupDescription(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandComputeRegionNetworkEndpointGroupNetworkEndpointType(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandComputeRegionNetworkEndpointGroupCloudRun(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedService, err := expandComputeRegionNetworkEndpointGroupCloudRunService(original["service"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedService); val.IsValid() && !isEmptyValue(val) {
		transformed["service"] = transformedService
	}

	transformedTag, err := expandComputeRegionNetworkEndpointGroupCloudRunTag(original["tag"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedTag); val.IsValid() && !isEmptyValue(val) {
		transformed["tag"] = transformedTag
	}

	transformedUrlMask, err := expandComputeRegionNetworkEndpointGroupCloudRunUrlMask(original["url_mask"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedUrlMask); val.IsValid() && !isEmptyValue(val) {
		transformed["urlMask"] = transformedUrlMask
	}

	return transformed, nil
}

func expandComputeRegionNetworkEndpointGroupCloudRunService(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandComputeRegionNetworkEndpointGroupCloudRunTag(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandComputeRegionNetworkEndpointGroupCloudRunUrlMask(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandComputeRegionNetworkEndpointGroupAppEngine(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 {
		return nil, nil
	}

	if l[0] == nil {
		transformed := make(map[string]interface{})
		return transformed, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedService, err := expandComputeRegionNetworkEndpointGroupAppEngineService(original["service"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedService); val.IsValid() && !isEmptyValue(val) {
		transformed["service"] = transformedService
	}

	transformedVersion, err := expandComputeRegionNetworkEndpointGroupAppEngineVersion(original["version"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedVersion); val.IsValid() && !isEmptyValue(val) {
		transformed["version"] = transformedVersion
	}

	transformedUrlMask, err := expandComputeRegionNetworkEndpointGroupAppEngineUrlMask(original["url_mask"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedUrlMask); val.IsValid() && !isEmptyValue(val) {
		transformed["urlMask"] = transformedUrlMask
	}

	return transformed, nil
}

func expandComputeRegionNetworkEndpointGroupAppEngineService(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandComputeRegionNetworkEndpointGroupAppEngineVersion(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandComputeRegionNetworkEndpointGroupAppEngineUrlMask(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandComputeRegionNetworkEndpointGroupCloudFunction(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedFunction, err := expandComputeRegionNetworkEndpointGroupCloudFunctionFunction(original["function"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedFunction); val.IsValid() && !isEmptyValue(val) {
		transformed["function"] = transformedFunction
	}

	transformedUrlMask, err := expandComputeRegionNetworkEndpointGroupCloudFunctionUrlMask(original["url_mask"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedUrlMask); val.IsValid() && !isEmptyValue(val) {
		transformed["urlMask"] = transformedUrlMask
	}

	return transformed, nil
}

func expandComputeRegionNetworkEndpointGroupCloudFunctionFunction(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandComputeRegionNetworkEndpointGroupCloudFunctionUrlMask(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandComputeRegionNetworkEndpointGroupRegion(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	f, err := parseGlobalFieldValue("regions", v.(string), "project", d, config, true)
	if err != nil {
		return nil, fmt.Errorf("Invalid value for region: %s", err)
	}
	return f.RelativeLink(), nil
}
