package gsoc2

func GetWorkflowTest() []byte { 
	return []byte(`{"actions":[{"app_name":"Gsoc2 Tools","app_version":"1.2.0","description":"Set a value to be saved to your organization in Gsoc2.","app_id":"b53109ec-2873-4076-9826-4e7f586dc714","errors":[],"id":"c93c2ce0-e42a-4d30-8a2e-e9adb7ee7cc4","is_valid":true,"isStartNode":true,"sharing":true,"label":"Change Me","public":true,"generated":false,"large_image":"","environment":"Gsoc2","name":"set_cache_value","parameters":[{"description":"The key to set the value for","id":"","name":"key","example":"timestamp","value":"$onprem_dashboard_testing","multiline":false,"options":null,"action_field":"","variant":"STATIC_VALUE","required":true,"configuration":false,"tags":null,"schema":{"type":"string"},"skip_multicheck":false,"value_replace":null,"unique_toggled":false},{"description":"The value to set","id":"","name":"value","example":"1621959545","value":"192.168.2.3 https://google.com","multiline":true,"options":null,"action_field":"","variant":"STATIC_VALUE","required":true,"configuration":false,"tags":null,"schema":{"type":"string"},"skip_multicheck":false,"value_replace":null,"unique_toggled":false}],"execution_variable":{"description":"","id":"","name":"","value":""},"position":{"x":-142.20343154942202,"y":130.5567750670353},"authentication_id":"","category":"","reference_url":"","sub_action":false,"source_workflow":"","run_magic_output":false,"run_magic_input":false,"execution_delay":0,"category_label":null,"suggestion":false},{"app_name":"Gsoc2 Tools","app_version":"1.2.0","description":"Get a value saved to your organization in Gsoc2","app_id":"b53109ec-2873-4076-9826-4e7f586dc714","errors":[],"id":"f8a44502-e350-4180-933c-f7c3d7e8460b","is_valid":true,"sharing":true,"label":"Gsoc2_Tools_3","public":true,"generated":false,"large_image":"","environment":"Gsoc2","name":"get_cache_value","parameters":[{"description":"The key to get","id":"","name":"key","example":"timestamp","value":"$onprem_dashboard_testing","multiline":false,"options":null,"action_field":"","variant":"STATIC_VALUE","required":true,"configuration":false,"tags":null,"schema":{"type":"string"},"skip_multicheck":false,"value_replace":null,"unique_toggled":false}],"execution_variable":{"description":"","id":"","name":"","value":""},"position":{"x":-133.57704208335156,"y":308.69403928684073},"authentication_id":"","category":"Other","reference_url":"","sub_action":false,"source_workflow":"","run_magic_output":false,"run_magic_input":false,"execution_delay":0,"category_label":null,"suggestion":false},{"app_name":"Gsoc2 Tools","app_version":"1.2.0","description":"Delete a value saved to your organization in Gsoc2","app_id":"b53109ec-2873-4076-9826-4e7f586dc714","errors":[],"id":"240b5c73-72eb-4ff0-b177-1dbf5a3cb854","is_valid":true,"sharing":true,"label":"Gsoc2_Tools_3_copy","public":true,"generated":false,"large_image":"","environment":"Gsoc2","name":"delete_cache_value","parameters":[{"description":"The key to delete","id":"","name":"key","example":"timestamp","value":"$onprem_dashboard_testing","multiline":false,"options":null,"action_field":"","variant":"STATIC_VALUE","required":true,"configuration":false,"tags":null,"schema":{"type":"string"},"skip_multicheck":false,"value_replace":null,"unique_toggled":false}],"execution_variable":{"description":"","id":"","name":"","value":""},"position":{"x":-130.2282403427722,"y":480.74311435295436},"authentication_id":"","category":"Other","reference_url":"","sub_action":false,"source_workflow":"","run_magic_output":false,"run_magic_input":false,"execution_delay":0,"category_label":null,"suggestion":false}],"branches":[{"destination_id":"f8a44502-e350-4180-933c-f7c3d7e8460b","id":"73ed3768-a385-4e8d-a8f5-d7fba4f6ea7e","source_id":"c93c2ce0-e42a-4d30-8a2e-e9adb7ee7cc4","label":"","has_errors":false,"conditions":null,"decorator":false},{"destination_id":"240b5c73-72eb-4ff0-b177-1dbf5a3cb854","id":"34a30674-1325-41fd-90ab-f517cbeb6aa0","source_id":"f8a44502-e350-4180-933c-f7c3d7e8460b","label":"","has_errors":false,"conditions":null,"decorator":false}],"visual_branches":[],"triggers":[],"schedules":[],"comments":[],"configuration":{"exit_on_error":false,"start_from_top":false,"skip_notifications":false},"created":1692126563,"edited":1697725624,"last_runtime":0,"due_date":1697587200,"tags":["test"],"id":"1cd69f13-5f82-462c-b8e1-91a5fbac4746","is_valid":true,"name":"Workflow Testing","description":"Used for workflow testing in Gsoc2 Onprem","start":"c93c2ce0-e42a-4d30-8a2e-e9adb7ee7cc4","owner":"7cff070a-e888-4e27-a575-39769b6102c2","sharing":"private","image":"","execution_org":{"name":"default","id":"ba4d38b7-db3f-4908-9ccb-47ec03f2963e","users":[],"role":"","creator_org":"","image":"","child_orgs":null,"region_url":""},"org_id":"ba4d38b7-db3f-4908-9ccb-47ec03f2963e","workflow_variables":[{"description":"","id":"7e50be7e-e774-4f4b-ac44-34c5f786eb32","name":"onprem_dashboard_testing","value":"onprem_dashboard_testing"}],"execution_environment":"","previously_saved":true,"categories":{"siem":{"name":"","count":0,"id":"","description":"","large_image":""},"communication":{"name":"","count":0,"id":"","description":"","large_image":""},"assets":{"name":"","count":0,"id":"","description":"","large_image":""},"cases":{"name":"","count":0,"id":"","description":"","large_image":""},"network":{"name":"","count":0,"id":"","description":"","large_image":""},"intel":{"name":"","count":0,"id":"","description":"","large_image":""},"edr":{"name":"","count":0,"id":"","description":"","large_image":""},"iam":{"name":"","count":0,"id":"","description":"","large_image":""},"email":{"name":"","count":0,"id":"","description":"","large_image":""},"other":{"name":"","count":2,"id":"","description":"","large_image":""}},"example_argument":"","public":false,"default_return_value":"","contact_info":{"name":"","url":""},"published_id":"","revision_id":"","usecase_ids":["Email management"],"blogpost":"","video":"","status":"test","workflow_type":"standalone","generated":false,"hidden":false,"updated_by":"admin"}`)
}
