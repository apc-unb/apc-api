# Projects

## Get all Projects from a student
* HTTP Request : ```GET http://api.com/project/{studentid}```
* Return a list of object in json format as follow

    ``` 
        [
			{
				"_id"           :	ObjectId,
				"studentid"     :   ObjectID,
				"projecttypeid" :   ObjectID,
                "monitorid"     :	ObjectId,
                "classid"       :	ObjectId,
                "createdat"     :	time.Time,
				"filename"      :	String,
                "status"        :   String
			}...
		]
    ```


## Update a project status
* HTTP Request : ```PUT http://api.com//project/status```
* Return a list of object in json format as follow

    ``` 
        {
            "_id"           :   ObjectId,
            "status"        :   String
        }
    ```


## Create Project
* HTTP Request : ```POST http://api.com/project```
* Send News's data in the request body in the follow format 

	``` 
        {
            "studentid"     :   ObjectID,
            "projecttypeid" :   ObjectID,
            "monitorid"     :	ObjectId,
            "classid"       :	ObjectId,
            "createdat"     :	time.Time,
            "filename"      :	String,
            "status"        :   String
        }
    ```

* http StatusCreated (201) will be sent if the project has been created correctly